package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return func (cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			diag := &z.RespDiag { Cat: "hlint", Sev: z.DIAG_INFO, Msg: srcref.Msg, Data: srcref.Data, PosLn: srcref.PosLn-1, PosCol: srcref.PosCol-1, Pos2Ln: srcref.Pos2Ln-1, Pos2Col: srcref.Pos2Col-1 }
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
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "rebuildfile:" + filerelpath, Sev: z.DIAG_ERR, PosLn: 9, PosCol: 2 })
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "filerebuild:" + filerelpath, Sev: z.DIAG_WARN, PosLn: 18, PosCol: 4 })
	return
}
