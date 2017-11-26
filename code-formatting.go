package z

type ICodeFormatting interface {
	IMetaCmds
}

type CodeFormattingBase struct {
}

func (me *CodeFormattingBase) Cmds() (cmds []*MetaCmd) {
	cmds = append(cmds, &MetaCmd{ID: "1", Title: "Choice One", Desc: "First choice.. First choice.. First choice.. First choice.. First choice.. First choice.. First choice.. ", Detail: "1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, "})
	cmds = append(cmds, &MetaCmd{ID: "2", Title: "Choice Two", Desc: "Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. ", Detail: "2 details, 2 details, 2 details, 2 details, 2 details, 2 details, 2 details, 2 details, "})
	cmds = append(cmds, &MetaCmd{ID: "3", Title: "Choice Tri", Desc: "Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. ", Detail: "3 details, 3 details, 3 details, 3 details, 3 details, 3 details, 3 details, 3 details, "})
	return
}

func (me *CodeFormattingBase) CmdsCategory() string {
	return "Formatting"
}
