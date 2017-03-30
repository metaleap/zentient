package z
import (
)


// the ONLY jsonish func to return a string-encoded-as-JSON-value
// thereby establishing convention/protocol for clients:
// if the response is such, it's to be interpreted as a reportable error
func jsonErrMsg (msg string) interface{} {
	return msg
}


func jsonStatus () interface{} {
	resp := map[string]interface{} {}
	resp["Root"] = Root
	resp["Zengines"] = jsonZengines()
	for zid, zengine := range Zengines {
		resp["Zengines["+zid+"]"] = zengine
	}
	resp["OpenFiles"] = OpenFiles
	resp["AllFiles"] = AllFiles
	return resp
}


func jsonZengines () interface{} {
	list := map[string][]string {} // ouch =)
	for zid, zengine := range Zengines {
		list[zid] = zengine.Ids()
	}
	return list
}
