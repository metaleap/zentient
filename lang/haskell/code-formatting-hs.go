package zhs

import (
	"github.com/metaleap/zentient"
)

type CodeFormatting struct {
	z.CodeFormattingBase
}

var (
	CodeFmt CodeFormatting
)

func init() {
	z.Lang.CodeFmt = &CodeFmt
}
