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

func (me *DiagItem) resetAndInferSrcActions(maybeOrigSrcRef *udev.SrcMsg) {
	me.SrcActions = nil
	if ilastcolon := ustr.Last(me.Msg, ":"); ilastcolon > 0 {
		if ustr.ToInt(me.Msg[ilastcolon+1:], 0) > 0 {
			if ifirstsep := ustr.IdxR(me.Msg, filepath.Separator); ifirstsep >= 0 {
				refpath := me.Msg[ifirstsep:]
				refpathf := refpath[:ustr.IdxR(refpath, ':')]
				if !ufs.IsFile(refpathf) {
					for i := ifirstsep - 1; i > 0; i-- {
						refpath = me.Msg[i:]
						if refpathf = refpath[:ustr.IdxR(refpath, ':')]; ufs.IsFile(refpathf) {
							break
						}
					}
				}
				if ufs.IsFile(refpathf) && !filepath.IsAbs(refpathf) {
					refpathf, _ = filepath.Abs(refpathf)
				}
				if ufs.IsFile(refpathf) {
					fpathref := refpathf + refpath[ustr.IdxR(refpath, ':'):]
					me.SrcActions = append(me.SrcActions, EditorAction{
						Cmd:       "zen.internal.openFileAt",
						Title:     Strf("Jump to %s", filepath.Base(fpathref)),
						Arguments: []string{fpathref},
					})
				}
			}
		}
	}
	if maybeOrigSrcRef != nil && maybeOrigSrcRef.Data != nil {
		if xfrom, xto, xnotes := maybeOrigSrcRef.Data["From"], maybeOrigSrcRef.Data["To"], maybeOrigSrcRef.Data["Note"]; xfrom != nil && xto != nil {
			from, _ := xfrom.(string)
			to, _ := xto.(string)
			notes, _ := xnotes.([]string)
			if from != "" && to != "" {
				me.SrcActions = append(me.SrcActions, EditorAction{
					Cmd:       "zen.internal.replaceText",
					Title:     "Apply Suggestion",
					Hint:      ustr.Join(notes, "\n"),
					Arguments: []string{from, to},
				})
			}
			me.Msg += " —\n" + ustr.Join(append([]string{"Instead of:\n\t" + from, "Consider:\n\t" + to}, notes...), "\n")
		}
	}
}

func (me DiagItems) propagate(lintDiags bool, diagsSticky bool, workspaceFiles WorkspaceFiles) {
	for _, diag := range me {
		f := workspaceFiles.ensure(diag.Loc.FilePath)
		fd := &f.Diags.Lint
		if (!lintDiags) && diag.Loc.Flag == int(DIAG_SEV_ERR) {
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
	// ISortable
	// fmt.Stringer
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

func (me *DiagJob) String() string {
	if str, _ := me.Target.(fmt.Stringer); str != nil {
		return str.String()
	}
	return Strf("%v", me.Target)
}

func (me *DiagBase) NewDiagItemFrom(srcRef *udev.SrcMsg, toolName string, fallbackFilePath func() string) (di *DiagItem) {
	di = &DiagItem{Msg: ustr.Trim(srcRef.Msg), Cat: toolName}
	di.Loc.Flag = srcRef.Flag
	di.Loc.SetFilePathAndPosOrRangeFrom(srcRef, fallbackFilePath)
	di.resetAndInferSrcActions(srcRef)
	return
}

func (me DiagItems) dropDupls() DiagItems {
	for i := 0; i < len(me); i++ {
		for j := i + 1; j < len(me); j++ {
			if me[i].equivTo(me[j]) {
				me = append(me[:j], me[j+1:]...)
				i = i - 1
				break
			}
		}
	}
	return me
}

func (me *DiagItem) equivTo(cmp *DiagItem) bool {
	if me != nil && cmp != nil && me != cmp {
		if !(me.Cat == cmp.Cat && me.Msg == cmp.Msg && me.StickyAuto == cmp.StickyAuto && me.StickyForce == cmp.StickyForce && len(me.Tags) == len(cmp.Tags) && len(me.Rel) == len(cmp.Rel) && len(me.SrcActions) == len(cmp.SrcActions) && me.Loc.equivTo(&cmp.Loc)) {
			return false
		} else {
			for i, t := range me.Tags {
				if t != cmp.Tags[i] {
					return false
				}
			}
			for i := range me.Rel {
				if !(me.Rel[i].equivTo(&cmp.Rel[i])) {
					return false
				}
			}
			for i := range me.SrcActions {
				if !me.SrcActions[i].equivTo(&cmp.SrcActions[i]) {
					return false
				}
			}
			return true
		}
	}
	return me == cmp
}

func (me *EditorAction) equivTo(cmp *EditorAction) bool {
	if me != nil && cmp != nil && me != cmp {
		if len(me.Arguments) != len(cmp.Arguments) {
			return false
		}
		for i := range me.Arguments {
			if me.Arguments[i] != cmp.Arguments[i] {
				return false
			}
		}
		return me.Cmd == cmp.Cmd && me.Hint == cmp.Hint && me.Title == cmp.Title
	}
	return me == cmp
}

func (me *SrcLens) equivTo(cmp *SrcLens) bool {
	if me != nil && cmp != nil && me != cmp {
		return me.CrLf == cmp.CrLf && me.Str == cmp.Str && me.Txt == cmp.Txt &&
			me.SrcLoc.equivTo(&cmp.SrcLoc)
	}
	return me == cmp
}

func (me *SrcLoc) equivTo(cmp *SrcLoc) bool {
	if me != nil && cmp != nil && me != cmp {
		return me.FilePath == cmp.FilePath && me.Flag == cmp.Flag &&
			me.Pos.equivTo(cmp.Pos) && me.Range.equivTo(cmp.Range)
	}
	return me == cmp
}

func (me *SrcPos) equivTo(cmp *SrcPos) bool {
	if me != nil && cmp != nil && me != cmp {
		return me.Off == cmp.Off || (me.Ln == cmp.Ln && me.Col == cmp.Col)
	}
	return me == cmp
}

func (me *SrcRange) equivTo(cmp *SrcRange) bool {
	if me != nil && cmp != nil && me != cmp {
		return me.Start.equivTo(&cmp.Start) && me.End.equivTo(&cmp.End)
	}
	return me == cmp
}
