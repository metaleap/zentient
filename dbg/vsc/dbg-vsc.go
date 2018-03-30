package zdbgvsc

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-leap/fs"
	"github.com/go-leap/run"
	"github.com/go-leap/str"
	"github.com/go-leap/sys"
	"github.com/metaleap/zentient/dbg"
	"github.com/metaleap/zentient/dbg/vsc/protocol"
)

type Dbg struct {
	Impl zdbg.IDbg
	sync.Mutex

	stdin                 *bufio.Scanner
	rawOut                *bufio.Writer
	vscLastInit           *zdbgvscp.InitializeRequestArguments
	sendseq               int
	logfile               *os.File
	waitIgnoreTermination bool
}

const (
	logToFile       = false
	logJsonIncoming = false
	logJsonOutgoing = false
)

var (
	bclen = []byte("Content-Length: ")
	bln   = []byte("\r\n\r\n")
)

func Main(impl zdbg.IDbg) {
	(&Dbg{Impl: impl}).main()
}

func (me *Dbg) main() {
	zdbgvscp.OnDisconnectRequest = me.onClientReq_Disconnect
	zdbgvscp.OnInitializeRequest = me.onClientReq_Initialize
	zdbgvscp.OnLaunchRequest = me.onClientReq_Launch
	zdbgvscp.OnThreadsRequest = me.onClientReq_Threads
	zdbgvscp.OnPauseRequest = me.onClientReq_Pause
	zdbgvscp.OnRestartRequest = me.onClientReq_Restart
	zdbgvscp.OnEvaluateRequest = me.onClientReq_Evaluate

	tmpdirpath := filepath.Join(usys.UserDataDirPath(true), filepath.Base(os.Args[0]))
	logdirpath := filepath.Join(tmpdirpath, "log")
	err := ufs.EnsureDir(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log"+ustr.Int64(time.Now().UnixNano())+".log.json")
	if logToFile && err == nil {
		me.logfile, err = os.Create(logfilepath)
	}
	if err != nil {
		panic(err)
	} else if me.logfile != nil {
		defer me.logfile.Close()
	}
	me.stdin, me.rawOut, _ = urun.SetupIpcPipes(1024*1024*4, zdbgvscp.IpcSplit_ContentLengthCrLfPlusJson, false)
	onerror := func(msg string) {
		me.Impl.PrintLn(true, msg)
		me.Impl.Kill()
		me.onServerEvt_Terminated()
		if me.logfile != nil {
			me.logfile.WriteString(msg)
		}
	}

	var srcfull string
	if len(os.Args) > 2 {
		srcfull = os.Args[2]
	}
	if err = me.Impl.Init(tmpdirpath, os.Args[1], srcfull, me.printLn); err != nil {
		onerror("Impl.Init:" + err.Error())
		return
	}
	defer me.Impl.Dispose()
	for me.stdin.Scan() {
		if err = me.stdin.Err(); err != nil {
			return
		}
		jsonin := me.stdin.Text()
		if logJsonIncoming && me.logfile != nil {
			_, _ = me.logfile.WriteString("\n\n\n\n\n" + jsonin)
		}
		var (
			req, resp interface{}
			respbase  *zdbgvscp.Response
			handled   bool
		)
		if req, err = zdbgvscp.TryUnmarshalRequest(jsonin); err != nil {
			onerror("TryUnmarshalRequest: " + err.Error())
			continue
		}
		if resp, respbase, handled, err = zdbgvscp.HandleRequest(req, me.initNewRespBase); resp == nil {
			onerror("BUG: resp returned was nil")
			continue
		} else if err != nil {
			respbase.Success, respbase.Message = false, err.Error()
			me.Impl.PrintLn(true, respbase.Message)
		} else if !handled {
			me.Impl.PrintLn(true, "zdbgNoHandlerYet:"+respbase.Command)
		}
		switch respbase.Command {
		case "initialize":
			me.onServerEvt_Initialized()
		case "disconnect":
			me.Impl.Kill()
		}
		me.send(resp)
		switch respbase.Command {
		case "disconnect":
			me.onServerEvt_Terminated()
			return
		}
	}
}

func (me *Dbg) printLn(isErr bool, msgLn string) {
	outcat := "stdout"
	if isErr {
		outcat = "stderr"
	}
	me.onServerEvt_Output(outcat, msgLn)
}

func (me *Dbg) send(item interface{}) {
	jsonout, err := json.Marshal(item)
	if err != nil {
		me.printLn(true, err.Error())
		return
	}
	me.Lock()
	defer me.Unlock()
	me.sendseq++
	if bresp := zdbgvscp.BaseResponse(item); bresp != nil {
		bresp.Seq = me.sendseq
	} else if bevt := zdbgvscp.BaseEvent(item); bevt != nil {
		bevt.Seq = me.sendseq
	} else if breq := zdbgvscp.BaseRequest(item); breq != nil {
		breq.Seq = me.sendseq
	}
	me.rawOut.Write(bclen)
	me.rawOut.Write([]byte(ustr.Int(len(jsonout))))
	me.rawOut.Write(bln)
	me.rawOut.Write(jsonout)
	me.rawOut.Flush()
	if logJsonOutgoing && me.logfile != nil {
		me.logfile.Write(bclen)
		me.logfile.Write([]byte(ustr.Int(len(jsonout))))
		me.logfile.Write(bln)
		me.logfile.Write(jsonout)
		me.logfile.Sync()
	}
}

func (me *Dbg) initNewRespBase(reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq
	respbase.Success = true
}
