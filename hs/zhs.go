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
	µ.ZengineBase.Init()
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



func (_ *zhs) Caps (cap string) (caps []*z.RespCap) {
	switch cap {
	case "fmt":
		caps = []*z.RespCap	{	&z.RespCap { Name: "stylish-haskell", Available: devhs.Has_stylish_haskell, InstHint: "`stack install stylish-haskell`" },
								&z.RespCap { Name: "hindent", Available: devhs.Has_hindent, InstHint: "`stack install hindent`" },
								&z.RespCap { Name: "brittany", Available: devhs.Has_brittany, InstHint: "`github.com/lspitzner/brittany`" },
							}
	case "lint":
		caps = []*z.RespCap	{	&z.RespCap { Name: "hlint", Available: devhs.Has_hlint, InstHint: "`stack install hlint`" },
							}
	}
	return caps
}

func (_ *zhs) DoFmt (src string, cmd string, tabsize int) (resp *z.RespFmt, err error) {
	var warns string
	var warnlns = true
	resp = &z.RespFmt{}
	do_stylish	:= func() { resp.Result, warns, err = ugo.CmdExecStdin(src, "", "stylish-haskell") }
	do_hindent	:= func() { resp.Result, warns, err = ugo.CmdExecStdin(src, "", "hindent", "--no-force-newline", "--indent-size", fmt.Sprint(tabsize)) }
	do_brittany	:= func() { resp.Result, warns, err = ugo.CmdExecStdin(src, "", "brittany", "--indent", fmt.Sprint(tabsize)) }
	do_custom	:= func() { resp.Result, warns, err = ugo.CmdExecStdin(src, "", cmd) }
	switch cmd {
	case "stylish-haskell":
		do_stylish()
	case "hindent":
		do_hindent()
	case "brittany":
		do_brittany()
	default:
		if len(cmd)>0 {
			do_custom()
		} else if devhs.Has_stylish_haskell {
			do_stylish()
		} else if devhs.Has_hindent {
			do_hindent()
		} else if devhs.Has_brittany {
			do_brittany()
		} else {
			resp = nil
		}
	}
	if resp!=nil {
		if warnlns { resp.Warnings = ustr.Split(warns, "\n") } else { resp.Warnings = []string { warns } }
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
