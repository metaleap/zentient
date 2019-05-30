package zat

import (
	"path/filepath"

	"github.com/metaleap/zentient"
)

var srcIntel atmoSrcIntel

type atmoSrcIntel struct {
	z.SrcIntelBase
}

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		if _, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {

		}
	}
	return
}
