package zat

import (
	// "github.com/go-leap/str"
	// "github.com/metaleap/atmo/0ld/session"
	"github.com/metaleap/zentient"
	// "time"
)

// const liveMode = true

// var Ctx *atmosess.Ctx

func OnPreInit() error {
	z.Lang.ID, z.Lang.Title = "atmo", "atmo"
	// var ctx atmosess.Ctx
	// ctx.On.SomeKitsRefreshed = diag.updateFromErrs
	// ctx.On.NewBackgroundMessages = onNewBackgroundMessages
	// kitimppath, err := ctx.Init(false, "")
	// if err != nil {
	// 	return err
	// }

	// Ctx, z.Lang.Live =
	// 	&ctx, liveMode
	// Ctx.KitsEnsureLoaded(true, kitimppath)
	// workspace.onPreInit()
	// diag.updateFromErrs(Ctx, true)
	// onNewBackgroundMessages(Ctx)
	// ctx.Options.FileModsCatchup.BurstLimit = 987 * time.Millisecond
	return nil
}

func OnPostInit() {
}

// func onNewBackgroundMessages(ctx *atmosess.Ctx) {
// 	msgs := ctx.BackgroundMessages(true)
// 	for i := range msgs {
// 		println(msgs[i].Time.Format("15:04:05") + "\t" + ustr.Join(msgs[i].Lines, "\n\t"))
// 	}
// }
