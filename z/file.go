package z
import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/metaleap/go-util-slice"
)

type File struct {
	RelPath		string
	FullPath	string
	DirRel		string
	DirFull		string
	µ			Zengine
	Proj		interface{}
}

var (
	AllFiles	= map[string]*File {}
	OpenFiles	= []string {}
)

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

func onFilesClosed (µ Zengine, relpaths []string) {
	for i,_ := range relpaths { relpaths[i] = normalizeFilePath(relpaths[i]) }
	OpenFiles = uslice.StrWithout(OpenFiles, false, relpaths...)
}

func onFilesOpened (µ Zengine, relpaths []string) {
	for _,relpath := range relpaths {
		relpath = normalizeFilePath(relpath)
		file := AllFiles[relpath]  ;  if file == nil {
			file = newFile(µ, relpath)  ;  µ.OnFile(file)  ;  AllFiles[relpath] = file
		}
		if isopened := !uslice.StrHas(OpenFiles, relpath) ; isopened {  OpenFiles = append(OpenFiles, relpath)  }
	}
}

func onFilesWritten (µ Zengine, relpaths []string) {
	for i,_ := range relpaths { relpaths[i] = normalizeFilePath(relpaths[i]) }
	µ.B().buildFrom(µ, relpaths)
}
