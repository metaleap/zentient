package z
import (
    "path/filepath"
)


type File struct {
    RelPath string
    FullPath string
    DataPath string
    Z Zengine
}


func NewFile (z Zengine, relpath string) *File {
    var me File
    me.Z = z
    me.RelPath = relpath
    me.FullPath = filepath.Join(ProjDir, relpath)
    me.DataPath = filepath.Join(DataProjDir, relpath)
    return &me
}
