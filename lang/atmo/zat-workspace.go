package zat

import (
	"path/filepath"
	"strings"

	"github.com/go-leap/str"
	"github.com/metaleap/atmo"
	"github.com/metaleap/zentient"
)

var (
	LoadKitsAsSoonAsFilesOpen = true

	workspace    atmoWorkspace
	goPathScopes []string
)

func init() {
	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
}

type atmoWorkspace struct {
	z.WorkspaceBase
}

func (*atmoWorkspace) onBeforeChanges(workspaceChanges *z.WorkspaceChanges, freshFiles []string, willAutoLint bool) {
	ondir := func(dirpath string) { _ = Ctx.KitByDirPath(dirpath, true) }
	for _, dirpath := range workspaceChanges.AddedDirs {
		ondir(dirpath)
	}
	if len(freshFiles) > 0 {
		for _, ffp := range freshFiles {
			if strings.ToLower(filepath.Ext(ffp)) == atmo.SrcFileExt {
				ondir(filepath.Dir(ffp))
			}
		}
	}
}

func (*atmoWorkspace) onAfterChanges(workspaceChanges *z.WorkspaceChanges) {
	Ctx.CatchUp(true)
	if LoadKitsAsSoonAsFilesOpen && len(workspaceChanges.OpenedFiles) > 0 {
		var kitstoload []string
		for _, srcfilepath := range workspaceChanges.OpenedFiles {
			if kit := Ctx.KitByDirPath(filepath.Dir(srcfilepath), true); kit != nil && !ustr.In(kit.ImpPath, kitstoload...) {
				kitstoload = append(kitstoload, kit.ImpPath)
			}
		}
		Ctx.KitsEnsureLoadedFully(false, kitstoload...)
	}
}

func (me *atmoWorkspace) onPreInit() {
	me.OnBeforeChanges, me.OnAfterChanges = me.onBeforeChanges, me.onAfterChanges
}
