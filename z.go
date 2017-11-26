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
	if err := Init(); err != nil {
		panic(err)
	}
	onPostInit()
	Serve()
}

func Serve() {
	var stdin *bufio.Scanner
	stdin, PipeIO.OutWriter, PipeIO.OutEncoder = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	for stdin.Scan() {
		go serveIncomingReq(stdin.Text())
	}
}

func serveIncomingReq(jsonreq string) {
	resp := reqDecodeAndRespond(jsonreq)
	err := PipeIO.OutEncoder.Encode(resp)
	if err == nil {
		err = PipeIO.OutWriter.Flush()
	}
	if err != nil {
		panic(err)
	}
}
