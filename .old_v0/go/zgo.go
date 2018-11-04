package zgo

import (
	"runtime"
	"strings"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient/z"
)

type zgo struct {
	z.Base
}

var (
	srcDir string
)

func Init() z.Zengine {
	if !udevgo.HasGoDevEnv() {
		return nil
	}
	srcDir = z.Ctx.SrcDir
	go refreshPkgs()
	me := &zgo{}
	me.Base.Init()
	return me
}

func (_ *zgo) EdLangIDs() []string {
	return []string{"go", "Go"}
}

func (me *zgo) B() *z.Base {
	return &me.Base
}

var (
	capsfmt, capsdiag, capsren, capsint, capsext []*z.RespCmd
	capsinit                                     = false
)

func (_ *zgo) Caps(cap string) (caps []*z.RespCmd) {
	if !capsinit {
		capsinit = true
		capsfmt = []*z.RespCmd{{Title: "goimports", Exists: udevgo.Has_goimports, Hint: "golang.org/x/tools/cmd/goimports"},
			{Title: "gofmt", Exists: udevgo.Has_gofmt},
		}
		capsdiag = []*z.RespCmd{{Title: "go vet", Exists: true},
			{Title: "golint", Exists: udevgo.Has_golint, Hint: "github.com/golang/lint"},
			{Title: "ineffassign", Exists: udevgo.Has_ineffassign, Hint: "github.com/gordonklaus/ineffassign"},
			{Title: "aligncheck", Exists: udevgo.Has_checkalign, Hint: "github.com/opennota/check"},
			{Title: "structcheck", Exists: udevgo.Has_checkstruct, Hint: "github.com/opennota/check"},
			{Title: "varcheck", Exists: udevgo.Has_checkvar, Hint: "github.com/opennota/check"},
			{Title: "errcheck", Exists: udevgo.Has_errcheck, Hint: "github.com/kisielk/errcheck"},
			{Title: "interfacer", Exists: udevgo.Has_golint, Hint: "github.com/mvdan/interfacer"},
			{Title: "unparam", Exists: udevgo.Has_unparam, Hint: "github.com/mvdan/unparam"},
			{Title: "unconvert", Exists: udevgo.Has_unconvert, Hint: "github.com/mdempsky/unconvert"},
			{Title: "maligned", Exists: udevgo.Has_maligned, Hint: "github.com/mdempsky/maligned"},
			{Title: "gosimple", Exists: udevgo.Has_gosimple, Hint: "github.com/dominikh/go-tools"},
			{Title: "unused", Exists: udevgo.Has_unused, Hint: "github.com/dominikh/go-tools"},
			{Title: "staticcheck", Exists: udevgo.Has_staticcheck, Hint: "github.com/dominikh/go-tools"},
		}
		capsren = []*z.RespCmd{{Title: "gorename", Exists: udevgo.Has_gorename, Hint: "golang.org/x/tools/cmd/gorename", More: "(affected files will be formatted gofmt-style)"}}
		capsext = []*z.RespCmd{{Title: "guru", Exists: udevgo.Has_guru, Hint: "golang.org/x/tools/cmd/guru", More: "via <i>Code Intel Extras</i>: Callees, Callers, Callstack, Free Variables, Types of Errors, Points-To, Channel Peers"},
			{Title: "go run", Exists: true, More: "via <i>Query Extras</i>: attempt to evaluate the given expression in the current package context"},
			{Title: "go doc", Exists: true, More: "via <i>Query Extras</i>"},
		}
		capsint = []*z.RespCmd{{Title: "gocode", Exists: udevgo.Has_gocode, Hint: "github.com/nsf/gocode", More: "Completion Suggest"},
			{Title: "guru", Exists: udevgo.Has_guru, Hint: "golang.org/x/tools/cmd/guru", More: "Go to Definition, Go to Type Definition, Go to Interfaces/Implementers, References Lookup, Symbols Lookup, Semantic Highlighting, Code Intel Extras"},
			{Title: "gogetdoc", Exists: udevgo.Has_gogetdoc, Hint: "github.com/zmb3/gogetdoc", More: "Hover Tips, Go to Definition, summaries for Completion Suggest"},
			{Title: "godef", Exists: udevgo.Has_godef, Hint: "github.com/rogpeppe/godef", More: "Go to Definition, Hover Tips"},
		}
	}
	switch cap {
	case "fmt":
		caps = capsfmt
	case "diag":
		caps = capsdiag
	case "ren":
		caps = capsren
	case "intel":
		caps = capsint
	case "extra":
		caps = capsext
	}
	return caps
}

func (me *zgo) DoFmt(src string, custcmd string, tabsize uint8) (*z.RespTxt, error) {
	return me.Base.DoFmt(src, custcmd,
		z.RespCmd{Exists: udevgo.Has_goimports, Name: "goimports", Args: []string{"-e"}},
		z.RespCmd{Exists: udevgo.Has_gofmt, Name: "gofmt", Args: []string{"-e", "-s"}})
}

func (me *zgo) DoRename(reqcmd string, relfilepath string, offset uint64, newname string, eol string, oldname string, off1 uint64, off2 uint64) (resp map[string]udev.SrcMsgs, err error) {
	if len(reqcmd) == 0 && !udevgo.Has_gorename {
		return nil, umisc.E("Couldn't find `gorename` command, and no custom tool was specified either.")
	}
	var fileedits udev.SrcMsgs
	if fileedits, err = udevgo.Gorename(reqcmd, relfilepath, offset, newname, eol); len(fileedits) > 0 {
		resp = map[string]udev.SrcMsgs{}
		for _, sr := range fileedits {
			ffp := sr.Ref
			sr.Ref = ""
			resp[ffp] = append(resp[ffp], sr)
		}
	}
	return
}

func (me *zgo) OnCfg(cfg map[string]interface{}) {
	me.Base.OnCfg(cfg)
}

func (_ *zgo) OnFile(newfile *z.File) {
	setFilePkgInfo(newfile)
}

func (_ *zgo) ReadyToBuildAndLint() bool {
	return udevgo.PkgsByDir != nil
}

func filePkg(relfilepath string) *udevgo.Pkg {
	if f := z.AllFiles[relfilepath]; f != nil {
		if pkg, ok := f.Proj.(*udevgo.Pkg); ok {
			return pkg
		}
	}
	return nil
}

func refreshPkgs() {
	udevgo.RefreshPkgs()
	for _, file := range z.AllFiles {
		setFilePkgInfo(file)
	}
}

func setFilePkgInfo(file *z.File) {
	dir := file.DirFull
	if runtime.GOOS == "windows" {
		dir = strings.ToLower(dir)
	}
	if pkg := udevgo.PkgsByDir[dir]; pkg != nil {
		file.Proj = pkg
	}
}
