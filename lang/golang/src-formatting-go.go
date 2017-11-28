package zgo

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
	me.knownFormatters = []*z.Tool{
		tools.gofmt, tools.goimports,
	}
}

func (me *srcFormatting) onPostInit() {
	if z.Prog.Cfg.FormatterName == "" && tools.gofmt.Installed {
		z.Prog.Cfg.FormatterName = "gofmt"
	}
}

func (me *srcFormatting) KnownFormatters() []*z.Tool {
	return me.knownFormatters
}
