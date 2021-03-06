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

func (me *DiagBase) Init() {
	me.cmdListDiags = &MenuItem{
		IpcID: IPCID_SRCDIAG_LIST,
		Title: Strf("Choose %s Auto-Linters", Lang.Title),
		Desc:  Strf("Select which (out of %d) lintish tools should run automatically (on file open/save)", me.Impl.KnownLinters().len(true)),
	}
	me.cmdListToggleAll = &MenuItem{}
	me.cmdRunDiagsOnCurFile = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_CURFILE, Title: Strf("Run Other %s Linters Now", Lang.Title)}
	me.cmdRunDiagsOnKnownFiles = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_ALLFILES, Title: me.cmdRunDiagsOnCurFile.Title, tag: "known"}
	me.cmdRunDiagsOnOpenFiles = &MenuItem{IpcID: IPCID_SRCDIAG_RUN_OPENFILES, Title: me.cmdRunDiagsOnCurFile.Title, tag: "opened"}
	me.cmdForgetAllDiags = &MenuItem{IpcID: IPCID_SRCDIAG_FORGETALL, Title: Strf("Forget All Currently Known %s Diagnostics", Lang.Title)}
	me.cmdPeekHiddenDiags = &MenuItem{IpcID: IPCID_SRCDIAG_PEEKHIDDEN, Title: Strf("Peek Hidden %s Lints", Lang.Title)}
}

func (me *DiagBase) MenuCategory() string {
	return "Diagnostics"
}

func (me *DiagBase) menuItems(srcLens *SrcLens) (menu MenuItems) {
	menu = make(MenuItems, 0, 5)
	autodiags, nonautodiags, workspacefiles := me.knownLinters(true), me.knownLinters(false), Lang.Workspace.Files()
	me.menuItemsUpdateHint(autodiags, me.cmdListDiags)
	menu = append(menu, me.cmdListDiags)
	if srcfilepath, numnonautos := "", nonautodiags.len(true); numnonautos > 0 {
		if srcLens != nil && srcLens.FilePath != "" {
			srcfilepath = srcLens.FilePath
			me.cmdRunDiagsOnCurFile.IpcArgs = srcfilepath
			me.cmdRunDiagsOnCurFile.Desc = Strf("➜ on: %s", Lang.Workspace.PrettyPath(srcfilepath))
			me.menuItemsUpdateHint(nonautodiags, me.cmdRunDiagsOnCurFile)
			menu = append(menu, me.cmdRunDiagsOnCurFile)
		}
		for menuitem, wfps := range map[*MenuItem][]string{me.cmdRunDiagsOnOpenFiles: workspacefiles.filePathsOpened(), me.cmdRunDiagsOnKnownFiles: workspacefiles.filePathsKnown()} {
			if l := len(wfps); l > 0 && (l > 1 || wfps[0] != srcfilepath) {
				menuitem.Desc = Strf("➜ on: %d currently-%s %s source file(s) in %d folder(s)",
					l, menuitem.tag, Lang.Title, workspacefiles.numDirs(func(f *WorkspaceFile) bool { return f.IsOpen || menuitem != me.cmdRunDiagsOnOpenFiles }))
				menuitem.Hint = ustr.Join(ustr.Map(wfps, filepath.Base), " · ")
				menu = append(menu, menuitem)
			}
		}
	}
	if ds := workspacefiles.diagsSummary(); ds != nil {
		hiddenlintnum, hiddenlintfiles, hiddenlintcats := 0, map[*WorkspaceFile]bool{}, map[string]bool{}
		me.cmdForgetAllDiags.Hint = Strf("for %d file(s) in %d folder(s)",
			len(ds.files), workspacefiles.numDirs(func(f *WorkspaceFile) bool { return ds.files[f] }))
		for dsf := range ds.files {
			if me.cmdForgetAllDiags.Hint += " — " + filepath.Base(dsf.Path); !dsf.IsOpen {
				if l := len(dsf.Diags.Lintishs.Items); l > 0 {
					hiddenlintfiles[dsf] = true
					for _, lintdiag := range dsf.Diags.Lintishs.Items {
						if !lintdiag.StickyAuto {
							hiddenlintnum, hiddenlintcats[lintdiag.Cat] = hiddenlintnum+1, true
						}
					}
				}
			}
		}
		me.cmdForgetAllDiags.Desc = Strf("Clears %d lint(s) and %d build-on-save diagnostic(s) ", ds.numLint, ds.numBuild)
		menu = append(menu, me.cmdForgetAllDiags)
		if me.cmdPeekHiddenDiags.IpcArgs = hiddenlintnum; hiddenlintnum > 0 {
			me.cmdPeekHiddenDiags.Hint = Strf("for %d file(s) in %d folder(s)",
				len(hiddenlintfiles), workspacefiles.numDirs(func(f *WorkspaceFile) bool { return hiddenlintfiles[f] }))
			for hlf := range hiddenlintfiles {
				me.cmdPeekHiddenDiags.Hint += " — " + filepath.Base(hlf.Path)
			}
			me.cmdPeekHiddenDiags.Desc = Strf("%d hidden lint(s) from", hiddenlintnum)
			for toolname := range hiddenlintcats {
				me.cmdPeekHiddenDiags.Desc += " — " + toolname
			}
			menu = append(menu, me.cmdPeekHiddenDiags)
		}
	}
	return
}

