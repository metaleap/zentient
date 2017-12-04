package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/slice"
)

type iCmdsProvider interface {
	iDispatcher

	Cmds(*SrcLens) []*coreCmd
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

type coreCmdMsgArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

type coreCmdResp struct {
	CoreCmdsMenu *coreCmdsMenu `json:"menu,omitempty"`
	WebsiteURL   string        `json:"url,omitempty"`
	NoteInfo     string        `json:"info,omitempty"`
	NoteWarn     string        `json:"warn,omitempty"`
	MsgAction    string        `json:"action,omitempty"`
}

type coreCmds struct {
}

func (me *coreCmds) dispatch(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_CORECMDS_PALETTE:
		me.onListAll(req, resp)
	default:
		return false
	}
	return true
}

func (me *coreCmds) onListAll(req *msgReq, resp *msgResp) {
	var cats sort.StringSlice
	m := coreCmdsMenu{Desc: "Showing: ", TopLevel: true}
	for _, cmds := range cmdProviders {
		for _, cmd := range cmds.Cmds(req.SrcLens) {
			if cmd.Category = cmds.CmdsCategory(); !uslice.StrHas(cats, cmd.Category) {
				cats = append(cats, cmd.Category)
			}
			m.Choices = append(m.Choices, cmd)
		}
	}
	sort.Sort(cats)
	m.Desc += strings.Join(cats, " · ")
	resp.CoreCmd = &coreCmdResp{CoreCmdsMenu: &m}
}

func (me *coreCmds) Init() {
}
