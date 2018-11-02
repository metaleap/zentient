package z

import (
	"path/filepath"

	"github.com/go-leap/str"
)

type DiagBase struct {
	Impl IDiag

	cmdListDiags            *MenuItem
	cmdListToggleAll        *MenuItem
	cmdRunDiagsOnCurFile    *MenuItem
	cmdRunDiagsOnOpenFiles  *MenuItem
	cmdRunDiagsOnKnownFiles *MenuItem
	cmdForgetAllDiags       *MenuItem
	cmdPeekHiddenDiags      *MenuItem
}

func (this *DiagBase) Init() {
	this.cmdListDiags = &MenuItem{
		IpcID: IPCID_SRCDIAG_LIST,
		Title: Strf("Choose %s Auto-Linters", Lang.Title),
		Desc:  Strf("Select which (out of %d) lintish tools should run automatically (on file open/save)", this.Impl.KnownLinters().len(true)),
	}
	this.cmdListToggleAll = &MenuItem{}
	this.cmdRunDiagsOnCurFile = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_CURFILE, Title: Strf("Run Other %s Linters Now", Lang.Title)}
	this.cmdRunDiagsOnKnownFiles = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_ALLFILES, Title: this.cmdRunDiagsOnCurFile.Title, tag: "known"}
	this.cmdRunDiagsOnOpenFiles = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_OPENFILES, Title: this.cmdRunDiagsOnCurFile.Title, tag: "opened"}
	this.cmdForgetAllDiags = &MenuItem{IpcID: IPCID_SRCDIAG_FORGETALL, Title: Strf("Forget All Currently Known %s Diagnostics", Lang.Title)}
	this.cmdPeekHiddenDiags = &MenuItem{IpcID: IPCID_SRCDIAG_PEEKHIDDEN, Title: Strf("Peek Hidden %s Lints", Lang.Title)}
}

func (this *DiagBase) MenuCategory() string {
	return "Linting"
}

func (this *DiagBase) menuItems(srcLens *SrcLens) (menu MenuItems) {
	menu = make(MenuItems, 0, 5)
	autodiags, nonautodiags, workspacefiles := this.knownLinters(true), this.knownLinters(false), Lang.Workspace.Files()
	this.menuItemsUpdateHint(autodiags, this.cmdListDiags)
	menu = append(menu, this.cmdListDiags)
	if srcfilepath, numnonautos := "", nonautodiags.len(true); numnonautos > 0 {
		if srcLens != nil && srcLens.FilePath != "" {
			srcfilepath = srcLens.FilePath
			this.cmdRunDiagsOnCurFile.IpcArgs = srcfilepath
			this.cmdRunDiagsOnCurFile.Desc = Strf("➜ on: %s", Lang.Workspace.PrettyPath(srcfilepath))
			this.menuItemsUpdateHint(nonautodiags, this.cmdRunDiagsOnCurFile)
			menu = append(menu, this.cmdRunDiagsOnCurFile)
		}
		for menuitem, wfps := range map[*MenuItem][]string{this.cmdRunDiagsOnOpenFiles: workspacefiles.filePathsOpened(), this.cmdRunDiagsOnKnownFiles: workspacefiles.filePathsKnown()} {
			if l := len(wfps); l > 0 && (l > 1 || wfps[0] != srcfilepath) {
				menuitem.Desc = Strf("➜ on: %d currently-%s %s source file(s) in %d folder(s)",
					l, menuitem.tag, Lang.Title, workspacefiles.numDirs(func(f *WorkspaceFile) bool { return f.IsOpen || menuitem != this.cmdRunDiagsOnOpenFiles }))
				menuitem.Hint = ustr.Join(ustr.Map(wfps, filepath.Base), " · ")
				menu = append(menu, menuitem)
			}
		}
	}
	if ds := workspacefiles.diagsSummary(); ds != nil {
		hiddenlintnum, hiddenlintfiles, hiddenlintcats := 0, map[*WorkspaceFile]bool{}, map[string]bool{}
		this.cmdForgetAllDiags.Hint = Strf("for %d file(s) in %d folder(s)",
			len(ds.files), workspacefiles.numDirs(func(f *WorkspaceFile) bool { return ds.files[f] }))
		for dsf := range ds.files {
			if this.cmdForgetAllDiags.Hint += " — " + filepath.Base(dsf.Path); !dsf.IsOpen {
				if l := len(dsf.Diags.Lint.Items); l > 0 {
					hiddenlintfiles[dsf] = true
					for _, lintdiag := range dsf.Diags.Lint.Items {
						if !lintdiag.StickyAuto {
							hiddenlintnum, hiddenlintcats[lintdiag.Cat] = hiddenlintnum+1, true
						}
					}
				}
			}
		}
		this.cmdForgetAllDiags.Desc = Strf("Clears %d lint(s) and %d build-on-save diagnostic(s) ", ds.numLint, ds.numBuild)
		menu = append(menu, this.cmdForgetAllDiags)
		if this.cmdPeekHiddenDiags.IpcArgs = hiddenlintnum; hiddenlintnum > 0 {
			this.cmdPeekHiddenDiags.Hint = Strf("for %d file(s) in %d folder(s)",
				len(hiddenlintfiles), workspacefiles.numDirs(func(f *WorkspaceFile) bool { return hiddenlintfiles[f] }))
			for hlf := range hiddenlintfiles {
				this.cmdPeekHiddenDiags.Hint += " — " + filepath.Base(hlf.Path)
			}
			this.cmdPeekHiddenDiags.Desc = Strf("%d hidden lint(s) from", hiddenlintnum)
			for toolname := range hiddenlintcats {
				this.cmdPeekHiddenDiags.Desc += " — " + toolname
			}
			menu = append(menu, this.cmdPeekHiddenDiags)
		}
	}
	return
}

