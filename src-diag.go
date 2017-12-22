package z

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/str"
)

type IDiag interface {
	IMenuItems

	KnownDiags() Tools
	OnUpdateBuildDiags(WorkspaceFiles, []string) DiagBuildJobs
	OnUpdateLintDiags(WorkspaceFiles, Tools, []string) DiagLintJobs
	RunBuildJobs(DiagBuildJobs) DiagItems
	RunLintJob(*DiagJobLint)
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
	ToolName   string
	FileRef    SrcLens
	Message    string
	SrcActions []EditorAction
}

func (me *DiagItem) resetAndInferSrcActions() {
	me.SrcActions = nil
	if ilastcolon := strings.LastIndex(me.Message, ":"); ilastcolon > 0 {
		if ilastnum := ustr.ToInt(me.Message[ilastcolon+1:]); ilastnum > 0 {
			if ifirstsep := strings.IndexRune(me.Message, filepath.Separator); ifirstsep >= 0 {
				refpath := me.Message[ifirstsep:]
				refpathf := refpath[:strings.IndexRune(refpath, ':')]
				if !ufs.FileExists(refpathf) {
					for i := ifirstsep - 1; i > 0; i-- {
						refpath = me.Message[i:]
						if refpathf = refpath[:strings.IndexRune(refpath, ':')]; ufs.FileExists(refpathf) {
							break
						}
					}
				}
				if ufs.FileExists(refpathf) && !filepath.IsAbs(refpathf) {
					refpathf, _ = filepath.Abs(refpathf)
				}
				if ufs.FileExists(refpathf) {
					cmd := EditorAction{Cmd: "zen.internal.openFileAt", Title: refpathf + refpath[strings.IndexRune(refpath, ':'):]}
					cmd.Arguments = append(cmd.Arguments, cmd.Title)
					cmd.Title = Strf("Jump to %s", filepath.Base(cmd.Title))
					me.SrcActions = append(me.SrcActions, cmd)
				}
			}
		}
	}
}

type DiagItems []*DiagItem

func (me DiagItems) propagate(lintDiags bool, workspaceFiles WorkspaceFiles) {
	for _, diag := range me {
		f := workspaceFiles.Ensure(diag.FileRef.FilePath)
		fd := &f.Diags.Lint
		if !lintDiags {
			fd = &f.Diags.Build
		}
		fd.UpToDate = true
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
		f := workspaceFiles.Ensure(filepath)
		fd := &f.Diags.Build
		if diagToolsIfLint != nil {
			fd = &f.Diags.Lint
		}
		fd.Forget(diagToolsIfLint)
		fd.UpToDate = true
	}
}

func (me *DiagJob) String() string { return me.Target.String() }

