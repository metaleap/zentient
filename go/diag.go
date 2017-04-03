package zgo
import (
	"path/filepath"
	"sync"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"

	"github.com/metaleap/zentient/z"
)


func newGoLint (pkgimppath string, filenames []string, cont func(map[string][]*z.RespDiag)) func() {
	return func() {
		filediags := map[string][]*z.RespDiag {}
		for _,srcref := range devgo.CmdExecOnSrc(true, false, nil, "golint", filenames...) {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: "golint", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
		}
		cont(filediags)
	}
}

	var govetcheckln = ustr.NotPrefixed("vet: ") // no need to recreate this lambda every time
func newGoVet (pkgimppath string, filenames []string, cont func(map[string][]*z.RespDiag)) func() {
	return func() {
		filediags := map[string][]*z.RespDiag {}
		cmdargs := []string { "tool", "vet", "-shadow=true", "-shadowstrict", "-all" }
		cmdargs = append(cmdargs, filenames...)
		for _,srcref := range devgo.CmdExecOnSrc(true, true, govetcheckln, "go", cmdargs...) {
			filediags[srcref.FilePath] = append(filediags[srcref.FilePath],
				&z.RespDiag { Cat: "go vet", Msg: srcref.Msg, PosLn: srcref.PosLn, PosCol: srcref.PosCol, Sev: z.DIAG_INFO })
		}
		cont(filediags)
	}
}


func (self *zgo) Lint (filerelpaths []string) (alldiags map[string][]*z.RespDiag) {
	alldiags = map[string][]*z.RespDiag {}
	if devgo.PkgsByDir==nil { return }

	pkgfiles := map[*devgo.Pkg][]string {}
	for _,frp := range filerelpaths {
		if pkg := devgo.PkgsByDir[filepath.Dir(filepath.Join(z.Ctx.SrcDir, frp))] ; pkg!=nil {
			pkgfiles[pkg] = append(pkgfiles[pkg], frp)
		}
	}
	var mutex sync.Mutex
	onlinted := func(linterdiags map[string][]*z.RespDiag) {
		mutex.Lock()  ;  defer mutex.Unlock()
		for frp,filediags := range linterdiags { alldiags[frp] = append(alldiags[frp], filediags...) }
	}
	funcs := []func() {}
	for fpkg,frps := range pkgfiles {
		if devgo.HasGoDevEnv() { funcs = append(funcs, newGoVet(fpkg.ImportPath, frps, onlinted)) }
		if devgo.Has_golint { funcs = append(funcs, newGoLint(fpkg.ImportPath, frps, onlinted)) }
	}
	ugo.WaitOn(funcs...)

	return
}

func (self *zgo) BuildFrom (filerelpath string) (diags []*z.RespDiag) {
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "rebuildfile:" + filerelpath, PosLn: 9, PosCol: 2, Sev: z.DIAG_ERR })
	diags = append(diags, &z.RespDiag { Cat: "devgo-mock", Msg: "filerebuild:" + filerelpath, PosLn: 18, PosCol: 4, Sev: z.DIAG_WARN })
	return
}


// func (self *zgo) refreshPkgDiags (rebuildfilerelpath string) {
// 	errs := devgo.RefreshPkgs()
// 	self.Base.DbgObjs = append(self.Base.DbgObjs, devgo.PkgsByDir)
// 	for _,err := range errs { self.Base.DbgMsgs = append(self.Base.DbgMsgs, err.Error()) }
// 	pd := map[string][]*z.RespDiag {}
// 	for _,pkg := range devgo.PkgsErrs {
// 		for _,pkgerr := range pkg.Errs {
// 			if len(pkgerr.RelPath)>0 && pkgerr.RelPath!=rebuildfilerelpath {
// 				pd[pkgerr.RelPath] = append(pd[pkgerr.RelPath],
// 					&z.RespDiag { Cat: "go list all", Msg: pkgerr.Msg, PosLn: pkgerr.PosLn, PosCol: pkgerr.PosCol, Sev: z.DIAG_ERR })
// 			}
// 		}
// 	}
// 	pkgdiags = pd
// }
