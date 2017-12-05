package z

import (
	"github.com/metaleap/go-util/fs"
)

type iSrcMod interface {
	iCmdsProvider

	CodeActions(*SrcLens) []EditorAction
	DoesStdoutWithFilePathArg(*Tool) bool
	KnownFormatters() Tools
	RunRenamer(*SrcLens, string) []*SrcLens
	RunFormatter(*Tool, string, string, string) (string, string)
}

type SrcModBase struct {
	cmdFmtSetDef    *coreCmd
	cmdFmtRunOnFile *coreCmd
	cmdFmtRunOnSel  *coreCmd

	Impl iSrcMod
}

func (me *SrcModBase) Init() {
	me.cmdFmtSetDef = &coreCmd{
		MsgID: MSGID_SRCMOD_FMT_SETDEFMENU,
		Title: "Change Default Formatter",
		Desc:  Strf("Specify your preferred default %s source formatter", Lang.Title),
	}
	me.cmdFmtRunOnFile = &coreCmd{
		MsgID: MSGID_SRCMOD_FMT_RUNONFILE,
		Title: "Format Document",
	}
	me.cmdFmtRunOnSel = &coreCmd{
		MsgID: MSGID_SRCMOD_FMT_RUNONSEL,
		Title: "Format Selection",
	}
}

