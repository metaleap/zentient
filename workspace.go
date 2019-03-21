package z

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/str"
	"github.com/go-leap/sys"
)

type IWorkspace interface {
	iDispatcher
	IObjSnap
	json.Marshaler
	sync.Locker

	Dirs() WorkspaceDirs
	Files() WorkspaceFiles
	PrettyPath(string, ...string) string
}

type WorkspaceDir struct {
	Path string
}

type WorkspaceDirs map[string]*WorkspaceDir

type WorkspaceFile struct {
	Path   string
	IsOpen bool `json:",omitempty"`
	Diags  struct {
		AutoLintUpToDate bool
		Build            diags
		Lint             diags
	}

	existsAtLastCheck bool
	modTime           int64
}

func (this *WorkspaceFile) updateModTime() (hasChanged bool) {
	fileinfo, err := os.Stat(this.Path)
	if this.existsAtLastCheck = (err == nil) && fileinfo.Mode().IsRegular(); this.existsAtLastCheck {
		modtime := fileinfo.ModTime().UnixNano()
		this.modTime, hasChanged = modtime, modtime > this.modTime
	}
	return
}

func (this *WorkspaceFile) resetDiags() {
	this.Diags.Build.forget(nil)
	this.Diags.Lint.forget(nil)
	this.Diags.AutoLintUpToDate = false
}

type WorkspaceFiles map[string]*WorkspaceFile

func (this WorkspaceFiles) ensure(fpath string) (file *WorkspaceFile) {
	if file = this[fpath]; file == nil {
		file = &WorkspaceFile{Path: fpath}
		if file.updateModTime(); file.existsAtLastCheck {
			this[fpath] = file
		}
	}
	return
}

type diagsSummary struct {
	numBuild int
	numLint  int
	files    map[*WorkspaceFile]bool
}

func (this WorkspaceFiles) diagsSummary() *diagsSummary {
	s := &diagsSummary{files: make(map[*WorkspaceFile]bool, len(this))}
	for _, f := range this {
		if nb, nl := len(f.Diags.Build.Items), len(f.Diags.Lint.Items); nb > 0 || nl > 0 {
			s.numBuild, s.numLint, s.files[f] = s.numBuild+nb, s.numLint+nl, true
		}
	}
	if s.numBuild == 0 && s.numLint == 0 {
		return nil
	}
	return s
}

func (this WorkspaceFiles) haveAnyDiags(buildDiags bool, lintDiags bool) bool {
	for _, f := range this {
		if lb, ll := len(f.Diags.Build.Items), len(f.Diags.Lint.Items); (buildDiags && lb > 0) || (lintDiags && ll > 0) {
			return true
		}
	}
	return false
}

func (this WorkspaceFiles) HasBuildDiags(filePath string) (has bool) {
	if f := this[filePath]; f != nil {
		has = len(f.Diags.Build.Items) > 0
	}
	return
}

func (this WorkspaceFiles) exists(fpath string) bool {
	return this[fpath] != nil
}

func (this WorkspaceFiles) filePathsOpened() (all []string) {
	all = make([]string, 0, len(this))
	for _, f := range this {
		if f.IsOpen {
			all = append(all, f.Path)
		}
	}
	return
}

func (this WorkspaceFiles) filePathsKnown() (all []string) {
	var i int
	all = make([]string, len(this))
	for fp := range this {
		all[i], i = fp, i+1
	}
	return
}

func (this WorkspaceFiles) numDirs(incl func(*WorkspaceFile) bool) int {
	filedirs := make(map[string]bool, len(this))
	for _, f := range this {
		if incl == nil || incl(f) {
			filedirs[filepath.Dir(f.Path)] = true
		}
	}
	return len(filedirs)
}

type WorkspaceChanges struct {
	AddedDirs    []string
	RemovedDirs  []string
	OpenedFiles  []string
	ClosedFiles  []string
	WrittenFiles []string
}

func (this *WorkspaceChanges) hasChanges() bool {
	return this.HasDirChanges() || this.hasFileChanges()
}

func (this *WorkspaceChanges) HasDirChanges() bool {
	return len(this.AddedDirs) > 0 || len(this.RemovedDirs) > 0
}

func (this *WorkspaceChanges) hasFileIn(filePath string, slices ...[]string) bool {
	if len(slices) == 0 {
		slices = [][]string{this.ClosedFiles, this.OpenedFiles, this.WrittenFiles}
	}
	for _, eventfiles := range slices {
		for _, fpath := range eventfiles {
			if fpath == filePath {
				return true
			}
		}
	}
	return false
}

func (this *WorkspaceChanges) hasFileChanges() bool {
	return len(this.OpenedFiles) > 0 || len(this.ClosedFiles) > 0 || len(this.WrittenFiles) > 0
}

type WorkspaceChangesBefore func(upd *WorkspaceChanges, freshFiles []string, willAutoLint bool)
type WorkspaceChangesAfter func(upd *WorkspaceChanges)

type WorkspaceBase struct {
	mutex sync.Mutex
	Impl  IWorkspace `json:"-"`

	OnBeforeChanges WorkspaceChangesBefore `json:"-"`
	OnAfterChanges  WorkspaceChangesAfter  `json:"-"`

	dirs  WorkspaceDirs
	files WorkspaceFiles

	pollingStarted bool
	pollingPaused  bool
}

func (this *WorkspaceBase) Init()                         { this.dirs, this.files = WorkspaceDirs{}, WorkspaceFiles{} }
func (this *WorkspaceBase) Dirs() (dirs WorkspaceDirs)    { dirs = this.dirs; return }
func (this *WorkspaceBase) Files() (files WorkspaceFiles) { files = this.files; return }

