package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

type CodeFormatting struct {
	z.CodeFormattingBase

	knownFormatters []string
}

var (
	CodeFmt CodeFormatting
)

func init() {
	CodeFmt.Self = &CodeFmt
	z.Lang.CodeFmt = &CodeFmt
	CodeFmt.knownFormatters = []string{"gofmt", "goimports", "gorename"}
}

func (me *CodeFormatting) IsInstalled(formatter string) bool {
	switch formatter {
	case "gofmt":
		return udevgo.Has_gofmt
	case "goimports":
		return udevgo.Has_goimports
	case "gorename":
		return udevgo.Has_gorename
	default:
		z.Bad("formatter", formatter)
	}
	return false
}

func (me *CodeFormatting) KnownFormatters() []string {
	return me.knownFormatters
}
