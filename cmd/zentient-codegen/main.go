package main

import (
	"github.com/metaleap/go-gent"
	"github.com/metaleap/go-gent/gents/enums"
)

func main() {
	pkgs := gent.MustLoadPkgs(map[string]string{
		"github.com/metaleap/zentient": "°_gent.go",
	})

	gents := gent.Gents{
		&gentenums.Gents.IsValid,
		&gentenums.Gents.Stringers,
	}
	gentenums.Gents.IsValid.RunOnlyForTypes.Named = []string{"IpcIDs"}
	gentenums.Gents.Stringers.RunNeverForTypes.Named = []string{"ToolCats"}
	gentenums.Gents.Stringers.All[0].SkipEarlyChecks = true

	pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
}
