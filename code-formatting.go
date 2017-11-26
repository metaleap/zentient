package z

import (
	"strings"
)

type iCodeFormatting interface {
	iCoreCmds

	IsInstalled(string) bool
	KnownFormatters() []string
}

type CodeFormattingBase struct {
	cmdListAll *coreCmd

	Self iCodeFormatting
}

func (me *CodeFormattingBase) Init() {
	me.cmdListAll = &coreCmd{
		ID: "lkf", MsgID: msgID_codeFmt_ListAll,
		Title: "List Known Formatters",
		Desc:  "Lists all known " + Lang.Title + " formatters and their installation info",
	}
}

func (me *CodeFormattingBase) Cmds() (cmds []*coreCmd) {
	if me.cmdListAll.Hint == "" {
		me.cmdListAll.Hint = strings.Join(me.Self.KnownFormatters(), " · ")
	}
	cmds = append(cmds, me.cmdListAll)
	return
}

func (me *CodeFormattingBase) CmdsCategory() string {
	return "Formatting"
}

func (me *CodeFormattingBase) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_codeFmt_ListAll:
		me.handle_ListAll(req, resp)
	default:
		return false
	}
	return true
}

func (me *CodeFormattingBase) handle_ListAll(req *msgReq, resp *msgResp) {
	cfmt := Lang.CodeFmt // need the interface impl, not the embedded base!
	m := coreCmdsMenu{Desc: strf("❬%s❭ · %s:", cfmt.CmdsCategory(), me.cmdListAll.Title)}
	for _, fmtname := range cfmt.KnownFormatters() {
		var cmd = coreCmd{Title: fmtname}
		if !cfmt.IsInstalled(fmtname) {
			cmd.Hint = "Not installed"
			cmd.Desc = "Click to open installation instructions"
		} else {
			cmd.Hint = "Installed"
			cmd.Desc = "Click to set as default formatter"
		}
		m.Choices = append(m.Choices, &cmd)
	}
	resp.CoreCmdsMenu = &m
}
