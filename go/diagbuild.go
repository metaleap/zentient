package zgo
import (
	"path/filepath"
	"sync"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"

	"github.com/metaleap/zentient/z"
)


var (
	laterebuilds sync.Mutex
)


func buildPkg (pkgimppath string, fallbackfilerelpath string, diags map[string]udev.SrcMsgs) bool {
	msgs := udev.CmdExecOnSrc(true, nil, "go", "install", pkgimppath)
	for _,srcref := range msgs { if srcref.Msg != "too many errors" {
		fpath := srcref.Ref  ;  srcref.Ref = "go install " + pkgimppath  ;  srcref.Flag = z.DIAG_SEV_ERR
		if !ufs.FileExists(filepath.Join(srcDir, fpath)) { srcref.Msg = fpath + ": " + srcref.Msg  ;  fpath = fallbackfilerelpath }
		diags[fpath] = append(diags[fpath], srcref)
	} }
	return len(msgs)==0
}


func (_ *zgo) BuildFrom (filerelpaths []string) (freshdiags map[string]udev.SrcMsgs) {
	pkgimppaths := []string {}  ;  pkgimpimppaths := []string {}  ;  devgo.GuruScopeExclPkgs = nil

	for _,frp := range filerelpaths { if pkg := filePkg(frp)  ;  pkg!=nil {
		if !(uslice.StrHas(pkgimppaths, pkg.ImportPath) || uslice.StrHas(pkgimpimppaths, pkg.ImportPath)) {
			pkgimppaths = append(pkgimppaths, pkg.ImportPath)
		}
		for _,pip := range pkg.Importers("") {
			if !(uslice.StrHas(pkgimppaths, pip) || uslice.StrHas(pkgimpimppaths, pip)) {
				pkgimpimppaths = append(pkgimpimppaths, pip)
			}
		}
	} }
	freshdiags = map[string]udev.SrcMsgs {}  ;  succeeded := []string {}
	for _,pkgimppath := range pkgimppaths {
		if success := buildPkg(pkgimppath, filerelpaths[0], freshdiags)  ;  success {
			succeeded = append(succeeded, pkgimppath)
		} else { return }
	}
	for _,pkgimppath := range pkgimpimppaths {
		if success := buildPkg(pkgimppath, filerelpaths[0], freshdiags)  ;  success {
			succeeded = append(succeeded, pkgimppath)
		}
	}
	refreshindirectdependants := func() {
		if asyncrebuilds := devgo.AllFinalDependants(succeeded)  ;  len(asyncrebuilds)>0 {
			defer refreshPkgs()  ;  laterebuilds.Lock()  ;  defer laterebuilds.Unlock()
			for _,pkgimppath := range asyncrebuilds {
				if !(uslice.StrHas(pkgimppaths, pkgimppath) || uslice.StrHas(pkgimpimppaths, pkgimppath) || uslice.StrHas(succeeded, pkgimppath)) {
					ugo.CmdExec("go", "install", pkgimppath)
				}
			}
		}
	}
	go refreshindirectdependants()
	return
}
