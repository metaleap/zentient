package zgo
import (
	"github.com/metaleap/go-devgo"

	"github.com/metaleap/zentient/z"
)


type zgo struct {
	z.Base
}


var (
	µ *zgo
)


func Init (ctx *z.Context) z.Zengine {
	if !devgo.HasGoDevEnv() {
		return nil
	}
	µ = &zgo{}
	µ.Base.Init()
	go devgo.RefreshPkgs(func(errs []error) {
		µ.Base.DbgObjs = append(µ.Base.DbgObjs, devgo.PkgsByDir, devgo.PkgsByImP)
		for _,err:= range errs { µ.Base.DbgMsgs = append(µ.Base.DbgMsgs, err.Error()) }
	})
	return µ
}


func (_ *zgo) EdLangIDs () []string {
	return []string { "go", "Go" }
}

func (self *zgo) B () *z.Base {
	return &self.Base
}


func (_ *zgo) Caps (cap string) (caps []*z.RespCmd) {
	switch cap {
	case "fmt":
		caps = []*z.RespCmd	{	&z.RespCmd { Title: "gofmt",	Exists: devgo.Has_gofmt,		Hint: "check your Go installation" },
							}
	case "lint":
		caps = []*z.RespCmd	{	&z.RespCmd { Title: "go vet",	Exists: devgo.HasGoDevEnv(),	Hint: "check your Go installation" },
								&z.RespCmd { Title: "golint",	Exists: devgo.Has_golint,		Hint: "`go get -u github.com/golang/lint/golint`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, custcmd string, tabsize uint8) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, custcmd, z.RespCmd { Exists: devgo.Has_gofmt, Name: "gofmt", Args: []string{"-e", "-s"} })
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (self *zgo) OnFileOpen (file *z.File) {
	self.Base.Diags[file.RelPath] = []*z.RespDiag {
		&z.RespDiag { Code: "W", Msg: "Mock warning for " + file.RelPath, PosLn: 2, PosCol: 1, Sev: z.DIAG_WARN, Cat: "devgo" },
		&z.RespDiag { Code: "I", Msg: "Mock info for " + file.RelPath, PosLn: 9, PosCol: 11, Sev: z.DIAG_INFO, Cat: "devgo" },
	}
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
