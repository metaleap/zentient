package z
import (
	"path/filepath"
	"runtime"
	"strings"
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
	f.FullPath = normalizeFilePath(filepath.Join(Ctx.SrcDir, relpath))  ;  f.DirFull = filepath.Dir(f.FullPath)
	return &f
}

func normalizeFilePath (fpath string) string {
	if runtime.GOOS=="windows" { return strings.ToLower(filepath.FromSlash(fpath)) }
	return fpath
}
