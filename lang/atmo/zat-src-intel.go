package zat

// import (
// 	"fmt"
// 	"path/filepath"

// 	"github.com/go-leap/dev/lex"
// 	"github.com/go-leap/str"
// 	. "github.com/metaleap/atmo/0ld/ast"
// 	. "github.com/metaleap/atmo/0ld/il"
// 	"github.com/metaleap/atmo/0ld/session"
// 	"github.com/metaleap/zentient"
// )

// var srcIntel atmoSrcIntel

// type atmoSrcIntel struct {
// 	z.SrcIntelBase
// }

// func init() {
// 	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
// }

// func (me *atmoSrcIntel) ComplItems(srcLens *z.SrcLens) (complItems z.SrcIntelCompls) {
// 	me.withCtxLockedAndKitAndInMemFileMod(srcLens, func(kit *atmosess.Kit) {
// 		namesinscope := kit.NamesInScope()
// 		if curtlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
// 			if curtld, ilnodes := kit.IrNodeOfAstNode(curtlc.Id(), astnodes[0]); len(ilnodes) > 0 {
// 				namesinscope = curtld.NamesInScopeAt(ilnodes[0], namesinscope, true)
// 			}
// 		}
// 		for name, nodes := range namesinscope {
// 			complItems = append(complItems, &z.SrcIntelCompl{
// 				Kind:          z.CMPL_COLOR,
// 				Documentation: &z.SrcIntelDoc{IsTrusted: true, Value: ustr.Plu(len(nodes), "doc")},
// 				Detail:        ustr.Plu(len(nodes), "detail"),
// 				Label:         name,
// 				SortText:      name,
// 			})
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) ComplDetails(srcLens *z.SrcLens, itemText string) *z.SrcIntelCompl {
// 	return me.SrcIntelBase.ComplDetails(srcLens, itemText)
// }

// func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (ret z.SrcLocs) {
// 	me.withCtxLockedAndKitAndInMemFileMod(srcLens, func(kit *atmosess.Kit) {
// 		if curtlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
// 			// HAPPY SMART PATH: already know the def(s) or def-arg the current name points to
// 			if curtld, ilnodes := kit.IrNodeOfAstNode(curtlc.Id(), astnodes[0]); len(ilnodes) > 0 {
// 				if ident, _ := ilnodes[0].(*IrIdentName); ident != nil {
// 					for _, cand := range ident.Ann.Candidates {
// 						tld, _ := cand.(*IrDef)
// 						if tld == nil {
// 							tld = curtld
// 							if defref, ok := cand.(atmosess.IrDefRef); ok {
// 								tld = defref.IrDef
// 							}
// 						}
// 						me.addLocFromNode(tld, &ret, cand)
// 					}
// 				} else { // not an ident, then point to input node itself
// 					me.addLocFromNode(curtld, &ret, ilnodes[0])
// 				}
// 				if len(ret) > 0 {
// 					return
// 				}
// 			}

// 			// FALL-BACK DUMB PATH: merely lexical (traversal up the original src AST, collect any & all local defs/def-args technically-in-scope and goal-named)
// 			if ident, _ := astnodes[0].(*AstIdent); ident == nil || !ident.IsName(true) {
// 				// not an ident, then point to input node itself
// 				me.addLocFromToks(curtlc, &ret, astnodes[0].Toks())
// 			} else {
// 				// points to local def-arg-in-scope or def-in-scope?
// 				for i := 1; i < len(astnodes); i++ {
// 					switch n := astnodes[i].(type) {
// 					case *AstDefArg:
// 						if nid, _ := n.NameOrConstVal.(*AstIdent); nid != nil && nid.Val == ident.Val {
// 							me.addLocFromToks(curtlc, &ret, nid.Tokens)
// 						}
// 					case *AstDef:
// 						if n.Name.Val == ident.Val {
// 							me.addLocFromToks(curtlc, &ret, n.Name.Tokens)
// 						} else {
// 							for da := range n.Args {
// 								if nid, _ := n.Args[da].NameOrConstVal.(*AstIdent); nid != nil && nid.Val == ident.Val {
// 									me.addLocFromToks(curtlc, &ret, nid.Tokens)
// 								}
// 							}
// 						}
// 					case *AstExprLet:
// 						for d := range n.Defs {
// 							if n.Defs[d].Name.Val == ident.Val {
// 								me.addLocFromToks(curtlc, &ret, n.Defs[d].Name.Tokens)
// 							}
// 						}
// 					}
// 				}
// 				// find all global goal-named defs -- loop isn't all that optimal but good-enough for interactive-and-only-occasional-use-case & succinct
// 				kitimports := kit.Imports()
// 				for _, k := range Ctx.Kits.All {
// 					if iscurkit := (k.ImpPath == kit.ImpPath); iscurkit || ustr.In(k.ImpPath, kitimports...) {
// 						if !iscurkit {
// 							Ctx.KitEnsureLoaded(k)
// 						}
// 						for _, srcfile := range k.SrcFiles {
// 							for i := range srcfile.TopLevel {
// 								if tlc := &srcfile.TopLevel[i]; tlc.Ast.Def.NameIfErr == ident.Val ||
// 									(tlc.Ast.Def.Orig != nil && tlc.Ast.Def.Orig.Name.Val == ident.Val) {
// 									toks := tlc.Ast.Tokens
// 									if tlc.Ast.Def.Orig != nil {
// 										toks = tlc.Ast.Def.Orig.Name.Tokens
// 									}
// 									me.addLocFromToks(tlc, &ret, toks)
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) Highlights(srcLens *z.SrcLens, curWord string) (ret z.SrcLocs) {
// 	me.withCtxLockedAndKitAndInMemFileMod(srcLens, func(kit *atmosess.Kit) {
// 		if tlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
// 			showscope := func(n IAstNode) { me.addLocFromToks(tlc, &ret, n.Toks()) }
// 			switch astnode := astnodes[0].(type) {
// 			case *AstExprLet:
// 				showscope(astnode)
// 			case *AstTopLevel:
// 				showscope(astnode)
// 			case *AstDef:
// 				showscope(astnode)
// 			case *AstExprAppl:
// 				showscope(astnode)
// 			case *AstExprCases:
// 				showscope(astnode)
// 			default:
// 				if tld, ilnodes := kit.IrNodeOfAstNode(tlc.Id(), astnode); len(ilnodes) > 0 {
// 					curfileonly := func(t *IrDef) bool { return t.AstFileChunk.SrcFile.SrcFilePath == srcLens.FilePath }
// 					var nodematches map[IIrNode]*IrDef
// 					switch ilnode := ilnodes[0].(type) {
// 					case *IrIdentDecl:
// 						nodematches = kit.SelectNodes(curfileonly, func(na []IIrNode, n IIrNode, nd []IIrNode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
// 							nid, _ := n.(*IrIdentName)
// 							ismatch = (nid != nil) && nid.ResolvesTo(ilnodes[1])
// 							return
// 						})
// 						nodematches[ilnode] = tld
// 					case *IrIdentName:
// 						nodematches = kit.SelectNodes(curfileonly, func(na []IIrNode, n IIrNode, nd []IIrNode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
// 							ismatch = ilnode.ResolvesTo(n)
// 							if nid, _ := n.(*IrIdentName); nid != nil && !ismatch {
// 								for _, cand := range ilnode.Ann.Candidates {
// 									if ismatch = nid.ResolvesTo(cand); ismatch {
// 										break
// 									}
// 								}
// 							}
// 							return
// 						})
// 						nodematches[ilnode] = tld
// 					default:
// 						nodematches = kit.SelectNodes(curfileonly, func(na []IIrNode, n IIrNode, nd []IIrNode) (ismatch bool, dontdescend bool, donetld bool, doneall bool) {
// 							ismatch = (ilnode == n || ilnode.EquivTo(n, false))
// 							return
// 						})
// 					}
// 					for mnode, mtld := range nodematches {
// 						me.addLocFromNode(mtld, &ret, mnode)
// 					}
// 				}
// 			}
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) Hovers(srcLens *z.SrcLens) (ret []z.SrcInfoTip) {
// 	me.withCtxLockedAndKitAndInMemFileMod(srcLens, func(kit *atmosess.Kit) {
// 		if tlc, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
// 			var nodetypenames string
// 			for _, n := range astnodes {
// 				nodetypenames += z.Strf("─ %T ", n)
// 			}
// 			ret = append(ret,
// 				z.SrcInfoTip{Value: tlc.Ast.Def.Orig.Name.Val},
// 				z.SrcInfoTip{Value: nodetypenames[4:]},
// 			)

// 			if tld, ilnodes := kit.IrNodeOfAstNode(tlc.Id(), astnodes[0]); len(ilnodes) > 0 {
// 				prednode := ilnodes[0]
// 				if decl, _ := prednode.(*IrIdentDecl); decl != nil {
// 					prednode = ilnodes[1]
// 				}
// 				if pred := Ctx.Preduce(kit, tld, prednode); pred != nil {
// 					ret = append(ret, z.SrcInfoTip{Value: "≡\n" + pred.String()})
// 				} else {
// 					ret = append(ret, z.SrcInfoTip{Value: "?\n" + fmt.Sprintf("%T", prednode)})
// 				}
// 				if nid, _ := ilnodes[0].(*IrIdentName); nid != nil {
// 					ret = append(ret, z.SrcInfoTip{Value: z.Strf("(resolves to %v candidate/s)", len(nid.Ann.Candidates))})
// 				}
// 			}
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) Annotactions(srcLens *z.SrcLens) (ret []*z.SrcAnnotaction) {
// 	ret = append(ret, &z.SrcAnnotaction{
// 		Title: "Foo Bar Baz", Desc: "It's yet another lens test",
// 		Range: z.SrcRange{Start: z.SrcPos{Ln: 8, Col: 1}, End: z.SrcPos{Ln: 8, Col: 8}},
// 	})
// 	return
// }

// func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (ret z.SrcLocs) {
// 	me.withCtxLockedAndKitAndInMemFileMod(srcLens, func(kit *atmosess.Kit) {
// 		if _, astnodes := me.astAt(kit, srcLens); len(astnodes) > 0 {
// 			var refs map[*IrDef][]IIrExpr
// 			if ident, _ := astnodes[0].(*AstIdent); ident != nil {
// 				refs = Ctx.KitsCollectReferences(true, ident.Val)
// 			} else if atom, _ := astnodes[0].(IAstExprAtomic); atom != nil {
// 				refs = Ctx.KitsCollectReferences(true, atom.String())
// 			}
// 			for tld, ilnodes := range refs {
// 				for _, node := range ilnodes {
// 					if tok := tld.AstOrigToks(node).First1(); tok != nil {
// 						ret.Add(tld.AstFileChunk.SrcFile.SrcFilePath, tok.OffPos(tld.AstFileChunk.PosOffsetLine(), tld.AstFileChunk.PosOffsetByte()))
// 					}
// 				}
// 			}
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (ret z.SrcLenses) {
// 	Ctx.Locked(func() {
// 		query = ustr.Lo(query)
// 		var kits atmosess.Kits
// 		curkit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true)
// 		if kits = Ctx.Kits.All; curFileOnly {
// 			kits = atmosess.Kits{curkit}
// 		}
// 		symbolForDef := func(tlc *AstFileChunk, def *AstDef) {
// 			var name string
// 			flag, usetopleveldef := z.SYM_FUNCTION, (def == nil)
// 			if usetopleveldef {
// 				def = tlc.Ast.Def.Orig
// 			}
// 			if def != nil {
// 				if len(query) == 0 || ustr.Has(ustr.Lo(def.Name.Val), query) {
// 					name = def.Name.Val
// 					if !usetopleveldef {
// 						flag = z.SYM_OBJECT
// 					} else if len(def.Args) == 0 {
// 						flag = z.SYM_FIELD
// 					} else if def.Name.IsOpish {
// 						flag = z.SYM_OPERATOR
// 					}
// 				}
// 			} else if usetopleveldef && tlc.Ast.Def.NameIfErr != "" {
// 				name, flag =
// 					tlc.Ast.Def.NameIfErr, z.SYM_EVENT
// 			}
// 			if name != "" {
// 				toks := tlc.Ast.Tokens
// 				if !usetopleveldef {
// 					toks = def.Tokens
// 				}
// 				ret = append(ret, &z.SrcLens{Str: name,
// 					Txt: "(description later)", SrcLoc: z.SrcLoc{
// 						FilePath: tlc.SrcFile.SrcFilePath,
// 						Flag:     int(flag),
// 						Range:    toksToRange(tlc, toks)}})
// 			}
// 		}
// 		if len(kits) > 0 && curkit != nil {
// 			me.withInMemFileMod(srcLens, curkit, func() {
// 				for _, kit := range kits {
// 					for _, srcfile := range kit.SrcFiles {
// 						if (!curFileOnly) || srcfile.SrcFilePath == srcLens.FilePath {
// 							for i := range srcfile.TopLevel {
// 								symbolForDef(&srcfile.TopLevel[i], nil)
// 								if curFileOnly /*&& srcfile.TopLevel[i].Encloses(srcLens.Byte0OffsetForPos(srcLens.Pos))*/ {
// 									if tld := srcfile.TopLevel[i].Ast.Def.Orig; tld != nil {
// 										if let, _ := tld.Body.(*AstExprLet); let != nil {
// 											for j := range let.Defs {
// 												symbolForDef(&srcfile.TopLevel[i], &let.Defs[j])
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			})
// 		}
// 	})
// 	return
// }

