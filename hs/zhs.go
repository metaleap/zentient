package zhs
import (
	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devhs"
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

	caps = append(caps, &z.RespCap { Name: "hindent", Available: devhs.Has_hindent, InstHint: "`stack install hindent`" })

	return caps
}

func (_ *zhs) DoFmt (src string) (resp *z.RespFmt, err error) {
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
