package zat

import (
	"time"

	"github.com/go-leap/str"
	"github.com/metaleap/atmo/session"
	"github.com/metaleap/zentient"
)

const liveMode = true

var Ctx *atmosess.Ctx

func OnPreInit() error {
	z.Lang.ID, z.Lang.Title = "atmo", "atmo"
	var ctx atmosess.Ctx
	ctx.On.SomeKitsRefreshed = diag.updateFromErrs
	ctx.On.NewBackgroundMessages = onNewBackgroundMessages
	ctx.Options.FileModsCatchup.BurstLimit = 456 * time.Millisecond
	kitimppath, err := ctx.Init(false, "")
	if err != nil {
		return err
	}
	Ctx, z.Lang.Live =
		&ctx, liveMode
	Ctx.KitsEnsureLoaded(true, kitimppath)
	workspace.onPreInit()
	diag.updateFromErrs(Ctx, true)
	onNewBackgroundMessages(Ctx)
	return nil
}

func OnPostInit() {
}

func onNewBackgroundMessages(ctx *atmosess.Ctx) {
	ctx.Locked(func() {
		msgs := ctx.BackgroundMessages(true)
		for i := range msgs {
			println(msgs[i].Time.Format("15:04:05") + "\t" + ustr.Join(msgs[i].Lines, "\n\t"))
		}
	})
}
