package zgo

import (
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
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
	if pkgs, shouldrefresh := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		if shouldrefresh {
			go caddyRunRefreshPkgs()
		}
		for _, pkg := range pkgs {
			jobs = append(jobs, z.DiagJob{AffectedFilePaths: pkg.GoFilePaths(), Target: pkg})
		}
	}
	return
}

func (me *goDiag) OnUpdateBuildDiags(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if pkgjobs := me.onUpdateDiagsPkgJobs(workspaceFiles, writtenFilePaths); len(pkgjobs) > 0 {
		for _, pj := range pkgjobs {
			jobs = append(jobs, &z.DiagJobBuild{DiagJob: pj, TargetCmp: ensureBuildOrder})
			for _, dependant := range pj.Target.(*udevgo.Pkg).Dependants() {
				if pkgdep := udevgo.PkgsByImP[dependant]; pkgdep != nil {
					jobs = append(jobs, &z.DiagJobBuild{DiagJob: z.DiagJob{Target: pkgdep, AffectedFilePaths: pkgdep.GoFilePaths()}, TargetCmp: ensureBuildOrder})
				}
			}
		}
		jobs = jobs.WithoutDuplicates()
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

func (*goDiag) fallbackFilePath(pkg *udevgo.Pkg) (filePath string) {
	workspacefiles := z.Lang.Workspace.Files()
	for _, fp := range pkg.GoFilePaths() {
		if filePath == "" {
			filePath = fp
		}
		if workspacefile, _ := workspacefiles[fp]; workspacefile != nil {
			if filePath = fp; workspacefile.IsOpen {
				break
			}
		}
	}
	return
}

func (me *goDiag) runBuildPkg(pkg *udevgo.Pkg) (diags z.DiagItems) {
	if msgs := udev.CmdExecOnSrc(true, nil, "go", "install", pkg.ImportPath); len(msgs) > 0 {
		diags = make(z.DiagItems, 0, len(msgs))
		fallbackfilepath, skipmsg := me.fallbackFilePath(pkg), "package "+pkg.ImportPath+":"
		for _, srcref := range msgs {
			if srcref.Msg != "too many errors" && !(srcref.Pos1Ch == 1 && srcref.Pos1Ln == 1 && srcref.Msg == skipmsg) {
				diags = append(diags, me.NewDiagItemFrom(srcref, "", fallbackfilepath))
			}
		}
	}
	return
}

func (me *goDiag) RunBuildJobs(jobs z.DiagBuildJobs) (diags z.DiagItems) {
	numjobs := len(jobs)
	failed, skipped := make(map[string]bool, numjobs), make(map[string]bool, numjobs)
	pkgnames := make([]string, 0, numjobs)
	for i := 0; i < numjobs; i++ {
		pkgnames = append(pkgnames, jobs[i].Target.(*udevgo.Pkg).ImportPath)
	}

	for i, pkgjob := range jobs {
		caddyBuildOnRunning(numjobs, i, pkgnames)
		skip, pkg := false, pkgjob.Target.(*udevgo.Pkg)
		if len(failed) > 0 {
			for _, pdep := range pkg.Deps {
				if skip, _ = failed[pdep]; skip {
					skipped[pkg.ImportPath] = true
					break
				}
			}
		}
		if !skip {
			pkgdiags := me.runBuildPkg(pkg)
			if diags, pkgjob.Succeeded = append(diags, pkgdiags...), len(pkgdiags) == 0; !pkgjob.Succeeded {
				failed[pkg.ImportPath] = true
			}
		}
	}
	caddyBuildOnDone(failed, skipped, pkgnames)
	go caddyRunRefreshPkgs()
	return
}

func (me *goDiag) RunLintJob(job *z.DiagJobLint) {
	defer job.Done()
	jt, pkg := job.Tool, job.Target.(*udevgo.Pkg)
	fallbackfilepath := me.fallbackFilePath(pkg)
	var msgs udev.SrcMsgs
	if jt == tools.gosimple {
		msgs = udevgo.LintGoSimple(pkg.ImportPath)
	} else if jt == tools.golint {
		msgs = udevgo.LintGolint(pkg.ImportPath)
	} else if jt == tools.goconst {
		msgs = udevgo.LintGoConst(pkg.Dir)
	} else if jt == tools.govet {
		msgs = udevgo.LintGoVet(pkg.ImportPath)
	} else {
		msgs = append(msgs, &udev.SrcMsg{Msg: z.BadMsg("lint tool", job.Tool.Name)})
	}
	for _, srcref := range msgs {
		srcref.Flag = int(job.Tool.DiagSev)
		job.Yield(me.NewDiagItemFrom(srcref, job.Tool.Name, fallbackfilepath))
	}
}
