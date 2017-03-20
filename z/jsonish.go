package z
import (
)


func jsonableZengines () interface{} {
    list := map[string][]string {}
    for id, zengine := range Zengines {
        list[id] = zengine.Ids()
    }
    return list
}
