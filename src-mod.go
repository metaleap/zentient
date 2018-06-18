package z

import (
	"strings"

	"github.com/go-leap/fs"
)

type ISrcMod interface {
	IMenuItems

	CodeActions(*SrcLens) []EditorAction
	DoesStdoutWithFilePathArg(*Tool) bool
	KnownFormatters() Tools
	RunRenamer(*SrcLens, string) SrcLenses
	RunFormatter(*Tool, string, *SrcFormattingClientPrefs, string, string) (string, string)
}

type SrcModEdits []srcModEdit

func (this *SrcModEdits) dropConflictingEdits() (droppedOffenders []srcModEdit) {
	all := *this
	for i := 0; i < len(all); i++ {
		for disedit, j := all[i], i+1; j < len(all); j++ {
			if datedit := all[j]; disedit.At.overlapsWith(datedit.At) {
				droppedOffenders = append(droppedOffenders, all[j])
				pref, suff := all[:j], all[j+1:]
				j, all = j-1, append(pref, suff...)
			}
		}
	}
	*this = all
	return
}

func (this SrcModEdits) Len() int          { return len(this) }
func (this SrcModEdits) Swap(i int, j int) { this[i], this[j] = this[j], this[i] }
func (this SrcModEdits) Less(i int, j int) bool {
	return this[i].At.Start.isSameOrGreaterThan(&this[j].At.End)
}

func (*SrcModEdits) lensForNewEdit(srcFilePath string) *SrcLens {
	var lens = SrcLens{SrcLoc: SrcLoc{FilePath: srcFilePath}}
	lens.EnsureSrcFull()
	return &lens
}

func (this *SrcModEdits) AddDeleteLine(srcFilePath string, lineAt *SrcPos) {
	lens := this.lensForNewEdit(srcFilePath)
	lens.Pos = lineAt
	edit := srcModEdit{At: &SrcRange{}}
	bo := lens.ByteOffsetForPos(lens.Pos)
	bo = strings.LastIndex(lens.Txt[:bo], "\n") + 1
	edit.At.Start.Off = lens.Rune1OffsetForByte0Offset(bo)
	bo2 := strings.IndexRune(lens.Txt[bo:], '\n') + 1
	edit.At.End.Off = lens.Rune1OffsetForByte0Offset(bo + bo2)
	*this = append(*this, edit)
}

func (this *SrcModEdits) AddInsert(srcFilePath string, atPos func(*SrcLens, *SrcPos) string) {
	lens := this.lensForNewEdit(srcFilePath)
	edit := srcModEdit{At: &SrcRange{}}
	if ins := atPos(lens, &edit.At.Start); ins != "" {
		edit.Val = ins
		*this = append(*this, edit)
	}
}

type srcModEdit struct {
	At  *SrcRange
	Val string // if not empty: inserts if At is pos, replaces if At is range. if empty: deletes if At range is range, errors if At is pos.
}

type SrcModBase struct {
	cmdFmtSetDef    *MenuItem
	cmdFmtRunOnFile *MenuItem
	cmdFmtRunOnSel  *MenuItem

	Impl ISrcMod
}

type SrcFormattingClientPrefs struct {
	InsertSpaces *bool
	TabSize      *int
}

func (this *SrcModBase) Init() {
	this.cmdFmtSetDef = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_SETDEFMENU,
		Title: Strf("Choose Default %s Formatter", Lang.Title),
		Desc:  "Specify your preferred default source formatter",
	}
	this.cmdFmtRunOnFile = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_RUNONFILE,
		Title: "Format Document",
	}
	this.cmdFmtRunOnSel = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_RUNONSEL,
		Title: "Format Selection",
	}
}

