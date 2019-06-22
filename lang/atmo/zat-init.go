package zat

import (
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

const liveMode = true

var Ctx *atmosess.Ctx

func OnPreInit() (err error) {
	var ctx atmosess.Ctx
	if err = ctx.Init(false, ""); err == nil {
		Ctx, z.Lang.Enabled, z.Lang.ID, z.Lang.Title, z.Lang.Live =
			&ctx, true, "atmo", "atmo", liveMode
		workspace.onPreInit()
		Ctx.Kits.OnFreshErrs = diag.updateFromErrs
		Ctx.Kits.OnSomeReprocessed = diag.updateFromErrs
		diag.updateFromErrs()
	}
	return
}

func OnPostInit() {
}
