package zhs

import (
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient/z"
)

func linterHlint(filerelpaths []string) func() map[string]udev.SrcMsgs {
	return func() map[string]udev.SrcMsgs {
		filediags := map[string]udev.SrcMsgs{}
		for _, srcref := range udevhs.LintHlint(filerelpaths) {
			fpath := srcref.Ref
			srcref.Ref = "hlint"
			srcref.Flag = z.DIAG_SEV_INFO
			filediags[fpath] = append(filediags[fpath], srcref)
		}
		return filediags
	}
}

func (me *zhs) Linters(filerelpaths []string, forcelinters ...string) (linters []func() map[string]udev.SrcMsgs) {
	cfgok := me.Base.CfgDiagToolEnabled(forcelinters)
	if udevhs.Has_hlint && cfgok("hlint") {
		linters = append(linters, linterHlint(filerelpaths))
	}
	return
}
