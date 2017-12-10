package zgo

import (
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
	me.knownDiags = z.Tools{}
}

func (me *goDiag) KnownDiags() z.Tools {
	return me.knownDiags
}
