package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient"
)

type srcFormatting struct {
	z.SrcFormattingBase

	knownFormatters []*z.SrcFormatterDesc
}

var (
	srcFmt srcFormatting
)

func init() {
	srcFmt.Self = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *srcFormatting) onPreInit() {
	srcFmt.knownFormatters = []*z.SrcFormatterDesc{
		&z.SrcFormatterDesc{Name: "stylish-haskell", Link: "http://github.com/jaspervdj/stylish-haskell#readme", Installed: udevhs.Has_stylish_haskell},
		&z.SrcFormatterDesc{Name: "hindent", Link: "http://github.com/commercialhaskell/hindent#readme", Installed: udevhs.Has_hindent},
		&z.SrcFormatterDesc{Name: "brittany", Link: "http://github.com/lspitzner/brittany#readme", Installed: udevhs.Has_brittany},
	}
}

func (me *srcFormatting) KnownFormatters() []*z.SrcFormatterDesc {
	return me.knownFormatters
}
