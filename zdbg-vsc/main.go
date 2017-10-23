package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"

	"github.com/metaleap/zentient/go"
	"github.com/metaleap/zentient/hs"
	"github.com/metaleap/zentient/z"
	"github.com/metaleap/zentient/zdbg-vsc/proto"
)

var (
	stdin  *bufio.Scanner
	rawOut *bufio.Writer

	vscLastInit *zdbgvscp.InitializeRequestArguments
	sendseq     = 0
	bclen       = []byte("Content-Length: ")
	bln         = []byte("\r\n\r\n")
	logfile     *os.File
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := z.Init(map[string]func() z.Zengine{"go": zgo.Init, "hs": zhs.Init})
	if err != nil {
		panic(err)
	}

	logdirpath := filepath.Join(z.Ctx.ConfigDir, "zdbglog")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log"+umisc.Str(time.Now().UnixNano())+".log")
	logfile, err = os.Create(logfilepath)
	if err != nil {
		panic(err)
	} else {
		defer logfile.Close()
	}
	logpanic := func(msg string) { logfile.WriteString(msg); panic(msg) }

	stdin, rawOut, _ = urun.SetupJsonProtoPipes(1024*1024*4, true, false)
	var req, resp interface{}
	var respbase *zdbgvscp.Response
	for stdin.Scan() {
		jsonin := stdin.Text()
		logfile.WriteString("\n\n\n\n\n" + jsonin)
		if req, err = zdbgvscp.TryUnmarshalRequest(jsonin); err != nil {
			logpanic("TryUnmarshalRequest: " + err.Error())
		}
		if resp, respbase, err = zdbgvscp.HandleRequest(req, initNewRespBase); resp == nil {
			logpanic("BUG: resp returned was nil")
		} else if err != nil {
			respbase.Success = false
			respbase.Message = err.Error()
		}
		switch respbase.Command {
		case "initialize":
			onServerEvt_Initialized()
		}
		send(resp)
		switch respbase.Command {
		case "disconnect":
			onServerEvt_Terminated()
			return
		}
	}

}

func send(item interface{}) {
	sendseq++
	if bresp := zdbgvscp.BaseResponse(item); bresp != nil {
		bresp.Seq = sendseq
	} else if bevt := zdbgvscp.BaseEvent(item); bevt != nil {
		bevt.Seq = sendseq
	} else if breq := zdbgvscp.BaseRequest(item); breq != nil {
		breq.Seq = sendseq
	}
	jsonout, err := json.Marshal(item)
	if err != nil {
		panic("json.Marshal: " + err.Error())
	}
	rawOut.Write(bclen)
	rawOut.Write([]byte(umisc.Str(len(jsonout))))
	rawOut.Write(bln)
	rawOut.Write(jsonout)
	rawOut.Flush()
	logfile.Write(bclen)
	logfile.Write([]byte(umisc.Str(len(jsonout))))
	logfile.Write(bln)
	logfile.Write(jsonout)
	logfile.Sync()
}

func initNewRespBase(reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq
	respbase.Success = true
}
