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
)

func main () {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	if err := z.Init(map[string]func()z.Zengine { "go": zgo.Init, "hs": zhs.Init })  ;  err!=nil { panic(err) }

	logdirpath := filepath.Join(z.Ctx.ConfigDir, "zdbglog")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log" + ugo.SPr(time.Now().UnixNano()) + ".log")
	logfile,err := os.Create(logfilepath)  ;  if err!=nil { panic(err) } else { defer logfile.Close() }
	logln := func(msg string) { logfile.WriteString(msg+"\n")  ;  logfile.Sync() }
	logpanic := func(msg string) {  logln(msg)  ;  panic(msg)  }

	sendseq := 0  ;  bclen := []byte("Content-Length: ")  ;  bln := []byte("\r\n\r\n")
	send := func (item interface{}) {
		sendseq++
		if bresp := zdbgvscp.BaseResponse(item)  ;  bresp!=nil {
			bresp.Seq = sendseq
		} else if bevt := zdbgvscp.BaseEvent(item)  ;  bevt!=nil {
			bevt.Seq = sendseq
		} else if breq := zdbgvscp.BaseRequest(item)  ;  breq!=nil {
			breq.Seq = sendseq
		}
		jsonout,err := json.Marshal(item)  ;  if err!=nil { logpanic("json.Marshal: " + err.Error()) }
		rawOut.Write(bclen)  ;  rawOut.Write([]byte(ugo.SPr(len(jsonout))))  ;  rawOut.Write(bln)  ;  rawOut.Write(jsonout)  ;  rawOut.Flush()
		logfile.Write(bclen)  ;  logfile.Write([]byte(ugo.SPr(len(jsonout))))  ;  logfile.Write(bln)  ;  logfile.Write(jsonout)  ;  logfile.Sync()
	}

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
	}

}

func initNewRespBase (reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq  ;  respbase.Success = true
}

func init () {
	zdbgvscp.OnDisconnectRequest = onDisconnect
	zdbgvscp.OnInitializeRequest = onInitialize
	zdbgvscp.OnLaunchRequest = onLaunch
}

func onDisconnect (req *zdbgvscp.DisconnectRequest, resp *zdbgvscp.DisconnectResponse) (err error) {
	if req.Arguments.Restart {
	}
	return
}

func onInitialize (req *zdbgvscp.InitializeRequest, resp *zdbgvscp.InitializeResponse) (err error) {
	resp.Body.SupportsRestartRequest = true
	resp.Body.SupportsConfigurationDoneRequest = true
	vscLastInit = &req.Arguments
	return
}

func onLaunch (req *zdbgvscp.LaunchRequest, resp *zdbgvscp.LaunchResponse) (err error) {
	for i := 0; i <9999999;i++ {
		err = ugo.E("Nah NOT ON " + ugo.SPr(i))
	}  // :" + " C=" + r.Arguments.C + " W=" + r.Arguments.W + " F=" + r.Arguments.F
	return
}