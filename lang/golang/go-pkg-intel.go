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
					return uslice.StrHas(curpkg.Imports, p.(*udevgo.Pkg).ImportPath)
				}
			} else if isdepi {
				lf.Pred = func(p z.IListItem) bool {
					return uslice.StrHas(curpkg.Deps, p.(*udevgo.Pkg).ImportPath)
				}
			} else if isimpd { // d meaning 'direct' (vs. 'indirect') --- not 'dependant' (vs 'importer')
				lf.Pred = func(p z.IListItem) bool {
					return uslice.StrHas(curpkg.Importers(), p.(*udevgo.Pkg).ImportPath)
				}
			} else if isimpi { // i meaning 'indirect' (vs. 'direct') --- not 'importer' (vs 'dependant')
				lf.Pred = func(p z.IListItem) bool {
					return uslice.StrHas(curpkg.Dependants(), p.(*udevgo.Pkg).ImportPath)
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
		if pkg.LoC > 0 {
			hints = append(hints, z.Strf("%d LoC", pkg.LoC))
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
		if pkgDir == "tmptest" {
			m := map[string]*udevgo.Pkg{}
			for _, p := range []string{"github.com/golamb/test-pscoreimp2go/Data/Lens/Iso/Newtype", "github.com/golamb/test-pscoreimp2go/Data/Exists", "github.com/golamb/test-pscoreimp2go", "github.com/golamb/test-pscoreimp2go/Data/List", "github.com/golamb/test-pscoreimp2go/Mini/TypeClass", "github.com/golamb/test-pscoreimp2go/Data/EuclideanRing", "github.com/golamb/test-pscoreimp2go/Control/Monad/RWS", "github.com/golamb/test-pscoreimp2go/Data/Identity", "github.com/golamb/test-pscoreimp2go/Mini/DataTypes", "github.com/golamb/test-pscoreimp2go/Data/CommutativeRing", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep", "github.com/go-forks/pgx", "github.com/golamb/test-pscoreimp2go/Optic/Types", "github.com/golamb/test-pscoreimp2go/Data/Functor/Product", "github.com/golamb/test-pscoreimp2go/Mini/TClsImpl3", "github.com/golamb/test-pscoreimp2go/Data/ArrayBuffer/Types", "github.com/golamb/test-pscoreimp2go/Data/Foreign/Index", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Store", "github.com/robertkrimen/godocdown", "fake.git.metrosystems.net/get-it-together/developer-assessment-support/GO/sample2/storage", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Clown", "github.com/golamb/test-pscoreimp2go/Data/Divisible", "github.com/golamb/test-pscoreimp2go/Data/Foreign/Lens", "github.com/golamb/test-pscoreimp2go/Control/Monad", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Alternate", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Wander", "github.com/golamb/test-pscoreimp2go/Optic/Internal/Prism", "github.com/golamb/test-pscoreimp2go/Control/Monad/Writer/Class", "github.com/golamb/test-pscoreimp2go/Data/Lens/Prism/Coproduct", "github.com/golamb/test-pscoreimp2go/Data/Lens/Lens/Tuple", "github.com/golamb/test-pscoreimp2go/Type/Equality", "github.com/golamb/test-pscoreimp2go/Control/Bind", "github.com/golamb/test-pscoreimp2go/Data/Enum", "github.com/capnproto/go-capnproto2/capnpc-go", "github.com/golamb/test-pscoreimp2go/Data/Unfoldable", "github.com/golamb/test-pscoreimp2go/Data/List/Lazy/Types", "github.com/golamb/test-pscoreimp2go/Optic/Getter", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Console", "github.com/go-forks/pgx/pgtype/testutil", "github.com/golamb/test-pscoreimp2go/Graphics/Canvas", "github.com/golamb/test-pscoreimp2go/Data/Lens/Prism/Maybe", "github.com/golamb/test-pscoreimp2go/Data/Distributive", "github.com/golamb/test-pscoreimp2go/Control/Monad/Gen", "github.com/golamb/test-pscoreimp2go/Control/Monad/ST", "github.com/capnproto/go-capnproto2/schemas", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Clown", "github.com/golamb/test-pscoreimp2go/Data/Lens/Indexed", "github.com/jackc/pgx/pgtype/ext/shopspring-numeric", "github.com/capnproto/go-capnproto2/rpc/internal/logtransport", "github.com/golamb/test-pscoreimp2go/Data/Foreign/Keys", "github.com/golamb/test-pscoreimp2go/Data/Map/Gen", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Trans/Class", "github.com/golamb/test-pscoreimp2go/Control/Monad/Cont/Class", "github.com/golamb/test-pscoreimp2go/Ohai/Greet", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Uncurried", "github.com/golamb/test-pscoreimp2go/Data/String/Regex", "github.com/golamb/test-pscoreimp2go/Control/MonadPlus", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Enum", "github.com/golamb/test-pscoreimp2go/Ohai/Howdy", "github.com/golamb/test-pscoreimp2go/Data/Array/ST/Iterator", "github.com/golamb/test-pscoreimp2go/Data/Set", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Indexed", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Costar", "github.com/go-forks/toml/cmd/tomlv", "github.com/golamb/test-pscoreimp2go/Color", "github.com/golamb/test-pscoreimp2go/Data/Lens/Grate", "github.com/golamb/test-pscoreimp2go/Data/Tuple", "github.com/golamb/test-pscoreimp2go/Data/Lens/At", "github.com/golamb/test-pscoreimp2go/Data/DivisionRing", "github.com/golamb/test-pscoreimp2go/Data/Semiring", "github.com/golamb/test-pscoreimp2go/Control/MonadZero", "github.com/golamb/test-pscoreimp2go/Data/Equivalence", "github.com/golamb/test-pscoreimp2go/Color/Scheme/MaterialDesign", "github.com/golamb/test-pscoreimp2go/Control/Monad/Reader", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Traversals", "github.com/golamb/test-pscoreimp2go/Control/Monad/Writer/Trans", "github.com/golamb/test-pscoreimp2go/Control/Lazy", "github.com/metro-cloud-opc/onb-poll", "github.com/golamb/test-pscoreimp2go/Control/Monad/Except", "github.com/golamb/test-pscoreimp2go/Data/Array/ST/Partial", "github.com/golamb/test-pscoreimp2go/Data/Ord/Unsafe", "github.com/go-forks/toml/cmd/toml-test-encoder", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Bounded", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Store/Class", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Strong", "github.com/golamb/test-pscoreimp2go/Mini/NewTypes", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/String", "github.com/golamb/test-pscoreimp2go/Type/Data/Boolean", "github.com/golamb/test-pscoreimp2go/Data/Field", "github.com/golamb/test-pscoreimp2go/Data/Functor/App", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Shop", "github.com/golamb/test-pscoreimp2go/Data/FoldableWithIndex", "github.com/golamb/test-pscoreimp2go/Data/Unit", "github.com/golamb/test-pscoreimp2go/Data/Lens/Index", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Cowrap", "github.com/metaleap/go-opengl/cmd/glfw3-minimal-app", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Timer", "github.com/golamb/test-pscoreimp2go/Data/BooleanAlgebra", "github.com/metaleap/go-geo-names/make-mongodb", "github.com/golamb/test-pscoreimp2go/Control/Monad/Gen/Class", "github.com/golamb/test-pscoreimp2go/Color/Scale/Perceptual", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/Pos", "fake.git.metrosystems.net/get-it-together/developer-assessment-support/GO/sample2/storage/fake", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Class", "github.com/golamb/test-pscoreimp2go/Data/Functor/Contravariant", "github.com/golamb/test-pscoreimp2go/Data/Record/Builder", "github.com/golamb/test-pscoreimp2go/Data/These", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Env/Class", "github.com/golamb/test-pscoreimp2go/Data/Enum/Gen", "github.com/golamb/test-pscoreimp2go/Mini/TypeAliases", "github.com/golamb/test-pscoreimp2go/Type/Row/Effect/Equality", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Ord", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser", "github.com/golamb/test-pscoreimp2go/Data/Lens/Fold", "github.com/golamb/test-pscoreimp2go/Data/List/Lazy", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Unsafe", "github.com/golamb/test-pscoreimp2go/Data/Lens/Fold/Partial", "github.com/capnproto/go-capnproto2/internal/fulfiller", "github.com/golamb/test-pscoreimp2go/Graphics/Drawing", "github.com/golamb/test-pscoreimp2go/Data/StrMap/Gen", "github.com/metaleap/go-opengl/cmd/opengl-minimal-app-glfw3", "github.com/golamb/test-pscoreimp2go/Control/Biapplicative", "github.com/golamb/test-pscoreimp2go/Data/Lens/Prism", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Grating", "github.com/golamb/test-pscoreimp2go/Data/Symbol", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Split", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Forget", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Exception/Unsafe", "github.com/golamb/test-pscoreimp2go/Control/Monad/Cont/Trans", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Costrong", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/Language", "github.com/golamb/test-pscoreimp2go/Control/Extend", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Parser", "github.com/golamb/test-pscoreimp2go/Color/Blending", "github.com/golamb/test-pscoreimp2go/Data/Lens/Prism/Either", "github.com/golamb/test-pscoreimp2go/Data/Array/ST", "github.com/golamb/test-pscoreimp2go/Data/Semigroup/Traversable", "github.com/golamb/da/ffi/ps2go/Data/Eq", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Tagged", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Traced/Class", "github.com/golamb/test-pscoreimp2go/Optic/Setter", "github.com/golamb/test-pscoreimp2go/Optic/Internal/Setter", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Eq", "github.com/golamb/test-pscoreimp2go/Control/Monad/Maybe/Trans", "github.com/golamb/test-pscoreimp2go/Data/Lens/Types", "github.com/jackc/pgx/log/zapadapter", "github.com/golamb/test-pscoreimp2go/Data/Either", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/JCursor/Gen", "github.com/metaleap/go-misctools/go-injecttest", "github.com/golamb/test-pscoreimp2go/Control/Semigroupoid", "github.com/golamb/test-pscoreimp2go/Control/Biapply", "github.com/golamb/test-pscoreimp2go/Data/String/Regex/Flags", "github.com/golamb/test-pscoreimp2go/Data/Divide", "github.com/golamb/test-pscoreimp2go/Data/Record/Unsafe", "github.com/golamb/test-pscoreimp2go/Control/Monad/RWS/Trans", "github.com/golamb/test-pscoreimp2go/Mini/ShowCls", "github.com/golamb/test-pscoreimp2go/Control/Comonad", "github.com/golamb/test-pscoreimp2go/Control/Alt", "github.com/golamb/test-pscoreimp2go/Data/Array", "github.com/remyoudompheng/go-misc/systemd/control", "github.com/metro-cloud-opc/onboarding-fault-detection", "github.com/golamb/test-pscoreimp2go/Data/Comparison", "github.com/golamb/test-pscoreimp2go/Data/Semigroup", "github.com/golamb/test-pscoreimp2go/Data/List/NonEmpty", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Joker", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Join", "github.com/golamb/test-pscoreimp2go/Data/Traversable/Accum", "github.com/golamb/test-pscoreimp2go/Optic/Lens", "github.com/golamb/test-pscoreimp2go/Data/Traversable/Accum/Internal", "github.com/golamb/test-pscoreimp2go/Data/Functor/Coproduct", "github.com/golamb/test-pscoreimp2go/Data/String/Extra", "github.com/golamb/test-pscoreimp2go/Mini/DataType", "github.com/golamb/test-pscoreimp2go/Data/Lens/Traversal", "github.com/golamb/test-pscoreimp2go/Partial/Unsafe", "github.com/golamb/test-pscoreimp2go/Data/Maybe/First", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Traced", "github.com/go-forks/pgx/log/zapadapter", "github.com/golamb/test-pscoreimp2go/Data/Tuple/Nested", "github.com/jackc/pgx/pgtype/testutil", "github.com/golamb/test-pscoreimp2go/Data/Foldable", "github.com/golamb/test-pscoreimp2go/Optic/Prism", "github.com/golamb/test-pscoreimp2go/Data/Lazy", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Conj", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Disj", "github.com/golamb/test-pscoreimp2go/Control/Monad/Writer", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff/Exception", "github.com/metro-cloud-opc/onb-cat", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Wrap", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Decode/Combinators", "github.com/golamb/test-pscoreimp2go/Data/Ordering", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Traced/Trans", "github.com/capnproto/go-capnproto2", "fake.git.metrosystems.net/get-it-together/developer-assessment-support/GO/sample2", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Store/Trans", "github.com/golamb/test-pscoreimp2go/Data/Decide", "github.com/golamb/test-pscoreimp2go/Color/Scale", "github.com/golamb/test-pscoreimp2go/Control/Monad/List/Trans", "github.com/golamb/test-pscoreimp2go/Data/String/CodePoints", "github.com/golamb/test-pscoreimp2go/Data/List/Partial", "github.com/golamb/test-pscoreimp2go/Type/Data/Symbol", "github.com/golamb/test-pscoreimp2go/Type/Row", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Prisms", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/Token", "github.com/golamb/test-pscoreimp2go/Data/String/Regex/Unsafe", "github.com/golamb/test-pscoreimp2go/Control/Monad/State/Class", "github.com/golamb/da/ffi/ps2go/Data/Show", "github.com/go-forks/toml/cmd/toml-test-decoder", "github.com/golamb/test-pscoreimp2go/Data/List/Types", "github.com/golamb/test-pscoreimp2go/Unsafe/Reference", "github.com/golamb/test-pscoreimp2go/Data/Functor", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Market", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Endo", "github.com/golamb/test-pscoreimp2go/Data/Lens/Getter", "github.com/golamb/test-pscoreimp2go/Mini/TCls3", "github.com/golamb/test-pscoreimp2go/Data/String", "github.com/capnproto/go-capnproto2/rpc", "github.com/golamb/test-pscoreimp2go/Data/Maybe", "github.com/golamb/test-pscoreimp2go/Main", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Decode/Class", "github.com/golamb/test-pscoreimp2go/Text/Parsing/StringParser", "github.com/golamb/test-pscoreimp2go/Data/Nullable", "github.com/golamb/test-pscoreimp2go/Data/List/ZipList", "github.com/golamb/test-pscoreimp2go/Data/Decidable", "github.com/golamb/test-pscoreimp2go/Control/Monad/Rec/Class", "github.com/golamb/test-pscoreimp2go/Text/Parsing/StringParser/Expr", "github.com/go-forks/pgx/examples/url_shortener", "github.com/golamb/test-pscoreimp2go/Data/Bitraversable", "github.com/metro-cloud-opc/travis-hook", "github.com/golamb/test-pscoreimp2go/Data/Char/Unicode", "github.com/golamb/test-pscoreimp2go/Data/Char/Unicode/Internal", "github.com/golamb/test-pscoreimp2go/Control/Monad/State/Trans", "github.com/golamb/test-pscoreimp2go/Data/Functor/Invariant", "github.com/coffeemug/lg", "github.com/golamb/test-pscoreimp2go/Control/Apply", "github.com/golamb/test-pscoreimp2go/Data/Bounded", "github.com/golamb/test-pscoreimp2go/Data/Array/Partial", "github.com/golamb/test-pscoreimp2go/Data/Functor/Compose", "github.com/golamb/test-pscoreimp2go/Color/Scheme/Clrs", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Decode/Generic", "github.com/golamb/test-pscoreimp2go/Control/Monad/Gen/Common", "github.com/golamb/test-pscoreimp2go/Data/Show", "github.com/golamb/test-pscoreimp2go/Control/Plus", "github.com/golamb/test-pscoreimp2go/Mini/Semifreakoid", "github.com/golamb/test-pscoreimp2go/Data/Ord", "github.com/golamb/test-pscoreimp2go/Data/Either/Nested", "github.com/golamb/test-pscoreimp2go/Control/Monad/Cont", "github.com/golamb/test-pscoreimp2go/Data/Lens/Record", "github.com/golamb/test-pscoreimp2go/Type/Proxy", "github.com/golamb/test-pscoreimp2go/Control/Monad/Eff", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Decode/Generic/Rep", "github.com/golamb/test-pscoreimp2go/Type/Data/Ordering", "github.com/golamb/test-pscoreimp2go/Control/Monad/Error/Class", "github.com/golamb/test-pscoreimp2go/Data/Eq", "github.com/golamb/test-pscoreimp2go/Control/Applicative", "github.com/capnproto/go-capnproto2/server", "github.com/golamb/test-pscoreimp2go/Data/Function", "github.com/golamb/test-pscoreimp2go/Control/Monad/Except/Trans", "github.com/go-forks/pgx/pgtype/ext/shopspring-numeric", "github.com/golamb/test-pscoreimp2go/Data/Lens/Zoom", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/Combinators", "github.com/golamb/test-pscoreimp2go/Control/Monad/Reader/Trans", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Closed", "github.com/golamb/test-pscoreimp2go/Data/Const", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Join", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Semigroup", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Re", "github.com/metaleap/go-opengl/cmd/gogl-minimal-app-glfw2", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Monoid", "github.com/golamb/test-pscoreimp2go/Data/Lens/Lens", "github.com/golamb/test-pscoreimp2go/Data/List/Lazy/NonEmpty", "github.com/golamb/test-pscoreimp2go/Data/Predicate", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Wrap", "github.com/golamb/test-pscoreimp2go/Data/Semigroup/Foldable", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Zipping", "github.com/golamb/test-pscoreimp2go/Data/NonEmpty", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Focusing", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Cochoice", "github.com/golamb/test-pscoreimp2go/Data/FunctorWithIndex", "github.com/golamb/test-pscoreimp2go/Data/Lens/Lens/Product", "github.com/go-forks/go-diff/diff", "github.com/remyoudompheng/go-misc/pkgsonames", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Env/Trans", "github.com/golamb/test-pscoreimp2go/Data/Bifoldable", "github.com/golamb/test-pscoreimp2go/Data/Map", "github.com/golamb/test-pscoreimp2go/Data/StrMap", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Encode/Generic/Rep", "github.com/golamb/test-pscoreimp2go/Data/Foreign", "github.com/capnproto/go-capnproto2/pogs", "github.com/golamb/test-pscoreimp2go/Data/TraversableWithIndex", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor", "github.com/golamb/test-pscoreimp2go/Sammy", "github.com/golamb/test-pscoreimp2go/Data/Monoid", "github.com/golamb/test-pscoreimp2go/Data/StrMap/ST", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Dual", "github.com/jackc/pgx/examples/url_shortener", "github.com/golamb/test-pscoreimp2go/Data/Functor/Product/Nested", "github.com/golamb/test-pscoreimp2go/Data/Int", "github.com/capnproto/go-capnproto2/encoding/text", "github.com/golamb/test-pscoreimp2go/Data/Void", "github.com/golamb/test-pscoreimp2go/Data/Traversable", "github.com/golamb/test-pscoreimp2go/PSCI/Support", "github.com/metaleap/go-opengl/cmd/gogl-minimal-app-glfw3", "github.com/golamb/test-pscoreimp2go/Data/Profunctor", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Joker", "github.com/golamb/test-pscoreimp2go/Data/Functor/Coproduct/Nested", "github.com/golamb/test-pscoreimp2go/Color/Scheme/Harmonic", "github.com/capnproto/go-capnproto2/internal/nodemap", "github.com/golamb/test-pscoreimp2go/Data/Lens/Lens/Unit", "github.com/golamb/test-pscoreimp2go/Data/Record/ST", "github.com/golamb/test-pscoreimp2go/Text/Parsing/StringParser/String", "github.com/golamb/test-pscoreimp2go/Data/Op", "github.com/golamb/test-pscoreimp2go/Control/Category", "github.com/golamb/test-pscoreimp2go/Control/Monad/Reader/Class", "github.com/metaleap/go-opengl/cmd/opengl-minimal-app-glfw2", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Encode/Class", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/JCursor", "github.com/golamb/test-pscoreimp2go/Color/Scheme/X11", "github.com/golamb/test-pscoreimp2go/Data/String/Gen", "github.com/golamb/test-pscoreimp2go/Data/Char/Gen", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Core", "github.com/golamb/test-pscoreimp2go/Data/Lens/Setter", "github.com/golamb/test-pscoreimp2go/Data/Generic", "github.com/golamb/test-pscoreimp2go/Optic/Core", "github.com/golamb/test-pscoreimp2go/Data/Record", "github.com/golamb/test-pscoreimp2go/Data/Lens/Internal/Exchange", "github.com/golamb/da/ffi/ps2go/Data/Semigroup", "github.com/golamb/test-pscoreimp2go/Text/Parsing/Parser/Expr", "github.com/golamb/test-pscoreimp2go/Graphics/Drawing/Font", "github.com/golamb/test-pscoreimp2go/Control/Comonad/Env", "github.com/golamb/test-pscoreimp2go/Control/Monad/Trans/Class", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Encode/Combinators", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Choice", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Gen", "github.com/golamb/test-pscoreimp2go/Text/Parsing/StringParser/Combinators", "github.com/golamb/test-pscoreimp2go/Control/Alternative", "github.com/golamb/test-pscoreimp2go/Data/Lens/Iso", "github.com/golamb/test-pscoreimp2go/Partial", "github.com/golamb/test-pscoreimp2go/Data/Ring", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Product", "github.com/golamb/test-pscoreimp2go/Mini/ShowImpl", "github.com/golamb/test-pscoreimp2go/Data/HeytingAlgebra", "github.com/golamb/test-pscoreimp2go/Optic/Laws/Lens", "github.com/golamb/test-pscoreimp2go/Data/Newtype", "github.com/metro-cloud-opc/onb-web", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Additive", "github.com/golamb/test-pscoreimp2go/Data/Options", "github.com/golamb/test-pscoreimp2go/Data/Argonaut/Encode/Generic", "github.com/golamb/test-pscoreimp2go/Data/Bifunctor/Flip", "github.com/golamb/test-pscoreimp2go/Data/Profunctor/Star", "github.com/golamb/test-pscoreimp2go/Data/Maybe/Last", "github.com/golamb/test-pscoreimp2go/Control/Monad/State", "github.com/golamb/test-pscoreimp2go/Data/Function/Uncurried", "github.com/golamb/test-pscoreimp2go/Mini/SimpleTypeClass", "github.com/golamb/test-pscoreimp2go/Data/Lens/Lens/Void", "github.com/golamb/test-pscoreimp2go/Data/Monoid/Multiplicative", "github.com/golamb/test-pscoreimp2go/Data/Generic/Rep/Show", "github.com/golamb/test-pscorefn2go/λ/hdgarrood/purescript-sequences/1.0.3/Data/Sequence/Internal", "github.com/golamb/test-pscorefn2go/λ/jutaro/purescript-typedarray/2.1.0/Data/TypedArray", "github.com/golamb/test-pscorefn2go/λ/purescript-contrib/purescript-profunctor-lenses/3.7.0/Data/Lens/Internal/Shop", "github.com/golamb/test-pscorefn2go/λ/nullobject/purescript-pqueue/1.0.0/Data/PQueue", "github.com/golamb/test-pscorefn2go/λ/purescript/purescript-prelude/3.1.0/Control/Apply", "github.com/glycerine/go-capnproto_test", "github.com/golamb/test-pscorefn2go/λ/sharkdp/purescript-quantities/7.0.0/Data/Quantity/Math", "github.com/golamb/test-pscorefn2go/λ/morganthomas/purescript-group/3.1.0/Data/Group", "github.com/golamb/test-pscorefn2go/λ/purescript/purescript-foldable-traversable/3.6.1/Data/FoldableWithIndex", "github.com/golamb/test-pscorefn2go/λ/purescript/purescript-control/3.3.1/Control/Extend", "github.com/golamb/test-pscorefn2go/λ/klangner/purescript-stats/0.2.0/Numeric/SpecFunctions", "github.com/golamb/test-pscorefn2go/λ/sharkdp/purescript-quantities/7.0.0/Data/Units/Time", "github.com/golamb/test-pscorefn2go/λ/purescript/purescript-exceptions/3.1.0/Control/Monad/Eff/Exception/Unsafe", "github.com/golamb/test-pscorefn2go/λ/rightfold/purescript-density-codensity/1.0.2/Control/Monad/Codensity", "github.com/golamb/test-pscorefn2go/λ/purescript/purescript-functors/2.2.0/Data/Functor/Product/Nested"} {
				m[p] = udevgo.PkgsByImP[p]
			}
			return m
		}
		return udevgo.PkgsByDir[pkgDir]
	}
	return nil
}
