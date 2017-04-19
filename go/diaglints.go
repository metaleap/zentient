package zgo
import (
	"path/filepath"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-slice"

	"github.com/metaleap/zentient/z"
)


func linter (dirrelpath string, diagcat string, diagsev int, each func() udev.SrcMsgs) func()map[string]udev.SrcMsgs {
	return func () map[string]udev.SrcMsgs {
		filediags := map[string]udev.SrcMsgs {}
		for _,srcref := range each() {  if fpath := srcref.Ref  ;  filepath.Dir(fpath)==dirrelpath {
			srcref.Ref = diagcat  ;  srcref.Flag = diagsev
			filediags[fpath] = append(filediags[fpath], srcref)
		} }
		return filediags
	}
}


func linterCheck (dirrelpath string, cmdname string, pkgimppath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintCheck(cmdname, pkgimppath) })
}
func linterMvDan (dirrelpath string, cmdname string, pkgimppath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintMvDan(cmdname, pkgimppath) })
}
func linterIneffAssign (dirrelpath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, "ineffassign", z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintIneffAssign(dirrelpath) })
}
func linterMDempsky (dirrelpath string, cmdname string, pkgimppath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintMDempsky(cmdname, pkgimppath) })
}
func linterGolint (dirrelpath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, "golint", z.DIAG_SEV_HINT,
		func () udev.SrcMsgs { return devgo.LintGolint(dirrelpath) })
}
func linterGoVet (dirrelpath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, "go vet", z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintGoVet(dirrelpath) })
}
func linterErrcheck (dirrelpath string, pkgimppath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, "errcheck", z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintErrcheck(pkgimppath) })
}
func linterHonnef (dirrelpath string, cmdname string, pkgimppath string) func()map[string]udev.SrcMsgs {
	return linter(dirrelpath, cmdname, z.DIAG_SEV_INFO,
		func () udev.SrcMsgs { return devgo.LintHonnef(cmdname, pkgimppath) })
}


func (me *zgo) Linters (filerelpaths []string, forcelinters ...string) (linters []func()map[string]udev.SrcMsgs) {
	pkgfiles := map[*devgo.Pkg][]string {}  ;  for _,frp := range filerelpaths {
		if pkg := filePkg(frp) ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	cfgok := me.Base.CfgDiagToolEnabled
	if len(forcelinters)>0 { cfgok = func (cmdname string) bool { return uslice.StrHas(forcelinters, cmdname) } }
	for fpkg,frps := range pkgfiles { dirrelpath := filepath.Dir(frps[0])
		if cfgok("go vet")									{ linters = append(linters, linterGoVet(dirrelpath)) }
		if devgo.Has_golint && cfgok("golint")				{ linters = append(linters, linterGolint(dirrelpath)) }
		if devgo.Has_ineffassign && cfgok("ineffassign")	{ linters = append(linters, linterIneffAssign(dirrelpath)) }
		if devgo.Has_interfacer && cfgok("interfacer")		{ linters = append(linters, linterMvDan(dirrelpath, "interfacer", fpkg.ImportPath)) }
		if devgo.Has_unparam && cfgok("unparam")			{ linters = append(linters, linterMvDan(dirrelpath, "unparam", fpkg.ImportPath)) }
		if devgo.Has_checkalign && cfgok("aligncheck")		{ linters = append(linters, linterCheck(dirrelpath, "aligncheck", fpkg.ImportPath)) }
		if devgo.Has_checkstruct && cfgok("structcheck")	{ linters = append(linters, linterCheck(dirrelpath, "structcheck", fpkg.ImportPath)) }
		if devgo.Has_checkvar && cfgok("varcheck")			{ linters = append(linters, linterCheck(dirrelpath, "varcheck", fpkg.ImportPath)) }
		if devgo.Has_unconvert && cfgok("unconvert")		{ linters = append(linters, linterMDempsky(dirrelpath, "unconvert", fpkg.ImportPath)) }
		if devgo.Has_maligned && cfgok("maligned")			{ linters = append(linters, linterMDempsky(dirrelpath, "maligned", fpkg.ImportPath)) }
		if devgo.Has_gosimple && cfgok("gosimple")			{ linters = append(linters, linterHonnef(dirrelpath, "gosimple", fpkg.ImportPath)) }
		if devgo.Has_unused && cfgok("unused")				{ linters = append(linters, linterHonnef(dirrelpath, "unused", fpkg.ImportPath)) }
		if devgo.Has_staticcheck && cfgok("staticcheck")	{ linters = append(linters, linterHonnef(dirrelpath, "staticcheck", fpkg.ImportPath)) }
		if devgo.Has_errcheck && cfgok("errcheck")			{ linters = append(linters, linterErrcheck(dirrelpath, fpkg.ImportPath)) }
	}
	return
}
