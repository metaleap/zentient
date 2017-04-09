package zhs
import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-devhs"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)


func (_ *zhs) BuildFrom (filerelpaths []string) (freshdiags map[string][]*z.RespDiag) {
	filerelpath := filerelpaths[0]
	filefullpath := filepath.Join(srcDir, filerelpath)  ;  dirfullpath := filepath.Dir(filefullpath)
	cmdargs := append(append([]string { "build" }, devhs.StackArgs...), devhs.StackArgsBuild...)
	raw,_ := ugo.CmdExecIn(dirfullpath, "stack", cmdargs...)  ;  lns := ustr.Split(raw, "\n")

	lnstackwarns := []string {}  ;  _w := "Warning: " + srcDir  ;  _wl := len(_w)
	for i := 0 ; i < len(lns) ; i++ {  if ustr.Pref(lns[i], _w) && strings.Index(lns[i], "stack.yaml: ") >= _wl {
		lnstackwarns = append(lnstackwarns, strings.TrimLeft(lns[i][_wl:], "/\\"))  ;  lns = append(lns[:i], lns[i+1:]...)  ;  i--
	} }

	for len(lns)>0 && len(lns[0])==0 { lns = lns[1:] }  ;  wasnoop := len(lns)==0
	if len(lns)>3 && ustr.Pref(lns[0], "Copying from ") && len(lns[1])==0 && ustr.Pref(lns[2], "Copied executables to ") {
		wasnoop = true  ;  for i := 3 ; i < len(lns) ; i++ { if !ustr.Pref(lns[i], "- ") { wasnoop = false } }
	}
	if !wasnoop {
		freshdiags = map[string][]*z.RespDiag {}

		for _,wln := range lnstackwarns { if parts := ustr.Split(wln, ":")  ;  len(parts)>1 {
			freshdiags[parts[0]] = append(freshdiags[parts[0]], &z.RespDiag { Sev: z.DIAG_SEV_WARN, SrcMsg: udev.SrcMsg { Ref: "stack", Msg: ustr.Join(parts[1:], ":"), PosLn: 1, PosCol: 1 } })
		}}

		if _p := "Could not parse '" + srcDir  ;  ustr.Pref(lns[0], _p) {
			stackyamlpath := strings.TrimRight(strings.TrimLeft(lns[0][len(_p):], "/\\"), "':")
			freshdiags[stackyamlpath] = append(freshdiags[stackyamlpath], &z.RespDiag { SrcMsg: udev.SrcMsg { Ref: "stack", Msg: ustr.Join(lns[1:], "\n"), PosLn: 1, PosCol: 1 } })
			return
		}

		if _e := "Error: "  ;  ustr.Pref(lns[0], _e) {
			freshdiags[filerelpath] = append(freshdiags[filerelpath], &z.RespDiag { SrcMsg: udev.SrcMsg { Ref: "stack", Msg: ustr.Join(lns, "\n")[len(_e):], PosLn: 1, PosCol: 1 } })
			return
		}

		var cur *z.RespDiag  ;  addlastcur := func() { if cur!=nil {
			fpath,_ := filepath.Rel(srcDir, cur.Ref)
			if len(fpath)==0 {  fpath = cur.Ref  }
			cur.Ref = "ghc"  ;  cur.Msg = strings.TrimSpace(cur.Msg)
			freshdiags[fpath] = append(freshdiags[fpath], cur)
		} }
		for _,ln := range lns { if len(ln)> 0 {
			if msg,isghcmsg := udev.SrcMsgFromLn(ln)  ;  isghcmsg && ustr.Pref(msg.Ref, srcDir) && ufs.FileExists(msg.Ref) {
				addlastcur()
				cur = &z.RespDiag {  SrcMsg: msg  }  ;  if _e := "error:"  ;  ustr.Pref(cur.Msg, _e) { cur.Msg = cur.Msg[len(_e):] }
				if _w := "warning: "  ;  ustr.Pref(cur.Msg, _w) { cur.Sev = z.DIAG_SEV_WARN  ;  cur.Msg = cur.Msg[len(_w):] }
			} else if ustr.Pref(ln, "    ") && cur!=nil {
				cur.Msg += ("\n" + ln)
			} else {
				addlastcur()  ;  cur = nil
			}
		} }
		addlastcur()
	}
	return
}
