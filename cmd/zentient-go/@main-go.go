package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/golang"
)

func main() {
	zgo.OnPreInit()
	z.InitAndServeOrPanic(zgo.OnPostInit)
}
