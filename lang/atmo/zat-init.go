package zat

import (
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

const liveMode = true

var Ctx *atmosess.Ctx

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title, l.Enabled, l.Live = "atmo", "atmo", true, liveMode
	var ctx atmosess.Ctx
	if err := ctx.Init(false, ""); err == nil {
		Ctx = &ctx
		workspace.onPreInit()
		Ctx.Kits.OnFreshErrs = diag.updateFromErrs
		Ctx.Kits.OnSomeReprocessed = diag.updateFromErrs
		diag.updateFromErrs()
	}
}

func OnPostInit() {
}
