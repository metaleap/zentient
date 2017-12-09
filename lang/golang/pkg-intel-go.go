package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var pkgIntel goPkgIntel

func init() {
	pkgIntel.listFilters = []*z.ListFilter{
		&z.ListFilter{ID: "opened", Pred: pkgIntel.isPkgOpened, Title: "In Workspace", Desc: "located somewhere in the current workspace"},
		&z.ListFilter{ID: "error", Pred: pkgIntel.isPkgError, Title: "With Errors", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "deperr", Pred: pkgIntel.isPkgDepErr, Title: "With Dependency Errors", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "command", Pred: pkgIntel.isPkgCommand, Title: "Commands", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "binary", Pred: pkgIntel.isPkgBinary, Title: "Binary Only", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "incomplete", Pred: pkgIntel.isPkgIncomplete, Title: "Incomplete", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "ignoreds", Pred: pkgIntel.isPkgIgnored, Title: "Ignored Go Files", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "invalids", Pred: pkgIntel.isPkgInvalid, Title: "Invalid Go Files", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "stale", Pred: pkgIntel.isPkgStale, Title: "Stale", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "standard", Pred: pkgIntel.isPkgStandard, Title: "Standard", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "goroot", Pred: pkgIntel.isPkgGoRoot, Title: "In GOROOT", Desc: "as reported by `go list`"},
	}
	z.Lang.PkgIntel, pkgIntel.Impl = &pkgIntel, &pkgIntel
}

type goPkgIntel struct {
	z.PkgIntelBase

	listFilters []*z.ListFilter
}

func (goPkgIntel) isPkgGoRoot(pkg interface{}) bool     { return pkg.(*udevgo.Pkg).Goroot }
func (goPkgIntel) isPkgBinary(pkg interface{}) bool     { return pkg.(*udevgo.Pkg).BinaryOnly }
func (goPkgIntel) isPkgCommand(pkg interface{}) bool    { return pkg.(*udevgo.Pkg).IsCommand() }
func (goPkgIntel) isPkgStandard(pkg interface{}) bool   { return pkg.(*udevgo.Pkg).Standard }
func (goPkgIntel) isPkgStale(pkg interface{}) bool      { return pkg.(*udevgo.Pkg).Stale }
func (goPkgIntel) isPkgIncomplete(pkg interface{}) bool { return pkg.(*udevgo.Pkg).Incomplete }
func (goPkgIntel) isPkgError(pkg interface{}) bool      { return pkg.(*udevgo.Pkg).Error != nil }
func (goPkgIntel) isPkgDepErr(pkg interface{}) bool     { return len(pkg.(*udevgo.Pkg).DepsErrors) > 0 }
func (goPkgIntel) isPkgIgnored(pkg interface{}) bool    { return len(pkg.(*udevgo.Pkg).IgnoredGoFiles) > 0 }
func (goPkgIntel) isPkgInvalid(pkg interface{}) bool    { return len(pkg.(*udevgo.Pkg).InvalidGoFiles) > 0 }
func (*goPkgIntel) isPkgOpened(pkg interface{}) bool    { return false }

func (me *goPkgIntel) UnfilteredDesc() string {
	return "in your GOPATH"
}

func (me *goPkgIntel) Count(allFilters z.ListFilters) int {
	if udevgo.PkgsByDir == nil {
		return -1
	}
	return len(me.Impl.List(allFilters))
}

func (me *goPkgIntel) List(allFilters z.ListFilters) (results []interface{}) {
	if udevgo.PkgsByDir != nil {
		for _, pkg := range udevgo.PkgsByDir {
			if pkg != nil {
				allpredicatesmatch := true
				if allFilters != nil {
					for filter, desired := range allFilters {
						if satisfiesfilter := filter.Pred(pkg); satisfiesfilter != desired {
							allpredicatesmatch = false
							break
						}
					}
				}
				if allpredicatesmatch {
					results = append(results, pkg)
				}
			}
		}
	}
	return
}

func (me *goPkgIntel) Filters() []*z.ListFilter {
	return pkgIntel.listFilters
}
