package zgo

import (
	"path/filepath"

	"github.com/go-leap/dev"
	"github.com/go-leap/dev/go"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
)

var (
	xIntelGuruCallers = z.ExtrasItem{ID: "guru.callers", Label: "❲guru❳ — Callers",
		Detail: "For this function / method implementation ➜ find possible callers."}
	xIntelGuruCallees = z.ExtrasItem{ID: "guru.callees", Label: "❲guru❳ — Callees",
		Detail: "For this function / method call ➜ find possible implementations to which it might dispatch."}
	xIntelGuruCallstack = z.ExtrasItem{ID: "guru.callstack", Label: "❲guru❳ — Call Stack",
		Detail: "For this function / method ➜ find an arbitrary path from a `main` call-graph root."}
	xIntelGuruFreevars = z.ExtrasItem{ID: "guru.freevars", Label: "❲guru❳ — Free Variables",
		Detail: "For this selection ➜ find variables referenced but not defined within it."}
	xIntelGuruErrtypes = z.ExtrasItem{ID: "guru.whicherrs", Label: "❲guru❳ — Error Types",
		Detail: "For this `error` value ➜ find its possible types."}
	xIntelGuruPointeeTypes = z.ExtrasItem{ID: "guru.pointsto.types", Label: "❲guru❳ — Pointees (Types)",
		Detail: "For this reference-typed expression ➜ find underlying types."}
	xIntelGuruPointeeVals = z.ExtrasItem{ID: "guru.pointsto.vals", Label: "❲guru❳ — Pointees (Allocations)",
		Detail: "For this reference-typed expression ➜ find related allocations."}
	xIntelGuruChanpeers = z.ExtrasItem{ID: "guru.peers", Label: "❲guru❳ — Channel Peers",
		Detail: "For this `<-` operator's channel ➜ find associated allocations, sends, receives and closes."}
)

