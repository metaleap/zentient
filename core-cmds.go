package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/slice"
)

type iCoreCmds interface {
	iHandler

	Cmds(srcLoc *SrcLoc) []*coreCmd
	CmdsCategory() string
}

type coreCmdsMenu struct {
	Desc     string     `json:"d,omitempty"`
	TopLevel bool       `json:"tl,omitempty"`
	Choices  []*coreCmd `json:"c,omitempty"`
}

type coreCmd struct {
	MsgID    msgIDs      `json:"mi,omitempty"`
	MsgArgs  interface{} `json:"ma,omitempty"`
	Category string      `json:"c,omitempty"`
	Title    string      `json:"t"`
	Desc     string      `json:"d,omitempty"`
	Hint     string      `json:"h,omitempty"`
}

type coreCmdsHandler struct {
}

func (me *coreCmdsHandler) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_coreCmds_ListAll:
		me.handle_ListAll(req, resp)
	default:
		return false
	}
	return true
}

func (me *coreCmdsHandler) handle_ListAll(req *msgReq, resp *msgResp) {
	var cats sort.StringSlice
	m := coreCmdsMenu{Desc: "Showing: ", TopLevel: true}
	for _, cmds := range cmdProviders {
		for _, cmd := range cmds.Cmds(req.SrcLoc) {
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

func (me *coreCmdsHandler) Init() {
	l := &Lang
	if l.SrcFmt != nil {
		cmdProviders = append(cmdProviders, l.SrcFmt)
	}

	for _, cmds := range cmdProviders {
		cmds.Init()
		handlers = append(handlers, cmds)
	}
}
