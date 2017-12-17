package zgo

import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/zentient"
)

var diag goDiag

func init() {
	diag.Impl, z.Lang.Diag = &diag, &diag
}

type goDiag struct {
	z.DiagBase

	knownDiags z.Tools
}

func (me *goDiag) onPreInit() {
	me.knownDiags = tools.KnownToolsFor(z.TOOLS_CAT_DIAGS)
}

func (me *goDiag) KnownDiags() z.Tools {
	return me.knownDiags
}

func (me *goDiag) runDiag(tool *z.Tool, pkg *udevgo.Pkg, yield chan *z.DiagItem, onjobdone func()) {
	defer onjobdone()

	diag := func(i int, fpath string, found string) *z.DiagItem {
		return &z.DiagItem{Message: "Found " + found, ToolName: tool.Name, FileRef: z.SrcLens{FilePath: fpath, Flag: int(z.DIAG_SEV_WARN), Pos: &z.SrcPos{Off: i + 1}}}
	}

	for _, filename := range pkg.GoFiles {
		fpath := filepath.Join(pkg.Dir, filename)
		filesrc := ufs.ReadTextFile(fpath, true, "")
		if idx := strings.Index(filesrc, "/sys"); tool.Name == "gosimple" && idx >= 0 {
			yield <- diag(idx, fpath, "/sys")
		}
		if idx := strings.Index(filesrc, "/run"); tool.Name == "goconst" && idx >= 0 {
			yield <- diag(idx, fpath, "/run")
		}
		if idx := strings.Index(filesrc, "/slice"); tool.Name == "golint" && idx >= 0 {
			yield <- diag(idx, fpath, "/slice")
		}
	}
}

func (me *goDiag) UpdateLintDiags(workspaceFiles z.WorkspaceFiles, diagTools z.Tools, filePaths []string) {
	if pkgs := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		numjobs, await := 0, make(chan *z.DiagItem)
		numdone, onjobdone := 0, func() { await <- nil }
		for _, pkg := range pkgs {
			for _, diagtool := range diagTools {
				numjobs++
				go me.runDiag(diagtool, pkg, await, onjobdone)
			}
			for _, filename := range pkg.GoFiles {
				if f, _ := workspaceFiles[filepath.Join(pkg.Dir, filename)]; f != nil {
					f.Diags.Lint.Forget(diagTools)
					f.Diags.Lint.UpToDate = true
				}
			}
		}

		var diagitems []*z.DiagItem
		for numdone < numjobs {
			select {
			case diagitem := <-await:
				if diagitem == nil {
					numdone++
				} else {
					diagitems = append(diagitems, diagitem)
				}
			}
		}
		for _, diag := range diagitems {
			f := workspaceFiles.Ensure(diag.FileRef.FilePath)
			f.Diags.Lint.UpToDate = true
			f.Diags.Lint.Items = append(f.Diags.Lint.Items, diag)
		}
	}
}
