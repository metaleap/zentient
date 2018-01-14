package zdbgvsc

import (
	"time"
)

func (me *Dbg) procStart() (err error) {
	stdout, stdin, stderr := &ProcInOut{dbg: me, outcat: "stdout"}, &ProcInOut{dbg: me}, &ProcInOut{dbg: me, outcat: "stderr"}
	if err = me.Impl.Start(stdout, stdin, stderr); err == nil {
		go me.procWait()
	}
	return
}

func (me *Dbg) procKill() error {
	return me.Impl.Kill()
}

func (me *Dbg) procWait() {
	me.waitIgnoreTermination = false
	err := me.Impl.Wait()
	if me.waitIgnoreTermination {
		me.waitIgnoreTermination = false
	} else {
		me.onServerEvt_Terminated()
	}
	if err != nil {
		me.onServerEvt_Output("stderr", "IDbg.Wait: "+err.Error())
	}
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
	if cmdevalexpr := me.dbg.Impl.Dequeue(); cmdevalexpr == "" {
		time.Sleep(time.Millisecond * 23) // reduces this program's CPU% from "too high" (~12+% here) to "fine" (under-1% here) --- the delay-time itself is arbitrary, lower means sooner writes to the sub-proc's stdin of course, but too-low (around 1ms/under) negates the CPU% benefit
	} else {
		n = copy(p, cmdevalexpr+"\n")
	}
	return
}
