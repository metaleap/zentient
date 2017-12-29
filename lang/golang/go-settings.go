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
}

func (me *goSettings) onChangedGuruScopeExcl(oldVal interface{}) {
	mod, guruscopeexclpkgs := false, udevgo.GuruScopeExclPkgs
	if oldval, _ := oldVal.([]string); len(oldval) > 0 {
		mod = true
		for _, oldpat := range oldval {
			guruscopeexclpkgs[oldpat] = false
			delete(guruscopeexclpkgs, oldpat)
		}
	}
	if newval, _ := me.cfgGuruScopeExcl.ValCfg.([]string); len(newval) > 0 {
		for _, pat := range newval {
			mod, guruscopeexclpkgs[pat] = true, true
		}
	}
	if mod {
		udevgo.GuruScopeExclPkgs = guruscopeexclpkgs
	}
}

func (me *goSettings) onReloadedGuruScopeExcl() {
	me.onChangedGuruScopeExcl(nil)
}

func (*goSettings) onChangingGuruScopeExcl(newVal interface{}) {
	if patterns := newVal.([]string); udevgo.PkgsByImP == nil {
		panic("PackageTracker not yet live — try again in a few seconds.")
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
	me.cfgGuruScopeExcl = &z.Setting{Id: "cfgGuruScopeExcl", ValDef: []string{}, Title: "Guru Exclusions", Desc: "Package patterns (`some/pkg/path/...`) to always exclude from guru `-scope`, space-delimited"}
	me.cfgGuruScopeExcl.OnChanging, me.cfgGuruScopeExcl.OnChanged, me.cfgGuruScopeExcl.OnReloaded = me.onChangingGuruScopeExcl, me.onChangedGuruScopeExcl, me.onReloadedGuruScopeExcl
	me.allSettings = []*z.Setting{me.cfgGuruScopeExcl}
}

func (me *goSettings) KnownSettings() z.Settings {
	return me.allSettings
}
