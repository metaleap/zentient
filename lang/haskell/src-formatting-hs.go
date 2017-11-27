package zhs

import (
	"github.com/metaleap/zentient"
)

type srcFormatting struct {
	z.SrcFormattingBase

	knownFormatters []*z.Tool
}

var (
	srcFmt srcFormatting
)

func init() {
	srcFmt.Self = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *srcFormatting) onPreInit() {
	srcFmt.knownFormatters = []*z.Tool{
		tools.hindent, tools.stylishhaskell, tools.brittany,
	}
}

func (me *srcFormatting) KnownFormatters() []*z.Tool {
	return me.knownFormatters
}
