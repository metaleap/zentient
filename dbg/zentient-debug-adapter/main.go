package main
import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/metaleap/go-util-misc"

	"github.com/metaleap/zentient/z"
	"github.com/metaleap/zentient/go"
	"github.com/metaleap/zentient/hs"
)

var (
	stdin *bufio.Scanner
	rawOut *bufio.Writer
	jOut *json.Encoder
)

func main () {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := z.Init(map[string]func()z.Zengine { "go": zgo.Init, "hs": zhs.Init })
	if err!=nil { panic(err) }

	logfilepath := filepath.Join(z.Ctx.ConfigDir, "log" + ugo.SPr(time.Now().UnixNano()) + ".log")
	panic(logfilepath)
	logfile,err := os.Create(logfilepath)  ;  if err!=nil { panic(err) } else { defer logfile.Close() }
	stdin,rawOut,jOut = ugo.SetupJsonProtoPipes(1024*1024*4)

}
