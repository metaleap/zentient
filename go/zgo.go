package zgo
import (
	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devgo"
)


type zgo struct {
	z.Base
}

var (
	µ *zgo
)


func New (root *z.RootInfo) z.Zengine {
	if !devgo.HasGoDevEnv() { return nil }

	µ = &zgo{}
	µ.Base.Init()
	return µ
}




func (_ *zgo) Ids () []string {
	return []string { "go", "Go" }
}


func (_ *zgo) Caps (cap string) (caps []*z.CmdInfo) {
	switch cap {
	case "fmt":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "gofmt", I: devgo.Has_gofmt, H: "check your Go installation" },
							}
	case "lint":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "go vet", I: devgo.HasGoDevEnv(), H: "check your Go installation" },
								&z.CmdInfo { N: "golint", I: devgo.Has_golint, H: "`go get -u github.com/golang/lint/golint`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, custcmd string, tabsize int) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, custcmd, z.CmdInfo { I: devgo.Has_gofmt, C: "gofmt", A: []string{"-e", "-s"} })
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (_ *zgo) OnFileOpen (file *z.File) {
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
