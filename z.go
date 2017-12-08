package z

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/sys"
)

const (
	pretendSlow = false
)

var (
	Err  = errors.New
	Errf = fmt.Errorf
	Strf = fmt.Sprintf

	menuProviders []iMenuProvider
	dispatchers   = []iDispatcher{
		&mainMenu{},
	}

	Lang struct {
		Enabled  bool
		ID       string
		Title    string
		SrcMod   iSrcMod
		SrcIntel iSrcIntel
		Extras   iExtras
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

func Bad(what string, which string) {
	panic(Strf("%s: invalid %s %s '%s'", Prog.name, Lang.Title, what, which))
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
		wellknowndispatchers := []iDispatcher{
			Lang.SrcIntel, Lang.SrcMod, Lang.Extras,
		}
		for _, disp := range wellknowndispatchers {
			if disp != nil {
				dispatchers = append(dispatchers, disp)
				disp.Init()
			}
		}

		wellknownmenuproviders := []iMenuProvider{
			Lang.SrcMod,
		}
		for _, menu := range wellknownmenuproviders {
			if menu != nil {
				menuProviders = append(menuProviders, menu)
			}
		}
	}
	return
}

func InitAndServe(onPreInit func(), onPostInit func()) (err error) {
	// note to self: don't ADD any further logic in here, do it only in either Init() or Serve()
	onPreInit()
	if err = Init(); err == nil {
		onPostInit()
		err = Serve()
	}
	return
}

func InitAndServeOrPanic(onPreInit func(), onPostInit func()) {
	if err := InitAndServe(onPreInit, onPostInit); err != nil {
		panic(err)
	}
}

func catch(err *error) {
	if except := recover(); except != nil {
		if e, ok := except.(error); ok {
			*err = e
		} else {
			*err = Errf("%v", except)
		}
	}
}

func Serve() (err error) {
	var stdin *bufio.Scanner
	stdin, pipeIO.outWriter, pipeIO.outEncoder = urun.SetupJsonProtoPipes(1024*1024*4, false, true)

	// we allow our sub-ordinate go-routines to panic and just before we return, we recover() the `err` to return (if any)
	defer catch(&err)

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
	if pretendSlow {
		time.Sleep(time.Millisecond * 2345) // temporary artificial delay to briefly see client-side progress-bars etc
	}
	resp := reqDecodeAndRespond(jsonreq)

	// err only covers: either resp couldn't be json-encoded, or stdout write/flush problem
	// (both would indicate bigger problems --- still recover()ed in Serve() though)
	// any other kind of error, reqDecodeAndRespond will record into resp.ErrMsg to report it back to the client
	err := pipeIO.outEncoder.Encode(resp)
	if err == nil {
		err = pipeIO.outWriter.Flush()
	}
	if err != nil {
		panic(err)
	}
}
