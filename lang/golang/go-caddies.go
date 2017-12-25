package zgo

import (
	"time"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var (
	caddyRefreshPkgs = z.Caddy{ID: "goPkgs", Title: "Go Package Tracker", Icon: ""}
	caddyBuildJobs   = z.Caddy{ID: "goInsts", Title: "Go Build-On-Save", Icon: ""}
)

func init() {
	caddyBuildJobs.OnReady = caddyBuildOnInit
	caddyRefreshPkgs.OnReady = caddyRunRefreshPkgs
	z.Lang.Caddies = append(z.Lang.Caddies, &caddyRefreshPkgs, &caddyBuildJobs)
}

func caddyBuildOnInit() {
	caddyBuildJobs.Status.Flag, caddyBuildJobs.Status.Desc, caddyBuildJobs.Details, caddyBuildJobs.UxActionID =
		z.CADDY_GOOD, "Nothing rebuilt yet in this session", "", "workbench.action.problems.focus"
	caddyBuildJobs.OnStatusChanged()
}

func caddyBuildOnRunning(numJobs int, cur int, all string) {
	caddyBuildJobs.Status.Flag, caddyBuildJobs.Status.Desc, caddyBuildJobs.Details =
		z.CADDY_BUSY, z.Strf("Rebuilding Go packages: %d/%d…", cur+1, numJobs), all
	caddyBuildJobs.OnStatusChanged()
}

func caddyBuildOnDone(failed map[string]bool, skipped map[string]bool, all []string, timeTaken time.Duration) {
	numbuilt := len(all) - (len(skipped) + len(failed))
	caddyBuildJobs.Status.Desc = z.Strf("out of %d packages ➜ \n\t\t%d rebuilt, %d failed, %d skipped in %s", len(all), numbuilt, len(failed), len(skipped), timeTaken)
	if len(failed) > 0 {
		caddyBuildJobs.Status.Flag = z.CADDY_ERROR
	} else {
		caddyBuildJobs.Status.Flag = z.CADDY_GOOD
	}

	caddyBuildJobs.Details = ""
	for _, pkgimppath := range all {
		if failed[pkgimppath] {
			caddyBuildJobs.Details += "FAILED:\t\t\t\t"
		} else if skipped[pkgimppath] {
			caddyBuildJobs.Details += "Skipped:\t\t"
		} else {
			caddyBuildJobs.Details += "Rebuilt:\t\t\t"
		}
		caddyBuildJobs.Details += pkgimppath + "\n"
	}
	caddyBuildJobs.OnStatusChanged()
}

func caddyRunRefreshPkgs() {
	caddyRefreshPkgs.ShouldReRunWhenNextDone = false
	caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc, caddyRefreshPkgs.Details, caddyRefreshPkgs.UxActionID =
		z.CADDY_BUSY, "refreshing", "", "zen.menus.main."+z.Lang.PkgIntel.MenuCategory()
	caddyRefreshPkgs.OnStatusChanged()
	firstrun := (udevgo.PkgsByDir == nil)

	if err := udevgo.RefreshPkgs(); err != nil {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_ERROR, "error: "+err.Error()
	} else {
		caddyRefreshPkgs.Status.Flag, caddyRefreshPkgs.Status.Desc =
			z.CADDY_GOOD, z.Strf("%d packages (at least %d broken)", len(udevgo.PkgsByDir), len(udevgo.PkgsErrs))
	}
	caddyRefreshPkgs.OnStatusChanged()
	if firstrun && (udevgo.PkgsByDir != nil) && (z.Lang.Diag != nil) {
		time.Sleep(time.Millisecond * 123)
		z.Lang.Workspace.Lock()
		defer z.Lang.Workspace.Unlock()
		z.Lang.Diag.UpdateLintDiagsIfAndAsNeeded(z.Lang.Workspace.Files(), true)
	}
	if caddyRefreshPkgs.ShouldReRunWhenNextDone {
		go caddyRunRefreshPkgs()
	}
}
