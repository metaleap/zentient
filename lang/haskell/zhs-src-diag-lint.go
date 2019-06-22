package zhs

import (
	"github.com/go-leap/dev/hs"
	"github.com/metaleap/zentient"
)

func (me *hsDiag) KnownLinters() z.Tools {
	return me.knownTools
}

func (me *hsDiag) PrepLintJobs(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
	for _, dt := range diagTools {
		jobs = append(jobs, &z.DiagJobLint{
			DiagJob: z.DiagJob{AffectedFilePaths: filePaths, Target: filePaths},
			Tool:    dt,
		})
	}
	return
}

func (me *hsDiag) RunLintJob(job *z.DiagJobLint, workspaceFiles z.WorkspaceFiles) {
	if job.Tool == tools.hlint {
		fpaths := job.Target.([]string)
		for _, srcref := range udevhs.LintHlint(fpaths) {
			srcref.Flag = int(job.Tool.DiagSev)
			job.Yield(me.DiagBase.NewDiagItemFrom(srcref, job.Tool.Name, func() string { return fpaths[0] }))
		}
	}
}
