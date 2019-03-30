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

func (me *hsDiag) onPreInit() {
	me.knownTools = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}
