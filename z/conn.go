package z
import (
	"encoding/json"
	"strings"

	"github.com/metaleap/go-util-str"
)

const (
	MSG_ZEN_STATUS	= "ZS:"
	MSG_ZEN_LANGS	= "ZL:"

	MSG_CAP_FMT		= "CF:"

	MSG_DO_FMT		= "DF:"

	MSG_FILE_OPEN	= "FO:"
	MSG_FILE_CLOSE	= "FC:"
	MSG_FILE_WRITE	= "FW:"
)


// globals set from main-app on init. 'bad style', but ok for this personal pet project
var (
	Out         *json.Encoder
)




func out (v interface{}) error {
	return Out.Encode(v)
}


func HandleRequest (queryln string) (e error) {
	var str string
	msgid,msgargs := ustr.BreakAt(queryln, 3)
	switch msgid {
		//  each case is ideally just a single func-call out, rpc-like
		//  anything else in a case then is only to furnish proper func args from msg-argstr

		//  FIRST: CASES THAT EXPECT A RESPONSE
		case MSG_ZEN_LANGS:
			e = out(jsonZengines())
		case MSG_ZEN_STATUS:
			e = out(jsonStatus())
		case MSG_CAP_FMT:
			zids := strings.Split(msgargs, ",")
			resp := map[string][]string {}
			for _, zid := range zids {
				if µ := Zengines[zid] ; µ != nil {
					resp[zid] = µ.Caps("fmt")
				}
			}
			e = out(resp)
		case MSG_DO_FMT:
			zid,injson := ustr.BreakOn(msgargs, ":")
			resp := map[string]string {}
			if e = json.Unmarshal([]byte(injson), &str)  ;  e == nil {
				if µ := Zengines[zid]  ;  µ != nil {  resp[zid] = µ.DoFmt(str)  }
				e = out(resp)
			}


		//  LAST: CASES THAT RECEIVE NO RESPONSE
		//  no error reporting to client either, for now. with some luck, it can all stay that way
		case MSG_FILE_OPEN:
			if z, relpath := fromZidMsg(msgargs) ; z != nil {  onFileOpen(z, relpath)  }
		case MSG_FILE_CLOSE:
			if z, relpath := fromZidMsg(msgargs) ; z != nil {  onFileClose(z, relpath)  }
		case MSG_FILE_WRITE:
			if z, relpath := fromZidMsg(msgargs) ; z != nil {  onFileWrite(z, relpath)  }

		//  NOTHING MATCHED? A BUG IN CLIENT, throw at client
		default:
			e = out(jsonErrMsg("Unknown MSG-ID `" + msgid + "` --- for diagnostics, msg-args were: " + msgargs))
	}
	return
}
