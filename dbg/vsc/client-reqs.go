package zdbgvsc

import (
	"errors"
	"time"

	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

func (this *Dbg) onClientReq_Initialize(req *zdbgvscp.InitializeRequest, resp *zdbgvscp.InitializeResponse) (err error) {
	resp.Body.SupportsRestartRequest = true
	resp.Body.SupportsConfigurationDoneRequest = true
	this.vscLastInit = &req.Arguments
	return
}

func (this *Dbg) onClientReq_Threads(req *zdbgvscp.ThreadsRequest, resp *zdbgvscp.ThreadsResponse) (err error) {
	resp.Body.Threads = []zdbgvscp.Thread{{Id: 1, Name: "DummyThread"}}
	return
}

func (this *Dbg) onClientReq_Launch(req *zdbgvscp.LaunchRequest, resp *zdbgvscp.LaunchResponse) (err error) {
	err = this.procStart()
	return
}

func (this *Dbg) onClientReq_Evaluate(req *zdbgvscp.EvaluateRequest, resp *zdbgvscp.EvaluateResponse) (err error) {
	this.Impl.Enqueue(req.Arguments.Expression)
	return
}

func (this *Dbg) onClientReq_Pause(req *zdbgvscp.PauseRequest, resp *zdbgvscp.PauseResponse) (err error) {
	err = errors.New("Not currently supported: Pause")
	return
}

func (this *Dbg) onClientReq_Restart(req *zdbgvscp.RestartRequest, resp *zdbgvscp.RestartResponse) (err error) {
	this.waitIgnoreTermination = true
	_ = this.procKill()
	for this.waitIgnoreTermination {
		time.Sleep(time.Millisecond)
	}
	err = this.procStart()
	return
}

func (this *Dbg) onClientReq_Disconnect(req *zdbgvscp.DisconnectRequest, resp *zdbgvscp.DisconnectResponse) (err error) {
	_ = this.procKill()
	if req.Arguments.Restart {
		err = this.procStart()
	}
	return
}