type DiagBase struct {
	Impl IDiag

	cmdListDiags     *MenuItem
	cmdListToggleAll *MenuItem
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
	me.cmdListToggleAll = &MenuItem{}
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

func (me *DiagBase) MenuItems(srcLens *SrcLens) (menu MenuItems) {
	me.menuItemsUpdateHint(me.knownDiags(true), me.cmdListDiags)
	menu = append(menu, me.cmdListDiags)
	if srcLens != nil && srcLens.FilePath != "" {
		nonautodiags, srcfilepath := me.knownDiags(false), srcLens.FilePath
		srcfilepath = Lang.Workspace.PrettyPath(srcfilepath)
		me.cmdRunDiagsOther.Desc = Strf("➜ run %d tool(s) on: %s", nonautodiags.Len(true), srcfilepath)
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

func (me *DiagBase) NewDiagItemFrom(srcRef *udev.SrcMsg, toolName string, fallbackFilePath string) (di *DiagItem) {
	di = &DiagItem{Message: srcRef.Msg, ToolName: toolName}
	di.FileRef.Flag = srcRef.Flag
	if srcRef.Pos2Ch > 0 && srcRef.Pos2Ln > 0 {
		di.FileRef.Range = &SrcRange{Start: SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch},
			End: SrcPos{Ln: srcRef.Pos2Ln, Col: srcRef.Pos2Ch}}
	} else {
		di.FileRef.Pos = &SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch}
	}
	if di.FileRef.FilePath = srcRef.Ref; di.FileRef.FilePath != "" && !filepath.IsAbs(di.FileRef.FilePath) {
		if absfilepath, err := filepath.Abs(di.FileRef.FilePath); err != nil {
			di.FileRef.FilePath = fallbackFilePath
		} else {
			di.FileRef.FilePath = absfilepath
		}
	}
	if di.FileRef.FilePath == "" || !ufs.FileExists(di.FileRef.FilePath) {
		di.FileRef.FilePath = fallbackFilePath
	}
	di.resetAndInferSrcActions()
	return
}

func (me *DiagBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCDIAG_LIST:
		me.onListAll(resp)
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

func (me *DiagBase) onListAll(resp *ipcResp) {
	resp.Menu = &MenuResp{SubMenu: &Menu{Desc: me.cmdListDiags.Desc}}
	knowndiagsauto, knowndiagsmanual := me.knownDiags(true), me.knownDiags(false)
	itemdesc := "Currently running automatically. ➜ Pick to turn this off."
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
		itemdesc = "Not currently running automatically. ➜ Pick to turn this on."
	}

	if len(resp.Menu.SubMenu.Items) > 0 {
		if len(knowndiagsauto) > 0 {
			me.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_NONE
			me.cmdListToggleAll.Title = "Disable All Auto-Diagnostics"
			me.cmdListToggleAll.Desc = "➜ if picked, no diagnostics tools will ever run on open and on save."
		} else {
			me.cmdListToggleAll.IpcID = IPCID_SRCDIAG_AUTO_ALL
			me.cmdListToggleAll.Title = "Enable All Auto-Diagnostics"
			me.cmdListToggleAll.Desc = "➜ if picked, all of the below diagnostics tools will run on open and on save."
		}
		resp.Menu.SubMenu.Items = append(MenuItems{me.cmdListToggleAll}, resp.Menu.SubMenu.Items...)
	}
	return
}

func (me *DiagBase) onToggleAll(enableAll bool, resp *ipcResp) {
	me.cmdRunDiagsOther.Hint, me.cmdListDiags.Hint = "", ""
	Prog.Cfg.AutoDiags = nil
	if enableAll {
		for _, diagtool := range me.Impl.KnownDiags() {
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
	resp.Menu = &MenuResp{NoteInfo: Strf("From now on, %s known-and-installed %s diagnostics tools will run automatically on open/save.", s, Lang.Title)}
	go me.onToggled()
}

func (me *DiagBase) onToggle(toolName string, resp *ipcResp) {
	me.cmdRunDiagsOther.Hint, me.cmdListDiags.Hint = "", ""
	if diagtool := me.Impl.KnownDiags().ByName(toolName); diagtool == nil {
		resp.ErrMsg = BadMsg(Lang.Title+" diagnostics tool name", toolName)
	} else if err := diagtool.ToggleInAutoDiags(); err != nil {
		resp.ErrMsg = err.Error()
	} else if diagtool.IsInAutoDiags() {
		resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` will run automatically on open/save.", Lang.Title, toolName)}
	} else {
		resp.Menu = &MenuResp{NoteInfo: Strf("The %s diagnostics tool `%s` won't run automatically on open/save.", Lang.Title, toolName)}
	}
	go me.onToggled()
}

func (me *DiagBase) onToggled() {
	Lang.Workspace.Lock()
	defer Lang.Workspace.Unlock()
	files := Lang.Workspace.Files()
	for _, f := range files {
		f.Diags.Lint.Forget(nil)
	}
	me.send(false)
	me.Impl.UpdateLintDiagsIfAndAsNeeded(files, true)
}

func (me *DiagBase) send(onlyBuildDiags bool) {
	files := Lang.Workspace.Files()
	resp := &DiagResp{LangID: Lang.ID, All: make(DiagItemsBy, len(files))}
	if !onlyBuildDiags {
		for _, f := range files {
			if onlyBuildDiags = len(f.Diags.Build.Items) > 0; onlyBuildDiags {
				break
			}
		}
	}
	for _, f := range files {
		fdiagitems := f.Diags.Lint.Items
		if onlyBuildDiags {
			fdiagitems = f.Diags.Build.Items
		}
		resp.All[f.Path] = fdiagitems
	}
	send(&ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: resp})
}