func (me *DiagBase) menuItemsUpdateHint(diags Tools, item *MenuItem) {
	if item.Hint == "" {
		toolnames := []string{}
		for _, dt := range diags {
			toolnames = append(toolnames, dt.Name)
		}
		if len(toolnames) == 0 {
			item.Hint = "(none)"
		} else {
			item.Hint = Strf("(%d/%d)  · %s", len(diags), len(me.Impl.KnownLinters()), ustr.Join(toolnames, " · "))
		}
	}
}

func (me *DiagBase) dispatch(req *IpcReq, resp *IpcResp) bool {
	switch req.IpcID {
	case IPCID_SRCDIAG_LIST:
		me.onListAll(resp)
	case IPCID_SRCDIAG_RUN_CURFILE:
		me.onRunManually([]string{req.IpcArgs.(string)}, resp)
	case IPCID_SRCDIAG_RUN_OPENFILES:
		me.onRunManually(nil, resp)
	case IPCID_SRCDIAG_RUN_ALLFILES:
		me.onRunManually([]string{}, resp)
	case IPCID_SRCDIAG_FORGETALL:
		me.onForgetAll()
	case IPCID_SRCDIAG_PEEKHIDDEN:
		me.onPeekHidden(int(req.IpcArgs.(float64)), resp.withMenu())
	case IPCID_SRCDIAG_AUTO_TOGGLE:
		me.onToggle(req.IpcArgs.(string), resp)
	case IPCID_SRCDIAG_AUTO_ALL:
		me.onToggleAll(true, resp)
	case IPCID_SRCDIAG_AUTO_NONE:
		me.onToggleAll(false, resp)
	default:
		return false
	}
	return true
}

func (me *DiagBase) onPeekHidden(approxNum int, resp *MenuResponse) {
	workspacefiles := Lang.Workspace.Files()
	resp.Refs = make(SrcLocs, 0, approxNum)
	for _, f := range workspacefiles {
		if (!f.IsOpen) && len(f.Diags.Lintishs.Items) > 0 {
			for _, lintdiag := range f.Diags.Lintishs.Items {
				if !lintdiag.StickyAuto {
					resp.Refs = append(resp.Refs, &lintdiag.Loc)
				}
			}
		}
	}
}

func (me *DiagBase) onForgetAll() {
	workspacefiles := Lang.Workspace.Files()
	for _, f := range workspacefiles {
		f.resetDiags()
	}
	me.send(workspacefiles, false)
}

var onRunManuallyInfoNoteAlreadyShownOnceInThisSession, onRunManuallyAlreadyCurrentlyRunning = false, false

