package z

import (
	"path/filepath"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/fs"
)

type Config struct {
	Misc          map[string]string `json:",omitempty"`
	FormatterName string            `json:",omitempty"`
	FormatterProg string            `json:",omitempty"`
	AutoDiags     []string          `json:",omitempty"`

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
		me.Misc = map[string]string{}
		me.filePath = filepath.Join(Prog.dir.config, Prog.name+".config.json")

		// 2. load
		if ufs.FileExists(me.filePath) { // otherwise, it's a fresh setup
			if me.err = umisc.JsonDecodeFromFile(me.filePath, me); me.err == nil {
				me.timeLastLoaded = time.Now().UnixNano()
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

func (me *Config) Save() error {
	return umisc.JsonEncodeToFile(me, me.filePath)
}
