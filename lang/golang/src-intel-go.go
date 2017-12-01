package zgo

import (
	"github.com/metaleap/zentient"
)

type goSrcIntel struct {
	z.SrcIntelBase
}

var srcIntel goSrcIntel

func init() {
	srcIntel.Self = &srcIntel
	z.Lang.SrcIntel = &srcIntel
}
