package z
import (
	"path/filepath"
)


type File struct {
	RelPath		string
	FullPath	string
	DirRel		string
	DirFull		string
	µ			Zengine
	Proj		interface{}
}


func newFile (z Zengine, relpath string) *File {
	var f File
	f.µ = z
	f.RelPath = relpath  ;  f.DirRel = filepath.Dir(f.RelPath)
	f.FullPath = filepath.Join(Ctx.SrcDir, relpath)  ;  f.DirFull = filepath.Dir(f.FullPath)
	return &f
}
