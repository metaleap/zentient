package zhs

import (
	"fmt"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient/z"
)

type zhs struct {
	z.Base
}

var (
	srcDir string
)

func Init() z.Zengine {
	if !udevhs.HasHsDevEnv() {
		return nil
	}
	srcDir = z.Ctx.SrcDir

	me := &zhs{}
	me.Base.Init()
	return me
}

func (_ *zhs) EdLangIDs() []string {
	return []string{"haskell", "Haskell"}
}

func (me *zhs) B() *z.Base {
	return &me.Base
}

func (_ *zhs) Caps(cap string) (caps []*z.RespCmd) {
	switch cap {
	case "fmt":
		caps = []*z.RespCmd{{Title: "stylish-haskell", Exists: udevhs.Has_stylish_haskell, Hint: "hackage.haskell.org/package/stylish-haskell"},
			{Title: "hindent", Exists: udevhs.Has_hindent, Hint: "hackage.haskell.org/package/hindent"},
			{Title: "brittany", Exists: udevhs.Has_brittany, Hint: "github.com/lspitzner/brittany"},
		}
	case "diag":
		caps = []*z.RespCmd{{Title: "hlint", Exists: udevhs.Has_hlint, Hint: "hackage.haskell.org/package/hlint"}}
	}
	return caps
}

func (me *zhs) DoFmt(src string, custcmd string, tabsize uint8) (resp *z.RespTxt, err error) {
	ts := fmt.Sprint(tabsize)
	return me.Base.DoFmt(src, custcmd,
		z.RespCmd{Exists: udevhs.Has_stylish_haskell, Name: "stylish-haskell", Args: []string{}},
		z.RespCmd{Exists: udevhs.Has_hindent, Name: "hindent", Args: []string{"--no-force-newline", "--indent-size", ts}},
		z.RespCmd{Exists: udevhs.Has_brittany, Name: "brittany", Args: []string{"--indent", ts}},
	)
}

func (me *zhs) DoRename(reqcmd string, relfilepath string, offset uint64, newname string, eol string, oldname string, off1 uint64, off2 uint64) (resp map[string]udev.SrcMsgs, err error) {
	err = umisc.E("Renaming symbol `" + oldname + "` in " + relfilepath + " at :" + umisc.Str(offset) + " (" + umisc.Str(off1) + " - " + umisc.Str(off2) + ") to `" + newname + "` rejected")
	return
}

func (me *zhs) OnCfg(cfg map[string]interface{}) {
	me.Base.OnCfg(cfg)
}

func (_ *zhs) OnFile(newfile *z.File) {
}

func (_ *zhs) QueryTools() []*z.RespPick {
	return nil
}

func (_ *zhs) QueryTool(req *z.ReqIntel) *z.RespTxt {
	return nil
}

func (_ *zhs) ReadyToBuildAndLint() bool {
	return true
}
