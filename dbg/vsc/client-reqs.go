package zdbgvsc

import (
	"time"

	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

func (me *Dbg) onClientReq_Initialize(req *zdbgvscp.InitializeRequest, resp *zdbgvscp.InitializeResponse) (err error) {
	resp.Body.SupportsRestartRequest = true
	resp.Body.SupportsConfigurationDoneRequest = true
	me.vscLastInit = &req.Arguments
	return
}

func (me *Dbg) onClientReq_Threads(req *zdbgvscp.ThreadsRequest, resp *zdbgvscp.ThreadsResponse) (err error) {
	resp.Body.Threads = []zdbgvscp.Thread{zdbgvscp.Thread{Id: 1, Name: "DummyThread"}}
	return
}

func (me *Dbg) onClientReq_Launch(req *zdbgvscp.LaunchRequest, resp *zdbgvscp.LaunchResponse) (err error) {
	err = me.procStart()
	return
}

func (me *Dbg) onClientReq_Evaluate(req *zdbgvscp.EvaluateRequest, resp *zdbgvscp.EvaluateResponse) (err error) {
	me.Impl.Enqueue(req.Arguments.Expression)
	return
}

func (me *Dbg) onClientReq_Pause(req *zdbgvscp.PauseRequest, resp *zdbgvscp.PauseResponse) (err error) {
	//	req.Arguments.ThreadId
	return
}

func (me *Dbg) onClientReq_Restart(req *zdbgvscp.RestartRequest, resp *zdbgvscp.RestartResponse) (err error) {
	me.waitIgnoreTermination = true
	_ = me.procKill()
	for me.waitIgnoreTermination {
		time.Sleep(time.Millisecond)
	}
	err = me.procStart()
	return
}

func (me *Dbg) onClientReq_Disconnect(req *zdbgvscp.DisconnectRequest, resp *zdbgvscp.DisconnectResponse) (err error) {
	_ = me.procKill()
	if req.Arguments.Restart {
		err = me.procStart()
	}
	return
}
