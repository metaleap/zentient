package zhs
import (
    "github.com/metaleap/zentient/z"
)

type zhs struct {}

var Zengine z.Zengine = zhs{}


func (_ zhs) Ids () []string {
    return []string { "haskell", "Haskell" }
}

func (_ zhs) OnFileActive (file *z.File) {
}

func (_ zhs) OnFileOpen (file *z.File) {
}

func (_ zhs) OnFileWrite (file *z.File) {
}
