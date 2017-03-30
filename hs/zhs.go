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

func (self *zhs) DoFmt (src string, cmd string, tabsize int) (resp *z.RespFmt, err error) {
	ts := fmt.Sprint(tabsize)
	return self.Base.DoFmt(src, cmd,
		z.Fmt { I: devhs.Has_stylish_haskell,	C: "stylish-haskell",	A: []string{} },
		z.Fmt { I: devhs.Has_hindent,			C: "hindent",			A: []string{"--no-force-newline", "--indent-size", ts} },
		z.Fmt { I: devhs.Has_brittany,			C: "brittany",			A: []string{"--indent", ts} },
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
