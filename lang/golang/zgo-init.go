package zgo

import (
	"errors"

	"github.com/go-leap/dev/go"
	"github.com/metaleap/zentient"
)

func OnPreInit() (err error) {
	if !udevgo.HasGoDevEnv() {
		err = errors.New("Go does not appear to be installed.")
	} else {
		z.Lang.ID, z.Lang.Title, z.Lang.Enabled = "go", "Go", true
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
