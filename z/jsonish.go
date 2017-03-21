package z
import (
)


func jsonStatus () interface{} {
    resp := map[string]interface{} {}
    resp["dirs"] = map[string]string { "proj": ProjDir, "data": DataDir, "dataproj": DataProjDir }
    resp["zengines"] = jsonZengines()
    return resp
}


func jsonZengines () interface{} {
    list := map[string][]string {}
    for id, zengine := range Zengines {
        list[id] = zengine.Ids()
    }
    return list
}
