package z
import (
	"sync"
	"time"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	builddiags	map[string][]*RespDiag
	lintdiags	map[string][]*RespDiag
	livediags	map[string][]*RespDiag
	lintmutex	sync.Mutex
	linttime	int64
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


func (self *Base) buildFrom (µ Zengine, filerelpath string) {
	if µ.ReadyToBuildAndLint() {
		self.livediags = nil
		self.lintdiags = nil
		fromfiles := append([]string { filerelpath }, uslice.StrWithout(openFiles(µ), false, filerelpath)...)
		self.builddiags = µ.BuildFrom(fromfiles)
	}
}


func (self *Base) liveDiags (µ Zengine, closedfrp string, openedfrp string) map[string][]*RespDiag {
	if len(openedfrp)>0 || len(closedfrp)>0 { self.livediags = nil }
	openfiles := openFiles(µ)  ;  livediags := self.livediags  ;  if livediags==nil {
		livediags = map[string][]*RespDiag {}
		if self.builddiags!=nil { for frp,fdiags := range self.builddiags { livediags[frp] = fdiags } }
		if lintdiags := self.lintdiags  ;  lintdiags!=nil {
			for _,frp := range openfiles { livediags[frp] = append(livediags[frp], lintdiags[frp]...) }
		}
	}
	if len(openedfrp)>0 || self.lintdiags==nil { self.linttime = time.Now().UnixNano()  ;  go self.relint(µ, self.linttime) }
	self.livediags = livediags  ;  return livediags
}


func (self *Base) relint (µ Zengine, mytime int64) {
	if µ.ReadyToBuildAndLint() && mytime>=self.linttime {
		self.lintmutex.Lock()  ;  defer self.lintmutex.Unlock() // we won't race ourselves doing the same work n times over

		if mytime>=self.linttime {
			lintdiags := self.lintdiags  ;  if lintdiags==nil { // so we can check at the end whether we're already outdated
				lintdiags = map[string][]*RespDiag {}  ;  self.lintdiags = lintdiags
			}
			freshdiags := map[string][]*RespDiag {}  ;  lintfiles := []string {}
			for _,frp := range openFiles(µ) { if _,alreadylinted := lintdiags[frp]  ;  !alreadylinted { lintfiles = append(lintfiles, frp) } }
			for _,frp := range lintfiles { freshdiags[frp] = []*RespDiag {} } // init to non-nil so our next alreadylinted above will be correct

			if len(lintfiles)>0 && mytime>=self.linttime {
				self.runLinters(µ.Linters(lintfiles), freshdiags)
				if lintdiags = self.lintdiags  ;  lintdiags!=nil && mytime>=self.linttime {
					for frp,fdiags := range freshdiags { lintdiags[frp] = fdiags }
					if self.lintdiags!=nil && mytime>=self.linttime {
						self.lintdiags , self.livediags = lintdiags , nil
					}
				}
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
	ugo.WaitOn(lintjobs...)
}