func (this *WorkspaceBase) Lock() {
	this.pollingPaused = true
	this.mutex.Lock()
}

func (this *WorkspaceBase) Unlock() {
	this.mutex.Unlock()
	this.pollingPaused = false
}

func (this *WorkspaceBase) MarshalJSON() ([]byte, error) {
	var obj struct {
		Dirs  WorkspaceDirs
		Files WorkspaceFiles
	}
	obj.Dirs, obj.Files = this.dirs, this.files
	return json.Marshal(&obj)
}

func (this *WorkspaceBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_PROJ_CHANGED:
		this.onChanges(req.ProjUpd)
	default:
		return false
	}
	return true
}

func (*WorkspaceBase) analyzeChanges(files WorkspaceFiles, upd *WorkspaceChanges) (freshFiles []string, hasFreshFiles bool, hasDiedFiles bool, dirsChanged bool, needsFreshAutoLints bool) {
	for _, eventfiles := range [][]string{upd.OpenedFiles, upd.ClosedFiles, upd.WrittenFiles} {
		for _, fpath := range eventfiles {
			if !files.exists(fpath) {
				freshFiles = append(freshFiles, fpath)
			}
		}
	}
	for _, file := range files { // we really only check all-known-files for changes seemingly-somewhat-redudantly here to cover changes outside the editor (from code-gens etc)
		if haschanged := file.updateModTime(); !file.existsAtLastCheck {
			hasDiedFiles = true
		} else if haschanged && !upd.hasFileIn(file.Path, upd.WrittenFiles) {
			upd.WrittenFiles = append(upd.WrittenFiles, file.Path)
		}
	}
	hasFreshFiles, dirsChanged = len(freshFiles) > 0, upd.HasDirChanges()
	needsFreshAutoLints = hasFreshFiles || len(upd.WrittenFiles) > 0 || len(upd.OpenedFiles) > 0
	return
}

func (this *WorkspaceBase) onChanges(upd *WorkspaceChanges) {
	if upd != nil && upd.hasChanges() {
		dirs, files := this.dirs, this.files
		freshfiles, hasfreshfiles, hasdiedfiles, dirschanged, needsfreshautolints := this.analyzeChanges(files, upd)
		if needsfreshautolints || hasfreshfiles || hasdiedfiles || dirschanged {
			this.Lock()
			defer this.Unlock()
		}
		if this.OnBeforeChanges != nil {
			this.OnBeforeChanges(upd, freshfiles, needsfreshautolints)
		}

		if dirschanged {
			dirs = make(WorkspaceDirs, len(this.dirs))
			for k, v := range this.dirs {
				dirs[k] = v
			}

			for _, gonedirpath := range upd.RemovedDirs {
				delete(dirs, gonedirpath)
			}
			for _, newdirpath := range upd.AddedDirs {
				if dir := dirs[newdirpath]; dir == nil {
					dir = &WorkspaceDir{Path: ustr.TrimR(newdirpath, "/\\")}
					dirs[newdirpath] = dir
				}
			}
		}

		if hasfreshfiles || hasdiedfiles {
			files = make(WorkspaceFiles, len(this.files))
			for k, v := range this.files {
				if v.existsAtLastCheck {
					files[k] = v
				}
			}
		}

		for _, gonefilepath := range upd.ClosedFiles {
			files.ensure(gonefilepath).IsOpen = false
		}
		for _, freshfilepath := range upd.OpenedFiles {
			files.ensure(freshfilepath).IsOpen = true
		}
		for _, modfilepath := range upd.WrittenFiles {
			files.ensure(modfilepath).resetDiags()
		}
		if Lang.Diag != nil {
			if len(upd.WrittenFiles) > 0 {
				Lang.Diag.UpdateBuildDiagsAsNeeded(files, upd.WrittenFiles)
			}
			if needsfreshautolints {
				Lang.Diag.UpdateLintDiagsIfAndAsNeeded(files, true)
			}
		}
		this.dirs, this.files = dirs, files
		if this.OnAfterChanges != nil {
			this.OnAfterChanges(upd)
		}
	}
	if !this.pollingStarted {
		this.pollingStarted = true
		go this.pollFileEventsForever()
	}
}

func (this *WorkspaceBase) ObjSnap(string) interface{} { return this.Impl }

func (*WorkspaceBase) ObjSnapPrefix() string { return Lang.ID + ".proj." }

func (this *WorkspaceBase) pollFileEventsForever() {
	interval := 1234 * time.Millisecond
	msgraw, _ := json.Marshal(&ipcResp{IpcID: IPCID_PROJ_POLLEVTS})
	msgraw = append(msgraw, '\n')
	for {
		if time.Sleep(interval); !this.pollingPaused {
			if !canSend() {
				return
			} else if err := sendRaw(msgraw); err != nil {
				return
			}
		}
	}
}

func (*WorkspaceBase) prettyPathRel(path string, fsPath string) string {
	if path != "" {
		if rp, err := filepath.Rel(path, fsPath); err == nil && rp != "" && !ustr.Pref(rp, ".") {
			return rp
		}
	}
	return ""
}

func (this *WorkspaceBase) PrettyPath(fsPath string, otherEnvs ...string) string {
	if fsPath != "" {
		rel := func(path string) string { return this.prettyPathRel(path, fsPath) }

		candidates := []string{}
		for _, d := range this.dirs {
			if rp := rel(d.Path); rp != "" {
				candidates = append(candidates, filepath.Join("...", filepath.Base(d.Path), rp))
			}
		}
		if len(candidates) > 0 {
			return ustr.Shortest(candidates)
		}

		for _, gopath := range udevgo.Gopaths() {
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
