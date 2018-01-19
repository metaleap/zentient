package zdbgvsc

import (
	"github.com/metaleap/zentient/dbg"
)

func (me *Dbg) procStart() (err error) {
	stdout, stdin, stderr :=
		&zdbg.ProcInOut{Dbg: me.Impl}, &zdbg.ProcInOut{Dbg: me.Impl}, &zdbg.ProcInOut{Dbg: me.Impl, IsStdErr: true}
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
		me.Impl.PrintLn(true, "IDbg.Wait: "+err.Error())
	}
}
