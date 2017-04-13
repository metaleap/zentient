package zhs
import (
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-misc"

	"github.com/metaleap/zentient/z"
)

func (self *zhs) IntelDefLoc (fullsrcfilepath string, srcin string, offset string) (*udev.SrcMsg, error) {
	return nil , ugo.E("No applicable Code Intel tools yet.")
}

func (self *zhs) IntelHovs (fullsrcfilepath string, srcin string, offset string) (hovs []*z.RespHov) {
	hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools yet." })
	return
}

func (self *zhs) IntelCmpl (fullsrcfilepath string, srcin string, offset string) (cmpls []*z.RespCmpl) {
	cmpls = append(cmpls, &z.RespCmpl { Label: "BarBaz", Kind: z.CMPL_METHOD })
	return
}
