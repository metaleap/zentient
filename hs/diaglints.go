package zhs
import (
	"github.com/metaleap/go-devhs"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/zentient/z"
)


func linterHlint (filerelpaths []string) func()map[string][]*udev.SrcMsg {
	return func () map[string][]*udev.SrcMsg {
		filediags := map[string][]*udev.SrcMsg {}
		for _,srcref := range devhs.LintHlint(filerelpaths) {
			fpath := srcref.Ref  ;  srcref.Ref = "hlint"  ;  srcref.Flag = z.DIAG_SEV_INFO
			if srcref.Data != nil { if _md,_ := srcref.Data["_md"].(string)  ;  len(_md)>0 {
				srcref.Ref = srcref.Ref + " » " + _md  ;  delete(srcref.Data, "_md")
			} }
			filediags[fpath] = append(filediags[fpath], srcref)
		}
		return filediags
	}
}


func (me *zhs) Linters (filerelpaths []string) (linters []func()map[string][]*udev.SrcMsg) {
	cfgok := me.Base.CfgDiagToolEnabled
	if devhs.Has_hlint && cfgok("hlint") { linters = append(linters, linterHlint(filerelpaths)) }
	return
}
