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


func Init (ctx *z.Context) z.Zengine {
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
		caps = []*z.CmdInfo	{	&z.CmdInfo { Title: "stylish-haskell",	Exists: devhs.Has_stylish_haskell,	Hint: "`stack install stylish-haskell`" },
								&z.CmdInfo { Title: "hindent",			Exists: devhs.Has_hindent,			Hint: "`stack install hindent`" },
								&z.CmdInfo { Title: "brittany",			Exists: devhs.Has_brittany,			Hint: "`github.com/lspitzner/brittany`" },
							}
	case "lint":
		caps = []*z.CmdInfo	{	&z.CmdInfo { Title: "hlint",			Exists: devhs.Has_hlint,			Hint: "`stack install hlint`" },
							}
	}
	return caps
}

func (self *zhs) DoFmt (src string, custcmd string, tabsize int) (resp *z.RespFmt, err error) {
	ts := fmt.Sprint(tabsize)
	return self.Base.DoFmt(src, custcmd,
		z.CmdInfo { Exists: devhs.Has_stylish_haskell,	Name: "stylish-haskell",	Args: []string{} },
		z.CmdInfo { Exists: devhs.Has_hindent,			Name: "hindent",			Args: []string{"--no-force-newline", "--indent-size", ts} },
		z.CmdInfo { Exists: devhs.Has_brittany,			Name: "brittany",			Args: []string{"--indent", ts} },
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
