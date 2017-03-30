package z
import (
	"github.com/metaleap/go-util-slice"
)


type RootInfo struct {
	SrcDir		string
	CacheDir	string
	ConfigDir	string
}

type RespCap struct {
	Name		string
	Available	bool
	InstHint	string
}

type RespFmt struct {
	Result		string
	Warnings	[]string
}


type Zengine interface {
	Ids () []string
	Jsonish () interface{}

	Base () *ZengineBase

	Caps (string) []*RespCap
	DoFmt (string, string, int) (*RespFmt, error)
	OnFileActive (*File)
	OnFileClose (*File)
	OnFileOpen (*File)
	OnFileWrite (*File)
}


type ZengineBase struct {
	Projs []*Proj
}


var (
	Root		= &RootInfo{}
	AllFiles	= map[string]*File {}
	OpenFiles	= []string {}
	Zengines	= map[string]Zengine {}
)


func (self *ZengineBase) Init () {
	self.Projs = []*Proj {}
}


func doFmt (zid string, reqsrc string, reqcmd string, reqtabsize int) (resp map[string]*RespFmt, err error) {
	if µ := Zengines[zid] ; µ != nil && len(reqsrc)>0 {
		var rfmt *RespFmt
		if rfmt,err = µ.DoFmt(reqsrc, reqcmd, reqtabsize) ; rfmt!=nil && err==nil {
			resp = map[string]*RespFmt { zid: rfmt }
		}
	}
	return
}

func onFileActive (file* File) {
	file.Z.OnFileActive(file)
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
