package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return func (cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			diag := &z.RespDiag { Cat: "hlint", Sev: z.DIAG_INFO, Msg: srcref.Msg, PosLn: srcref.PosLn-1, PosCol: srcref.PosCol-1 }
			if srcref.Sev=="Warning" { diag.Sev = z.DIAG_WARN }
			if srcref.Sev=="Error" { diag.Sev = z.DIAG_ERR }
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath], diag)
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


func (_ *zhs) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	freshdiags = map[string][]*z.RespDiag {}
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}
