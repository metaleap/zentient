package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/golang"
)

func main() {
	z.InitAndServeOrPanic(zgo.OnPreInit, zgo.OnPostInit)
}
