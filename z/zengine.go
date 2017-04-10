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
	Linters ([]string) []func()map[string][]*RespDiag
	ReadyToBuildAndLint () bool
	BuildFrom ([]string) map[string][]*RespDiag
	OnFile (*File)
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
}

func onFileOpen (µ Zengine, relpath string) {
	relpath = filepath.FromSlash(relpath)
	file := AllFiles[relpath]  ;  if file == nil {
		file = NewFile(µ, relpath)
		µ.OnFile(file)
		AllFiles[relpath] = file
	}
	if isnew := !uslice.StrHas(OpenFiles, relpath) ; isnew {
		OpenFiles = append(OpenFiles, relpath)
	}
}

func onFileWrite (µ Zengine, relpath string) {
	relpath = filepath.FromSlash(relpath)
	µ.B().buildFrom(µ, relpath)
}


// func refreshAllDiags() {
// 	funcs := []func() {}
// 	// for _,zeng := range Zengines { µ := zeng  ;  funcs = append(funcs, func() { µ.B().RefreshDiags(µ, "", "") }) }
// 	ugo.WaitOn(funcs...)
// }
