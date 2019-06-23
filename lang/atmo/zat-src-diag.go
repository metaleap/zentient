package zat

import (
	"sync"

	"github.com/metaleap/atmo"
	"github.com/metaleap/zentient"
)

var diag atmoDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type atmoDiag struct {
	z.DiagBase
	sync.Mutex

	errDiags z.DiagItems
}

func (me *atmoDiag) updateFromErrs(_ bool) {
	var errdiags z.DiagItems
	for _, kit := range Ctx.Kits.All {
		errs2srcs := make(map[error][]byte, 4)
		for _, err := range kit.Errors(errs2srcs) {
			errdiag := &z.DiagItem{Msg: err.Error()}
			if e, _ := err.(*atmo.Error); e != nil {
				errdiag.Msg, errdiag.Cat = e.Msg(), e.Cat().String()
				if e.Cat() == atmo.ErrCatUnreachable {
					errdiag.Tags = []int{1}
				}
				if pos, src := e.Pos(), string(errs2srcs[err]); pos != nil {
					errdiag.Loc.FilePath, errdiag.Loc.Pos = pos.Filename, &z.SrcPos{}
					if errdiag.Loc.Pos.Ln, errdiag.Loc.Pos.Col = pos.Line, pos.Column; len(src) > 0 {
						errdiag.Loc.Pos.SetRune1OffFromByte0Off(pos.Offset, src)
					} else if pos.Line < 1 || pos.Column < 1 {
						errdiag.Loc.Pos.Off = 1 + pos.Offset
					}
					if errlen := e.Len(); errlen > 1 && len(src) > 0 {
						errdiag.Loc.Range = &z.SrcRange{}
						errdiag.Loc.Range.Start.SetRune1OffFromByte0Off(pos.Offset, src)
						errdiag.Loc.Range.End.SetRune1OffFromByte0Off(pos.Offset+errlen, src)
					}
				}
			}
			errdiags = append(errdiags, errdiag)
		}
	}
	me.Lock()
	me.errDiags = errdiags
	me.Unlock()
}

func (*atmoDiag) KnownLinters() z.Tools { return nil }

func (me *atmoDiag) PrepIssueJobs(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) z.DiagBuildJobs {
	var job z.DiagJobBuild
	job.AffectedFilePaths = Ctx.Kits.All.SrcFilePaths()
	return z.DiagBuildJobs{&job}
}

func (me *atmoDiag) RunIssueJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) (errdiags z.DiagItems) {
	me.Lock()
	errdiags = me.errDiags
	me.Unlock()
	return
}

func (*atmoDiag) PrepLintJobs(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
	return
}

func (*atmoDiag) RunLintJob(job *z.DiagJobLint, workspaceFiles z.WorkspaceFiles) {}
