package zgo

import (
	"fmt"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var (
	caddyRefreshPkgs = z.Caddy{ID: "goPkgs", Title: "GOPATH tracker", Icon: ""}
)

func caddiesInit() {
	caddyRefreshPkgs.OnReady = caddyRunRefreshPkgs
	z.Lang.Caddies = append(z.Lang.Caddies, &caddyRefreshPkgs)
}

func caddyRunRefreshPkgs() {
	caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc, caddyRefreshPkgs.Details, caddyRefreshPkgs.ClientCmdID =
		z.CADDY_BUSY, "refreshing", "", ""
	caddyRefreshPkgs.OnStatusChanged()

	if err := udevgo.RefreshPkgs(); err != nil {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_ERROR, "error: "+err.Error()
	} else {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_GOOD, fmt.Sprintf("%d Go packages (%d broken)", len(udevgo.PkgsByDir), len(udevgo.PkgsErrs))
		if len(udevgo.PkgsErrs) > 0 {
			caddyRefreshPkgs.ClientCmdID = "zen.extras.query"
			for _, pkg := range udevgo.PkgsErrs {
				caddyRefreshPkgs.Details += pkg.ImportPath + "\n"
			}
		}
	}
	caddyRefreshPkgs.OnStatusChanged()
}
