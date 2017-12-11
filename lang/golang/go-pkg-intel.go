package zgo

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/zentient"
)

var pkgIntel goPkgIntel

func init() {
	pkgIntel.Impl, z.Lang.PkgIntel = &pkgIntel, &pkgIntel
	pkgIntel.listFilterImps = &z.ListFilter{ID: "imps", Pred: pkgIntel.isPkgNope, Title: "Dependants", OnSrcLens: pkgIntel.onSrcLens}
	pkgIntel.listFilterDeps = &z.ListFilter{ID: "deps", Pred: pkgIntel.isPkgNope, Title: "Dependencies", OnSrcLens: pkgIntel.onSrcLens}
	pkgIntel.listFilterOpen = &z.ListFilter{ID: "open", Pred: pkgIntel.isPkgOpened, Title: "In Workspace", Desc: "located somewhere in the current workspace"}
	pkgIntel.listFilters = []*z.ListFilter{
		pkgIntel.listFilterImps,
		pkgIntel.listFilterDeps,
		pkgIntel.listFilterOpen,
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
}

type goPkgIntel struct {
	z.PkgIntelBase

	listFilterImps *z.ListFilter
	listFilterDeps *z.ListFilter
	listFilterOpen *z.ListFilter
	listFilters    []*z.ListFilter
}

func (goPkgIntel) isPkgNope(pkg z.ListItem) bool       { return false }
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
func (*goPkgIntel) isPkgOpened(pkg z.ListItem) bool {
	p := pkg.(*udevgo.Pkg)
	for dirpath, _ := range workspace.OpenDirs {
		if p.Dir == dirpath || strings.HasPrefix(p.Dir, strings.TrimRight(dirpath, "/\\")+string(filepath.Separator)) {
			return true
		}
	}
	return false
}

func (me *goPkgIntel) UnfilteredDesc() string {
	return "in your GOPATH"
}

func (me *goPkgIntel) onSrcLens(lf *z.ListFilter, srcLens *z.SrcLens) {
	curpkgdesc, isdeps, isimps := "?", lf == me.listFilterDeps, lf == me.listFilterImps
	lf.Pred, lf.Desc = me.isPkgNope, "?"

	if srcLens != nil && srcLens.FilePath != "" {
		if pkg := udevgo.PkgsByDir[filepath.Dir(srcLens.FilePath)]; pkg != nil {
			curpkgdesc = pkg.ImportPath
			if isdeps {
				lf.Pred = func(p z.ListItem) bool {
					return uslice.StrHas(pkg.Deps, p.(*udevgo.Pkg).ImportPath)
				}
			} else if isimps {
				lf.Pred = func(p z.ListItem) bool {
					importers := pkg.Dependants()
					return uslice.StrHas(importers, p.(*udevgo.Pkg).ImportPath)
				}
			}
		}
	}

	if isdeps {
		lf.Desc = "imported by `"
	} else if isimps {
		lf.Desc = "that import `"
	}
	lf.Desc += curpkgdesc + "`"
}

func (me *goPkgIntel) Count(filters z.ListFilters) (count int) {
	count = -1
	me.list(filters, &count)
	return
}

func (me *goPkgIntel) list(filters z.ListFilters, count *int) (results z.ListItems) {
	if udevgo.PkgsByDir != nil {
		if count != nil {
			*count = 0
		}
		for _, pkg := range udevgo.PkgsByDir {
			if pkg != nil {
				allpredicatesmatch := true
				if filters != nil {
					for filter, desired := range filters {
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

func (me *goPkgIntel) List(filters z.ListFilters) (results z.ListItems) {
	return me.list(filters, nil)
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
		if pkgtarget := z.Lang.Workspace.PrettyPath(pkg.Target); me.isPkgCommand(pkg) && pkgtarget != "" {
			if hint := "Target: " + pkgtarget; item.Desc == "" {
				item.Desc = hint
			} else {
				hints = append(hints, hint)
			}
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
		item.IpcID = z.IPCID_OBJ_SNAPSHOT
		item.IpcArgs = me.ObjSnapPrefix() + pkg.Dir
	}
	return
}

func (me *goPkgIntel) Filters() []*z.ListFilter {
	return pkgIntel.listFilters
}

func (me *goPkgIntel) ObjSnap(pkgDir string) interface{} {
	if udevgo.PkgsByDir != nil {
		return udevgo.PkgsByDir[pkgDir]
	}
	return nil
}