func (this *DiagBase) menuItemsUpdateHint(diags Tools, item *MenuItem) {
	if item.Hint == "" {
		toolnames := []string{}
		for _, dt := range diags {
			toolnames = append(toolnames, dt.Name)
		}
		if len(toolnames) == 0 {
			item.Hint = "(none)"
		} else {
			item.Hint = Strf("(%d/%d)  · %s", len(diags), len(this.Impl.KnownLinters()), ustr.Join(toolnames, " · "))
		}
	}
}

func (this *DiagBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCDIAG_LIST:
		this.onListAll(resp)
	case IPCID_SRCDIAG_RUN_CURFILE:
		this.onRunManually([]string{req.IpcArgs.(string)}, resp)
	case IPCID_SRCDIAG_RUN_OPENFILES:
		this.onRunManually(nil, resp)
	case IPCID_SRCDIAG_RUN_ALLFILES:
		this.onRunManually([]string{}, resp)
	case IPCID_SRCDIAG_FORGETALL:
		this.onForgetAll()
	case IPCID_SRCDIAG_PEEKHIDDEN:
		this.onPeekHidden(int(req.IpcArgs.(float64)), resp.withMenu())
	case IPCID_SRCDIAG_AUTO_TOGGLE:
		this.onToggle(req.IpcArgs.(string), resp)
	case IPCID_SRCDIAG_AUTO_ALL:
		this.onToggleAll(true, resp)
	case IPCID_SRCDIAG_AUTO_NONE:
		this.onToggleAll(false, resp)
	default:
		return false
	}
	return true
}

func (this *DiagBase) onPeekHidden(approxNum int, resp *menuResp) {
	workspacefiles := Lang.Workspace.Files()
	resp.Refs = make(SrcLocs, 0, approxNum)
	for _, f := range workspacefiles {
		if (!f.IsOpen) && len(f.Diags.Lint.Items) > 0 {
			for _, lintdiag := range f.Diags.Lint.Items {
				if !lintdiag.StickyAuto {
					resp.Refs = append(resp.Refs, &lintdiag.Loc)
				}
			}
		}
	}
}

func (this *DiagBase) onForgetAll() {
	workspacefiles := Lang.Workspace.Files()
	for _, f := range workspacefiles {
		f.resetDiags()
	}
	go this.send(workspacefiles, false)
}

var onRunManuallyInfoNoteAlreadyShownOnceInThisSession, onRunManuallyAlreadyCurrentlyRunning = false, false

func (this *DiagBase) onRunManually(filePaths []string, resp *ipcResp) {
	if onRunManuallyAlreadyCurrentlyRunning {
		resp.Menu = &menuResp{NoteWarn: "Declined: previous batch of lintish jobs still running, please wait until those have finished."}
	} else {
		workspacefiles := Lang.Workspace.Files()
		if filePaths == nil {
			filePaths = workspacefiles.filePathsOpened()
		} else if len(filePaths) == 0 {
			filePaths = workspacefiles.filePathsKnown()
		}
		if workspacefiles.haveAnyDiags(true, false) {
			resp.Menu = &menuResp{NoteWarn: "Any lintish findings will not display as long as the currently shown build problems remain unresolved in the workspace."}
		} else if !onRunManuallyInfoNoteAlreadyShownOnceInThisSession {
			onRunManuallyInfoNoteAlreadyShownOnceInThisSession = true
			resp.Menu = &menuResp{NoteInfo: Strf("All lintish findings (if any) will show up shortly and remain visible until invalidated.")}
		}
		go this.Impl.UpdateLintDiagsIfAndAsNeeded(workspacefiles, false, filePaths...)
	}
}

