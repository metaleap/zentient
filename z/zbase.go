package z
import (
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)

type Base struct {
	TmpStuff interface{}
}



func (self *Base) Init () {
}


func (self *Base) DoFmt (src string, custcmd string, cmds ...CmdInfo) (resp *RespFmt, err error) {
	var (	cmdoutstderr string
			c = -1
			run = false
			c2f = func(c CmdInfo) func() { return func() {
					resp.Result, cmdoutstderr, err = ugo.CmdExecStdin(src, "", c.C, c.A...)  } }
		)
	resp = &RespFmt{}
	for i, cmd := range cmds {
		if cmd.f = c2f(cmd)  ;  len(cmd.N)==0 {  cmd.N = cmd.C  }
		if cmds[i] = cmd  ;  cmd.N==custcmd && c<0 {  c = i  }
	}
	if run = (c>=0) ; run {
		cmds[c].f()
	} else if run = (len(custcmd)>0) ; run {
		c2f(CmdInfo{ N: custcmd, C: custcmd, I: true, A: []string{} })()
	} else {
		for i, cmd := range cmds {
			if run = cmd.I ; run {  cmds[i].f()  ;  break  }
		}
	}
	if (run) {
		resp.Warnings = ustr.Split(cmdoutstderr, "\n")
	} else {
		resp = nil
	}
	return
}
