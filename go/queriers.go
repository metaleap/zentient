package zgo
import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)

var queryTools = []*z.RespPick {
		&z.RespPick{ Label: "GoDoc", Detail: "[package] [member name] – shows the specified item's documentation topic.", Desc: "godoc" },
	}
func (me *zgo) QueryTools () []*z.RespPick {
	return queryTools
}


func (_ *zgo) QueryTool (req *z.ReqIntel) (resp *z.RespTxt) {
	resp = &z.RespTxt{ Id: req.Id }
	switch req.Id {
		case "godoc":
			req.Sym2 = ustr.Trim(req.Sym2)
			if i1,i2 := ustr.Idx(req.Sym2, ".") , ustr.Idx(req.Sym2, " ")  ;  i1>0 && (i2<0 || i2>i1) { req.Sym2 = req.Sym2[:i1] + " " + req.Sym2[i1+1:] }
			var cmd = ustr.Split(req.Sym2, " ")  ;  if devgo.PkgsByImP!=nil && ustr.IsLower(cmd[0][:1]) && devgo.PkgsByImP[cmd[0]]==nil {
				if dp := filepath.Join(srcDir, cmd[0])  ;  devgo.PkgsByDir!=nil && devgo.PkgsByDir[dp]!=nil {
					cmd[0] = devgo.PkgsByDir[dp].ImportPath
				} else { for _,pkg := range devgo.PkgsByImP {
					if pkg.Name==cmd[0] { cmd[0] = pkg.ImportPath  ;  break }
				} }
			}
			if devgo.Has_godoc { cmd = append(ustr.N("godoc", "-ex"), cmd...) } else { cmd = append(ustr.N("go", "doc"), cmd...) }
			cmdout,cmderr,_ := ugo.CmdExecStdin ("", filepath.Dir(req.Ffp), cmd[0], cmd[1:]...)
			resp.Warnings = uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim)
			if ustr.Pref(cmdout, "use 'godoc cmd/") { cmdout = cmdout[ustr.Idx(cmdout, "\n"):] }
			prep,p,lns := "" , 0 , ustr.Split(cmdout, "\n")  ;  for i,ln := range lns {
				if len(ln)>0 && len(ustr.LettersOnly(ln))>0 && !ustr.Pref(ln, " ") {
					if ln = ustr.Trim(ln)  ;  ustr.IsUpper(ln) { lns[i] = "</p><h1>" + ln + "</h1><p>" } else {
						lns[i] = "</p><h2 id='" + ugo.SPr(p) + "'>" + ln + "</h2><p>"
						prep += "<li><a href='#" + ugo.SPr(p) + "'>" + ln + "</a></li>"  ;  p++
					}
				} else { lns[i] = ustr.Trim(ln) }
			}
			if cmdout = strings.Replace(ustr.Trim(ustr.Join(lns, "\n")), "\n\n", "</p><p>", -1)  ;  len(cmdout)>0 {
				if !ustr.Pref(cmdout, "<p>") { cmdout = "<p>" + cmdout }  ;  if !ustr.Suff(cmdout, "</p>") { cmdout = cmdout + "</p>" }
			}
			if ustr.Has(prep, "</li><li>") {  cmdout = "<ul>" + prep + "</ul>" + cmdout  }
			resp.Result = cmdout
		default:
			resp.Warnings = []string{ "Unknown querier: " + req.Id }
	}
	return
}
