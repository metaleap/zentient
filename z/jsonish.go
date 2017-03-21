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
    resp["dirs"] = map[string]string { "proj": ProjDir, "data": DataDir, "dataproj": DataProjDir }
    resp["zengines"] = jsonZengines()
    resp["files"] = jsonFiles()
    return resp
}

func jsonFiles () interface{} {
    return AllFiles
}


func jsonZengines () interface{} {
    list := map[string][]string {}
    for id, zengine := range Zengines {
        list[id] = zengine.Ids()
    }
    return list
}
