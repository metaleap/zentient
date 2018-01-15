package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var diag goDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type goDiag struct {
	z.DiagBase

	knownTools z.Tools
}

func (me *goDiag) onPreInit() {
	me.knownTools = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}

func (me *goDiag) onUpdateDiagsPrepPkgJobs(filePaths []string) (jobs []z.DiagJob) {
	if pkgs, shouldrefresh := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		if shouldrefresh {
			go caddyRunRefreshPkgs()
		}
		for _, pkg := range pkgs {
			if !(pkg.Standard || pkg.BinaryOnly) {
				if pkggofilepaths := pkg.GoFilePaths(true); len(pkggofilepaths) > 0 {
					jobs = append(jobs, z.DiagJob{AffectedFilePaths: pkggofilepaths, Target: pkg})
				}
			}
		}
	}
	return
}

func (*goDiag) fallbackFilePath(pkg *udevgo.Pkg, workspaceFiles z.WorkspaceFiles) (filePath string) {
	for _, fp := range pkg.GoFilePaths(false) {
		if filePath == "" {
			filePath = fp
		}
		if workspacefile, _ := workspaceFiles[fp]; workspacefile != nil {
			if filePath = fp; workspacefile.IsOpen {
				break
			}
		}
	}
	return
}
