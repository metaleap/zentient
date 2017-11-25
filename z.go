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
	err := Init("go")
	if err == nil {
		err = Serve()
	}
	if err != nil {
		panic(err)
	}
}

func Serve() (err error) {
	var resp *MsgResp
	var jsonresp string
	var stdin *bufio.Scanner

	stdin, PipeIO.RawOut, PipeIO.Out = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	for stdin.Scan() {
		resp = reqDecodeAndHandle(stdin.Text())
		if jsonresp, err = resp.encode(); err != nil {
			return
		} else if _, err = fmt.Println(jsonresp); err != nil {
			return
		}
	}
	return
}
