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

func (me *hsDiag) OnUpdateBuildDiags(writtenFilePaths []string) (jobs z.DiagBuildJobs) {
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

func (me *hsDiag) RunBuildJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) (diags z.DiagItems) {
	for _, job := range jobs {
		stackyamldirpath := job.Target.(string)
		stackyamlfilepath := filepath.Join(stackyamldirpath, "stack.yaml")
		cmdargs := append(append([]string{"build"}, udevhs.StackArgs...), udevhs.StackArgsBuild...)
		stdout, stderr, err := urun.CmdExecIn(stackyamldirpath, "stack", cmdargs...)
		if err != nil {
			stderr = err.Error()
		}
		lns := ustr.Split(stderr+"\n"+ustr.Trim(stdout), "\n")

		lnstackwarns, _w := []string{}, "Warning: "+stackyamlfilepath+": "
		for i := 0; i < len(lns); i++ {
			if ustr.Pref(lns[i], _w) {
				lnstackwarns = append(lnstackwarns, lns[i][len(_w):])
				i, lns = i-1, append(lns[:i], lns[i+1:]...)
			}
		}

		for len(lns) > 0 && lns[0] == "" {
			lns = lns[1:]
		}
		wasnoop := len(lns) == 0
		if len(lns) > 3 && ustr.Pref(lns[0], "Copying from ") && len(lns[1]) == 0 && ustr.Pref(lns[2], "Copied executables to ") {
			wasnoop = true
			for i := 3; i < len(lns); i++ {
				if !ustr.Pref(lns[i], "- ") {
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
				diags = append(diags, &z.DiagItem{
					Cat: "stack", Msg: ustr.Join(lns[1:], "\n"), Loc: z.SrcLoc{
						FilePath: fpath, Flag: int(z.DIAG_SEV_ERR), Pos: &z.SrcPos{Col: 1, Ln: 1},
					},
				})
				return
			}

			if _e := "Error: "; ustr.Pref(lns[0], _e) {
				diags = append(diags, &z.DiagItem{
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
					if i := ustr.Pos(cur.Msg, "]"); ustr.Pref(cur.Msg, "[") && i > 0 {
						toolname = toolname + "  » " + cur.Msg[1:i]
						cur.Msg = ustr.Trim(cur.Msg[i+1:])
					}
					diag := me.DiagBase.NewDiagItemFrom(cur, toolname, func() string { return stackyamlfilepath })
					diags = append(diags, diag)
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
	}
	return
}