func (me *DiagBase) onRunManually(filePaths []string, resp *IpcResp) {
	if onRunManuallyAlreadyCurrentlyRunning {
		resp.Menu = &MenuResponse{NoteWarn: "Declined: previous batch of lintish jobs still running, please wait until those have finished."}
	} else {
		workspacefiles := Lang.Workspace.Files()
		if filePaths == nil {
			filePaths = workspacefiles.filePathsOpened()
		} else if len(filePaths) == 0 {
			filePaths = workspacefiles.filePathsKnown()
		}
		if workspacefiles.haveAnyDiags(true, false) {
			resp.Menu = &MenuResponse{NoteWarn: "Any lintish findings will not display as long as the currently shown build problems remain unresolved in the workspace."}
		} else if !onRunManuallyInfoNoteAlreadyShownOnceInThisSession {
			onRunManuallyInfoNoteAlreadyShownOnceInThisSession = true
			resp.Menu = &MenuResponse{NoteInfo: Strf("All lintish findings (if any) will show up shortly and remain visible until invalidated.")}
		}
		go me.Impl.UpdateLintDiagsIfAndAsNeeded(workspacefiles, false, filePaths...)
	}
}

func (me *DiagBase) onListAll(resp *IpcResp) {
	resp.Menu = &MenuResponse{SubMenu: &Menu{Desc: me.cmdListDiags.Desc}}
	knowndiagsauto, knowndiagsmanual := me.knownLinters(true), me.knownLinters(false)
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
			me.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_NONE
			me.cmdListToggleAll.Title = "Disable All Auto-Linters"
			me.cmdListToggleAll.Desc = "➜ pick to have no lintish tools ever run on file open/save."
		} else {
			me.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_ALL
			me.cmdListToggleAll.Title = "Enable All Auto-Linters"
			me.cmdListToggleAll.Desc = "➜ pick to have all of the below lintish tools run on file open/save."
		}
		resp.Menu.SubMenu.Items = append(MenuItems{me.cmdListToggleAll}, resp.Menu.SubMenu.Items...)
	}
}

func (me *DiagBase) onToggleAll(enableAll bool, resp *IpcResp) {
	me.cmdRunDiagsOnCurFile.Hint, me.cmdListDiags.Hint = "", ""
	if Prog.Cfg.AutoDiags = nil; enableAll {
		for _, diagtool := range me.Impl.KnownLinters() {
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
	resp.Menu = &MenuResponse{NoteInfo: Strf("From now on, %s known-and-installed %s lintish tools will run automatically on file open/save.", s, Lang.Title)}
	go me.onToggled()
}

func (me *DiagBase) onToggle(toolName string, resp *IpcResp) {
	me.cmdRunDiagsOnCurFile.Hint, me.cmdListDiags.Hint = "", ""
	if diagtool := me.Impl.KnownLinters().byName(toolName); diagtool == nil {
		resp.ErrMsg = BadMsg(Lang.Title+" lintish tool name", toolName)
	} else if err := diagtool.toggleInAutoDiags(); err != nil {
		resp.ErrMsg = err.Error()
	} else if diagtool.isInAutoDiags() {
		resp.Menu = &MenuResponse{NoteInfo: Strf("The %s lintish tool `%s` will run automatically on file open/save.", Lang.Title, toolName)}
	} else {
		resp.Menu = &MenuResponse{NoteInfo: Strf("The %s lintish tool `%s` won't run automatically on file open/save.", Lang.Title, toolName)}
	}
	go me.onToggled()
}

func (me *DiagBase) onToggled() {
	Lang.Workspace.Lock()
	defer Lang.Workspace.Unlock()
	workspaceFiles := Lang.Workspace.Files()
	for _, f := range workspaceFiles {
		f.Diags.Lintishs.forget(nil)
		f.Diags.AutoLintUpToDate = false
	}
	me.Impl.UpdateLintDiagsIfAndAsNeeded(workspaceFiles, true)
}

func (me *DiagBase) send(workspaceFiles WorkspaceFiles, onlyProbDiags bool) {
	resp := &Diags{LangID: Lang.ID, All: make(DiagItemsBy, len(workspaceFiles))}
	onlyProbDiags = onlyProbDiags || workspaceFiles.haveAnyDiags(true, false)
	for _, f := range workspaceFiles {
		fdiagitems := f.Diags.Lintishs.Items
		if onlyProbDiags {
			fdiagitems = f.Diags.Problems.Items
		}
		if fdiagitems != nil {
			resp.All[f.Path] = fdiagitems.dropDupls()
		}
	}
	send(&IpcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: resp})
}
