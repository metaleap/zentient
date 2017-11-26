package z

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/sys"
)

var (
	strf = fmt.Sprintf

	LangID   string
	CacheDir string
	PipeIO   struct {
		Out    *json.Encoder
		RawOut *bufio.Writer
	}
)

func Init(langID string) (err error) {
	LangID = langID
	CacheDir = filepath.Join(usys.UserDataDirPath(), os.Args[0])
	err = ufs.EnsureDirExists(CacheDir)
	return
}

func InitAndServeOrPanic(langID string) {
	if err := Init("go"); err != nil {
		panic(err)
	}
	Serve()
}

func Serve() {
	var stdin *bufio.Scanner
	stdin, PipeIO.RawOut, PipeIO.Out = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	for stdin.Scan() {
		go Handle(stdin.Text())
	}
}

func Handle(jsonreq string) {
	resp := reqDecodeAndRespond(jsonreq)
	time.Sleep(time.Millisecond * 444)
	if jsonresp, err := resp.encode(); err != nil {
		panic(err)
	} else if _, err = fmt.Println(jsonresp); err != nil {
		panic(err)
	}
}
