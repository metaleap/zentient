package z
import (
	"sync"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	alldiags map[string][]*RespDiag
	curdiags map[string][]*RespDiag

	DbgMsgs []string
	DbgObjs []interface{}

	zid string
}



func (self *Base) Init () {
	self.curdiags = map[string][]*RespDiag {}
	self.alldiags = map[string][]*RespDiag {}

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

func (self *Base) refreshDiags (µ Zengine, closedfilerelpath string, openedfilerelpath string, writtenfilerelpath string) {
	var mutex sync.Mutex
	lintfiles := []string {}
	funcs := []func() {}
	freshdiags := map[string][]*RespDiag {}
	openfiles := uslice.StrFiltered(OpenFiles, func(relpath string) bool {
		file := AllFiles[relpath]
		return file!=nil && file.µ == µ
	})

	if len(closedfilerelpath)>0 {
		delete(self.curdiags, closedfilerelpath)
	}
	if len(openedfilerelpath)>0 {
		if _,cached := self.alldiags[openedfilerelpath] ; !cached {
			lintfiles = append(lintfiles, openedfilerelpath)
		}
	}
	if len(writtenfilerelpath)>0 {
		self.curdiags = map[string][]*RespDiag {}
		self.alldiags = map[string][]*RespDiag {}
		lintfiles = openfiles
		funcs = append(funcs, func() {
			diagsfrombuild := µ.BuildFrom(writtenfilerelpath)
			mutex.Lock()  ;  defer mutex.Unlock()
			for relfilepath,filediags := range diagsfrombuild { freshdiags[relfilepath] = append(freshdiags[relfilepath], filediags...) }
		})
	} else { // edge-case: there may be openfiles without existing diags if they were opened before diagnostics were ready to run: attempt to catch up now
		for _,relfilepath := range openfiles {
			if _,cached := self.alldiags[relfilepath] ; (!cached) && !uslice.StrHas(lintfiles, relfilepath) {
				lintfiles = append(lintfiles, relfilepath)
			}
		}
	}
	if len(lintfiles)>0 {
		funcs = append(funcs, func() {
			diagsfromlint := µ.Lint(lintfiles)
			mutex.Lock()  ;  defer mutex.Unlock()
			for relfilepath,filediags := range diagsfromlint { freshdiags[relfilepath] = append(freshdiags[relfilepath], filediags...) }
		})
	}
	ugo.WaitOn(funcs...)

	for relfilepath,filediags := range freshdiags {
		self.alldiags[relfilepath] = filediags
	}
	for _,relfilepath := range openfiles {
		self.curdiags[relfilepath] = self.alldiags[relfilepath]
	}
	allcurdiags[self.zId()] = self.curdiags
}
