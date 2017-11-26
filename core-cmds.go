package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/slice"
)

type iCoreCmds interface {
	iResponder

	Cmds() []*coreCmd
	CmdsCategory() string
}

type coreCmdsMenu struct {
	Desc    string     `json:"d"`
	Choices []*coreCmd `json:"c"`
}

type coreCmd struct {
	ID       string `json:"i"`
	MsgID    msgIDs `json:"m,omitempty"`
	Category string `json:"c"`
	Title    string `json:"t"`
	Desc     string `json:"d,omitempty"`
	Hint     string `json:"h,omitempty"`
}

type coreCmds struct {
}

func (me *coreCmds) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_coreCmds_ListAll:
		me.handle_ListAll(req, resp)
	default:
		return false
	}
	return true
}

func (me *coreCmds) handle_ListAll(req *msgReq, resp *msgResp) {
	var cats sort.StringSlice
	m := coreCmdsMenu{Desc: "Categories: "}
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
	resp.CoreCmdsMenu = &m
}

func (me *coreCmds) Init() {
	l := &Lang

	if l.CodeFmt != nil {
		l.cmdProviders = append(l.cmdProviders, l.CodeFmt)
	}

	for _, cmds := range l.cmdProviders {
		cmds.Init()
		handlers = append(handlers, cmds)
	}
}
