package zat

import (
	"path/filepath"

	"github.com/go-leap/dev/lex"
	"github.com/go-leap/str"
	"github.com/metaleap/atmo/lang"
	"github.com/metaleap/atmo/lang/irfun"
	"github.com/metaleap/atmo/session"
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
		if panicked := Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, func() {
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
		}); panicked != nil {
			panic(panicked)
		}
	}
	return
}

func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		if panicked := Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, func() {
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
		}); panicked != nil {
			panic(panicked)
		}
	}
	return
}

func (me *atmoSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (allsyms z.SrcLenses) {
	query = ustr.Lo(query)
	var kits atmosess.Kits
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); !curFileOnly {
		kits = Ctx.Kits.All
	} else if kit != nil {
		kits = atmosess.Kits{kit}
		Ctx.KitEnsureLoaded(kit)
	}
	if panicked := Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, func() {
		for _, kit := range kits {
			for _, srcfile := range kit.SrcFiles {
				if srcfile.SrcFilePath == srcLens.FilePath || !curFileOnly {
					for i := range srcfile.TopLevel {
						if def := srcfile.TopLevel[i].Ast.Def.Orig; def != nil {
							if len(query) == 0 || ustr.Has(ustr.Lo(def.Name.Val), query) {
								allsyms = append(allsyms, &z.SrcLens{Str: def.Name.Val, Txt: "(description later)", SrcLoc: z.SrcLoc{
									FilePath: srcfile.SrcFilePath, Flag: int(z.SYM_FUNCTION), Range: toksToRange(def.Tokens)}})
							}
						}
					}
				}
			}
		}
	}); panicked != nil {
		panic(panicked)
	}
	return
}

func (me *atmoSrcIntel) addLocFromToks(locs *z.SrcLocs, srcFilePath string, toks udevlex.Tokens) *z.SrcLoc {
	if tok := toks.First(nil); tok != nil {
		return locs.Add(srcFilePath, &tok.Meta.Position)
	}
	return nil
}

func (me *atmoSrcIntel) addLocFromNode(locs *z.SrcLocs, srcFilePath string, node atmolang_irfun.IAstNode) *z.SrcLoc {
	toks := node.OrigToks()
	if def := node.IsDef(); def != nil {
		if ts := def.Name.OrigToks(); len(ts) > 0 {
			toks = ts
		}
	}
	return me.addLocFromToks(locs, srcFilePath, toks)
}

func tokToPos(tok *udevlex.Token) *z.SrcPos {
	return &z.SrcPos{Off: tok.Meta.Offset + 1, Ln: tok.Meta.Line, Col: tok.Meta.Column}
}

func toksToRange(toks udevlex.Tokens) (sr *z.SrcRange) {
	sr = &z.SrcRange{Start: *tokToPos(toks.First(nil))}
	tok := toks.Last(nil)
	l := len(tok.Meta.Orig)
	sr.End.Off, sr.End.Ln, sr.End.Col = tok.Meta.Offset+l, tok.Meta.Line, tok.Meta.Column+l
	return
}
