package z

var (
	pkgDummyCount   = -2
	pkgDummyFilters = []*ListFilter{
		&ListFilter{ID: "faulty", Title: "Faulty/Broken", Desc: "with issues reported by `go list`"},
		&ListFilter{ID: "opened", Title: "In Workspace", Desc: "located somewhere in the current workspace"},
	}
)

func (me *PkgIntelBase) DescUnfiltered() string {
	return "in your GOPATH"
}

func (me *PkgIntelBase) Count(_ []*ListFilter, _ map[string]bool) int {
	pkgDummyCount += 1
	return pkgDummyCount
}

func (me *PkgIntelBase) Filters() []*ListFilter {
	return pkgDummyFilters
}