func (me *SrcModBase) Cmds(srcLens *SrcLens) (cmds []*coreCmd) {
	if srcLens != nil {
		desc := "(" + me.cmdFmtSetDef.Desc + " first)"
		if me.hasFormatter() {
			if desc = "➜ using "; me.isFormatterCustom() {
				desc += "'" + Prog.Cfg.FormatterProg + "' like "
			}
			desc += "'" + Prog.Cfg.FormatterName + "'"
		}

		if srcLens.FilePath != "" || srcLens.SrcFull != "" {
			me.cmdFmtRunOnFile.Desc = desc
			if me.cmdFmtRunOnFile.Hint = srcLens.FilePath; me.cmdFmtRunOnFile.Hint == "" {
				me.cmdFmtRunOnFile.Hint = srcLens.SrcFull
			}
			cmds = append(cmds, me.cmdFmtRunOnFile)
		}
		if srcLens.SrcSel != "" {
			me.cmdFmtRunOnSel.Desc = desc
			me.cmdFmtRunOnSel.Hint = srcLens.SrcSel
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

func (*SrcModBase) CmdsCategory() string {
	return "Formatting"
}

func (*SrcModBase) DoesStdoutWithFilePathArg(*Tool) bool {
	return true
}

func (*SrcModBase) CodeActions(srcLens *SrcLens) (all []EditorAction) {
	all = append(all, EditorAction{Title: "Open Zentient Menu", Cmd: "zen.core.cmds.listall", Hint: "Should open the main Palette menu"})
	return
}

func (*SrcModBase) RunRenamer(srcLens *SrcLens, newName string) (all []*SrcLens) {
	panic(Strf("Rename not yet implemented for __%s__.", Lang.Title))
}

func (*SrcModBase) hasFormatter() bool {
	return Prog.Cfg.FormatterName != ""
}

func (*SrcModBase) isFormatterCustom() bool {
	return Prog.Cfg.FormatterProg != "" && Prog.Cfg.FormatterProg != Prog.Cfg.FormatterName
}

func (me *SrcModBase) dispatch(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_SRCMOD_FMT_SETDEFMENU:
		me.onSetDefMenu(req, resp)
	case MSGID_SRCMOD_FMT_SETDEFPICK:
		me.onSetDefPick(req, resp)
	case MSGID_SRCMOD_FMT_RUNONFILE, MSGID_SRCMOD_FMT_RUNONSEL:
		me.onRunFormatter(req, resp)
	case MSGID_SRCMOD_RENAME:
		me.onRename(req, resp)
	case MSGID_SRCMOD_ACTIONS:
		me.onActions(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcModBase) onActions(req *msgReq, resp *msgResp) {
	resp.SrcActions = me.Impl.CodeActions(req.SrcLens)
}

func (me *SrcModBase) onRename(req *msgReq, resp *msgResp) {
	newname, _ := req.MsgArgs.(string)
	if newname == "" {
		resp.ErrMsg = "Rename: missing new-name"
	} else {
		resp.SrcMods = me.Impl.RunRenamer(req.SrcLens, newname)
	}
}

func (me *SrcModBase) onRunFormatter(req *msgReq, resp *msgResp) {
	var hasopt = false
	opt, _ := req.MsgArgs.(map[string]interface{})
	if opt != nil {
		_, tabSize := opt["tabSize"]
		_, insertSpaces := opt["insertSpaces"]
		hasopt = tabSize || insertSpaces
	}
	if !hasopt {
		resp.CoreCmd = &coreCmdResp{}
	}

	formatter := me.Impl.KnownFormatters().ByName(Prog.Cfg.FormatterName)
	if formatter == nil {
		if resp.CoreCmd == nil {
			resp.ErrMsg = "Select a Default Formatter first via the Zentient 'Palette' menu."
		} else {
			resp.CoreCmd.NoteWarn = "Select a Default Formatter first, either via the Zentient 'Palette' menu or:"
			resp.MsgID = MSGID_SRCMOD_FMT_SETDEFMENU
			resp.CoreCmd.MsgAction = Strf("Pick your preferred Zentient default %s formatter…", Lang.Title)
		}
		return
	}

	srcfilepath := req.SrcLens.FilePath
	withfilepathcmdarg := me.Impl.DoesStdoutWithFilePathArg(formatter)
	if !(ufs.FileExists(srcfilepath) && withfilepathcmdarg) {
		srcfilepath = ""
	}
	src := &req.SrcLens.SrcSel
	if *src == "" {
		src = &req.SrcLens.SrcFull
	}
	if (*src == "") && req.SrcLens.FilePath != "" && ufs.FileExists(req.SrcLens.FilePath) && !withfilepathcmdarg {
		req.SrcLens.ensureSrcFull()
		src = &req.SrcLens.SrcFull
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

	if srcformatted, stderr := me.Impl.RunFormatter(formatter, cmdname, srcfilepath, *src); stderr != "" {
		resp.ErrMsg = stderr
	} else {
		*src = srcformatted
		resp.SrcMods = []*SrcLens{req.SrcLens}
	}
}

func (me *SrcModBase) onSetDefMenu(req *msgReq, resp *msgResp) {
	m := coreCmdsMenu{Desc: "First pick a known formatter, then optionally specify a custom tool name:"}
	for _, kf := range me.Impl.KnownFormatters() {
		var cmd = coreCmd{Title: kf.Name, MsgID: MSGID_SRCMOD_FMT_SETDEFPICK}
		cmd.MsgArgs = map[string]interface{}{"fn": kf.Name, "fp": coreCmdMsgArgPrompt{Placeholder: kf.Name,
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
		m.Choices = append(m.Choices, &cmd)
	}
	resp.CoreCmd = &coreCmdResp{CoreCmdsMenu: &m}
}

func (me *SrcModBase) onSetDefPick(req *msgReq, resp *msgResp) {
	m := req.MsgArgs.(map[string]interface{})
	Prog.Cfg.FormatterName = m["fn"].(string)
	if Prog.Cfg.FormatterProg = m["fp"].(string); Prog.Cfg.FormatterProg == Prog.Cfg.FormatterName {
		Prog.Cfg.FormatterProg = ""
	}
	if err := Prog.Cfg.Save(); err != nil {
		resp.ErrMsg = err.Error()
	} else {
		resp.CoreCmd = &coreCmdResp{}
		resp.CoreCmd.NoteInfo = Strf("Default %s formatter changed to '%s'", Lang.Title, Prog.Cfg.FormatterName)
		if me.isFormatterCustom() {
			resp.CoreCmd.NoteInfo += Strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
		resp.CoreCmd.NoteInfo += "."
	}
}
