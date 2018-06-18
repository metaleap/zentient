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

func (this *Dbg) main() {
	zdbgvscp.OnDisconnectRequest = this.onClientReq_Disconnect
	zdbgvscp.OnInitializeRequest = this.onClientReq_Initialize
	zdbgvscp.OnLaunchRequest = this.onClientReq_Launch
	zdbgvscp.OnThreadsRequest = this.onClientReq_Threads
	zdbgvscp.OnPauseRequest = this.onClientReq_Pause
	zdbgvscp.OnRestartRequest = this.onClientReq_Restart
	zdbgvscp.OnEvaluateRequest = this.onClientReq_Evaluate

	tmpdirpath := filepath.Join(usys.UserDataDirPath(true), filepath.Base(os.Args[0]))
	logdirpath := filepath.Join(tmpdirpath, "log")
	err := ufs.EnsureDir(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log"+ustr.Int64(time.Now().UnixNano())+".log.json")
	if logToFile && err == nil {
		this.logfile, err = os.Create(logfilepath)
	}
	if err != nil {
		panic(err)
	} else if this.logfile != nil {
		defer this.logfile.Close()
	}
	this.stdin, this.rawOut, _ = urun.SetupIpcPipes(1024*1024*4, zdbgvscp.IpcSplit_ContentLengthCrLfPlusJson, false)
	onerror := func(msg string) {
		this.Impl.PrintLn(true, msg)
		this.Impl.Kill()
		this.onServerEvt_Terminated()
		if this.logfile != nil {
			this.logfile.WriteString(msg)
		}
	}

	var srcfull string
	if len(os.Args) > 2 {
		srcfull = os.Args[2]
	}
	if err = this.Impl.Init(tmpdirpath, os.Args[1], srcfull, this.printLn); err != nil {
		onerror("Impl.Init:" + err.Error())
		return
	}
	defer this.Impl.Dispose()
	for this.stdin.Scan() {
		if err = this.stdin.Err(); err != nil {
			return
		}
		jsonin := this.stdin.Text()
		if logJsonIncoming && this.logfile != nil {
			_, _ = this.logfile.WriteString("\n\n\n\n\n" + jsonin)
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
		if resp, respbase, handled, err = zdbgvscp.HandleRequest(req, this.initNewRespBase); resp == nil {
			onerror("BUG: resp returned was nil")
			continue
		} else if err != nil {
			respbase.Success, respbase.Message = false, err.Error()
			this.Impl.PrintLn(true, respbase.Message)
		} else if !handled {
			this.Impl.PrintLn(true, "zdbgNoHandlerYet:"+respbase.Command)
		}
		switch respbase.Command {
		case "initialize":
			this.onServerEvt_Initialized()
		case "disconnect":
			this.Impl.Kill()
		}
		this.send(resp)
		switch respbase.Command {
		case "disconnect":
			this.onServerEvt_Terminated()
			return
		}
	}
}

func (this *Dbg) printLn(isErr bool, msgLn string) {
	outcat := "stdout"
	if isErr {
		outcat = "stderr"
	}
	this.onServerEvt_Output(outcat, msgLn)
}

func (this *Dbg) send(item interface{}) {
	jsonout, err := json.Marshal(item)
	if err != nil {
		this.printLn(true, err.Error())
		return
	}
	this.Lock()
	defer this.Unlock()
	this.sendseq++
	if bresp := zdbgvscp.BaseResponse(item); bresp != nil {
		bresp.Seq = this.sendseq
	} else if bevt := zdbgvscp.BaseEvent(item); bevt != nil {
		bevt.Seq = this.sendseq
	} else if breq := zdbgvscp.BaseRequest(item); breq != nil {
		breq.Seq = this.sendseq
	}
	this.rawOut.Write(bclen)
	this.rawOut.Write([]byte(ustr.Int(len(jsonout))))
	this.rawOut.Write(bln)
	this.rawOut.Write(jsonout)
	this.rawOut.Flush()
	if logJsonOutgoing && this.logfile != nil {
		this.logfile.Write(bclen)
		this.logfile.Write([]byte(ustr.Int(len(jsonout))))
		this.logfile.Write(bln)
		this.logfile.Write(jsonout)
		this.logfile.Sync()
	}
}

func (this *Dbg) initNewRespBase(reqbase *zdbgvscp.Request, respbase *zdbgvscp.Response) {
	respbase.Request_seq = reqbase.Seq
	respbase.Success = true
}
