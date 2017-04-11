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

func onFilesClosed (µ Zengine, relpaths []string) {
	for i,_ := range relpaths { relpaths[i] = filepath.FromSlash(relpaths[i]) }
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpaths...)
}

func onFilesOpened (µ Zengine, relpaths []string) {
	for _,relpath := range relpaths {
		relpath = filepath.FromSlash(relpath)
		file := AllFiles[relpath]  ;  if file == nil {
			file = NewFile(µ, relpath)  ;  µ.OnFile(file)  ;  AllFiles[relpath] = file
		}
		if isopened := !uslice.StrHas(OpenFiles, relpath) ; isopened {  OpenFiles = append(OpenFiles, relpath)  }
	}
}

func onFilesWritten (µ Zengine, relpaths []string) {
	for i,_ := range relpaths { relpaths[i] = filepath.FromSlash(relpaths[i]) }
	µ.B().buildFrom(µ, relpaths)
}
