package main

import (
	"fmt"

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

	timetotal, statsperpkg := pkgs.MustRunGentsAndGenerateOutputFiles(nil, gents)
	fmt.Println("total time taken for all parallel runs and INCL. gofmt + file-write :\n\t\t" + timetotal.String())
	for pkg, stats := range statsperpkg {
		fmt.Println("time taken for " + pkg.ImportPath + " EXCL. file-write:\n\t\tconstruct=" + stats.DurationOf.Constructing.String() + "\t\temit=" + stats.DurationOf.Emitting.String() + "\t\tformat=" + stats.DurationOf.Formatting.String() + "")
	}
}
