package z

import (
	"strings"
)

type IDiag interface {
	IMenuItems

	KnownDiags() Tools
}

type DiagBase struct {
	Impl IDiag

	cmdListDiags     *MenuItem
	cmdRunDiagsOther *MenuItem
}

func (me *DiagBase) Init() {
	me.cmdListDiags = &MenuItem{
		IpcID: IPCID_SRCDIAG_LIST,
		Title: "Choose Auto-Diagnostics",
		Desc:  Strf("Select out of %d which %s diagnostics tools should run automatically (on open and on save)", len(me.Impl.KnownDiags()), Lang.Title),
	}
	me.cmdRunDiagsOther = &MenuItem{
		IpcID: IPCID_SRCDIAG_RUN,
		Title: "Run Non-Auto-Diagnostics Now",
	}
}

func (me *DiagBase) knownDiags(auto bool) (diags Tools) {
	for _, dt := range me.Impl.KnownDiags() {
		if dt.IsInAutoDiags() == auto {
			diags = append(diags, dt)
		}
	}
	return
}

func (me *DiagBase) MenuCategory() string {
	return "Diagnostics"
}

func (me *DiagBase) MenuItems(srcLens *SrcLens) (menu []*MenuItem) {
	updatehint := func(diags Tools, item *MenuItem) {
		if item.Hint == "" {
			toolnames := []string{}
			for _, dt := range diags {
				toolnames = append(toolnames, dt.Name)
			}
			if len(toolnames) == 0 {
				item.Hint = "(none)"
			} else {
				item.Hint = Strf("(%d/%d)  · ", len(diags), len(me.Impl.KnownDiags())) + strings.Join(toolnames, " · ")
			}
		}
	}

	if srcLens != nil && srcLens.FilePath != "" {
		nonautodiags, srcfilepath := me.knownDiags(false), srcLens.FilePath
		if Lang.Workspace != nil {
			srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
		}
		me.cmdRunDiagsOther.Desc = Strf("➜ run %d diagnostics tools on: %s", len(nonautodiags), srcfilepath)
		updatehint(nonautodiags, me.cmdRunDiagsOther)
		menu = append(menu, me.cmdRunDiagsOther)
	}
	updatehint(me.knownDiags(true), me.cmdListDiags)
	menu = append(menu, me.cmdListDiags)
	return
}

func (me *DiagBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCDIAG_LIST:
		me.onListAll(resp)
	case IPCID_SRCDIAG_TOGGLE:
		me.onToggle(req.IpcArgs.(string), resp)
	default:
		return false
	}
	return true
}

func (me *DiagBase) onListAll(resp *ipcResp) {
	resp.Menu = &MenuResp{SubMenu: &Menu{Desc: me.cmdListDiags.Desc}}
	isinautodiags := true
	for _, kd := range []Tools{me.knownDiags(true), me.knownDiags(false)} {
		for _, dt := range kd {
			item := &MenuItem{Title: dt.Name, Hint: dt.Website, IpcID: IPCID_SRCDIAG_TOGGLE, IpcArgs: dt.Name}
			if dt.Installed {
				item.Hint = "Installed  ·  " + item.Hint
			} else {
				item.Hint = "Not Installed  ·  " + item.Hint
			}
			if isinautodiags {
				item.Desc = "Currently running automatically. ➜ Pick to turn this off."
			} else {
				item.Desc = "Not currently running automatically. ➜ Pick to turn this on."
			}
			resp.Menu.SubMenu.Items = append(resp.Menu.SubMenu.Items, item)
		}
		isinautodiags = false
	}
	return
}

func (me *DiagBase) onToggle(toolName string, resp *ipcResp) {
	me.cmdRunDiagsOther.Hint, me.cmdListDiags.Hint = "", ""
	if tool := me.Impl.KnownDiags().ByName(toolName); tool == nil {
		resp.ErrMsg = BadMsg(Lang.Title+" diagnostics tool name", toolName)
	} else if err := tool.ToggleInAutoDiags(); err != nil {
		resp.ErrMsg = err.Error()
	} else if tool.IsInAutoDiags() {
		resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` will be run automatically on open/save.", Lang.Title, toolName)}
	} else {
		resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` won't be run automatically on open/save, but may be invoked manually via the Zentient Main Menu.", Lang.Title, toolName)}
	}
}
