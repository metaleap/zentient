package zgo
import (
	"github.com/metaleap/zentient/z"

	"github.com/metaleap/go-devgo"
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




func (_ zgo) Ids () []string {
	return []string { "go", "Go" }
}

func (self* zgo) Jsonish () interface{} {
	return self
}


func (self* zgo) Base () *z.ZengineBase {
	return &self.ZengineBase
}


func (_ zgo) Caps (string) []string {
	caps := []string {}
	if devgo.GoFmt { caps = append(caps, "gofmt") }
	return caps
}

func (_ zgo) OnFileActive (file *z.File) {
}

func (_ zgo) OnFileOpen (file *z.File) {
}

func (_ zgo) OnFileClose (file *z.File) {
}

func (_ zgo) OnFileWrite (file *z.File) {
}
