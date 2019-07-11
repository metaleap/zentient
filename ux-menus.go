package z

import (
	"github.com/go-leap/str"
)

type IMenuItems interface {
	iDispatcher

	menuItems(*SrcLens) MenuItems
	MenuCategory() string
}

type menuItemIpcArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

type mainMenu struct {
}

func (me *mainMenu) dispatch(req *IpcReq, resp *IpcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_MAIN:
		me.onMainMenu(req, resp)
	default:
		return false
	}
	return true
}

func (*mainMenu) onMainMenu(req *IpcReq, resp *IpcResp) {
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
	resp.Menu = &MenuResponse{SubMenu: &m}
}

func (*mainMenu) Init() {
}
