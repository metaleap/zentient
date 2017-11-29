package zhs

import (
	"github.com/metaleap/zentient"
)

type srcFormatting struct {
	z.SrcFormattingBase

	knownFormatters z.Tools
}

var (
	srcFmt srcFormatting
)

func init() {
	srcFmt.Self = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *srcFormatting) onPreInit() {
	srcFmt.knownFormatters = z.Tools{
		tools.hindent, tools.stylishhaskell, tools.brittany,
	}
}

func (me *srcFormatting) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *srcFormatting) RunFormatter(formatter *z.Tool, customProgName string, srcFilePath string, srcFull string) (string, string, error) {
	return "", "", z.Errf("Not yet implemented")
}
