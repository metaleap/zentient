package zat

import (
	"path/filepath"
	"text/scanner"

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

			// happy smart path: already know the def(s) or def-arg the current name points to
			println("ONE")
			if irnodes := kit.AstNodeIrFunFor(tlc.Id(), nodes[0]); len(irnodes) > 0 {
				println("TWO")
				if ident, _ := irnodes[0].(*atmolang_irfun.AstIdentName); ident != nil && len(ident.Anns.ResolvesTo) > 0 {
					println("TRI", len(ident.Anns.ResolvesTo))
					for _, node := range ident.Anns.ResolvesTo {
						println("HUH", node.Print().Toks().String(), "HAH")
						def, _ := node.(*atmolang_irfun.AstDef)
						if def == nil {
							if deftop, _ := node.(*atmolang_irfun.AstDefTop); deftop != nil {
								def = &deftop.AstDef
							}
						}
						tok := node.OrigToks().First(nil)
						if def != nil {
							if t := def.Name.OrigToks().First(nil); t != nil {
								tok = t
							}
						}
						if tok != nil {
							locs.Add(tlc.SrcFile.SrcFilePath, &tok.Meta.Position)
						} else {
							locs.Add(tlc.SrcFile.SrcFilePath, &scanner.Position{Line: 1, Column: 1})
						}
					}
				}
			} else {

			}
			return

			// fall-back dumb path: traversal along the original src AST
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
