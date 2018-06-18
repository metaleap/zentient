package z

import (
	"sort"
	"strings"
	"time"
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

func (this DiagBuildJobs) Len() int               { return len(this) }
func (this DiagBuildJobs) Swap(i int, j int)      { this[i], this[j] = this[j], this[i] }
func (this DiagBuildJobs) Less(i int, j int) bool { return this[i].IsSortedPriorTo(this[j]) }

func (this DiagBuildJobs) withoutDuplicates() (nu DiagBuildJobs) {
	nu = make(DiagBuildJobs, 0, len(this))
	done := make(map[string]bool, len(this))
	for _, job := range this {
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

func (this *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool {
	c := cmp.(*DiagJobBuild)
	if this.TargetCmp != nil {
		return this.TargetCmp(this.Target, c.Target)
	}
	if sortish, _ := this.Target.(ISortable); sortish != nil {
		return sortish.IsSortedPriorTo(c.Target)
	}
	return false
}

type BuildProgress struct {
	NumJobs   int
	StartTime time.Time
	Failed    map[string]bool
	Skipped   map[string]bool
	PkgNames  []string
}

func NewBuildProgress(numJobs int) *BuildProgress {
	return &BuildProgress{NumJobs: numJobs, StartTime: time.Now(), PkgNames: make([]string, 0, numJobs), Failed: make(map[string]bool, numJobs), Skipped: make(map[string]bool, numJobs)}
}

func (this *BuildProgress) AddPkgName(pkgName string) {
	this.PkgNames = append(this.PkgNames, pkgName)
}

func (this *BuildProgress) OnDone() {
	CaddyBuildOnDone(this.Failed, this.Failed, this.PkgNames, time.Since(this.StartTime))
}

func (this *BuildProgress) OnJob(i int) {
	CaddyBuildOnRunning(this.NumJobs, i, this.String())
}

func (this *BuildProgress) String() string {
	return strings.Join(this.PkgNames, "\n")
}

func (*DiagBase) FixerUppers() []FixerUpper { return nil }

func (this *DiagBase) fixUps(diags DiagItems) {
	fixers := this.Impl.FixerUppers()
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

func (this *DiagBase) UpdateBuildDiagsAsNeeded(workspaceFiles WorkspaceFiles, writtenFiles []string) {
	if jobs := this.Impl.OnUpdateBuildDiags(writtenFiles).withoutDuplicates(); len(jobs) > 0 {
		sort.Sort(jobs)
		for _, job := range jobs {
			job.forgetPrevDiags(nil, false, workspaceFiles)
		}
		go this.send(workspaceFiles, true)
		diagitems := this.Impl.RunBuildJobs(jobs, workspaceFiles)
		diagitems.propagate(false, true, workspaceFiles)
		if len(diagitems) > 0 {
			go this.fixUps(diagitems)
		}
	}
	go this.send(workspaceFiles, false)
}
