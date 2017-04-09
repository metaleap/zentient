package zgo
import (
	"os"
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
	relpathprefix = "." + string(os.PathSeparator)
	laterebuilds sync.Mutex
)


func (_ *zgo) BuildFrom (filerelpath string) (freshdiags map[string][]*z.RespDiag) {
	dirrelpath := filepath.Dir(filerelpath)  ;  freshdiags = map[string][]*z.RespDiag {}
	dirrelpaths := devgo.ImportersOf(filepath.Join(srcDir, dirrelpath), srcDir)
	dirrelpathsmin := append([]string { dirrelpath }, devgo.ShakeOutIntermediateDepsViaDir(dirrelpaths, srcDir)...)

	succeeded := []string {}  ;  for _,dirrelpath := range dirrelpathsmin {
		msgs := udev.CmdExecOnSrc(true, nil, "go", "install", relpathprefix + dirrelpath) // filepath.Join NOT good here: would remove ./ that `go install` does need to use dirrelpath instead of an imp-path
		for _,srcref := range msgs { if srcref.Msg != "too many errors" {
			d := &z.RespDiag { Sev: z.DIAG_SEV_ERR, SrcMsg: srcref }  ;  d.Ref = dirrelpath
			fpath := srcref.Ref  ;  if !ufs.FileExists(filepath.Join(srcDir, fpath)) {
				fpath = filerelpath  ;  d.Msg = srcref.Ref + ": " + d.Msg
			}
			freshdiags[fpath] = append(freshdiags[fpath], d) } }
		if success := len(msgs)==0  ;  success {
			succeeded = append(succeeded, dirrelpath)
		} else { return }
	}

	rebuildindirectdependantsasync := func() {
		asyncrebuilds := []string {}
		for _,dirrelpath := range dirrelpaths { if !uslice.StrHas(dirrelpathsmin, dirrelpath) { asyncrebuilds = append(asyncrebuilds, dirrelpath) } }
		asyncrebuilds = uslice.StrMap(append(asyncrebuilds, succeeded...), func(drp string) string { return filepath.Join(srcDir, drp) })
		if len(asyncrebuilds)>0 {
			defer devgo.RefreshPkgs()  ;  laterebuilds.Lock()  ;  defer laterebuilds.Unlock()
			for _,pkgimppath := range devgo.AllFinalDependants(asyncrebuilds) {
				ugo.CmdExec("go", "install", pkgimppath)
			}
		}
	}
	go rebuildindirectdependantsasync()
	return
}
