package zgo
import (
	"runtime"
	"strings"

	"github.com/metaleap/go-devgo"

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
	self := &zgo{}
	self.Base.Init()
	return self
}


func (_ *zgo) EdLangIDs () []string {
	return []string { "go", "Go" }
}

func (self *zgo) B () *z.Base {
	return &self.Base
}


func (_ *zgo) Caps (cap string) (caps []*z.RespCmd) {
	switch cap {
	case "fmt":
		caps = []*z.RespCmd	{	{ Title: "goimports",	Exists: devgo.Has_goimports,	Hint: "`go get golang.org/x/tools/cmd/goimports`" },
								{ Title: "gofmt",		Exists: devgo.Has_gofmt,		Hint: "fix your Go installation" },
							}
	case "diag":
		caps = []*z.RespCmd	{	{ Title: "go install",	Exists: true },
								{ Title: "go vet",		Exists: true },
								{ Title: "golint",		Exists: devgo.Has_golint,		Hint: "`go get -u github.com/golang/lint/golint`" },
								{ Title: "ineffassign",	Exists: devgo.Has_ineffassign,	Hint: "`go get -u github.com/gordonklaus/ineffassign`" },
								{ Title: "aligncheck",	Exists: devgo.Has_checkalign,	Hint: "`github.com/opennota/check`" },
								{ Title: "structcheck",	Exists: devgo.Has_checkstruct,	Hint: "`github.com/opennota/check`" },
								{ Title: "varcheck",	Exists: devgo.Has_checkvar,		Hint: "`github.com/opennota/check`" },
								{ Title: "interfacer",	Exists: devgo.Has_golint,		Hint: "`github.com/mvdan/interfacer`" },
								{ Title: "unconvert",	Exists: devgo.Has_unconvert,	Hint: "`github.com/mdempsky/unconvert`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, custcmd string, tabsize uint8) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, custcmd,
		z.RespCmd { Exists: devgo.Has_goimports, Name: "goimports", Args: []string { "-e" } },
		z.RespCmd { Exists: devgo.Has_gofmt, Name: "gofmt", Args: []string {"-e", "-s"} })
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
