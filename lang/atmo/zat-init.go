package zat

import (
	"github.com/go-leap/str"
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

const liveMode = true

var Ctx *atmosess.Ctx

func OnPreInit() error {
	z.Lang.ID, z.Lang.Title = "atmo", "atmo"
	var ctx atmosess.Ctx
	kitimppath, err := ctx.Init(false, "")
	if err != nil {
		return err
	}
	Ctx, z.Lang.Live =
		&ctx, liveMode
	Ctx.KitsEnsureLoaded(true, kitimppath)
	workspace.onPreInit()
	Ctx.On.SomeKitsRefreshed = diag.updateFromErrs
	Ctx.On.NewBackgroundMessages = onNewBackgroundMessages
	diag.updateFromErrs(true)
	onNewBackgroundMessages()
	return nil
}

func OnPostInit() {
}

func onNewBackgroundMessages() {
	Ctx.Locked(func() {
		msgs := Ctx.BackgroundMessages(true)
		for i := range msgs {
			println(msgs[i].Time.Format("15:04:05") + "\t" + ustr.Join(msgs[i].Lines, "\n\t"))
		}
	})
}