// func (me *atmoSrcIntel) addLocFromToks(tlc *AstFileChunk, dst *z.SrcLocs, toks udevlex.Tokens) *z.SrcLoc {
// 	if tok := toks.First1(); tok != nil {
// 		pos := tok.OffPos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
// 		return dst.Add(tlc.SrcFile.SrcFilePath, pos)
// 	}
// 	return nil
// }

// func (me *atmoSrcIntel) addLocFromNode(tld *IrDef, dst *z.SrcLocs, node IIrNode) *z.SrcLoc {
// 	toks := tld.AstOrigToks(node)
// 	if def := node.IsDef(); def != nil {
// 		if ts := tld.AstOrigToks(&def.Ident); len(ts) > 0 {
// 			toks = ts
// 		}
// 	}
// 	return me.addLocFromToks(tld.AstFileChunk, dst, toks)
// }

// func (me *atmoSrcIntel) astAt(kit *atmosess.Kit, srcLens *z.SrcLens) (topLevelChunk *AstFileChunk, theNodeAndItsAncestors []IAstNode) {
// 	if topLevelChunk, theNodeAndItsAncestors = kit.AstNodeAt(srcLens.FilePath, srcLens.Byte0OffsetForPos(srcLens.Pos)); topLevelChunk != nil && topLevelChunk.Ast.Def.Orig == nil {
// 		theNodeAndItsAncestors = nil
// 	}
// 	return
// }

