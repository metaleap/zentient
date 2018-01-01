package zgo

import (
	"strings"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var settings goSettings

func init() {
	settings.Impl, z.Lang.Settings = &settings, &settings
}

type goSettings struct {
	z.SettingsBase

	allSettings      z.Settings
	cfgGuruScopeExcl *z.Setting
	cfgGuruScopeMin  *z.Setting
}

func (me *goSettings) onChangedGuruScopeExcl(oldVal interface{}) {
	if oldval, _ := oldVal.([]string); len(oldval) > 0 {
		for _, oldpat := range oldval {
			udevgo.GuruScopeExclPkgs[oldpat] = false
			delete(udevgo.GuruScopeExclPkgs, oldpat)
		}
	}
	if newval, _ := me.cfgGuruScopeExcl.ValCfg.([]string); len(newval) > 0 {
		for _, pat := range newval {
			udevgo.GuruScopeExclPkgs[pat] = true
		}
	}
}

func (me *goSettings) onReloadedGuruScopeExcl() {
	me.onChangedGuruScopeExcl(nil)
}

/*
github.com/golamb/... github.com/capnproto/... github.com/robertkrimen/... github.com/metaleap/go-opengl/... github.com/metro-cloud-opc/... github.com/arangodb/... github.com/waigani/... github.com/metaleap/go-misctools/... github.com/remyoudompheng/... github.com/jackc/... github.com/metaleap/go-geo-names/... fake.git.metrosystems.net/... github.com/go-forks/... github.com/coffeemug/... github.com/golang/dep/... github.com/sirupsen/... labix.org/... sourcegraph.com/... github.com/juju/... github.com/metaleap/go-util/... github.com/glycerine/... github.com/metaleap/gonad-coreimp/...
*/

func (*goSettings) onChangingGuruScopeExcl(newVal interface{}) {
	if patterns := newVal.([]string); udevgo.PkgsByImP == nil {
		panic(_PKG_NOT_READY_MSG)
	} else {
		for _, pat := range patterns {
			if pkg := udevgo.PkgsByImP[pat]; pkg == nil {
				if !strings.HasSuffix(pat, "/...") {
					z.BadPanic("guru `-scope` exclusion pattern (no `/...` pattern and no such import-path exists) — ", pat)
				}
				var found bool
				pref, self := pat[:len(pat)-3], pat[:len(pat)-4]
				for _, pkg = range udevgo.PkgsByImP {
					if found = strings.HasPrefix(pkg.ImportPath, pref) || pkg.ImportPath == self; found {
						break
					}
				}
				if !found {
					z.BadPanic("guru `-scope` exclusion pattern (no existing import-path matches `/...` pattern) — ", pat)
				}
			}
		}
	}
	return
}

func (me *goSettings) onPreInit() {
	me.cfgGuruScopeExcl = &z.Setting{Id: "cfgGuruScopeExcl", ValDef: []string{}, Title: "Guru: Scopes Exclusions", Desc: "Package patterns (`some/pkg/path/...`) to always exclude from guru `-scope`, space-delimited"}
	me.cfgGuruScopeExcl.OnChanging, me.cfgGuruScopeExcl.OnChanged, me.cfgGuruScopeExcl.OnReloaded = me.onChangingGuruScopeExcl, me.onChangedGuruScopeExcl, me.onReloadedGuruScopeExcl
	me.cfgGuruScopeMin = &z.Setting{Id: "cfgGuruScopeMin", ValDef: false, Title: "Guru: Minimal Scopes", Desc: "If `true`, CodeIntel queries scope to current-and-subordinate packages instead of workspace"}
	me.allSettings = []*z.Setting{me.cfgGuruScopeMin, me.cfgGuruScopeExcl}
}

func (me *goSettings) KnownSettings() z.Settings {
	return append(me.allSettings, me.SettingsBase.KnownSettings()...)
}
