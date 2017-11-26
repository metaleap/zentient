package z

import (
	"strings"
)

type iCodeFormatting interface {
	iResponder
	iMetaCmds

	IsInstalled(string) bool
	KnownFormatters() []string
}

type CodeFormattingBase struct {
	cmdListAll *metaCmd

	Self iCodeFormatting
}

func (me *CodeFormattingBase) Init() {
	me.cmdListAll = &metaCmd{
		ID: "lkf", MsgID: msgID_codeFmt_ListAll,
		Title: "List Known Formatters",
		Desc:  "Lists all known " + Lang.Title + " formatters and their installation info",
	}
}

func (me *CodeFormattingBase) Cmds() (cmds []*metaCmd) {
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

func (CodeFormattingBase) handle_ListAll(req *msgReq, resp *msgResp) {
	m := metaCmdsMenu{Desc: "List of formatters:"}
	cfmt := Lang.CodeFmt // need the interface impl, not the embedded base!
	cat := cfmt.CmdsCategory()
	for _, fmtname := range cfmt.KnownFormatters() {
		var cmd = metaCmd{Category: cat, Title: fmtname}
		if !cfmt.IsInstalled(fmtname) {
			cmd.Hint = "Not installed"
			cmd.Desc = "Click to open installation instructions"
		} else {
			cmd.Hint = "Installed"
			cmd.Desc = "Click to set as default formatter"
		}
		m.Choices = append(m.Choices, &cmd)
	}
	resp.MetaCmdsMenu = &m
}
