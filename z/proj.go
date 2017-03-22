package z
import (
)


type Proj struct {
    NameShort       string      //  semantics of both depend on the zengine
    NameQualified   string

    Dir             string      //  root-relative src directory path
    File            string      //  root-relative project file path --- semantics depend on zengine, could be empty for Golang or either a stack.yaml or project.cabal for Haskell
}
