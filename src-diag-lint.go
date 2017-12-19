package z

type DiagLintJobs []*DiagJobLint

type DiagJobLint struct {
	DiagJob
	Tool     *Tool
	lintChan chan *DiagItem
}

func (me *DiagJobLint) Done()                { me.lintChan <- nil }
func (me *DiagJobLint) Yield(diag *DiagItem) { me.lintChan <- diag }

func (me *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool) {
	if diagtools := me.knownDiags(autos); len(diagtools) > 0 {
		var filepaths []string
		for _, f := range workspaceFiles {
			if f != nil && f.IsOpen && !f.Diags.Lint.UpToDate {
				filepaths = append(filepaths, f.Path)
			}
		}
		if len(filepaths) > 0 {
			me.updateLintDiags(workspaceFiles, diagtools, filepaths)
		}
	}
	go me.send(false)
}

func (me *DiagBase) updateLintDiags(workspaceFiles WorkspaceFiles, diagTools Tools, filePaths []string) {
	if jobs := me.Impl.OnUpdateLintDiags(workspaceFiles, diagTools, filePaths); len(jobs) > 0 {
		numjobs, numdone, await := 0, 0, make(chan *DiagItem)
		for _, job := range jobs {
			numjobs++
			job.lintChan = await
			go me.Impl.RunLintJob(job)
			job.forgetAndMarkUpToDate(diagTools, workspaceFiles)
		}

		var diagitems DiagItems
		for diagitem := range await {
			if diagitem != nil {
				diagitems = append(diagitems, diagitem)
			} else if numdone++; numdone == numjobs {
				break
			}
		}
		diagitems.propagate(true, workspaceFiles)
	}
}
