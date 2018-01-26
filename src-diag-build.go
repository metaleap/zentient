package z

import (
	"sort"
)

type FixerUpper func(*DiagItem) *FixUp

type fixUpsByFile map[string][]*FixUp

type FixUp struct {
	Name  string
	Items []string
	Edits SrcModEdits
}

type fixUps struct {
	FilePath string
	Desc     map[string][]string
	Edits    SrcModEdits
	Dropped  []srcModEdit
}

type IDiagBuild interface {
	FixerUppers() []FixerUpper
	OnUpdateBuildDiags([]string) DiagBuildJobs
	RunBuildJobs(DiagBuildJobs, WorkspaceFiles) DiagItems
	UpdateBuildDiagsAsNeeded(WorkspaceFiles, []string)
}

type DiagBuildJobs []*DiagJobBuild

func (me DiagBuildJobs) Len() int               { return len(me) }
func (me DiagBuildJobs) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me DiagBuildJobs) Less(i int, j int) bool { return me[i].IsSortedPriorTo(me[j]) }

func (me DiagBuildJobs) withoutDuplicates() (nu DiagBuildJobs) {
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
}

func (me *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool {
	c := cmp.(*DiagJobBuild)
	if me.TargetCmp != nil {
		return me.TargetCmp(me.Target, c.Target)
	}
	if sortish, _ := me.Target.(ISortable); sortish != nil {
		return sortish.IsSortedPriorTo(c.Target)
	}
	return false
}

func (*DiagBase) FixerUppers() []FixerUpper { return nil }

func (me *DiagBase) fixUps(diags DiagItems) {
	fixers := me.Impl.FixerUppers()
	if len(fixers) == 0 {
		return
	}
	fixupsbyfile := fixUpsByFile{}
	for _, d := range diags {
		for _, f := range fixers {
			if fixup := f(d); fixup != nil && len(fixup.Edits) > 0 {
				fixupsbyfile[d.Loc.FilePath] = append(fixupsbyfile[d.Loc.FilePath], fixup)
			}
		}
	}
	if len(fixupsbyfile) > 0 {
		dr := &diagResp{LangID: Lang.ID, FixUps: make([]*fixUps, 0, len(fixupsbyfile))}
		for filepath, filefixups := range fixupsbyfile {
			fixups := &fixUps{FilePath: filepath, Desc: map[string][]string{}}
			for _, fixup := range filefixups {
				fixups.Desc[fixup.Name] = append(fixups.Desc[fixup.Name], fixup.Items...)
				fixups.Edits = append(fixups.Edits, fixup.Edits...)
			}
			if fixups.Dropped = fixups.Edits.dropConflictingEdits(); fixups.Dropped == nil { // be nice to the client-side here..
				fixups.Dropped = []srcModEdit{}
			}
			sort.Sort(fixups.Edits)
			dr.FixUps = append(dr.FixUps, fixups)
		}
		send(&ipcResp{IpcID: IPCID_SRCDIAG_PUB, SrcDiags: dr})
	}
}

func (me *DiagBase) UpdateBuildDiagsAsNeeded(workspaceFiles WorkspaceFiles, writtenFiles []string) {
	if jobs := me.Impl.OnUpdateBuildDiags(writtenFiles).withoutDuplicates(); len(jobs) > 0 {
		sort.Sort(jobs)
		for _, job := range jobs {
			job.forgetPrevDiags(nil, false, workspaceFiles)
		}
		go me.send(workspaceFiles, true)
		diagitems := me.Impl.RunBuildJobs(jobs, workspaceFiles)
		diagitems.propagate(false, true, workspaceFiles)
		if len(diagitems) > 0 {
			go me.fixUps(diagitems)
		}
	}
	go me.send(workspaceFiles, false)
}
