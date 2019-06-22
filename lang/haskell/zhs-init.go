package zhs

import (
	"errors"

	"github.com/go-leap/dev/hs"
	"github.com/metaleap/zentient"
)

func OnPreInit() (err error) {
	if !udevhs.HasHsDevEnv() {
		err = errors.New("Haskell `stack` does not appear to be installed.")
	} else {
		z.Lang.Enabled, z.Lang.ID, z.Lang.Title = true, "haskell", "Haskell"
		tools.onPreInit()
		diag.onPreInit()
		srcMod.onPreInit()
	}
	return
}

func OnPostInit() {
}
