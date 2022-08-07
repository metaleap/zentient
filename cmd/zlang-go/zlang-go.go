package main

import (
	"os"

	z "github.com/metaleap/zentient"
	zgo "github.com/metaleap/zentient/lang/golang"
)

func main() {
	os.Setenv("GO111MODULE", "off")
	z.InitAndServe(zgo.OnPreInit, zgo.OnPostInit)
}
