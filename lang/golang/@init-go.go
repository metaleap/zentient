package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "go"
	l.Title = "Go"
	if l.Enabled = udevgo.HasGoDevEnv(); l.Enabled {
		go udevgo.RefreshPkgs()
		srcFmt.onPreInit()
	}
}

func OnPostInit() {
}
