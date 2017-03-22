package zhs
import (
    "github.com/metaleap/zentient/z"
)


type zhs struct {
}

var (
    zengine *zhs
)


func New (proj *z.ProjInfo) z.Zengine {
    zengine = &zhs{}
    return zengine
}



func (_ zhs) Ids () []string {
    return []string { "haskell", "Haskell" }
}

func (me* zhs) Jsonish () interface{} {
    return me
}

func (_ zhs) OnFileActive (file *z.File) {
}

func (_ zhs) OnFileOpen (file *z.File) {
}

func (_ zhs) OnFileWrite (file *z.File) {
}
