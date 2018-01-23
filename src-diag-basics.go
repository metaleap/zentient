package z

import (
	"fmt"
	"path/filepath"

	"github.com/go-leap/dev"
	"github.com/go-leap/fs"
	"github.com/go-leap/str"
)

type IDiag interface {
	IDiagBuild
	IDiagLint
	IMenuItems
	send(WorkspaceFiles, bool)
}

type diags struct {
	Items DiagItems `json:",omitempty"`
}

func (me *diags) forget(onlyFor Tools) {
	if len(onlyFor) == 0 {
		me.Items = nil
	} else {
		for i := 0; i < len(me.Items); i++ {
			if onlyFor.has(me.Items[i].Cat) {
				pre, post := me.Items[:i], me.Items[i+1:]
				i, me.Items = i-1, append(pre, post...)
			}
		}
	}
}

type diagItemsBy map[string]DiagItems

type DiagItem struct {
	Cat         string `json:",omitempty"`
	Loc         SrcLoc
	Msg         string
	SrcActions  []EditorAction `json:",omitempty"`
	StickyForce bool           `json:"-"`
	StickyAuto  bool           `json:"Sticky,omitempty"`
}

func (me *DiagItem) resetAndInferSrcActions() {
	me.SrcActions = nil
	if ilastcolon := ustr.Last(me.Msg, ":"); ilastcolon > 0 {
		if ustr.ToInt(me.Msg[ilastcolon+1:], 0) > 0 {
			if ifirstsep := ustr.Idx(me.Msg, filepath.Separator); ifirstsep >= 0 {
				refpath := me.Msg[ifirstsep:]
				refpathf := refpath[:ustr.Idx(refpath, ':')]
				if !ufs.IsFile(refpathf) {
					for i := ifirstsep - 1; i > 0; i-- {
						refpath = me.Msg[i:]
						if refpathf = refpath[:ustr.Idx(refpath, ':')]; ufs.IsFile(refpathf) {
							break
						}
					}
				}
				if ufs.IsFile(refpathf) && !filepath.IsAbs(refpathf) {
					refpathf, _ = filepath.Abs(refpathf)
				}
				if ufs.IsFile(refpathf) {
					cmd := EditorAction{Cmd: "zen.internal.openFileAt", Title: refpathf + refpath[ustr.Idx(refpath, ':'):]}
					cmd.Arguments = append(cmd.Arguments, cmd.Title)
					cmd.Title = Strf("Jump to %s", filepath.Base(cmd.Title))
					me.SrcActions = append(me.SrcActions, cmd)
				}
			}
		}
	}
}

type DiagItems []*DiagItem

func (me DiagItems) propagate(lintDiags bool, diagsSticky bool, workspaceFiles WorkspaceFiles) {
	for _, diag := range me {
		f := workspaceFiles.ensure(diag.Loc.FilePath)
		fd := &f.Diags.Lint
		if !lintDiags {
			fd = &f.Diags.Build
		}
		if diag.StickyForce, fd.Items = diagsSticky, append(fd.Items, diag); diagsSticky {
			diag.StickyAuto = true
		} else {
			diag.StickyAuto = uint64(diag.Loc.Flag) <= cfgLintStickiness.ValUInt()
		}
	}
}

type IDiagJobTarget interface {
	ISortable
	fmt.Stringer
}

type DiagJob struct {
	AffectedFilePaths []string
	Target            IDiagJobTarget
}

func (me *DiagJob) forgetPrevDiags(diagToolsIfLint Tools, setAutoUpToDateToTrueIfLint bool, workspaceFiles WorkspaceFiles) {
	forbuild := len(diagToolsIfLint) == 0
	var f *WorkspaceFile
	for _, fpath := range me.AffectedFilePaths {
		if setAutoUpToDateToTrueIfLint {
			f = workspaceFiles.ensure(fpath)
		} else {
			f = workspaceFiles[fpath]
		}
		if f != nil {
			if forbuild {
				f.Diags.Build.forget(nil)
				f.Diags.AutoLintUpToDate = false
			} else if setAutoUpToDateToTrueIfLint {
				f.Diags.AutoLintUpToDate = true
			}
			f.Diags.Lint.forget(diagToolsIfLint)
		}
	}
}

func (me *DiagJob) String() string { return me.Target.String() }

type diagResp struct {
	All    diagItemsBy
	FixUps []*fixUps
	LangID string
}

func (me *DiagBase) NewDiagItemFrom(srcRef *udev.SrcMsg, toolName string, fallbackFilePath func() string) (di *DiagItem) {
	di = &DiagItem{Msg: ustr.Trim(srcRef.Msg), Cat: toolName}
	di.Loc.Flag = srcRef.Flag
	di.Loc.SetFilePathAndPosOrRangeFrom(srcRef, fallbackFilePath)
	di.resetAndInferSrcActions()
	return
}
