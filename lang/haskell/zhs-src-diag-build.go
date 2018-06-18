package zhs

import (
	"path/filepath"

	"github.com/go-leap/dev"
	"github.com/go-leap/dev/hs"
	"github.com/go-leap/fs"
	"github.com/go-leap/run"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
)

func (this *hsDiag) OnUpdateBuildDiags(writtenFilePaths []string) (jobs z.DiagBuildJobs) {
	stackyamlfilepaths := map[string]bool{}
	for _, fp := range writtenFilePaths {
		if stackyamlfilepath := ufs.Locate(fp, "stack.yaml"); stackyamlfilepath != "" {
			stackyamlfilepaths[stackyamlfilepath] = true
		}
	}
	for stackyamlfilepath := range stackyamlfilepaths {
		stackyamldirpath, job := filepath.Dir(stackyamlfilepath), &z.DiagJobBuild{}
		job.Target = stackyamldirpath
		job.AffectedFilePaths = ufs.AllFilePathsIn(stackyamldirpath, ".stack-work", "*.hs")
		jobs = append(jobs, job)
	}
	return
}

func (this *hsDiag) RunBuildJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) (diags z.DiagItems) {
	progress := z.NewBuildProgress(len(jobs))
	for i := 0; i < progress.NumJobs; i++ {
		progress.AddPkgName(jobs[i].Target.(string))
	}

	for i, job := range jobs {
		progress.OnJob(i)
		stackyamldirpath := job.Target.(string)
		stackyamlfilepath := filepath.Join(stackyamldirpath, "stack.yaml")
		haderrs, cmdargs := false, append(append([]string{"build"}, udevhs.StackArgs...), udevhs.StackArgsBuild...)
		stdout, stderr, err := urun.CmdExecIn(stackyamldirpath, "stack", cmdargs...)
		if err != nil {
			stderr = err.Error()
		}
		lns := ustr.Split(stderr+"\n"+ustr.Trim(stdout), "\n")

		lnstackwarns, _w := []string{}, "Warning: "+stackyamlfilepath+": "
		for j := 0; j < len(lns); j++ {
			if ustr.Pref(lns[j], _w) {
				lnstackwarns = append(lnstackwarns, lns[j][len(_w):])
				j, lns = j-1, append(lns[:j], lns[j+1:]...)
			}
		}

		for len(lns) > 0 && lns[0] == "" {
			lns = lns[1:]
		}
		wasnoop := len(lns) == 0
		if len(lns) > 3 && ustr.Pref(lns[0], "Copying from ") && len(lns[1]) == 0 && ustr.Pref(lns[2], "Copied executables to ") {
			wasnoop = true
			for j := 3; j < len(lns); j++ {
				if !ustr.Pref(lns[j], "- ") {
					wasnoop = false
				}
			}
		}

		if !wasnoop {
			for _, wln := range lnstackwarns {
				if parts := ustr.Split(wln, ":"); len(parts) > 1 {
					diags = append(diags, &z.DiagItem{
						Cat: "stack", Msg: ustr.Join(parts[1:], ":"), Loc: z.SrcLoc{
							FilePath: parts[0], Flag: int(z.DIAG_SEV_WARN), Pos: &z.SrcPos{Col: 1, Ln: 1},
						},
					})
				}
			}

			if _p := "Could not parse '"; ustr.Pref(lns[0], _p) {
				fpath := ustr.TrimR(lns[0][len(_p):], "':")
				haderrs, diags = true, append(diags, &z.DiagItem{
					Cat: "stack", Msg: ustr.Join(lns[1:], "\n"), Loc: z.SrcLoc{
						FilePath: fpath, Flag: int(z.DIAG_SEV_ERR), Pos: &z.SrcPos{Col: 1, Ln: 1},
					},
				})
				return
			}

			if _e := "Error: "; ustr.Pref(lns[0], _e) {
				haderrs, diags = true, append(diags, &z.DiagItem{
					Cat: "stack", Msg: ustr.Join(lns, "\n")[len(_e):], Loc: z.SrcLoc{
						FilePath: stackyamlfilepath, Flag: int(z.DIAG_SEV_ERR), Pos: &z.SrcPos{Col: 1, Ln: 1},
					},
				})
				return
			}

			var cur *udev.SrcMsg
			addlastcur := func() {
				if cur != nil {
					cur.Msg = ustr.Trim(cur.Msg)
					toolname := "ghc"
					if !ustr.Suff(cur.Ref, ".hs") {
						toolname = "stack"
					}
					if j := ustr.Pos(cur.Msg, "]"); ustr.Pref(cur.Msg, "[") && j > 0 {
						toolname = toolname + "  » " + cur.Msg[1:j]
						cur.Msg = ustr.Trim(cur.Msg[j+1:])
					}
					diag := this.DiagBase.NewDiagItemFrom(cur, toolname, func() string { return stackyamlfilepath })
					if diags = append(diags, diag); diag.Loc.Flag == int(z.DIAG_SEV_ERR) {
						haderrs = true
					}
					cur = nil
				}
			}
			for _, ln := range lns {
				if len(ln) > 0 {
					if msg := udev.SrcMsgFromLn(ln); msg != nil && ustr.Pref(msg.Ref, stackyamldirpath+string(filepath.Separator)) && ufs.IsFile(msg.Ref) {
						addlastcur()
						cur = msg
						if ustr.Pref(cur.Msg, "error: ") {
							cur.Msg = cur.Msg[7:]
						}
						if ustr.Pref(cur.Msg, "warning: ") {
							cur.Flag = int(z.DIAG_SEV_WARN)
							cur.Msg = cur.Msg[9:]
						}
					} else if ustr.Pref(ln, "    ") && cur != nil {
						cur.Msg += ("\n" + ln)
					} else {
						addlastcur()
					}
				}
			}
			addlastcur()
		}
		if haderrs {
			progress.Failed[stackyamldirpath] = true
		}
	}
	progress.OnDone()
	return
}
