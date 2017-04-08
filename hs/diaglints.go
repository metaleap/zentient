package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return func (cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			d := &z.RespDiag { SrcMsg: srcref }  ;  d.Ref = "hlint" ; d.Sev = z.DIAG_SEV_INFO
			d.PosLn = srcref.PosLn-1  ;  d.PosCol = srcref.PosCol-1  ;  d.Pos2Ln = srcref.Pos2Ln-1  ;  d.Pos2Col = srcref.Pos2Col-1
			fpath := srcref.Ref  ;  filediags[fpath] = append(filediags[fpath], d)
		}
		cont(filediags)
	}
}


func (self *zhs) Lint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) (freshdiags map[string][]*z.RespDiag) {
	funcs := []func(func(map[string][]*z.RespDiag)) {}  ;  latefuncs := []func(func(map[string][]*z.RespDiag)) {}
	if devhs.Has_hlint {
		funcs = append(funcs, linterHlint(filerelpaths))
	}
	return self.Base.Lint(funcs, latefuncs, ondelayedlintersdone)
}

func (_ *zhs) LintReady () bool {
	return true
}
