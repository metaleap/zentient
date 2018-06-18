package zhs

import (
	"github.com/metaleap/zentient"
)

var diag hsDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type hsDiag struct {
	z.DiagBase

	knownTools z.Tools
}

func (this *hsDiag) onPreInit() {
	this.knownTools = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}
