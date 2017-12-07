package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/slice"
)

type iMenuProvider interface {
	iDispatcher

	MenuItems(*SrcLens) []*MenuItem
	MenuCategory() string
}

type Menu struct {
	Desc     string      `json:"d,omitempty"`
	TopLevel bool        `json:"tl,omitempty"`
	Items    []*MenuItem `json:"i,omitempty"`
}

type MenuItem struct {
	MsgID    msgIDs      `json:"mi,omitempty"`
	MsgArgs  interface{} `json:"ma,omitempty"`
	Category string      `json:"c,omitempty"`
	Title    string      `json:"t"`
	Desc     string      `json:"d,omitempty"`
	Hint     string      `json:"h,omitempty"`
}

type MenuItemMsgArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

type MenuResp struct {
	SubMenu    *Menu  `json:"menu,omitempty"`
	WebsiteURL string `json:"url,omitempty"`
	NoteInfo   string `json:"info,omitempty"`
	NoteWarn   string `json:"warn,omitempty"`
	MsgAction  string `json:"action,omitempty"`
}

type mainMenu struct {
}

func (me *mainMenu) dispatch(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_MENUS_MAIN:
		me.onListAll(req, resp)
	default:
		return false
	}
	return true
}

func (me *mainMenu) onListAll(req *msgReq, resp *msgResp) {
	var cats sort.StringSlice
	m := Menu{Desc: "Showing: ", TopLevel: true}
	for _, menu := range menuProviders {
		for _, item := range menu.MenuItems(req.SrcLens) {
			if item.Category = menu.MenuCategory(); !uslice.StrHas(cats, item.Category) {
				cats = append(cats, item.Category)
			}
			m.Items = append(m.Items, item)
		}
	}
	sort.Sort(cats)
	m.Desc += strings.Join(cats, " · ")
	resp.Menu = &MenuResp{SubMenu: &m}
}

func (me *mainMenu) Init() {
}
