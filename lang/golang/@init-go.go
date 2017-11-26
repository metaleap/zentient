package zgo

import (
	"github.com/metaleap/zentient"
)

var (
	CodeFmt CodeFormatting
)

func OnPreInit() {
	z.Lang.CodeFmt = &CodeFmt
}

func OnPostInit() {
}
