package zgo

import (
	"sort"
	"strings"

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
		&z.ListFilter{ID: "ignoreds", Pred: pkgIntel.isPkgIgnored, Title: "With Ignored Go Files", Desc: "as reported by `go list`"},
		&z.ListFilter{ID: "invalids", Pred: pkgIntel.isPkgInvalid, Title: "With Invalid Go Files", Desc: "as reported by `go list`"},
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

func (goPkgIntel) isPkgGoRoot(pkg z.ListItem) bool     { return pkg.(*udevgo.Pkg).Goroot }
func (goPkgIntel) isPkgBinary(pkg z.ListItem) bool     { return pkg.(*udevgo.Pkg).BinaryOnly }
func (goPkgIntel) isPkgCommand(pkg z.ListItem) bool    { return pkg.(*udevgo.Pkg).IsCommand() }
func (goPkgIntel) isPkgStandard(pkg z.ListItem) bool   { return pkg.(*udevgo.Pkg).Standard }
func (goPkgIntel) isPkgStale(pkg z.ListItem) bool      { return pkg.(*udevgo.Pkg).Stale }
func (goPkgIntel) isPkgIncomplete(pkg z.ListItem) bool { return pkg.(*udevgo.Pkg).Incomplete }
func (goPkgIntel) isPkgError(pkg z.ListItem) bool      { return pkg.(*udevgo.Pkg).Error != nil }
func (goPkgIntel) isPkgDepErr(pkg z.ListItem) bool     { return len(pkg.(*udevgo.Pkg).DepsErrors) > 0 }
func (goPkgIntel) isPkgIgnored(pkg z.ListItem) bool    { return len(pkg.(*udevgo.Pkg).IgnoredGoFiles) > 0 }
func (goPkgIntel) isPkgInvalid(pkg z.ListItem) bool    { return len(pkg.(*udevgo.Pkg).InvalidGoFiles) > 0 }
func (*goPkgIntel) isPkgOpened(pkg z.ListItem) bool    { return false }

func (me *goPkgIntel) UnfilteredDesc() string {
	return "in your GOPATH"
}

func (me *goPkgIntel) Count(allFilters z.ListFilters) (count int) {
	count = -1
	me.list(allFilters, &count)
	return
}

func (me *goPkgIntel) list(allFilters z.ListFilters, count *int) (results z.ListItems) {
	if udevgo.PkgsByDir != nil {
		if count != nil {
			*count = 0
		}
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
					if count == nil {
						results = append(results, pkg)
					} else {
						*count = (*count) + 1
					}
				}
			}
		}
		sort.Sort(results)
	}
	return
}

func (me *goPkgIntel) List(allFilters z.ListFilters) (results z.ListItems) {
	return me.list(allFilters, nil)
}

func (me *goPkgIntel) ListItemToMenuItem(p z.ListItem) (item *z.MenuItem) {
	if pkg, _ := p.(*udevgo.Pkg); pkg != nil {
		delim, hints := " · ", []string{}
		item = &z.MenuItem{Category: pkg.Name, Desc: pkg.Doc, Title: pkg.ImportPath}
		if item.Category == "" {
			item.Category = "  "
		}
		if len(pkg.Errs) > 0 {
			item.Desc = ""
			for _, e := range pkg.Errs {
				item.Desc += delim + e.Msg
			}
		} else if pkg.Error != nil {
			item.Desc = pkg.Error.Err
		} else if len(pkg.DepsErrors) > 0 {
			item.Desc = ""
			for _, e := range pkg.DepsErrors {
				item.Desc += delim + e.Err
			}
		}
		if item.Desc == "" {
			item.Desc = pkg.StaleReason
		} else if strings.HasPrefix(item.Desc, delim) {
			item.Desc = item.Desc[len(delim):]
		} else if pref := "Package " + pkg.Name + " "; strings.HasPrefix(item.Desc, pref) {
			item.Desc = item.Desc[len(pref):]
		}

		if suffix := ": " + pkg.StaleReason; me.isPkgStale(pkg) {
			if item.Desc == pkg.StaleReason {
				suffix = ""
			}
			hints = append(hints, "Stale"+suffix)
		}
		__ := func(f z.ListItemPredicate) *z.ListItemPredicate { return &f }
		for f, s := range map[*z.ListItemPredicate]string{
			__(me.isPkgBinary):     "Binary",
			__(me.isPkgCommand):    "Command",
			__(me.isPkgIncomplete): "Incomplete",
			__(me.isPkgStandard):   "Standard",
			__(me.isPkgGoRoot):     "In GOROOT",
		} {
			if (*f)(pkg) {
				hints = append(hints, s)
			}
		}
		if me.isPkgError(pkg) {
			if len(pkg.Errs) == 0 {
				hints = append(hints, "Error")
			} else {
				hints = append(hints, z.Strf("%d error(s)", len(pkg.Errs)))
			}
		}
		if l := len(pkg.DepsErrors); l > 0 {
			hints = append(hints, z.Strf("%d dependency error(s)", l))
		}
		if l := len(pkg.IgnoredGoFiles); l > 0 {
			hints = append(hints, z.Strf("%d ignored file(s)", l))
		}
		if l := len(pkg.InvalidGoFiles); l > 0 {
			hints = append(hints, z.Strf("%d invalid file(s)", l))
		}
		item.Hint = strings.Join(hints, delim)
		item.Tag = pkg
	}
	return
}

func (me *goPkgIntel) Filters() []*z.ListFilter {
	return pkgIntel.listFilters
}
