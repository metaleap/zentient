package zgo

import (
	"path/filepath"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var (
	xIntelGuruCallees = z.ExtrasItem{ID: "guru.callees", Label: "Callees",
		Detail: "For this function / method call, lists possible implementations to which it might dispatch."}
	xIntelGuruCallers = z.ExtrasItem{ID: "guru.callers", Label: "Callers",
		Detail: "For this function / method implementation, lists possible callers."}
	xIntelGuruCallstack = z.ExtrasItem{ID: "guru.callstack", Label: "Call Stack",
		Detail: "For this function / method, shows an arbitrary path to the root of the call graph."}
	xIntelGuruFreevars = z.ExtrasItem{ID: "guru.freevars", Label: "Free Variables",
		Detail: "For this selection, lists variables referenced but not defined within it."}
	xIntelGuruErrtypes = z.ExtrasItem{ID: "guru.whicherrs", Label: "Error Types",
		Detail: "For this `error` value, lists its possible types."}
	xIntelGuruPointees = z.ExtrasItem{ID: "guru.pointsto", Label: "Pointees",
		Detail: "For this pointer-typed / container-typed expression, lists possible associated types / symbols."}
	xIntelGuruChanpeers = z.ExtrasItem{ID: "guru.peers", Label: "Channel Peers",
		Detail: "For this `<-` operation's channel, lists associated allocations, sends, receives and closes."}
)

func (me *goExtras) runIntel_Guru(guruCmd string, srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if !tools.guru.Installed {
		z.ToolGonePanic("guru")
	}
	bpos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	bp1, bp2 := ustr.FromInt(bpos), ""
	if srcLens.Range != nil {
		bpos1, bpos2 := srcLens.ByteOffsetForPosWithRuneOffset(&srcLens.Range.Start), srcLens.ByteOffsetForPosWithRuneOffset(&srcLens.Range.End)
		bp1, bp2 = ustr.FromInt(bpos1), ustr.FromInt(bpos2)
	}
	guruscope := ""
	if settings.cfgGuruScopeMin.ValBool() {
		pkgs, shouldrefresh := udevgo.PkgsForFiles(srcLens.FilePath)
		for _, pkg := range pkgs {
			guruscope = pkg.ImportPath + "/..."
			break
		}
		if guruscope == "" {
			go caddyRunRefreshPkgs()
			panic("Not part of a Go package: " + filepath.Base(srcLens.FilePath))
		} else if shouldrefresh {
			go caddyRunRefreshPkgs()
		}
	}
	var err error
	switch guruCmd {
	case "callees":
		if gcs, e := udevgo.QueryCallees_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			resp.Refs = make(z.SrcLenses, 0, len(gcs.Callees))
			for _, gc := range gcs.Callees {
				if srcref := udev.SrcMsgFromLn(gc.Pos); srcref != nil {
					resp.Refs.AddFrom(srcref, nil)
				}
			}
		}
	default:
		z.BadPanic("`guru` command", guruCmd)
	}
	if err != nil {
		errmsg, chkmsg := err.Error(), "guru: couldn't load packages due to errors: "
		if cml, i := len(chkmsg), ustr.Idx(errmsg, chkmsg); i >= 0 {
			oldnumscopeexcl, errpkgimppaths := len(udevgo.GuruScopeExclPkgs), ustr.Split(errmsg[i+cml:], ", ")
			if len(errpkgimppaths) > 0 {
				errpkgimppaths[len(errpkgimppaths)-1] = ustr.Before(errpkgimppaths[len(errpkgimppaths)-1], " ", false)
				for _, epkg := range errpkgimppaths {
					udevgo.GuruScopeExclPkgs[epkg] = true
				}
				if len(udevgo.GuruScopeExclPkgs) > oldnumscopeexcl {
					go z.SendNotificationMessageToClient(z.DIAG_SEV_WARN, z.Strf("guru complained about %d packages, re-running with those excluded: %v", len(errpkgimppaths), errpkgimppaths))
					resp.Refs = nil
					me.runIntel_Guru(guruCmd, srcLens, arg, resp)
				}
			}
		}
	}
	if err != nil {
		panic(err)
	}
}
