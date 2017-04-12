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



func (self *zgo) DoRename (reqcmd string, relfilepath string, offset uint64, newname string, oldname string, off1 uint64, off2 uint64) (resp map[string][]*z.RespRen, err error) {
	// err = ugo.E("Renaming symbol `" + oldname + "` in " + relfilepath + " at :" + ugo.SPr(offset) + " (" + ugo.SPr(off1) + " - " + ugo.SPr(off2) + ") to `" + newname + "` rejected")
	cmdargs := []string{ "-d", "-to", newname, "-offset", fmt.Sprintf("%s:#%d", relfilepath, offset) }
	var renout string  ;  if len(reqcmd)==0 {  reqcmd = "gorename"  }
	if renout,err = ugo.CmdExec(reqcmd, cmdargs...)  ;  err!=nil && len(renout)>0 {  err = ugo.E(renout)  } else if err==nil {
		if idx := strings.Index(renout, "--- ")  ;  idx<0 { err = ugo.E(renout) } else {
			renout = renout[idx+4:]  ;  rendiffs := uslice.StrMap(ustr.Split(renout, "--- "), strings.TrimSpace)
			resp = map[string][]*z.RespRen {}
			for _,rendiff := range rendiffs { if difflns := ustr.Split(rendiff, "\n")  ;  len(difflns)>0 { if idx := strings.Index(difflns[0], "\t")  ;  idx>0 { if ffp := difflns[0][:idx]  ;  ufs.FileExists(ffp) {
				resp[ffp] = []*z.RespRen { { NewText: "fooboo\n", Dbg: difflns[1:], StartLn: 0, StartChr: 0, EndLn: 0, EndChr: 0 } }
			} } } }
		}
	}
	return
}
