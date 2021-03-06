package main

import (
	"path/filepath"
	"time"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/fs"
	"github.com/metaleap/go-fromjsonschema"
)

const (
	origuri = "https://raw.githubusercontent.com/Microsoft/vscode-debugadapter-node/master/debugProtocol.json"
	srcpath = "github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol/vscdbgprotocol.json"
	dstpath = "github.com/metaleap/zentient/dbg/vsc/protocol/protocol.go"
)

func main() {
	gopathsrc := filepath.Join(udevgo.Gopaths()[0], "src")
	jsonschemaraw := ufs.ReadTextFileOrPanic(filepath.Join(gopathsrc, srcpath))

	fromjsd.GoPkgDesc = "Package codegen'd on " + time.Now().Format(time.RFC3339) + " from " + srcpath + " with github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol"
	jsd, err := fromjsd.NewJsonSchema(jsonschemaraw)
	if err != nil {
		panic(err)
	}

	jsd.Defs["DisconnectArguments"].EnsureProps(nil)
	jsd.ForceCopyProps("Request", "Response", "command")

	unmarshalHints := map[string]string{"ProtocolMessage": "type", "Event": "event", "Request": "command", "Response": "command"}
	handlerBaseTypes := map[string]string{"Request": "Response"}
	gosrc := jsd.Generate("zdbgvscp", unmarshalHints, handlerBaseTypes, "Event", "Response", "Request")

	if err = ufs.WriteTextFile(filepath.Join(gopathsrc, dstpath), gosrc); err != nil {
		panic(err)
	}
}
