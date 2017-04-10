package z
import (
	"sync"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	builddiags	map[string][]*RespDiag
	lintdiags	map[string][]*RespDiag
	livediags	map[string][]*RespDiag
	lintmutex	sync.Mutex

	zid			string
}



func (self *Base) Init () {
}

func (self *Base) zId () string {
	if len(self.zid)==0 { for zid,µ := range Zengines { if µ.B()==self { self.zid = zid ; break } } }
	return self.zid
}


func (self *Base) DoFmt (src string, custcmd string, cmds ...RespCmd) (resp *RespFmt, err error) {
	var (	cmdoutstderr string
			c = -1
			run = false
			c2f = func(c RespCmd) func() { return func() {
					resp.Result, cmdoutstderr, err = ugo.CmdExecStdin(src, "", c.Name, c.Args...)  } }
		)
	resp = &RespFmt{}
	for i, cmd := range cmds {
		if cmd.f = c2f(cmd); len(cmd.Title)==0 {  cmd.Title = cmd.Name  }
		if cmds[i] = cmd; cmd.Title==custcmd && c<0 {  c = i  }
	}
	if run = (c>=0) ; run {
		cmds[c].f()
	} else if run = (len(custcmd)>0) ; run {
		c2f(RespCmd{ Title: custcmd, Name: custcmd, Exists: true, Args: []string{} })()
	} else {
		for i, cmd := range cmds {
			if run = cmd.Exists ; run {  cmds[i].f()  ;  break  }
		}
	}
	if (run) {
		resp.Warnings = ustr.Split(cmdoutstderr, "\n")
	} else {
		resp = nil
	}
	return
}


func openFiles (µ Zengine) (openfiles []string) {
	for _,frp := range OpenFiles {
		if file := AllFiles[frp] ; file!=nil && file.µ==µ { openfiles = append(openfiles, frp) }
	}
	return
}
func (self *Base) OpenFiles () []string {
	return openFiles(Zengines[self.zId()])
}


// func (self *Base) Lint (linterslate []func(func(map[string][]*RespDiag)), ondelayedlintersdone func(map[string][]*RespDiag)) {
// 	var mutex sync.Mutex  ;  var mutexlate sync.Mutex
// 	freshdiags = map[string][]*RespDiag {}  ;  var ldiags map[string][]*RespDiag

// 	// funcs := []func() {}
// 	// onsinglelinterdone := func(linterdiags map[string][]*RespDiag) {
// 	// 	mutex.Lock()  ;  defer mutex.Unlock()
// 	// 	for frp,filediags := range linterdiags { freshdiags[frp] = append(freshdiags[frp], filediags...) }
// 	// }
// 	// for _,linter := range linters { fn := linter  ;  funcs = append(funcs, func() { fn(onsinglelinterdone) } ) }
// 	// ugo.WaitOn(funcs...)

// 	latefuncs := []func() {}
// 	var runlatefuncs func()
// 	runlatefuncs = func () {
// 		latestart := time.Now().Nanosecond()  ;  ldiags = map[string][]*RespDiag {}
// 		if ugo.WaitOn(latefuncs...)  ;  latestart > self.latediagt { // don't do another run if stale, can wreck machine perf in extremely rapid progressions of repeated `onFileWrite`s
// 			ondelayedlintersdone(ldiags)
// 		}
// 	}
// 	onsinglelinterdonelate := func(linterdiags map[string][]*RespDiag) {
// 		mutexlate.Lock()  ;  defer mutexlate.Unlock()
// 		for frp,filediags := range linterdiags { ldiags[frp] = append(ldiags[frp], filediags...) }
// 	}
// 	for _,linter := range linterslate { fn := linter  ;  latefuncs = append(latefuncs, func() { fn(onsinglelinterdonelate) }) }
// 	go runlatefuncs() // we run this only now so that the above returns potentially a bit quicker
// 	return
// }


// func (self *Base) RefreshDiags (µ Zengine, closedfilerelpath string, writtenfilerelpath string) {
// 	if (!µ.ReadyToBuildOrLint()) { return }
// 	relintfiles := []string {}
// 	freshdiags := map[string][]*RespDiag {}
// 	openfiles := openFiles(µ)

// 	if len(writtenfilerelpath)>0 {
// 		relintfiles = openfiles
// 		if diagsfrombuild := µ.BuildFrom(writtenfilerelpath, self.onDelayedLintersDone)  ;  diagsfrombuild!=nil {
// 			self.resetAllDiags() // one file changed, then all cached diags across files are outdated
// 			for frp,filediags := range diagsfrombuild { freshdiags[frp] = append(freshdiags[frp], filediags...) }
// 		}
// 	} else { // covers the newly opened file if any, plus catchup on any openfiles without existing diags if they were opened before diagnostics were ready to run
// 		self.diagmutex.Lock()
// 		for _,frp := range openfiles {
// 			if _,cached := self.alldiags[frp] ; (!cached) && frp!=closedfilerelpath {
// 				relintfiles = append(relintfiles, frp)
// 				self.alldiags[frp] = []*RespDiag {}
// 			}
// 		}
// 		self.diagmutex.Unlock()
// 	}

// 	if len(relintfiles)>0 {
// 		diagsfromlint := µ.Lint(relintfiles, self.onDelayedLintersDone)
// 		for frp,filediags := range diagsfromlint { freshdiags[frp] = append(freshdiags[frp], filediags...) }
// 	}

