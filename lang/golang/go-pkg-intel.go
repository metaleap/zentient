package zgo

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
)

const _PKG_NOT_READY_MSG = "Package Tracker not yet ready — try again in a few seconds."

var pkgIntel goPkgIntel

func init() {
	pkgIntel.Impl, z.Lang.PkgIntel = &pkgIntel, &pkgIntel

	dynfilter := func(id string, title string) *z.ListFilter {
		return &z.ListFilter{ID: id, Pred: pkgIntel.isPkgNope, Title: title, OnSrcLens: pkgIntel.onSrcLens}
	}
	pkgIntel.listFilterSelf = dynfilter("self", "Current & Ancillary")
	pkgIntel.listFilterImpD = dynfilter("impd", "Direct Dependants")
	pkgIntel.listFilterImpI = dynfilter("impi", "Direct & Indirect Dependants")
	pkgIntel.listFilterDepD = dynfilter("depd", "Direct Dependencies")
	pkgIntel.listFilterDepI = dynfilter("depi", "Direct & Indirect Dependencies")
	pkgIntel.listFilterOpen = &z.ListFilter{ID: "open", Pred: pkgIntel.isPkgOpened, Title: "In Workspace", Desc: "located somewhere in one of the currently opened workspace-folders"}

	pkgIntel.listFilters = []*z.ListFilter{
		pkgIntel.listFilterSelf,
		pkgIntel.listFilterImpD,
		pkgIntel.listFilterImpI,
		pkgIntel.listFilterDepD,
		pkgIntel.listFilterDepI,
		pkgIntel.listFilterOpen,
		{ID: "error", Pred: pkgIntel.isPkgError, Title: "With Errors", Desc: "as reported by `go list`"},
		{ID: "deperr", Pred: pkgIntel.isPkgDepErr, Title: "With Dependency Errors", Desc: "as reported by `go list`"},
		{ID: "command", Pred: pkgIntel.isPkgCommand, Title: "Commands", Desc: "as reported by `go list`"},
		{ID: "binary", Pred: pkgIntel.isPkgBinary, Title: "Binary Only", Desc: "as reported by `go list`"},
		{ID: "incomplete", Pred: pkgIntel.isPkgIncomplete, Title: "Incomplete", Desc: "as reported by `go list`"},
		{ID: "ignoreds", Pred: pkgIntel.isPkgIgnored, Title: "With Ignored Go Files", Desc: "as reported by `go list`"},
		{ID: "invalids", Pred: pkgIntel.isPkgInvalid, Title: "With Invalid Go Files", Desc: "as reported by `go list`"},
		{ID: "stale", Pred: pkgIntel.isPkgStale, Title: "Stale", Desc: "as reported by `go list`"},
		{ID: "standard", Pred: pkgIntel.isPkgStandard, Title: "Standard", Desc: "as reported by `go list`"},
		{ID: "goroot", Pred: pkgIntel.isPkgGoRoot, Title: "In GOROOT", Desc: "as reported by `go list`"},
	}
}

type goPkgIntel struct {
	z.PkgIntelBase

	listFilterSelf *z.ListFilter
	listFilterImpD *z.ListFilter
	listFilterImpI *z.ListFilter
	listFilterDepD *z.ListFilter
	listFilterDepI *z.ListFilter
	listFilterOpen *z.ListFilter
	listFilters    []*z.ListFilter
}

