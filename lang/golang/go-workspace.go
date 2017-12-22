package zgo

import (
	"path/filepath"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var workspace goWorkspace

func init() {
	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
}

type goWorkspace struct {
	z.WorkspaceBase
}

func (me *goWorkspace) onBeforeChanges(upd *z.WorkspaceChanges, dirsChanged bool, freshFiles []string, willAutoLint bool) {
	if hasnewpkgs := false; udevgo.PkgsByDir != nil {
		for _, nfp := range freshFiles {
			if hasnewpkgs = (nil == udevgo.PkgsByDir[filepath.Dir(nfp)]); hasnewpkgs {
				break
			}
		}
		if hasnewpkgs && caddyRefreshPkgs.Ready() && !caddyRefreshPkgs.PendingOrBusy() {
			caddyRunRefreshPkgs()
		}
	}
}

func (me *goWorkspace) onPreInit() {
	me.OnBeforeChanges = me.onBeforeChanges
}
