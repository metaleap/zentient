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
	z.InitZBase(&µ.ZengineBase)
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


func (_ *zgo) Caps (string) []*z.RespCap {
	caps := []*z.RespCap {}

	caps = append(caps, &z.RespCap { Name: "gofmt", Available: true, InstHint: "check your Go installation" })

	return caps
}

func (_ *zgo) DoFmt (src string) (resp *z.RespFmt, err error) {
	var warns string
	resp = &z.RespFmt{}
	resp.Result, warns, err = ugo.CmdExecStdin(src, "", "gomt", "-e", "-s")
	resp.Warnings = ustr.Split(warns, "\n")
	// resp = nil
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
