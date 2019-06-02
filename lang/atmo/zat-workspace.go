package zat

import (
	"path/filepath"
	"strings"

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
	newpotentialkitsimppaths := make([]string, 0, 2)
	ondir := func(dirpath string) {
		if kit := Ctx.KitByDirPath(dirpath, true); kit != nil && !ustr.In(kit.ImpPath, newpotentialkitsimppaths...) {
			newpotentialkitsimppaths = append(newpotentialkitsimppaths, kit.ImpPath)
		}
	}

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
	Ctx.KitsEnsureLoaded(false, newpotentialkitsimppaths...)

	if len(workspaceChanges.WrittenFiles) > 0 {
		Ctx.CatchUp(true)
	}
}

func (me *atmoWorkspace) onPreInit() {
	me.OnBeforeChanges = me.onBeforeChanges
}
