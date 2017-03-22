package zgo
import (
    "github.com/metaleap/zentient/z"
)


type zgo struct {
}

var (
    zengine *zgo
)


func New (proj *z.ProjInfo) z.Zengine {
    zengine = &zgo{}
    return zengine
}



func (_ zgo) Ids () []string {
    return []string { "go", "Go" }
}

func (me* zgo) Jsonish () interface{} {
    return me
}

func (_ zgo) OnFileActive (file *z.File) {
}

func (_ zgo) OnFileOpen (file *z.File) {
}

func (_ zgo) OnFileWrite (file *z.File) {
}
