package zat

import (
	"path/filepath"
	"strings"

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
	ondir := func(dirpath string) { Ctx.KitByDirPath(dirpath, true) }
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

func (*atmoWorkspace) onAfterChanges(*z.WorkspaceChanges) {
	Ctx.CatchUp(true)
}

func (me *atmoWorkspace) onPreInit() {
	me.OnBeforeChanges, me.OnAfterChanges = me.onBeforeChanges, me.onAfterChanges
}
