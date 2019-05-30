package zat

import (
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

var Ctx *atmosess.Ctx

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title, l.Enabled = "atmo", "atmo", true
	var ctx atmosess.Ctx
	if err := ctx.Init(false, ""); err == nil {
		Ctx = &ctx
		workspace.onPreInit()
	}
}

func OnPostInit() {
}
