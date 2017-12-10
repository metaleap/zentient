package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title = "haskell", "Haskell"
	if l.Enabled = udevhs.HasHsDevEnv(); l.Enabled {
		tools.onPreInit()
		srcMod.onPreInit()
	}
}

func OnPostInit() {
}
