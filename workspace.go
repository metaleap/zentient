package z

import (
	"sync"
	"time"
)

type IWorkspace interface {
	iDispatcher
	IObjSnap

	PollFileEventsEvery(int64)
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
