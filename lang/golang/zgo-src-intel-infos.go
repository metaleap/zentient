package zgo

import (
	"path/filepath"
	"strings"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
)

var (
	srcIntel     goSrcIntel
	cmplPatterns = map[string]z.Completion{
		"*":          z.CMPL_UNIT,
		"map[":       z.CMPL_TYPEPARAMETER,
		"[]":         z.CMPL_TYPEPARAMETER,
		"func(":      z.CMPL_EVENT,
		"interface{": z.CMPL_INTERFACE,
		"struct{":    z.CMPL_CLASS,
	}
)

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

type goSrcIntel struct {
	z.SrcIntelBase
}

func (*goSrcIntel) ComplItemsShouldSort(*z.SrcLens) bool { return true }

func (*goSrcIntel) ComplItems(srcLens *z.SrcLens) (all z.SrcIntelCompls) {
	if !tools.gocode.Installed {
		return
	}
	rawresp, err := udevgo.QueryCmplSugg_Gocode(srcLens.FilePath, srcLens.Txt, z.Strf("c%d", srcLens.Pos.Off-1))
	if err != nil {
		panic(err)
	}
	if len(rawresp) > 0 {
		all = make(z.SrcIntelCompls, 0, len(rawresp))
		for _, raw := range rawresp {
			if c, n, t := raw["class"], raw["name"], raw["type"]; n != "" && !(c == "import" && strings.Contains(n, "/internal/")) {
				cmpl := &z.SrcIntelCompl{Detail: t, Label: n, Kind: z.CMPL_COLOR}
				switch c {
				case "PANIC":
					continue
				case "package", "import":
					cmpl.SortPrio, cmpl.Kind = 2, z.CMPL_FOLDER
				case "var":
					cmpl.SortPrio, cmpl.Kind = 3, z.CMPL_FIELD
				case "func":
					cmpl.SortPrio, cmpl.Kind = 4, z.CMPL_FUNCTION
				case "const":
					cmpl.SortPrio, cmpl.Kind = 20, z.CMPL_CONSTANT
				case "type":
					cmpl.SortPrio = 19
					switch t {
					case "interface":
						cmpl.SortPrio, cmpl.Kind = 10, z.CMPL_INTERFACE
					case "struct":
						cmpl.SortPrio, cmpl.Kind = 11, z.CMPL_CLASS
					case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
						cmpl.Kind = z.CMPL_ENUMMEMBER
					case "float32", "float64", "complex64", "complex128", "int64", "uint64":
						cmpl.Kind = z.CMPL_OPERATOR
					case "string", "rune", "[]byte":
						cmpl.Kind = z.CMPL_UNIT
					case "built-in":
						switch n {
						case "byte", "float32", "float64", "complex64", "complex128", "int64", "uint64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
							cmpl.Kind = z.CMPL_OPERATOR
						case "string", "rune", "bool", "uintptr":
							cmpl.Kind = z.CMPL_UNIT
						default:
							cmpl.Kind = z.CMPL_FILE
						}
					default:
						cmpl.SortPrio, cmpl.Kind = 12, z.CMPL_TYPEPARAMETER
						for pref, ck := range cmplPatterns {
							if strings.HasPrefix(t, pref) {
								cmpl.Kind = ck
								break
							}
						}
					}
				default:
					cmpl.SortPrio, cmpl.Detail = 1, "CMPLCLS:["+c+"]\n"+cmpl.Detail
				}
				all = append(all, cmpl)
			}
		}
	}
	return
}

