package zhs
import (
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/zentient/z"
)


func (_ *zhs) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	freshdiags = map[string][]*z.RespDiag {}
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Sev: z.DIAG_SEV_ERR, SrcMsg: udev.SrcMsg { Ref: "devhs-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2 } })
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Sev: z.DIAG_SEV_WARN, SrcMsg: udev.SrcMsg { Ref: "devhs-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4 } })
	return
}
