package z

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/sys"
)

var (
	strf = fmt.Sprintf

	cmdProviders []iCoreCmds
	handlers     = []iHandler{
		&coreCmdsHandler{},
	}

	Lang struct {
		Enabled bool
		ID      string
		Title   string
		SrcFmt  iSrcFormatting
	}
	Prog struct {
		Cfg  Config
		name string
		dir  struct {
			cache  string
			config string
		}
	}
	pipeIO struct {
		outEncoder *json.Encoder
		outWriter  *bufio.Writer
	}
)

func Bad(what, which string) {
	panic(strf("%s: invalid %s %s '%s'", Prog.name, Lang.Title, what, which))
}

func Init() (err error) {
	Prog.name = os.Args[0]
	Prog.dir.config = filepath.Join(usys.UserDataDirPath(false), Prog.name)
	Prog.dir.cache = filepath.Join(usys.UserDataDirPath(true), Prog.name)
	if err = ufs.EnsureDirExists(Prog.dir.config); err != nil {
		return
	} else if err = ufs.EnsureDirExists(Prog.dir.cache); err != nil {
		return
	}
	if Prog.Cfg.reload(); Prog.Cfg.err == nil {
		for _, preDefinedHandler := range handlers {
			preDefinedHandler.Init()
		}
	}
	return
}

func InitAndServeOrPanic(onPreInit func(), onPostInit func()) {
	// note to self: don't ADD any further logic in here, either in Init() or in Serve()
	onPreInit()
	err := Init()
	if err == nil {
		onPostInit()
		err = Serve()
	}
	if err != nil {
		panic(err)
	}
}

func catch(err *error) {
	if except := recover(); except != nil {
		if e, ok := except.(error); ok {
			*err = e
		} else {
			*err = fmt.Errorf("%v", except)
		}
	}
}

func Serve() (err error) {
	var stdin *bufio.Scanner
	stdin, pipeIO.outWriter, pipeIO.outEncoder = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	// we don't directly wire up a json.Decoder to stdin but read individual lines in as strings first:
	// - this enforces our line-based protocol
	// - allows decoding in separate go-routine
	// - bad lines are simply reported without decoder in confused/error state; and without needing to exit
	defer catch(&err)
	for stdin.Scan() {
		if err = stdin.Err(); err == nil {
			go serveIncomingReq(stdin.Text())
		} else {
			return
		}
	}
	return
}

func serveIncomingReq(jsonreq string) {
	// err only covers: either resp couldn't be json-encoded, or stdout write/flush problem
	// (both would indicate bigger problems --- still recover()ed in Serve() though)
	// any other kind of error, reqDecodeAndRespond will record into resp.ErrMsg to report it back to the client
	resp := reqDecodeAndRespond(jsonreq)
	err := pipeIO.outEncoder.Encode(resp)
	if err == nil {
		err = pipeIO.outWriter.Flush()
	}
	if err != nil {
		panic(err)
	}
}
