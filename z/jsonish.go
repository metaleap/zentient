package z
import (
	"github.com/metaleap/go-util-dev"
)


type RespCmd struct {
	Name	string		//	actual cmd name
	Args	[]string	//	args

	Title	string		//	display name, eg: N = "go vet" when C = "go" with A = ["vet"]  ;  if empty fall back to C
	Exists	bool		//	installed?
	Hint	string		//	install hint

	f	func()		//	tmp field used in Base.DoFmt()
}

type RespDiag struct {
	udev.SrcMsg
	Sev uint8
}


type RespFmt struct {
	Result		string
	Warnings	[]string
}


const (
	DIAG_SEV_ERR	uint8 = 0
	DIAG_SEV_WARN	uint8 = 1
	DIAG_SEV_INFO	uint8 = 2
	DIAG_SEV_HINT	uint8 = 3
)


func jsonLiveDiags (closedfrp string, openedfrp string) (livediags map[string]map[string][]*RespDiag) {
	livediags = map[string]map[string][]*RespDiag {}
	fc := AllFiles[closedfrp]  ;  fo := AllFiles[openedfrp]
	for zid,µ := range Zengines {
		var fcrp, forp string
		if fc!=nil && fc.µ==µ { fcrp = fc.RelPath } else { fcrp = "" }
		if fo!=nil && fo.µ==µ { forp = fo.RelPath } else { forp = "" }
		livediags[zid] = µ.B().liveDiags(µ, fcrp, forp)
	}
	return
}


// the ONLY jsonish func to return a string-encoded-as-JSON-value
// thereby establishing convention/protocol for clients:
// if the response is such, it's to be interpreted as a reportable error
func jsonErrMsg (msg string) interface{} {
	return msg
}


func jsonStatus () interface{} {
	resp := map[string]interface{} {}
	resp["livediags"] = Zengines["go"].B().livediags
	resp["lintdiags"] = Zengines["go"].B().lintdiags
	resp["builddiags"] = Zengines["go"].B().builddiags
	// resp["Ctx"] = Ctx
	// resp["OpenFiles"] = OpenFiles
	// resp["AllFiles"] = AllFiles
	// resp["Zengines"] = jsonZengines()
	// for zid, zengine := range Zengines {
	// 	resp["Zengines["+zid+"]"] = zengine
	// }
	return resp
}


func jsonZengines () interface{} {
	list := map[string][]string {} // ouch =)
	for zid, zengine := range Zengines {
		list[zid] = zengine.EdLangIDs()
	}
	return list
}