// func tokToPos(tlc *AstFileChunk, tok *udevlex.Token) (ret *z.SrcPos) {
// 	pos := tok.OffPos(tlc.PosOffsetLine(), tlc.PosOffsetByte())
// 	ret = &z.SrcPos{Ln: pos.Ln1, Col: pos.Col1}
// 	ret.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)
// 	return
// }

// func toksToRange(tlc *AstFileChunk, toks udevlex.Tokens) *z.SrcRange {
// 	var sr z.SrcRange
// 	tlcoffl, tlcoffb := tlc.PosOffsetLine(), tlc.PosOffsetByte()
// 	pos := toks.First1().OffPos(tlcoffl, tlcoffb)
// 	sr.Start.Col, sr.Start.Ln = pos.Col1, pos.Ln1
// 	sr.Start.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)

// 	pos = toks.Last1().OffPosEnd(tlcoffl, tlcoffb)
// 	sr.End.Col, sr.End.Ln = pos.Col1, pos.Ln1
// 	sr.End.SetRune1OffFromByte0Off(pos.Off0, tlc.SrcFile.LastLoad.Src)

// 	return &sr
// }

// func (me *atmoSrcIntel) withCtxLockedAndKitAndInMemFileMod(srcLens *z.SrcLens, do func(*atmosess.Kit)) {
// 	Ctx.Locked(func() {
// 		if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
// 			me.withInMemFileMod(srcLens, kit, func() {
// 				do(kit)
// 			})
// 		}
// 	})
// }

// func (me *atmoSrcIntel) withInMemFileMod(srcLens *z.SrcLens, kit *atmosess.Kit, do func()) {
// 	Ctx.KitEnsureLoaded(kit)
// 	if srcLens.Txt == "" {
// 		do()
// 	} else if liveMode {
// 		if srcfile := kit.SrcFiles.ByFilePath(srcLens.FilePath); srcfile != nil {
// 			srcfile.Options.TmpAltSrc = []byte(srcLens.Txt)
// 			Ctx.CatchUpOnFileMods(srcfile)
// 		}
// 		do()
// 	} else if panicked := Ctx.WithInMemFileMod(srcLens.FilePath, srcLens.Txt, do); panicked != nil {
// 		panic(panicked)
// 	}
// }
