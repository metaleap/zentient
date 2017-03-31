package z
import (
	"github.com/metaleap/go-util-slice"
)


type Context struct {
	SrcDir		string
	CacheDir	string
	ConfigDir	string
}

type RespCmd struct {
	Name	string		//	actual cmd name
	Args	[]string	//	args

	Title	string		//	display name, eg: N = "go vet" when C = "go" with A = ["vet"]  ;  if empty fall back to C
	Exists	bool		//	installed?
	Hint	string		//	install hint

	f	func()		//	tmp field used in Base.DoFmt()
}

type RespFmt struct {
	Result		string
	Warnings	[]string
}


type Zengine interface {
	Ids () []string

	Caps (string) []*RespCmd
	DoFmt (string, string, int) (*RespFmt, error)
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
)



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
