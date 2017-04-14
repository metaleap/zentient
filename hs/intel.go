package zhs
import (
	"github.com/metaleap/go-util-dev"

	"github.com/metaleap/zentient/z"
)

func (_ *zhs) IntelDefLoc (req *z.ReqIntel) *udev.SrcMsg {
	return nil
}

func (_ *zhs) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools yet." })
	return
}

func (_ *zhs) IntelCmpl (req *z.ReqIntel) (cmpls []*z.RespCmpl) {
	cmpls = append(cmpls, &z.RespCmpl { Label: "BarBaz", Kind: z.CMPL_METHOD })
	return
}
