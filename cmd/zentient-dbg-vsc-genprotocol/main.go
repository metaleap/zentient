package main

import (
	"path/filepath"

	"github.com/metaleap/go-fromjsonschema"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
)

const (
	srcpath = "github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol/vscdbgprotocol.json"
	dstpath = "github.com/metaleap/zentient/dbg/vsc/protocol/protocol.go"
)

func main() {
	gopathsrc := filepath.Join(udevgo.AllGoPaths()[0], "src")
	jsonschemaraw := ufs.ReadTextFile(filepath.Join(gopathsrc, srcpath), true, "")

	fromjsd.GoPkgDesc = "Package codegen'd from " + srcpath + " with github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol"
	jsd, err := fromjsd.NewJsonSchema(jsonschemaraw)
	if err != nil {
		panic(err)
	}

	jsd.Defs["DisconnectArguments"].EnsureProps(map[string]string{"restart": "boolean"})
	// jsd.Defs["LaunchRequestArguments"].EnsureProps(map[string]string{})
	jsd.ForceCopyProps("Request", "Response", "command")

	unmarshalHints := map[string]string{"ProtocolMessage": "type", "Event": "event", "Request": "command", "Response": "command"}
	handlerBaseTypes := map[string]string{"Request": "Response"}
	gosrc := jsd.Generate("zdbgvscp", unmarshalHints, handlerBaseTypes, "Event", "Response", "Request")

	if err = ufs.WriteTextFile(filepath.Join(gopathsrc, dstpath), gosrc); err != nil {
		panic(err)
	}
}
