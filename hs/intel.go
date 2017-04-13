package zhs
import (
	"github.com/metaleap/go-util-dev"

	"github.com/metaleap/zentient/z"
)

func (self *zhs) IntelDefLoc (req *z.ReqIntel) *udev.SrcMsg {
	return nil
}

func (self *zhs) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools yet." })
	return
}

func (self *zhs) IntelCmpl (req *z.ReqIntel) (cmpls []*z.RespCmpl) {
	cmpls = append(cmpls, &z.RespCmpl { Label: "BarBaz", Kind: z.CMPL_METHOD })
	return
}
