package zps

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "purescript"
	l.Title = "PureScript"
}

func OnPostInit() {
}
