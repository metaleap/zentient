package main

import (
	"path/filepath"

	"github.com/metaleap/go-fromjsonschema"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
)

const srcpath = "github.com/metaleap/zentient/_notes_misc_etc/vscdbgprotocol.json"

func main() {
	gopath := ugo.GoPaths()[0]
	jsonschemaraw := ufs.ReadTextFile(filepath.Join(gopath, "src/"+srcpath), true, "")

	fromjsd.GoPkgDesc = "Package codegen'd from " + srcpath + " via github.com/metaleap/zentient/dbg/zentient-debug-protocol-gen"
	gosrc, err := fromjsd.Generate("zdbgproto", jsonschemaraw)
	if err != nil {
		panic(err)
	}

	ufs.WriteTextFile(filepath.Join(gopath, "src/github.com/metaleap/zentient/dbg/proto/proto.go"), gosrc)
}
