package zgo

import (
	"github.com/metaleap/zentient"
)

type goExtraIntel struct {
	z.ExtrasBase
}

var extraIntel goExtraIntel

func init() {
	extraIntel.Impl = &extraIntel
	z.Lang.ExtraIntel = &extraIntel
}
