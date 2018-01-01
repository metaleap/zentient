package z

import (
	"time"

	"github.com/metaleap/go-util/slice"
)

var (
	cfgLintStickiness *Setting
)

type IDiagLint interface {
	KnownLinters() Tools
	OnUpdateLintDiags(WorkspaceFiles, Tools, []string) DiagLintJobs
	RunLintJob(*DiagJobLint)
	UpdateLintDiagsIfAndAsNeeded(WorkspaceFiles, bool, ...string)
}

func init() {
	if cfgLintStickiness == nil {
		cfgLintStickiness = &Setting{Id: "cfgLintStickiness", Title: "Sticky-Lints Level", ValDef: uint64(0),
			Desc: "Lints shown even for closed files ➜ 0: errors only — 1: and warnings — 2: and infos — 3: and hints."}
		cfgLintStickiness.OnChanging = func(nv interface{}) {
			if l, ok := nv.(uint64); (!ok) || l > 3 {
				panic("Wanted: a level of at-least 0 and at-most 3.")
			}
		}
		cfgLintStickiness.OnChanged = func(_ interface{}) {
			val, workspacefiles := cfgLintStickiness.ValUInt(), Lang.Workspace.Files()
			for _, f := range workspacefiles {
				for _, d := range f.Diags.Lint.Items {
					d.StickyAuto = uint64(d.Loc.Flag) <= (val)
				}
			}
			go Lang.Diag.send(workspacefiles, false)
		}
	}
}

type DiagLintJobs []*DiagJobLint

type DiagJobLint struct {
	DiagJob
	Tool        *Tool
	lintChan    chan *DiagItem
	timeStarted time.Time
	timeTaken   time.Duration
}

func (me *DiagJobLint) Yield(diag *DiagItem) { me.lintChan <- diag }
func (me *DiagJobLint) Done() {
	me.timeTaken = time.Since(me.timeStarted)
	me.lintChan <- nil
}

func (me *DiagBase) knownLinters(auto bool) (diags Tools) {
	for _, dt := range me.Impl.KnownLinters() {
		if dt.IsInAutoDiags() == auto {
			diags = append(diags, dt)
		}
	}
	return
}

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool, onlyFilePaths ...string) {
	if nonautos, diagtools := !autos, me.knownLinters(autos); len(diagtools) > 0 {
		var filepaths []string
		for _, f := range workspaceFiles {
			if autos && len(f.Diags.Build.Items) > 0 {
				return
			} else if f.IsOpen && (nonautos || !f.Diags.AutoLintUpToDate) {
				if len(onlyFilePaths) == 0 || uslice.StrHas(onlyFilePaths, f.Path) {
					filepaths = append(filepaths, f.Path)
				}
			}
		}
		if len(filepaths) > 0 {
			me.updateLintDiags(workspaceFiles, diagtools, autos, filepaths).propagate(true, nonautos, workspaceFiles)
		}
	}
	go me.send(workspaceFiles, false)
}

func (me *DiagBase) updateLintDiags(workspaceFiles WorkspaceFiles, diagTools Tools, autos bool, filePaths []string) (diagitems DiagItems) {
	jobs := me.Impl.OnUpdateLintDiags(workspaceFiles, diagTools, filePaths)
	if numjobs, nonautos := len(jobs), !autos; numjobs > 0 {
		numdone, await, descs := 0, make(chan *DiagItem), make([]string, numjobs)
		for _, job := range jobs { // separate loop from the go-routines below to prevent concurrent-map-read+write as forgetPrevDiags() calls workspaceFiles.ensure()
			job.WorkspaceFiles = workspaceFiles
			job.forgetPrevDiags(diagTools, autos, workspaceFiles)
		}
		go me.send(workspaceFiles, false)
		if nonautos {
			onRunManuallyAlreadyCurrentlyRunning = true
		}
		for i, job := range jobs {
			job.lintChan, job.timeStarted = await, time.Now()
			go me.Impl.RunLintJob(job)
			descs[i] = job.Tool.Name + " ➜ " + job.String()
		}
		send(&ipcResp{IpcID: IPCID_SRCDIAG_STARTED, ObjSnapshot: descs})
		for diagitem := range await {
			if diagitem != nil {
				diagitems = append(diagitems, diagitem)
			} else if numdone++; numdone == numjobs {
				break
			}
		}
		for i, job := range jobs {
			descs[i] += Strf(" \n\t\t%s", job.timeTaken)
		}
		go send(&ipcResp{IpcID: IPCID_SRCDIAG_FINISHED, ObjSnapshot: descs})
		if nonautos {
			onRunManuallyAlreadyCurrentlyRunning = false
		}
	}
	return
}
