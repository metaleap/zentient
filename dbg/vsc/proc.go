package zdbgvsc

import (
	"github.com/metaleap/zentient/dbg"
)

func (this *Dbg) procStart() (err error) {
	stdout, stdin, stderr :=
		&zdbg.ProcInOut{Dbg: this.Impl}, &zdbg.ProcInOut{Dbg: this.Impl}, &zdbg.ProcInOut{Dbg: this.Impl, IsStdErr: true}
	if err = this.Impl.Start(stdout, stdin, stderr); err == nil {
		go this.procWait()
	}
	return
}

func (this *Dbg) procKill() error {
	return this.Impl.Kill()
}

func (this *Dbg) procWait() {
	this.waitIgnoreTermination = false
	err := this.Impl.Wait()
	if this.waitIgnoreTermination {
		this.waitIgnoreTermination = false
	} else {
		this.onServerEvt_Terminated()
	}
	if err != nil {
		this.Impl.PrintLn(true, "IDbg.Wait: "+err.Error())
	}
}
