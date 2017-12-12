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

type WorkspaceFiles map[string]*WorkspaceFile

func (me WorkspaceFiles) ensure(fpath string) (file *WorkspaceFile) {
	if file, _ = me[fpath]; file == nil {
		file = &WorkspaceFile{Path: fpath}
		me[fpath] = file
	}
	return
}

func (me WorkspaceFiles) exists(fpath string) bool {
	f, _ := me[fpath]
	return f != nil
}

type WorkspaceFile struct {
	Path   string
	IsOpen bool `json:",omitempty"`
	Diags  struct {
		Build Diags
		Lint  Diags
	}
}

func (me *WorkspaceFile) ForgetDiags() {
	me.Diags.Build.Forget()
	me.Diags.Lint.Forget()
}

type WorkspaceChanges struct {
	AddedDirs    []string
	RemovedDirs  []string
	OpenedFiles  []string
	ClosedFiles  []string
	WrittenFiles []string
}

type WorkspaceChangesBefore func(upd *WorkspaceChanges, dirsChanged bool, newFiles bool, willAutoLint bool)
type WorkspaceChangesAfter func(upd *WorkspaceChanges)

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

	OnBeforeChanges WorkspaceChangesBefore `json:"-"`
	OnAfterChanges  WorkspaceChangesAfter  `json:"-"`

	Dirs  map[string]*WorkspaceDir
	Files WorkspaceFiles

	pollingPaused bool
}

func (me *WorkspaceBase) Init() {
	me.Dirs = map[string]*WorkspaceDir{}
	me.Files = WorkspaceFiles{}
}

func (me *WorkspaceBase) lockAndPause() {
	me.mutex.Lock()
	me.pollingPaused = true
}

func (me *WorkspaceBase) unlockAndUnpause() {
	me.pollingPaused = false
	me.mutex.Unlock()
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
		hasnewfile, dirs, files, dirschanged :=
			false, me.Dirs, me.Files, len(upd.AddedDirs) > 0 || len(upd.RemovedDirs) > 0

		for _, eventfiles := range [][]string{upd.OpenedFiles, upd.ClosedFiles, upd.WrittenFiles} {
			for _, fpath := range eventfiles {
				if hasnewfile = !files.exists(fpath); hasnewfile {
					break
				}
			}
			if hasnewfile {
				break
			}
		}
		needsfreshautolints := hasnewfile || len(upd.WrittenFiles) > 0

		if needsfreshautolints || hasnewfile || dirschanged {
			me.lockAndPause()
			defer me.unlockAndUnpause()
		}
		if me.OnBeforeChanges != nil {
			me.OnBeforeChanges(upd, dirschanged, hasnewfile, needsfreshautolints)
		}

		if dirschanged {
			dirs = make(map[string]*WorkspaceDir, len(me.Dirs))
			for k, v := range me.Dirs {
				dirs[k] = v
			}

			for _, gonedirpath := range upd.RemovedDirs {
				delete(dirs, gonedirpath)
			}
			for _, newdirpath := range upd.AddedDirs {
				if dir, _ := dirs[newdirpath]; dir == nil {
					dir = &WorkspaceDir{Path: newdirpath}
					dirs[newdirpath] = dir
				}
			}
		}

		if hasnewfile {
			files = make(WorkspaceFiles, len(me.Files))
			for k, v := range me.Files {
				files[k] = v
			}
		}

		for _, gonefilepath := range upd.ClosedFiles {
			files.ensure(gonefilepath).IsOpen = false
		}
		for _, newfilepath := range upd.OpenedFiles {
			files.ensure(newfilepath).IsOpen = true
		}
		for _, modfilepath := range upd.WrittenFiles {
			files.ensure(modfilepath).ForgetDiags()
		}
		if needsfreshautolints && Lang.Diag != nil {
			Lang.Diag.UpdateLintDiagsIfAndAsNeeded(files, true)
		}

		me.Dirs, me.Files = dirs, files
		if me.OnAfterChanges != nil {
			me.OnAfterChanges(upd)
		}
	}
}

func (me *WorkspaceBase) ObjSnap(_ string) interface{} {
	return me.Impl
}

func (me *WorkspaceBase) ObjSnapPrefix() string {
	return Lang.ID + ".proj."
}

func (me *WorkspaceBase) PollFileEventsEvery(milliseconds int64) {
	interval := time.Millisecond * time.Duration(milliseconds)
	msg := &ipcResp{IpcID: IPCID_PROJ_POLLEVTS}
	for {
		if time.Sleep(interval); !me.pollingPaused {
			if !canSend() {
				return
			} else if err := send(msg); err != nil {
				return
			}
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
		for _, d := range me.Dirs {
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
