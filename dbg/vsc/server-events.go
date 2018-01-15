package zdbgvsc

import (
	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

func (me *Dbg) onServerEvt_Initialized() {
	evtInitialized := zdbgvscp.NewInitializedEvent()
	me.send(evtInitialized)
}

func (me *Dbg) onServerEvt_Output(cat string, msg string) {
	evtOutput := zdbgvscp.NewOutputEvent()
	evtOutput.Body.Category, evtOutput.Body.Output = cat, msg
	me.send(evtOutput)
}

func (me *Dbg) onServerEvt_Stopped() {
	evtStopped := zdbgvscp.NewStoppedEvent()
	me.send(evtStopped)
}

func (me *Dbg) onServerEvt_Exited() {
	evtExited := zdbgvscp.NewExitedEvent()
	me.send(evtExited)
}

func (me *Dbg) onServerEvt_Terminated() {
	evtTerminated := zdbgvscp.NewTerminatedEvent()
	me.send(evtTerminated)
}
