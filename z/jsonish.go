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

type RespRen struct {
	NewText		string
	Dbg			interface{}
	StartLn		uint64
	StartChr	uint64
	EndLn		uint64
	EndChr		uint64
}


const (
	DIAG_SEV_ERR	uint8 = 0
	DIAG_SEV_WARN	uint8 = 1
	DIAG_SEV_INFO	uint8 = 2
	DIAG_SEV_HINT	uint8 = 3
)


var (
	newlivediags = true
)


func jsonLiveDiags (frpszid string, closedfrps []string, openedfrps []string) (jld map[string]map[string][]*RespDiag) {
	if len(closedfrps)>0 || len(openedfrps)>0 {  newlivediags = true  }
	if newlivediags {
		diagsready := true  ;  jld = map[string]map[string][]*RespDiag {}  ;  var fc, fo []string
		for zid,µ := range Zengines {
			if (!µ.ReadyToBuildAndLint()) { diagsready = false }
			if zid==frpszid { fc,fo = closedfrps,openedfrps } else { fc,fo = nil,nil }
			jld[zid] = µ.B().liveDiags(µ, fc, fo)
		}
		if diagsready { newlivediags = false }
	}
	return // if diags haven't changed since last req, send nil
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
