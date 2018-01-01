package z

import (
	"sort"
)

type FixerUpper func(*DiagItem) *FixUp

type IDiagBuild interface {
	FixUps(DiagItems, func([]*FixUp))
	OnUpdateBuildDiags([]string) DiagBuildJobs
	RunBuildJobs(DiagBuildJobs) DiagItems
	UpdateBuildDiagsAsNeeded(WorkspaceFiles, []string)
}

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

type FixUpsByFile map[string][]*FixUp

type FixUp struct {
	Name      string
	Item      string
	Mod       SrcLens
	ModDropLn bool
}

func (me *DiagJobBuild) Yield(diag *DiagItem) { me.diags = append(me.diags, diag) }

func (me *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool {
	c := cmp.(*DiagJobBuild)
	if me.TargetCmp != nil {
		return me.TargetCmp(me.Target, c.Target)
	}
	return me.Target.IsSortedPriorTo(c.Target)
}

func (me *DiagBase) onFixUps(all []*FixUp) {
	if len(all) > 0 {
		fixups := make(FixUpsByFile, len(all))
		for _, fix := range all {
			fixups[fix.Mod.FilePath] = append(fixups[fix.Mod.FilePath], fix)
		}
		send(&ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: &DiagResp{FixUps: fixups, LangID: Lang.ID}})
	}
}

func (*DiagBase) FixUps(DiagItems, func([]*FixUp)) {
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
		if len(diagitems) > 0 {
			go me.Impl.FixUps(diagitems, me.onFixUps)
		}
	}
	go me.send(workspaceFiles, false)
}
