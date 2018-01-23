package zgo

import (
	"path/filepath"
	"strings"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/fs"
	"github.com/metaleap/zentient"
)

var (
	workspace    goWorkspace
	goPathScopes []string
)

func init() {
	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
}

type goWorkspace struct {
	z.WorkspaceBase
}

func (me *goWorkspace) onAfterChanges(upd *z.WorkspaceChanges) {
	if sep := string(filepath.Separator); upd.HasDirChanges() {
		goPathScopes, udevgo.GuruScopes = nil, ""
		for _, dp := range me.Dirs() {
			dp_ := dp.Path + sep
			for _, gp := range udevgo.Gopaths() {
				if gpsrc := filepath.Join(gp, "src"); strings.HasPrefix(gp, dp_) && ufs.IsDir(gpsrc) {
					ufs.WalkDirsIn(gpsrc, func(gopathsubdir string) bool {
						goPathScopes = append(goPathScopes, gopathsubdir[len(gpsrc)+1:]+"/...")
						return true
					})
				} else if gpsrc_ := gpsrc + sep; strings.HasPrefix(dp_, gpsrc_) {
					goPathScopes = append(goPathScopes, dp_[len(gpsrc_):]+"...")
				}
			}
		}
		if len(goPathScopes) > 0 {
			udevgo.GuruScopes = strings.Join(goPathScopes, ",")
		}
	}
}

func (me *goWorkspace) onBeforeChanges(upd *z.WorkspaceChanges, freshFiles []string, willAutoLint bool) {
	if hasnewpkgs := false; udevgo.PkgsByDir != nil && len(freshFiles) > 0 {
		for _, ffp := range freshFiles {
			if hasnewpkgs = strings.ToLower(filepath.Ext(ffp)) == ".go" && (nil == udevgo.PkgsByDir[filepath.Dir(ffp)]); hasnewpkgs {
				break
			}
		}
		if hasnewpkgs && caddyRefreshPkgs.IsReady() {
			if caddyRefreshPkgs.IsPendingOrBusy() {
				caddyRefreshPkgs.ShouldReRunWhenNextDone = true
			} else {
				go caddyRunRefreshPkgs()
			}
		}
	}
}

func (me *goWorkspace) onPreInit() {
	me.OnBeforeChanges, me.OnAfterChanges = me.onBeforeChanges, me.onAfterChanges
}
