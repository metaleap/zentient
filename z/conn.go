package z
import (
	"encoding/json"

	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
)

const (
	MSG_ZEN_STATUS	= "ZS:"
	MSG_ZEN_LANGS	= "ZL:"

	MSG_CAPS		= "CA:"

	MSG_DO_FMT		= "DF:"

	MSG_FILE_OPEN	= "FO:"
	MSG_FILE_CLOSE	= "FC:"
	MSG_FILE_WRITE	= "FW:"
)


// globals set from main-app on init. 'bad style', but ok for this personal pet project
var (
	Out *json.Encoder
)




func out (v interface{}) error {
	return Out.Encode(v)
}


func HandleRequest (queryln string) (e error) {
	// var instr string
	var inany interface{}
	var inobj map[string]interface{}

	msgid,msgrest := ustr.BreakAt(queryln, 3)
	msgzids,msgargs := ustr.BreakOn(msgrest, ":")
	zids := ustr.Split(msgzids, ",")
	if len(msgargs)>1 && (msgargs[0]=='"' || msgargs[0]=='{' || msgargs[0]=='[' || msgargs[0]=='(' || msgargs[0]=='\'') {
		json.Unmarshal([]byte(msgargs), &inany)
		// instr,_ = inany.(string)
		inobj,_ = inany.(map[string]interface{})
	}
	switch msgid {
		//  each case is ideally just a single func-call out, rpc-like
		//  anything else in a case then is only to furnish proper func args from msg-argstr


		case MSG_ZEN_LANGS:
			e = out(jsonZengines())
		case MSG_ZEN_STATUS:
			e = out(jsonStatus())
		case MSG_FILE_WRITE:
			onFileWrite(Zengines[zids[0]], msgargs)
			e = out(allcurdiags)
		case MSG_FILE_OPEN:
			onFileOpen(Zengines[zids[0]], msgargs)
			e = out(allcurdiags)
		case MSG_FILE_CLOSE:
			onFileClose(Zengines[zids[0]], msgargs)
			e = out(allcurdiags)
		case MSG_CAPS:
			resp := map[string][]*RespCmd {}
			for _, zid := range zids { if µ := Zengines[zid] ; µ != nil {
				resp[zid] = µ.Caps(msgargs)  }  }
			e = out(resp)
		case MSG_DO_FMT:
			if resp,err := doFmt(zids[0], ugo.S(inobj["s"]), ugo.S(inobj["c"]), uint8(ugo.F(inobj["t"])))  ;  (err != nil) {
				e = out(err.Error())  } else {  e = out(resp)  }


		//  nothing matched? a bug in client, throw at client
		default:
			e = out(jsonErrMsg("Unknown MSG-ID `" + msgid + "` --- for diagnostics, msg-args were: " + msgargs))
	}
	return
}
