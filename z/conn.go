package z
import (
    "encoding/json"
)

const (
    MSG_ZEN_STATE   = "ZS:"
    MSG_ZEN_LANGS   = "ZL:"

    MSG_FILE_OPEN   = "FO:"
    MSG_FILE_CLOSE  = "FC:"
)


// globals set from main-app on init. 'bad style', but.. hey it's golang anyway
var (
    Out         *json.Encoder
    ProjDir     string
    DataDir     string
    DataProjDir string
)




func out (v interface{}) error {
    return Out.Encode(v)
}


func HandleRequest (queryln string) (e error) {
    switch msg := queryln[:3]  ;  msg {

    case MSG_ZEN_STATE:
        e=out("zen status summary")


    case MSG_ZEN_LANGS:
        e=out(jsonableZengines())


    case MSG_FILE_OPEN:


    case MSG_FILE_CLOSE:


    default:
        e=out([]string { DataDir, DataProjDir, msg, queryln })

    }
    return
}
