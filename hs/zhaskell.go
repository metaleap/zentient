package zhs
import (
    "github.com/metaleap/zentient/z"
)

type zhs struct {}

var Zengine z.Zengine = zhs{}


func (_ zhs) Ids () []string {
    return []string { "haskell", "Haskell" }
}
