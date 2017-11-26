package zgo

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID = "go"
	l.Title = "Go"
}

func OnPostInit() {
}
