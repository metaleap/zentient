package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/haskell"
)

func main() {
	z.InitAndServe(zhs.OnPreInit, zhs.OnPostInit)
}
