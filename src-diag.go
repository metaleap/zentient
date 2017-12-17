package z

import (
	"fmt"
	"sort"
	"strings"
)

type IDiag interface {
	IMenuItems

	KnownDiags() Tools
	OnUpdateBuildDiags(WorkspaceFiles, []string) DiagBuildJobs
	OnUpdateLintDiags(WorkspaceFiles, Tools, []string) DiagLintJobs
	RunBuildJobs(DiagBuildJobs) DiagItems
	RunLintJob(*DiagJobLint)
	send()
	UpdateBuildDiagsAsNeeded(WorkspaceFiles, []string)
	UpdateLintDiagsIfAndAsNeeded(WorkspaceFiles, bool)
}

type Diags struct {
	upToDate bool      `json:",omitempty"`
	Items    DiagItems `json:",omitempty"`
}

func (me *Diags) Forget(onlyFor Tools) {
	if len(onlyFor) == 0 {
		me.upToDate, me.Items = false, nil
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

type DiagItemsBy map[string]DiagItems

type DiagItem struct {
	ToolName string
	FileRef  SrcLens
	Message  string
}

type DiagItems []*DiagItem

func (me DiagItems) propagate(lintDiags bool, workspaceFiles WorkspaceFiles) {
	for _, diag := range me {
		f := workspaceFiles.Ensure(diag.FileRef.FilePath)
		fd := &f.Diags.Lint
		if !lintDiags {
			fd = &f.Diags.Build
		}
		fd.upToDate = true
		fd.Items = append(fd.Items, diag)
	}
}

type IDiagJobTarget interface {
	ISortable
	fmt.Stringer
}

type DiagJob struct {
	AffectedFilePaths []string
	Target            IDiagJobTarget
}

func (me *DiagJob) forgetAndMarkUpToDate(diagToolsIfLint Tools, workspaceFiles WorkspaceFiles) {
	for _, filepath := range me.AffectedFilePaths {
		if f, _ := workspaceFiles[filepath]; f != nil {
			fd := &f.Diags.Build
			if diagToolsIfLint != nil {
				fd = &f.Diags.Lint
			}
			fd.Forget(diagToolsIfLint)
			fd.upToDate = true
		}
	}

}

func (me *DiagJob) String() string { return me.Target.String() }

type DiagJobBuild struct {
	DiagJob
	TargetCmp func(IDiagJobTarget, IDiagJobTarget) bool
	Succeeded bool
	diags     DiagItems
}

func (me *DiagJobBuild) Yield(diag *DiagItem) { me.diags = append(me.diags, diag) }

func (me *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool {
	c := cmp.(*DiagJobBuild)
	if me.TargetCmp != nil {
		return me.TargetCmp(me.Target, c.Target)
	}
	return me.Target.IsSortedPriorTo(c.Target)
}

type DiagJobLint struct {
	DiagJob
	Tool     *Tool
	lintChan chan *DiagItem
}

func (me *DiagJobLint) Done()                { me.lintChan <- nil }
func (me *DiagJobLint) Yield(diag *DiagItem) { me.lintChan <- diag }

type DiagLintJobs []*DiagJobLint

type DiagBuildJobs []*DiagJobBuild

func (me DiagBuildJobs) Len() int               { return len(me) }
func (me DiagBuildJobs) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me DiagBuildJobs) Less(i int, j int) bool { return me[i].IsSortedPriorTo(me[j]) }

func (me DiagBuildJobs) Find(check func(*DiagJobBuild) bool) *DiagJobBuild {
	for _, dj := range me {
		if check(dj) {
			return dj
		}
	}
	return nil
}

func (me *DiagBuildJobs) RemoveAt(i int) {
	m := *me
	pref, suff := m[:i], m[i+1:]
	*me = append(pref, suff...)
}

type DiagBase struct {
	Impl IDiag

	cmdListDiags     *MenuItem
	cmdRunDiagsOther *MenuItem
}

type DiagResp struct {
	All    DiagItemsBy
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
	me.menuItemsUpdateHint(me.knownDiags(true), me.cmdListDiags)
	menu = append(menu, me.cmdListDiags)
	if srcLens != nil && srcLens.FilePath != "" {
		nonautodiags, srcfilepath := me.knownDiags(false), srcLens.FilePath
		srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
		me.cmdRunDiagsOther.Desc = Strf("➜ run %d tools on: %s", nonautodiags.Len(true), srcfilepath)
		me.menuItemsUpdateHint(nonautodiags, me.cmdRunDiagsOther)
		menu = append(menu, me.cmdRunDiagsOther)
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
			item.Hint = Strf("(%d/%d)  · %s", len(diags), len(me.Impl.KnownDiags()), strings.Join(toolnames, " · "))
		}
	}
}

func (me *DiagBase) UpdateBuildDiagsAsNeeded(workspaceFiles WorkspaceFiles, writtenFiles []string) {
	if jobs := me.Impl.OnUpdateBuildDiags(workspaceFiles, writtenFiles); len(jobs) > 0 {
		sort.Sort(jobs)
		for _, job := range jobs {
			job.forgetAndMarkUpToDate(nil, workspaceFiles)
		}
		diagitems := me.Impl.RunBuildJobs(jobs)
		diagitems.propagate(false, workspaceFiles)
	}
	go me.send()
}

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool) {
	if diagtools := me.knownDiags(autos); len(diagtools) > 0 {
		var filepaths []string
		for _, f := range workspaceFiles {
			if f != nil && f.IsOpen && !f.Diags.Lint.upToDate {
				filepaths = append(filepaths, f.Path)
			}
		}
		if len(filepaths) > 0 {
			me.updateLintDiags(workspaceFiles, diagtools, filepaths)
		}
	}
	go me.send()
}

func (me *DiagBase) updateLintDiags(workspaceFiles WorkspaceFiles, diagTools Tools, filePaths []string) {
	if jobs := me.Impl.OnUpdateLintDiags(workspaceFiles, diagTools, filePaths); len(jobs) > 0 {
		numjobs, numdone, await := 0, 0, make(chan *DiagItem)
		for _, job := range jobs {
			numjobs++
			job.lintChan = await
			go me.Impl.RunLintJob(job)
			job.forgetAndMarkUpToDate(diagTools, workspaceFiles)
		}

		var diagitems DiagItems
		for numdone < numjobs {
			select {
			case diagitem := <-await:
				if diagitem == nil {
					numdone++
				} else {
					diagitems = append(diagitems, diagitem)
				}
			}
		}
		diagitems.propagate(true, workspaceFiles)
	}
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
	hasbuilddiags, files := false, Lang.Workspace.Files()
	resp := &DiagResp{LangID: Lang.ID, All: make(DiagItemsBy, len(files))}
	for _, f := range files {
		if hasbuilddiags = len(f.Diags.Build.Items) > 0; hasbuilddiags {
			break
		}
	}
	for _, f := range files {
		fdiagitems := f.Diags.Lint.Items
		if hasbuilddiags {
			fdiagitems = f.Diags.Build.Items
		}
		resp.All[f.Path] = fdiagitems
	}
	send(&ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: resp})
}
