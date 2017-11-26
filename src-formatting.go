package z

import (
	"strings"
)

type iSrcFormatting interface {
	iCoreCmds

	KnownFormatters() []*SrcFormatterDesc
}

type SrcFormattingBase struct {
	cmdListAll *coreCmd

	Self iSrcFormatting
}

type SrcFormatterDesc struct {
	Name      string
	Link      string
	Installed bool
}

func (me *SrcFormattingBase) Init() {
	me.cmdListAll = &coreCmd{
		MsgID: msgID_srcFmt_ListAll,
		Title: "List All Known Formatters",
		Desc:  "Lists all known " + Lang.Title + " formatters and their installation info / status",
	}
}

func (me *SrcFormattingBase) Cmds() (cmds []*coreCmd) {
	if me.cmdListAll.Hint == "" {
		kfnames := []string{}
		for _, kf := range me.Self.KnownFormatters() {
			kfnames = append(kfnames, kf.Name)
		}
		me.cmdListAll.Hint = strings.Join(kfnames, " · ")
	}
	cmds = append(cmds, me.cmdListAll)
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
			resp.Note = strf("After installing '%s', reload Zentient to recognize it.", kf.Name)
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
		var cmd = coreCmd{Title: kf.Name, MsgArgs: kf.Name}
		if !kf.Installed {
			cmd.MsgID = msgID_srcFmt_InfoLink
			cmd.Hint = "Not installed"
			cmd.Desc = "➜ Open installation infos at " + kf.Link
		} else {
			cmd.Hint = "Installed"
			cmd.Desc = "➜ Set as default formatter"
		}
		m.Choices = append(m.Choices, &cmd)
	}
	resp.CoreCmdsMenu = &m
}
