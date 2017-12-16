package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var diag goDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type goDiag struct {
	z.DiagBase

	knownDiags z.Tools
}

func (me *goDiag) onPreInit() {
	me.knownDiags = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}

func (me *goDiag) KnownDiags() z.Tools {
	return me.knownDiags
}

func (me *goDiag) runDiag(tool *z.Tool, pkg *udevgo.Pkg, diagchan chan *z.DiagItem, onjobdone func()) {
	defer onjobdone()
	// gosimple golint goconst

	diagchan <- &z.DiagItem{Message: "Msg1from:" + tool.Name}
	println(tool.Name + ">>>" + pkg.ImportPath)
	diagchan <- &z.DiagItem{Message: "Msg2from:" + tool.Name}
}

func (me *goDiag) UpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) {
	if pkgs := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		var diagitems []*z.DiagItem
		numjobs, numdone, diagchan := 0, 0, make(chan *z.DiagItem)
		onjobdone := func() { diagchan <- nil }
		for _, pkg := range pkgs {
			for _, diagtool := range diagTools {
				numjobs++
				go me.runDiag(diagtool, pkg, diagchan, onjobdone)
			}
		}
		for numdone < numjobs {
			select {
			case diagitem := <-diagchan:
				if diagitem == nil {
					numdone++
				} else {
					diagitems = append(diagitems, diagitem)
				}
			}
		}
		println(len(diagitems))
	}
}
