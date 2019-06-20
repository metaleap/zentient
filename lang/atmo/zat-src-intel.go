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

func (me *atmoSrcIntel) DefSym(srcLens *z.SrcLens) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if curtlc, nodes := me.astAt(kit, srcLens); len(nodes) > 0 {
				// HAPPY SMART PATH: already know the def(s) or def-arg the current name points to
				if curtld, irnodes := kit.AstNodeIrFunFor(curtlc.Id(), nodes[0]); len(irnodes) > 0 {
					if ident, _ := irnodes[0].(*atmoil.AstIdentName); ident != nil {
						for _, node := range ident.Anns.Candidates {
							tld, _ := node.(*atmoil.AstDefTop)
							if tld == nil {
								tld = curtld
								if adr, ok := node.(atmosess.AstDefRef); ok {
									tld = adr.AstDefTop
								}
							}
							me.addLocFromNode(tld.OrigTopLevelChunk, &locs, tld.OrigTopLevelChunk.SrcFile.SrcFilePath, tld, node)
						}
					} else { // not an ident, then point to input node itself
						me.addLocFromNode(curtlc, &locs, curtlc.SrcFile.SrcFilePath, curtld, irnodes[0])
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
								me.addLocFromToks(curtlc, &locs, curtlc.SrcFile.SrcFilePath, nid.Tokens)
							}
						case *atmolang.AstDef:
							if n.Name.Val == ident.Val {
								me.addLocFromToks(curtlc, &locs, curtlc.SrcFile.SrcFilePath, n.Name.Tokens)
							} else {
								for da := range n.Args {
									if nid, _ := n.Args[da].NameOrConstVal.(*atmolang.AstIdent); nid != nil && nid.Val == ident.Val {
										me.addLocFromToks(curtlc, &locs, curtlc.SrcFile.SrcFilePath, nid.Tokens)
									}
								}
							}
						case *atmolang.AstExprLet:
							for d := range n.Defs {
								if n.Defs[d].Name.Val == ident.Val {
									me.addLocFromToks(curtlc, &locs, curtlc.SrcFile.SrcFilePath, n.Defs[d].Name.Tokens)
								}
							}
						}
					}
					// find all global goal-named defs --- loop isn't all that optimal but good-enough & succinct
					for _, k := range Ctx.Kits.All {
						if k.ImpPath == kit.ImpPath || ustr.In(k.ImpPath, kit.Imports...) {
							for _, def := range kit.Defs(ident.Val) {
								me.addLocFromNode(curtlc, &locs, def.OrigTopLevelChunk.SrcFile.SrcFilePath, def, def)
							}
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) Hovers(srcLens *z.SrcLens) (infoTips []z.InfoTip) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if tlc, nodes := me.astAt(kit, srcLens); len(nodes) > 0 {
				var nodetypenames string
				for _, n := range nodes {
					nodetypenames += z.Strf("— %T ", n)
				}
				infoTips = append(infoTips,
					z.InfoTip{Value: tlc.Ast.Def.Orig.Name.Val},
					z.InfoTip{Value: nodetypenames[4:]},
				)

				if _, ilnodes := kit.AstNodeIrFunFor(tlc.Id(), nodes[0]); len(ilnodes) > 0 {
					for _, n := range ilnodes {
						if nid, _ := n.(*atmoil.AstIdentName); nid != nil {
							infoTips = append(infoTips, z.InfoTip{Value: z.Strf("(resolves to %v candidate/s)", len(nid.Anns.Candidates))})
						}
						infoTips = append(infoTips,
							z.InfoTip{Value: z.Strf("%T:\n%s", n, n.Facts().Description()), Language: "plain"},
						)
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (locs z.SrcLocs) {
	if kit := Ctx.KitByDirPath(filepath.Dir(srcLens.FilePath), true); kit != nil {
		me.withInMemFileMod(srcLens, kit, func() {
			if _, nodes := me.astAt(kit, srcLens); len(nodes) > 0 {
				var refs map[*atmoil.AstDefTop][]atmoil.IAstExpr
				if ident, _ := nodes[0].(*atmolang.AstIdent); ident != nil {
					refs = Ctx.KitsCollectReferences(true, ident.Val)
				} else if atom, _ := nodes[0].(atmolang.IAstExprAtomic); atom != nil {
					refs = Ctx.KitsCollectReferences(true, atom.String())
				}
				for tld, nodes := range refs {
					for _, node := range nodes {
						if tok := tld.OrigToks(node).First(nil); tok != nil {
							locs.Add(tld.OrigTopLevelChunk.SrcFile.SrcFilePath, tok.Pos(tld.OrigTopLevelChunk.PosOffsetLine(), tld.OrigTopLevelChunk.PosOffsetByte()))
						}
					}
				}
			}
		})
	}
	return
}

func (me *atmoSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (allsyms z.SrcLenses) {
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
									allsyms = append(allsyms, &z.SrcLens{Str: def.Name.Val, Txt: "(description later)", SrcLoc: z.SrcLoc{
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

func (me *atmoSrcIntel) addLocFromToks(tlc *atmolang.SrcTopChunk, locs *z.SrcLocs, srcFilePath string, toks udevlex.Tokens) *z.SrcLoc {
	if tok := toks.First(nil); tok != nil {
		return locs.Add(srcFilePath, tok.Pos(tlc.PosOffsetLine(), tlc.PosOffsetByte()))
	}
	return nil
}

func (me *atmoSrcIntel) addLocFromNode(tlc *atmolang.SrcTopChunk, locs *z.SrcLocs, srcFilePath string, tld *atmoil.AstDefTop, node atmoil.IAstNode) *z.SrcLoc {
	toks := tld.OrigToks(node)
	if def := node.IsDef(); def != nil {
		if ts := tld.OrigToks(&def.Name); len(ts) > 0 {
			toks = ts
		}
	}
	return me.addLocFromToks(tlc, locs, srcFilePath, toks)
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
