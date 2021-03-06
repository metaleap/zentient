package z

import (
	"fmt"
	"runtime/debug"

	"github.com/go-leap/run"
	"github.com/go-leap/str"
)

func canSend() bool {
	return Prog.pipeIO.stdoutEncoder != nil && Prog.pipeIO.stdoutWriter != nil
}

func send(resp *IpcResp) (err error) {
	Prog.pipeIO.mutex.Lock()
	defer Prog.pipeIO.mutex.Unlock()

	if err = Prog.pipeIO.stdoutEncoder.Encode(resp); err == nil {
		err = Prog.pipeIO.stdoutWriter.Flush()
	}
	return
}

func sendRaw(jsonResp []byte) (err error) {
	Prog.pipeIO.mutex.Lock()
	defer Prog.pipeIO.mutex.Unlock()

	if _, err = Prog.pipeIO.stdoutWriter.Write(jsonResp); err == nil {
		err = Prog.pipeIO.stdoutWriter.Flush()
	}
	return
}

func catch(set *error) {
	Prog.pipeIO.stdinReadLn, Prog.pipeIO.stdoutWriter, Prog.pipeIO.stdoutEncoder = nil, nil, nil
	if except := recover(); except != nil {
		debug.PrintStack()
		if err, _ := except.(error); err != nil {
			*set = err
		} else {
			*set = fmt.Errorf("%v", except)
		}
	}
}

func Serve() (err error) {
	// ensure that the returned `err` will capture a sub-ordinate go-routine's panic, if any:
	defer catch(&err)

	Prog.pipeIO.stdinReadLn, Prog.pipeIO.stdoutWriter, Prog.pipeIO.stdoutEncoder =
		urun.SetupIpcPipes(1024*1024*1, nil, true)

	// announce each caddy's existence
	for _, c := range Lang.Caddies {
		send(&IpcResp{CaddyUpdate: c})
	}
	// only now are the caddies notified that their status changes may now be broadcast
	for _, c := range Lang.Caddies {
		c.ready = true
		go c.OnReady()
	}

	// we don't directly wire up a json.Decoder to stdin but read individual lines in as strings first:
	// - this enforces our line-delimited (rather than 'json-delimited') protocol
	// - edit: dropped for now /* was: "allows json-decoding in separate go-routine" */
	// - bad lines are simply reported to client without having a single 'global' decoder in confused/error state / without needing to exit
	for Prog.pipeIO.stdinReadLn.Scan() {
		go serveIncomingReq(append(make([]byte, 0, 64), Prog.pipeIO.stdinReadLn.Bytes()...))
	}
	err = Prog.pipeIO.stdinReadLn.Err()
	return
}

func serveIncomingReq(jsonreq []byte) {
	resp := ipcDecodeReqAndRespond(jsonreq)

	// err only covers: either resp couldn't be json-encoded, or stdout write/flush problem:
	// both would indicate bigger problems -- still recover()ed in Serve(), but program-ending.
	// any other kind of error, above ipcDecodeReqAndRespond call will record into resp.ErrMsg to report it back to the client and the program stays running
	if err := send(resp); err != nil {
		if errmsg := err.Error(); ustr.Has(errmsg, "in string escape code") {
			SendNotificationMessageToClient(DIAG_SEV_ERR, ustr.Int(int(resp.ReqID)))
		} else {
			panic(err)
		}
	}
}
