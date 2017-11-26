package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/purescript"
)

func main() {
	zps.OnPreInit()
	z.InitAndServeOrPanic(zps.OnPostInit)
}
