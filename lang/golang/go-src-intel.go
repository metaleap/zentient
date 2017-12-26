package zgo

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var (
	srcIntel     goSrcIntel
	cmplPatterns = map[string]z.Completion{
		"*":          z.CMPL_UNIT,
		"map[":       z.CMPL_VARIABLE,
		"[]":         z.CMPL_VARIABLE,
		"func(":      z.CMPL_EVENT,
		"interface{": z.CMPL_INTERFACE,
		"struct{":    z.CMPL_CLASS,
	}
	symsPatterns = map[string]z.Symbol{
		"*":          z.SYM_NULL,
		"map[":       z.SYM_OBJECT,
		"[]":         z.SYM_ARRAY,
		"func(":      z.SYM_EVENT,
		"interface{": z.SYM_INTERFACE,
		"struct{":    z.SYM_CLASS,
	}
	cmplCharsFunc = []string{"("}
)

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

type goSrcIntel struct {
	z.SrcIntelBase
}

func (*goSrcIntel) ComplItems(srcLens *z.SrcLens) (all []*z.SrcIntelCompl) {
	if !tools.gocode.Installed {
		return
	}
	rawresp, err := udevgo.QueryCmplSugg_Gocode(srcLens.FilePath, srcLens.Txt, z.Strf("c%d", srcLens.Pos.Off-1))
	if err != nil {
		panic(err)
	}
	if len(rawresp) > 0 {
		all = make([]*z.SrcIntelCompl, 0, len(rawresp))
		for _, raw := range rawresp {
			if c, n, t := raw["class"], raw["name"], raw["type"]; n != "" {
				cmpl := &z.SrcIntelCompl{Detail: t, Label: n, Kind: z.CMPL_COLOR, FilterText: strings.ToLower(n)}
				switch c {
				case "func":
					cmpl.Kind = z.CMPL_FUNCTION
					cmpl.SortText = "9" + cmpl.Label
					cmpl.CommitChars = cmplCharsFunc
				case "package":
					cmpl.Kind = z.CMPL_FOLDER
					cmpl.SortText = "1" + cmpl.Label
				case "var":
					cmpl.Kind = z.CMPL_FIELD
					cmpl.SortText = "4" + cmpl.Label
				case "const":
					cmpl.Kind = z.CMPL_CONSTANT
					cmpl.SortText = "3" + cmpl.Label
				case "type":
					cmpl.SortText = "2" + cmpl.Label
					switch t {
					case "built-in":
						switch n {
						case "byte", "float32", "float64", "float", "complex64", "complex128", "int64", "uint64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
							cmpl.Kind = z.CMPL_OPERATOR
						case "string", "rune", "bool", "uintptr":
							cmpl.Kind = z.CMPL_UNIT
						default:
							cmpl.Kind = z.CMPL_FILE
						}
					case "float32", "float64", "float", "complex64", "complex128", "int64", "uint64":
						cmpl.Kind = z.CMPL_OPERATOR
					case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
						cmpl.Kind = z.CMPL_ENUMMEMBER
					case "string", "rune", "[]byte":
						cmpl.Kind = z.CMPL_UNIT
					case "struct":
						cmpl.Kind = z.CMPL_CLASS
					case "interface":
						cmpl.Kind = z.CMPL_INTERFACE
					default:
						cmpl.Kind = z.CMPL_TYPEPARAMETER
						for pref, ck := range cmplPatterns {
							if strings.HasPrefix(t, pref) {
								cmpl.Kind = ck
								break
							}
						}
					}
				default:
					cmpl.SortText = "0" + cmpl.Label
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
	pos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	rs := srcLens.ByteOffsetForPosWithRuneOffset(&srcLens.Range.Start)
	re := srcLens.ByteOffsetForPosWithRuneOffset(&srcLens.Range.End)
	srcLens.Txt = srcLens.Txt[:rs] + itemText + srcLens.Txt[re:]
	itemDoc = &z.SrcIntelCompl{
		Documentation: &z.SrcIntelDoc{IsTrusted: true},
	}
	decl, spos := "", ustr.FromInt(pos)
	if tools.gogetdoc.Installed {
		ggd := udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, spos, true)
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
	if itemDoc.Documentation.Value == "" {
		itemDoc.Documentation.Value = " " // z.Strf("(No docs for `%s` — at least if inserted here)", ggd.Name)
	}
	return
}

func (*goSrcIntel) Highlights(srcLens *z.SrcLens, curWord string) (all z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	byteoff := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	gw, err := udevgo.QueryWhat_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(byteoff))
	if err != nil {
		panic(err)
	}
	all = make(z.SrcLenses, 0, len(gw.SameIDs))
	for _, sameid := range gw.SameIDs {
		if srcref := udev.SrcMsgFromLn(sameid); srcref != nil {
			sl := &z.SrcLens{}
			if sl.SetFilePathAndPosOrRangeFrom(srcref, nil); sl.FilePath == srcLens.FilePath && (sl.Range != nil || sl.Pos != nil) {
				sl.FilePath = ""
				all = append(all, sl)
			}
		}
	}
	if len(all) == 0 && len(gw.Enclosing) > 0 {
		srcraw := []byte(srcLens.Txt)
		bpos2rpos := func(bytepos int) int {
			length := 0
			for _ = range string(srcraw[:bytepos]) {
				length++
			}
			return length // here's what *won't* work for multi-byte/unicode/etc: just bytepos, or even len(string(srcraw[:bytepos]))
		}
		var check func(num int, checks ...string) bool
		check = func(num int, checks ...string) bool {
			if num <= 0 {
				for _, chk := range checks {
					if check(1, chk, chk) {
						return true
					}
				}
				return false
			}
			if ustr.AnyOf(gw.Enclosing[0].Description, checks[:num]...) {
				for _, syntaxnode := range gw.Enclosing {
					if ustr.AnyOf(syntaxnode.Description, checks[num:]...) {
						all = append(all, &z.SrcLens{Range: &z.SrcRange{
							Start: z.SrcPos{Off: 1 + bpos2rpos(syntaxnode.Start)}, End: z.SrcPos{Off: 1 + bpos2rpos(syntaxnode.End)}}})
						return true
					}
				}
			}
			return false
		}
		if check(1, "break statement",
			"range loop", "for loop", "select statement", "switch statement") {
			return
		}
		if check(1, "case clause",
			"select statement", "switch statement") {
			return
		}
		if check(1, "continue statement",
			"range loop", "for loop") {
			return
		}
		if check(2, "defer statement", "return statement",
			"function literal", "function declaration") {
			return
		}
		check(-1, "if statement", "select statement", "switch statement", "go statement", "range loop", "for loop", "struct type", "interface type", "map type", "function type", "slice type", "type specification", "type declaration", "function declaration", "field/method/parameter", "field/method/parameter list", "function call", "function call (or conversion)", "basic literal", "composite literal", "variable declaration", "constant declaration")
	}
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

func (me *goSrcIntel) Hovers(srcLens *z.SrcLens) (hovs []z.InfoTip) {
	var ggd *udevgo.Gogetdoc
	var decl *z.InfoTip
	offset := z.Strf("%d", srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos))

	if !tools.gogetdoc.Installed {
		hovs = append(hovs, z.InfoTip{Value: tools.gogetdoc.NotInstalledMessage()})
	} else {
		if ggd = udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, offset, false); ggd != nil {
			curpkgdir := filepath.Dir(srcLens.FilePath)
			ispkglocal := strings.HasPrefix(ggd.Pos, curpkgdir)
			if ggd.Err != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.Err})
			}
			// if ggd.ErrMsgs != "" {
			// 	// typically uninteresting here, ie. parse errors from transient editing state
			// 	hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.ErrMsgs})
			// }
			if headline := ggd.ImpN; false && headline != "" && !ispkglocal {
				headline = udevgo.PkgImpPathsToNamesInLn(headline, curpkgdir)
				hovs = append(hovs, z.InfoTip{Value: "### " + headline})
			}
			if ggd.Decl = me.goFuncDeclLineBreaks(ggd.Decl, 42); ggd.Decl != "" {
				if ggd.ImpP != "" {
					ggd.Decl = strings.Replace(ggd.Decl, ggd.ImpP+".", "", -1)
				}
				ggd.Decl = udevgo.PkgImpPathsToNamesInLn(ggd.Decl, curpkgdir)
				if strings.HasPrefix(ggd.Decl, "field ") { // ensure syntax-highlighting:
					ggd.Decl = z.Strf("//ℤ/ struct field:\n{ %s }\n//ℤ/ field context (tags etc.) not shown", ggd.Decl[6:])
				}
				decl = &z.InfoTip{Language: z.Lang.ID, Value: ggd.Decl}
				hovs = append(hovs, *decl)
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
					impdoc = "[" + impdoc + "](http://godoc.org/" + ggd.DocUrl + ")"
				}
				hovs = append(hovs, z.InfoTip{Value: ustr.Both(impdoc, "\n\n", ggd.Doc)})
			}
		}
	}
	if decl == nil && tools.godef.Installed {
		if defdecl := udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.Txt, offset); defdecl != "" {
			decl = &z.InfoTip{Language: z.Lang.ID, Value: me.goFuncDeclLineBreaks(defdecl, 42)}
			hovs = append([]z.InfoTip{*decl}, hovs...)
		}
	}
	return
}

