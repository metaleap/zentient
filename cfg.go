package z

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-leap/fs"
	"github.com/go-leap/std"
)

type ISettings interface {
	IMenuItems

	KnownSettings() Settings
}

type Settings []*Setting

func (this Settings) byId(id string) *Setting {
	for _, s := range this {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func (this Settings) numCust() (n int) {
	for _, s := range this {
		if s.ValCfg != nil {
			n++
		}
	}
	return
}

func (this Settings) titles() (all []string) {
	all = make([]string, len(this))
	for i, s := range this {
		all[i] = s.Title
	}
	return
}

type Setting struct {
	Id         string
	Title      string
	Desc       string
	ValCfg     interface{}
	ValDef     interface{}
	OnChanging func(newVal interface{}) `json:"-"`
	OnChanged  func(oldVal interface{}) `json:"-"`
	OnReloaded func()                   `json:"-"`

	menuItem *MenuItem
}

func (this *Setting) Val() interface{} {
	if this.ValCfg != nil {
		return this.ValCfg
	}
	return this.ValDef
}

func (this *Setting) ValBool() (val bool) {
	val, _ = this.Val().(bool)
	return
}

func (this *Setting) ValInt() (val int64) {
	val, _ = this.Val().(int64)
	return
}

func (this *Setting) ValUInt() (val uint64) {
	val, _ = this.Val().(uint64)
	return
}

func (this *Setting) ValStr() (val string) {
	v := this.Val()
	switch vx := v.(type) {
	case string:
		val = vx
	case []string:
		val = strings.Join(vx, " ")
	default:
		val = Strf("%v", vx)
	}
	return
}

func (this *Setting) ValStrs() (val []string) {
	val, _ = this.Val().([]string)
	return
}

type Config struct {
	Internal      map[string]interface{} `json:",omitempty"`
	FormatterName string                 `json:",omitempty"`
	FormatterProg string                 `json:",omitempty"`
	AutoDiags     []string               `json:",omitempty"`

	err error
	// recallFilePath string
	filePath       string
	timeLastLoaded int64
}

func (this *Config) reload() {
	if stale, _ := ufs.IsNewerThanTime(this.filePath, this.timeLastLoaded); stale {
		// 1. re-initialize me
		var empty Config
		*this = empty
		this.filePath = filepath.Join(Prog.Dir.Config, Prog.Name+".config.json")

		// 2. load
		if ufs.IsFile(this.filePath) { // otherwise, it's a fresh setup
			if this.err = ustd.JsonDecodeFromFile(this.filePath, this); this.err == nil {
				this.timeLastLoaded = time.Now().UnixNano()
				if Lang.Settings != nil && this.Internal != nil {
					for _, ks := range Lang.Settings.KnownSettings() {
						if val, ok := this.Internal[ks.Id]; ok {
							switch vx := val.(type) {
							case []interface{}:
								strs := make([]string, len(vx))
								for i := range vx {
									strs[i] = vx[i].(string)
								}
								ks.ValCfg = strs
							case float64:
								switch ks.ValDef.(type) {
								case int64:
									ks.ValCfg = (int64)(vx)
								case uint64:
									ks.ValCfg = (uint64)(vx)
								default:
									ks.ValCfg = vx
								}
							default:
								ks.ValCfg = val
							}
							if ks.OnReloaded != nil {
								ks.OnReloaded()
							}
						}
					}
					this.Internal = nil
				}
			}
		}
	}
}

// func (this *Config) recall() {
// 	this.recallFilePath = filepath.Join(Prog.Dir.Cache, Prog.Name+".recall.json")
// 	if ufs.IsFile(this.recallFilePath) {
// 		ustd.JsonDecodeFromFile(this.recallFilePath, &Prog.recall)
// 	}
// 	if Prog.recall.i64 == nil {
// 		Prog.recall.i64 = map[string]int64{}
// 	}
// }

// func (this *Config) saveRecall() {
// 	ustd.JsonEncodeToFile(&Prog.recall, this.recallFilePath)
// }

func (this *Config) Save() (err error) {
	if Lang.Settings != nil {
		this.Internal = map[string]interface{}{}
		for _, ks := range Lang.Settings.KnownSettings() {
			if ks.ValCfg != nil {
				this.Internal[ks.Id] = ks.ValCfg
			}
		}
		if len(this.Internal) == 0 {
			this.Internal = nil
		}
	}
	err = ustd.JsonEncodeToFile(this, this.filePath)
	this.Internal = nil
	return
}

type SettingsBase struct {
	Impl ISettings

	cmdListAll  *MenuItem
	cmdResetAll *MenuItem
}

func (this *SettingsBase) Init() {
	if Lang.Settings != nil {
		ks := this.Impl.KnownSettings()
		this.cmdListAll = &MenuItem{IpcID: IPCID_CFG_LIST, Title: Strf("%s-Specific", Lang.Title), Hint: Strf("%d setting(s)", len(ks)), Desc: Strf("Customize ➜ %s", strings.Join(ks.titles(), " · "))}
		this.cmdResetAll = &MenuItem{IpcID: IPCID_CFG_RESETALL, Title: "Reset All", Hint: Strf("%s-Specific Settings", Lang.Title)}
		this.cmdResetAll.Confirm = Strf("Are you sure you want to %s %s?", this.cmdResetAll.Title, this.cmdResetAll.Hint)
		for _, s := range ks {
			s.menuItem = &MenuItem{Title: s.Title, Desc: s.Desc, IpcID: IPCID_CFG_SET}
		}
	}
}

func (this *SettingsBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_CFG_LIST:
		this.onListAll(resp.withMenu())
	case IPCID_CFG_SET:
		args := req.IpcArgs.(map[string]interface{})
		this.onSet(args["id"].(string), args["val"].(string), resp.withMenu())
	case IPCID_CFG_RESETALL:
		if num, err := this.onResetAll(); err != nil {
			resp.ErrMsg = err.Error()
		} else {
			resp.withMenu().NoteInfo = Strf("%d customized setting(s) just reset to factory defaults.", num)
		}
	default:
		return false
	}
	return true
}

func (this *SettingsBase) KnownSettings() Settings {
	if Lang.Diag == nil {
		return nil
	}
	return Settings{cfgLintStickiness}
}

func (this *SettingsBase) onSet(cfgId string, cfgVal string, menu *menuResp) {
	info, setting := "changed", this.Impl.KnownSettings().byId(cfgId)
	if setting == nil {
		BadPanic("setting ID", cfgId)
	}
	var err error
	var newval interface{}
	if cfgVal = strings.TrimSpace(cfgVal); cfgVal == "" {
		info = "reset"
	} else {
		switch setting.ValDef.(type) {
		case string:
			newval = cfgVal
		case []string:
			newval = strings.Split(cfgVal, " ")
		case bool:
			newval, err = strconv.ParseBool(cfgVal)
		case int64:
			newval, err = strconv.ParseInt(cfgVal, 10, 64)
		case uint64:
			newval, err = strconv.ParseUint(cfgVal, 10, 64)
		case float64:
			newval, err = strconv.ParseFloat(cfgVal, 64)
		default:
			BadPanic(Strf("setting'%s'.ValDef.(type)", setting.Id), Strf("%T", setting.ValDef))
		}
		if err == nil && setting.OnChanging != nil {
			setting.OnChanging(newval)
		}
	}
	if oldval := setting.ValCfg; err == nil {
		setting.ValCfg = newval
		if err = Prog.Cfg.Save(); err == nil {
			if menu.NoteInfo = Strf("Setting `%s` was successfully %s.", setting.Title, info); setting.OnChanged != nil {
				go setting.OnChanged(oldval)
			}
		}
	}
	if err != nil {
		panic(err)
	}
}

func (this *SettingsBase) onListAll(menu *menuResp) {
	menu.SubMenu = &Menu{Desc: Strf("%s — %s:", this.MenuCategory(), this.cmdListAll.Title)}
	for _, ks := range this.Impl.KnownSettings() {
		svdef, svcur := "(empty)", "(default)"
		if ks.ValDef != nil && ks.ValDef != "" {
			svdef = Strf("%v", ks.ValDef)
		}
		if ks.ValCfg != nil && ks.ValCfg != "" {
			svcur = Strf("%v", ks.ValCfg)
		}
		ks.menuItem.Hint = Strf("Default: %s — Current: %s", svdef, svcur)
		ks.menuItem.IpcArgs = map[string]interface{}{"id": ks.Id, "val": menuItemIpcArgPrompt{Placeholder: ks.Desc,
			Prompt: "Specify as instructed, or clear to reset.", Value: ks.ValStr()}}
		menu.SubMenu.Items = append(menu.SubMenu.Items, ks.menuItem)
	}
}

func (this *SettingsBase) onResetAll() (num int, err error) {
	for _, ks := range Lang.Settings.KnownSettings() {
		if ks.ValCfg != nil {
			num, ks.ValCfg = num+1, nil
		}
	}
	err = Prog.Cfg.Save()
	return
}

func (this *SettingsBase) menuItems(*SrcLens) (menuItems MenuItems) {
	if Lang.Settings != nil {
		menuItems = MenuItems{this.cmdListAll}
		if num := Lang.Settings.KnownSettings().numCust(); num > 0 {
			this.cmdResetAll.Desc = Strf("Forgets %d current customization(s)", num)
			menuItems = append(menuItems, this.cmdResetAll)
		}
	}
	return
}

func (*SettingsBase) MenuCategory() string {
	return "Settings"
}
