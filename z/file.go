package z
import (
	"path/filepath"
)


type File struct {
	RelPath		string
	FullPath	string
	µ			Zengine
}


func NewFile (z Zengine, relpath string) *File {
	var f File
	f.µ = z
	f.RelPath = relpath
	f.FullPath = filepath.Join(Ctx.SrcDir, relpath)
	return &f
}
