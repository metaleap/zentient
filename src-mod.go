package z

import (
	"github.com/metaleap/go-util/fs"
)

type ISrcMod interface {
	IMenuItems

	CodeActions(*SrcLens) []EditorAction
	DoesStdoutWithFilePathArg(*Tool) bool
	KnownFormatters() Tools
	RunRenamer(*SrcLens, string) SrcLenses
	RunFormatter(*Tool, string, *SrcFormattingClientPrefs, string, string) (string, string)
}

type SrcModEdits []SrcModEdit

func (me *SrcModEdits) DropConflictingEdits() (droppedOffenders []SrcModEdit) {
	all, mod := *me, false
	for again := true; again; {
		again = false
		for i, disedit := range all {
			for j, datedit := range all {
				if i != j && disedit.At.OverlapsWith(datedit.At) {
					droppedOffenders = append(droppedOffenders, all[i])
					pref, suff := all[:i], all[i+1:]
					again, mod, all = true, true, append(pref, suff...)
					break
				}
			}
			if again {
				break
			}
		}
	}
	if mod {
		*me = all
	}
	return
}

func (me SrcModEdits) Len() int               { return len(me) }
func (me SrcModEdits) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me SrcModEdits) Less(i int, j int) bool { return me[i].At.Start.ComesBehind(&me[j].At.End) }

func (me *SrcModEdits) AddEdit_DeleteLine(srcFilePath string, lineAt *SrcPos) {
	var lens = SrcLens{SrcLoc: SrcLoc{FilePath: srcFilePath, Pos: lineAt}}
	lens.EnsureSrcFull()
	edit := SrcModEdit{}
	*me = append(*me, edit)
}

type SrcModEdit struct {
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

func (me *SrcModBase) Init() {
	me.cmdFmtSetDef = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_SETDEFMENU,
		Title: Strf("Choose Default %s Formatter", Lang.Title),
		Desc:  "Specify your preferred default source formatter",
	}
	me.cmdFmtRunOnFile = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_RUNONFILE,
		Title: "Format Document",
	}
	me.cmdFmtRunOnSel = &MenuItem{
		IpcID: IPCID_SRCMOD_FMT_RUNONSEL,
		Title: "Format Selection",
	}
}

func (me *SrcModBase) MenuItems(srcLens *SrcLens) (cmds MenuItems) {
	if srcLens != nil {
		srcfilepath, hint := srcLens.FilePath, "("+me.cmdFmtSetDef.Desc+" first)"
		if me.hasFormatter() {
			if hint = "➜ using "; me.isFormatterCustom() {
				hint += "'" + Prog.Cfg.FormatterProg + "' like "
			}
			hint += "'" + Prog.Cfg.FormatterName + "'"
		}

		if isfp := srcfilepath != ""; isfp || srcLens.Txt != "" {
			srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
			if me.cmdFmtRunOnFile.Desc, me.cmdFmtRunOnFile.Hint = srcfilepath, hint; !isfp {
				me.cmdFmtRunOnFile.Desc = srcLens.Txt
			}
			cmds = append(cmds, me.cmdFmtRunOnFile)
		}
		if srcLens.Str != "" {
			me.cmdFmtRunOnSel.Desc = srcLens.Str
			me.cmdFmtRunOnSel.Hint = hint
			cmds = append(cmds, me.cmdFmtRunOnSel)
		}
	}

	if me.cmdFmtSetDef.Hint = "(none)"; me.hasFormatter() {
		if me.cmdFmtSetDef.Hint = "'" + Prog.Cfg.FormatterName + "'"; me.isFormatterCustom() {
			me.cmdFmtSetDef.Hint += "-compatible '" + Prog.Cfg.FormatterProg + "'"
		}
	}
	me.cmdFmtSetDef.Hint = "Current: " + me.cmdFmtSetDef.Hint
	cmds = append(cmds, me.cmdFmtSetDef)
	return
}

func (*SrcModBase) MenuCategory() string {
	return "Formatting"
}

func (*SrcModBase) CodeActions(srcLens *SrcLens) (all []EditorAction) {
	return
}

func (*SrcModBase) RunRenamer(srcLens *SrcLens, newName string) (all SrcLenses) {
	panic(Strf("Rename not yet implemented for __%s__.", Lang.Title))
}

func (*SrcModBase) hasFormatter() bool {
	return Prog.Cfg.FormatterName != ""
}

func (*SrcModBase) isFormatterCustom() bool {
	return Prog.Cfg.FormatterProg != "" && Prog.Cfg.FormatterProg != Prog.Cfg.FormatterName
}

