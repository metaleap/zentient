package z
import (
)


type Zengine interface {
    Ids() []string
}


var Zengines map[string]Zengine
