package zhs

import (
	"github.com/metaleap/zentient"
)

var (
	srcIntel hsSrcIntel
)

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

type hsSrcIntel struct {
	z.SrcIntelBase
}
