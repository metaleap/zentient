package main
import (
	"bytes"
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"

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
	stdin.Split(func(data []byte, ateof bool) (advance int, token []byte, err error) {
		if i_cl1 := bytes.Index(data, []byte("Content-Length: "))  ;  i_cl1>=0 {
			datafromclen := data[i_cl1+16:]  ;  if i_cl2 := bytes.IndexAny(datafromclen, "\r\n")  ;  i_cl2>0 {
				if clen := ustr.ToInt(string(datafromclen[:i_cl2]))  ;  clen<=0 { panic(clen) } else {
					if i_js1 := bytes.Index(datafromclen, []byte("{\""))  ;  i_js1 > i_cl2 {
						if i_js2 := i_js1+clen  ;  len(datafromclen)>=i_js2 {
							advance = i_cl1 + 16 + i_js2  ;  token = datafromclen[i_js1:i_js2]
						}
					}
				}
			}
		}
		return
	})
	var reqln string
	for stdin.Scan() {
		reqln = stdin.Text()
		logfile.WriteString(reqln+"\n")
		logfile.Sync()
		rawOut.WriteString(reqln)
		rawOut.Flush()
	}

}