func (this *DiagBase) onListAll(resp *ipcResp) {
	resp.Menu = &menuResp{SubMenu: &Menu{Desc: this.cmdListDiags.Desc}}
	knowndiagsauto, knowndiagsmanual := this.knownLinters(true), this.knownLinters(false)
	itemdesc := "WILL run automatically on file open/save. ➜ Pick to turn this off."
	for _, knowndiags := range []Tools{knowndiagsauto, knowndiagsmanual} {
		for _, dt := range knowndiags {
			item := &MenuItem{Title: dt.Name}
			if dt.Installed {
				item.Hint = "Installed  ·  " + dt.Website
				item.IpcID, item.IpcArgs = IPCID_SRCDIAG_AUTO_TOGGLE, dt.Name
				item.Desc = itemdesc
			} else {
				item.Hint = "Not Installed"
				item.Desc = "➜ " + dt.Website
				item.IpcArgs = dt.Website
			}
			resp.Menu.SubMenu.Items = append(resp.Menu.SubMenu.Items, item)
		}
		itemdesc = "WON'T run automatically on file open/save. ➜ Pick to turn this on."
	}

	if len(resp.Menu.SubMenu.Items) > 0 {
		if len(knowndiagsauto) > 0 {
			this.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_NONE
			this.cmdListToggleAll.Title = "Disable All Auto-Linters"
			this.cmdListToggleAll.Desc = "➜ pick to have no lintish tools ever run on file open/save."
		} else {
			this.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_ALL
			this.cmdListToggleAll.Title = "Enable All Auto-Linters"
			this.cmdListToggleAll.Desc = "➜ pick to have all of the below lintish tools run on file open/save."
		}
		resp.Menu.SubMenu.Items = append(MenuItems{this.cmdListToggleAll}, resp.Menu.SubMenu.Items...)
	}
}

func (this *DiagBase) onToggleAll(enableAll bool, resp *ipcResp) {
	this.cmdRunDiagsOnCurFile.Hint, this.cmdListDiags.Hint = "", ""
	if Prog.Cfg.AutoDiags = nil; enableAll {
		for _, diagtool := range this.Impl.KnownLinters() {
			Prog.Cfg.AutoDiags = append(Prog.Cfg.AutoDiags, diagtool.Name)
		}
	}
	if err := Prog.Cfg.Save(); err != nil {
		resp.ErrMsg = err.Error()
	}
	s := "no"
	if enableAll {
		s = "all"
	}
	resp.Menu = &menuResp{NoteInfo: Strf("From now on, %s known-and-installed %s lintish tools will run automatically on file open/save.", s, Lang.Title)}
	go this.onToggled()
}

func (this *DiagBase) onToggle(toolName string, resp *ipcResp) {
	this.cmdRunDiagsOnCurFile.Hint, this.cmdListDiags.Hint = "", ""
	if diagtool := this.Impl.KnownLinters().byName(toolName); diagtool == nil {
		resp.ErrMsg = BadMsg(Lang.Title+" lintish tool name", toolName)
	} else if err := diagtool.toggleInAutoDiags(); err != nil {
		resp.ErrMsg = err.Error()
	} else if diagtool.isInAutoDiags() {
		resp.Menu = &menuResp{NoteInfo: Strf("The %s lintish tool `%s` will run automatically on file open/save.", Lang.Title, toolName)}
	} else {
		resp.Menu = &menuResp{NoteInfo: Strf("The %s lintish tool `%s` won't run automatically on file open/save.", Lang.Title, toolName)}
	}
	go this.onToggled()
}

func (this *DiagBase) onToggled() {
	Lang.Workspace.Lock()
	defer Lang.Workspace.Unlock()
	workspaceFiles := Lang.Workspace.Files()
	for _, f := range workspaceFiles {
		f.Diags.Lint.forget(nil)
		f.Diags.AutoLintUpToDate = false
	}
	this.Impl.UpdateLintDiagsIfAndAsNeeded(workspaceFiles, true)
}

func (this *DiagBase) send(workspaceFiles WorkspaceFiles, onlyBuildDiags bool) {
	resp := &diagResp{LangID: Lang.ID, All: make(diagItemsBy, len(workspaceFiles))}
	onlyBuildDiags = onlyBuildDiags || workspaceFiles.haveAnyDiags(true, false)
	for _, f := range workspaceFiles {
		fdiagitems := f.Diags.Lint.Items
		if onlyBuildDiags {
			fdiagitems = f.Diags.Build.Items
		}
		resp.All[f.Path] = fdiagitems
	}
	send(&ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: resp})
}
