package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "haskell"
	l.Title = "Haskell"
	l.Enabled = udevhs.HasHsDevEnv()
}

func OnPostInit() {
}
