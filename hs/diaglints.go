package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func()map[string][]*z.RespDiag {
	return func () map[string][]*z.RespDiag {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			d := &z.RespDiag { Sev: z.DIAG_SEV_INFO , SrcMsg: srcref }  ;  d.Ref = "hlint"
			d.PosLn = srcref.PosLn  ;  d.PosCol = srcref.PosCol  ;  d.Pos2Ln = srcref.Pos2Ln  ;  d.Pos2Col = srcref.Pos2Col
			fpath := srcref.Ref  ;  filediags[fpath] = append(filediags[fpath], d)
		}
		return filediags
	}
}


func (self *zhs) Linters (filerelpaths []string) (linters []func()map[string][]*z.RespDiag) {
	if devhs.Has_hlint {
		linters = append(linters, linterHlint(filerelpaths))
	}
	return
}


// func (self *zhs) OldLint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) (freshdiags map[string][]*z.RespDiag) {
// 	latefuncs := []func(func(map[string][]*z.RespDiag)) {}
// 	if devhs.Has_hlint {
// 		latefuncs = append(latefuncs, linterHlint(filerelpaths))
// 	}
// 	return // self.Base.Lint(latefuncs, ondelayedlintersdone)
// }
