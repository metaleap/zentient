package z

import (
	"strings"
)

type IDiag interface {
	IMenuItems

	KnownDiags() Tools
	OnSend() *DiagResp
	UpdateLintDiags(WorkspaceFiles, Tools, []string)
	UpdateLintDiagsIfAndAsNeeded(WorkspaceFiles, bool)
}

type Diags struct {
	UpToDate bool      `json:",omitempty"`
	Items    DiagItems `json:",omitempty"`
}

func (me *Diags) Forget(onlyFor Tools) {
	if len(onlyFor) == 0 {
		me.UpToDate, me.Items = false, nil
	} else {
		for i := 0; i < len(me.Items); i++ {
			if onlyFor.Has(me.Items[i].ToolName) {
				pre, post := me.Items[:i], me.Items[i+1:]
				me.Items = append(pre, post...)
				i--
			}
		}
	}
}

type diagItems map[string]DiagItems

type DiagItem struct {
	ToolName string
	FileRef  SrcLens
	Message  string
}

type DiagItems []*DiagItem

type DiagBase struct {
	Impl IDiag

	cmdListDiags     *MenuItem
	cmdRunDiagsOther *MenuItem
}

type DiagResp struct {
	All    diagItems
	LangID string
}

func (me *DiagBase) Init() {
	me.cmdListDiags = &MenuItem{
		IpcID: IPCID_SRCDIAG_LIST,
		Title: "Choose Auto-Diagnostics",
		Desc:  Strf("Select which (out of %d) %s diagnostics tools should run automatically (on open and on save)", me.Impl.KnownDiags().Len(true), Lang.Title),
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
				item.Hint = Strf("(%d/%d)  · %s", len(diags), len(me.Impl.KnownDiags()), strings.Join(toolnames, " · "))
			}
		}
	}

	updatehint(me.knownDiags(true), me.cmdListDiags)
	menu = append(menu, me.cmdListDiags)
	if srcLens != nil && srcLens.FilePath != "" {
		nonautodiags, srcfilepath := me.knownDiags(false), srcLens.FilePath
		srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
		me.cmdRunDiagsOther.Desc = Strf("➜ run %d tools on: %s", nonautodiags.Len(true), srcfilepath)
		updatehint(nonautodiags, me.cmdRunDiagsOther)
		menu = append(menu, me.cmdRunDiagsOther)
	}
	return
}

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(files WorkspaceFiles, autos bool) {
	var diagtools = me.knownDiags(autos)

	if len(diagtools) > 0 {
		var filepaths []string
		for _, f := range files {
			if f != nil && f.IsOpen && !f.Diags.Lint.UpToDate {
				filepaths = append(filepaths, f.Path)
			}
		}
		if len(filepaths) > 0 {
			me.Impl.UpdateLintDiags(files, diagtools, filepaths)
		}
	}
	go me.send()
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
			item := &MenuItem{Title: dt.Name}
			if dt.Installed {
				item.Hint = "Installed  ·  " + dt.Website
				item.IpcID, item.IpcArgs = IPCID_SRCDIAG_TOGGLE, dt.Name
				if isinautodiags {
					item.Desc = "Currently running automatically. ➜ Pick to turn this off."
				} else {
					item.Desc = "Not currently running automatically. ➜ Pick to turn this on."
				}
			} else {
				item.Hint = "Not Installed"
				item.Desc = "➜ " + dt.Website
				item.IpcArgs = dt.Website
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
	} else {
		if tool.IsInAutoDiags() {
			resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` will be run automatically on open/save.", Lang.Title, toolName)}
		} else {
			resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` won't be run automatically on open/save, but may be invoked manually via the Zentient Main Menu.", Lang.Title, toolName)}
		}
		me.onToggled()
	}
}

func (me *DiagBase) onToggled() {
	Lang.Workspace.Lock()
	defer Lang.Workspace.Unlock()
	files := Lang.Workspace.Files()
	for _, f := range files {
		f.Diags.Lint.Forget(nil)
	}
	me.Impl.UpdateLintDiagsIfAndAsNeeded(files, true)
}

func (me *DiagBase) send() {
	msg := &ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: me.Impl.OnSend()}
	send(msg)
}

func (me *DiagBase) OnSend() (diags *DiagResp) {
	files := Lang.Workspace.Files()
	diags = &DiagResp{LangID: Lang.ID, All: make(diagItems, len(files))}
	for _, f := range files {
		if num := len(f.Diags.Build.Items) + len(f.Diags.Lint.Items); num > 0 {
			filediags := make(DiagItems, 0, num)
			filediags = append(filediags, f.Diags.Build.Items...)
			filediags = append(filediags, f.Diags.Lint.Items...)
			diags.All[f.Path] = filediags
		}
	}
	return
}
