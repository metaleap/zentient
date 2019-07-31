package zat

import (
	"fmt"
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
	Ctx.Locked(func() {
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
						if len(ret) > 0 {
							return
						}
					}

					// FALL-BACK DUMB PATH: merely lexical (traversal up the original src AST, collect any & all local defs/def-args technically-in-scope and goal-named)
					if ident, _ := astnodes[0].(*atmolang.AstIdent); ident == nil || !ident.IsName(true) {
						// not an ident, then point to input node itself
						me.addLocFromToks(curtlc, &ret, astnodes[0].Toks())
					} else {
						// points to local def-arg-in-scope or def-in-scope?
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
						// find all global goal-named defs -- loop isn't all that optimal but good-enough for interactive-and-only-occasional-use-case & succinct
						kitimports := kit.Imports()
						for _, k := range Ctx.Kits.All {
							if iscurkit := (k.ImpPath == kit.ImpPath); iscurkit || ustr.In(k.ImpPath, kitimports...) {
								if !iscurkit {
									Ctx.KitEnsureLoaded(k)
								}
								for _, srcfile := range k.SrcFiles {
									for i := range srcfile.TopLevel {
										if tlc := &srcfile.TopLevel[i]; tlc.Ast.Def.NameIfErr == ident.Val ||
											(tlc.Ast.Def.Orig != nil && tlc.Ast.Def.Orig.Name.Val == ident.Val) {
											toks := tlc.Ast.Tokens
											if tlc.Ast.Def.Orig != nil {
												toks = tlc.Ast.Def.Orig.Name.Tokens
											}
											me.addLocFromToks(tlc, &ret, toks)
										}
									}
								}
							}
						}
					}
				}
			})
		}
	})
	return
}

func (me *atmoSrcIntel) Highlights(srcLens *z.SrcLens, curWord string) (ret z.SrcLocs) {
	Ctx.Locked(func() {
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
							curfileonly := func(t *atmoil.IrDefTop) bool { return t.OrigTopChunk.SrcFile.SrcFilePath == srcLens.FilePath }
							var nodematches map[atmoil.INode]*atmoil.IrDefTop
							switch ilnode := ilnodes[0].(type) {
							case *atmoil.IrIdentDecl:
								nodematches = kit.SelectNodes(curfileonly, func(na []atmoil.INode, n atmoil.INode, nd []atmoil.INode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
									if nid, _ := n.(*atmoil.IrIdentName); nid != nil {
										ismatch = nid.ResolvesTo(ilnodes[1])
									}
									return
								})
								nodematches[ilnode] = tld
							case *atmoil.IrIdentName:
								nodematches = kit.SelectNodes(curfileonly, func(na []atmoil.INode, n atmoil.INode, nd []atmoil.INode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
									for _, cand := range ilnode.Anns.Candidates {
										if ismatch = (cand == n); ismatch {
											break
										}
									}
									if nid, _ := n.(*atmoil.IrIdentName); nid != nil && !ismatch {
										for _, cand := range ilnode.Anns.Candidates {
											if ismatch = nid.ResolvesTo(cand); ismatch {
												break
											}
										}
									}
									return
								})
								nodematches[ilnode] = tld
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
	})
	return
}

func (me *atmoSrcIntel) Hovers(srcLens *z.SrcLens) (ret []z.SrcInfoTip) {
	Ctx.Locked(func() {
		if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
			me.withInMemFileMod(srcLens, kit, func() {
				if tlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
					var nodetypenames string
					for _, n := range astnodes {
						nodetypenames += z.Strf("─ %T ", n)
					}
					ret = append(ret,
						z.SrcInfoTip{Value: tlc.Ast.Def.Orig.Name.Val},
						z.SrcInfoTip{Value: nodetypenames[4:]},
					)

					if tld, ilnodes := kit.IrNodeOfAstNode(tlc.Id(), astnodes[0]); len(ilnodes) > 0 {
						if tld.Anns.Preduced != nil {
							ret = append(ret, z.SrcInfoTip{Value: "≡\n" + tld.Anns.Preduced.SummaryCompact() + ">>>" + fmt.Sprintf("%T", ilnodes[0])})
						}
						if nid, _ := ilnodes[0].(*atmoil.IrIdentName); nid != nil {
							ret = append(ret, z.SrcInfoTip{Value: z.Strf("(resolves to %v candidate/s)", len(nid.Anns.Candidates))})
						}
					}
				}
			})
		}
	})
	return
}

func (me *atmoSrcIntel) Annotactions(srcLens *z.SrcLens) (ret []*z.SrcAnnotaction) {
	ret = append(ret, &z.SrcAnnotaction{
		Title: "Foo Bar Baz", Desc: "It's yet another lens test",
		Range: z.SrcRange{Start: z.SrcPos{Ln: 8, Col: 1}, End: z.SrcPos{Ln: 8, Col: 8}},
	})
	return
}

