package zhs
import (
	"github.com/metaleap/zentient/z"
)


type zhs struct {
	z.ZengineBase
}

var (
	µ *zhs
)


func New (root *z.RootInfo) z.Zengine {
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



func (_ *zhs) Caps (string) []string {
	return []string {}
}

func (_ *zhs) DoFmt (src string) *z.FmtResp {
	return nil
}

func (_ *zhs) OnFileActive (file *z.File) {
}

func (_ *zhs) OnFileOpen (file *z.File) {
}

func (_ *zhs) OnFileClose (file *z.File) {
}

func (_ *zhs) OnFileWrite (file *z.File) {
}