func (me *goSrcIntel) ComplDetails(srcLens *z.SrcLens, itemText string) (itemDoc *z.SrcIntelCompl) {
	if !(tools.gogetdoc.Installed || tools.godef.Installed) {
		return
	}
	pos := srcLens.Byte0OffsetForPos(srcLens.Pos)
	rs := srcLens.Byte0OffsetForPos(&srcLens.Range.Start)
	re := srcLens.Byte0OffsetForPos(&srcLens.Range.End)
	srcLens.Txt = srcLens.Txt[:rs] + itemText + srcLens.Txt[re:]
	itemDoc = &z.SrcIntelCompl{
		Documentation: &z.SrcIntelDoc{IsTrusted: true},
	}
	decl, spos := "", ustr.Int(pos)
	if tools.gogetdoc.Installed {
		ggd := udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, spos, true, true)
		if decl = ggd.Decl; decl != "" {
			decl = me.goFuncDeclLineBreaks(udevgo.PkgImpPathsToNamesInLn(decl, filepath.Dir(srcLens.FilePath)), 23)
		}
		if ggd.Doc != "" {
			itemDoc.Documentation.Value = strings.TrimSpace(ggd.Doc)
		} else if ggd.Err != "" {
			itemDoc.Documentation.Value = ggd.Err
		} else if ggd.ErrMsgs != "" {
			// typically uninteresting here, ie. parse errors from transient editing state
			// itemDoc.Documentation.Value = ggd.ErrMsgs
		}
	}
	if decl == "" && tools.godef.Installed {
		if decl = udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.Txt, spos); decl != "" {
			decl = me.goFuncDeclLineBreaks(decl, 23)
		}
	}
	itemDoc.Detail = me.goDeclSnip(decl)
	return
}
func (*goSrcIntel) goFuncDeclLineBreaks(decl string, maxlen int) string {
	if len(decl) > maxlen && !strings.Contains(decl, "\n") {
		dl, dr := decl[:6], decl[6:]
		next := func() int { return strings.IndexRune(dr, '(') }
		for i := next(); i >= 0; i = next() {
			isemptyparens := i <= (len(dr)-2) && dr[i:i+2] == "()"
			isfuncsig1 := i >= 4 && dr[i-4:i] == "func"
			isfuncsig2 := i >= 5 && dr[i-5:i] == "func "
			if ignore := isemptyparens || isfuncsig1 || isfuncsig2; ignore {
				dl += dr[:i+2]
				dr = dr[i+2:]
			} else {
				dl += dr[:i] + "\n  ("
				dr = dr[i+1:]
			}
		}
		decl = dl + dr
	}
	return decl
}

func (me *goSrcIntel) CanIntelForCmplOrHover(lex *z.SrcIntelLex) bool {
	return lex == nil || lex.Ident != "" || lex.Other != ""
}

func (me *goSrcIntel) Hovers(srcLens *z.SrcLens) (hovs []z.SrcInfoTip) {
	var ggd *udevgo.Gogetdoc
	var decl *z.SrcInfoTip
	offset := z.Strf("%d", srcLens.Byte0OffsetForPos(srcLens.Pos))

	if !tools.gogetdoc.Installed {
		hovs = append(hovs, z.SrcInfoTip{Value: tools.gogetdoc.NotInstalledMessage()})
	} else {
		if ggd = udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, offset, false, true); ggd != nil {
			curpkgdir := filepath.Dir(srcLens.FilePath)
			ispkglocal := strings.HasPrefix(ggd.Pos, curpkgdir)
			if ggd.Err != "" {
				hovs = append(hovs, z.SrcInfoTip{Language: "plaintext", Value: ggd.Err})
			}
			// if ggd.ErrMsgs != "" {
			// 	// typically uninteresting here, ie. parse errors from transient editing state
			// 	hovs = append(hovs, z.SrcInfoTip{Language: "plaintext", Value: ggd.ErrMsgs})
			// }
			if headline := ggd.ImpN; false && headline != "" && !ispkglocal {
				headline = udevgo.PkgImpPathsToNamesInLn(headline, curpkgdir)
				hovs = append(hovs, z.SrcInfoTip{Value: "### " + headline})
			}
			if ggd.Decl = me.goFuncDeclLineBreaks(ggd.Decl, 42); ggd.Decl != "" {
				if ggd.ImpP != "" {
					ggd.Decl = strings.Replace(ggd.Decl, ggd.ImpP+".", "", -1)
				}
				ggd.Decl = udevgo.PkgImpPathsToNamesInLn(ggd.Decl, curpkgdir)
				if strings.HasPrefix(ggd.Decl, "field ") { // ensure syntax-highlighting:
					ggd.Decl = z.Strf("//ℤ/ struct field:\n{ %s }\n//ℤ/ field context (tags etc.) not shown", ggd.Decl[6:])
				}
				decl = &z.SrcInfoTip{Language: z.Lang.ID, Value: ggd.Decl}
			}
			if impdoc := ggd.ImpP; ggd.Doc != "" || impdoc != "" {
				if ispkglocal {
					impdoc = ""
				} else if impdoc != "" && ggd.DocUrl != "" {
					if impdoc != ggd.Pkg {
						impdoc = z.Strf("`import %s %q`", ggd.Pkg, impdoc)
					} else if impdoc != "builtin" {
						impdoc = z.Strf("`import %q`", impdoc)
					}
					docuri := "zentient://" + z.Lang.ID + "/godoc/pkg/" + ggd.DocUrl
					impdoc = z.Strf("[%s](%s)", impdoc, pages.linkifyUri(docuri))
				}
				hovs = append(hovs, z.SrcInfoTip{Value: ustr.Combine(impdoc, "\n\n", ggd.Doc)})
			}
		}
	}
	if decl == nil && tools.godef.Installed {
		if defdecl := udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.Txt, offset); defdecl != "" {
			decl = &z.SrcInfoTip{Language: z.Lang.ID, Value: me.goFuncDeclLineBreaks(defdecl, 42)}
		}
	}
	if decl != nil {
		if strings.Count(decl.Value, "\n") > 3 {
			hovs = append(hovs, *decl)
		} else {
			hovs = append([]z.SrcInfoTip{*decl}, hovs...)
		}
	}
	return
}

