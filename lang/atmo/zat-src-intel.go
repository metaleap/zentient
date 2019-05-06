package zat

import (
	"github.com/metaleap/zentient"
)

var srcIntel atmoSrcIntel

type atmoSrcIntel struct {
	z.SrcIntelBase
}

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}
