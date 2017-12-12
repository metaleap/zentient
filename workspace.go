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

type WorkspaceFile struct {
	Path   string
	IsOpen bool `json:",omitempty"`
	Diags  struct {
		UpToDate bool        `json:",omitempty"`
		Build    []*DiagItem `json:",omitempty"`
		Lint     []*DiagItem `json:",omitempty"`
	}
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

	Dirs  map[string]*WorkspaceDir
	Files WorkspaceFiles
}

func (me *WorkspaceBase) Init() {
	noop := func(_ *WorkspaceChanges) {}
	me.OnBeforeChanges, me.OnAfterChanges = noop, noop
	me.Dirs = map[string]*WorkspaceDir{}
	me.Files = WorkspaceFiles{}
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
		dirs, files := me.Dirs, me.Files

		if len(upd.AddedDirs) > 0 || len(upd.RemovedDirs) > 0 {
			dirs = make(map[string]*WorkspaceDir, len(me.Dirs))
			for k, v := range me.Dirs {
				dirs[k] = v
			}

			for _, gonedir := range upd.RemovedDirs {
				delete(dirs, gonedir)
			}
			for _, newdir := range upd.AddedDirs {
				if dir, _ := dirs[newdir]; dir == nil {
					dir = &WorkspaceDir{Path: newdir}
					dirs[newdir] = dir
				}
			}
		}

		if hasnewfile, needsrelints := false, len(upd.WrittenFiles) > 0; needsrelints || len(upd.OpenedFiles) > 0 || len(upd.ClosedFiles) > 0 {
			for _, sl := range [][]string{upd.OpenedFiles, upd.ClosedFiles, upd.WrittenFiles} {
				for _, fp := range sl {
					if f, _ := files[fp]; f == nil {
						hasnewfile = true
						break
					}
				}
				if hasnewfile {
					break
				}
			}
			if needsrelints = hasnewfile || needsrelints; hasnewfile {
				files = make(WorkspaceFiles, len(me.Files))
				for k, v := range me.Files {
					files[k] = v
				}
			}

			for _, gonefile := range upd.ClosedFiles {
				files.ensure(gonefile).IsOpen = false
			}
			for _, newfile := range upd.OpenedFiles {
				files.ensure(newfile).IsOpen = true
			}
			for _, modfile := range upd.WrittenFiles {
				file := files.ensure(modfile)
				file.Diags.UpToDate, file.Diags.Build, file.Diags.Lint = false, nil, nil
			}
			if Lang.Diag != nil && needsrelints {
				go Lang.Diag.UpdateLintDiagsIfAndAsNeeded(files, true)
			}
		}

		me.Dirs, me.Files = dirs, files
		me.OnAfterChanges(upd)
	}
}

func (me WorkspaceFiles) ensure(fpath string) (file *WorkspaceFile) {
	if file, _ = me[fpath]; file == nil {
		file = &WorkspaceFile{Path: fpath}
		me[fpath] = file
	}
	return
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
