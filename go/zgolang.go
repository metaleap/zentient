package zgo
import (
    "github.com/metaleap/zentient/z"
)

type zgo struct {}

var Zengine z.Zengine = zgo{}


func (_ zgo) Ids () []string {
    return []string { "go", "Go" }
}
