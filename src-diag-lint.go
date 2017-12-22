package z

import (
	"strings"

	"github.com/metaleap/go-util/slice"
)

type DiagLintJobs []*DiagJobLint

type DiagJobLint struct {
	DiagJob
	Tool     *Tool
	lintChan chan *DiagItem
}

func (me *DiagJobLint) Done()                { me.lintChan <- nil }
func (me *DiagJobLint) Yield(diag *DiagItem) { me.lintChan <- diag }

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool, onlyFilePaths ...string) {
	if nonautos, diagtools := !autos, me.knownDiags(autos); len(diagtools) > 0 {
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
			me.updateLintDiags(workspaceFiles, diagtools, !autos, filepaths)
		}
	}
	go me.send(false)
}

func (me *DiagBase) updateLintDiags(workspaceFiles WorkspaceFiles, diagTools Tools, sticky bool, filePaths []string) {
	if jobs := me.Impl.OnUpdateLintDiags(workspaceFiles, diagTools, filePaths); len(jobs) > 0 {
		numjobs, numdone, await, descs := 0, 0, make(chan *DiagItem), make([]string, len(jobs))
		for i, job := range jobs {
			numjobs++
			job.lintChan = await
			go me.Impl.RunLintJob(job)
			job.forgetPrevDiags(diagTools, workspaceFiles)
			descs[i] = job.Tool.Name + " " + job.String()
		}
		go send(&ipcResp{IpcID: IPCID_SRCDIAG_STARTED, ObjSnapshot: strings.Join(descs, "\n")})
		var diagitems DiagItems
		for diagitem := range await {
			if diagitem != nil {
				diagitems = append(diagitems, diagitem)
			} else if numdone++; numdone == numjobs {
				break
			}
		}
		diagitems.propagate(true, sticky, workspaceFiles)
		go send(&ipcResp{IpcID: IPCID_SRCDIAG_FINISHED, ObjSnapshot: ""})
	}
}
