package zgo

import (
	"strings"
	"time"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

func ensureBuildOrder(dis z.IDiagJobTarget, dat z.IDiagJobTarget) bool {
	return dis.(*udevgo.Pkg).IsSortedPriorToByDeps(dat.(*udevgo.Pkg))
}

func (me *goDiag) OnUpdateBuildDiags(writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if pkgjobs := me.onUpdateDiagsPrepPkgJobs(writtenFilePaths); len(pkgjobs) > 0 {
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

func (me *goDiag) FixerUppers() []z.FixerUpper {
	return []z.FixerUpper{me.tryFixupImpNotFound}
}

func (me *goDiag) tryFixupImpNotFound(d *z.DiagItem) (fix *z.FixUp) {
	//cannot find package "foo" in any of:
	pref, i := "cannot find package \"", strings.Index(d.Msg, "\" in any of:")
	if strings.HasPrefix(d.Msg, pref) && i > len(pref) {
		badimpname := d.Msg[:i][len(pref):]
		fix = &z.FixUp{Name: "Removes invalid imports", Items: []string{badimpname}}
		fix.Edits.AddEdit_DeleteLine(d.Loc.FilePath, d.Loc.Pos)
	}
	return
}
