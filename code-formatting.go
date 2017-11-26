package z

type iCodeFormatting interface {
	iMetaCmds

	KnownFormatters() []string
}

type CodeFormattingBase struct {
	cmdListAll *metaCmd
}

func (me *CodeFormattingBase) Init() {
	me.cmdListAll = &metaCmd{
		ID: "lkf", MsgID: msgID_codeFmt_ListAll,
		Title: "List Known Formatters",
		Hint:  "gofmt goreturns etc",
		Desc:  "Lists all known " + Lang.Title + " formatters and their installation info",
	}
}

func (me *CodeFormattingBase) Cmds() (cmds []*metaCmd) {
	cmds = append(cmds, me.cmdListAll)
	return
}

func (me *CodeFormattingBase) CmdsCategory() string {
	return "Formatting"
}

func (me *CodeFormattingBase) KnownFormatters() []string {
	return nil
}

func codeFmtHandle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_codeFmt_ListAll:
		codeFmtHandleListAll(req, resp)
	default:
		return false
	}
	return true
}

func codeFmtHandleListAll(req *msgReq, resp *msgResp) {
	m := metaCmdsMenu{Desc: "List of formatters:"}
	m.Choices = append(m.Choices, &metaCmd{Title: "Foo Title", Category: "Foo Cat"})
	resp.MetaCmdsMenu = &m
}
