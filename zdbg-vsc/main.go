package main
import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"

	"github.com/metaleap/zentient/z"
	"github.com/metaleap/zentient/go"
	"github.com/metaleap/zentient/hs"
	"github.com/metaleap/zentient/zdbg-vsc/proto"
)

var (
	stdin *bufio.Scanner
	rawOut *bufio.Writer

	vscLastInit *zdbgvscp.InitializeRequestArguments
	sendseq = 0
	bclen = []byte("Content-Length: ")
	bln = []byte("\r\n\r\n")
	logfile *os.File
)

func main () {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := z.Init(map[string]func()z.Zengine { "go": zgo.Init, "hs": zhs.Init })  ;  if err!=nil { panic(err) }

	logdirpath := filepath.Join(z.Ctx.ConfigDir, "zdbglog")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log" + ugo.SPr(time.Now().UnixNano()) + ".log")
	logfile,err = os.Create(logfilepath)  ;  if err!=nil { panic(err) } else { defer logfile.Close() }
	logpanic := func(msg string) {  logfile.WriteString(msg)  ;  panic(msg)  }

	stdin,rawOut,_ = ugo.SetupJsonProtoPipes(1024*1024*4, true, false)
	var req, resp interface{}  ;  var respbase *zdbgvscp.Response
	for stdin.Scan() {
		jsonin := stdin.Text()
		logfile.WriteString("\n\n\n\n\n"+jsonin)
		if req,err = zdbgvscp.TryUnmarshalRequest(jsonin)  ;  err!=nil { logpanic("TryUnmarshalRequest: " + err.Error()) }
		if resp,respbase,err = zdbgvscp.HandleRequest(req, initNewRespBase)  ;  resp==nil {
			logpanic("BUG: resp returned was nil")
		} else if err!=nil { respbase.Success = false  ;  respbase.Message = err.Error() }
		switch respbase.Command {
			case "initialize":
				send(zdbgvscp.NewInitializedEvent())
		}
		send(resp)
		switch respbase.Command{
		case "launch":
			oe := zdbgvscp.NewOutputEvent()
			oe.Body.Output = "Testing stdout"  ;  oe.Body.Category = "stdout"
			send(oe)
			oe.Body.Output = "Testing stderr"  ;  oe.Body.Category = "stderr"
			send(oe)
			oe.Body.Output = "Testing conlog"  ;  oe.Body.Category = "console"
			send(oe)
		case "disconnect":
			send(zdbgvscp.NewTerminatedEvent())
			return
		}
	}

}

func send (item interface{}) {
	sendseq++
	if bresp := zdbgvscp.BaseResponse(item)  ;  bresp!=nil {
		bresp.Seq = sendseq
	} else if bevt := zdbgvscp.BaseEvent(item)  ;  bevt!=nil {
		bevt.Seq = sendseq
	} else if breq := zdbgvscp.BaseRequest(item)  ;  breq!=nil {
		breq.Seq = sendseq
	}
	jsonout,err := json.Marshal(item)  ;  if err!=nil { panic("json.Marshal: " + err.Error()) }
	rawOut.Write(bclen)  ;  rawOut.Write([]byte(ugo.SPr(len(jsonout))))  ;  rawOut.Write(bln)  ;  rawOut.Write(jsonout)  ;  rawOut.Flush()
	logfile.Write(bclen)  ;  logfile.Write([]byte(ugo.SPr(len(jsonout))))  ;  logfile.Write(bln)  ;  logfile.Write(jsonout)  ;  logfile.Sync()
}

func initNewRespBase (reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq  ;  respbase.Success = true
}

func init () {
	zdbgvscp.OnDisconnectRequest = onClientReqDisconnect
	zdbgvscp.OnInitializeRequest = onClientReqInitialize
	zdbgvscp.OnLaunchRequest = onClientReqLaunch
	zdbgvscp.OnThreadsRequest = onClientReqThreads
	zdbgvscp.OnPauseRequest = onClientReqPause
	zdbgvscp.OnRestartRequest = onClientReqRestart
}

func onClientReqDisconnect (req *zdbgvscp.DisconnectRequest, resp *zdbgvscp.DisconnectResponse) (err error) {
	if req.Arguments.Restart {
	}
	return
}

func onClientReqInitialize (req *zdbgvscp.InitializeRequest, resp *zdbgvscp.InitializeResponse) (err error) {
	resp.Body.SupportsRestartRequest = true
	resp.Body.SupportsConfigurationDoneRequest = true
	vscLastInit = &req.Arguments
	return
}

func onClientReqLaunch (req *zdbgvscp.LaunchRequest, resp *zdbgvscp.LaunchResponse) (err error) {
	return
}

var dummyThread = []zdbgvscp.Thread { zdbgvscp.Thread { Id: 1, Name: "DummyThread" } }
func onClientReqThreads (req *zdbgvscp.ThreadsRequest, resp *zdbgvscp.ThreadsResponse) (err error) {
	resp.Body.Threads = dummyThread
	return
}

func onClientReqPause (req *zdbgvscp.PauseRequest, resp *zdbgvscp.PauseResponse) (err error) {
	//	req.Arguments.ThreadId
	return
}

func onClientReqRestart (req *zdbgvscp.RestartRequest, resp *zdbgvscp.RestartResponse) (err error) {
	return
}
