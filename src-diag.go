package z

import (
	"fmt"
	"sort"
	"strings"
)

type IDiag interface {
	IMenuItems

	KnownDiags() Tools
	OnSend() *DiagResp
	OnUpdateBuildDiags(WorkspaceFiles, []string) DiagJobs
	OnUpdateLintDiags(WorkspaceFiles, Tools, []string) DiagJobs
	RunDiag(*DiagJob)
	UpdateBuildDiagsAsNeeded(WorkspaceFiles, []string)
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

type DiagItemsBy map[string]DiagItems

type DiagItem struct {
	ToolName string
	FileRef  SrcLens
	Message  string
}

type DiagItems []*DiagItem

type IDiagJobTarget interface {
	ISortable
	fmt.Stringer
}

type DiagJob struct {
	AffectedFilePaths []string
	Target            IDiagJobTarget
	Tool              *Tool
	TargetCmp         func(IDiagJobTarget, IDiagJobTarget) bool

	ch chan *DiagItem
}

func (me *DiagJob) Done()                { me.ch <- nil }
func (me *DiagJob) Yield(diag *DiagItem) { me.ch <- diag }
func (me *DiagJob) String() string       { return me.Target.String() }
func (me *DiagJob) IsSortedPriorTo(cmp interface{}) bool {
	c, _ := cmp.(*DiagJob)
	if me != nil && c != nil {
		if me.Target != nil && c.Target != nil {
			if me.TargetCmp != nil {
				return me.TargetCmp(me.Target, c.Target)
			} else {
				return me.Target.IsSortedPriorTo(c.Target)
			}
		}
		return me.Target == nil && c.Target == nil
	}
	return me == nil && c == nil
}

type DiagJobs []*DiagJob

func (me DiagJobs) Len() int               { return len(me) }
func (me DiagJobs) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me DiagJobs) Less(i int, j int) bool { return me[i].IsSortedPriorTo(me[j]) }

func (me DiagJobs) Find(check func(*DiagJob) bool) *DiagJob {
	for _, dj := range me {
		if check(dj) {
			return dj
		}
	}
	return nil
}

func (me *DiagJobs) RemoveAt(i int) {
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
	jobs := me.Impl.OnUpdateBuildDiags(workspaceFiles, writtenFiles)
	sort.Sort(jobs)
	println(Strf("%v", jobs))
}

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool) {
	if diagtools := me.knownDiags(autos); len(diagtools) > 0 {
		var filepaths []string
		for _, f := range workspaceFiles {
			if f != nil && f.IsOpen && !f.Diags.Lint.UpToDate {
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
			job.ch = await
			go me.Impl.RunDiag(job)
			for _, filepath := range job.AffectedFilePaths {
				if f, _ := workspaceFiles[filepath]; f != nil {
					f.Diags.Lint.Forget(diagTools)
					f.Diags.Lint.UpToDate = true
				}
			}
		}

		var diagitems []*DiagItem
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
		for _, diag := range diagitems {
			f := workspaceFiles.Ensure(diag.FileRef.FilePath)
			f.Diags.Lint.UpToDate = true
			f.Diags.Lint.Items = append(f.Diags.Lint.Items, diag)
		}
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
	msg := &ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: me.Impl.OnSend()}
	send(msg)
}

func (me *DiagBase) OnSend() *DiagResp {
	files := Lang.Workspace.Files()
	diags := &DiagResp{LangID: Lang.ID, All: make(DiagItemsBy, len(files))}
	for _, f := range files {
		if num := len(f.Diags.Build.Items) + len(f.Diags.Lint.Items); num > 0 {
			filediags := make(DiagItems, 0, num)
			filediags = append(filediags, f.Diags.Build.Items...)
			filediags = append(filediags, f.Diags.Lint.Items...)
			diags.All[f.Path] = filediags
		}
	}
	return diags
}
