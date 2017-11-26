package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/haskell"
)

func main() {
	zhs.OnPreInit()
	z.InitAndServeOrPanic(zhs.OnPostInit)
}
