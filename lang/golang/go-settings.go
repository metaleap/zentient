package zgo

import (
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

func (me *goSettings) onPreInit() {
	me.cfgGuruScopeExcl = &z.Setting{ID: "cfgGuruScopeExcl", ValDef: "", Title: "Guru Exclusions", Desc: "Package patterns (`github.com/foo/...`) to always exclude from guru `-scope`, space-delimited"}
	me.allSettings = []*z.Setting{me.cfgGuruScopeExcl}
}

func (me *goSettings) KnownSettings() z.Settings {
	return me.allSettings
}
