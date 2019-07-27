package zgo

import (
	"errors"

	"github.com/go-leap/dev/go"
	"github.com/metaleap/zentient"
)

func OnPreInit() (err error) {
	z.Lang.ID, z.Lang.Title, z.Lang.Misc.BacktickStrings = "go", "Go", true
	if !udevgo.HasGoDevEnv() {
		err = errors.New("Go does not appear to be installed.")
	} else {
		settings.onPreInit()
		workspace.onPreInit()
		tools.onPreInit()
		diag.onPreInit()
		srcMod.onPreInit()
	}
	return
}

func OnPostInit() {
	srcMod.onPostInit()
}
