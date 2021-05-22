package main

import (
	"os"

	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/golang"
)

func main() {
	os.Setenv("GO111MODULE", "off")
	z.InitAndServe(zgo.OnPreInit, zgo.OnPostInit)
}
