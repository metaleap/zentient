package z
import (
	"path/filepath"

	"github.com/metaleap/go-util-misc"
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
	DoRename (string, string, uint64, string, string, string, uint64, uint64) (map[string][]*RespRen, error)
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


func doFmt (zid string, reqsrc string, reqcmd string, reqtabsize uint8) (resp *RespFmt, err error) {
	if µ := Zengines[zid]  ;  µ==nil || len(reqsrc)==0 { err = ugo.E("Bad zid or input src") } else { resp,err = µ.DoFmt(reqsrc, reqcmd, reqtabsize) }
	return
}

func doRename (zid string, reqcmd string, relfilepath string, offset uint64, newname string, eol string, oldname string, off1 uint64, off2 uint64) (resp map[string][]*RespRen, err error) {
	µ := Zengines[zid]   ;  if µ==nil {  err = ugo.E("Bad zid: " + zid)  }  ;  if len(newname)==0 {  err = ugo.E("No newname given")  }
	resp,err = µ.DoRename(reqcmd, relfilepath, offset, newname, eol, oldname, off1, off2)
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
