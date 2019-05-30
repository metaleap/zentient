package zat

import (
	"fmt"
	"path/filepath"

	"github.com/metaleap/atmo/lang"
	"github.com/metaleap/atmo/lang/irfun"
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
		Ctx.KitEnsureLoaded(kit)
		if _, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {
			if ident, _ := nodes[0].(*atmolang.AstIdent); ident != nil && ident.IsName(true) {
				for tld, nodes := range Ctx.KitsCollectReferences(true, ident.Val) {
					for _, node := range nodes {
						if tok := node.OrigToks().First(nil); tok != nil {
							locs.Add(tld.OrigTopLevelChunk.SrcFile.SrcFilePath, &tok.Meta.Position)
						}
					}
				}
			}
		}
	}
	return
}

func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		Ctx.KitEnsureLoaded(kit)
		if tlc, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {
			if irnodes := kit.AstNodeIrFunFor(tlc.Id(), nodes[0]); len(irnodes) > 0 {
				if ident, _ := irnodes[0].(*atmolang_irfun.AstIdentName); ident != nil {
					z.SendNotificationMessageToClient(2, fmt.Sprintf("%v", len(ident.Anns.ResolvesTo)))
				}
			}

			if ident, _ := nodes[0].(*atmolang.AstIdent); ident != nil && ident.IsName(true) {
				// points to parent def-arg or def-in-scope?
				for i := 1; i < len(nodes); i++ {
					switch n := nodes[i].(type) {
					case *atmolang.AstDefArg:
						if nid, _ := n.NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
							locs.Add(tlc.SrcFile.SrcFilePath, &nid.Tokens[0].Meta.Position)
						}
					case *atmolang.AstDef:
						if n.Name.Val == ident.Val {
							locs.Add(tlc.SrcFile.SrcFilePath, &n.Name.Tokens[0].Meta.Position)
						} else {
							for da := range n.Args {
								if nid, _ := n.Args[da].NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
									locs.Add(tlc.SrcFile.SrcFilePath, &nid.Tokens[0].Meta.Position)
								}
							}
						}
					case *atmolang.AstExprLet:
						for d := range n.Defs {
							if n.Defs[d].Name.Val == ident.Val {
								locs.Add(tlc.SrcFile.SrcFilePath, &n.Defs[d].Name.Tokens[0].Meta.Position)
							}
						}
					}
				}

				// find all global goal-named defs
				for _, kit := range Ctx.Kits.All {
					for _, def := range kit.Defs(ident.Val) {
						if tok := def.Name.OrigToks().First(nil); tok != nil {
							locs.Add(def.OrigTopLevelChunk.SrcFile.SrcFilePath, &tok.Meta.Position)
						}
					}
				}
			}
		}
	}
	return
}
