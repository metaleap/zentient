package zat

import (
	"path/filepath"

	"github.com/go-leap/dev/lex"
	"github.com/go-leap/str"
	"github.com/metaleap/atmo/il"
	"github.com/metaleap/atmo/lang"
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

func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (ret z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if curtlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
				// HAPPY SMART PATH: already know the def(s) or def-arg the current name points to
				if curtld, ilnodes := kit.IrNodeOfAstNode(curtlc.Id(), astnodes[0]); len(ilnodes) > 0 {
					if ident, _ := ilnodes[0].(*atmoil.IrIdentName); ident != nil {
						for _, node := range ident.Anns.Candidates {
							tld, _ := node.(*atmoil.IrDefTop)
							if tld == nil {
								tld = curtld
								if adr, ok := node.(atmosess.IrDefRef); ok {
									tld = adr.IrDefTop
								}
							}
							me.addLocFromNode(tld, &ret, node)
						}
					} else { // not an ident, then point to input node itself
						me.addLocFromNode(curtld, &ret, ilnodes[0])
					}
					return
				}

				// FALL-BACK DUMB PATH: merely lexical (traversal up the original src AST, collect any & all defs/def-args technically-in-scope and goal-named)
				if ident, _ := astnodes[0].(*atmolang.AstIdent); ident != nil && ident.IsName(true) {
					// points to parent def-arg or def-in-scope?
					for i := 1; i < len(astnodes); i++ {
						switch n := astnodes[i].(type) {
						case *atmolang.AstDefArg:
							if nid, _ := n.NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
								me.addLocFromToks(curtlc, &ret, nid.Tokens)
							}
						case *atmolang.AstDef:
							if n.Name.Val == ident.Val {
								me.addLocFromToks(curtlc, &ret, n.Name.Tokens)
							} else {
								for da := range n.Args {
									if nid, _ := n.Args[da].NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
										me.addLocFromToks(curtlc, &ret, nid.Tokens)
									}
								}
							}
						case *atmolang.AstExprLet:
							for d := range n.Defs {
								if n.Defs[d].Name.Val == ident.Val {
									me.addLocFromToks(curtlc, &ret, n.Defs[d].Name.Tokens)
								}
							}
						}
					}
					// find all global goal-named defs --- loop isn't all that optimal but good-enough & succinct
					for _, k := range Ctx.Kits.All {
						if k.ImpPath == kit.ImpPath || ustr.In(k.ImpPath, kit.Imports...) {
							for _, def := range kit.Defs(ident.Val) {
								me.addLocFromNode(def, &ret, def)
							}
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) Highlights(srcLens *z.SrcLens, curWord string) (ret z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if tlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
				showscope := func(n atmolang.IAstNode) { me.addLocFromToks(tlc, &ret, n.Toks()) }
				switch astnode := astnodes[0].(type) {
				case *atmolang.AstExprLet:
					showscope(astnode)
				case *atmolang.AstTopLevel:
					showscope(astnode)
				case *atmolang.AstDef:
					showscope(astnode)
				case *atmolang.AstExprAppl:
					showscope(astnode)
				case *atmolang.AstExprCases:
					showscope(astnode)
				default:
					if len(astnodes) > 1 {
						if astarg, _ := astnodes[1].(*atmolang.AstDefArg); astarg != nil {
							// TODO: handle astarg.Affix or non-name astarg.NameOrConstVal
						}
					}
					if tld, ilnodes := kit.IrNodeOfAstNode(tlc.Id(), astnode); len(ilnodes) > 0 {
						curfileonly := func(t *atmoil.IrDefTop) bool { return t.OrigTopLevelChunk.SrcFile.SrcFilePath == srcLens.FilePath }
						var nodematches map[atmoil.INode]*atmoil.IrDefTop
						switch ilnode := ilnodes[0].(type) {
						case *atmoil.IrIdentDecl:
							nodematches = kit.SelectNodes(curfileonly, func(na []atmoil.INode, n atmoil.INode, nd []atmoil.INode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
								if nid, _ := n.(*atmoil.IrIdentName); nid != nil {
									ismatch = nid.ResolvesToPotentially(ilnodes[1])
								}
								return
							})
							if len(nodematches) > 0 {
								nodematches[ilnode] = tld
							}
						case *atmoil.IrIdentName:
						default:
							nodematches = kit.SelectNodes(curfileonly, func(na []atmoil.INode, n atmoil.INode, nd []atmoil.INode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
								ismatch = (ilnode == n || ilnode.EquivTo(n))
								return
							})
						}
						for mnode, mtld := range nodematches {
							me.addLocFromNode(mtld, &ret, mnode)
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) Hovers(srcLens *z.SrcLens) (ret []z.InfoTip) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if tlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
				var nodetypenames string
				for _, n := range astnodes {
					nodetypenames += z.Strf("— %T ", n)
				}
				ret = append(ret,
					z.InfoTip{Value: tlc.Ast.Def.Orig.Name.Val},
					z.InfoTip{Value: nodetypenames[4:]},
				)

				if _, ilnodes := kit.IrNodeOfAstNode(tlc.Id(), astnodes[0]); len(ilnodes) > 0 {
					for _, n := range ilnodes {
						if nid, _ := n.(*atmoil.IrIdentName); nid != nil {
							ret = append(ret, z.InfoTip{Value: z.Strf("(resolves to %v candidate/s)", len(nid.Anns.Candidates))})
						}
						ret = append(ret,
							z.InfoTip{Value: z.Strf("%T:\n%s", n, n.Facts().Description()), Language: "plain"},
						)
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (ret z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if _, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
				var refs map[*atmoil.IrDefTop][]atmoil.IExpr
				if ident, _ := astnodes[0].(*atmolang.AstIdent); ident != nil {
					refs = Ctx.KitsCollectReferences(true, ident.Val)
				} else if atom, _ := astnodes[0].(atmolang.IAstExprAtomic); atom != nil {
					refs = Ctx.KitsCollectReferences(true, atom.String())
				}
				for tld, ilnodes := range refs {
					for _, node := range ilnodes {
						if tok := tld.OrigToks(node).First(nil); tok != nil {
							ret.Add(tld.OrigTopLevelChunk.SrcFile.SrcFilePath, tok.Pos(tld.OrigTopLevelChunk.PosOffsetLine(), tld.OrigTopLevelChunk.PosOffsetByte()))
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (ret z.SrcLenses) {
	query = ustr.Lo(query)
	var kits atmosess.Kits
	curkit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true)
	if kits = Ctx.Kits.All; curFileOnly {
		kits = atmosess.Kits{curkit}
	}
	if len(kits) > 0 && curkit != nil {
		me.withInMemFileMod(srcLens, curkit, func() {
			for _, kit := range kits {
				for _, srcfile := range kit.SrcFiles {
					if srcfile.SrcFilePath == srcLens.FilePath || !curFileOnly {
						for i := range srcfile.TopLevel {
							tlc := &srcfile.TopLevel[i]
							if def := tlc.Ast.Def.Orig; def != nil {
								if len(query) == 0 || ustr.Has(ustr.Lo(def.Name.Val), query) {
									ret = append(ret, &z.SrcLens{Str: def.Name.Val, Txt: "(description later)", SrcLoc: z.SrcLoc{
										FilePath: srcfile.SrcFilePath, Flag: int(z.SYM_FUNCTION), Range: toksToRange(tlc, def.Tokens)}})
								}
							}
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) addLocFromToks(tlc *atmolang.SrcTopChunk, locs *z.SrcLocs, toks udevlex.Tokens) *z.SrcLoc {
	if tok := toks.First(nil); tok != nil {
		return locs.Add(tlc.SrcFile.SrcFilePath, tok.Pos(tlc.PosOffsetLine(), tlc.PosOffsetByte()))
	}
	return nil
}

func (me *atmoSrcIntel) addLocFromNode(tld *atmoil.IrDefTop, locs *z.SrcLocs, node atmoil.INode) *z.SrcLoc {
	toks := tld.OrigToks(node)
	if def := node.IsDef(); def != nil {
		if ts := tld.OrigToks(&def.Name); len(ts) > 0 {
			toks = ts
		}
	}
	return me.addLocFromToks(tld.OrigTopLevelChunk, locs, toks)
}

func (me *atmoSrcIntel) astAt(kit *atmosess.Kit, srcLens *z.SrcLens) (topLevelChunk *atmolang.SrcTopChunk, theNodeAndItsAncestors []atmolang.IAstNode) {
	if topLevelChunk, theNodeAndItsAncestors = kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); topLevelChunk != nil && topLevelChunk.Ast.Def.Orig == nil {
		theNodeAndItsAncestors = nil
	}
	return
}

func tokToPos(tlc *atmolang.SrcTopChunk, tok *udevlex.Token) *z.SrcPos {
	pos := tok.Pos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
	return &z.SrcPos{Off: pos.Offset + 1, Ln: pos.Line, Col: pos.Column}
}

func toksToRange(tlc *atmolang.SrcTopChunk, toks udevlex.Tokens) (sr *z.SrcRange) {
	sr = &z.SrcRange{Start: *tokToPos(tlc, toks.First(nil))}
	tok := toks.Last(nil)
	l, pos := len(tok.Meta.Orig), tok.Pos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
	sr.End.Off, sr.End.Ln, sr.End.Col = pos.Offset+l, pos.Line, pos.Column+l
	return
}

func (me *atmoSrcIntel) withInMemFileMod(srcLens *z.SrcLens, kit *atmosess.Kit, do func()) {
	Ctx.KitEnsureLoaded(kit)
	if srcLens.Txt == "" {
		do()
	} else if liveMode {
		if srcfile := kit.SrcFiles.ByFilePath(srcLens.FilePath); srcfile != nil {
			srcfile.Options.TmpAltSrc = []byte(srcLens.Txt)
			Ctx.CatchUpOnFileMods(srcfile)
		}
		do()
	} else if panicked := Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, do); panicked != nil {
		panic(panicked)
	}
}
