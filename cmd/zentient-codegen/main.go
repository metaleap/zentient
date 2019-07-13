package main

import (
	"github.com/go-leap/dev/go/gen"
	"github.com/metaleap/go-gent"
	"github.com/metaleap/go-gent/gents/enums"
	"github.com/metaleap/go-gent/gents/json"
	"github.com/metaleap/go-gent/gents/structs"
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
		&gentstructs.Gents.StructFieldsTrav,
		&gentstructs.Gents.StructFieldsGetSet,
	}
	gentenums.Gents.IsValid.RunOnlyForTypes.Named = []string{"IpcIDs"}
	gentenums.Gents.Stringers.RunNeverForTypes.Named = []string{"ToolCats"}
	gentenums.Gents.Stringers.All[0].SkipEarlyChecks = true

	// temporaries..
	gentjson.Gents.Structs.Marshal.MethodName, gentjson.Gents.Structs.Unmarshal.MethodName = "preview_"+gentjson.Gents.Structs.Marshal.MethodName, "preview_"+gentjson.Gents.Structs.Unmarshal.MethodName
	gentjson.Gents.Structs.RunOnlyForTypes.Named = []string{"fooResp"}
	gentstructs.Gents.StructFieldsTrav.RunOnlyForTypes.Named = []string{"fooResp"}
	gentstructs.Gents.StructFieldsGetSet.RunOnlyForTypes.Named = []string{"fooResp"}

	pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
}
