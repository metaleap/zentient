package z

import (
	"sync"
	"time"
)

type IWorkspace interface {
	iDispatcher
	IObjSnap
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
	sync.RWMutex
	Impl IWorkspace `json:"-"`

	OnBeforeChanges func(*WorkspaceChanges) `json:"-"`
	OnAfterChanges  func(*WorkspaceChanges) `json:"-"`

	OpenDirs  map[string]*WorkspaceDir
	OpenFiles map[string]*WorkspaceFile
}

func (me *WorkspaceBase) Init() {
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
		me.lock()
		defer me.unlock()

		if on := me.OnBeforeChanges; on != nil {
			on(upd)
		}
		for _, gonedir := range upd.RemovedDirs {
			delete(me.OpenDirs, gonedir)
		}
		for _, closedfile := range upd.ClosedFiles {
			delete(me.OpenFiles, closedfile)
		}
		for _, newdir := range upd.AddedDirs {
			if dir, _ := me.OpenDirs[newdir]; dir == nil {
				dir = &WorkspaceDir{Path: newdir}
				me.OpenDirs[newdir] = dir
			}
		}
		for _, newfile := range upd.OpenedFiles {
			if file, _ := me.OpenFiles[newfile]; file == nil {
				file = &WorkspaceFile{Path: newfile}
				me.OpenFiles[newfile] = file
			}
		}

		if on := me.OnAfterChanges; on != nil {
			on(upd)
		}
	}
}

func (me *WorkspaceBase) lock() {
	me.RLock()
	me.Lock()
}

func (me *WorkspaceBase) unlock() {
	me.Unlock()
	me.RUnlock()
}

func (me *WorkspaceBase) ObjSnap(_ string) interface{} {
	return me.Impl
}

func (me *WorkspaceBase) ObjSnapPrefix() string {
	return Lang.ID + ".proj."
}

func workspacePollFileEventsEvery(milliseconds int64) {
	interval := time.Millisecond * time.Duration(milliseconds)
	msg := &ipcResp{IpcID: IPCID_PROJ_POLLEVTS}
	for canSend() {
		send(msg)
		time.Sleep(interval)
	}
}
