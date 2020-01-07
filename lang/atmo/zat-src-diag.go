package zat

// import (
// 	"sync"

// 	"github.com/metaleap/atmo/0ld"
// 	"github.com/metaleap/atmo/0ld/session"
// 	"github.com/metaleap/zentient"
// )

// var diag atmoDiag

// func init() {
// 	diag.Impl, z.Lang.Diag = &diag, &diag
// }

// type atmoDiag struct {
// 	z.DiagBase
// 	sync.Mutex

// 	errDiags z.DiagItems
// }

// func (me *atmoDiag) updateFromErrs(ctx *atmosess.Ctx, hadFreshErrs bool) {
// 	var errdiags z.DiagItems
// 	for _, kit := range ctx.Kits.All {
// 		errs2srcs := make(map[*atmo.Error][]byte, 4)
// 		for _, err := range kit.Errors(errs2srcs) {
// 			errdiag := &z.DiagItem{Cat: err.CodeAndCat(), Msg: err.Msg()}
// 			if err.Cat() == atmo.ErrCatUnreachable {
// 				errdiag.Tags = []int{1}
// 			}
// 			if pos, src := err.Pos(), errs2srcs[err]; pos != nil {
// 				errdiag.Loc.FilePath, errdiag.Loc.Pos = pos.FilePath, &z.SrcPos{}
// 				if errdiag.Loc.Pos.Ln, errdiag.Loc.Pos.Col = pos.Ln1, pos.Col1; len(src) > 0 {
// 					errdiag.Loc.Pos.SetRune1OffFromByte0Off(pos.Off0, src)
// 				} else if pos.Ln1 < 1 || pos.Col1 < 1 {
// 					errdiag.Loc.Pos.Off = 1 + pos.Off0
// 				}
// 				if errlen := err.Len(); errlen > 1 && len(src) > 0 {
// 					errdiag.Loc.Range = &z.SrcRange{}
// 					errdiag.Loc.Range.Start.SetRune1OffFromByte0Off(pos.Off0, src)
// 					errdiag.Loc.Range.End.SetRune1OffFromByte0Off(pos.Off0+errlen, src)
// 				}
// 			}
// 			errdiags = append(errdiags, errdiag)
// 		}
// 	}
// 	me.Lock()
// 	me.errDiags = errdiags
// 	me.Unlock()
// }

// func (*atmoDiag) KnownLinters() z.Tools { return nil }

// func (me *atmoDiag) PrepProbJobs(workspaceFiles z.WorkspaceFiles, writtenFilePaths []string) z.DiagBuildJobs {
// 	var job z.DiagJobBuild
// 	Ctx.Locked(func() { job.AffectedFilePaths = Ctx.Kits.All.SrcFilePaths() })
// 	return z.DiagBuildJobs{&job}
// }

// func (me *atmoDiag) RunProbJobs(jobs z.DiagBuildJobs, workspaceFiles z.WorkspaceFiles) (errdiags z.DiagItems) {
// 	me.Lock()
// 	errdiags = me.errDiags
// 	me.Unlock()
// 	return
// }

// func (*atmoDiag) PrepLintJobs(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) (jobs z.DiagLintJobs) {
// 	return
// }

// func (*atmoDiag) RunLintJob(job *z.DiagJobLint, workspaceFiles z.WorkspaceFiles) {}