func (me *goSrcIntel) Signature(srcLens *z.SrcLens) (sig *z.SrcIntelSigHelp) {
	sig = &z.SrcIntelSigHelp{Signatures: []z.SrcIntelSigInfo{z.SrcIntelSigInfo{}}}
	sig0 := &sig.Signatures[0]
	if !(tools.guru.Installed && (tools.gogetdoc.Installed || tools.godef.Installed)) {
		sig0.Label, sig0.Documentation.Value = z.ToolsMsgGone("guru or one of gogetdoc/godef"), z.ToolsMsgMore("(tool name)")
		return
	}
	pos, posmax := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos), -1
	gw, err := udevgo.QueryWhat_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(pos))
	if err != nil {
		sig0.Label, sig0.Documentation.Value = "Error running guru", err.Error()
		return
	}
	pos, posmax = -1, pos
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
			decl, spos := "", ustr.FromInt(poss[len(poss)-1])
			if tools.gogetdoc.Installed {
				ggd := udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, spos, true)
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

func (me *goSrcIntel) Symbols(sL *z.SrcLens, query string, curFileOnly bool) (allsyms z.SrcLenses) {
	onerr := func(label string, detail string) z.SrcLenses {
		return z.SrcLenses{&z.SrcLens{Flag: int(z.SYM_EVENT), Str: label, Txt: detail, FilePath: sL.FilePath, Pos: sL.Pos, Range: sL.Range}}
	}
	if !tools.guru.Installed {
		return onerr(z.ToolsMsgGone("guru"), z.ToolsMsgMore("guru"))
	}
	sL.EnsureSrcFull()
	srclns := strings.Split(sL.Txt, "\n")
	bytepos := 8 + sL.ByteOffsetForFirstLineBeginningWith("package ")
	gd, err := udevgo.QueryDesc_Guru(sL.FilePath, sL.Txt, ustr.FromInt(bytepos))
	if err != nil {
		return onerr("Error running guru:", err.Error())
	} else if gd.Package == nil {
		return onerr("Error running guru:", "not in a Go package")
	}

	// no more early-returns, now get busy
	if !curFileOnly {
		sort.Sort(gd) // sort doesn't seem to help improve vsc's ctrl+t ux for now, but maybe in some future vsc release..
	}
	anyfilegoes, curpkgdir, numpms := !curFileOnly, filepath.Dir(sL.FilePath), len(gd.Package.Members)
	query, allsyms = strings.ToLower(query), make(z.SrcLenses, 0, numpms) // numpms will never be a 'good' cap in any of these, but hey any number beats the default cap of 0..
	for _, pm := range gd.Package.Members {
		ispmlisted := false
		if srcref := udev.SrcMsgFromLn(pm.Pos); srcref != nil && (anyfilegoes || srcref.Ref == sL.FilePath) && (query == "" || gd.Matches(pm, query)) {
			ispmlisted = true
			pmtype, sym := pm.Type, &z.SrcLens{Str: pm.Name}
			pmuntyped := strings.HasPrefix(pmtype, "untyped ")
			{
				if pmuntyped {
					pmtype = pmtype[8:]
				} else if strings.HasPrefix(pmtype, "struct{") {
					next := func() int { return strings.Index(pmtype, ` "json:\"`) }
					for ij1 := next(); ij1 > 0; ij1 = next() {
						if ij2 := strings.Index(pmtype[ij1+9:], `\""`); ij2 >= 0 {
							pref, suff := pmtype[:ij1], pmtype[ij1+9+3+ij2:]
							pmtype = pref + suff
						} else {
							break
						}
					}
				}
				pmtype = udevgo.PkgImpPathsToNamesInLn(pmtype, curpkgdir)
			}
			sym.SetFilePathAndPosOrRangeFrom(srcref, nil)
			switch pm.Kind {
			case "const":
				if sym.Flag, sym.Txt = int(z.SYM_CONSTANT), pmtype+" = "+pm.Value; !pmuntyped {
					sym.Str, sym.Flag = "▶   "+pm.Name, int(z.SYM_NUMBER)
				}
			case "var":
				sym.Flag, sym.Txt = int(z.SYM_VARIABLE), pmtype
			case "func":
				fnargs, fnret := me.symFuncSigBreak(pmtype)
				sym.Flag, sym.Txt = int(z.SYM_FUNCTION), udevgo.PkgImpPathsToNamesInLn(fnret, curpkgdir)
				sym.Str += "  " + udevgo.PkgImpPathsToNamesInLn(strings.TrimPrefix(fnargs, "func"), curpkgdir)
			case "type":
				sym.Txt, sym.Flag = pmtype, int(z.SYM_TYPEPARAMETER)
				switch pmtype {
				case "float32", "float64", "float", "complex64", "complex128", "int64", "uint64":
					sym.Flag = int(z.SYM_NUMBER)
				case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
					sym.Flag = int(z.SYM_ENUMMEMBER)
				case "rune", "string", "[]byte":
					sym.Flag = int(z.SYM_STRING)
				case "bool":
					sym.Flag = int(z.SYM_BOOLEAN)
				case "uintptr":
					sym.Flag = int(z.SYM_NULL)
				default:
					for pref, enum := range symsPatterns {
						if strings.HasPrefix(pmtype, pref) {
							sym.Flag = int(enum)
							break
						}
					}
				}
			default:
				z.BadPanic("guru.DescribeMember.Kind", pm.Kind)
			}
			allsyms = append(allsyms, sym)
		}
		for _, method := range pm.Methods {
			if isok := ispmlisted || !strings.HasPrefix(pm.Type, "interface{"); isok && (query == "" || strings.Contains(strings.ToLower(method.Name), query)) {
				if srcref := udev.SrcMsgFromLn(method.Pos); srcref != nil && (anyfilegoes || srcref.Ref == sL.FilePath) {
					p1, p2 := strings.Index(method.Name, " ("), strings.Index(method.Name, ") ")
					methodtype, methodtitle := method.Name[:p2][p1+2:], method.Name[p2+2:]
					methodname := strings.TrimSpace(methodtitle[:strings.Index(methodtitle, "(")])
					if strings.HasPrefix(pm.Type, "interface{") && !(strings.Contains(pm.Type, "{"+methodname+"(") || strings.Contains(pm.Type, "; "+methodname+"(")) {
						// guru reports an embedded interface's methods redundantly for each embedder, all pointing to the embeddee's original loc. we skip these
						continue
					}
					if curFileOnly && srcref.Ref == sL.FilePath && strings.HasPrefix(pm.Type, "struct{") && srcref.Pos1Ln > 0 && srcref.Pos1Ln <= len(srclns) {
						if srcln := srclns[srcref.Pos1Ln-1]; !strings.Contains(srcln, methodtype+") "+methodname+"(") {
							// guru reports an embedded struct's methods redundantly for each embedder, all pointing to the embeddee's original loc. we skip these --- at least in "file symbols" mode
							continue
						}
					}
					lens := &z.SrcLens{Flag: int(z.SYM_METHOD), Str: "▶  " + methodtitle}
					lens.SetFilePathAndPosOrRangeFrom(srcref, nil)
					// if !ispmlisted { // if method's receiver type not in the symbols listing, prepend it's name to the pretend-indentation
					lens.Str = methodtype + "  " + lens.Str
					// }
					lens.Str, lens.Txt = me.symFuncSigBreak(lens.Str)
					lens.Str, lens.Txt = udevgo.PkgImpPathsToNamesInLn(lens.Str, curpkgdir), udevgo.PkgImpPathsToNamesInLn(lens.Txt, curpkgdir)
					if i := strings.Index(lens.Str, "("); i > 0 { // insert some spacing between name and args
						lens.Str = lens.Str[:i] + "  " + lens.Str[i:]
					}
					allsyms = append(allsyms, lens)
				}
			}
		}
	}
	return
}

func (*goSrcIntel) symFuncSigBreak(fnsig string) (fnargs string, fnret string) {
	fnargs, fnret = fnsig, " "
	co, cc, pos := 0, 0, 0
	for i, r := range fnsig {
		if r == '(' {
			co++
		} else if r == ')' {
			cc++
		}
		if cc > 0 && co > 0 && cc == co {
			pos = i
			break
		}
	}
	if pos > 0 && pos < len(fnsig)-1 {
		fnargs, fnret = fnsig[:pos+1], fnsig[pos+2:]
	}
	return
}

func (*goSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (refs z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	bytepos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	if gr := udevgo.QueryRefs_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(bytepos)); len(gr) > 0 {
		refs = make(z.SrcLenses, 0, len(gr))
		for _, gref := range gr {
			if srcref := udev.SrcMsgFromLn(gref.Pos); srcref != nil {
				refloc := &z.SrcLens{}
				refloc.SetFilePathAndPosOrRangeFrom(srcref, nil)
				refs = append(refs, refloc)
			}
		}
	}
	return
}
