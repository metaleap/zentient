package zat

import (
	"path/filepath"

	"github.com/go-leap/str"
	"github.com/metaleap/atmo"
	"github.com/metaleap/zentient"
)

var (
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
	var kitimppaths []string
	for _, ffp := range append(append(freshFiles, workspaceChanges.WrittenFiles...), workspaceChanges.OpenedFiles...) {
		if filepath.Ext(ffp) == atmo.SrcFileExt {
			dirpath := filepath.Dir(ffp)
			if kit := Ctx.KitByDirPath(dirpath, true); kit != nil && !ustr.In(kit.ImpPath, kitimppaths...) {
				kitimppaths = append(kitimppaths, kit.ImpPath)
			}
		}
	}
	Ctx.CatchUpOnFileMods()
	Ctx.KitsEnsureLoaded(false, kitimppaths...)
}

func (*atmoWorkspace) onAfterChanges(workspaceChanges *z.WorkspaceChanges) {
}

func (me *atmoWorkspace) onPreInit() {
	me.OnBeforeChanges, me.OnAfterChanges = me.onBeforeChanges, me.onAfterChanges
}
