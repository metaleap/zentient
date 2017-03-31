package zgo
import (
	"encoding/json"
	"go/build"

	"github.com/metaleap/zentient/z"
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)


type zgo struct {
	z.Base
	DbgJsonOut string
	PkgsByDir map[string]*Pkg
	PkgsByImP map[string]*Pkg
}

type Pkg struct {
	build.Package

	//	the below all copied over from `go list`  src because it outputs this stuff but our above build.Package doesn't have it:

	Target		string			`json:",omitempty"` // install path
	Shlib		string			`json:",omitempty"` // the shared library that contains this package (only set when -linkshared)
	Standard	bool			`json:",omitempty"` // is this package part of the standard Go library?
	Stale		bool			`json:",omitempty"` // would 'go install' do anything for this package?
	StaleReason	string			`json:",omitempty"` // why is Stale true?
	Incomplete	bool			`json:",omitempty"` // was there an error loading this package or dependencies?
	Error		*PkgErr			`json:",omitempty"` // error loading this package (not dependencies)
	DepsErrors	[]*PkgErr		`json:",omitempty"` // errors loading dependencies
}

type PkgErr struct {
	ImportStack	[]string	// shortest path from package named on command line to this one
	Pos			string		// position of error (if present, file:line:col)
	Err			string		// the error itself
}


var (
	µ *zgo
)


func New (root *z.RootInfo) z.Zengine {
	if !devgo.HasGoDevEnv() {
		return nil
	}
	µ = &zgo{}
	µ.Base.Init()
	go µ.refreshPkgs()
	return µ
}


func (self *zgo) refreshPkgs () {
	pkgsbydir,pkgsbyimp := map[string]*Pkg {} , map[string]*Pkg {}

	if cmdout,_ := ugo.CmdExec("go", "list", "-e", "-json", "all") ; len(cmdout)>0 {
		if jsonobjstrs := ustr.Split(cmdout, "}\n{") ; len(jsonobjstrs)>0 {
			for _, jsonobjstr := range jsonobjstrs {
				if jsonobjstr[0]!='{'					{ jsonobjstr = "{" + jsonobjstr }
				if jsonobjstr[len(jsonobjstr)-1]!='}'	{ jsonobjstr = jsonobjstr + "}" }
				var pkg Pkg // re-decl in loop isn't the most efficient but the most defensive =)
				if err := json.Unmarshal([]byte(jsonobjstr), &pkg) ; err==nil {
					pkgsbydir[pkg.Dir] = &pkg
					pkgsbyimp[pkg.ImportPath] = &pkg
				} else {
					if err.Error()=="unexpected end of JSON input" {
						self.DbgJsonOut = jsonobjstr
					}
					// self.DbgJsonOut += (err.Error() + "\n\n")
				}
			}
			if len(pkgsbydir)==0 { pkgsbydir = nil ; pkgsbyimp = nil }
			self.PkgsByDir,self.PkgsByImP = pkgsbydir,pkgsbyimp
		}
	}
}


func (_ *zgo) Ids () []string {
	return []string { "go", "Go" }
}


func (_ *zgo) Caps (cap string) (caps []*z.CmdInfo) {
	switch cap {
	case "fmt":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "gofmt", I: devgo.Has_gofmt, H: "check your Go installation" },
							}
	case "lint":
		caps = []*z.CmdInfo	{	&z.CmdInfo { N: "go vet", I: devgo.HasGoDevEnv(), H: "check your Go installation" },
								&z.CmdInfo { N: "golint", I: devgo.Has_golint, H: "`go get -u github.com/golang/lint/golint`" },
							}
	}
	return caps
}

func (self *zgo) DoFmt (src string, custcmd string, tabsize int) (*z.RespFmt, error) {
	return self.Base.DoFmt(src, custcmd, z.CmdInfo { I: devgo.Has_gofmt, C: "gofmt", A: []string{"-e", "-s"} })
}

func (_ *zgo) OnFileActive (file *z.File) {
}

func (_ *zgo) OnFileOpen (file *z.File) {
}

func (_ *zgo) OnFileClose (file *z.File) {
}

func (_ *zgo) OnFileWrite (file *z.File) {
}
