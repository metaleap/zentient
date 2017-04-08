package z
import (
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
	Data	string
	Msg		string
	PosLn	uint32
	PosCol	uint32
	Pos2Ln	uint32
	Pos2Col	uint32
	Sev		uint8
	Cat		string
}

type RespFmt struct {
	Result		string
	Warnings	[]string
}


const (
	DIAG_ERR	uint8 = 0
	DIAG_WARN	uint8 = 1
	DIAG_INFO	uint8 = 2
	DIAG_HINT	uint8 = 3
)


func jsonCurDiags () (allcurdiags map[string]map[string][]*RespDiag) {
	allcurdiags = map[string]map[string][]*RespDiag {}
	for zid,µ := range Zengines { allcurdiags[zid] = µ.B().curdiags }
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
	resp["Ctx"] = Ctx
	resp["OpenFiles"] = OpenFiles
	resp["AllFiles"] = AllFiles
	resp["Zengines"] = jsonZengines()
	for zid, zengine := range Zengines {
		resp["Zengines["+zid+"]"] = zengine
	}
	return resp
}


func jsonZengines () interface{} {
	list := map[string][]string {} // ouch =)
	for zid, zengine := range Zengines {
		list[zid] = zengine.EdLangIDs()
	}
	return list
}
