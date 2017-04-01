package z
import (
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
	OnFileActive (*File)
	OnFileClose (*File)
	OnFileOpen (*File)
	OnFileWrite (*File)
}


var (
	Ctx			= &Context{}
	AllFiles	= map[string]*File {}
	OpenFiles	= []string {}
	Zengines	= map[string]Zengine {}

	allDiags	= map[string]map[string][]*RespDiag {}
)


func AllDiags () map[string]map[string][]*RespDiag {
	for zid,µ := range Zengines {
		allDiags[zid] = µ.B().Diags
	}
	return allDiags
}


func doFmt (zid string, reqsrc string, reqcmd string, reqtabsize uint8) (resp map[string]*RespFmt, err error) {
	if µ := Zengines[zid] ; µ != nil && len(reqsrc)>0 {
		var rfmt *RespFmt
		if rfmt,err = µ.DoFmt(reqsrc, reqcmd, reqtabsize) ; rfmt!=nil && err==nil {
			resp = map[string]*RespFmt { zid: rfmt }
		}
	}
	return
}

func onFileActive (file* File) {
	file.µ.OnFileActive(file)
}

func onFileClose (µ Zengine, relpath string) {
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpath)
	file := AllFiles[relpath]
	µ.OnFileClose(file)
}

func onFileOpen (µ Zengine, relpath string) {
	uslice.StrAppendUnique(&OpenFiles, relpath)
	file := AllFiles[relpath]
	if file == nil {
		file = NewFile(µ, relpath)
		AllFiles[relpath] = file
		µ.OnFileOpen(file)
	}
	onFileActive(file)
}

func onFileWrite (µ Zengine, relpath string) {
	file := AllFiles[relpath]
	if (file == nil) {
		onFileOpen(µ, relpath)
		file = AllFiles[relpath]
	}
	µ.OnFileWrite(file)
}
