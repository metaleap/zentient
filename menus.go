package z

import (
	"strings"

	"github.com/metaleap/go-util/slice"
)

type IMenuItems interface {
	iDispatcher

	MenuItems(*SrcLens) MenuItems
	MenuCategory() string
}

type Menu struct {
	Desc     string    `json:"desc,omitempty"`
	TopLevel bool      `json:"topLevel,omitempty"`
	Items    MenuItems `json:"items"`
}

type MenuItems []*MenuItem

type MenuItem struct {
	IpcID    IpcIDs      `json:"ii,omitempty"`
	IpcArgs  interface{} `json:"ia,omitempty"`
	Category string      `json:"c,omitempty"`
	Title    string      `json:"t"`
	Desc     string      `json:"d,omitempty"`
	Hint     string      `json:"h,omitempty"`
}

type MenuItemIpcArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

type MenuResp struct {
	SubMenu       *Menu  `json:"menu,omitempty"`
	WebsiteURL    string `json:"url,omitempty"`
	NoteInfo      string `json:"info,omitempty"`
	NoteWarn      string `json:"warn,omitempty"`
	UxActionLabel string `json:"uxActionLabel,omitempty"`
}

type mainMenu struct {
}

func (me *mainMenu) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_MAIN:
		me.onMainMenu(req, resp)
	default:
		return false
	}
	return true
}

func (me *mainMenu) onMainMenu(req *ipcReq, resp *ipcResp) {
	var cats []string
	catfilter, _ := req.IpcArgs.(string)
	m := Menu{Desc: "Categories:  ", TopLevel: true}
	if catfilter != "" {
		m.Desc = "Category:  "
	}
	for _, menu := range Prog.menus {
		for _, item := range menu.MenuItems(req.SrcLens) {
			if item.Category = menu.MenuCategory(); catfilter == "" || item.Category == catfilter {
				if !uslice.StrHas(cats, item.Category) {
					cats = append(cats, item.Category)
				}
				m.Items = append(m.Items, item)
			}
		}
	}
	m.Desc += strings.Join(cats, "  ·  ")
	resp.Menu = &MenuResp{SubMenu: &m}
}

func (me *mainMenu) Init() {
}
