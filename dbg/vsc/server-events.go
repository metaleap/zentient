package zdbgvsc

import (
	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

func (this *Dbg) onServerEvt_Initialized() {
	evtInitialized := zdbgvscp.NewInitializedEvent()
	this.send(evtInitialized)
}

func (this *Dbg) onServerEvt_Output(cat string, msg string) {
	evtOutput := zdbgvscp.NewOutputEvent()
	evtOutput.Body.Category, evtOutput.Body.Output = cat, msg
	this.send(evtOutput)
}

func (this *Dbg) onServerEvt_Stopped() {
	evtStopped := zdbgvscp.NewStoppedEvent()
	this.send(evtStopped)
}

func (this *Dbg) onServerEvt_Exited() {
	evtExited := zdbgvscp.NewExitedEvent()
	this.send(evtExited)
}

func (this *Dbg) onServerEvt_Terminated() {
	evtTerminated := zdbgvscp.NewTerminatedEvent()
	this.send(evtTerminated)
}
