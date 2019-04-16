package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/atmo"
)

func main() {
	z.InitAndServeOrPanic(zat.OnPreInit, zat.OnPostInit)
}
