package main

import (
	"github.com/go-leap/dev/go/gen"
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
		&gentjson.Gents.Structs,
	}
	gentenums.Gents.IsValid.RunOnlyForTypes.Named = []string{"IpcIDs"}
	gentenums.Gents.Stringers.RunNeverForTypes.Named = []string{"ToolCats"}
	gentenums.Gents.Stringers.All[0].SkipEarlyChecks = true

	// temporaries..
	gentjson.Gents.Structs.Marshal.Name, gentjson.Gents.Structs.Unmarshal.Name = "preview_"+gentjson.Gents.Structs.Marshal.Name, "preview_"+gentjson.Gents.Structs.Unmarshal.Name
	gentjson.Gents.Structs.RunOnlyForTypes.Named = []string{
		"SrcPos", "SrcRange", "WorkspaceChanges", "SrcModEdit", "SrcLoc",
		"SrcLens", "SrcIntelSigHelp", "SrcIntelSigInfo", "SrcIntelSigParam",
		"SrcIntelDoc", "SrcIntel", "SrcIntels", "SrcIntelCompl", "SrcInfoTip",
		"SrcAnnotaction", "MenuItem", "Menu", "ExtrasItem", "EditorAction",
		"DiagItem", "DiagFixUps", "Caddy", "MenuResponse", "Extras", "Diags",
		"IpcResp",
	}

	pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
}
