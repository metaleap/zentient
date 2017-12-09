package z

import (
	"bufio"

	"github.com/metaleap/go-util/run"
)

func catch(err *error) {
	if except := recover(); except != nil {
		if e, ok := except.(error); ok {
			*err = e
		} else {
			*err = Errf("%v", except)
		}
	}
}

func send(resp *ipcResp) (err error) {
	if err = pipeIO.outEncoder.Encode(resp); err == nil {
		err = pipeIO.outWriter.Flush()
	}
	return
}

func Serve() (err error) {
	var stdin *bufio.Scanner
	stdin, pipeIO.outWriter, pipeIO.outEncoder = urun.SetupJsonIpcPipes(1024*1024*4, false, true)

	// we allow our sub-ordinate go-routines to panic and just before we return, we recover() the `err` to return (if any)
	defer catch(&err)

	// announce each caddy's existence
	for _, c := range Lang.Caddies {
		send(&ipcResp{CaddyUpdate: c})
	}
	// only now are the caddies notified that their status changes may now be broadcast
	for _, c := range Lang.Caddies {
		go c.OnReady()
	}

	// we don't directly wire up a json.Decoder to stdin but read individual lines in as strings first:
	// - this enforces our line-delimited (rather than 'json-delimited') protocol
	// - allows json-decoding in separate go-routine
	// - bad lines are simply reported to client without having a single 'global' decoder in confused/error state / without needing to exit
	for stdin.Scan() {
		go serveIncomingReq(stdin.Text())
	}
	err = stdin.Err()
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
