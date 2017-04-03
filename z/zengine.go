package z
import (
	"path/filepath"

	"github.com/metaleap/go-util-slice"
)


type Context struct {
	SrcDir		string
	CacheDir	string
	ConfigDir	string
}


type Zengine interface {
	EdLangIDs () []string
	B () *Base

	Caps (string) []*RespCmd
	DoFmt (string, string, uint8) (*RespFmt, error)
	Lint ([]string) map[string][]*RespDiag
	OnFileClose (*File)
	OnFileOpen (*File)
	OnFileWrite (*File)
	BuildFrom (string) map[string][]*RespDiag
}


var (
	Ctx			= &Context{}
	AllFiles	= map[string]*File {}
	OpenFiles	= []string {}
	Zengines	= map[string]Zengine {}
)


func doFmt (zid string, reqsrc string, reqcmd string, reqtabsize uint8) (resp map[string]*RespFmt, err error) {
	if µ := Zengines[zid] ; µ != nil && len(reqsrc)>0 {
		var rfmt *RespFmt
		if rfmt,err = µ.DoFmt(reqsrc, reqcmd, reqtabsize) ; rfmt!=nil && err==nil {
			resp = map[string]*RespFmt { zid: rfmt }
		}
	}
	return
}

func onFileClose (µ Zengine, relpath string) {
	relpath = filepath.FromSlash(relpath)
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpath)
	µ.B().refreshDiags(µ, relpath, "", "")
	µ.OnFileClose(AllFiles[relpath])
}

func onFileOpen (µ Zengine, relpath string) {
	relpath = filepath.FromSlash(relpath)
	file := AllFiles[relpath]
	if file == nil {
		file = NewFile(µ, relpath)
		AllFiles[relpath] = file
	}
	if isnew := !uslice.StrHas(OpenFiles, relpath) ; isnew {
		OpenFiles = append(OpenFiles, relpath)
		µ.B().refreshDiags(µ, "", relpath, "")
	}
	µ.OnFileOpen(file)
}

func onFileWrite (µ Zengine, relpath string) {
	relpath = filepath.FromSlash(relpath)
	file := AllFiles[relpath]
	µ.B().refreshDiags(µ, "", "", relpath)
	µ.OnFileWrite(file)
}
