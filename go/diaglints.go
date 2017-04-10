package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"

	"github.com/metaleap/zentient/z"
)


func linter (diagcat string, diagsev uint8, each func() []udev.SrcMsg) func()map[string][]*z.RespDiag {
	return func () map[string][]*z.RespDiag {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range each() {  fpath := srcref.Ref
			d := &z.RespDiag { Sev: diagsev, SrcMsg: srcref }  ;  d.Ref = diagcat
			filediags[fpath] = append(filediags[fpath], d)
		}
		return filediags
	}
}


func linterCheck (cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(cmdname, z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintCheck(cmdname, pkgimppath)
	})
}
func linterMvDan (cmdname string, pkgimppath string) func()map[string][]*z.RespDiag {
	return linter(cmdname, z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintMvDan(cmdname, pkgimppath)
	})
}
func linterIneffAssign (filerelpaths []string) func()map[string][]*z.RespDiag {
	return linter("ineffassign", z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintIneffAssign(filerelpaths)
	})
}
func linterGoLint (filerelpaths []string) func()map[string][]*z.RespDiag {
	return linter("golint", z.DIAG_SEV_HINT, func () []udev.SrcMsg {
		return devgo.LintGolint(filerelpaths)
	})
}
func linterGoVet (filerelpaths []string) func()map[string][]*z.RespDiag {
	return linter("go vet", z.DIAG_SEV_INFO, func () []udev.SrcMsg {
		return devgo.LintGoVet(filerelpaths)
	})
}


func (self *zgo) Linters (filerelpaths []string) (linters []func()map[string][]*z.RespDiag) {
	pkgfiles := map[*devgo.Pkg][]string {}  ;  for _,frp := range filerelpaths {
		if pkg := filePkg(frp) ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	for fpkg,frps := range pkgfiles {
		linters = append(linters, linterGoVet(frps))
		if devgo.Has_golint { linters = append(linters, linterGoLint(frps)) }
		if devgo.Has_ineffassign { linters = append(linters, linterIneffAssign(frps)) }
		if devgo.Has_interfacer { linters = append(linters, linterMvDan("interfacer", fpkg.ImportPath)) }
		if devgo.Has_unparam { linters = append(linters, linterMvDan("unparam", fpkg.ImportPath)) }
		if devgo.Has_checkalign { linters = append(linters, linterCheck("aligncheck", fpkg.ImportPath)) }
		if devgo.Has_checkstruct { linters = append(linters, linterCheck("structcheck", fpkg.ImportPath)) }
		if devgo.Has_checkvar { linters = append(linters, linterCheck("varcheck", fpkg.ImportPath)) }
	}
	return
}
