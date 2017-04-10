package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"

	"github.com/metaleap/zentient/z"
)


func linter (diagcat string, diagsev uint8, each func() []udev.SrcMsg) func(func(map[string][]*z.RespDiag)) {
	return func(cont func(map[string][]*z.RespDiag)) {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range each() {  fpath := srcref.Ref
			d := &z.RespDiag { Sev: diagsev, SrcMsg: srcref }  ;  d.Ref = diagcat
			filediags[fpath] = append(filediags[fpath], d)
		}
		cont(filediags)
	}
}


func linterCheck (cmdname string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	return linter(cmdname, z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintCheck(cmdname, pkgimppath)
	})
}
func linterMvDan (cmdname string, pkgimppath string) func(func(map[string][]*z.RespDiag)) {
	return linter(cmdname, z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintMvDan(cmdname, pkgimppath)
	})
}
func linterIneffAssign (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("ineffassign", z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintIneffAssign(filerelpaths)
	})
}
func linterGoLint (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("golint", z.DIAG_SEV_HINT, func () []udev.SrcMsg {
		return devgo.LintGolint(filerelpaths)
	})
}
func linterGoVet (filerelpaths []string) func(func(map[string][]*z.RespDiag)) {
	return linter("go vet", z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintGoVet(filerelpaths)
	})
}


func (self *zgo) Lint (filerelpaths []string, ondelayedlintersdone func(map[string][]*z.RespDiag)) map[string][]*z.RespDiag {
	latefuncs := []func(func(map[string][]*z.RespDiag)) {}
	pkgfiles := map[*devgo.Pkg][]string {}  ;  for _,frp := range filerelpaths {
		if pkg := filePkg(frp) ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}

	for fpkg,frps := range pkgfiles {
		latefuncs = append(latefuncs, linterGoVet(frps))
		if devgo.Has_interfacer		{ latefuncs = append(latefuncs, linterMvDan("interfacer", fpkg.ImportPath)) }
		if devgo.Has_unparam		{ latefuncs = append(latefuncs, linterMvDan("unparam", fpkg.ImportPath)) }
		if devgo.Has_checkalign		{ latefuncs = append(latefuncs, linterCheck("aligncheck", fpkg.ImportPath)) }
		if devgo.Has_checkstruct	{ latefuncs = append(latefuncs, linterCheck("structcheck", fpkg.ImportPath)) }
		if devgo.Has_checkvar		{ latefuncs = append(latefuncs, linterCheck("varcheck", fpkg.ImportPath)) }
		if devgo.Has_ineffassign	{ latefuncs = append(latefuncs, linterIneffAssign(frps)) }
		if devgo.Has_golint			{ latefuncs = append(latefuncs, linterGoLint(frps)) }
	}
	return nil // self.Base.Lint(latefuncs, ondelayedlintersdone)
}
