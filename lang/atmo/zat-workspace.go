package zat

// import (
// 	"path/filepath"

// 	"github.com/go-leap/str"
// 	"github.com/metaleap/atmo/0ld"
// 	. "github.com/metaleap/atmo/0ld/ast"
// 	"github.com/metaleap/atmo/0ld/session"
// 	"github.com/metaleap/zentient"
// )

// var (
// 	workspace    atmoWorkspace
// 	goPathScopes []string
// )

// func init() {
// 	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
// }

// type atmoWorkspace struct {
// 	z.WorkspaceBase
// }

// func (*atmoWorkspace) onBeforeChanges(workspaceChanges *z.WorkspaceChanges, freshFiles []string, willAutoLint bool) {
// 	Ctx.Locked(func() {
// 		var kitimppaths []string
// 		gatherkits2refresh := func(ffps ...string) (lastkit *atmosess.Kit) {
// 			for _, ffp := range ffps {
// 				if filepath.Ext(ffp) == atmo.SrcFileExt {
// 					dirpath := filepath.Dir(ffp)
// 					if lastkit = Ctx.KitByDirPath(dirpath, true); lastkit != nil && !ustr.In(lastkit.ImpPath, kitimppaths...) {
// 						kitimppaths = append(kitimppaths, lastkit.ImpPath)
// 					}
// 				}
// 			}
// 			return
// 		}
// 		gatherkits2refresh(freshFiles...)
// 		gatherkits2refresh(workspaceChanges.WrittenFiles...)
// 		gatherkits2refresh(workspaceChanges.OpenedFiles...)

// 		var livesrcfiles []*AstFile
// 		if liveMode {
// 			updatesrcfile := func(srcfilepath string, srctxt []byte) {
// 				if kit := gatherkits2refresh(srcfilepath); kit != nil {
// 					srcfile := kit.SrcFiles.ByFilePath(srcfilepath)
// 					if srcfile == nil {
// 						Ctx.KitEnsureLoaded(kit)
// 						srcfile = kit.SrcFiles.ByFilePath(srcfilepath)
// 					}
// 					if srcfile != nil {
// 						srcfile.Options.TmpAltSrc = srctxt
// 						livesrcfiles = append(livesrcfiles, srcfile)
// 					}
// 				}
// 			}
// 			for _, srcfilepath := range workspaceChanges.WrittenFiles {
// 				updatesrcfile(srcfilepath, nil)
// 			}
// 			for srcfilepath, srctxt := range workspaceChanges.LiveFiles {
// 				updatesrcfile(srcfilepath, []byte(srctxt))
// 			}
// 		}
// 		Ctx.CatchUpOnFileMods(livesrcfiles...)
// 		Ctx.KitsEnsureLoaded(false, kitimppaths...)
// 	})
// }

// func (*atmoWorkspace) onAfterChanges(workspaceChanges *z.WorkspaceChanges) {
// }

// func (me *atmoWorkspace) onPreInit() {
// 	me.OnBeforeChanges, me.OnAfterChanges = me.onBeforeChanges, me.onAfterChanges
// }
