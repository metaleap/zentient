package z

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/sys"
)

type Context struct {
	LangID   string
	CacheDir string
}

var (
	Ctx    Context
	PipeIO struct {
		Out    *json.Encoder
		RawOut *bufio.Writer
	}
)

func Init(langID string) (err error) {
	Ctx.LangID = langID
	Ctx.CacheDir = filepath.Join(usys.UserDataDirPath(), os.Args[0])
	err = ufs.EnsureDirExists(Ctx.CacheDir)
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
	var respjson string
	var stdin *bufio.Scanner

	stdin, PipeIO.RawOut, PipeIO.Out = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	for stdin.Scan() {
		resp = handleReq(stdin.Text())
		if respjson, err = resp.encode(); err != nil {
			return
		} else {
			println(respjson)
		}
	}
	return
}