func (me *goExtras) runIntel_Guru(guruCmd string, srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if !tools.guru.Installed {
		z.ToolGonePanic("guru")
	}
	bpos := srcLens.ByteOffsetForPos(srcLens.Pos)
	bp1, bp2 := ustr.Int(bpos), ""
	if srcLens.Range != nil {
		bpos1, bpos2 := srcLens.ByteOffsetForPos(&srcLens.Range.Start), srcLens.ByteOffsetForPos(&srcLens.Range.End)
		if bp1 = ustr.Int(bpos1); bpos2 != bpos1 {
			bp2 = ustr.Int(bpos2)
		}
	}
	guruscope := ""
	if settings.cfgGuruScopeMin.ValBool() {
		pkgimppath, shouldrefresh := udevgo.GuruMinimalScopeFor(srcLens.FilePath)
		if pkgimppath != "" {
			guruscope = pkgimppath + "/..."
		}
		if nope, notyet := (guruscope == ""), (udevgo.PkgsByImP == nil); nope || shouldrefresh {
			go caddyRunRefreshPkgs()
			if nope {
				if notyet {
					resp.Warns = append(resp.Warns, _PKG_NOT_READY_MSG)
				} else {
					panic("Not part of a Go package: " + filepath.Base(srcLens.FilePath) + ".")
				}
			}
			return
		}
	}
	var err error
	curpkgdir := filepath.Dir(srcLens.FilePath)
	switch guruCmd {
	case "callees":
		resp.Desc = xIntelGuruCallees.Detail
		if gcs, e := udevgo.QueryCallees_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			resp.Refs = make(z.SrcLocs, 0, len(gcs.Callees))
			for _, gc := range gcs.Callees {
				resp.Refs.AddFrom(udev.SrcMsgFromLn(gc.Pos), nil)
			}
		}
	case "callers":
		resp.Desc = xIntelGuruCallers.Detail
		if gcs, e := udevgo.QueryCallers_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			resp.Refs = make(z.SrcLocs, 0, len(gcs))
			for _, gc := range gcs {
				resp.Refs.AddFrom(udev.SrcMsgFromLn(gc.Pos), nil)
			}
		}
	case "freevars":
		resp.Desc = xIntelGuruFreevars.Label + " ➜ " + xIntelGuruFreevars.Detail
		if gfvs, e := udevgo.QueryFreevars_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2); e != nil {
			err = e
		} else {
			resp.Items = make([]*z.ExtrasItem, 0, len(gfvs))
			for _, gfv := range gfvs {
				resp.Items = append([]*z.ExtrasItem{{Desc: udevgo.PkgImpPathsToNamesInLn(gfv.Type, curpkgdir), Label: gfv.Ref,
					Detail: z.Lang.Workspace.PrettyPath(gfv.Pos), FilePos: gfv.Pos}}, resp.Items...)
			}
		}
	case "callstack":
		resp.Desc = xIntelGuruCallstack.Detail
		if gcs, e := udevgo.QueryCallstack_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			resp.Desc = udevgo.PkgImpPathsToNamesInLn(gcs.Target, curpkgdir) + " ➜ " + xIntelGuruCallstack.Label
			resp.Items = make([]*z.ExtrasItem, 0, len(gcs.Callers))
			for _, gc := range gcs.Callers {
				resp.Items = append([]*z.ExtrasItem{{Desc: gc.Desc, Label: udevgo.PkgImpPathsToNamesInLn(gc.Caller, curpkgdir),
					Detail: z.Lang.Workspace.PrettyPath(gc.Pos), FilePos: gc.Pos}}, resp.Items...)
			}
		}
	case "whicherrs":
		resp.Desc = xIntelGuruErrtypes.Detail
		if gwe, e := udevgo.QueryWhicherrs_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			for _, gwec := range gwe.Constants {
				resp.Refs.AddFrom(udev.SrcMsgFromLn(gwec), nil)
			}
			for _, gweg := range gwe.Globals {
				resp.Refs.AddFrom(udev.SrcMsgFromLn(gweg), nil)
			}
			for _, gwet := range gwe.Types {
				resp.Refs.AddFrom(udev.SrcMsgFromLn(gwet.Position), nil)
			}
		}
	case "pointsto.types", "pointsto.vals":
		ispt, ispv := guruCmd == "pointsto.types", guruCmd == "pointsto.vals"
		if ispt {
			resp.Desc = xIntelGuruPointeeTypes.Detail
		} else if ispv {
			resp.Desc = xIntelGuruPointeeVals.Detail
		}
		if gpts, e := udevgo.QueryPointsto_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			for _, gpt := range gpts {
				if ispt {
					resp.Refs.AddFrom(udev.SrcMsgFromLn(gpt.NamePos), nil)
				} else if ispv {
					for _, gptl := range gpt.Labels {
						resp.Refs.AddFrom(udev.SrcMsgFromLn(gptl.Pos), nil)
					}
				}
			}
		}
	case "peers":
		resp.Desc = xIntelGuruChanpeers.Detail
		if gp, e := udevgo.QueryPeers_Guru(srcLens.FilePath, srcLens.Txt, bp1, bp2, guruscope); e != nil {
			err = e
		} else {
			resp.Desc = udevgo.PkgImpPathsToNamesInLn(gp.Type, curpkgdir) + " ➜ " + xIntelGuruChanpeers.Label
			for _, locs := range [][]string{gp.Allocs, gp.Closes, gp.Receives, gp.Sends} {
				for _, loc := range locs {
					resp.Refs.AddFrom(udev.SrcMsgFromLn(loc), nil)
				}
			}
		}
	default:
		z.BadPanic("`guru` command", guruCmd)
	}
	if err != nil {
		errmsg, chkmsg := err.Error(), "guru: couldn't load packages due to errors: "
		if cml, i := len(chkmsg), ustr.Pos(errmsg, chkmsg); i >= 0 {
			err = nil
			oldnumscopeexcl, errpkgimppaths := len(udevgo.GuruScopeExclPkgs), ustr.Split(errmsg[i+cml:], ", ")
			if len(errpkgimppaths) > 0 {
				errpkgimppathlast := errpkgimppaths[len(errpkgimppaths)-1]
				errpkgimppaths[len(errpkgimppaths)-1] = ustr.BeforeFirst(errpkgimppathlast, " ", errpkgimppathlast)
				for _, epkg := range errpkgimppaths {
					udevgo.GuruScopeExclPkgs[epkg] = true
				}
				if len(udevgo.GuruScopeExclPkgs) > oldnumscopeexcl {
					go z.SendNotificationMessageToClient(z.DIAG_SEV_WARN, z.Strf("guru complained about %d packages, re-running with those excluded: %v", len(errpkgimppaths), errpkgimppaths))
					resp.Refs = nil
					me.runIntel_Guru(guruCmd, srcLens, arg, resp)
				}
			}
		} else {
			resp.Warns = append(resp.Warns, err.Error())
		}
	} else if len(resp.Info) == 0 && resp.Items == nil && len(resp.Warns) == 0 && len(resp.Refs) == 0 {
		resp.Items = []*z.ExtrasItem{}
	}
}
