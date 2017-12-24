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
	symsPatterns = map[string]z.Symbol{
		"*":          z.SYM_NULL,
		"map[":       z.SYM_OBJECT,
		"[]":         z.SYM_ARRAY,
		"func(":      z.SYM_EVENT,
		"interface{": z.SYM_INTERFACE,
		"struct{":    z.SYM_STRUCT,
	}
)

func init() {
	srcIntel.Impl, z.Lang.SrcIntel = &srcIntel, &srcIntel
}

type goSrcIntel struct {
	z.SrcIntelBase
}

func (*goSrcIntel) hoverDeclLineBreaks(decl string) string {
	if len(decl) > 50 && !strings.Contains(decl, "\n") {
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
		if ggd = udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, offset); ggd != nil {
			curpkgdir := filepath.Dir(srcLens.FilePath)
			ispkglocal := strings.HasPrefix(ggd.Pos, curpkgdir)
			if ggd.Err != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.Err})
			}
			if ggd.ErrMsgs != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.ErrMsgs})
			}
			if headline := ggd.ImpN; false && headline != "" && !ispkglocal {
				headline = udevgo.PkgImpPathsToNamesInLn(headline, curpkgdir)
				hovs = append(hovs, z.InfoTip{Value: "### " + headline})
			}
			if ggd.Decl = me.hoverDeclLineBreaks(ggd.Decl); ggd.Decl != "" {
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
	if tools.godef.Installed && decl == nil {
		if defdecl := udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.Txt, offset); defdecl != "" {
			decl = &z.InfoTip{Language: z.Lang.ID, Value: me.hoverDeclLineBreaks(defdecl)}
			hovs = append([]z.InfoTip{*decl}, hovs...)
		}
	}
	return
}

func (me *goSrcIntel) Symbols(sL *z.SrcLens, query string, curFileOnly bool) (allsyms z.SrcLenses) {
	onerr := func(label string, detail string) z.SrcLenses {
		return z.SrcLenses{&z.SrcLens{Flag: int(z.SYM_EVENT), Str: label, Txt: detail, FilePath: sL.FilePath, Pos: sL.Pos, Range: sL.Range}}
	}
	if !udevgo.Has_guru {
		return onerr("Not installed: guru", "for more information, see: Zentient Main Menu / Tooling / guru.")
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
				fnargs, fnret := goSrcFuncSigBreak(pmtype)
				sym.Flag, sym.Txt = int(z.SYM_FUNCTION), udevgo.PkgImpPathsToNamesInLn(fnret, curpkgdir)
				sym.Str += "  " + udevgo.PkgImpPathsToNamesInLn(strings.TrimPrefix(fnargs, "func"), curpkgdir)
			case "type":
				sym.Txt, sym.Flag = pmtype, int(z.SYM_CLASS)
				switch pmtype {
				case "float32", "float64", "float", "complex", "int64", "uint64":
					sym.Flag = int(z.SYM_NUMBER)
				case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
					sym.Flag = int(z.SYM_ENUMMEMBER)
				case "rune", "string", "[]byte":
					sym.Flag = int(z.SYM_STRING)
				case "bool":
					sym.Flag = int(z.SYM_BOOLEAN)
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
					lens := &z.SrcLens{Flag: int(z.SYM_METHOD), Str: "▶   " + methodtitle}
					lens.SetFilePathAndPosOrRangeFrom(srcref, nil)
					// if !ispmlisted { // if method's receiver type not in the symbols listing, prepend it's name to the pretend-indentation
					lens.Str = methodtype + " " + lens.Str
					// }
					lens.Str, lens.Txt = goSrcFuncSigBreak(lens.Str)
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

func goSrcFuncSigBreak(fnsig string) (fnargs string, fnret string) {
	if i := strings.Index(fnsig, ") "); i > 0 { // func sig has return args
		fnargs, fnret = fnsig[:i+1], fnsig[i+2:]
	} else { // void func sig (has no return args)
		fnargs, fnret = fnsig, " " // " " instead of "" circumvents a VScode quirk in its 'Workspace Symbols' UX
	}
	return
}