func (me *SrcModBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCMOD_FMT_SETDEFMENU:
		me.onSetDefMenu(req, resp)
	case IPCID_SRCMOD_FMT_SETDEFPICK:
		me.onSetDefPick(req, resp)
	case IPCID_SRCMOD_FMT_RUNONFILE, IPCID_SRCMOD_FMT_RUNONSEL:
		me.onRunFormatter(req, resp)
	case IPCID_SRCMOD_RENAME:
		me.onRename(req, resp)
	case IPCID_SRCMOD_ACTIONS:
		me.onActions(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcModBase) onActions(req *ipcReq, resp *ipcResp) {
	resp.SrcActions = me.Impl.CodeActions(req.SrcLens)
}

func (me *SrcModBase) onRename(req *ipcReq, resp *ipcResp) {
	newname, _ := req.IpcArgs.(string)
	if newname == "" {
		resp.ErrMsg = "Rename: missing new-name"
	} else {
		resp.SrcMods = me.Impl.RunRenamer(req.SrcLens, newname)
	}
}

func (me *SrcModBase) onRunFormatter(req *ipcReq, resp *ipcResp) {
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
		resp.Menu = &MenuResp{}
	}

	formatter := me.Impl.KnownFormatters().ByName(Prog.Cfg.FormatterName)
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
	withfilepathcmdarg := me.Impl.DoesStdoutWithFilePathArg(formatter)
	if !(ufs.FileExists(srcfilepath) && withfilepathcmdarg) {
		srcfilepath = ""
	}
	src := &req.SrcLens.Str
	if *src == "" {
		src = &req.SrcLens.Txt
	}
	if (*src == "") && req.SrcLens.FilePath != "" && ufs.FileExists(req.SrcLens.FilePath) && !withfilepathcmdarg {
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

	if srcformatted, stderr := me.Impl.RunFormatter(formatter, cmdname, prefs, srcfilepath, *src); srcformatted != "" {
		*src, resp.SrcMods = srcformatted, SrcLenses{req.SrcLens}
	} else if stderr != "" {
		resp.ErrMsg = stderr
	}
}

func (me *SrcModBase) onSetDefMenu(req *ipcReq, resp *ipcResp) {
	m := Menu{Desc: "First pick a known formatter, then optionally specify a custom tool name:"}
	for _, kf := range me.Impl.KnownFormatters() {
		var cmd = MenuItem{Title: kf.Name, IpcID: IPCID_SRCMOD_FMT_SETDEFPICK}
		cmd.IpcArgs = map[string]interface{}{"fn": kf.Name, "fp": MenuItemIpcArgPrompt{Placeholder: kf.Name,
			Prompt: Strf("Optionally enter the name of an alternative '%s'-compatible equivalent tool to use", kf.Name)}}
		cmd.Desc = Strf("➜ Pick to use '%s' (or compatible equivalent) as the default %s formatter", kf.Name, Lang.Title)
		if kf.Name != Prog.Cfg.FormatterName || !me.isFormatterCustom() {
			if cmd.Hint = "· Installed "; !kf.Installed {
				cmd.Hint = "· Not Installed "
			}
		}
		if kf.Name == Prog.Cfg.FormatterName {
			if cmd.Hint += "· Current Default "; me.isFormatterCustom() {
				cmd.Hint += "— Using '" + Prog.Cfg.FormatterProg + "' "
			}
		}
		cmd.Hint += "· " + kf.Website
		m.Items = append(m.Items, &cmd)
	}
	resp.Menu = &MenuResp{SubMenu: &m}
}

func (me *SrcModBase) onSetDefPick(req *ipcReq, resp *ipcResp) {
	m := req.IpcArgs.(map[string]interface{})
	Prog.Cfg.FormatterName = m["fn"].(string)
	if Prog.Cfg.FormatterProg = m["fp"].(string); Prog.Cfg.FormatterProg == Prog.Cfg.FormatterName {
		Prog.Cfg.FormatterProg = ""
	}
	if err := Prog.Cfg.Save(); err != nil {
		resp.ErrMsg = err.Error()
	} else {
		resp.Menu = &MenuResp{}
		resp.Menu.NoteInfo = Strf("Default %s formatter changed to '%s'", Lang.Title, Prog.Cfg.FormatterName)
		if me.isFormatterCustom() {
			resp.Menu.NoteInfo += Strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
		resp.Menu.NoteInfo += "."
	}
}
