package zat

import (
	"github.com/go-leap/str"
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

const liveMode = true

var Ctx *atmosess.Ctx

func OnPreInit() (err error) {
	z.Lang.ID, z.Lang.Title = "atmo", "atmo"
	var ctx atmosess.Ctx
	if err = ctx.Init(false, ""); err == nil {
		Ctx, z.Lang.Live =
			&ctx, liveMode
		workspace.onPreInit()
		Ctx.On.SomeKitsRefreshed = diag.updateFromErrs
		Ctx.On.NewBackgroundMessages = onNewBackgroundMessages
		diag.updateFromErrs(true)
		onNewBackgroundMessages()
	}
	return
}

func OnPostInit() {
}

func onNewBackgroundMessages() {
	msgs := Ctx.BackgroundMessages(true)
	for i := range msgs {
		println(msgs[i].Time.Format("15:04:05") + "\t" + ustr.Join(msgs[i].Lines, "\n\t"))
	}
}
