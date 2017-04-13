package z
import (
	"github.com/metaleap/go-util-dev"
)



type ReqIntel struct {
	Ffp	string
	Pos	string
	EoL	string
	Src	string
}

type RespCmd struct {
	Name	string		//	actual cmd name
	Args	[]string	//	args

	Title	string		//	display name, eg: N = "go vet" when C = "go" with A = ["vet"]  ;  if empty fall back to C
	Exists	bool		//	installed?
	Hint	string		//	install hint
	For		string

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

type RespCmpl struct {
	Label		string	`json:"label"`
	Kind		int		`json:"kind"`
	Detail		string	`json:"detail"`
	Doc			string	`json:"documentation"`
	SortTxt		string	`json:"sortText"`
	FilterTxt	string	`json:"filterText"`
	InsertTxt	string	`json:"insertText"`
}

type RespHov struct {
	Txt		string	`json:"value"`
	Lang	string	`json:"language"`
}


const (
	DIAG_SEV_ERR	= 0
	DIAG_SEV_WARN	= 1
	DIAG_SEV_INFO	= 2
	DIAG_SEV_HINT	= 3
)

const (
	CMPL_TEXT			= 0
	CMPL_METHOD			= 1
	CMPL_FUNCTION		= 2
	CMPL_CONSTRUCTOR	= 3
	CMPL_FIELD			= 4
	CMPL_VARIABLE		= 5
	CMPL_CLASS			= 6
	CMPL_INTERFACE		= 7
	CMPL_STRUCT			= 21
	CMPL_MODULE			= 8
	CMPL_PROPERTY		= 9
	CMPL_UNIT			= 10
	CMPL_VALUE			= 11
	CMPL_CONSTANT		= 20
	CMPL_ENUM			= 12
	CMPL_ENUMMEMBER		= 19
	CMPL_KEYWORD		= 13
	CMPL_SNIPPET		= 14
	CMPL_COLOR			= 15
	CMPL_REFERENCE		= 17
	CMPL_FILE			= 16
	CMPL_FOLDER			= 18
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
