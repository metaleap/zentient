package zgo

import (
	"github.com/metaleap/zentient"
)

var (
	pages goPages
)

func init() {
	pages.Impl, z.Lang.Pages = &pages, &pages
}

type goPages struct {
	z.PagesBase
}
