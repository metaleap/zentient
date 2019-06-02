package zat

import (
	"path/filepath"

	"github.com/go-leap/dev/lex"
	"github.com/go-leap/str"
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
		Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, func() {
			Ctx.KitEnsureLoaded(kit)
			if _, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {
				var refs map[*atmolang_irfun.AstDefTop][]atmolang_irfun.IAstExpr
				if ident, _ := nodes[0].(*atmolang.AstIdent); ident != nil {
					refs = Ctx.KitsCollectReferences(true, ident.Val)
				} else if atom, _ := nodes[0].(atmolang.IAstExprAtomic); atom != nil {
					refs = Ctx.KitsCollectReferences(true, atom.String())
				}
				for tld, nodes := range refs {
					for _, node := range nodes {
						if tok := node.OrigToks().First(nil); tok != nil {
							locs.Add(tld.OrigTopLevelChunk.SrcFile.SrcFilePath, &tok.Meta.Position)
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, func() {
			Ctx.KitEnsureLoaded(kit)
			if tlc, nodes := kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); len(nodes) > 0 {
				// HAPPY SMART PATH: already know the def(s) or def-arg the current name points to
				if irnodes := kit.AstNodeIrFunFor(tlc.Id(), nodes[0]); len(irnodes) > 0 {
					if ident, _ := irnodes[0].(*atmolang_irfun.AstIdentName); ident == nil {
						me.addLocFromNode(&locs, tlc.SrcFile.SrcFilePath, irnodes[0])
					} else {
						for _, node := range ident.Anns.ResolvesTo {
							me.addLocFromNode(&locs, tlc.SrcFile.SrcFilePath, node)
						}
					}
					return
				}

				// FALL-BACK DUMB PATH: merely lexical (traversal up the original src AST, collect any & all defs/def-args technically-in-scope and goal-named)
				if ident, _ := nodes[0].(*atmolang.AstIdent); ident != nil && ident.IsName(true) {
					// points to parent def-arg or def-in-scope?
					for i := 1; i < len(nodes); i++ {
						switch n := nodes[i].(type) {
						case *atmolang.AstDefArg:
							if nid, _ := n.NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
								me.addLocFromToks(&locs, tlc.SrcFile.SrcFilePath, nid.Tokens)
							}
						case *atmolang.AstDef:
							if n.Name.Val == ident.Val {
								me.addLocFromToks(&locs, tlc.SrcFile.SrcFilePath, n.Name.Tokens)
							} else {
								for da := range n.Args {
									if nid, _ := n.Args[da].NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
										me.addLocFromToks(&locs, tlc.SrcFile.SrcFilePath, nid.Tokens)
									}
								}
							}
						case *atmolang.AstExprLet:
							for d := range n.Defs {
								if n.Defs[d].Name.Val == ident.Val {
									me.addLocFromToks(&locs, tlc.SrcFile.SrcFilePath, n.Defs[d].Name.Tokens)
								}
							}
						}
					}
					// find all global goal-named defs --- loop isn't all that optimal but good-enough & succinct
					for _, k := range Ctx.Kits.All {
						if k.ImpPath == kit.ImpPath || ustr.In(k.ImpPath, kit.Imports...) {
							for _, def := range kit.Defs(ident.Val) {
								me.addLocFromNode(&locs, def.OrigTopLevelChunk.SrcFile.SrcFilePath, def)
							}
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) addLocFromToks(locs *z.SrcLocs, srcFilePath string, toks udevlex.Tokens) {
	if tok := toks.First(nil); tok != nil {
		locs.Add(srcFilePath, &tok.Meta.Position)
	}
}

func (me *atmoSrcIntel) addLocFromNode(locs *z.SrcLocs, srcFilePath string, node atmolang_irfun.IAstNode) {
	toks := node.OrigToks()
	if def := node.IsDef(); def != nil {
		if ts := def.Name.OrigToks(); len(ts) > 0 {
			toks = ts
		}
	}
	me.addLocFromToks(locs, srcFilePath, toks)
}
