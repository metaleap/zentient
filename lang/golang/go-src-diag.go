package zgo

import (
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

func ensureBuildOrder(dis z.IDiagJobTarget, dat z.IDiagJobTarget) bool {
	return dis.(*udevgo.Pkg).IsSortedPriorToByDeps(dat.(*udevgo.Pkg))
}

func (me *goDiag) onUpdateDiagsPkgJobs(workspaceFiles z.WorkspaceFiles, filePaths []string) (jobs []z.DiagJob) {
	if pkgs := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		for _, pkg := range pkgs {
			jobs = append(jobs, z.DiagJob{AffectedFilePaths: pkg.GoFilePaths(), Target: pkg})
		}
	}
	return
}

func (me *goDiag) OnUpdateBuildDiags(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if dependant, pkgjobs := "", me.onUpdateDiagsPkgJobs(workspaceFiles, writtenFilePaths); len(pkgjobs) > 0 {
		byimppath := func(job *z.DiagJobBuild) bool { return job.Target.(*udevgo.Pkg).ImportPath == dependant }
		for _, pj := range pkgjobs {
			jobs = append(jobs, &z.DiagJobBuild{DiagJob: pj, TargetCmp: ensureBuildOrder})
			for _, dependant = range pj.Target.(*udevgo.Pkg).Dependants() {
				if pkgdep := udevgo.PkgsByImP[dependant]; pkgdep != nil && jobs.Find(byimppath) == nil {
					jobs = append(jobs, &z.DiagJobBuild{DiagJob: z.DiagJob{Target: pkgdep, AffectedFilePaths: pkgdep.GoFilePaths()}, TargetCmp: ensureBuildOrder})
				}
			}
		}
	}
	return
}

func (me *goDiag) OnUpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
	if pkgjobs := me.onUpdateDiagsPkgJobs(workspaceFiles, filePaths); len(pkgjobs) > 0 {
		for _, pj := range pkgjobs {
			skippkg := false
			for _, fpath := range pj.Target.(*udevgo.Pkg).GoFilePaths() {
				if skippkg = workspaceFiles.HasBuildDiags(fpath); skippkg {
					break
				}
			}
			if !skippkg {
				for _, dt := range diagTools {
					jobs = append(jobs, &z.DiagJobLint{DiagJob: pj, Tool: dt})
				}
			}
		}
	}
	return
}

func (me *goDiag) RunBuildJobs(jobs z.DiagBuildJobs) (diags z.DiagItems) {
	justfailed := make(map[string]bool, len(jobs))
	mockdiag := func(i int, fpath string, found string) *z.DiagItem {
		return &z.DiagItem{Message: "Found " + found, ToolName: "go install", FileRef: z.SrcLens{FilePath: fpath, Flag: int(z.DIAG_SEV_ERR), Pos: &z.SrcPos{Off: i + 1}}}
	}

	for _, pkgjob := range jobs {
		skip, pkg := false, pkgjob.Target.(*udevgo.Pkg)
		if len(justfailed) > 0 {
			for _, pdep := range pkg.Deps {
				if skip, _ = justfailed[pdep]; skip {
					break
				}
			}
		}
		if !skip {
			for _, fpath := range pkg.GoFilePaths() {
				filesrc := strings.ToLower(ufs.ReadTextFile(fpath, true, ""))
				if idx := strings.Index(filesrc, "fo"+"ol"); idx >= 0 {
					justfailed[pkg.ImportPath] = true
					diags = append(diags, mockdiag(idx, fpath, "fo"+"ol"))
				}
			}
		}
	}
	return
}

func (me *goDiag) RunLintJob(job *z.DiagJobLint) {
	defer job.Done()
	pkg, mockdiag := job.Target.(*udevgo.Pkg), func(i int, fpath string, found string) *z.DiagItem {
		return &z.DiagItem{Message: "Found " + found, ToolName: job.Tool.Name, FileRef: z.SrcLens{FilePath: fpath, Flag: int(job.Tool.DiagSev), Pos: &z.SrcPos{Off: i + 1}}}
	}
	for _, fpath := range pkg.GoFilePaths() {
		filesrc := ufs.ReadTextFile(fpath, true, "")
		if idx := strings.Index(filesrc, "/sys"); job.Tool.Name == "gosimple" && idx >= 0 {
			job.Yield(mockdiag(idx, fpath, "/sys"))
		}
		if idx := strings.Index(filesrc, "/run"); job.Tool.Name == "goconst" && idx >= 0 {
			job.Yield(mockdiag(idx, fpath, "/run"))
		}
		if idx := strings.Index(filesrc, "/slice"); job.Tool.Name == "golint" && idx >= 0 {
			job.Yield(mockdiag(idx, fpath, "/slice"))
		}
	}
}
