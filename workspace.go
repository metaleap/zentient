package z

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/sys"
)

type IWorkspace interface {
	iDispatcher
	IObjSnap

	PollFileEventsEvery(int64)
	PrettyPath(string, ...string) string
}

type WorkspaceDir struct {
	Path string
}

type WorkspaceFile struct {
	Path string
}

type WorkspaceChanges struct {
	AddedDirs    []string
	RemovedDirs  []string
	OpenedFiles  []string
	ClosedFiles  []string
	WrittenFiles []string
}

func (me *WorkspaceChanges) hasChanges() bool {
	return len(me.AddedDirs) > 0 ||
		len(me.RemovedDirs) > 0 ||
		len(me.OpenedFiles) > 0 ||
		len(me.ClosedFiles) > 0 ||
		len(me.WrittenFiles) > 0
}

type WorkspaceBase struct {
	mutex sync.Mutex
	Impl  IWorkspace `json:"-"`

	OnBeforeChanges func(*WorkspaceChanges) `json:"-"`
	OnAfterChanges  func(*WorkspaceChanges) `json:"-"`

	OpenDirs  map[string]*WorkspaceDir
	OpenFiles map[string]*WorkspaceFile
}

func (me *WorkspaceBase) Init() {
	noop := func(_ *WorkspaceChanges) {}
	me.OnBeforeChanges, me.OnAfterChanges = noop, noop
	me.OpenDirs = map[string]*WorkspaceDir{}
	me.OpenFiles = map[string]*WorkspaceFile{}
}

func (me *WorkspaceBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_PROJ_CHANGED:
		me.onChanges(req.ProjUpd)
	default:
		return false
	}
	return true
}

func (me *WorkspaceBase) onChanges(upd *WorkspaceChanges) {
	if upd != nil && upd.hasChanges() {
		me.OnBeforeChanges(upd)

		me.mutex.Lock()
		defer me.mutex.Unlock()
		opendirs, openfiles := me.OpenDirs, me.OpenFiles

		if len(upd.AddedDirs) > 0 || len(upd.RemovedDirs) > 0 {
			opendirs = make(map[string]*WorkspaceDir, len(me.OpenDirs))
			for k, v := range me.OpenDirs {
				opendirs[k] = v
			}
			for _, gonedir := range upd.RemovedDirs {
				delete(opendirs, gonedir)
			}
			for _, newdir := range upd.AddedDirs {
				if dir, _ := opendirs[newdir]; dir == nil {
					dir = &WorkspaceDir{Path: newdir}
					opendirs[newdir] = dir
				}
			}
		}

		if len(upd.OpenedFiles) > 0 || len(upd.ClosedFiles) > 0 {
			openfiles = make(map[string]*WorkspaceFile, len(me.OpenFiles))
			for k, v := range me.OpenFiles {
				openfiles[k] = v
			}
			for _, gonefile := range upd.ClosedFiles {
				delete(openfiles, gonefile)
			}
			for _, newfile := range upd.OpenedFiles {
				if file, _ := openfiles[newfile]; file == nil {
					file = &WorkspaceFile{Path: newfile}
					openfiles[newfile] = file
				}
			}
		}

		me.OpenDirs, me.OpenFiles = opendirs, openfiles
		me.OnAfterChanges(upd)
	}
}

func (me *WorkspaceBase) ObjSnap(_ string) interface{} {
	return me.Impl
}

func (me *WorkspaceBase) ObjSnapPrefix() string {
	return Lang.ID + ".proj."
}

func (*WorkspaceBase) PollFileEventsEvery(milliseconds int64) {
	interval := time.Millisecond * time.Duration(milliseconds)
	msg := &ipcResp{IpcID: IPCID_PROJ_POLLEVTS}
	for {
		time.Sleep(interval)
		if !canSend() {
			return
		} else if err := send(msg); err != nil {
			return
		}
	}
}

func (me *WorkspaceBase) PrettyPath(fsPath string, otherEnvs ...string) string {
	if fsPath != "" {
		rel := func(path string) string {
			if rp, err := filepath.Rel(path, fsPath); path != "" && err == nil && rp != "" && !strings.HasPrefix(rp, ".") {
				return rp
			}
			return ""
		}

		candidates := []string{}
		for _, d := range me.OpenDirs {
			if rp := rel(d.Path); rp != "" {
				candidates = append(candidates, filepath.Join("…", filepath.Base(d.Path), rp))
			}
		}
		if shortest := ""; len(candidates) > 0 {
			for _, c := range candidates {
				if shortest == "" || len(c) < len(shortest) {
					shortest = c
				}
			}
			return shortest
		}

		for _, gopath := range udevgo.GoPaths {
			if rp := rel(gopath); rp != "" {
				return filepath.Join("$GOPATH", rp)
			}
		}

		for _, envname := range otherEnvs {
			if envval := os.Getenv(envname); envval != "" {
				if envpaths := filepath.SplitList(envval); len(envpaths) > 0 {
					for _, envpath := range envpaths {
						if rp := rel(envpath); rp != "" {
							return filepath.Join("$"+envname, rp)
						}
					}
				}
			}
		}

		if rp := rel(usys.UserHomeDirPath()); rp != "" {
			return filepath.Join("~", rp)
		}
	}
	return fsPath
}
