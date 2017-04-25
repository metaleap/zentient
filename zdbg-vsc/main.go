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
	respseq int
	stdin *bufio.Scanner
	rawOut *bufio.Writer

	vscLastInit *zdbgvscp.InitializeRequestArguments
)

func main () {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := z.Init(map[string]func()z.Zengine { "go": zgo.Init, "hs": zhs.Init })
	if err!=nil { panic(err) }

	logdirpath := filepath.Join(z.Ctx.ConfigDir, "zdbglog")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log" + ugo.SPr(time.Now().UnixNano()) + ".log")
	logfile,err := os.Create(logfilepath)  ;  if err!=nil { panic(err) } else { defer logfile.Close() }
	logln := func(msg string) { logfile.WriteString(msg+"\n")  ;  logfile.Sync() }
	logpanic := func(msg string) {  logln(msg)  ;  panic(msg)  }

	stdin,rawOut,_ = ugo.SetupJsonProtoPipes(1024*1024*4, true, false)
	bclen := []byte("Content-Length: ")  ;  bln := []byte("\n\n")
	for stdin.Scan() {
		injson := stdin.Text()  ;  req,err := zdbgvscp.TryUnmarshalRequest(injson)
		if err!=nil { logln(err.Error())  ;  panic(err) }
		var resp interface{}  ;  switch r := req.(type) {
			case *zdbgvscp.DisconnectRequest:
				resp = onDisconnect(r)
			case *zdbgvscp.InitializeRequest:
				resp = onInitialize(r)
			default:
				logpanic(injson)
		}
		if resp==nil { logpanic("BUG: resp returned was nil") }
		outjson,err := json.Marshal(resp)
		if err!=nil { logpanic("json.Marshal: " + err.Error()) }
		rawOut.Write(bclen)  ;  rawOut.Write([]byte(ugo.SPr(len(outjson))))  ;  rawOut.Write(bln)  ;  rawOut.Write(outjson)
	}

}

func newRespBase (reqbase *zdbgvscp.Request) (respbase zdbgvscp.Response) {
	respbase.ProtocolMessage.Seq = respseq  ;  respseq++
	respbase.ProtocolMessage.Type = "response"  ;  respbase.Type = "response"
	respbase.Request_seq = reqbase.Seq  ;  respbase.Command = reqbase.Command
	respbase.Success = true
	return
}

func onDisconnect (r *zdbgvscp.DisconnectRequest) (resp *zdbgvscp.DisconnectResponse) {
	resp = &zdbgvscp.DisconnectResponse { Response: newRespBase(&r.Request) }
	// if r.Arguments.Restart {}
	return
}

func onInitialize (r *zdbgvscp.InitializeRequest) (resp *zdbgvscp.InitializeResponse) {
	vscLastInit = &r.Arguments
	resp = &zdbgvscp.InitializeResponse { Response: newRespBase(&r.Request) }
	resp.Body.SupportsRestartRequest = true
	return
}
