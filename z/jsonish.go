package z
import (
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)



type ReqIntel struct {
	Ffp		string	`json:",omitempty"`
	Pos		string	`json:",omitempty"`
	Src		string	`json:",omitempty"`
	Sym1 	string	`json:",omitempty"`
	Sym2 	string	`json:",omitempty"`
	CrLf	bool	`json:",omitempty"`
}

func (self *ReqIntel) RunePosToBytePos () {
	if len(self.Src)==0 {  self.Src = ufs.ReadTextFile(self.Ffp, false, "")  }
	if off := int(ustr.ParseInt(self.Pos))  ;  off>0 && len(self.Src)>0 {
		reoff := func() int { r := 0  ;  for i, _ := range self.Src { if r==off { return i }  ;  { r++ } }  ;  return len([]byte(self.Src)) }
		self.Pos = ugo.SPr(reoff())
	}
}


type RespCmd struct {
	Name	string		`json:",omitempty"`		//	actual cmd name
	Args	[]string	`json:",omitempty"`	//	args

	Title	string		`json:",omitempty"`	//	display name, eg: N = "go vet" when C = "go" with A = ["vet"]  ;  if empty fall back to C
	Exists	bool		`json:",omitempty"`	//	installed?
	Hint	string		`json:",omitempty"`	//	install hint
	More	string		`json:",omitempty"`

	f	func()		//	tmp field used in Base.DoFmt()
}

type RespDiag struct {
	udev.SrcMsg
	Sev uint8
}


type RespTxt struct {
	Result		string		`json:",omitempty"`
	Warnings	[]string	`json:",omitempty"`
}

type RespCmpl struct {
	Label		string		`json:"label,omitempty"`
	Kind		int			`json:"kind,omitempty"`
	Detail		string		`json:"detail,omitempty"`
	Doc			string		`json:"documentation,omitempty"`
	SortTxt		string		`json:"sortText,omitempty"`
	FilterTxt	string		`json:"filterText,omitempty"`
	InsertTxt	string		`json:"insertText,omitempty"`
	CommitChars	[]string	`json:"commitCharacters,omitempty"`
}

type RespHov struct {
	Txt		string	`json:"value,omitempty"`
	Lang	string	`json:"language,omitempty"`
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
	// resp["livediags"] = Zengines["go"].B().livediags
	// resp["lintdiags"] = Zengines["go"].B().lintdiags
	// resp["builddiags"] = Zengines["go"].B().builddiags
	resp["Ctx"] = Ctx
	resp["OpenFiles"] = OpenFiles
	resp["AllFiles"] = AllFiles
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
