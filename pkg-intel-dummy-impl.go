package z

var (
	pkgDummyCount   = -2
	pkgDummyFilters = []*PkgIntelFilter{
		&PkgIntelFilter{ID: "faulty", Title: "Faulty/Broken", Desc: "with issues reported by `go list`"},
		&PkgIntelFilter{ID: "opened", Title: "Opened", Desc: "located somewhere in the current workspace"},
	}
)

func (me *PkgIntelBase) ListAllDesc() string {
	return "in your GOPATH"
}

func (me *PkgIntelBase) ListCount(_ []*PkgIntelFilter, _ map[string]bool) int {
	pkgDummyCount += 1
	return pkgDummyCount
}

func (me *PkgIntelBase) ListFilters() []*PkgIntelFilter {
	return pkgDummyFilters
}
