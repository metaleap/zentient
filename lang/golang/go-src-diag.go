package zgo

import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/zentient"
)

var diag goDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type goDiag struct {
	z.DiagBase

	knownDiags z.Tools
}

func (me *goDiag) onPreInit() {
	me.knownDiags = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}

func (me *goDiag) KnownDiags() z.Tools {
	return me.knownDiags
}

func (me *goDiag) OnUpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagJobs) {
	if pkgs := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		for _, pkg := range pkgs {
			pkgfilepaths := make([]string, 0, len(pkg.GoFiles))
			for _, filename := range pkg.GoFiles {
				pkgfilepaths = append(pkgfilepaths, filepath.Join(pkg.Dir, filename))
			}
			for _, diagtool := range diagTools {
				jobs = append(jobs, &z.DiagJob{Tool: diagtool, AffectedFilePaths: pkgfilepaths, Target: pkg})
			}
		}
	}
	return
}

func (me *goDiag) RunDiag(job *z.DiagJob, yield z.DiagItemsChan) {
	defer yield.Done()
	pkg, mockdiag := job.Target.(*udevgo.Pkg), func(i int, fpath string, found string) *z.DiagItem {
		return &z.DiagItem{Message: "Found " + found, ToolName: job.Tool.Name, FileRef: z.SrcLens{FilePath: fpath, Flag: int(job.Tool.DiagSev), Pos: &z.SrcPos{Off: i + 1}}}
	}
	for _, filename := range pkg.GoFiles {
		fpath := filepath.Join(pkg.Dir, filename)
		filesrc := ufs.ReadTextFile(fpath, true, "")
		if idx := strings.Index(filesrc, "/sys"); job.Tool.Name == "gosimple" && idx >= 0 {
			yield <- mockdiag(idx, fpath, "/sys")
		}
		if idx := strings.Index(filesrc, "/run"); job.Tool.Name == "goconst" && idx >= 0 {
			yield <- mockdiag(idx, fpath, "/run")
		}
		if idx := strings.Index(filesrc, "/slice"); job.Tool.Name == "golint" && idx >= 0 {
			yield <- mockdiag(idx, fpath, "/slice")
		}
	}
}
