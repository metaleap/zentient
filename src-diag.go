package z

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
		Title: "Choose Auto-Diagnostics",
		Desc:  Strf("Select which %s diagnostics should run automatically (on open and on save)", Lang.Title),
	}
	me.cmdRunDiagsOther = &MenuItem{
		Title: "Run Non-Auto-Diagnostics Now",
		Desc:  Strf("Runs all %s diagnostics that do not run automatically.", Lang.Title),
	}
}

func (me *DiagBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	default:
		return false
	}
	return true
}

func (me *DiagBase) MenuCategory() string {
	return "Diagnostics"
}

func (me *DiagBase) MenuItems(srcLens *SrcLens) (menu []*MenuItem) {
	menu = append(menu, me.cmdListDiags)
	if srcLens != nil && srcLens.FilePath != "" {
		menu = append(menu, me.cmdRunDiagsOther)
	}
	return
}
