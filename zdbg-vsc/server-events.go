package main
import (
	"github.com/metaleap/zentient/zdbg-vsc/proto"
)

	var evtInitialized = zdbgvscp.NewInitializedEvent()
func onServerEvt_Initialized () {
	send(evtInitialized)
}

	var evtOutput = zdbgvscp.NewOutputEvent()
func onServerEvt_Output (cat string, msg string) {
	evtOutput.Body.Category , evtOutput.Body.Output= cat , msg
	send(evtOutput)
}

	var evtTerminated = zdbgvscp.NewTerminatedEvent()
func onServerEvt_Terminated () {
	send(evtTerminated)
}
