package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/slice"
)

type iMetaCmds interface {
	Cmds() []*metaCmd
	CmdsCategory() string
	Init()
}

type metaCmdsMenu struct {
	Desc    string     `json:"d"`
	Choices []*metaCmd `json:"c"`
}

type metaCmd struct {
	ID       string `json:"i"`
	MsgID    msgIDs `json:"m,omitempty"`
	Category string `json:"c"`
	Title    string `json:"t"`
	Desc     string `json:"d,omitempty"`
	Hint     string `json:"h,omitempty"`
}

func metaCmdsProvidersInit() { // assumes Lang.cmdProviders is empty
	l := &Lang

	if l.CodeFmt != nil {
		l.cmdProviders = append(l.cmdProviders, l.CodeFmt)
	}

	for _, cmds := range l.cmdProviders {
		cmds.Init()
	}
}

func metaCmdsHandle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_metaCmds_ListAll:
		metaCmdsHandleListAll(req, resp)
	default:
		return false
	}
	return true
}

func metaCmdsHandleListAll(req *msgReq, resp *msgResp) {
	var cats sort.StringSlice
	m := metaCmdsMenu{Desc: "Categories: "}
	for _, cmds := range Lang.cmdProviders {
		for _, cmd := range cmds.Cmds() {
			if cmd.Category = cmds.CmdsCategory(); !uslice.StrHas(cats, cmd.Category) {
				cats = append(cats, cmd.Category)
			}
			m.Choices = append(m.Choices, cmd)
		}
	}
	sort.Sort(cats)
	m.Desc += strings.Join(cats, " · ")
	resp.MetaCmdsMenu = &m
}
