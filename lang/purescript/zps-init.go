package zps

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title = "purescript", "PureScript"
}

func OnPostInit() {
}
