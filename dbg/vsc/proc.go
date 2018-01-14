package zdbgvsc

import (
	"os/exec"
)

func (me *Dbg) procLaunch() (err error) {
	me.procKill()
	me.cmd = exec.Command("go-stdinoutdummy")
	me.cmd.Stdout = &ProcInOut{dbg: me, outcat: "stdout"}
	me.cmd.Stderr = &ProcInOut{dbg: me, outcat: "stderr"}
	me.cmd.Stdin = &ProcInOut{dbg: me}
	if err = me.cmd.Start(); err != nil {
		me.procKill()
	} else {
		go me.procWait(me.cmd)
	}
	return
}

func (me *Dbg) procWait(cmd *exec.Cmd) {
	if cmd != nil {
		if err := cmd.Wait(); err != nil {
			me.onServerEvt_Output("stderr", "ERR:"+err.Error())
		}
		me.onServerEvt_Terminated()
	}
}

func (me *Dbg) procKill() (err error) {
	if me.cmd != nil && me.cmd.Process != nil {
		err = me.cmd.Process.Kill()
	}
	me.cmd = nil
	return
}

type ProcInOut struct {
	outcat string
	dbg    *Dbg
}

func (me *ProcInOut) Write(p []byte) (n int, err error) {
	n = len(p)
	me.dbg.onServerEvt_Output(me.outcat, string(p))
	return
}

func (me *ProcInOut) Read(p []byte) (n int, err error) {
	if cmdexprs := me.dbg.cmdExprs; len(cmdexprs) > 0 {
		expr := []byte(cmdexprs[0] + "\n")
		me.dbg.cmdExprs = me.dbg.cmdExprs[1:]
		if n = len(expr); n > 1 {
			for i := 0; i < n; i++ {
				p[i] = expr[i]
			}
		}
	}
	return
}
