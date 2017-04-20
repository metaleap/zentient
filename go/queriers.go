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
			var cmd = ustr.Split(req.Sym2, " ")
			if devgo.Has_godoc { cmd = append(ustr.N("godoc"), cmd...) } else { cmd = append(ustr.N("go", "doc"), cmd...) }
			cmdout,cmderr,_ := ugo.CmdExecStdin ("", filepath.Dir(req.Ffp), cmd[0], cmd[1:]...)
			resp.Warnings = uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim)
			//use 'godoc cmd/strings' for documentation on the strings command \n\n
			if ustr.Pref(cmdout, "use 'godoc cmd/") { cmdout = cmdout[ustr.Idx(cmdout, "\n"):] }
			lns := ustr.Split(cmdout, "\n")
			for i,ln := range lns {
				if len(ln)>0 && !ustr.Pref(ln, " ") { lns[i] = "</p><h2>" + ustr.Trim(ln) + "</h2><p>" } else { lns[i] = ustr.Trim(ln) }
			}
			cmdout = strings.Replace(ustr.Trim(ustr.Join(lns, "\n")), "\n\n", "</p><p>", -1)
			if !ustr.Pref(cmdout, "<p>") { cmdout = "<p>" + cmdout }  ;  if !ustr.Suff(cmdout, "</p>") { cmdout = cmdout + "</p>" }
			resp.Result = cmdout
		default:
			resp.Warnings = []string{ "Unknown querier: " + req.Id }
	}
	return
}
