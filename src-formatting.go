package z

import (
	"strings"
)

type iSrcFormatting interface {
	iCoreCmds

	KnownFormatters() []*Tool
}

type SrcFormattingBase struct {
	cmdListAll   *coreCmd
	cmdSetDef    *coreCmd
	cmdRunOnFile *coreCmd
	cmdRunOnSel  *coreCmd

	Self iSrcFormatting
}

func (me *SrcFormattingBase) Init() {
	me.cmdListAll = &coreCmd{
		MsgID: msgID_srcFmt_ListAll,
		Title: "List All Known Formatters",
		Desc:  "Lists all known " + Lang.Title + " formatters and their installation info / status",
	}
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
	if me.cmdListAll.Hint == "" {
		kfnames := []string{}
		for _, kf := range me.Self.KnownFormatters() {
			kfnames = append(kfnames, kf.Name)
		}
		me.cmdListAll.Hint = strings.Join(kfnames, " · ")
	}

	if srcLoc != nil {
		desc := "(" + me.cmdSetDef.Desc + " first)"
		if Prog.Cfg.FormatterName != "" {
			if desc = "➜ using "; Prog.Cfg.isFormatterCustom() {
				desc += "'" + Prog.Cfg.FormatterProg + "' in place of "
			}
			desc += "'" + Prog.Cfg.FormatterName + "'"
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
	if me.cmdSetDef.Hint = "(none)"; Prog.Cfg.FormatterName != "" {
		if me.cmdSetDef.Hint = "'" + Prog.Cfg.FormatterName + "'"; Prog.Cfg.isFormatterCustom() {
			me.cmdSetDef.Hint += "-compatible '" + Prog.Cfg.FormatterProg + "'"
		}
	}
	me.cmdSetDef.Hint = "Current: " + me.cmdSetDef.Hint
	cmds = append(cmds, me.cmdSetDef, me.cmdListAll)
	return
}

func (me *SrcFormattingBase) CmdsCategory() string {
	return "Formatting"
}

func (me *SrcFormattingBase) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_srcFmt_ListAll:
		me.handle_ListAll(req, resp)
	case msgID_srcFmt_InfoLink:
		me.handle_InfoLink(req, resp)
	case msgID_srcFmt_SetDefMenu:
		me.handle_SetDefMenu(req, resp)
	case msgID_srcFmt_SetDefPick:
		me.handle_SetDefPick(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcFormattingBase) handle_InfoLink(req *msgReq, resp *msgResp) {
	fmtname := req.MsgArgs.(string)
	for _, kf := range me.Self.KnownFormatters() {
		if kf.Name == fmtname {
			if !kf.Installed {
				resp.Note = strf("After installing '%s', reload Zentient to recognize it.", kf.Name)
			}
			resp.WebsiteURL = kf.Website
			return
		}
	}
	Bad("formatter", fmtname)
}

func (me *SrcFormattingBase) handle_ListAll(req *msgReq, resp *msgResp) {
	cfmt := Lang.SrcFmt // need the interface impl, not the embedded base!
	m := coreCmdsMenu{Desc: strf("❬%s❭ · %s:", cfmt.CmdsCategory(), me.cmdListAll.Title)}
	for _, kf := range cfmt.KnownFormatters() {
		var cmd = coreCmd{Title: kf.Name, MsgArgs: kf.Name, MsgID: msgID_srcFmt_InfoLink}
		cmd.Desc = "➜ Open website at ⟨ " + kf.Website + " ⟩"
		if !kf.Installed {
			cmd.Desc += " for installation help"
			cmd.Hint = "Not installed"
		} else {
			cmd.Hint = "Installed"
		}
		if kf.Name == Prog.Cfg.FormatterName {
			if cmd.Hint += " · Set as Default"; Prog.Cfg.isFormatterCustom() {
				cmd.Hint += " (via '" + Prog.Cfg.FormatterProg + "')"
			}
		}
		m.Choices = append(m.Choices, &cmd)
	}
	resp.CoreCmdsMenu = &m
}

func (me *SrcFormattingBase) handle_SetDefMenu(req *msgReq, resp *msgResp) {
	me.handle_ListAll(req, resp)
	resp.CoreCmdsMenu.Desc = "First pick a known formatter, then optionally specify a custom tool name:"
	for _, cmd := range resp.CoreCmdsMenu.Choices {
		toolname := cmd.MsgArgs.(string)
		cmd.Hint = strf(" — or specify a custom alternative (but '%s'-compatible) equivalent tool next", toolname)
		cmd.Desc = strf("➜ Pick to use '%s' (or compatible equivalent) as the default %s formatter", toolname, Lang.Title)
		cmd.MsgID = msgID_srcFmt_SetDefPick
		cmd.MsgArgs = map[string]interface{}{
			"fn": toolname,
			"fp": msgArgPrompt{Prompt: strf("Optionally enter the name of an alternative '%s'-compatible equivalent tool to use", toolname), Placeholder: toolname},
		}
	}
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
		if Prog.Cfg.FormatterProg != "" {
			resp.Note += strf("-compatible equivalent '%s'", Prog.Cfg.FormatterProg)
		}
	}
}
