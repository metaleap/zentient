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
	livemutex	sync.Mutex

	zid			string
}



func (self *Base) Init () {
	self.resetAllDiags()
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
	self.resetAllDiags()
	if builddiags := µ.BuildFrom(filerelpath)  ;  builddiags!=nil {
		self.resetAllDiags()
		self.builddiags = builddiags
	}
}


func (self *Base) liveDiags (µ Zengine) map[string][]*RespDiag {
	self.livemutex.Lock() ; defer self.livemutex.Unlock()
	if self.livediags==nil {
		openfiles := openFiles(µ)
		self.livediags = map[string][]*RespDiag {}
		for frp,fdiags := range self.builddiags { self.livediags[frp] = fdiags }
		for frp,fdiags := range self.lintdiags {
			if uslice.StrHas(openfiles, frp) { self.livediags[frp] = append(self.livediags[frp], fdiags...) }
		}
	}
	return self.livediags
}


func (self *Base) resetAllDiags () {
	self.builddiags = map[string][]*RespDiag {}
	self.lintdiags = map[string][]*RespDiag {}
	self.resetLiveDiags()
}

func (self *Base) resetLiveDiags () {
	self.livemutex.Lock() ; defer self.livemutex.Unlock()
	self.livediags = nil
}
