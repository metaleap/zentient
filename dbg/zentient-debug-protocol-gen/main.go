package main

import (
	"path/filepath"

	"github.com/metaleap/go-fromjsonschema"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
)

const srcpath = "github.com/metaleap/zentient/_notes_misc_etc/vscdbgprotocol.json"
const dstpath = "github.com/metaleap/zentient/dbg/proto/proto.go"

func main() {
	gopath := filepath.Join(ugo.GoPaths()[0], "src")
	jsonschemaraw := ufs.ReadTextFile(filepath.Join(gopath, srcpath), true, "")

	fromjsd.GoPkgDesc = "Package codegen'd from " + srcpath + " with github.com/metaleap/zentient/dbg/zentient-debug-protocol-gen"
	jdefs, err := fromjsd.DefsFromJsonSchema(jsonschemaraw)
	if err != nil { panic(err) }
	gosrc := fromjsd.Generate("zdbgproto", jdefs, map[string]string { "ProtocolMessage": "type", "Event": "event", "Request": "command" })

	if err = ufs.WriteTextFile(filepath.Join(gopath, dstpath), gosrc); err != nil {
		panic(err)
	}
}
