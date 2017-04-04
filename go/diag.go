package zgo
import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"

	"github.com/metaleap/zentient/z"
)


func lnrelify (ln string) string {
	if ufs.PathPrefix(ln, z.Ctx.SrcDir) {
		if ln = ln[len(z.Ctx.SrcDir):] ; ln[0]=='\\' || ln[0]=='/' { ln = ln[1:] }
		return ln
	}
	return ""
}


func newGoCheck (chk string, pkgimppath string, cont func(map[string][]*z.RespDiag)) func() {
	reline := func (pkgimppath string) func(string)string {
		return func (ln string) string {
			if strings.HasPrefix(ln, pkgimppath + ": ") { return lnrelify(ln[len(pkgimppath)+2:]) }
			return ""
		}
	}
	return func() {
		cmdname := chk + "check"
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devgo.CmdExecOnSrc(true, true, true, reline(pkgimppath), cmdname, pkgimppath) {
			if strings.HasPrefix(srcref.Msg, pkgimppath+".") { srcref.Msg = srcref.Msg[len(pkgimppath)+1:] }
			if chk!="align" { srcref.Msg = "(unused & unexported) " + srcref.Msg }
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: cmdname, Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
		}
		cont(filediags)
	}
}

func newIneffAssign (filerelpaths []string, cont func(map[string][]*z.RespDiag)) func() {
	reline := lnrelify
	return func() {
		filediags := map[string][]*z.RespDiag {}
		for _,filerelpath := range filerelpaths {
			for _,srcref := range devgo.CmdExecOnSrc(true, true, false, reline, "ineffassign", filerelpath) {
				filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
					&z.RespDiag { Cat: "ineffassign", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
			}
		}
		cont(filediags)
	}
}

func newGoLint (filerelpaths []string, cont func(map[string][]*z.RespDiag)) func() {
	censored := func (msg string) (skip bool) {
		skip = skip || msg == "if block ends with a return statement, so drop this else and outdent its block"
		skip = skip || ustr.Has(msg, "ALL_CAPS")
		skip = skip || ustr.Has(msg, "underscore")
		skip = skip || ustr.Has(msg, "CamelCase")
		skip = skip || ustr.Has(msg, " should have comment ")
		skip = skip || ustr.Has(msg, "package comment should be of the form \"")
		skip = skip || ustr.Has(msg, "should omit 2nd value from range; this loop is equivalent to ")
		skip = skip || ustr.Has(msg, "don't use generic names")
		skip = skip || ustr.DistBetween(msg, "comment on exported ", " should be of the form \"") > 0
		return
	}
	return func() {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devgo.CmdExecOnSrc(true, true, false, nil, "golint", filerelpaths...) {
			if !censored(srcref.Msg) {
				filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
					&z.RespDiag { Cat: "golint", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_HINT })
			}
		}
		cont(filediags)
	}
}

func newGoVet (filerelpaths []string, cont func(map[string][]*z.RespDiag)) func() {
	reline := func (ln string) string {  if strings.HasPrefix(ln, "vet: ") { return "" } else { return ln }  }
	return func() {
		filediags := map[string][]*z.RespDiag {}
		cmdargs := []string { "tool", "vet", "-shadow=true", "-shadowstrict", "-all" }
		cmdargs = append(cmdargs, filerelpaths...)
		for _,srcref := range devgo.CmdExecOnSrc(true, false, true, reline, "go", cmdargs...) {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: "go vet", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_WARN })
		}
		cont(filediags)
	}
}


func (self *zgo) Lint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) (freshdiags map[string][]*z.RespDiag) {
	if devgo.PkgsByDir==nil { return }
	freshdiags = map[string][]*z.RespDiag {}  ;  latediags := map[string][]*z.RespDiag {}

	pkgfiles := map[*devgo.Pkg][]string {}
	for _,frp := range filerelpaths {
		if pkg := devgo.PkgsByDir[filepath.Dir(filepath.Join(z.Ctx.SrcDir, frp))] ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	var mutex sync.Mutex  ;  var mutexlate sync.Mutex
	onlints := func(linterdiags map[string][]*z.RespDiag) {
		mutex.Lock()  ;  defer mutex.Unlock()
		for frp,filediags := range linterdiags { freshdiags[frp] = append(freshdiags[frp], filediags...) }
	}
	onlintslate := func(linterdiags map[string][]*z.RespDiag) {
		mutexlate.Lock()  ;  defer mutexlate.Unlock()
		for frp,filediags := range linterdiags { latediags[frp] = append(latediags[frp], filediags...) }
	}
	funcs := []func() {}  ;  latefuncs := []func() {}
	runlatefuncs := func () { ugo.WaitOn(latefuncs...)  ;  ondelayedlintersdone(latediags) }
	for fpkg,frps := range pkgfiles {
		if devgo.Has_checkalign		{ latefuncs = append(latefuncs, newGoCheck("align", fpkg.ImportPath, onlintslate)) }
		if devgo.Has_checkstruct	{ latefuncs = append(latefuncs, newGoCheck("struct", fpkg.ImportPath, onlintslate)) }
		if devgo.Has_checkvar		{ latefuncs = append(latefuncs, newGoCheck("var", fpkg.ImportPath, onlintslate)) }
		if devgo.Has_ineffassign	{ funcs = append(funcs, newIneffAssign(frps, onlints)) }
		if devgo.HasGoDevEnv()		{ funcs = append(funcs, newGoVet(frps, onlints)) }
		if devgo.Has_golint			{ funcs = append(funcs, newGoLint(frps, onlints)) }
	}
	go runlatefuncs()
	ugo.WaitOn(funcs...)
	return
}

func (self *zgo) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	freshdiags = map[string][]*z.RespDiag {}
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devgo-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { Cat: "devgo-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}
