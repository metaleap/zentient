package zgo

import (
	"path/filepath"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"

	"github.com/metaleap/zentient/z"
)

func linter(dirrelpath string, diagcat string, diagsev int, each func() udev.SrcMsgs) func() map[string]udev.SrcMsgs {
	return func() map[string]udev.SrcMsgs {
		filediags := map[string]udev.SrcMsgs{}
		for _, srcref := range each() {
			if fpath := srcref.Ref; filepath.Dir(fpath) == dirrelpath {
				srcref.Ref = diagcat
				srcref.Flag = diagsev
				filediags[fpath] = append(filediags[fpath], srcref)
			}
		}
		return filediags
	}
}

func linterCheck(dirrelpath string, cmdname string, pkgimppath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintCheck(cmdname, pkgimppath) })
}
func linterMvDan(dirrelpath string, cmdname string, pkgimppath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintMvDan(cmdname, pkgimppath) })
}
func linterIneffAssign(dirrelpath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, "ineffassign", z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintIneffAssign(dirrelpath) })
}
func linterMDempsky(dirrelpath string, cmdname string, pkgimppath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintMDempsky(cmdname, pkgimppath) })
}
func linterGolint(dirrelpath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, "golint", z.DIAG_SEV_HINT,
		func() udev.SrcMsgs { return udevgo.LintGolint(dirrelpath) })
}
func linterGoVet(dirrelpath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, "go vet", z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintGoVet(dirrelpath) })
}
func linterErrcheck(dirrelpath string, pkgimppath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, "errcheck", z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintErrcheck(pkgimppath) })
}
func linterHonnef(dirrelpath string, cmdname string, pkgimppath string) func() map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func() udev.SrcMsgs { return udevgo.LintHonnef(cmdname, pkgimppath) })
}

func (me *zgo) Linters(filerelpaths []string, forcelinters ...string) (linters []func() map[string]udev.SrcMsgs) {
	pkgfiles := map[*udevgo.Pkg][]string{}
	for _, frp := range filerelpaths {
		if pkg := filePkg(frp); pkg != nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	cfgok := me.Base.CfgDiagToolEnabled(forcelinters)
	for fpkg, frps := range pkgfiles {
		dirrelpath := filepath.Dir(frps[0])
		if cfgok("go vet") {
			linters = append(linters, linterGoVet(dirrelpath))
		}
		if udevgo.Has_golint && cfgok("golint") {
			linters = append(linters, linterGolint(dirrelpath))
		}
		if udevgo.Has_ineffassign && cfgok("ineffassign") {
			linters = append(linters, linterIneffAssign(dirrelpath))
		}
		if udevgo.Has_interfacer && cfgok("interfacer") {
			linters = append(linters, linterMvDan(dirrelpath, "interfacer", fpkg.ImportPath))
		}
		if udevgo.Has_unparam && cfgok("unparam") {
			linters = append(linters, linterMvDan(dirrelpath, "unparam", fpkg.ImportPath))
		}
		if udevgo.Has_checkalign && cfgok("aligncheck") {
			linters = append(linters, linterCheck(dirrelpath, "aligncheck", fpkg.ImportPath))
		}
		if udevgo.Has_checkstruct && cfgok("structcheck") {
			linters = append(linters, linterCheck(dirrelpath, "structcheck", fpkg.ImportPath))
		}
		if udevgo.Has_checkvar && cfgok("varcheck") {
			linters = append(linters, linterCheck(dirrelpath, "varcheck", fpkg.ImportPath))
		}
		if udevgo.Has_unconvert && cfgok("unconvert") {
			linters = append(linters, linterMDempsky(dirrelpath, "unconvert", fpkg.ImportPath))
		}
		if udevgo.Has_maligned && cfgok("maligned") {
			linters = append(linters, linterMDempsky(dirrelpath, "maligned", fpkg.ImportPath))
		}
		if udevgo.Has_gosimple && cfgok("gosimple") {
			linters = append(linters, linterHonnef(dirrelpath, "gosimple", fpkg.ImportPath))
		}
		if udevgo.Has_unused && cfgok("unused") {
			linters = append(linters, linterHonnef(dirrelpath, "unused", fpkg.ImportPath))
		}
		if udevgo.Has_staticcheck && cfgok("staticcheck") {
			linters = append(linters, linterHonnef(dirrelpath, "staticcheck", fpkg.ImportPath))
		}
		if udevgo.Has_errcheck && cfgok("errcheck") {
			linters = append(linters, linterErrcheck(dirrelpath, fpkg.ImportPath))
		}
	}
	return
}
