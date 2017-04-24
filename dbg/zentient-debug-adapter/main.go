package main
import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/metaleap/go-util-fs"
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

	logdirpath := filepath.Join(z.Ctx.ConfigDir, "zdbglog")
	ufs.EnsureDirExists(logdirpath)
	logfilepath := filepath.Join(logdirpath, "log" + ugo.SPr(time.Now().UnixNano()) + ".log")
	logfile,err := os.Create(logfilepath)  ;  if err!=nil { panic(err) } else { defer logfile.Close() }

	stdin,rawOut,jOut = ugo.SetupJsonProtoPipes(1024*1024*4)
	var reqln string
	for stdin.Scan() {
		reqln = stdin.Text()
		logfile.WriteString("⟨" + reqln+"⟩")
		logfile.Sync()
		// rawOut.WriteString(reqln)
		// rawOut.Flush()
	}

}