func (this *SrcModBase) menuItems(srcLens *SrcLens) (cmds MenuItems) {
	if srcLens != nil {
		srcfilepath, hint := srcLens.FilePath, "("+this.cmdFmtSetDef.Desc+" first)"
		if this.hasFormatter() {
			if hint = "➜ using "; this.isFormatterCustom() {
				hint += "'" + Prog.Cfg.FormatterProg + "' like "
			}
			hint += "'" + Prog.Cfg.FormatterName + "'"
		}

		if isfp := srcfilepath != ""; isfp || srcLens.Txt != "" {
			srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
			if this.cmdFmtRunOnFile.Desc, this.cmdFmtRunOnFile.Hint = srcfilepath, hint; !isfp {
				this.cmdFmtRunOnFile.Desc = srcLens.Txt
			}
			cmds = append(cmds, this.cmdFmtRunOnFile)
		}
		if srcLens.Str != "" {
			this.cmdFmtRunOnSel.Desc = srcLens.Str
			this.cmdFmtRunOnSel.Hint = hint
			cmds = append(cmds, this.cmdFmtRunOnSel)
		}
	}

	if this.cmdFmtSetDef.Hint = "(none)"; this.hasFormatter() {
		if this.cmdFmtSetDef.Hint = "'" + Prog.Cfg.FormatterName + "'"; this.isFormatterCustom() {
			this.cmdFmtSetDef.Hint += "-compatible '" + Prog.Cfg.FormatterProg + "'"
		}
	}
	this.cmdFmtSetDef.Hint = "Current: " + this.cmdFmtSetDef.Hint
	cmds = append(cmds, this.cmdFmtSetDef)
	return
}

func (*SrcModBase) MenuCategory() string {
	return "Formatting"
}

func (*SrcModBase) CodeActions(srcLens *SrcLens) (all []EditorAction) {
	return
}

func (*SrcModBase) RunRenamer(srcLens *SrcLens, newName string) (all SrcLenses) {
	panic(Strf("Rename not yet implemented for %s.", Lang.Title))
}

func (*SrcModBase) hasFormatter() bool {
	return Prog.Cfg.FormatterName != ""
}

func (*SrcModBase) isFormatterCustom() bool {
	return Prog.Cfg.FormatterProg != "" && Prog.Cfg.FormatterProg != Prog.Cfg.FormatterName
}

func (this *SrcModBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCMOD_FMT_SETDEFMENU:
		this.onSetDefMenu(req, resp)
	case IPCID_SRCMOD_FMT_SETDEFPICK:
		this.onSetDefPick(req, resp)
	case IPCID_SRCMOD_FMT_RUNONFILE, IPCID_SRCMOD_FMT_RUNONSEL:
		this.onRunFormatter(req, resp)
	case IPCID_SRCMOD_RENAME:
		this.onRename(req, resp)
	case IPCID_SRCMOD_ACTIONS:
		this.onActions(req, resp)
	default:
		return false
	}
	return true
}

func (this *SrcModBase) onActions(req *ipcReq, resp *ipcResp) {
	resp.SrcActions = this.Impl.CodeActions(req.SrcLens)
}

func (this *SrcModBase) onRename(req *ipcReq, resp *ipcResp) {
	newname, _ := req.IpcArgs.(string)
	if newname == "" {
		resp.ErrMsg = "Rename: missing new-name"
	} else {
		resp.SrcMods = this.Impl.RunRenamer(req.SrcLens, newname)
	}
}

