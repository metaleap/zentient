package zgo

import (
	"strings"
	"time"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/zentient"
)

func ensureBuildOrder(dis z.IDiagJobTarget, dat z.IDiagJobTarget) bool {
	return dis.(*udevgo.Pkg).IsSortedPriorToByDeps(dat.(*udevgo.Pkg))
}

func (me *goDiag) OnUpdateBuildDiags(writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if pkgjobs := me.onUpdateDiagsPrepPkgJobs(writtenFilePaths); len(pkgjobs) > 0 {
		for _, pj := range pkgjobs {
			job := &z.DiagJobBuild{DiagJob: pj, TargetCmp: ensureBuildOrder}
			for _, dependant := range pj.Target.(*udevgo.Pkg).Dependants() {
				if pkgdep := udevgo.PkgsByImP[dependant]; pkgdep != nil {
					jobs = append(jobs, &z.DiagJobBuild{DiagJob: z.DiagJob{Target: pkgdep, AffectedFilePaths: pkgdep.GoFilePaths()}, TargetCmp: ensureBuildOrder})
				}
			}
			for _, dep := range pj.Target.(*udevgo.Pkg).Deps {
				// this sub-optimal loop in practice unneeded in ~90+% of usage but sometimes *is* to prevent ugly diag duplications of existing unaddressed dep diags
				// ie: you remove an issue from main, there remains one in its imported dep, but we rely on `go install` for that one to rebuild
				// (there was no build-on-save signal for the dep, just the main) --- we could also do lengthy "duplicate check"s on all diags but
				// mildly cleaner to mark all go files of all dependencies as "affected" aka "may-produce-diags", meaning we clear those deps too (not just dependants as done above)
				if pkgdep := udevgo.PkgsByImP[dep]; pkgdep != nil {
					for _, gfp := range pkgdep.GoFilePaths() {
						if !uslice.StrHas(job.AffectedFilePaths, gfp) {
							job.AffectedFilePaths = append(job.AffectedFilePaths, gfp)
						}
					}
				}
			}
			jobs = append(jobs, job)
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
	return []z.FixerUpper{me.tryFixImpNotFound, me.tryFixImpMissing}
}

func (me *goDiag) tryFixImpMissing(d *z.DiagItem) (fix *z.FixUp) {
	if pkgname, pref, pkgs, i := "", "undefined: ", udevgo.PkgsByImP, strings.Index(d.Msg, " is not a package"); pkgs != nil {
		if strings.HasPrefix(d.Msg, pref) {
			pkgname = d.Msg[len(pref):]
		} else if i > 0 {
			pkgname = d.Msg[:i]
		}
		if mpkgs := []string{}; pkgname != "" {
			for _, p := range pkgs {
				if p.Name == pkgname {
					mpkgs = append(mpkgs, p.ImportPath)
				}
			}
			if pkg := udevgo.PkgsByImP[uslice.StrShortest(mpkgs)]; pkg != nil {
				fix = &z.FixUp{Name: "Add missing imports", Items: []string{pkg.ImportPath}}
				fix.Edits.AddEdit_Insert(d.Loc.FilePath, func(srclens *z.SrcLens, set *z.SrcPos) (ins string) {
					if i := strings.Index(srclens.Txt, "\nimport (\n"); i > 0 {
						idot, j := strings.IndexRune(pkg.ImportPath, '.'), strings.Index(srclens.Txt[i:], "\n)\n")
						if ins = z.Strf("\t%#v\n", pkg.ImportPath); j > 0 && idot >= 0 && idot < strings.IndexRune(pkg.ImportPath, '/') {
							i = i + j
						} else {
							i += 9
						}
						set.Off = srclens.Rune1OffsetForByte0Offset(i + 1)
					}
					return
				})
			}
		}
	}
	return
}

func (me *goDiag) tryFixImpNotFound(d *z.DiagItem) (fix *z.FixUp) {
	var badimpname string
	if pref1, i1 := "cannot find package \"", strings.Index(d.Msg, "\" in any of:"); strings.HasPrefix(d.Msg, pref1) && i1 > len(pref1) {
		//cannot find package "foo" in any of:
		badimpname = d.Msg[:i1][len(pref1):]
	} else if pref2, i2 := "invalid import path: \"", strings.LastIndex(d.Msg, "\""); strings.HasPrefix(d.Msg, pref2) && i2 > len(pref2) {
		//invalid import path: "moyhoar hoaryya baddabam fam"
		badimpname = d.Msg[:i2][len(pref2):]
	} else if pref3, i3 := "imported and not used: \"", strings.LastIndex(d.Msg, "\""); strings.HasPrefix(d.Msg, pref3) && i3 > len(pref3) {
		badimpname = d.Msg[:i3][len(pref3):]
	}
	if badimpname != "" {
		fix = &z.FixUp{Name: "Removes invalid imports", Items: []string{badimpname}}
		fix.Edits.AddEdit_DeleteLine(d.Loc.FilePath, d.Loc.Pos)
	}
	return
}
