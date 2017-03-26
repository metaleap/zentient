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


func (_ *zgo) Caps (string) []string {
	caps := []string {}
	if devgo.GoFmt { caps = append(caps, "gofmt") }
	return caps
}

func (_ *zgo) DoFmt (src string) *z.FmtResp {
	var warns string
	resp := &z.FmtResp{}
	resp.Result, warns, resp.Error = ugo.CmdExecStdin(src, "", "gofmt", "-e", "-s")

	resp.Warnings = ustr.Split(warns, "\n")
	return resp
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (_ *zgo) OnFileOpen (file *z.File) {
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
