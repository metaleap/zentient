package zhs
import (
	"fmt"

	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devhs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)


type zhs struct {
	z.ZengineBase
}

var (
	µ *zhs
)


func New (root *z.RootInfo) z.Zengine {
	if !devhs.HasHsDevEnv() { return nil }

	µ = &zhs{}
	z.InitZBase(&µ.ZengineBase)
	return µ
}




func (_ *zhs) Ids () []string {
	return []string { "haskell", "Haskell" }
}

func (self *zhs) Jsonish () interface{} {
	return self
}


func (self *zhs) Base () *z.ZengineBase {
	return &self.ZengineBase
}



func (_ *zhs) Caps (string) []*z.RespCap {
	caps := []*z.RespCap {}

	caps = append(caps, &z.RespCap { Name: "stylish-haskell", Available: devhs.Has_stylish_haskell, InstHint: "`stack install stylish-haskell`" })
	caps = append(caps, &z.RespCap { Name: "hindent", Available: devhs.Has_hindent, InstHint: "`stack install hindent`" })
	caps = append(caps, &z.RespCap { Name: "brittany", Available: devhs.Has_brittany, InstHint: "`github.com/lspitzner/brittany`" })

	return caps
}

func (_ *zhs) DoFmt (src string, cmd string, tabsize int) (resp *z.RespFmt, err error) {
	var warns string
	resp = &z.RespFmt{}
	if (len(cmd)>0) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", cmd)
	} else if (devhs.Has_stylish_haskell) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", "stylish-haskell")
	} else if (devhs.Has_hindent) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", "hindent", "--no-force-newline", "--indent-size", fmt.Sprint(tabsize))
	} else if (devhs.Has_brittany) {
		resp.Result, warns, err = ugo.CmdExecStdin(src, "", "brittany", "--indent", fmt.Sprint(tabsize))
	} else {
		resp = nil
	}
	if (resp != nil) {
		resp.Warnings = ustr.Split(warns, "\n")
	}
	return
}

func (_ *zhs) OnFileActive (file *z.File) {
}

func (_ *zhs) OnFileOpen (file *z.File) {
}

func (_ *zhs) OnFileClose (file *z.File) {
}

func (_ *zhs) OnFileWrite (file *z.File) {
}
