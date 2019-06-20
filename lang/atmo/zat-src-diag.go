package zat

import (
	"text/scanner"

	"github.com/metaleap/atmo"
	"github.com/metaleap/zentient"
)

var diag atDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type atDiag struct {
	z.DiagBase
	errDiags z.DiagItems
}

func (me *atDiag) onSomeKitsReprocessed() {
	var diags z.DiagItems
	for _, kit := range Ctx.Kits.All {
		errs2srcs := make(map[error][]byte, 2)
		for _, err := range kit.Errors(errs2srcs) {
			errdiag := &z.DiagItem{Msg: err.Error(), Misc: make([]interface{}, 3)}
			if e, _ := err.(*atmo.Error); e != nil {
				errdiag.Misc[0] = e
				errdiag.Msg, errdiag.Cat = e.Msg(), e.Cat().String()
				if e.Cat() == atmo.ErrCatUnreachable {
					errdiag.Tags = []int{1}
				}
				if pos, src := e.Pos(), string(errs2srcs[err]); pos != nil {
					errdiag.Misc[1], errdiag.Misc[2] = pos, src
					errdiag.Loc.FilePath, errdiag.Loc.Pos = pos.Filename, &z.SrcPos{}
					if errdiag.Loc.Pos.Ln, errdiag.Loc.Pos.Col = pos.Line, pos.Column; len(src) > 0 {
						errdiag.Loc.Pos.SetRune1OffFromByte0Off(pos.Offset, src)
					} else if pos.Line < 1 || pos.Column < 1 {
						errdiag.Loc.Pos.Off = 1 + pos.Offset
					}
				}
			}
			diags = append(diags, errdiag)
		}
	}
	me.errDiags = diags
}

func (*atDiag) KnownLinters() z.Tools {
	return nil
}

func (*atDiag) ShouldOnFileOpen() bool { return true }

func (me *atDiag) OnUpdateBuildDiags(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) z.DiagBuildJobs {
	var job z.DiagJobBuild
	job.AffectedFilePaths = Ctx.Kits.All.SrcFilePaths()
	return z.DiagBuildJobs{&job}
}

func (me *atDiag) RunBuildJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) z.DiagItems {
	for _, errdiag := range me.errDiags {
		if errdiag.Loc.Range == nil && errdiag.Loc.Pos != nil {
			if e, _ := errdiag.Misc[0].(*atmo.Error); e != nil {
				if pos, _ := errdiag.Misc[1].(*scanner.Position); pos != nil {
					if src, ok := errdiag.Misc[2].(string); ok {
						isfileopen := workspaceFiles.IsOpen(pos.Filename)
						if l := e.Len(); l > 1 && len(src) > 0 && isfileopen {
							errdiag.Loc.Range = &z.SrcRange{}
							errdiag.Loc.Range.Start.SetRune1OffFromByte0Off(pos.Offset, src)
							errdiag.Loc.Range.End.SetRune1OffFromByte0Off(pos.Offset+l, src)
						}
					}
				}
			}
		}
	}
	return me.errDiags
}

func (*atDiag) OnUpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {

	return
}

func (*atDiag) RunLintJob(job *z.DiagJobLint, workspaceFiles z.WorkspaceFiles) {
}
