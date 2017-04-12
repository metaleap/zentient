package zgo
import (
	"path/filepath"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"

	"github.com/metaleap/zentient/z"
)


func linter (dirrelpath string, diagcat string, diagsev uint8, each func() []udev.SrcMsg) func()map[string][]*z.RespDiag {
	return func () map[string][]*z.RespDiag {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range each() {  if fpath := srcref.Ref  ;  filepath.Dir(fpath)==dirrelpath {
			d := &z.RespDiag { Sev: diagsev, SrcMsg: srcref }  ;  d.Ref = diagcat
			filediags[fpath] = append(filediags[fpath], d)
		} }
		return filediags
	}
}


func linterCheck (dirrelpath string, cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintCheck(cmdname, pkgimppath) })
}
func linterMvDan (dirrelpath string, cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintMvDan(cmdname, pkgimppath) })
}
func linterIneffAssign (dirrelpath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, "ineffassign", z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintIneffAssign(dirrelpath) })
}
func linterMDempsky (dirrelpath string, cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintMDempsky(cmdname, pkgimppath) })
}
func linterGolint (dirrelpath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, "golint", z.DIAG_SEV_HINT,
		func () []udev.SrcMsg { return devgo.LintGolint(dirrelpath) })
}
func linterGoVet (dirrelpath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, "go vet", z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintGoVet(dirrelpath) })
}
func linterErrcheck (dirrelpath string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, "errcheck", z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintErrcheck(pkgimppath) })
}
func linterHonnef (dirrelpath string, cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () []udev.SrcMsg { return devgo.LintHonnef(cmdname, pkgimppath) })
}


func (self *zgo) Linters (filerelpaths []string) (linters []func()map[string][]*z.RespDiag) {
	pkgfiles := map[*devgo.Pkg][]string {}  ;  for _,frp := range filerelpaths {
		if pkg := filePkg(frp) ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	for fpkg,frps := range pkgfiles { dirrelpath := filepath.Dir(frps[0])
		linters = append(linters, linterGoVet(dirrelpath))
		if devgo.Has_golint { linters = append(linters, linterGolint(dirrelpath)) }
		if devgo.Has_ineffassign { linters = append(linters, linterIneffAssign(dirrelpath)) }
		if devgo.Has_interfacer { linters = append(linters, linterMvDan(dirrelpath, "interfacer", fpkg.ImportPath)) }
		if devgo.Has_unparam { linters = append(linters, linterMvDan(dirrelpath, "unparam", fpkg.ImportPath)) }
		if devgo.Has_checkalign { linters = append(linters, linterCheck(dirrelpath, "aligncheck", fpkg.ImportPath)) }
		if devgo.Has_checkstruct { linters = append(linters, linterCheck(dirrelpath, "structcheck", fpkg.ImportPath)) }
		if devgo.Has_checkvar { linters = append(linters, linterCheck(dirrelpath, "varcheck", fpkg.ImportPath)) }
		if devgo.Has_unconvert { linters = append(linters, linterMDempsky(dirrelpath, "unconvert", fpkg.ImportPath)) }
		if devgo.Has_maligned { linters = append(linters, linterMDempsky(dirrelpath, "maligned", fpkg.ImportPath)) }
		if devgo.Has_gosimple { linters = append(linters, linterHonnef(dirrelpath, "gosimple", fpkg.ImportPath)) }
		if devgo.Has_unused { linters = append(linters, linterHonnef(dirrelpath, "unused", fpkg.ImportPath)) }
		if devgo.Has_staticcheck { linters = append(linters, linterHonnef(dirrelpath, "staticcheck", fpkg.ImportPath)) }
		if devgo.Has_errcheck { linters = append(linters, linterErrcheck(dirrelpath, fpkg.ImportPath)) }
	}
	return
}
