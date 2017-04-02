package zgo
import (
	"github.com/metaleap/go-devgo"

	"github.com/metaleap/zentient/z"
)


func (self *zgo) Lint (filerelpath string) (diags []*z.RespDiag) {
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "isopenfile:" + filerelpath, PosLn: 19, PosCol: 1, Sev: z.DIAG_HINT })
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "isfileopen:" + filerelpath, PosLn: 17, PosCol: 3, Sev: z.DIAG_INFO })
	return
}

func (self *zgo) BuildFrom (filerelpath string) (diags []*z.RespDiag) {
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}


func (self *zgo) refreshPkgDiags (rebuildfilerelpath string) {
	/*errs :=*/ devgo.RefreshPkgs()
	// self.Base.DbgObjs = append(self.Base.DbgObjs, devgo.PkgsByDir)
	// for _,err := range errs { self.Base.DbgMsgs = append(self.Base.DbgMsgs, err.Error()) }
	// pd := map[string][]*z.RespDiag {}
	// for _,pkg := range devgo.PkgsErrs {
	// 	for _,pkgerr := range pkg.Errs {
	// 		if len(pkgerr.RelPath)>0 && pkgerr.RelPath!=rebuildfilerelpath {
	// 			pd[pkgerr.RelPath] = append(pd[pkgerr.RelPath],
	// 				&z.RespDiag { Cat: "go list all", Msg: pkgerr.Msg, PosLn: pkgerr.PosLn, PosCol: pkgerr.PosCol, Sev: z.DIAG_ERR })
	// 		}
	// 	}
	// }
	// pkgdiags = pd
}
