package zgo
import (
	"github.com/metaleap/go-devgo"

	"github.com/metaleap/zentient/z"
)


func (self *zgo) refreshDiags(rebuildfile *z.File, lintfile *z.File, unlintfile *z.File) {
	errs := devgo.RefreshPkgs()
	if len(errs)>0 {}
	µ.Base.DbgObjs = append(µ.Base.DbgObjs, devgo.PkgsByDir)
	for _,err:= range errs { µ.Base.DbgMsgs = append(µ.Base.DbgMsgs, err.Error()) }

	self.Base.Diags = map[string][]*z.RespDiag {}

	if rebuildfile!=nil {
		self.Base.Diags[rebuildfile.RelPath] = append(self.Base.Diags[rebuildfile.RelPath],
			&z.RespDiag { Cat: "mock", Msg: "rebuildfile:" + rebuildfile.RelPath, PosLn: 9, PosCol: 2, Sev: z.DIAG_WARN })
	}
	if lintfile!=nil {
		self.Base.Diags[lintfile.RelPath] = append(self.Base.Diags[lintfile.RelPath],
			&z.RespDiag { Cat: "mock", Msg: "lintfile:" + lintfile.RelPath, PosLn: 9, PosCol: 2, Sev: z.DIAG_WARN })
	}
	if unlintfile!=nil {
		self.Base.Diags[unlintfile.RelPath] = append(self.Base.Diags[unlintfile.RelPath],
			&z.RespDiag { Cat: "mock", Msg: "unlintfile:" + unlintfile.RelPath, PosLn: 9, PosCol: 2, Sev: z.DIAG_WARN })
	}

	for _,pkg := range devgo.PkgsErrs {
		for _,pkgerr := range pkg.Errs {
			if len(pkgerr.RelPath)>0 {
				self.Base.Diags[pkgerr.RelPath] = append(self.Base.Diags[pkgerr.RelPath],
					&z.RespDiag { Cat: "go list", Msg: pkgerr.Msg, PosLn: pkgerr.PosLn, PosCol: pkgerr.PosCol, Sev: z.DIAG_ERR })
			}
		}
	}
}
