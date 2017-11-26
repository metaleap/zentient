package zgo

import (
	"github.com/metaleap/go-util/dev/go"
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
		&z.SrcFormatterDesc{Name: "gofmt", Link: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt},
		&z.SrcFormatterDesc{Name: "goimports", Link: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports},
	}
}

func (me *srcFormatting) KnownFormatters() []*z.SrcFormatterDesc {
	return me.knownFormatters
}
