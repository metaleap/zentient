package z
import (
	"github.com/metaleap/go-util-slice"
)


type RootInfo struct {
	SrcDir      string
	CacheDir    string
	ConfigDir   string
}


type Zengine interface {
	Ids () []string
	Jsonish () interface{}

	Base () *ZengineBase

	Caps (string) []string
	OnFileActive (*File)
	OnFileClose (*File)
	OnFileOpen (*File)
	OnFileWrite (*File)
}


type ZengineBase struct {
	Projs []*Proj
}


var (
	Root        = &RootInfo{}
	AllFiles    = map[string]*File {}
	OpenFiles   = []string {}
	Zengines    = map[string]Zengine {}
)


func InitZBase (base *ZengineBase) {
	base.Projs = []*Proj {}
}


func fromZidMsg (msgargs string) (z Zengine, argstr string) {
	zid := msgargs[:2]
	if z = Zengines[zid] ; z != nil {
		argstr = msgargs[3:]
	}
	return
}

func onFileActive (file* File) {
	file.Z.OnFileActive(file)
}

func onFileClose (z Zengine, relpath string) {
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpath)
	file := AllFiles[relpath]
	z.OnFileClose(file)
}

func onFileOpen (z Zengine, relpath string) {
	uslice.StrAppendUnique(&OpenFiles, relpath)
	file := AllFiles[relpath]
	if file == nil {
		file = NewFile(z, relpath)
		AllFiles[relpath] = file
		z.OnFileOpen(file)
	}
	onFileActive(file)
}

func onFileWrite (z Zengine, relpath string) {
	file := AllFiles[relpath]
	if (file == nil) {
		onFileOpen(z, relpath)
		file = AllFiles[relpath]
	}
	z.OnFileWrite(file)
}
