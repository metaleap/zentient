package zgo

import (
	"strings"
	"time"

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

func (me *goDiag) onUpdateDiagsPkgJobs(filePaths []string) (jobs []z.DiagJob) {
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

func (me *goDiag) OnUpdateBuildDiags(writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if pkgjobs := me.onUpdateDiagsPkgJobs(writtenFilePaths); len(pkgjobs) > 0 {
		for _, pj := range pkgjobs {
			jobs = append(jobs, &z.DiagJobBuild{DiagJob: pj, TargetCmp: ensureBuildOrder})
			for _, dependant := range pj.Target.(*udevgo.Pkg).Dependants() {
				if pkgdep := udevgo.PkgsByImP[dependant]; pkgdep != nil {
					jobs = append(jobs, &z.DiagJobBuild{DiagJob: z.DiagJob{Target: pkgdep, AffectedFilePaths: pkgdep.GoFilePaths()}, TargetCmp: ensureBuildOrder})
				}
			}
		}
	}
	return
}

func (me *goDiag) OnUpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
	if pkgjobs := me.onUpdateDiagsPkgJobs(filePaths); len(pkgjobs) > 0 {
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

func (*goDiag) fallbackFilePath(pkg *udevgo.Pkg, workspaceFiles z.WorkspaceFiles) (filePath string) {
	for _, fp := range pkg.GoFilePaths() {
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

func (me *goDiag) runBuildPkg(pkg *udevgo.Pkg, workspaceFiles z.WorkspaceFiles) (diags z.DiagItems) {
	if msgs := udev.CmdExecOnSrc(true, nil, "go", "install", pkg.ImportPath); len(msgs) > 0 {
		diags = make(z.DiagItems, 0, len(msgs))
		skipmsg, fallbackfilepath := "package "+pkg.ImportPath+":", func() string { return me.fallbackFilePath(pkg, workspaceFiles) }
		for _, srcref := range msgs {
			if srcref.Msg != "too many errors" && !(srcref.Pos1Ch == 1 && srcref.Pos1Ln == 1 && srcref.Msg == skipmsg) {
				diags = append(diags, me.NewDiagItemFrom(srcref, "", fallbackfilepath))
			}
		}
	}
	return
}

func (me *goDiag) RunBuildJobs(jobs z.DiagBuildJobs) (diags z.DiagItems) {
	numjobs, starttime, numbuilt := len(jobs), time.Now(), 0
	failed, skipped := make(map[string]bool, numjobs), make(map[string]bool, numjobs)
	pkgnames := make([]string, 0, numjobs)
	for i := 0; i < numjobs; i++ {
		pkgnames = append(pkgnames, jobs[i].Target.(*udevgo.Pkg).ImportPath)
	}
	allpkgnames := strings.Join(pkgnames, "\n")

	for i, pkgjob := range jobs {
		caddyBuildOnRunning(numjobs, i, allpkgnames)
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
			pkgdiags := me.runBuildPkg(pkg, pkgjob.WorkspaceFiles)
			if pkgjob.Succeeded, diags = len(pkgdiags) == 0, append(diags, pkgdiags...); pkgjob.Succeeded {
				numbuilt++
			} else {
				failed[pkg.ImportPath] = true
			}
		}
	}
	caddyBuildOnDone(failed, skipped, pkgnames, time.Since(starttime))
	if numbuilt > 0 {
		go caddyRunRefreshPkgs()
	}
	return
}

func (me *goDiag) RunLintJob(job *z.DiagJobLint) {
	defer job.Done()
	if !job.Tool.Installed {
		return
	}
	jt, pkg := job.Tool, job.Target.(*udevgo.Pkg)
	var msgs udev.SrcMsgs
	if jt == tools.gosimple {
		msgs = udevgo.LintGoSimple(pkg.ImportPath)
	} else if jt == tools.golint {
		msgs = udevgo.LintGolint(pkg.ImportPath)
	} else if jt == tools.goconst {
		msgs = udevgo.LintGoConst(pkg.Dir)
	} else if jt == tools.govet {
		msgs = udevgo.LintGoVet(pkg.ImportPath)
	} else if jt == tools.ineffassign {
		msgs = udevgo.LintIneffAssign(pkg.Dir)
	} else if jt == tools.maligned {
		msgs = udevgo.LintViaPkgImpPath("maligned", string(pkg.ImportPath), false)
	} else if jt == tools.unconvert {
		msgs = udevgo.LintViaPkgImpPath("unconvert", pkg.ImportPath, false)
	} else if jt == tools.errcheck {
		msgs = udevgo.LintErrcheck(pkg.ImportPath)
	} else if jt == tools.checkstruct {
		msgs = udevgo.LintCheck("structcheck", pkg.ImportPath)
	} else if jt == tools.checkalign {
		msgs = udevgo.LintCheck("aligncheck", pkg.ImportPath)
	} else if jt == tools.checkvar {
		msgs = udevgo.LintCheck("varcheck", pkg.ImportPath)
	} else if jt == tools.unparam {
		msgs = udevgo.LintMvDan("unparam", pkg.ImportPath)
	} else if jt == tools.interfacer {
		msgs = udevgo.LintMvDan("interfacer", pkg.ImportPath)
	} else if jt == tools.unindent {
		msgs = udevgo.LintMvDan("unindent", pkg.ImportPath)
	} else if jt == tools.deadcode {
		msgs = udevgo.LintViaPkgImpPath("deadcode", pkg.ImportPath, true)
	} else if jt == tools.unused {
		msgs = udevgo.LintHonnef("unused", pkg.ImportPath)
	} else if jt == tools.staticcheck {
		msgs = udevgo.LintHonnef("staticcheck", pkg.ImportPath)
	} else {
		msgs = append(msgs, &udev.SrcMsg{Msg: z.BadMsg("lint tool", job.Tool.Name)})
	}
	if len(msgs) > 0 {
		fallbackfilepath := func() string { return me.fallbackFilePath(pkg, job.WorkspaceFiles) }
		for _, srcref := range msgs {
			srcref.Flag = int(job.Tool.DiagSev)
			job.Yield(me.NewDiagItemFrom(srcref, job.Tool.Name, fallbackfilepath))
		}
	}
}