// 	//	duplicates are very much possible onFileWrite as we might rebuild directly-dependant pkgs/libs/projs, so detect them
// 	if len(writtenfilerelpath)>0 {
// 		frds := freshdiags  ;  for frp,fds := range frds {
// 			mod := false  ;  for i,fd := range fds {  for j := i+1 ;  j<len(fds)  ;  j++ {
// 				if fds[j].Msg==fd.Msg && fds[j].Sev==fd.Sev && fds[j].Ref==fd.Ref && fds[j].PosLn==fd.PosLn && fds[j].PosCol==fd.PosCol && fds[j].Pos2Ln==fd.Pos2Ln && fds[j].Pos2Col==fd.Pos2Col {
// 					fds[j] = fds[len(fds)-1]  ;  fds = fds[:len(fds)-1]  ;  mod = true  ;  j--  } } }
// 			if mod { freshdiags[frp] = fds }
// 		}
// 	}
// 	openfiles = openFiles(µ) //	a refresh is prudent here
// 	self.diagmutex.Lock()  ;  defer self.diagmutex.Unlock()
// 	for frp,filediags := range freshdiags { self.alldiags[frp] = filediags }
// 	self.curdiags = map[string][]*RespDiag {}
// 	for frp,filediags := range self.alldiags { if !uslice.StrHas(openfiles, frp) {
// 		// file closed or not --- any "real issues" (typically, errors) that we already found earlier WILL continue to be shown:
// 		for _,diag := range filediags { if µ.DiagResident(diag.Sev) { self.curdiags[frp] = append(self.curdiags[frp], diag) } }
// 	} else { self.curdiags[frp] = self.alldiags[frp] } }
// 	for frp,filediags := range self.latediags {
// 		if uslice.StrHas(openfiles, frp) { self.curdiags[frp] = append(self.curdiags[frp], filediags...) }
// 	}
// }

// func (self *Base) onDelayedLintersDone (ldiags map[string][]*RespDiag) {
// 	self.diagmutex.Lock() ; defer self.diagmutex.Unlock()
// 	for frp,fdiags := range ldiags {  self.latediags[frp] = fdiags  }
// }


func (self *Base) buildFrom (µ Zengine, filerelpath string) {
	if µ.ReadyToBuildAndLint() {
		self.livediags = nil
		self.lintdiags = nil
		fromfiles := append([]string { filerelpath }, uslice.StrWithout(openFiles(µ), false, filerelpath)...)
		self.builddiags = µ.BuildFrom(fromfiles)
	}
}


func (self *Base) liveDiags (µ Zengine, closedfrp string, openedfrp string) map[string][]*RespDiag {
	relint := len(openedfrp)>0  ;  openfiles := openFiles(µ)
	if len(openedfrp)>0 || len(closedfrp)>0 { self.livediags = nil }
	livediags := self.livediags  ;  if livediags==nil {
		livediags = map[string][]*RespDiag {}
		if self.builddiags!=nil { for frp,fdiags := range self.builddiags { livediags[frp] = fdiags } }
		if lintdiags := self.lintdiags  ;  lintdiags==nil { relint = true } else {
			for _,frp := range openfiles { livediags[frp] = append(livediags[frp], lintdiags[frp]...) }
		}
	}
	if relint { go self.relint(µ, openfiles) }
	self.livediags = livediags  ;  return livediags
}


func (self *Base) relint (µ Zengine, openfiles []string) {
	if µ.ReadyToBuildAndLint() {
		self.lintmutex.Lock()  ;  defer self.lintmutex.Unlock() // we won't race ourselves doing the same work n times over

		lintdiags := self.lintdiags  ;  if lintdiags==nil { // so we can check at the end whether we're already outdated
			lintdiags = map[string][]*RespDiag {}  ;  self.lintdiags = lintdiags
		}
		freshdiags := map[string][]*RespDiag {}  ;  lintfiles := []string {}
		for _,frp := range openfiles { if _,alreadylinted := lintdiags[frp]  ;  !alreadylinted { lintfiles = append(lintfiles, frp) } }
		for _,frp := range lintfiles { freshdiags[frp] = []*RespDiag {} } // init to non-nil so our next alreadylinted above will be correct

		self.runLinters(µ.Linters(lintfiles), freshdiags)

		if lintdiags = self.lintdiags  ;  lintdiags!=nil { // else there was a rebuild later on and all our lints are already outdated
			for frp,fdiags := range freshdiags { lintdiags[frp] = fdiags }
			if self.lintdiags!=nil { // again with feeling, because buildFrom can strike any time!
				self.lintdiags = lintdiags  ;  self.livediags = nil
			}
		}
	}
}


func (self *Base) runLinters (linters []func()map[string][]*RespDiag, freshdiags map[string][]*RespDiag) {
	var mutex sync.Mutex
	lintjobs := []func() {}
	for _,linter := range linters {
		lf := linter  ;  lintjobs = append(lintjobs, func() {
			if linterdiags := lf()  ;  linterdiags!=nil {
				mutex.Lock()  ;  defer mutex.Unlock()
				for frp,fdiags := range linterdiags {
					freshdiags[frp] = append(freshdiags[frp], fdiags...)
				}
			}
		})
	}
	ugo.WaitOn_(lintjobs...)
}
