package zdbg

import (
	"time"
)

type ProcInOut struct {
	Dbg IDbg

	IsStdErr bool
}

func (this *ProcInOut) Write(p []byte) (n int, err error) {
	n = len(p)
	this.Dbg.PrintLn(this.IsStdErr, string(p))
	return
}

func (this *ProcInOut) Read(p []byte) (n int, err error) {
	if cmdevalexpr := this.Dbg.Dequeue(); cmdevalexpr == "" {
		time.Sleep(time.Millisecond * 23) // reduces CPU% from "too high" (~12+% here) to "fine" (under-1% here) --- the delay-time itself is arbitrary, lower means sooner writes to the sub-proc's stdin of course, but too-low (around 1ms/under) and no more CPU% benefit
	} else {
		n = copy(p, cmdevalexpr+"\n")
	}
	return
}