func (goPkgIntel) isPkgNope(pkg z.IListItem) bool       { return false }
func (goPkgIntel) isPkgGoRoot(pkg z.IListItem) bool     { return pkg.(*udevgo.Pkg).Goroot }
func (goPkgIntel) isPkgBinary(pkg z.IListItem) bool     { return pkg.(*udevgo.Pkg).BinaryOnly }
func (goPkgIntel) isPkgCommand(pkg z.IListItem) bool    { return pkg.(*udevgo.Pkg).IsCommand() }
func (goPkgIntel) isPkgStandard(pkg z.IListItem) bool   { return pkg.(*udevgo.Pkg).Standard }
func (goPkgIntel) isPkgIncomplete(pkg z.IListItem) bool { return pkg.(*udevgo.Pkg).Incomplete }
func (goPkgIntel) isPkgStale(pkg z.IListItem) bool      { return pkg.(*udevgo.Pkg).Stale }
func (goPkgIntel) isPkgError(pkg z.IListItem) bool      { return pkg.(*udevgo.Pkg).Error != nil }
func (goPkgIntel) isPkgDepErr(pkg z.IListItem) bool     { return len(pkg.(*udevgo.Pkg).DepsErrors) > 0 }
func (goPkgIntel) isPkgIgnored(pkg z.IListItem) bool    { return len(pkg.(*udevgo.Pkg).IgnoredGoFiles) > 0 }
func (goPkgIntel) isPkgInvalid(pkg z.IListItem) bool    { return len(pkg.(*udevgo.Pkg).InvalidGoFiles) > 0 }
func (*goPkgIntel) isPkgOpened(pkg z.IListItem) bool {
	p := pkg.(*udevgo.Pkg)
	for dirpath := range workspace.Dirs() {
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
	curpkgdesc, isdepd, isdepi, isimpd, isimpi, isself := "?", lf == me.listFilterDepD, lf == me.listFilterDepI, lf == me.listFilterImpD, lf == me.listFilterImpI, lf == me.listFilterSelf
	lf.Pred, lf.Desc = me.isPkgNope, "?"

	if srcLens != nil && srcLens.FilePath != "" {
		if curpkg := udevgo.PkgsByDir[filepath.Dir(srcLens.FilePath)]; curpkg != nil {
			curpkgdesc = curpkg.ImportPath
			if isdepd {
				lf.Pred = func(p z.IListItem) bool {
					return ustr.In(p.(*udevgo.Pkg).ImportPath, curpkg.Imports...)
				}
			} else if isdepi {
				lf.Pred = func(p z.IListItem) bool {
					return ustr.In(p.(*udevgo.Pkg).ImportPath, curpkg.Deps...)
				}
			} else if isimpd { // d meaning 'direct' (vs. 'indirect') --- not 'dependant' (vs 'importer')
				lf.Pred = func(p z.IListItem) bool {
					return ustr.In(p.(*udevgo.Pkg).ImportPath, curpkg.Importers()...)
				}
			} else if isimpi { // i meaning 'indirect' (vs. 'direct') --- not 'importer' (vs 'dependant')
				lf.Pred = func(p z.IListItem) bool {
					return ustr.In(p.(*udevgo.Pkg).ImportPath, curpkg.Dependants()...)
				}
			} else if isself {
				lf.Pred = func(p z.IListItem) (iscurpkg bool) {
					pkg := p.(*udevgo.Pkg)
					imppath := pkg.ImportPath
					if iscurpkg = imppath == curpkg.ImportPath || strings.HasPrefix(imppath, curpkg.ImportPath+"/"); iscurpkg {
						pkg.CountLoC()
					}
					return
				}
			}
		}
	}

	if isdepd {
		lf.Desc = "explicitly imported by `"
	} else if isdepi {
		lf.Desc = "explicitly or implicitly imported by `"
	} else if isimpd {
		lf.Desc = "that directly import `"
	} else if isimpi {
		lf.Desc = "that directly or indirectly depend on `"
	} else if isself {
		lf.Desc = "that appear to belong to `"
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

func (me *goPkgIntel) ListItemToMenuItem(p z.IListItem) (item *z.MenuItem) {
	descsmighthavepaths := false
	if pkg, _ := p.(*udevgo.Pkg); pkg != nil {
		delim, hints := " · ", []string{}
		if pkg.ApproxLoC > 0 {
			hints = append(hints, z.Strf("~%d LoC", pkg.ApproxLoC))
		}
		item = &z.MenuItem{Category: pkg.Name, Desc: pkg.Doc, Title: pkg.ImportPath}
		if item.Category == "" {
			item.Category = "  "
		}
		if len(pkg.Errs) > 0 {
			item.Desc, descsmighthavepaths = "", true
			for _, e := range pkg.Errs {
				item.Desc += delim + e.Msg
			}
		} else if pkg.Error != nil {
			item.Desc, descsmighthavepaths = pkg.Error.Err, true
		} else if len(pkg.DepsErrors) > 0 {
			item.Desc, descsmighthavepaths = "", true
			for _, e := range pkg.DepsErrors {
				item.Desc += delim + e.Err
			}
		}
		if item.Desc == "" {
			item.Desc, descsmighthavepaths = pkg.StaleReason, true
		} else if strings.HasPrefix(item.Desc, delim) {
			item.Desc = item.Desc[len(delim):]
		} else if pref := "Package " + pkg.Name + " "; strings.HasPrefix(item.Desc, pref) {
			item.Desc, descsmighthavepaths = item.Desc[len(pref):], true
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
				suffix, descsmighthavepaths = "", true
			}
			hints = append(hints, "Stale"+suffix)
		}
		ª := func(f z.ListItemPredicate) *z.ListItemPredicate { return &f }
		for f, s := range map[*z.ListItemPredicate]string{
			ª(me.isPkgBinary):     "Binary",
			ª(me.isPkgCommand):    "Command",
			ª(me.isPkgIncomplete): "Incomplete",
			ª(me.isPkgStandard):   "Standard",
			ª(me.isPkgGoRoot):     "In GOROOT",
			ª(me.isPkgOpened):     "In Workspace",
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
	if descsmighthavepaths {
		item.Desc = z.PrettifyPathsIn(item.Desc)
	}
	return
}

func (me *goPkgIntel) Filters() []*z.ListFilter {
	return pkgIntel.listFilters
}

func (me *goPkgIntel) ObjSnap(pkgDir string) interface{} {
	if udevgo.PkgsByDir != nil {
		if pkg := udevgo.PkgsByDir[pkgDir]; pkg != nil {
			pkg.CountLoC()
			return pkg
		}
	}
	return nil
}
