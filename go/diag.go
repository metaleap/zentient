package zgo
import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"

	"github.com/metaleap/zentient/z"
)

var (
	dirrelpathpref = "." + string(os.PathSeparator)
	laterebuilds sync.Mutex
)


func lnrelify (ln string) string {
	if ufs.PathPrefix(ln, srcDir) {
		if ln = ln[len(srcDir):] ; ln[0]=='\\' || ln[0]=='/' { ln = ln[1:] }
		return ln
	}
	return ""
}


func linterCheck (chk string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	reline := func (pkgimppath string) func(string)string {
		return func (ln string) string {
			if strings.HasPrefix(ln, pkgimppath + ": ") { return lnrelify(ln[len(pkgimppath)+2:]) }
			return ""
		}
	}
	return func(cont func(map[string][]*z.RespDiag)) {
		cmdname := chk + "check"
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range udev.CmdExecOnSrc(true, true, true, reline(pkgimppath), cmdname, pkgimppath) {
			if strings.HasPrefix(srcref.Msg, pkgimppath+".") { srcref.Msg = srcref.Msg[len(pkgimppath)+1:] }
			if chk!="align" { srcref.Msg = "(unused & unexported) " + srcref.Msg }
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: cmdname, Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
		}
		cont(filediags)
	}
}

func linterMvDan (cmdname string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	reline := func (ln string) string {
		if rln := lnrelify(ln) ; len(rln)>0 { return rln }
		return ln
	}
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range udev.CmdExecOnSrc(true, true, false, reline, cmdname, pkgimppath) {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: cmdname, Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
		}
		cont(filediags)
	}
}

func linterIneffAssign (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	reline := lnrelify
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,filerelpath := range filerelpaths {
			for _,srcref := range udev.CmdExecOnSrc(true, true, false, reline, "ineffassign", filerelpath) {
				filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
					&z.RespDiag { Cat: "ineffassign", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
			}
		}
		cont(filediags)
	}
}

func linterGoLint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	censored := func (msg string) (skip bool) {
		skip = skip || ustr.Has(msg, " should have comment ")
		skip = skip || msg == "if block ends with a return statement, so drop this else and outdent its block"
		skip = skip || ustr.Has(msg, "ALL_CAPS")
		skip = skip || ustr.Has(msg, "underscore")
		skip = skip || ustr.Has(msg, "CamelCase")
		skip = skip || ustr.Has(msg, "package comment should be of the form \"")
		skip = skip || ustr.Has(msg, "should omit 2nd value from range; this loop is equivalent to ")
		skip = skip || ustr.Has(msg, "don't use generic names")
		skip = skip || ustr.DistBetween(msg, "comment on exported ", " should be of the form \"") > 0
		return
	}
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range udev.CmdExecOnSrc(true, true, false, nil, "golint", filerelpaths...) {
			if !censored(srcref.Msg) {
				filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
					&z.RespDiag { Cat: "golint", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_HINT })
			}
		}
		cont(filediags)
	}
}

func linterGoVet (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	reline := func (ln string) string {  if strings.HasPrefix(ln, "vet: ") { return "" } else { return ln }  }
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		cmdargs := []string { "tool", "vet", "-shadow=true", "-shadowstrict", "-all" }
		cmdargs = append(cmdargs, filerelpaths...)
		for _,srcref := range udev.CmdExecOnSrc(true, false, true, reline, "go", cmdargs...) {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: "go vet", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_WARN })
		}
		cont(filediags)
	}
}



func (self *zgo) Lint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) map[string][]*z.RespDiag {
	funcs := []func(func(map[string][]*z.RespDiag)) {}  ;  latefuncs := []func(func(map[string][]*z.RespDiag)) {}
	pkgfiles := map[*devgo.Pkg][]string {}
	for _,frp := range filerelpaths {
		if pkg := devgo.PkgsByDir[strings.ToLower(filepath.Dir(filepath.Join(srcDir, frp)))] ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}

	for fpkg,frps := range pkgfiles {
		funcs = append(funcs, linterGoVet(frps))
		if devgo.Has_interfacer		{ latefuncs = append(latefuncs, linterMvDan("interfacer", fpkg.ImportPath)) }
		if devgo.Has_unparam		{ latefuncs = append(latefuncs, linterMvDan("unparam", fpkg.ImportPath)) }
		if devgo.Has_checkalign		{ latefuncs = append(latefuncs, linterCheck("align", fpkg.ImportPath)) }
		if devgo.Has_checkstruct	{ latefuncs = append(latefuncs, linterCheck("struct", fpkg.ImportPath)) }
		if devgo.Has_checkvar		{ latefuncs = append(latefuncs, linterCheck("var", fpkg.ImportPath)) }
		if devgo.Has_ineffassign	{ funcs = append(funcs, linterIneffAssign(frps)) }
		if devgo.Has_golint			{ funcs = append(funcs, linterGoLint(frps)) }
	}
	return self.Base.Lint(funcs, latefuncs, ondelayedlintersdone)
}

func (_ *zgo) LintReady () bool {
	return devgo.PkgsByDir!=nil
}


func (_ *zgo) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	dirrelpath := filepath.Dir(filerelpath)  ;  freshdiags = map[string][]*z.RespDiag {}
	dirrelpaths := devgo.ImportersOf(filepath.Join(srcDir, dirrelpath), srcDir)
	dirrelpathsmin := append([]string { dirrelpath }, devgo.ShakeOutIntermediateDepsViaDir(dirrelpaths, srcDir)...)

	succeeded := []string {}  ;  for i,dirrelpath := range dirrelpathsmin {
		msgs := udev.CmdExecOnSrc(true, false, true, nil, "go", "install", dirrelpathpref + dirrelpath) // filepath.Join NOT good here: would remove ./ that `go install` does need to use dirrelpath instead of an imp-path
		for _,srcref := range msgs { freshdiags[srcref.FilePath] = append(freshdiags[srcref.FilePath],
			&z.RespDiag { Cat: "gc", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_ERR }) }
		if success := len(msgs)==0  ;  success {
			succeeded = append(succeeded, dirrelpath)
		} else if i==0 { return }
	}

	rebuildindirectdependantsasync := func() {
		asyncrebuilds := []string {}
		for _,dirrelpath := range dirrelpaths { if !uslice.StrHas(dirrelpathsmin, dirrelpath) { asyncrebuilds = append(asyncrebuilds, dirrelpath) } }
		asyncrebuilds = uslice.StrMap(append(asyncrebuilds, succeeded...), func(drp string) string { return filepath.Join(srcDir, drp) })
		if len(asyncrebuilds)>0 {
			defer devgo.RefreshPkgs()  ;  laterebuilds.Lock()  ;  defer laterebuilds.Unlock()
			for _,pkgimppath := range devgo.AllFinalDependants(asyncrebuilds) { ugo.CmdExec("go", "install", pkgimppath) }
		}
	}
	go rebuildindirectdependantsasync()
	return
}
