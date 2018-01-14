package zdbgvsc

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/sys"
	"github.com/metaleap/zentient/dbg"
	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

var (
	bclen = []byte("Content-Length: ")
	bln   = []byte("\r\n\r\n")
)

func Main(impl zdbg.IDbg) {
	(&Dbg{Impl: impl}).main()
}

type Dbg struct {
	Impl zdbg.IDbg

	stdin       *bufio.Scanner
	rawOut      *bufio.Writer
	cmd         *exec.Cmd
	cmdExprs    []string
	vscLastInit *zdbgvscp.InitializeRequestArguments
	sendseq     int
	logfile     *os.File
}

func (me *Dbg) main() {
	zdbgvscp.OnDisconnectRequest = me.onClientReq_Disconnect
	zdbgvscp.OnInitializeRequest = me.onClientReq_Initialize
	zdbgvscp.OnLaunchRequest = me.onClientReq_Launch
	zdbgvscp.OnThreadsRequest = me.onClientReq_Threads
	zdbgvscp.OnPauseRequest = me.onClientReq_Pause
	zdbgvscp.OnRestartRequest = me.onClientReq_Restart
	zdbgvscp.OnEvaluateRequest = me.onClientReq_Evaluate

	var err error
	logdirpath := filepath.Join(filepath.Join(usys.UserDataDirPath(true), filepath.Base(os.Args[0])), "log")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log"+umisc.Str(time.Now().UnixNano())+".log")
	me.logfile, err = os.Create(logfilepath)
	if err != nil {
		panic(err)
	} else {
		defer me.logfile.Close()
	}
	logpanic := func(msg string) { me.logfile.WriteString(msg); panic(msg) }

	me.stdin, me.rawOut, _ = urun.SetupJsonIpcPipes(1024*1024*4, true, false)
	var req, resp interface{}
	var respbase *zdbgvscp.Response
	for me.stdin.Scan() {
		if err = me.stdin.Err(); err != nil {
			break
		}
		jsonin := me.stdin.Text()
		me.logfile.WriteString("\n\n\n\n\n" + jsonin)
		if req, err = zdbgvscp.TryUnmarshalRequest(jsonin); err != nil {
			logpanic("TryUnmarshalRequest: " + err.Error())
		}
		if resp, respbase, err = zdbgvscp.HandleRequest(req, me.initNewRespBase); resp == nil {
			logpanic("BUG: resp returned was nil")
		} else if err != nil {
			respbase.Success = false
			respbase.Message = err.Error()
		}
		switch respbase.Command {
		case "initialize":
			me.onServerEvt_Initialized()
		}
		me.send(resp)
		switch respbase.Command {
		case "disconnect":
			me.onServerEvt_Terminated()
			return
		}
	}

}

func (me *Dbg) send(item interface{}) {
	me.sendseq++
	if bresp := zdbgvscp.BaseResponse(item); bresp != nil {
		bresp.Seq = me.sendseq
	} else if bevt := zdbgvscp.BaseEvent(item); bevt != nil {
		bevt.Seq = me.sendseq
	} else if breq := zdbgvscp.BaseRequest(item); breq != nil {
		breq.Seq = me.sendseq
	}
	jsonout, err := json.Marshal(item)
	if err != nil {
		panic("json.Marshal: " + err.Error())
	}
	me.rawOut.Write(bclen)
	me.rawOut.Write([]byte(umisc.Str(len(jsonout))))
	me.rawOut.Write(bln)
	me.rawOut.Write(jsonout)
	me.rawOut.Flush()
	me.logfile.Write(bclen)
	me.logfile.Write([]byte(umisc.Str(len(jsonout))))
	me.logfile.Write(bln)
	me.logfile.Write(jsonout)
	me.logfile.Sync()
}

func (me *Dbg) initNewRespBase(reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq
	respbase.Success = true
}
