package z

import (
	"github.com/metaleap/go-util/fs"
)

type iSrcFormatting interface {
	iCoreCmds

	KnownFormatters() Tools
	RunFormatter(*Tool, string, string, string) (string, string, error)
}

type SrcFormattingBase struct {
	cmdSetDef    *coreCmd
	cmdRunOnFile *coreCmd
	cmdRunOnSel  *coreCmd

	Self iSrcFormatting
}

func (me *SrcFormattingBase) Init() {
	me.cmdSetDef = &coreCmd{
		MsgID: msgID_srcFmt_SetDefMenu,
		Title: "Change Default Formatter",
		Desc:  Strf("Specify your preferred default %s source formatter", Lang.Title),
	}
	me.cmdRunOnFile = &coreCmd{
		MsgID: msgID_srcFmt_RunOnFile,
		Title: "Format Document",
	}
	me.cmdRunOnSel = &coreCmd{
		MsgID: msgID_srcFmt_RunOnSel,
		Title: "Format Selection",
	}
}

func (me *SrcFormattingBase) Cmds(srcLens *SrcLens) (cmds []*coreCmd) {
	if srcLens != nil {
		desc := "(" + me.cmdSetDef.Desc + " first)"
		if me.hasFormatter() {
			if desc = "➜ using "; me.isFormatterCustom() {
				desc += "'" + Prog.Cfg.FormatterProg + "' like "
			}
			desc += "'" + Prog.Cfg.FormatterName + "'"
		}

		if srcLens.FilePath != "" || srcLens.SrcFull != "" {
			me.cmdRunOnFile.Desc = desc
			if me.cmdRunOnFile.Hint = srcLens.FilePath; me.cmdRunOnFile.Hint == "" {
				me.cmdRunOnFile.Hint = srcLens.SrcFull
			}
			cmds = append(cmds, me.cmdRunOnFile)
		}
		if srcLens.SrcSel != "" {
			me.cmdRunOnSel.Desc = desc
			me.cmdRunOnSel.Hint = srcLens.SrcSel
			cmds = append(cmds, me.cmdRunOnSel)
		}
	}

	if me.cmdSetDef.Hint = "(none)"; me.hasFormatter() {
		if me.cmdSetDef.Hint = "'" + Prog.Cfg.FormatterName + "'"; me.isFormatterCustom() {
			me.cmdSetDef.Hint += "-compatible '" + Prog.Cfg.FormatterProg + "'"
		}
	}
	me.cmdSetDef.Hint = "Current: " + me.cmdSetDef.Hint
	cmds = append(cmds, me.cmdSetDef)
	return
}

func (me *SrcFormattingBase) CmdsCategory() string {
	return "Formatting"
}

func (me *SrcFormattingBase) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_srcFmt_SetDefMenu:
		me.handle_SetDefMenu(req, resp)
	case msgID_srcFmt_SetDefPick:
		me.handle_SetDefPick(req, resp)
	case msgID_srcFmt_RunOnFile, msgID_srcFmt_RunOnSel:
		me.handle_RunFormatter(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcFormattingBase) handle_RunFormatter(req *msgReq, resp *msgResp) {
	self := me.Self
	formatter := self.KnownFormatters().ByName(Prog.Cfg.FormatterName)
	if formatter == nil {
		resp.NoteWarn = "Select a Default Formatter first, either via the Zentient 'Palette' menu or:"
		resp.MsgID = msgID_srcFmt_SetDefMenu
		resp.MsgAction = Strf("Pick your preferred Zentient default %s formatter…", Lang.Title)
		return
	}

	srcfilepath := req.SrcLens.FilePath
	if !ufs.FileExists(srcfilepath) {
		srcfilepath = ""
	}

	src := &req.SrcLens.SrcSel
	if *src == "" {
		src = &req.SrcLens.SrcFull
	}
	// won't use this until we encounter a formatter that doesn't itself support file-path arguments:
	// if *src == "" && req.SrcLens.FilePath != "" && ufs.FileExists(req.SrcLens.FilePath) {
	// 	req.SrcLens.SrcFull = ufs.ReadTextFile(req.SrcLens.FilePath, true, "")
	// 	src = &req.SrcLens.SrcFull
	// }
	if *src != "" {
		srcfilepath = ""
	} else if srcfilepath == "" {
		resp.ErrMsg = "Nothing to format?!"
		return
	}

	if srcformatted, stderr, err := self.RunFormatter(formatter, Prog.Cfg.FormatterProg, srcfilepath, *src); err != nil {
		resp.ErrMsg = err.Error()
	} else if stderr != "" {
		resp.ErrMsg = stderr
		resp.ErrMsgFromTool = true
	} else {
		*src = srcformatted
		resp.SrcMod = req.SrcLens
	}
}

func (me *SrcFormattingBase) handle_SetDefMenu(req *msgReq, resp *msgResp) {
	m := coreCmdsMenu{Desc: "First pick a known formatter, then optionally specify a custom tool name:"}
	for _, kf := range me.Self.KnownFormatters() {
		var cmd = coreCmd{Title: kf.Name, MsgID: msgID_srcFmt_SetDefPick}
		cmd.MsgArgs = map[string]interface{}{"fn": kf.Name, "fp": msgArgPrompt{Placeholder: kf.Name,
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
	resp.CoreCmdsMenu = &m
}

func (me *SrcFormattingBase) handle_SetDefPick(req *msgReq, resp *msgResp) {
	m := req.MsgArgs.(map[string]interface{})
	Prog.Cfg.FormatterName = m["fn"].(string)
	if Prog.Cfg.FormatterProg = m["fp"].(string); Prog.Cfg.FormatterProg == Prog.Cfg.FormatterName {
		Prog.Cfg.FormatterProg = ""
	}
	if err := Prog.Cfg.Save(); err != nil {
		resp.ErrMsg = err.Error()
	} else {
		resp.NoteInfo = Strf("Default %s formatter changed to '%s'", Lang.Title, Prog.Cfg.FormatterName)
		if me.isFormatterCustom() {
			resp.NoteInfo += Strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
		resp.NoteInfo += "."
	}
}

func (*SrcFormattingBase) hasFormatter() bool {
	return Prog.Cfg.FormatterName != ""
}
func (*SrcFormattingBase) isFormatterCustom() bool {
	return Prog.Cfg.FormatterProg != "" && Prog.Cfg.FormatterProg != Prog.Cfg.FormatterName
}
