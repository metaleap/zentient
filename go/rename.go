package zgo
import (
	"fmt"
	"strings"

	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)



func (self *zgo) DoRename (reqcmd string, relfilepath string, offset uint64, newname string, eol string, oldname string, off1 uint64, off2 uint64) (resp map[string][]*z.RespRen, err error) {
	cmdargs := []string{ "-d", "-to", newname, "-offset", fmt.Sprintf("%s:#%d", relfilepath, offset) }
	var renout string  ;  if len(reqcmd)==0 {  reqcmd = "gorename"  }
	if renout,err = ugo.CmdExec(reqcmd, cmdargs...)  ;  err!=nil { if len(renout)>0 { err = ugo.E(renout) }  ;  return }
	i := ustr.Idx(renout, "--- ")  ;  if i<0 { err = ugo.E(renout)  ;  return }
	resp = map[string][]*z.RespRen {}  ;  renout = renout[i+4:]  ;  rendiffs := uslice.StrMap(ustr.Split(renout, "--- "), strings.TrimSpace)
	if len(rendiffs)==0 { return nil , ugo.E("Renaming aborted: no diffs could be obtained.") }

	for _,rendiff := range rendiffs {
		if i = ustr.Idx(rendiff, "\t")  ;  i<=0 { return nil , ugo.E("Renaming aborted: could not detect file path in diffs.") }
		if ffp := rendiff[:i]  ;  !ufs.FileExists(ffp) { return nil , ugo.E("Renaming aborted: bad absolute file path `" + ffp + "` in diffs.") } else {
			if i = ustr.Idx(rendiff, "@@ -")  ;  i<=0 { return nil , ugo.E("Renaming aborted: `@@ -` expected.") } else {
				feds := []*z.RespRen {}  ;  for _,hunkchunk := range ustr.Split(rendiff[i+4:], "@@ -") { if lns := ustr.Split(hunkchunk, "\n")  ;  len(lns)>0 {
					i = ustr.Idx(lns[0], ",")  ;  lb := ustr.ParseInt(lns[0][:i])  ;  s := lns[0][i+1:]  ;  ll := ustr.ParseInt(s[:ustr.Idx(s, " +")])
					if lb==0 || ll==0 { return nil , ugo.E("Renaming aborted: diffs contained invalid or unparsable line hints.") } else {
						fed := &z.RespRen{ StartLn: lb-1, StartChr: 0, EndLn: lb-1+ll, EndChr: 0 }
						for _,ln := range lns[1:] { if ustr.Pref(ln, " ") || ustr.Pref(ln, "+") { fed.NewText = fed.NewText + ln[1:] + eol } }
						feds = append(feds, fed)
					}
				} else { return nil , ugo.E("Renaming aborted: expected something between one `@@ -` and the next.") } }
				if len(feds)>0 {  resp[ffp] = feds  } else {  return nil , ugo.E("Renaming aborted: a diff without effective edits.")  }
			}
		}
	}
	return
}
