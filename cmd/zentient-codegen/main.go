package main

import (
	"github.com/go-leap/dev/go/gen"
	"github.com/go-leap/str"
	"github.com/metaleap/go-gent"
	"github.com/metaleap/go-gent/gents/enums"
	"github.com/metaleap/go-gent/gents/json"
)

func main() {
	udevgogen.Self.Name = "me"
	pkgs := gent.MustLoadPkgs(map[string]string{
		"github.com/metaleap/zentient": "°gent.go",
	})

	gents := gent.Gents{
		&gentenums.Gents.IsValid,
		&gentenums.Gents.Stringers,
		&gentjson.Gents.OtherTypes,
	}

	gentenums.Gents.IsValid.RunOnlyForTypes.Named = []string{"IpcIDs"}

	gentenums.Gents.Stringers.RunNeverForTypes.Named = []string{"ToolCats"}
	gentenums.Gents.Stringers.All[0].SkipEarlyChecks = true

	gentjson.Gents.OtherTypes.Marshal.Name, gentjson.Gents.OtherTypes.Unmarshal.Name =
		"preview_"+gentjson.Gents.OtherTypes.Marshal.Name, "preview_"+gentjson.Gents.OtherTypes.Unmarshal.Name
	marshalTypes, unmarshalTypes := []string{
		"IpcResp", "SrcPos", "SrcRange", "SrcModEdit", "SrcLoc", "SrcLens",
		"SrcIntelSigHelp", "SrcIntelSigInfo", "SrcIntelSigParam", "SrcIntelDoc",
		"SrcIntel", "SrcIntels", "SrcIntelCompl", "SrcInfoTip", "SrcAnnotaction",
		"MenuItem", "Menu", "ExtrasItem", "MenuResponse", "Extras",
		"DiagItem", "DiagFixUps", "Caddy", "Diags", "EditorAction",
	}, []string{
		"IpcReq", "WorkspaceChanges", "SrcLens", "SrcLoc", "SrcPos", "SrcRange",
	}
	gentjson.Gents.OtherTypes.Marshal.MayGenFor = func(t *gent.Type) bool {
		return ustr.In(t.Name, marshalTypes...)
	}
	gentjson.Gents.OtherTypes.Unmarshal.MayGenFor = func(t *gent.Type) bool {
		return ustr.In(t.Name, unmarshalTypes...)
	}
	gentjson.Gents.OtherTypes.RunOnlyForTypes.Named =
		append(marshalTypes, unmarshalTypes...)

	pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
}
