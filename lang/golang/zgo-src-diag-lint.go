package zgo

import (
	"github.com/go-leap/dev"
	"github.com/go-leap/dev/go"
	"github.com/metaleap/zentient"
)

func (me *goDiag) KnownLinters() z.Tools {
	return me.knownTools
}

func (me *goDiag) PrepLintJobs(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
	if pkgjobs := me.onUpdateDiagsPrepPkgJobs(filePaths); len(pkgjobs) > 0 {
		for _, pj := range pkgjobs {
			skippkg := false
			for _, fpath := range pj.Target.(*udevgo.Pkg).GoFilePaths(true) {
				if skippkg = workspaceFiles.HasProbDiags(fpath); skippkg {
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

func (me *goDiag) RunLintJob(job *z.DiagJobLint, workspaceFiles z.WorkspaceFiles) {
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
		fallbackfilepath := func() string { return me.fallbackFilePath(pkg, workspaceFiles) }
		for _, srcref := range msgs {
			srcref.Flag = int(job.Tool.DiagSev)
			job.Yield(me.NewDiagItemFrom(srcref, job.Tool.Name, fallbackfilepath))
		}
	}
}
