package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/purescript"
)

func main() {
	z.InitAndServeOrPanic(zps.OnPreInit, zps.OnPostInit)
}