func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (ret z.SrcLocs) {
	Ctx.Locked(func() {
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
							if tok := tld.OrigToks(node).First1(); tok != nil {
								ret.Add(tld.OrigTopChunk.SrcFile.SrcFilePath, tok.OffPos(tld.OrigTopChunk.PosOffsetLine(), tld.OrigTopChunk.PosOffsetByte()))
							}
						}
					}
				}
			})
		}
	})
	return
}

func (me *atmoSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (ret z.SrcLenses) {
	Ctx.Locked(func() {
		query = ustr.Lo(query)
		var kits atmosess.Kits
		curkit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true)
		if kits = Ctx.Kits.All; curFileOnly {
			kits = atmosess.Kits{curkit}
		}
		symbolForTopLevelDef := func(tlc *atmolang.SrcTopChunk) {
			var name string
			flag := z.SYM_FUNCTION
			if tld := tlc.Ast.Def.Orig; tld != nil {
				if len(query) == 0 || ustr.Has(ustr.Lo(tld.Name.Val), query) {
					name = tld.Name.Val
					if len(tld.Args) == 0 {
						flag = z.SYM_FIELD
					} else if tld.Name.IsOpish {
						flag = z.SYM_OPERATOR
					}
				}
			} else if tlc.Ast.Def.NameIfErr != "" {
				name, flag =
					tlc.Ast.Def.NameIfErr, z.SYM_EVENT
			}
			if name != "" {
				ret = append(ret, &z.SrcLens{Str: name,
					Txt: "(description later)", SrcLoc: z.SrcLoc{
						FilePath: tlc.SrcFile.SrcFilePath,
						Flag:     int(flag),
						Range:    toksToRange(tlc, tlc.Ast.Tokens)}})
			}
		}
		if len(kits) > 0 && curkit != nil {
			me.withInMemFileMod(srcLens, curkit, func() {
				for _, kit := range kits {
					for _, srcfile := range kit.SrcFiles {
						if (!curFileOnly) || srcfile.SrcFilePath == srcLens.FilePath {
							for i := range srcfile.TopLevel {
								symbolForTopLevelDef(&srcfile.TopLevel[i])
							}
						}
					}
				}
			})
		}
	})
	return
}

func (me *atmoSrcIntel) addLocFromToks(tlc *atmolang.SrcTopChunk, dst *z.SrcLocs, toks udevlex.Tokens) *z.SrcLoc {
	if tok := toks.First1(); tok != nil {
		pos := tok.OffPos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
		return dst.Add(tlc.SrcFile.SrcFilePath, pos)
	}
	return nil
}

func (me *atmoSrcIntel) addLocFromNode(tld *atmoil.IrDefTop, dst *z.SrcLocs, node atmoil.INode) *z.SrcLoc {
	toks := tld.OrigToks(node)
	if def := node.IsDef(); def != nil {
		if ts := tld.OrigToks(&def.Name); len(ts) > 0 {
			toks = ts
		}
	}
	return me.addLocFromToks(tld.OrigTopChunk, dst, toks)
}

func (me *atmoSrcIntel) astAt(kit *atmosess.Kit, srcLens *z.SrcLens) (topLevelChunk *atmolang.SrcTopChunk, theNodeAndItsAncestors []atmolang.IAstNode) {
	if topLevelChunk, theNodeAndItsAncestors = kit.AstNodeAt(srcLens.FilePath, srcLens.ByteOffsetForPos(srcLens.Pos)); topLevelChunk != nil && topLevelChunk.Ast.Def.Orig == nil {
		theNodeAndItsAncestors = nil
	}
	return
}

func tokToPos(tlc *atmolang.SrcTopChunk, tok *udevlex.Token) (ret *z.SrcPos) {
	pos := tok.OffPos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
	ret = &z.SrcPos{Ln: pos.Ln1, Col: pos.Col1}
	ret.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)
	return
}

func toksToRange(tlc *atmolang.SrcTopChunk, toks udevlex.Tokens) *z.SrcRange {
	var sr z.SrcRange
	tlcoffl, tlcoffb := tlc.PosOffsetLine(), tlc.PosOffsetByte()

	pos := toks.First1().OffPos(tlcoffl, tlcoffb)
	sr.Start.Col, sr.Start.Ln = pos.Col1, pos.Ln1
	sr.Start.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)

	pos = toks.Last1().OffPosEnd(tlcoffl, tlcoffb)
	sr.End.Col, sr.End.Ln = pos.Col1, pos.Ln1
	sr.End.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)

	return &sr
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
