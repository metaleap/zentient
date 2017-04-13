package z
import (
	"encoding/json"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)

const (
	REQ_ZEN_STATUS		= "ZS:"
	REQ_ZEN_LANGS		= "ZL:"
	REQ_ZEN_CONFIG		= "ZC:"

	REQ_QUERY_CAPS		= "QC:"
	REQ_QUERY_DIAGS		= "QD:"

	REQ_INTEL_DEFLOC	= "IL:"
	REQ_INTEL_HOVER		= "IH:"
	REQ_INTEL_CMPL		= "IC:"

	REQ_DO_FMT			= "DF:"
	REQ_DO_RENAME		= "DR:"

	REQ_FILES_OPENED	= "FO:"
	REQ_FILES_CLOSED	= "FC:"
	REQ_FILES_WRITTEN	= "FW:"
)


// globals set from main-app on init. 'bad style', but ok for this personal pet project
var (
	Out *json.Encoder
)




func out (v interface{}) error {
	return Out.Encode(v)
}


func HandleRequest (queryln string) (e error) {
	var inlst []string
	var inmap map[string]interface{}

	msgid,msgrest := ustr.BreakAt(queryln, 3)
	msgzids,msgargs := ustr.BreakOn(msgrest, ":")
	zids := ustr.Split(msgzids, ",")
	if len(msgargs)>1 {
		if (msgargs[0]=='{') { json.Unmarshal([]byte(msgargs), &inmap) } else
			if (msgargs[0]=='[') { json.Unmarshal([]byte(msgargs), &inlst) }
	}
	switch msgid {
		//  each case is ideally just a single func-call out, rpc-like
		//  anything else in a case then is only to furnish proper func args from msg-argstr

		case REQ_INTEL_DEFLOC:
			if resp,err := Zengines[zids[0]].IntelDefLoc(ugo.S(inmap["ffp"]), ugo.S(inmap["i"]), ugo.S(inmap["o"]))  ;  resp!=nil {
				e = out(resp) } else if err!=nil { e = out(err.Error()) } else { e = out(nil) }
		case REQ_INTEL_HOVER:
			e = out(Zengines[zids[0]].IntelHovs(ugo.S(inmap["ffp"]), ugo.S(inmap["i"]), ugo.S(inmap["o"])))
		case REQ_INTEL_CMPL:
			e = out(Zengines[zids[0]].IntelCmpl(ugo.S(inmap["ffp"]), ugo.S(inmap["i"]), ugo.S(inmap["o"])))

		case REQ_FILES_WRITTEN:
			onFilesWritten(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], nil, nil))
		case REQ_FILES_OPENED:
			onFilesOpened(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], nil, inlst))
		case REQ_FILES_CLOSED:
			onFilesClosed(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], inlst, nil))

		case REQ_DO_FMT:
			if resp,err := doFmt(zids[0], ugo.S(inmap["s"]), ugo.S(inmap["c"]), uint8(ugo.F(inmap["t"])))  ;  (err != nil) {
				e = out(err.Error())  } else {  e = out(resp)  }
		case REQ_DO_RENAME:
			if resp,err := doRename(zids[0], ugo.S(inmap["c"]), ugo.S(inmap["rfp"]), uint64(ustr.ParseInt(ugo.S(inmap["o"]))), ugo.S(inmap["nn"]), ugo.S(inmap["e"]), ugo.S(inmap["no"]), uint64(ustr.ParseInt(ugo.S(inmap["o1"]))), uint64(ustr.ParseInt(ugo.S(inmap["o2"]))))  ;  (err != nil) {
				e = out(err.Error())  } else {  e = out(resp)  }

		case REQ_QUERY_DIAGS:
			e = out(jsonLiveDiags("", nil, nil))
		case REQ_QUERY_CAPS:
			resp := map[string][]*RespCmd {}
			for _, zid := range zids { if µ := Zengines[zid] ; µ != nil {
				resp[zid] = µ.Caps(msgargs)  }  }
			e = out(resp)

		case REQ_ZEN_LANGS:
			e = out(jsonZengines())
		case REQ_ZEN_STATUS:
			e = out(jsonStatus())
		case REQ_ZEN_CONFIG:
			Zengines[zids[0]].OnCfg(inmap)
			e = out(nil)

		//  nothing matched? a bug in client, throw at client
		default:
			e = out(jsonErrMsg("Unknown MSG-ID `" + msgid + "` --- for diagnostics, msg-args were: " + msgargs))
	}
	return
}
