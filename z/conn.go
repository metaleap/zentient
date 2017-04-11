package z
import (
	"encoding/json"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)

const (
	MSG_ZEN_STATUS		= "ZS:"
	MSG_ZEN_LANGS		= "ZL:"

	MSG_QUERY_CAPS		= "QC:"
	MSG_QUERY_DIAGS		= "QD:"

	MSG_DO_FMT			= "DF:"
	MSG_DO_RENAME		= "DR:"

	MSG_FILES_OPENED	= "FO:"
	MSG_FILES_CLOSED	= "FC:"
	MSG_FILES_WRITTEN	= "FW:"
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

		case MSG_ZEN_LANGS:
			e = out(jsonZengines())
		case MSG_ZEN_STATUS:
			e = out(jsonStatus())

		case MSG_QUERY_DIAGS:
			e = out(jsonLiveDiags("", nil, nil))
		case MSG_QUERY_CAPS:
			resp := map[string][]*RespCmd {}
			for _, zid := range zids { if µ := Zengines[zid] ; µ != nil {
				resp[zid] = µ.Caps(msgargs)  }  }
			e = out(resp)

		case MSG_FILES_WRITTEN:
			onFilesWritten(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], nil, nil))
		case MSG_FILES_OPENED:
			onFilesOpened(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], nil, inlst))
		case MSG_FILES_CLOSED:
			onFilesClosed(Zengines[zids[0]], inlst)
			e = out(jsonLiveDiags(zids[0], inlst, nil))

		case MSG_DO_FMT:
			if resp,err := doFmt(zids[0], ugo.S(inmap["s"]), ugo.S(inmap["c"]), uint8(ugo.F(inmap["t"])))  ;  (err != nil) {
				e = out(err.Error())  } else {  e = out(resp)  }
		case MSG_DO_RENAME:
			e = out("No formatter")

		//  nothing matched? a bug in client, throw at client
		default:
			e = out(jsonErrMsg("Unknown MSG-ID `" + msgid + "` --- for diagnostics, msg-args were: " + msgargs))
	}
	return
}
