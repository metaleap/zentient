package zgo
import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"

	"github.com/metaleap/zentient/z"
)


func lnrelify (ln string) string {
	if ufs.PathPrefix(ln, srcDir) {
		if ln = ln[len(srcDir):] ; ln[0]=='\\' || ln[0]=='/' { ln = ln[1:] }
		return ln
	}
	return ""
}


func linter (diagcat string, diagsev uint8, each func() []udev.SrcRefMsg) func(func(map[string][]*z.RespDiag)) {
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range each() {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: diagcat, Sev: diagsev, Msg: srcref.Msg, Data: srcref.Data, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Pos2Ln: srcref.Pos2Ln, Pos2Col: srcref.Pos2Col })
		}
		cont(filediags)
	}
}


func linterCheck (cmdname string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	return linter(cmdname, z.DIAG_INFO, func () []udev.SrcRefMsg {
		return devgo.LintCheck(cmdname, pkgimppath)
	})
}
func linterMvDan (cmdname string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	return linter(cmdname, z.DIAG_INFO, func () []udev.SrcRefMsg {
		return devgo.LintMvDan(cmdname, pkgimppath)
	})
}
func linterIneffAssign (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("ineffassign", z.DIAG_INFO, func () []udev.SrcRefMsg {
		return devgo.LintIneffAssign(filerelpaths)
	})
}
func linterGoLint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("golint", z.DIAG_HINT, func () []udev.SrcRefMsg {
		return devgo.LintGolint(filerelpaths)
	})
}
func linterGoVet (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("go vet", z.DIAG_WARN, func () []udev.SrcRefMsg {
		return devgo.LintGoVet(filerelpaths)
	})
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
		if devgo.Has_checkalign		{ latefuncs = append(latefuncs, linterCheck("aligncheck", fpkg.ImportPath)) }
		if devgo.Has_checkstruct	{ latefuncs = append(latefuncs, linterCheck("structcheck", fpkg.ImportPath)) }
		if devgo.Has_checkvar		{ latefuncs = append(latefuncs, linterCheck("varcheck", fpkg.ImportPath)) }
		if devgo.Has_ineffassign	{ funcs = append(funcs, linterIneffAssign(frps)) }
		if devgo.Has_golint			{ funcs = append(funcs, linterGoLint(frps)) }
	}
	return self.Base.Lint(funcs, latefuncs, ondelayedlintersdone)
}

func (_ *zgo) LintReady () bool {
	return devgo.PkgsByDir!=nil
}
