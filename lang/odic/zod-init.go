package zod

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() {
	l := &z.Lang
	l.ID, l.Title = "odic", "odic"
}

func OnPostInit() {
}
