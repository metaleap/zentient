package zgo
import (
	"github.com/metaleap/go-devgo"

	"github.com/metaleap/zentient/z"
)


type zgo struct {
	z.Base
}


var (
	srcDir string
)


func Init () z.Zengine {
	if !devgo.HasGoDevEnv() { return nil }
	srcDir = z.Ctx.SrcDir
	self := &zgo{}
	self.Base.Init()
	go devgo.RefreshPkgs()
	return self
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
	case "diag":
		caps = []*z.RespCmd	{	&z.RespCmd { Title: "go install",	Exists: devgo.HasGoDevEnv(),	Hint: "check your Go installation" },
								// &z.RespCmd { Title: "go list",	Exists: devgo.HasGoDevEnv(),	Hint: "check your Go installation" },
								&z.RespCmd { Title: "go vet",		Exists: devgo.HasGoDevEnv(),	Hint: "check your Go installation" },
								&z.RespCmd { Title: "golint",		Exists: devgo.Has_golint,		Hint: "`go get -u github.com/golang/lint/golint`" },
								&z.RespCmd { Title: "ineffassign",	Exists: devgo.Has_ineffassign,	Hint: "`go get -u github.com/gordonklaus/ineffassign`" },
								&z.RespCmd { Title: "aligncheck",	Exists: devgo.Has_checkalign,	Hint: "`github.com/opennota/check`" },
								&z.RespCmd { Title: "structcheck",	Exists: devgo.Has_checkstruct,	Hint: "`github.com/opennota/check`" },
								&z.RespCmd { Title: "varcheck",		Exists: devgo.Has_checkvar,		Hint: "`github.com/opennota/check`" },
								&z.RespCmd { Title: "interfacer",	Exists: devgo.Has_golint,		Hint: "`github.com/mvdan/interfacer`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, custcmd string, tabsize uint8) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, custcmd, z.RespCmd { Exists: devgo.Has_gofmt, Name: "gofmt", Args: []string{"-e", "-s"} })
}

func (self *zgo) OnFileOpen (file *z.File) {
}

func (self *zgo) OnFileClose (file *z.File) {
}

func (self *zgo) OnFileWrite (file *z.File) {
	go devgo.RefreshPkgs()
}
