package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
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
	CodeFmt.knownFormatters = []string{"stylish-haskell", "hindent", "brittany"}
}

func (me *CodeFormatting) IsInstalled(formatter string) bool {
	switch formatter {
	case "hindent":
		return udevhs.Has_hindent
	case "brittany":
		return udevhs.Has_brittany
	case "stylish-haskell":
		return udevhs.Has_stylish_haskell
	default:
		z.Bad("formatter", formatter)
	}
	return false
}

func (me *CodeFormatting) KnownFormatters() []string {
	return me.knownFormatters
}
