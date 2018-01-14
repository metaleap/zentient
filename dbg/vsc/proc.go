package zdbgvsc

import (
	"os/exec"
	"time"
)

func (me *Dbg) procLaunch() (err error) {
	_ = me.procKill()
	me.cmd = exec.Command("go-stdinoutdummy")
	me.cmd.Stdout = &ProcInOut{dbg: me, outcat: "stdout"}
	me.cmd.Stderr = &ProcInOut{dbg: me, outcat: "stderr"}
	me.cmd.Stdin = &ProcInOut{dbg: me}
	if err = me.cmd.Start(); err != nil {
		_ = me.procKill()
	} else {
		go me.procWait()
	}
	return
}

func (me *Dbg) procWait() {
	_, err := me.cmd.Process.Wait() // cmd.Wait() hung forever in this specific scenario
	me.onServerEvt_Terminated()
	if err != nil {
		me.onServerEvt_Output("stderr", "exec.Cmd.Wait: "+err.Error())
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
	if len(me.dbg.cmdExprs) == 0 {
		time.Sleep(time.Millisecond * 23) // reduces this program's CPU% from "too high" (~12+% here) to "fine" (under-1% here) --- the delay-time itself is arbitrary, lower means sooner writes to the sub-proc's stdin of course, but too-low (around 1ms/under) negates the CPU% benefit
	} else {
		expr := me.dbg.cmdExprs[0]
		me.dbg.cmdExprs = me.dbg.cmdExprs[1:]
		if len(expr) > 0 {
			n = copy(p, expr+"\n")
		}
	}
	return
}
