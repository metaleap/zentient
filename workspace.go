package z

type IWorkspace interface {
	iDispatcher
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
	Impl IWorkspace `json:"-"`

	OnBeforeChanges []func(*WorkspaceChanges) `json:"-"`
	OnAfterChanges  []func(*WorkspaceChanges) `json:"-"`

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
	// case IPCID_PROJ_SNAPSHOT:
	// 	resp.ObjSnapshot = me.onSnapshot()
	default:
		return false
	}
	return true
}

func (me *WorkspaceBase) onChanges(upd *WorkspaceChanges) {
	if upd != nil && upd.hasChanges() {
		for _, on := range me.OnBeforeChanges {
			on(upd)
		}

		for _, gonedir := range upd.RemovedDirs {
			delete(me.OpenDirs, gonedir)
		}
		for _, closedfile := range upd.ClosedFiles {
			delete(me.OpenFiles, closedfile)
		}

		for _, newdir := range upd.AddedDirs {
			dir, _ := me.OpenDirs[newdir]
			if dir == nil {
				dir = &WorkspaceDir{Path: newdir}
				me.OpenDirs[newdir] = dir
			}
		}
		for _, newfile := range upd.OpenedFiles {
			file, _ := me.OpenFiles[newfile]
			if file == nil {
				file = &WorkspaceFile{Path: newfile}
				me.OpenFiles[newfile] = file
			}
		}

		for _, on := range me.OnAfterChanges {
			on(upd)
		}
	}
}

func (me *WorkspaceBase) onSnapshot() IWorkspace {
	return me.Impl
}
