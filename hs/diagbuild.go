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


func (_ *zhs) BuildFrom (filerelpaths []string) (freshdiags map[string][]*udev.SrcMsg) {
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
		freshdiags = map[string][]*udev.SrcMsg {}

		for _,wln := range lnstackwarns { if parts := ustr.Split(wln, ":")  ;  len(parts)>1 {
			freshdiags[parts[0]] = append(freshdiags[parts[0]], &udev.SrcMsg { Flag: z.DIAG_SEV_WARN, Ref: "stack", Msg: ustr.Join(parts[1:], ":"), Pos1Ln: 1, Pos1Ch: 1 })
		}}

		if _p := "Could not parse '" + srcDir  ;  ustr.Pref(lns[0], _p) {
			stackyamlpath := strings.TrimRight(strings.TrimLeft(lns[0][len(_p):], "/\\"), "':")
			freshdiags[stackyamlpath] = append(freshdiags[stackyamlpath], &udev.SrcMsg { Ref: "stack", Msg: ustr.Join(lns[1:], "\n"), Pos1Ln: 1, Pos1Ch: 1 })
			return
		}

		if _e := "Error: "  ;  ustr.Pref(lns[0], _e) {
			freshdiags[filerelpath] = append(freshdiags[filerelpath], &udev.SrcMsg { Ref: "stack", Msg: ustr.Join(lns, "\n")[len(_e):], Pos1Ln: 1, Pos1Ch: 1 })
			return
		}

		var cur *udev.SrcMsg  ;  addlastcur := func() { if cur!=nil {
			fpath,_ := filepath.Rel(srcDir, cur.Ref)
			if len(fpath)==0 {  fpath = cur.Ref  }
			if ustr.Suff(fpath, ".hs") {  cur.Misc = "ghc"  }
			cur.Ref = "stack"  ;  cur.Msg = strings.TrimSpace(cur.Msg)
			if i := strings.Index(cur.Msg, "]")  ;  ustr.Pref(cur.Msg, "[") && i>0 {
				cur.Ref = cur.Ref + "  » " + cur.Msg[1:i]  ;  cur.Msg = strings.TrimSpace(cur.Msg[i+1:])
			}
			freshdiags[fpath] = append(freshdiags[fpath], cur)
		} }
		for _,ln := range lns { if len(ln)> 0 {
			if msg := udev.SrcMsgFromLn(ln)  ;  msg!=nil && ustr.Pref(msg.Ref, srcDir) && ufs.FileExists(msg.Ref) {
				addlastcur()
				cur = msg  ;  if ustr.Pref(cur.Msg, "error: ") { cur.Msg = cur.Msg[7:] }
				if ustr.Pref(cur.Msg, "warning: ") { cur.Flag = z.DIAG_SEV_WARN  ;  cur.Msg = cur.Msg[9:] }
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
