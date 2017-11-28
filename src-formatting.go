package z

type iSrcFormatting interface {
	iCoreCmds

	KnownFormatters() []*Tool
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
		Desc:  strf("Specify your preferred default %s source formatter", Lang.Title),
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

func (me *SrcFormattingBase) Cmds(srcLoc *SrcLoc) (cmds []*coreCmd) {
	if srcLoc != nil {
		desc := "(" + me.cmdSetDef.Desc + " first)"
		if me.hasFormatter() {
			if desc = "➜ using "; me.isFormatterCustom() {
				desc += "'" + Prog.Cfg.FormatterProg + "' as a "
			}
			desc += "'" + Prog.Cfg.FormatterName + "' equivalent"
		}

		if srcLoc.FilePath != "" || srcLoc.SrcFull != "" {
			me.cmdRunOnFile.Desc = desc
			if me.cmdRunOnFile.Hint = srcLoc.FilePath; me.cmdRunOnFile.Hint == "" {
				me.cmdRunOnFile.Hint = srcLoc.SrcFull
			}
			cmds = append(cmds, me.cmdRunOnFile)
		}
		if srcLoc.SrcSel != "" {
			me.cmdRunOnSel.Desc = desc
			me.cmdRunOnSel.Hint = srcLoc.SrcSel
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
	default:
		return false
	}
	return true
}

func (me *SrcFormattingBase) handle_SetDefMenu(req *msgReq, resp *msgResp) {
	cfmt := Lang.SrcFmt // need the interface impl, not the embedded base!
	m := coreCmdsMenu{Desc: "First pick a known formatter, then optionally specify a custom tool name:"}
	for _, kf := range cfmt.KnownFormatters() {
		var cmd = coreCmd{Title: kf.Name, MsgID: msgID_srcFmt_SetDefPick}
		cmd.MsgArgs = map[string]interface{}{"fn": kf.Name, "fp": msgArgPrompt{Placeholder: kf.Name,
			Prompt: strf("Optionally enter the name of an alternative '%s'-compatible equivalent tool to use", kf.Name)}}
		cmd.Desc = strf("➜ Pick to use '%s' (or compatible equivalent) as the default %s formatter", kf.Name, Lang.Title)
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
		resp.Note = strf("Default %s formatter changed to '%s'", Lang.Title, Prog.Cfg.FormatterName)
		if me.isFormatterCustom() {
			resp.Note += strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
		resp.Note += "."
	}
}

func (*SrcFormattingBase) hasFormatter() bool {
	return Prog.Cfg.FormatterName != ""
}
func (*SrcFormattingBase) isFormatterCustom() bool {
	return Prog.Cfg.FormatterProg != "" && Prog.Cfg.FormatterProg != Prog.Cfg.FormatterName
}
