package zat

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title, l.Enabled = "atmo", "atmo", true
}

func OnPostInit() {
}
