package zgo

import (
	"fmt"
	"time"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var (
	caddyRefreshPkgs = z.Caddy{ID: "goPkgs", Title: "Go tracker", Icon: ""}
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "go"
	l.Title = "Go"
	if l.Enabled = udevgo.HasGoDevEnv(); l.Enabled {
		caddyRefreshPkgs.OnReady = refreshPkgs
		l.Caddies = append(l.Caddies, &caddyRefreshPkgs)
		toolsInit()
		srcMod.onPreInit()
	}
}

func OnPostInit() {
	srcMod.onPostInit()
}

func refreshPkgs() {
	time.Sleep(time.Second * 3)
	caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
		z.CADDY_BUSY, "refreshing"
	caddyRefreshPkgs.OnStatusChanged()
	time.Sleep(time.Second * 3)

	if err := udevgo.RefreshPkgs(); err != nil {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_ERROR, "error: "+err.Error()
	} else {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_GOOD, fmt.Sprintf("%d packages (%d broken)", len(udevgo.PkgsByDir), len(udevgo.PkgsErrs))
	}
	caddyRefreshPkgs.OnStatusChanged()
}
