package zgo
import (
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
			resp.Warnings = []string{ "use 'godoc cmd/strings' for documentation on the strings command" }
			resp.Result = "func ToLower(s string) string\n    ToLower returns a copy of the string s with all Unicode letters mapped\n    to their lower case.\n\n\n\nBUGS\n\n   The rule Title uses for word boundaries does not handle Unicode\n   punctuation properly.\n\n"
		default:
			resp.Warnings = []string{ "Unknown querier: " + req.Id }
	}
	return
}
