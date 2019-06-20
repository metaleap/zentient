package zgo

import (
	"strings"

	"github.com/go-leap/dev"
	"github.com/go-leap/dev/go"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
)

func ensureBuildOrder(dis z.IDiagJobTarget, dat z.IDiagJobTarget) bool {
	return dis.(*udevgo.Pkg).IsSortedPriorToByDeps(dat.(*udevgo.Pkg))
}

func (me *goDiag) OnUpdateBuildDiags(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	if pkgjobs, pkgsbyimp := me.onUpdateDiagsPrepPkgJobs(writtenFilePaths), udevgo.PkgsByImP; len(pkgjobs) > 0 && pkgsbyimp != nil {
		for _, pj := range pkgjobs {
			job := &z.DiagJobBuild{DiagJob: pj, TargetCmp: ensureBuildOrder}
			for _, dependant := range pj.Target.(*udevgo.Pkg).Dependants() {
				if pkgdep := pkgsbyimp[dependant]; pkgdep != nil {
					jobs = append(jobs, &z.DiagJobBuild{DiagJob: z.DiagJob{Target: pkgdep, AffectedFilePaths: pkgdep.GoFilePaths(true)}, TargetCmp: ensureBuildOrder})
				}
			}
			jobs = append(jobs, job)
		}
		for _, job := range jobs {
			// somewhat inelegant-seeming loop prevents accumulation of duplicate build-diags
			for _, dep := range job.Target.(*udevgo.Pkg).Deps {
				if pkgdep := pkgsbyimp[dep]; pkgdep != nil {
					for _, gfp := range pkgdep.GoFilePaths(true) {
						if !ustr.In(gfp, job.AffectedFilePaths...) {
							job.AffectedFilePaths = append(job.AffectedFilePaths, gfp)
						}
					}
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

func (me *goDiag) RunBuildJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) (diags z.DiagItems) {
	numbuilt, progress := 0, z.NewBuildProgress(len(jobs))
	for i := 0; i < progress.NumJobs; i++ {
		progress.AddPkgName(jobs[i].Target.(*udevgo.Pkg).ImportPath)
	}

	for i, pkgjob := range jobs {
		progress.OnJob(i)
		skip, pkg := false, pkgjob.Target.(*udevgo.Pkg)
		if len(progress.Failed) > 0 {
			for _, pdep := range pkg.Deps {
				if skip, _ = progress.Failed[pdep]; skip {
					progress.Skipped[pkg.ImportPath] = true
					break
				}
			}
		}
		if !skip {
			pkgdiags := me.runBuildPkg(pkg, workspaceFiles)
			if pkgjob.Succeeded, diags = len(pkgdiags) == 0, append(diags, pkgdiags...); pkgjob.Succeeded {
				numbuilt++
			} else {
				progress.Failed[pkg.ImportPath] = true
			}
		}
	}
	progress.OnDone()
	if numbuilt > 0 {
		go caddyRunRefreshPkgs()
		if tools.godocdown.Installed {
			for _, pkgjob := range jobs {
				if pkgjob.Succeeded {
					go tools.execGodocdown(pkgjob.Target.(*udevgo.Pkg))
				}
			}
		}
	}
	return
}

func (me *goDiag) FixerUppers() []z.FixerUpper {
	return []z.FixerUpper{me.tryFixImpNotFound, me.tryFixImpMissing}
}

func (*goDiag) tryFixImpMissing(d *z.DiagItem) (fix *z.FixUp) {
	if pkgname, pref, pkgs, idx := "", "undefined: ", udevgo.PkgsByImP, strings.Index(d.Msg, " is not a package"); pkgs != nil {
		if strings.HasPrefix(d.Msg, pref) {
			pkgname = d.Msg[len(pref):]
		} else if idx > 0 {
			pkgname = d.Msg[:idx]
		}
		if pkgname != "" {
			mpkgs := udevgo.PkgsByName(pkgname)
			if pkg := pkgs[ustr.Fewest(mpkgs, "/", ustr.Shortest)]; pkg != nil {
				fix = &z.FixUp{Name: "Add missing imports", Items: []string{pkg.ImportPath}}
				fix.Edits.AddInsert(d.Loc.FilePath, func(srclens *z.SrcLens, set *z.SrcPos) (ins string) {
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

func (*goDiag) tryFixImpNotFound(d *z.DiagItem) (fix *z.FixUp) {
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
		fix.Edits.AddDeleteLine(d.Loc.FilePath, d.Loc.Pos)
	}
	return
}
