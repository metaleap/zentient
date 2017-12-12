package z

import (
	"fmt"

	"github.com/metaleap/go-util/run"
)

func canSend() bool {
	return Prog.pipeIO.stdoutEncoder != nil && Prog.pipeIO.stdoutWriter != nil
}

func send(resp *ipcResp) (err error) {
	Prog.pipeIO.mutex.Lock()
	defer Prog.pipeIO.mutex.Unlock()
	if err = Prog.pipeIO.stdoutEncoder.Encode(resp); err == nil {
		err = Prog.pipeIO.stdoutWriter.Flush()
	}
	return
}

func catch(set *error) {
	Prog.pipeIO.stdinReadLn, Prog.pipeIO.stdoutWriter, Prog.pipeIO.stdoutEncoder = nil, nil, nil
	if except := recover(); except != nil {
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
		urun.SetupJsonIpcPipes(1024*1024*16, false, true)

	// announce each caddy's existence
	for _, c := range Lang.Caddies {
		send(&ipcResp{CaddyUpdate: c})
	}
	// only now are the caddies notified that their status changes may now be broadcast
	for _, c := range Lang.Caddies {
		go c.OnReady()
	}

	go Lang.Workspace.PollFileEventsEvery(INTERVALMS_POLL_EDITORFILEEVENTS)

	// we don't directly wire up a json.Decoder to stdin but read individual lines in as strings first:
	// - this enforces our line-delimited (rather than 'json-delimited') protocol
	// - allows json-decoding in separate go-routine
	// - bad lines are simply reported to client without having a single 'global' decoder in confused/error state / without needing to exit
	for Prog.pipeIO.stdinReadLn.Scan() {
		go serveIncomingReq(Prog.pipeIO.stdinReadLn.Text())
	}
	err = Prog.pipeIO.stdinReadLn.Err()
	return
}

func serveIncomingReq(jsonreq string) {
	resp := ipcDecodeReqAndRespond(jsonreq)

	// err only covers: either resp couldn't be json-encoded, or stdout write/flush problem:
	// both would indicate bigger problems --- still recover()ed in Serve(), but program-ending.
	// any other kind of error, above ipcDecodeReqAndRespond will record into resp.ErrMsg to report it back to the client and the program stays running
	if err := send(resp); err != nil {
		panic(err)
	}
}
