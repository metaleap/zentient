package z

import (
	"github.com/go-leap/str"
)

type IMenuItems interface {
	iDispatcher

	menuItems(*SrcLens) MenuItems
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
	Confirm  string      `json:"q,omitempty"`

	tag string
}

type menuItemIpcArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

type menuResp struct {
	SubMenu       *Menu   `json:",omitempty"`
	WebsiteURL    string  `json:",omitempty"`
	NoteInfo      string  `json:",omitempty"`
	NoteWarn      string  `json:",omitempty"`
	UxActionLabel string  `json:",omitempty"`
	Refs          SrcLocs `json:",omitempty"`
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

func (*mainMenu) onMainMenu(req *ipcReq, resp *ipcResp) {
	var cats []string
	catfilter, _ := req.IpcArgs.(string)
	m := Menu{Desc: "Categories:  ", TopLevel: true}
	if catfilter != "" {
		m.Desc = "Category:  "
	}
	for _, menu := range Prog.menus {
		for _, item := range menu.menuItems(req.SrcLens) {
			if item.Category = menu.MenuCategory(); catfilter == "" || item.Category == catfilter {
				if !ustr.In(item.Category, cats...) {
					cats = append(cats, item.Category)
				}
				m.Items = append(m.Items, item)
			}
		}
	}
	m.Desc += ustr.Join(cats, "  ·  ")
	resp.Menu = &menuResp{SubMenu: &m}
}

func (*mainMenu) Init() {
}
