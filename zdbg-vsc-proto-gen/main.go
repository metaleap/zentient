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
	jsd, err := fromjsd.NewJsonSchema(jsonschemaraw)
	if err != nil { panic(err) }

	jsd.Defs["DisconnectArguments"].EnsureProps(map[string]string { "restart": "boolean" })
	jsd.Defs["LaunchRequestArguments"].EnsureProps(map[string]string { "w": "string", "c": "string", "f": "string" })
	jsd.ForceCopyProps("Request", "Response", "command")

	unmarshalHints := map[string]string { "ProtocolMessage": "type", "Event": "event", "Request": "command", "Response": "command" }
	handlerBaseTypes := map[string]string{ "Request": "Response" }
	gosrc := jsd.Generate("zdbgvscp", unmarshalHints, handlerBaseTypes, "Event", "Response", "Request")

	if err = ufs.WriteTextFile(filepath.Join(gopath, dstpath), gosrc); err != nil {
		panic(err)
	}
}
