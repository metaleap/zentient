package z

import (
	"time"
)

var (
	caddyBuildJobs = Caddy{ID: "buildOnSave", Icon: "⚗", Title: "Build-On-Save"}
)

func (me *Caddy) onInit() {
	me.Status.Flag, me.Status.Desc = CADDY_PENDING, "pending"
}

func (me *Caddy) OnStatusChanged() {
	go send(&IpcResp{CaddyUpdate: me})
}

func (me *Caddy) IsPendingOrBusy() bool {
	return me.Status.Flag == CADDY_BUSY || me.Status.Flag == CADDY_PENDING
}
func (me *Caddy) IsReady() bool {
	return me.ready
}

func init() {
	caddyBuildJobs.OnReady = caddyBuildOnInit
	Lang.Caddies = append(Lang.Caddies, &caddyBuildJobs)
}

func caddyBuildOnInit() {
	caddyBuildJobs.Title = Lang.Title + " " + caddyBuildJobs.Title
	caddyBuildJobs.Status.Flag, caddyBuildJobs.Status.Desc, caddyBuildJobs.Details, caddyBuildJobs.UxActionID =
		CADDY_GOOD, "Nothing built-on-save yet in this session", "", "workbench.action.problems.focus"
	caddyBuildJobs.OnStatusChanged()
}

func CaddyBuildOnRunning(numJobs int, cur int, all string) {
	caddyBuildJobs.Status.Flag, caddyBuildJobs.Status.Desc, caddyBuildJobs.Details =
		CADDY_BUSY, Strf("Rebuilding: %d/%d...", cur+1, numJobs), all
	caddyBuildJobs.OnStatusChanged()
}

func CaddyBuildOnDone(failed map[string]bool, skipped map[string]bool, all []string, timeTaken time.Duration) {
	numbuilt := len(all) - (len(skipped) + len(failed))
	caddyBuildJobs.Status.Desc = Strf("out of %d packages ➜ \n\t\t%d rebuilt, %d failed, %d skipped in %s", len(all), numbuilt, len(failed), len(skipped), timeTaken)
	if len(failed) > 0 {
		caddyBuildJobs.Status.Flag = CADDY_ERROR
	} else {
		caddyBuildJobs.Status.Flag = CADDY_GOOD
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
