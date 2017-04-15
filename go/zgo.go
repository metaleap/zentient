package zgo
import (
	"runtime"
	"strings"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/zentient/z"
)


type zgo struct {
	z.Base
}


var (
	srcDir string
)


func Init () z.Zengine {
	if !devgo.HasGoDevEnv() { return nil }
	srcDir = z.Ctx.SrcDir
	go refreshPkgs()
	me := &zgo{}
	me.Base.Init()
	return me
}


func (_ *zgo) EdLangIDs () []string {
	return []string { "go", "Go" }
}

func (me *zgo) B () *z.Base {
	return &me.Base
}


func (_ *zgo) Caps (cap string) (caps []*z.RespCmd) {
	switch cap {
	case "fmt":
		caps = []*z.RespCmd	{	{ Title: "goimports",	Exists: devgo.Has_goimports,	Hint: "golang.org/x/tools/cmd/goimports" },
								{ Title: "gofmt",		Exists: devgo.Has_gofmt },
							}
	case "diag":
		caps = []*z.RespCmd	{	{ Title: "go vet",		Exists: true },
								{ Title: "golint",		Exists: devgo.Has_golint,		Hint: "github.com/golang/lint" },
								{ Title: "ineffassign",	Exists: devgo.Has_ineffassign,	Hint: "github.com/gordonklaus/ineffassign" },
								{ Title: "aligncheck",	Exists: devgo.Has_checkalign,	Hint: "github.com/opennota/check" },
								{ Title: "structcheck",	Exists: devgo.Has_checkstruct,	Hint: "github.com/opennota/check" },
								{ Title: "varcheck",	Exists: devgo.Has_checkvar,		Hint: "github.com/opennota/check" },
								{ Title: "errcheck",	Exists: devgo.Has_errcheck,		Hint: "github.com/kisielk/errcheck" },
								{ Title: "interfacer",	Exists: devgo.Has_golint,		Hint: "github.com/mvdan/interfacer" },
								{ Title: "unparam",		Exists: devgo.Has_unparam,		Hint: "github.com/mvdan/unparam" },
								{ Title: "unconvert",	Exists: devgo.Has_unconvert,	Hint: "github.com/mdempsky/unconvert" },
								{ Title: "maligned",	Exists: devgo.Has_maligned,		Hint: "github.com/mdempsky/maligned" },
								{ Title: "gosimple",	Exists: devgo.Has_gosimple,		Hint: "github.com/dominikh/go-tools" },
								{ Title: "unused",		Exists: devgo.Has_unused,		Hint: "github.com/dominikh/go-tools" },
								{ Title: "staticcheck",	Exists: devgo.Has_staticcheck,	Hint: "github.com/dominikh/go-tools" },
							}
	case "ren":
		caps = []*z.RespCmd	{	{ Title: "gorename",	Exists: devgo.Has_gorename,		Hint: "golang.org/x/tools/cmd/gorename",	More: "(affected files will be formatted gofmt-style)" },
							}
	case "intel":
		caps = []*z.RespCmd	{	{ Title: "gocode",		Exists: devgo.Has_gocode,		Hint: "github.com/nsf/gocode",				More: "Completion Suggest" },
								{ Title: "gogetdoc",	Exists: devgo.Has_gogetdoc,		Hint: "github.com/zmb3/gogetdoc",			More: "Hover Tips, Go to Definition, summaries for Completion Suggest" },
								{ Title: "godef",		Exists: devgo.Has_godef,		Hint: "github.com/rogpeppe/godef",			More: "Go to Definition, Hover Tips" },
							}
	}
	return caps
}

func (me *zgo) DoFmt (src string, custcmd string, tabsize uint8) (*z.RespTxt, error) {
	return me.Base.DoFmt(src, custcmd,
		z.RespCmd { Exists: devgo.Has_goimports, Name: "goimports", Args: []string { "-e" } },
		z.RespCmd { Exists: devgo.Has_gofmt, Name: "gofmt", Args: []string {"-e", "-s"} })
}

func (me *zgo) DoRename (reqcmd string, relfilepath string, offset uint64, newname string, eol string, oldname string, off1 uint64, off2 uint64) (resp map[string][]*udev.SrcMsg, err error) {
	if len(reqcmd)==0 && !devgo.Has_gorename { return nil , ugo.E("Couldn't find `gorename` command, and no custom tool was specified either.") }
	var fileedits []*udev.SrcMsg  ;  if fileedits,err = devgo.Gorename(reqcmd, relfilepath, offset, newname, eol)  ;  len(fileedits)>0 {
		resp = map[string][]*udev.SrcMsg {}  ;  for _,sr := range fileedits {
			ffp := sr.Ref  ;  sr.Ref = ""  ;  resp[ffp] = append(resp[ffp], sr)
		}
	}
	return
}

func (me *zgo) OnCfg (cfg map[string]interface{}) {
	me.Base.OnCfg(cfg)
}

func (_ *zgo) OnFile (newfile *z.File) {
	setFilePkgInfo(newfile)
}

func (_ *zgo) ReadyToBuildAndLint () bool {
	return devgo.PkgsByDir!=nil
}


func filePkg (relfilepath string) *devgo.Pkg {
	if f := z.AllFiles[relfilepath]  ;  f != nil {
		if pkg,ok := f.Proj.(*devgo.Pkg)  ;  ok { return pkg }
	}
	return nil
}

func refreshPkgs () {
	devgo.RefreshPkgs()
	for _,file := range z.AllFiles { setFilePkgInfo(file) }
}

func setFilePkgInfo (file *z.File) {
	dir := file.DirFull  ;  if runtime.GOOS=="windows" { dir = strings.ToLower(dir) }
	if pkg := devgo.PkgsByDir[dir]  ;  pkg!=nil { file.Proj = pkg }
}
