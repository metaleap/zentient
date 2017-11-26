package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/haskell"
)

func main() {
	z.InitAndServeOrPanic(zhs.OnPreInit, zhs.OnPostInit)
}
