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
				pref := pat[:len(pat)-3]
				for _, pkg = range udevgo.PkgsByImP {
					if found = strings.HasPrefix(pkg.ImportPath, pref); found {
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
	me.cfgGuruScopeExcl = &z.Setting{Id: "cfgGuruScopeExcl", ValDef: []string{}, Title: "Guru Exclusions", Desc: "Package patterns (`github.com/foo/...`) to always exclude from guru `-scope`, space-delimited"}
	me.cfgGuruScopeExcl.OnChanging = me.onChangingGuruScopeExcl
	me.allSettings = []*z.Setting{me.cfgGuruScopeExcl}
}

func (me *goSettings) KnownSettings() z.Settings {
	return me.allSettings
}
