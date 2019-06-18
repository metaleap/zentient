package zex

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title = "examplelang", "ExampleLang"
}

func OnPostInit() {
}
