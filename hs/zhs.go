package zhs
import (
	"fmt"

	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devhs"
)


type zhs struct {
	z.Base
}


func Init () z.Zengine {
	if !devhs.HasHsDevEnv() { return nil }

	self := &zhs{}
	self.Base.Init()
	return self
}




func (_ *zhs) EdLangIDs () []string {
	return []string { "haskell", "Haskell" }
}

func (self *zhs) B () *z.Base {
	return &self.Base
}



func (_ *zhs) Caps (cap string) (caps []*z.RespCmd) {
	switch cap {
	case "fmt":
		caps = []*z.RespCmd	{	&z.RespCmd { Title: "stylish-haskell",	Exists: devhs.Has_stylish_haskell,	Hint: "`stack install stylish-haskell`" },
								&z.RespCmd { Title: "hindent",			Exists: devhs.Has_hindent,			Hint: "`stack install hindent`" },
								&z.RespCmd { Title: "brittany",			Exists: devhs.Has_brittany,			Hint: "`github.com/lspitzner/brittany`" },
							}
	case "diag":
		caps = []*z.RespCmd	{	&z.RespCmd { Title: "stack install",	Exists: devhs.HasHsDevEnv(),	Hint: "check your Stack installation" },
								&z.RespCmd { Title: "hlint",			Exists: devhs.Has_hlint,		Hint: "`stack install hlint`" },
							}
	}
	return caps
}

func (self *zhs) DoFmt (src string, custcmd string, tabsize uint8) (resp *z.RespFmt, err error) {
	ts := fmt.Sprint(tabsize)
	return self.Base.DoFmt(src, custcmd,
		z.RespCmd { Exists: devhs.Has_stylish_haskell,	Name: "stylish-haskell",	Args: []string{} },
		z.RespCmd { Exists: devhs.Has_hindent,			Name: "hindent",			Args: []string{"--no-force-newline", "--indent-size", ts} },
		z.RespCmd { Exists: devhs.Has_brittany,			Name: "brittany",			Args: []string{"--indent", ts} },
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
