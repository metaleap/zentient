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

	err            error
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

func (me *Config) Save() error {
	return umisc.JsonEncodeToFile(me, me.filePath)
}
