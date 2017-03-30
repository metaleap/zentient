package zgo
import (
	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)


type zgo struct {
	z.ZengineBase
}

var (
	µ *zgo
)


func New (root *z.RootInfo) z.Zengine {
	if !devgo.HasGoDevEnv() { return nil }

	µ = &zgo{}
	µ.ZengineBase.Init()
	return µ
}




func (_ *zgo) Ids () []string {
	return []string { "go", "Go" }
}

func (self *zgo) Jsonish () interface{} {
	return self
}


func (self *zgo) Base () *z.ZengineBase {
	return &self.ZengineBase
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

func (_ *zgo) DoFmt (src string, cmd string, tabsize int) (resp *z.RespFmt, err error) {
	var warns string
	resp = &z.RespFmt{}
	if (len(cmd)>0) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", cmd)
	} else if (devgo.Has_gofmt) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", "gofmt", "-e", "-s")
	} else {
		resp = nil
	}
	if (resp != nil) {
		resp.Warnings = ustr.Split(warns, "\n")
	}
	return
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (_ *zgo) OnFileOpen (file *z.File) {
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
