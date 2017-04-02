package zgo
import (
	"github.com/metaleap/go-devgo"

	"github.com/metaleap/zentient/z"
)


func (self *zgo) RefreshDiags (rebuildfilerelpath string, openfiles []string) {
	diags := map[string][]*z.RespDiag {}
	if len(rebuildfilerelpath)>0 {
		diags[rebuildfilerelpath] = append(diags[rebuildfilerelpath],
			&z.RespDiag { Cat: "devgo-mock", Msg: "rebuildfile:" + rebuildfilerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	}
	for _,filerelpath := range openfiles {
		diags[filerelpath] = append(diags[filerelpath],
			&z.RespDiag { Cat: "devgo-mock", Msg: "isopenfile:" + filerelpath, PosLn: 19, PosCol: 1, Sev: z.DIAG_WARN })
		diags[filerelpath] = append(diags[filerelpath],
			&z.RespDiag { Cat: "devgo-mock", Msg: "isfileopen:" + filerelpath, PosLn: 17, PosCol: 3, Sev: z.DIAG_INFO })
	}

	errs := devgo.RefreshPkgs()
	self.Base.DbgObjs = append(self.Base.DbgObjs, devgo.PkgsByDir)
	for _,err := range errs { self.Base.DbgMsgs = append(self.Base.DbgMsgs, err.Error()) }

	for _,pkg := range devgo.PkgsErrs {
		for _,pkgerr := range pkg.Errs {
			if len(pkgerr.RelPath)>0 {
				diags[pkgerr.RelPath] = append(diags[pkgerr.RelPath],
					&z.RespDiag { Cat: "go list", Msg: pkgerr.Msg, PosLn: pkgerr.PosLn, PosCol: pkgerr.PosCol, Sev: z.DIAG_ERR })
			}
		}
	}

	self.Base.Diags = diags
}
