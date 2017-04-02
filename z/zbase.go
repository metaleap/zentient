package z
import (
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	Diags map[string][]*RespDiag

	DbgMsgs []string
	DbgObjs []interface{}
}



func (self *Base) Init () {
	self.Diags = map[string][]*RespDiag {}

	self.DbgMsgs = []string {}
	self.DbgObjs = []interface{} {}
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

func (self *Base) refreshDiags (µ Zengine, rebuildfilerelpath string) (diags map[string][]*RespDiag) {
	openfiles := uslice.StrFilter(OpenFiles, func(relpath string) bool {
		file := AllFiles[relpath]
		return file!=nil && file.µ == µ
	})
	if !uslice.StrHas(openfiles, rebuildfilerelpath) {  rebuildfilerelpath = ""  }
	diags = µ.B().Diags
	for relfilepath,filediags := range diags {
		filediagsnu := []*RespDiag {}
		if relfilepath!=rebuildfilerelpath { for _,fd := range filediags {
			if fd.Sev==DIAG_ERR || fd.Sev==DIAG_WARN { filediagsnu = append(filediagsnu, fd) } } }
		diags[relfilepath] = filediagsnu
	}
	funcs := []func() { func() {
		for relfilepath,filediags := range µ.Lint(openfiles) {
			diags[relfilepath] = append(diags[relfilepath], filediags...)
		}
	} }
	if isrebuild := len(rebuildfilerelpath)>0 ; isrebuild {
		funcs = append(funcs, func() { diags[rebuildfilerelpath] = append(diags[rebuildfilerelpath], µ.BuildFrom(rebuildfilerelpath)...) })
	}
	ugo.WaitOn(funcs...)
	return
}
