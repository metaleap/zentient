package z

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/fs"
)

type ISettings interface {
	IMenuItems

	KnownSettings() Settings
}

type Settings []*Setting

func (me Settings) numCust() (n int) {
	for _, s := range me {
		if s.ValCfg != nil && s.ValCfg != s.ValDef {
			n++
		}
	}
	return
}

func (me Settings) titles() (all []string) {
	all = make([]string, len(me))
	for i, s := range me {
		all[i] = s.Title
	}
	return
}

type Setting struct {
	ID     string
	Title  string
	Desc   string
	ValCfg interface{}
	ValDef interface{}

	menuItem *MenuItem
}

func (me *Setting) Val() interface{} {
	if me.ValCfg != nil {
		return me.ValCfg
	}
	return me.ValDef
}

func (me *Setting) ValStr() (val string) {
	val, _ = me.Val().(string)
	return
}

type Config struct {
	Internal      map[string]interface{} `json:",omitempty"`
	FormatterName string                 `json:",omitempty"`
	FormatterProg string                 `json:",omitempty"`
	AutoDiags     []string               `json:",omitempty"`

	err            error
	recallFilePath string
	filePath       string
	timeLastLoaded int64
}

func (me *Config) reload() {
	if stale, _ := ufs.IsNewerThanTime(me.filePath, me.timeLastLoaded); stale {
		// 1. re-initialize me
		var empty Config
		*me = empty
		me.filePath = filepath.Join(Prog.dir.config, Prog.name+".config.json")

		// 2. load
		if ufs.FileExists(me.filePath) { // otherwise, it's a fresh setup
			if me.err = umisc.JsonDecodeFromFile(me.filePath, me); me.err == nil {
				me.timeLastLoaded = time.Now().UnixNano()
				if Lang.Settings != nil && me.Internal != nil {
					for _, ks := range Lang.Settings.KnownSettings() {
						if val, ok := me.Internal[ks.ID]; ok {
							ks.ValCfg = val
						}
					}
					me.Internal = nil
				}
			}
		}
	}
	return
}

func (me *Config) recall() {
	me.recallFilePath = filepath.Join(Prog.dir.cache, Prog.name+".recall.json")
	if ufs.FileExists(me.recallFilePath) {
		umisc.JsonDecodeFromFile(me.recallFilePath, &Prog.recall)
	}
	if Prog.recall.i64 == nil {
		Prog.recall.i64 = map[string]int64{}
	}
}

func (me *Config) saveRecall() {
	umisc.JsonEncodeToFile(&Prog.recall, me.recallFilePath)
}

func (me *Config) Save() (err error) {
	if Lang.Settings != nil {
		me.Internal = map[string]interface{}{}
		for _, ks := range Lang.Settings.KnownSettings() {
			if ks.ValCfg != nil && ks.ValCfg != ks.ValDef {
				me.Internal[ks.ID] = ks.ValCfg
			}
		}
		if len(me.Internal) == 0 {
			me.Internal = nil
		}
	}
	err = umisc.JsonEncodeToFile(me, me.filePath)
	me.Internal = nil
	return
}

type SettingsBase struct {
	Impl ISettings

	cmdListAll  *MenuItem
	cmdResetAll *MenuItem
}

func (me *SettingsBase) Init() {
	if Lang.Settings != nil {
		ks := me.Impl.KnownSettings()
		me.cmdListAll = &MenuItem{IpcID: IPCID_CFG_LIST, Title: Strf("%s-Specific", Lang.Title), Hint: Strf("%d setting(s)", len(ks)), Desc: Strf("Customize: %s", strings.Join(ks.titles(), " · "))}
		me.cmdResetAll = &MenuItem{IpcID: IPCID_CFG_RESETALL, Title: "Reset All", Hint: Strf("%s-Specific Settings", Lang.Title)}
		for _, s := range ks {
			s.menuItem = &MenuItem{Title: s.Title, Desc: s.Desc, IpcID: IPCID_CFG_SET}
		}
	}
}

func (me *SettingsBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_CFG_LIST:
		me.onListAll(resp.withMenu())
	case IPCID_CFG_RESETALL:
		if num, err := me.onResetAll(); err != nil {
			resp.ErrMsg = err.Error()
		} else {
			resp.withMenu().NoteInfo = Strf("%d customized setting(s) just reset to factory defaults.", num)
		}
	default:
		return false
	}
	return true
}

func (me *SettingsBase) onListAll(menu *MenuResp) {
	menu.SubMenu = &Menu{Desc: Strf("%s — %s:", me.MenuCategory(), me.cmdListAll.Title)}
	for _, ks := range me.Impl.KnownSettings() {
		svdef := Strf("%v", ks.ValDef)
		if ks.ValDef == nil || svdef == "" {
			svdef = "(empty)"
		}
		svcur := Strf("%v", ks.ValCfg)
		if ks.ValCfg == nil || svcur == "" {
			svcur = "(default)"
		}
		ks.menuItem.Hint = Strf("Default: %s  ·  Current: %s", svdef, svcur)
		ks.menuItem.IpcArgs = map[string]interface{}{"id": ks.ID, "val": MenuItemIpcArgPrompt{Placeholder: ks.Desc,
			Prompt: "Clear to reset.", Value: ks.ValStr()}}
		menu.SubMenu.Items = append(menu.SubMenu.Items, ks.menuItem)
	}
}

func (me *SettingsBase) onResetAll() (num int, err error) {
	for _, ks := range Lang.Settings.KnownSettings() {
		if ks.ValCfg != nil {
			num, ks.ValCfg = num+1, nil
		}
	}
	err = Prog.Cfg.Save()
	return
}

func (me *SettingsBase) MenuItems(*SrcLens) (menuItems MenuItems) {
	if Lang.Settings != nil {
		menuItems = MenuItems{me.cmdListAll}
		if num := Lang.Settings.KnownSettings().numCust(); num > 0 || true {
			me.cmdResetAll.Desc = Strf("Forgets %d current customization(s)", num)
			menuItems = append(menuItems, me.cmdResetAll)
		}
	}
	return
}

func (*SettingsBase) MenuCategory() string {
	return "Settings"
}
