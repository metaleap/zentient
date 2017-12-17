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

func (me *goDiag) OnUpdateLintDiags(workspaceFiles z.WorkspaceFiles, filePaths []string) (targets z.DiagTargets) {
	if pkgs := udevgo.PkgsForFiles(filePaths...); len(pkgs) > 0 {
		for _, pkg := range pkgs {
			target := &z.DiagTarget{Target: pkg}
			for _, filename := range pkg.GoFiles {
				target.AffectedFilePaths = append(target.AffectedFilePaths, filepath.Join(pkg.Dir, filename))
			}
			targets = append(targets, target)
		}
	}
	return
}

func (me *goDiag) RunDiag(tool *z.Tool, target *z.DiagTarget, yield z.DiagItemsChan) {
	defer yield.Done()
	pkg, diag := target.Target.(*udevgo.Pkg), func(i int, fpath string, found string) *z.DiagItem {
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
