package z
import (
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
	OnFileActive (*File)
	OnFileClose (*File)
	OnFileOpen (*File)
	OnFileWrite (*File)
	RefreshDiags(string, []string) map[string][]*RespDiag
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
	refreshAllDiags("")
}

func onFileClose (µ Zengine, relpath string) {
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpath)
	file := AllFiles[relpath]
	µ.OnFileClose(file)
	refreshAllDiags("")
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
	refreshAllDiags(relpath)
}

func each (fn func (Zengine) func()) (funcs []func()) {
	for _,µ := range Zengines {  funcs = append(funcs, fn(µ))  }
	return
}

func refreshAllDiags (rebuildfilerelpath string) {
	funcs := each(func(µ Zengine) func() { return func() {
		µ.B().Diags = µ.B().refreshDiags(µ, rebuildfilerelpath)
	} })
	ugo.WaitOn(funcs...)
}
