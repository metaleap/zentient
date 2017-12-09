package zgo

import (
	"github.com/metaleap/zentient"
)

var pkgIntel goPkgIntel

func init() {
	pkgIntel.listFilters = []*z.ListFilter{
		&z.ListFilter{ID: "faulty", Title: "Faulty/Broken", Desc: "with issues reported by `go list`"},
		&z.ListFilter{ID: "opened", Title: "In Workspace", Desc: "located somewhere in the current workspace"},
	}
	z.Lang.PkgIntel, pkgIntel.Impl = &pkgIntel, &pkgIntel
}

type goPkgIntel struct {
	z.PkgIntelBase

	listFilters []*z.ListFilter
}

func (me *goPkgIntel) DescUnfiltered() string {
	return "in your GOPATH"
}

func (me *goPkgIntel) Count(_ []*z.ListFilter, _ map[string]bool) int {
	return 0
}

func (me *goPkgIntel) Filters() []*z.ListFilter {
	return pkgIntel.listFilters
}
