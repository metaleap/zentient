package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func()map[string][]*z.RespDiag {
	return func () map[string][]*z.RespDiag {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			diag := &z.RespDiag { Sev: z.DIAG_SEV_INFO , SrcMsg: srcref }  ;  diag.Ref = "hlint"
			diag.PosLn = srcref.PosLn  ;  diag.PosCol = srcref.PosCol  ;  diag.Pos2Ln = srcref.Pos2Ln  ;  diag.Pos2Col = srcref.Pos2Col
			if srcref.Data != nil {
				if _md,ok := srcref.Data["_md"].(string)  ;  ok {
					diag.Ref = diag.Ref + " » " + _md
					delete(srcref.Data, "_md")
				}
			}
			fpath := srcref.Ref  ;  filediags[fpath] = append(filediags[fpath], diag)
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
