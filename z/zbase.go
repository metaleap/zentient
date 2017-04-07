package z
import (
	"sync"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	alldiags	map[string][]*RespDiag
	curdiags	map[string][]*RespDiag
	latediags	map[string][]*RespDiag
	diagmutex	sync.Mutex

	DbgMsgs		[]string
	DbgObjs		[]interface{}

	zid			string
}



func (self *Base) Init () {
	self.resetAllDiags()

	self.DbgMsgs = []string {}
	self.DbgObjs = []interface{} {}
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


func (self *Base) Lint (linters []func(func(map[string][]*RespDiag)), linterslate []func(func(map[string][]*RespDiag)), ondelayedlintersdone func(map[string][]*RespDiag)) (freshdiags map[string][]*RespDiag) {
	var mutex sync.Mutex  ;  var mutexlate sync.Mutex
	freshdiags = map[string][]*RespDiag {}  ;  latediags := map[string][]*RespDiag {}
	onlinterdone := func(linterdiags map[string][]*RespDiag) {
		mutex.Lock()  ;  defer mutex.Unlock()
		for frp,filediags := range linterdiags { freshdiags[frp] = append(freshdiags[frp], filediags...) }
	}
	onlinterdonelate := func(linterdiags map[string][]*RespDiag) {
		mutexlate.Lock()  ;  defer mutexlate.Unlock()
		for frp,filediags := range linterdiags { latediags[frp] = append(latediags[frp], filediags...) }
	}
	funcs := []func() {}
	for _,linter := range linters { fn := linter  ;  funcs = append(funcs, func() { fn(onlinterdone) } ) }
	ugo.WaitOn(funcs...)

	latefuncs := []func() {}
	for _,linter := range linterslate { fn := linter  ;  latefuncs = append(latefuncs, func() { fn(onlinterdonelate) }) }
	runlatefuncs := func () { ugo.WaitOn(latefuncs...)  ;  ondelayedlintersdone(latediags) }
	go runlatefuncs() // we run this only now so that the above returns potentially a bit quicker
	return
}


func (self *Base) RefreshDiags (µ Zengine, closedfilerelpath string, openedfilerelpath string, writtenfilerelpath string) {
	if (!µ.LintReady()) { return }
	ondelayedlintersdone := func(ldiags map[string][]*RespDiag) {
		self.diagmutex.Lock() ; defer self.diagmutex.Unlock()
		for frp,fdiags := range ldiags {  self.latediags[frp] = fdiags  }
	}

	var mutex sync.Mutex
	lintfiles := []string {}
	funcs := []func() {}
	freshdiags := map[string][]*RespDiag {}

	if len(openedfilerelpath)>0 {
		if _,cached := self.alldiags[openedfilerelpath] ; !cached {
			lintfiles = append(lintfiles, openedfilerelpath)
		}
	}
	openfiles := openFiles(µ)
	if len(writtenfilerelpath)>0 {
		self.resetAllDiags()
		lintfiles = openfiles
		funcs = append(funcs, func() {
			diagsfrombuild := µ.BuildFrom(writtenfilerelpath)
			mutex.Lock()  ;  defer mutex.Unlock()
			for frp,filediags := range diagsfrombuild { freshdiags[frp] = append(freshdiags[frp], filediags...) }
		})
	} else { // edge-case: there may be openfiles without existing diags if they were opened before diagnostics were ready to run: attempt to catchup now
		for _,frp := range openfiles {
			if _,cached := self.alldiags[frp] ; (!cached) && !uslice.StrHas(lintfiles, frp) {
				lintfiles = append(lintfiles, frp)
			}
		}
	}
	if len(closedfilerelpath)>0 {
		lintfiles = uslice.StrWithout(lintfiles, false, closedfilerelpath)
	}
	if len(lintfiles)>0 {
		funcs = append(funcs, func() {
			diagsfromlint := µ.Lint(lintfiles, ondelayedlintersdone)
			mutex.Lock()  ;  defer mutex.Unlock()
			for frp,filediags := range diagsfromlint { freshdiags[frp] = append(freshdiags[frp], filediags...) }
		})
	}
	ugo.WaitOn(funcs...)
	for _,frp := range lintfiles {
		if _,hadlints := freshdiags[frp] ; !hadlints { freshdiags[frp] = []*RespDiag {} } // mustn't be nil so our catchup above works
	}

	self.diagmutex.Lock()  ;  defer self.diagmutex.Unlock()
	for frp,filediags := range freshdiags { self.alldiags[frp] = filediags }
	self.curdiags = map[string][]*RespDiag {}
	openfiles = openFiles(µ) //	a refresh is prudent here
	for _,frp := range openfiles { self.curdiags[frp] = self.alldiags[frp] }
	for frp,filediags := range self.alldiags { if !uslice.StrHas(openfiles, frp) {
		// file closed or not --- any errors that we already found earlier WILL continue to be shown:
		for _,diag := range filediags { if diag.Sev == DIAG_ERR { self.curdiags[frp] = append(self.curdiags[frp], diag) } }
	} }
	for frp,filediags := range self.latediags {
		if uslice.StrHas(openfiles, frp) {
			self.curdiags[frp] = append(self.curdiags[frp], filediags...)
		}
	}
	//	duplicates are very much possible onFileWrite as we rebuild dependant pkgs/libs/projs, so detect them
	scd := self.curdiags  ;  for frp,fds := range scd {
		mod := false  ;  for i,fd := range fds {
			for j := i+1 ;  j<len(fds)  ;  j++ {
				if fds[j].Msg==fd.Msg && fds[j].PosLn==fd.PosLn && fds[j].PosCol==fd.PosCol && fds[j].Sev==fd.Sev /*&& fds[j].Cat==fd.Cat && fds[j].Code==fd.Code*/ {
					fds[j] = fds[len(fds)-1]  ;  fds = fds[:len(fds)-1]  ;  mod = true  ;  j--
				}
			}
		}
		if mod { self.curdiags[frp] = fds }
	}
}


func (self *Base) resetAllDiags () {
	self.diagmutex.Lock() ; defer self.diagmutex.Unlock()
	self.curdiags = map[string][]*RespDiag {}
	self.alldiags = map[string][]*RespDiag {}
	self.latediags = map[string][]*RespDiag {}
}
