package main

import (
	"path/filepath"

	"github.com/metaleap/go-fromjsonschema"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
)

const srcpath = "github.com/metaleap/zentient/zdbg-vsc/_notes_misc_etc/vscdbgprotocol.json"
const dstpath = "github.com/metaleap/zentient/zdbg-vsc/proto/proto.go"

func main() {
	gopath := filepath.Join(ugo.GoPaths()[0], "src")
	jsonschemaraw := ufs.ReadTextFile(filepath.Join(gopath, srcpath), true, "")

	fromjsd.GoPkgDesc = "Package codegen'd from " + srcpath + " with github.com/metaleap/zentient/zdbg-vsc-proto-gen"
	jdefs, err := fromjsd.NewJsonSchema(jsonschemaraw)
	if err != nil { panic(err) }

	jdefs.Defs["DisconnectArguments"].EnsureProps(map[string]string { "restart": "boolean" })
	jdefs.Defs["LaunchRequestArguments"].EnsureProps(map[string]string { "w": "string", "c": "string", "f": "string" })

	unmarshalHints := map[string]string { "ProtocolMessage": "type", "Event": "event", "Request": "command" }
	handlerBaseTypes := map[string]string{ "Request": "Response" }
	gosrc := jdefs.Generate("zdbgvscp", unmarshalHints, handlerBaseTypes, "Event", "Response")

	if err = ufs.WriteTextFile(filepath.Join(gopath, dstpath), gosrc); err != nil {
		panic(err)
	}
}
