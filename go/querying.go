package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
)

func (self *zgo) QueryDefLoc (fullsrcfilepath string, srcin string, offset string) (*udev.SrcMsg, error) {
	return devgo.QueryDefLoc_Godef(fullsrcfilepath, srcin, offset)
}
