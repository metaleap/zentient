package zgo

import (
	"github.com/go-leap/dev/go"
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title = "go", "Go"
	if l.Enabled = udevgo.HasGoDevEnv(); l.Enabled {
		settings.onPreInit()
		workspace.onPreInit()
		tools.onPreInit()
		diag.onPreInit()
		srcMod.onPreInit()
	}
}

func OnPostInit() {
	srcMod.onPostInit()
}
