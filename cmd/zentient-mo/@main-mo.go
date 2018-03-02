package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/mojo"
)

func main() {
	z.InitAndServeOrPanic(zps.OnPreInit, zps.OnPostInit)
}
