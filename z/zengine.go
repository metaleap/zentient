package z
import (
	"github.com/metaleap/go-util-slice"
)


type RootInfo struct {
	SrcDir		string
	CacheDir	string
	ConfigDir	string
}

type FmtResp struct {
	Result		string
	Warnings	[]string
	Error		error
}


type Zengine interface {
	Ids () []string
	Jsonish () interface{}

	Base () *ZengineBase

	Caps (string) []string
	DoFmt (string) *FmtResp
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


func InitZBase (base *ZengineBase) {
	base.Projs = []*Proj {}
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
