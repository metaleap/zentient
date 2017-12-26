package z

import (
	"sort"
)

type DiagBuildJobs []*DiagJobBuild

func (me DiagBuildJobs) Len() int               { return len(me) }
func (me DiagBuildJobs) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me DiagBuildJobs) Less(i int, j int) bool { return me[i].IsSortedPriorTo(me[j]) }

func (me DiagBuildJobs) WithoutDuplicates() (nu DiagBuildJobs) {
	nu = make(DiagBuildJobs, 0, len(me))
	done := make(map[string]bool, len(me))
	for _, job := range me {
		if s := job.String(); !done[s] {
			done[s], nu = true, append(nu, job)
		}
	}
	return
}

type DiagJobBuild struct {
	DiagJob
	TargetCmp func(IDiagJobTarget, IDiagJobTarget) bool
	Succeeded bool
	diags     DiagItems
}

func (me *DiagJobBuild) Yield(diag *DiagItem) { me.diags = append(me.diags, diag) }

func (me *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool {
	c := cmp.(*DiagJobBuild)
	if me.TargetCmp != nil {
		return me.TargetCmp(me.Target, c.Target)
	}
	return me.Target.IsSortedPriorTo(c.Target)
}

func (me *DiagBase) UpdateBuildDiagsAsNeeded(workspaceFiles WorkspaceFiles, writtenFiles []string) {
	if jobs := me.Impl.OnUpdateBuildDiags(writtenFiles).WithoutDuplicates(); len(jobs) > 0 {
		sort.Sort(jobs)
		for _, job := range jobs {
			job.WorkspaceFiles = workspaceFiles
			job.forgetPrevDiags(nil, false, workspaceFiles)
		}
		go me.send(workspaceFiles, true)
		diagitems := me.Impl.RunBuildJobs(jobs)
		diagitems.propagate(false, true, workspaceFiles)
	}
	go me.send(workspaceFiles, false)
}
