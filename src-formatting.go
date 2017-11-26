package z

import (
	"strings"
)

type iSrcFormatting interface {
	iCoreCmds

	KnownFormatters() []*SrcFormatterDesc
}

type SrcFormatterDesc struct {
	Name      string
	Link      string
	Installed bool
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
		MsgID: msgID_srcFmt_SetDef,
		Title: "Change Default Formatter",
		Desc:  strf("Specify your preferred default %s source formatter", Lang.Title),
		Hint:  "Current: (none)",
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
		if srcLoc.FilePath != "" || srcLoc.SrcFull != "" {
			cmds = append(cmds, me.cmdRunOnFile)
		}
		if srcLoc.SrcSel != "" {
			cmds = append(cmds, me.cmdRunOnSel)
		}
	}
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
			resp.WebsiteURL = kf.Link
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
		cmd.Desc = "➜ Open website at " + kf.Link
		if !kf.Installed {
			cmd.Hint = "Not installed"
		} else {
			cmd.Hint = "Installed"
		}
		m.Choices = append(m.Choices, &cmd)
	}
	resp.CoreCmdsMenu = &m
}
