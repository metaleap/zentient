package z
import (
    "path/filepath"
)


type File struct {
    RelPath     string
    FullPath    string
    Z           Zengine
}


func NewFile (z Zengine, relpath string) *File {
    var f File
    f.Z = z
    f.RelPath = relpath
    f.FullPath = filepath.Join(Root.SrcDir, relpath)
    return &f
}
