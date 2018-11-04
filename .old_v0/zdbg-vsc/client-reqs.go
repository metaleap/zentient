package main
import (
	"github.com/metaleap/zentient/zdbg-vsc/proto"
)

func init () {
	zdbgvscp.OnDisconnectRequest = onClientReq_Disconnect
	zdbgvscp.OnInitializeRequest = onClientReq_Initialize
	zdbgvscp.OnLaunchRequest = onClientReq_Launch
	zdbgvscp.OnThreadsRequest = onClientReq_Threads
	zdbgvscp.OnPauseRequest = onClientReq_Pause
	zdbgvscp.OnRestartRequest = onClientReq_Restart
	zdbgvscp.OnEvaluateRequest = onClientReq_Evaluate
}

func onClientReq_Evaluate (req *zdbgvscp.EvaluateRequest, resp *zdbgvscp.EvaluateResponse) (err error) {
	cmdExprs = append(cmdExprs, req.Arguments.Expression)
	return
}

func onClientReq_Disconnect (req *zdbgvscp.DisconnectRequest, resp *zdbgvscp.DisconnectResponse) (err error) {
	if req.Arguments.Restart {}
	err = terminateProc()
	return
}

func onClientReq_Initialize (req *zdbgvscp.InitializeRequest, resp *zdbgvscp.InitializeResponse) (err error) {
	resp.Body.SupportsRestartRequest = true
	resp.Body.SupportsConfigurationDoneRequest = true
	vscLastInit = &req.Arguments
	return
}

func onClientReq_Launch (req *zdbgvscp.LaunchRequest, resp *zdbgvscp.LaunchResponse) (err error) {
	if req.Arguments.S==" " { req.Arguments.S = "" } // vsc would cancel debug session if we sent "" so we work-around on the client so we catch it on the server too, ugh
	err = launchProc(req)
	return
}

var dummyThread = []zdbgvscp.Thread { zdbgvscp.Thread { Id: 1, Name: "DummyThread" } }
func onClientReq_Threads (req *zdbgvscp.ThreadsRequest, resp *zdbgvscp.ThreadsResponse) (err error) {
	resp.Body.Threads = dummyThread
	return
}

func onClientReq_Pause (req *zdbgvscp.PauseRequest, resp *zdbgvscp.PauseResponse) (err error) {
	//	req.Arguments.ThreadId
	return
}

func onClientReq_Restart (req *zdbgvscp.RestartRequest, resp *zdbgvscp.RestartResponse) (err error) {
	return
}
