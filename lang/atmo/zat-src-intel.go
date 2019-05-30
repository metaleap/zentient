package zat

import (
	"fmt"
	"path/filepath"

	// "github.com/metaleap/atmo"
	"github.com/metaleap/atmo/lang"
	"github.com/metaleap/zentient"
)

var srcIntel atmoSrcIntel

type atmoSrcIntel struct {
	z.SrcIntelBase
}

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		if tlc, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {
			if tok := nodes[len(nodes)-1].Toks().First(nil); tok != nil {
				locs = append(locs, &z.SrcLoc{FilePath: tlc.SrcFile.SrcFilePath,
					Pos: &z.SrcPos{Off: tok.Meta.Offset + 1, Ln: tok.Meta.Line, Col: tok.Meta.Column}})
			}

			if ident, _ := nodes[len(nodes)-1].(*atmolang.AstIdent); ident != nil && ident.IsName(true) {
				for _, kit := range Ctx.Kits.All {
					for _, def := range kit.Defs(ident.Val) {
						if tok := def.Name.OrigToks().First(nil); tok != nil {
							locs = append(locs, &z.SrcLoc{FilePath: def.OrigTopLevelChunk.SrcFile.SrcFilePath,
								Pos: &z.SrcPos{Off: tok.Meta.Offset + 1, Ln: tok.Meta.Line, Col: tok.Meta.Column}})
						}
					}
				}
			}
		}
	} else {
		z.SendNotificationMessageToClient(z.DIAG_SEV_INFO, fmt.Sprintf("%v", len(Ctx.Kits.All)))
	}
	return
}
