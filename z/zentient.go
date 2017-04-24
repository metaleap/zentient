package z
import (
	"os"
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
)


type Context struct {
	SrcDir		string
	CacheDir	string
	ConfigDir	string
}

var Ctx = &Context{}


func Init (zengineIniters map[string]func()Zengine) (err error) {
	Ctx.SrcDir,err = os.Getwd()
	udev.SrcDir = Ctx.SrcDir
	if err==nil { err = ensureDataDirs() }
	Zengines = map[string]Zengine {}
	inits := []func() {}
	for zid,zinit := range zengineIniters {
		_zid,_zinit := zid,zinit
		inits = append(inits, func() { regZ(_zid, _zinit()) })
	}
	ugo.WaitOn(inits...)
	return
}

func ensureDataDirs () error {
	var basedir, subdir string

	//  coming from VScode?
	if len(os.Args) > 1 && len(os.Args[1])>0 {
		const sep = string(os.PathSeparator)
		if editordatadir , index := os.Args[1] , strings.Index(os.Args[1], sep + "Code" + sep) ; index > 0 {
			basedir = editordatadir[0 : index]
		}
	}
	//  otherwise..
	if len(basedir) == 0 || !ufs.DirExists(basedir) {
		basedir = ugo.UserDataDirPath()
	}

	Ctx.ConfigDir = filepath.Join(basedir, "zentient")
	if volname := filepath.VolumeName(Ctx.SrcDir) ; len(volname) > 0 {
		subdir = strings.Replace(Ctx.SrcDir, volname, ufs.SanitizeFsName(volname), -1)
	} else {
		subdir = Ctx.SrcDir
	}
	Ctx.CacheDir = filepath.Join(Ctx.ConfigDir, subdir)
	return ufs.EnsureDirExists(Ctx.CacheDir) //  this also creates ConfigDir
}

func regZ (zid string, µ Zengine) {
	if µ != nil  {
		Zengines[zid] = µ
	}
}
