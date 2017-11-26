package z

type IMetaCmds interface {
	Cmds() []*MetaCmd
	CmdsCategory() string
}

type MetaCmdsMenu struct {
	Desc    string     `json:"d"`
	Choices []*MetaCmd `json:"c"`
}

type MetaCmd struct {
	ID       string `json:"i"`
	MsgID    MsgIDs `json:"m,omitempty"`
	Category string `json:"c"`
	Title    string `json:"t"`
	Desc     string `json:"d1,omitempty"`
	Detail   string `json:"d2,omitempty"`
}

func metaCmdsProvidersUpdate() {
	l := &Lang
	l.cmdProviders = []IMetaCmds{}

	if l.CodeFmt != nil {
		l.cmdProviders = append(l.cmdProviders, l.CodeFmt)
	}
}

func handleMetaCmdsListAll(req *MsgReq, resp *MsgResp) {
	m := MetaCmdsMenu{Desc: "Choose wisely, mister:"}
	for _, cmds := range Lang.cmdProviders {
		for _, cmd := range cmds.Cmds() {
			cmd.Category = cmds.CmdsCategory()
			m.Choices = append(m.Choices, cmd)
		}
	}
	resp.MetaCmdsMenu = &m
}
