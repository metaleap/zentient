package zgo

import (
	"github.com/go-leap/dev/go"
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

func (this *goDiag) onPreInit() {
	this.knownTools = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}

func (*goDiag) onUpdateDiagsPrepPkgJobs(filePaths []string) (jobs []z.DiagJob) {
	pkgs, shouldrefresh := udevgo.PkgsForFiles(filePaths...)
	if shouldrefresh {
		go caddyRunRefreshPkgs()
	}
	if len(pkgs) > 0 {
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
