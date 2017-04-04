package zhs
import (
	"github.com/metaleap/zentient/z"
)


func (self *zhs) Lint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) (freshdiags map[string][]*z.RespDiag) {
	freshdiags = map[string][]*z.RespDiag {}
	for _,filerelpath := range filerelpaths {
		freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "isopenfile:" + filerelpath, PosLn: 19, PosCol: 1, Sev: z.DIAG_HINT })
		freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "isfileopen:" + filerelpath, PosLn: 17, PosCol: 3, Sev: z.DIAG_INFO })
	}
	return
}


func (self *zhs) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	freshdiags = map[string][]*z.RespDiag {}
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devhs-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}
