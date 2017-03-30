package zhs
import (
	"fmt"

	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devhs"
)


type zhs struct {
	z.Base
}

var (
	µ *zhs
)


func New (root *z.RootInfo) z.Zengine {
	if !devhs.HasHsDevEnv() { return nil }

	µ = &zhs{}
	µ.Base.Init()
	return µ
}




func (_ *zhs) Ids () []string {
	return []string { "haskell", "Haskell" }
}



func (_ *zhs) Caps (cap string) (caps []*z.CmdInfo) {
	switch cap {
	case "fmt":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "stylish-haskell", I: devhs.Has_stylish_haskell, H: "`stack install stylish-haskell`" },
								&z.CmdInfo { N: "hindent", I: devhs.Has_hindent, H: "`stack install hindent`" },
								&z.CmdInfo { N: "brittany", I: devhs.Has_brittany, H: "`github.com/lspitzner/brittany`" },
							}
	case "lint":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "hlint", I: devhs.Has_hlint, H: "`stack install hlint`" },
							}
	}
	return caps
}

func (self *zhs) DoFmt (src string, custcmd string, tabsize int) (resp *z.RespFmt, err error) {
	ts := fmt.Sprint(tabsize)
	return self.Base.DoFmt(src, custcmd,
		z.CmdInfo { I: devhs.Has_stylish_haskell,	C: "stylish-haskell",	A: []string{} },
		z.CmdInfo { I: devhs.Has_hindent,			C: "hindent",			A: []string{"--no-force-newline", "--indent-size", ts} },
		z.CmdInfo { I: devhs.Has_brittany,			C: "brittany",			A: []string{"--indent", ts} },
		)
}

func (_ *zhs) OnFileActive (file *z.File) {
}

func (_ *zhs) OnFileOpen (file *z.File) {
}

func (_ *zhs) OnFileClose (file *z.File) {
}

func (_ *zhs) OnFileWrite (file *z.File) {
}
