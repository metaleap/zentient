package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "haskell"
	l.Title = "Haskell"
	if l.Enabled = udevhs.HasHsDevEnv(); l.Enabled {
		toolsInit()
		srcFmt.onPreInit()
	}
}

func OnPostInit() {
}
