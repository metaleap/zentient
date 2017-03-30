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


func (_ *zgo) Caps (cap string) (caps []*z.RespCap) {
	switch cap {
	case "fmt":
		caps = []*z.RespCap	{	&z.RespCap { Name: "gofmt", Available: devgo.Has_gofmt, InstHint: "check your Go installation" },
							}
	case "lint":
		caps = []*z.RespCap	{	&z.RespCap { Name: "go vet", Available: devgo.HasGoDevEnv(), InstHint: "check your Go installation" },
								&z.RespCap { Name: "golint", Available: devgo.Has_golint, InstHint: "`go get -u github.com/golang/lint/golint`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, cmd string, tabsize int) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, cmd, z.Fmt{	N: "gofmt", I: devgo.Has_gofmt, C: "gofmt", A: []string{"-e", "-s"} })
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (_ *zgo) OnFileOpen (file *z.File) {
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
