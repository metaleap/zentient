package zhs

import (
	"errors"

	"github.com/go-leap/dev/hs"
	"github.com/metaleap/zentient"
)

func OnPreInit() (err error) {
	if z.Lang.ID, z.Lang.Title = "haskell", "Haskell"; !udevhs.HasHsDevEnv() {
		err = errors.New("Haskell `stack` does not appear to be installed.")
	} else {
		tools.onPreInit()
		diag.onPreInit()
		srcMod.onPreInit()
	}
	return
}

func OnPostInit() {
}
