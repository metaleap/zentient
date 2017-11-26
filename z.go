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

	Lang struct {
		Enabled bool
		ID      string
		Title   string
		CodeFmt iCodeFormatting

		cmdProviders []iMetaCmds
	}

	Prog struct {
		Name     string
		CacheDir string
	}

	PipeIO struct {
		OutEncoder *json.Encoder
		OutWriter  *bufio.Writer
	}
)

func Bad(what, which string) {
	panic(strf("%s: bad %s %s '%s'", Prog.Name, Lang.Title, what, which))
}

func Init() (err error) {
	Prog.CacheDir = filepath.Join(usys.UserDataDirPath(), os.Args[0])
	err = ufs.EnsureDirExists(Prog.CacheDir)

	metaCmdsProvidersInit()
	return
}

func InitAndServeOrPanic(onPreInit func(), onPostInit func()) {
	// note to self: don't ADD any further logic in here, either in Init() or in Serve()
	Prog.Name = os.Args[0]
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
	stdin, PipeIO.OutWriter, PipeIO.OutEncoder = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	// we don't directly wire up a json.Decoder to stdin but read individual lines in as strings first:
	// - this enforces our line-based protocol
	// - allows decoding in separate go-routine
	// - bad lines are simply reported without decoder in confused/error state or needing to quit
	defer catch(&err)
	for stdin.Scan() {
		if err = stdin.Err(); err != nil {
			return
		}
		go serveIncomingReq(stdin.Text())
	}
	return
}

func serveIncomingReq(jsonreq string) {
	// err only covers: either resp couldn't be encoded, or stdout problem
	// all other errors are captured in resp.ErrMsg
	resp := reqDecodeAndRespond(jsonreq)
	err := PipeIO.OutEncoder.Encode(resp)
	if err == nil {
		err = PipeIO.OutWriter.Flush()
	}
	if err != nil {
		panic(err)
	}
}
