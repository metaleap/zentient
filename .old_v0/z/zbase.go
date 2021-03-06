package z

import (
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type Base struct {
	builddiags         map[string]udev.SrcMsgs
	lintdiags          map[string]udev.SrcMsgs
	livediags          map[string]udev.SrcMsgs
	lintmutex          sync.Mutex
	linttime           int64
	zid                string
	disabledToolsDiag  []string
	disabledToolsIntel []string
}

func capByName(caps []*RespCmd, name string) *RespCmd {
	for _, cap := range caps {
		if cap.Title == name {
			return cap
		}
	}
	return nil
}

func (me *Base) Init() {
}

func (me *Base) zId() string {
	if len(me.zid) == 0 {
		for zid, µ := range Zengines {
			if µ.B() == me {
				me.zid = zid
				break
			}
		}
	}
	return me.zid
}

func (me *Base) DoFmt(src string, custcmd string, cmds ...RespCmd) (resp *RespTxt, err error) {
	var (
		cmdoutstderr string
		c            = -1
		run          = false
		c2f          = func(c RespCmd) func() {
			return func() {
				resp.Result, cmdoutstderr, err = urun.CmdExecStdin(src, "", c.Name, c.Args...)
			}
		}
	)
	resp = &RespTxt{}
	for i, cmd := range cmds {
		if cmd.f = c2f(cmd); len(cmd.Title) == 0 {
			cmd.Title = cmd.Name
		}
		if cmds[i] = cmd; cmd.Title == custcmd && c < 0 {
			c = i
		}
	}
	if run = (c >= 0); run {
		cmds[c].f()
	} else if run = (len(custcmd) > 0); run {
		c2f(RespCmd{Title: custcmd, Name: custcmd, Exists: true, Args: []string{}})()
	} else {
		for i, cmd := range cmds {
			if run = cmd.Exists; run {
				cmds[i].f()
				break
			}
		}
	}
	if run {
		resp.Warnings = ustr.Split(cmdoutstderr, "\n")
	} else {
		resp = nil
	}
	return
}

func openFiles(µ Zengine) (openfiles []string) {
	for _, frp := range OpenFiles {
		if file := AllFiles[frp]; file != nil && file.µ == µ {
			openfiles = append(openfiles, frp)
		}
	}
	return
}
func (me *Base) OpenFiles() []string {
	return openFiles(Zengines[me.zId()])
}

func (me *Base) CfgDiagToolEnabled(forcecmdnames []string) func(string) bool {
	return func(cmdname string) bool {
		if len(forcecmdnames) > 0 {
			return uslice.StrHas(forcecmdnames, cmdname)
		}
		return !uslice.StrHas(me.disabledToolsDiag, cmdname)
	}
}

func (me *Base) CfgIntelToolEnabled(cmdname string) bool {
	return !uslice.StrHas(me.disabledToolsIntel, cmdname)
}

func (me *Base) OnCfg(cfg map[string]interface{}) {
	if cfg != nil {
		if s, ok := cfg["diag.disabled"].(string); ok {
			me.disabledToolsDiag = ustr.Split(s, ",")
		}
		if s, ok := cfg["intel.disabled"].(string); ok {
			me.disabledToolsIntel = ustr.Split(s, ",")
		}
		me.lintmutex.Lock()
		defer me.lintmutex.Unlock()
		me.lintdiags = nil
		me.livediags = nil
		newlivediags = true
	}
}

func (me *Base) buildFrom(µ Zengine, filerelpaths []string) {
	if µ.ReadyToBuildAndLint() {
		newlivediags = true
		me.livediags = nil
		me.lintdiags = nil
		me.builddiags = µ.BuildFrom(filerelpaths)
	}
}

func (me *Base) intelTools(µ Zengine) (tools []*RespPick) {
	dcaps := []*RespCmd{}
	cd := µ.Caps("diag")
	for _, dt := range me.disabledToolsDiag {
		if cap := capByName(cd, dt); cap != nil && cap.Exists {
			dcaps = append(dcaps, cap)
		}
	}
	if len(dcaps) > 0 {
		var xs, xn string
		for _, dc := range dcaps {
			xs = xs + "." + dc.Title
		}
		if len(dcaps) > 1 {
			xn = "all "
		}
		tools = append(tools, &RespPick{Label: "Additional Code Diagnostics", Desc: "__diags" + xs, Detail: "Runs: " + ustr.Join(ustr.Split(xs[1:], "."), " · ") + " (" + xn + "disabled for on-the-fly diagnostics as per `settings.json`)"})
	}
	return
}
func (me *Base) intelTool(µ Zengine, req *ReqIntel) (srcrefs udev.SrcMsgs, ran bool) {
	if ustr.Pref(req.Id, "__diags.") {
		ran = true
		tnames := ustr.Split(req.Id, ".")[1:]
		frp, err := filepath.Rel(Ctx.SrcDir, req.Ffp)
		if err == nil {
			for _, linter := range µ.Linters([]string{frp}, tnames...) {
				if fdiags := linter(); fdiags != nil {
					for frp, fd := range fdiags {
						for _, sr := range fd {
							if sr.Data != nil {
								if rf, ok := sr.Data["rf"].(string); ok {
									if rt, ok := sr.Data["rt"].(string); ok {
										sr.Msg = sr.Ref + ": " + sr.Msg
										sr.Ref = "`" + rt + "` instead of `" + rf + "`"
									}
								}
							}
							IntelToolAddResult(&srcrefs, sr, sr.Msg, sr.Ref).Ref = filepath.Join(Ctx.SrcDir, frp)
						}
					}
				}
			}
		}
	}
	return
}
func IntelToolAddResult(tosrcrefs *udev.SrcMsgs, sr *udev.SrcMsg, label string, desc string) *udev.SrcMsg {
	if sr != nil {
		sr.Msg, sr.Misc = label, desc
		*tosrcrefs = append(*tosrcrefs, sr)
	}
	return sr
}

func (me *Base) liveDiags(µ Zengine, closedfrps []string, openedfrps []string) map[string]udev.SrcMsgs {
	if len(openedfrps) > 0 || len(closedfrps) > 0 {
		newlivediags = true
		me.livediags = nil
	}
	openfiles := openFiles(µ)
	livediags := me.livediags
	if livediags == nil {
		livediags = map[string]udev.SrcMsgs{}
		if me.builddiags != nil {
			for frp, fdiags := range me.builddiags {
				livediags[frp] = fdiags
			}
		}
		if lintdiags := me.lintdiags; lintdiags != nil {
			for _, frp := range openfiles {
				livediags[frp] = append(livediags[frp], lintdiags[frp]...)
			}
		}
	}
	if len(openedfrps) > 0 || me.lintdiags == nil {
		now := time.Now().UnixNano()
		me.linttime = now
		go me.relint(µ, now)
	}
	me.livediags = livediags
	return livediags
}

func (me *Base) relint(µ Zengine, mytime int64) {
	if µ.ReadyToBuildAndLint() && mytime >= me.linttime {
		me.lintmutex.Lock()
		defer me.lintmutex.Unlock() // we won't race ourselves doing the same work n times over

		if mytime >= me.linttime {
			lintdiags := me.lintdiags
			if lintdiags == nil { // so we can check at the end whether we're already outdated
				lintdiags = map[string]udev.SrcMsgs{}
				me.lintdiags = lintdiags
			}
			freshdiags := map[string]udev.SrcMsgs{}
			lintfiles := []string{}
			for _, frp := range openFiles(µ) {
				if _, alreadylinted := lintdiags[frp]; !alreadylinted {
					lintfiles = append(lintfiles, frp)
				}
			}
			for _, frp := range lintfiles {
				freshdiags[frp] = udev.SrcMsgs{}
			} // init to non-nil so our next alreadylinted above will be correct

			if len(lintfiles) > 0 && mytime >= me.linttime {
				me.runLinters(µ.Linters(lintfiles), freshdiags)
				if lintdiags = me.lintdiags; lintdiags != nil && mytime >= me.linttime {
					prependnowtime := false
					nowstr := "" // turn on after major refactors to verify things stay cached as long as possible/permissable
					if prependnowtime {
						nowstr = ustr.After(time.Now().String(), " ", true)
						nowstr = (nowstr[:strings.LastIndex(nowstr, ":")]) + "\t"
					}
					for frp, fdiags := range freshdiags {
						if prependnowtime {
							for i, _ := range fdiags {
								fd := fdiags[i]
								fd.Ref = nowstr + fd.Ref
								fdiags[i] = fd
							}
						}
						lintdiags[frp] = fdiags
					}
					if me.lintdiags != nil && mytime >= me.linttime {
						me.lintdiags = lintdiags
						me.livediags = nil
						newlivediags = true
					}
				}
			}
		}
	}
}

func (me *Base) runLinters(linters []func() map[string]udev.SrcMsgs, freshdiags map[string]udev.SrcMsgs) {
	var mutex sync.Mutex
	lintjobs := []func(){}
	for _, linter := range linters {
		lf := linter
		lintjobs = append(lintjobs, func() {
			if linterdiags := lf(); linterdiags != nil {
				mutex.Lock()
				defer mutex.Unlock()
				for frp, fdiags := range linterdiags {
					freshdiags[frp] = append(freshdiags[frp], fdiags...)
				}
			}
		})
	}
	urun.WaitOn(lintjobs...)
}
