package zhs
import (
	"github.com/metaleap/zentient/z"
)


func (self *zhs) Lint (filerelpaths []string) (filediags map[string][]*z.RespDiag) {
	filediags = map[string][]*z.RespDiag {}
	for _,filerelpath := range filerelpaths {
		filediags[filerelpath] = append(filediags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "isopenfile:" + filerelpath, PosLn: 19, PosCol: 1, Sev: z.DIAG_HINT })
		filediags[filerelpath] = append(filediags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "isfileopen:" + filerelpath, PosLn: 17, PosCol: 3, Sev: z.DIAG_INFO })
	}
	return
}


func (self *zhs) BuildFrom (filerelpath string) (alldiags map[string][]*z.RespDiag) {
	alldiags = map[string][]*z.RespDiag {}
	alldiags[filerelpath] = append(alldiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	alldiags[filerelpath] = append(alldiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}
