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

	Dirs() WorkspaceDirs
	Files() WorkspaceFiles
	PollFileEventsEvery(int64)
	PrettyPath(string, ...string) string
}

type WorkspaceDir struct {
	Path string
}

type WorkspaceDirs map[string]*WorkspaceDir

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

type WorkspaceChangesBefore func(upd *WorkspaceChanges, dirsChanged bool, newFiles []string, willAutoLint bool)
type WorkspaceChangesAfter func(upd *WorkspaceChanges)

func (me *WorkspaceChanges) hasChanges() bool {
	return len(me.AddedDirs) > 0 || len(me.RemovedDirs) > 0 || len(me.OpenedFiles) > 0 || len(me.ClosedFiles) > 0 || len(me.WrittenFiles) > 0
}

type WorkspaceBase struct {
	mutex sync.Mutex
	Impl  IWorkspace `json:"-"`

	OnBeforeChanges WorkspaceChangesBefore `json:"-"`
	OnAfterChanges  WorkspaceChangesAfter  `json:"-"`

	dirs  WorkspaceDirs
	files WorkspaceFiles

	pollingPaused bool
}

func (me *WorkspaceBase) Init()                         { me.dirs, me.files = WorkspaceDirs{}, WorkspaceFiles{} }
func (me *WorkspaceBase) Dirs() (dirs WorkspaceDirs)    { dirs = me.dirs; return }
func (me *WorkspaceBase) Files() (files WorkspaceFiles) { files = me.files; return }

func (me *WorkspaceBase) lockAndPausePolling() {
	me.pollingPaused = true
	me.mutex.Lock()
}

func (me *WorkspaceBase) unlockAndResumePolling() {
	me.mutex.Unlock()
	me.pollingPaused = false
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

func (_ *WorkspaceBase) analyzeChanges(files WorkspaceFiles, upd *WorkspaceChanges) (newFiles []string, hasNewFiles bool, dirsChanged bool, needsFreshAutoLints bool) {
	for _, eventfiles := range [][]string{upd.OpenedFiles, upd.ClosedFiles, upd.WrittenFiles} {
		for _, fpath := range eventfiles {
			if !files.exists(fpath) {
				newFiles = append(newFiles, fpath)
			}
		}
	}
	hasNewFiles, dirsChanged = len(newFiles) > 0, len(upd.AddedDirs) > 0 || len(upd.RemovedDirs) > 0
	needsFreshAutoLints = hasNewFiles || len(upd.WrittenFiles) > 0
	return
}

func (me *WorkspaceBase) onChanges(upd *WorkspaceChanges) {
	if upd != nil && upd.hasChanges() {
		dirs, files := me.dirs, me.files
		newfiles, hasnewfiles, dirschanged, needsfreshautolints := me.analyzeChanges(files, upd)

		if needsfreshautolints || hasnewfiles || dirschanged {
			me.lockAndPausePolling()
			defer me.unlockAndResumePolling()
		}
		if me.OnBeforeChanges != nil {
			me.OnBeforeChanges(upd, dirschanged, newfiles, needsfreshautolints)
		}

		if dirschanged {
			dirs = make(WorkspaceDirs, len(me.dirs))
			for k, v := range me.dirs {
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

		if hasnewfiles {
			files = make(WorkspaceFiles, len(me.files))
			for k, v := range me.files {
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

		me.dirs, me.files = dirs, files
		if me.OnAfterChanges != nil {
			me.OnAfterChanges(upd)
		}
	}
}

func (me *WorkspaceBase) ObjSnap(_ string) interface{} { return me.Impl }

func (me *WorkspaceBase) ObjSnapPrefix() string { return Lang.ID + ".proj." }

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
		for _, d := range me.dirs {
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