func (this *SrcModBase) onRunFormatter(req *ipcReq, resp *ipcResp) {
	optraw, _ := req.IpcArgs.(map[string]interface{})
	var prefs *SrcFormattingClientPrefs
	if optraw != nil {
		tabSize, ok1 := optraw["tabSize"].(float64)
		insertSpaces, ok2 := optraw["insertSpaces"].(bool)
		if ok1 || ok2 {
			if prefs = (&SrcFormattingClientPrefs{}); ok2 {
				prefs.InsertSpaces = &insertSpaces
			}
			if tabsize := int(tabSize); ok1 {
				prefs.TabSize = &tabsize
			}
		}
	} else {
		resp.Menu = &menuResp{}
	}

	formatter := this.Impl.KnownFormatters().byName(Prog.Cfg.FormatterName)
	if formatter == nil {
		if resp.Menu == nil {
			resp.ErrMsg = "Select a Default Formatter first via the Zentient Main Menu."
		} else {
			resp.Menu.NoteWarn = "Select a Default Formatter first, either via the Zentient Main Menu or:"
			resp.IpcID = IPCID_SRCMOD_FMT_SETDEFMENU
			resp.Menu.UxActionLabel = Strf("Pick your preferred Zentient default %s formatter…", Lang.Title)
		}
		return
	}

	srcfilepath := req.SrcLens.FilePath
	withfilepathcmdarg := this.Impl.DoesStdoutWithFilePathArg(formatter)
	if !(ufs.IsFile(srcfilepath) && withfilepathcmdarg) {
		srcfilepath = ""
	}
	src := &req.SrcLens.Str
	if *src == "" {
		src = &req.SrcLens.Txt
	}
	if (*src == "") && req.SrcLens.FilePath != "" && ufs.IsFile(req.SrcLens.FilePath) && !withfilepathcmdarg {
		req.SrcLens.EnsureSrcFull()
		src = &req.SrcLens.Txt
	}

	if *src != "" {
		srcfilepath = ""
	} else if srcfilepath == "" {
		resp.ErrMsg = "Nothing to format?!"
		return
	}

	cmdname := formatter.Name
	if Prog.Cfg.FormatterProg != "" {
		cmdname = Prog.Cfg.FormatterProg
	}

	if srcformatted, stderr := this.Impl.RunFormatter(formatter, cmdname, prefs, srcfilepath, *src); srcformatted != "" {
		*src, resp.SrcMods = srcformatted, SrcLenses{req.SrcLens}
	} else if stderr != "" {
		resp.ErrMsg = stderr
	}
}

func (this *SrcModBase) onSetDefMenu(req *ipcReq, resp *ipcResp) {
	m := Menu{Desc: "First pick a known formatter, then optionally specify a custom tool name:"}
	for _, kf := range this.Impl.KnownFormatters() {
		var cmd = MenuItem{Title: kf.Name, IpcID: IPCID_SRCMOD_FMT_SETDEFPICK}
		cmd.IpcArgs = map[string]interface{}{"fn": kf.Name, "fp": menuItemIpcArgPrompt{Placeholder: kf.Name,
			Prompt: Strf("Optionally enter the name of an alternative '%s'-compatible equivalent tool to use", kf.Name)}}
		cmd.Desc = Strf("➜ Pick to use '%s' (or compatible equivalent) as the default %s formatter", kf.Name, Lang.Title)
		if kf.Name != Prog.Cfg.FormatterName || !this.isFormatterCustom() {
			if cmd.Hint = "· Installed "; !kf.Installed {
				cmd.Hint = "· Not Installed "
			}
		}
		if kf.Name == Prog.Cfg.FormatterName {
			if cmd.Hint += "· Current Default "; this.isFormatterCustom() {
				cmd.Hint += "— Using '" + Prog.Cfg.FormatterProg + "' "
			}
		}
		cmd.Hint += "· " + kf.Website
		m.Items = append(m.Items, &cmd)
	}
	resp.Menu = &menuResp{SubMenu: &m}
}

func (this *SrcModBase) onSetDefPick(req *ipcReq, resp *ipcResp) {
	m := req.IpcArgs.(map[string]interface{})
	Prog.Cfg.FormatterName = m["fn"].(string)
	if Prog.Cfg.FormatterProg = m["fp"].(string); Prog.Cfg.FormatterProg == Prog.Cfg.FormatterName {
		Prog.Cfg.FormatterProg = ""
	}
	if err := Prog.Cfg.Save(); err != nil {
		resp.ErrMsg = err.Error()
	} else {
		resp.Menu = &menuResp{}
		resp.Menu.NoteInfo = Strf("Default %s formatter changed to '%s'", Lang.Title, Prog.Cfg.FormatterName)
		if this.isFormatterCustom() {
			resp.Menu.NoteInfo += Strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
		resp.Menu.NoteInfo += "."
	}
}
