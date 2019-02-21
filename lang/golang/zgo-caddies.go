package zgo

import (
	"time"

	"github.com/go-leap/dev/go"
	"github.com/metaleap/zentient"
)

var (
	caddyRefreshPkgs = z.Caddy{ID: "goPkgs", Title: "Go Package Tracker", Icon: "📦"}
)

func init() {
	caddyRefreshPkgs.OnReady = caddyRunRefreshPkgs
	z.Lang.Caddies = append(z.Lang.Caddies, &caddyRefreshPkgs)
}

func caddyRunRefreshPkgs() {
	caddyRefreshPkgs.ShouldReRunWhenNextDone = false
	caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc, caddyRefreshPkgs.Details, caddyRefreshPkgs.UxActionID =
		z.CADDY_BUSY, "refreshing", "", "zen.menus.main."+z.Lang.PkgIntel.MenuCategory()
	caddyRefreshPkgs.OnStatusChanged()
	pkgsbydir := udevgo.PkgsByDir
	firstrun := (pkgsbydir == nil)

	if err := udevgo.RefreshPkgs(); err != nil {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_ERROR, "error: "+err.Error()
	} else {
		pkgsbydir = udevgo.PkgsByDir
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_GOOD, z.Strf("%d packages (at least %d broken)", len(pkgsbydir), len(udevgo.PkgsErrs))
	}
	caddyRefreshPkgs.OnStatusChanged()
	if firstrun && (pkgsbydir != nil) && (z.Lang.Diag != nil) {
		time.Sleep(time.Millisecond * 123)
		z.Lang.Workspace.Lock()
		defer z.Lang.Workspace.Unlock()
		z.Lang.Diag.UpdateLintDiagsIfAndAsNeeded(z.Lang.Workspace.Files(), true)
	}
	if caddyRefreshPkgs.ShouldReRunWhenNextDone {
		go caddyRunRefreshPkgs()
	}
}
