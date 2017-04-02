package zhs
import (
	// "github.com/metaleap/go-devhs"

	"github.com/metaleap/zentient/z"
)


func (self *zhs) RefreshDiags (rebuildfilerelpath string, openfiles []string) {
	diags := self.Base.Diags
	isrebuild := len(rebuildfilerelpath)>0

	for relfilepath,filediags := range diags {
		filediagsnu := []*z.RespDiag {}
		if relfilepath!=rebuildfilerelpath { for _,fd := range filediags {
			if fd.Sev==z.DIAG_ERR || fd.Sev==z.DIAG_WARN { filediagsnu = append(filediagsnu, fd) } } }
		diags[relfilepath] = filediagsnu
	}

	for _,filerelpath := range openfiles {
		diags[filerelpath] = append(diags[filerelpath],
			&z.RespDiag { Cat: "devhs-mock", Msg: "isopenfile:" + filerelpath, PosLn: 19, PosCol: 1, Sev: z.DIAG_HINT })
		diags[filerelpath] = append(diags[filerelpath],
			&z.RespDiag { Cat: "devhs-mock", Msg: "isfileopen:" + filerelpath, PosLn: 17, PosCol: 3, Sev: z.DIAG_INFO })
	}

	if isrebuild {
		diags[rebuildfilerelpath] = append(diags[rebuildfilerelpath],
			&z.RespDiag { Cat: "devhs-mock", Msg: "rebuildfile:" + rebuildfilerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
		diags[rebuildfilerelpath] = append(diags[rebuildfilerelpath],
			&z.RespDiag { Cat: "devhs-mock", Msg: "filerebuild:" + rebuildfilerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })

		// if (!DiagInclGoListPkgErrs) {
		// 	//	might still like a refreshed pkg metadata list in general but not urgently needed right now
		// 	go devhs.RefreshPkgs()
		// } else {
		// 	errs := devhs.RefreshPkgs()
		// 	self.Base.DbgObjs = append(self.Base.DbgObjs, devhs.PkgsByDir)
		// 	for _,err := range errs { self.Base.DbgMsgs = append(self.Base.DbgMsgs, err.Error()) }

		// 	for _,pkg := range devhs.PkgsErrs {
		// 		for _,pkgerr := range pkg.Errs {
		// 			if len(pkgerr.RelPath)>0 && pkgerr.RelPath!=rebuildfilerelpath {
		// 				diags[pkgerr.RelPath] = append(diags[pkgerr.RelPath],
		// 					&z.RespDiag { Cat: "go list all", Msg: pkgerr.Msg, PosLn: pkgerr.PosLn, PosCol: pkgerr.PosCol, Sev: z.DIAG_ERR })
		// 			}
		// 		}
		// 	}
		// }
	}
	self.Base.Diags = diags
}
