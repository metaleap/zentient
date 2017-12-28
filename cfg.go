package z

import (
	"path/filepath"
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

	menuItems   MenuItems
	cmdResetAll *MenuItem
}

func (me *SettingsBase) Init() {
	if Lang.Settings != nil {
		me.menuItems = MenuItems{&MenuItem{Title: "Reset All", Hint: Strf("%s-Specific Settings", Lang.Title), IpcID: IPCID_CFG_RESETALL}}
		for _, ks := range Lang.Settings.KnownSettings() {
			ks.menuItem = &MenuItem{Title: ks.Title, Desc: ks.Desc, IpcArgs: ks.ID}
			me.menuItems = append(me.menuItems, ks.menuItem)
		}
	}
}

func (me *SettingsBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_CFG_RESETALL:
		if num, err := me.onResetAll(); err != nil {
			resp.ErrMsg = err.Error()
		} else {
			resp.Menu = &MenuResp{NoteInfo: Strf("%d customized setting(s) just reset to factory defaults.", num)}
		}
	default:
		return false
	}
	return true
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

func (me *SettingsBase) MenuItems(*SrcLens) MenuItems {
	if len(me.menuItems) > 0 {
		me.menuItems[0].Desc = Strf("Forgets %d current customization(s)", Lang.Settings.KnownSettings().numCust())
	}
	return me.menuItems
}

func (*SettingsBase) MenuCategory() string {
	return "Settings"
}