func (me *goSrcIntel) Signature(srcLens *z.SrcLens) (sig *z.SrcIntelSigHelp) {
	sig = &z.SrcIntelSigHelp{Signatures: []z.SrcIntelSigInfo{{}}}
	sig0 := &sig.Signatures[0]
	if !(tools.guru.Installed && (tools.gogetdoc.Installed || tools.godef.Installed)) {
		sig0.Label, sig0.Documentation.Value = z.ToolsMsgGone("guru or one of gogetdoc/godef"), z.ToolsMsgMore("(tool name)")
		return
	}
	pos := srcLens.Byte0OffsetForPos(srcLens.Pos)
	gw, err := udevgo.QueryWhat_Guru(srcLens.FilePath, srcLens.Txt, ustr.Int(pos))
	if err != nil {
		sig0.Label, sig0.Documentation.Value = "Error running guru", err.Error()
		return
	}
	posmax := pos
	pos = -1
	for _, ge := range gw.Enclosing {
		if strings.HasPrefix(ge.Description, "function call") {
			pos = ge.Start
			break
		}
	}
	if pos < 0 {
		sig = nil
	} else {
		poss := []int{}
		for mpos := pos; mpos < posmax; mpos++ {
			if c := srcLens.Txt[mpos+1]; c == '(' {
				poss = append(poss, mpos)
			} else if l := len(poss); c == ')' && l > 1 {
				poss = poss[:l-1]
			}
		}
		if len(poss) == 0 {
			sig = nil
		} else {
			decl, spos := "", ustr.Int(poss[len(poss)-1])
			if tools.gogetdoc.Installed {
				ggd := udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, spos, true, true)
				if decl = ggd.Decl; decl != "" {
					decl = udevgo.PkgImpPathsToNamesInLn(decl, filepath.Dir(srcLens.FilePath))
				}
				if ggd.Doc != "" {
					sig0.Documentation.Value = strings.TrimSpace(ggd.Doc)
				} else if ggd.Err != "" {
					sig0.Documentation.Value = ggd.Err
				} else if ggd.ErrMsgs != "" {
					// typically uninteresting here, ie. parse errors from transient editing state
					// sig0.Documentation.Value = ggd.ErrMsgs
				}
			} else {
				sig0.Documentation.Value = z.ToolsMsgGone("gogetdoc")
			}
			if decl == "" && tools.godef.Installed {
				decl = udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.Txt, spos)
			}
			if decl == "" {
				sig = nil
			} else {
				sig0.Label = me.goDeclSnip(me.goFuncDeclLineBreaks(decl, 42))
			}
		}
	}
	return
}

func (*goSrcIntel) goDeclSnip(decl string) string {
	if strings.HasPrefix(decl, "var ") {
		decl = decl[4:]
	} else if strings.HasPrefix(decl, "field ") {
		decl = decl[6:]
	}
	if strings.HasPrefix(decl, "func ") {
		decl = decl[5:]
	} else if i, j := strings.Index(decl, " func("), strings.IndexRune(decl, ' '); i > 0 && i == j {
		decl = decl[:i] + decl[i+5:]
	} else {
		for {
			if i = strings.Index(decl, " `"); i <= 0 {
				if i = strings.Index(decl, "\t`"); i <= 0 {
					break
				}
			}
			if j = strings.Index(decl[i+2:], "\"`"); j <= 0 {
				break
			} else {
				decl = decl[:i] + decl[j+2+2+i:]
			}
		}
	}
	return decl
}
