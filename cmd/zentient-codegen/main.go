package main

import (
	. "github.com/go-leap/dev/go/gen"
	"github.com/go-leap/str"
	"github.com/metaleap/go-gent"
	"github.com/metaleap/go-gent/gents/enums"
	"github.com/metaleap/go-gent/gents/json"
)

func main() {
	Self.Name = "me"
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

	typeNames4Marshal := []string{
		"IpcResp", "SrcPos", "SrcRange", "SrcModEdit", "SrcModEdits",
		"SrcLoc", "SrcLens", "SrcLenses", "SrcLocs", "Caddy",
		"SrcIntelSigHelp", "SrcIntelSigInfo", "SrcIntelSigParam", "SrcIntelDoc",
		"SrcIntel", "SrcIntels", "SrcIntelCompl", "SrcIntelCompls",
		"MenuItem", "Menu", "MenuItems", "MenuResponse", "MenuItemArgPrompt",
		"DiagItem", "DiagFixUps", "Diags", "DiagItemsBy", "DiagItems",
		"EditorAction", "ExtrasItem", "Extras", "SrcInfoTip", "SrcAnnotaction",
	}
	typeNames4Unmarshal := []string{
		"IpcReq", "WorkspaceChanges", "SrcLens", "SrcLoc", "SrcPos", "SrcRange",
	}
	gentjson.Gents.OtherTypes.Marshal.MayGenFor = func(t *gent.Type) bool {
		return ustr.In(t.Name, typeNames4Marshal...)
	}
	gentjson.Gents.OtherTypes.Unmarshal.MayGenFor = func(t *gent.Type) bool {
		return ustr.In(t.Name, typeNames4Unmarshal...)
	}
	gentjson.Gents.OtherTypes.RunOnlyForTypes.Named = append(append([]string{}, typeNames4Marshal...), typeNames4Unmarshal...)

	gentjson.Gents.OtherTypes.Marshal.ResliceInsteadOfWhitespace = true
	gentjson.Gents.OtherTypes.Marshal.GenPanicImplsForOthers = true
	gentjson.Gents.OtherTypes.Marshal.OnStdlibFallbacks = onMarshalStdlibAddPrintlnStmt
	gentjson.Gents.OtherTypes.Marshal.TryInterfaceTypesBeforeStdlib = []*TypeRef{
		T.Empty.Interface,
		T.String,
		T.SliceOf.Strings,
		TLocal("MenuItemArgPrompt"),
		TMap(T.String, T.Empty.Interface),
	}
	gentjson.Gents.OtherTypes.Unmarshal.CommonTypesToExtractToHelpers = []*TypeRef{
		T.Empty.Interface,
		T.SliceOf.Strings,
		TSlice(T.Empty.Interface),
		T.String,
		T.Int,
	}
	gentjson.Gents.OtherTypes.Unmarshal.GenPanicImplsForOthers = true
	gentjson.Gents.OtherTypes.Unmarshal.OnStdlibFallbacks = onUnmarshalStdlibAddPrintlnStmt

	timetaken, _ := pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
	println(timetaken.String())
}

func onMarshalStdlibAddPrintlnStmt(ctx *gent.Ctx, fAcc ISyn, s ...ISyn) Syns {
	return append(s,
		B.Panic.Of(fAcc),
	)
}

func onUnmarshalStdlibAddPrintlnStmt(ctx *gent.Ctx, fAcc ISyn, s ...ISyn) Syns {
	return append(s,
		B.Panic.Of(fAcc),
	)
}
