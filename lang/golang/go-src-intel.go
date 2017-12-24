package zgo

import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var srcIntel goSrcIntel

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

func (me *goSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (allsyms []*z.SrcLens) {
	onerr := func(label string, detail string) []*z.SrcLens {
		return []*z.SrcLens{&z.SrcLens{Flag: int(z.SYM_EVENT), Str: label, Txt: detail, FilePath: srcLens.FilePath, Pos: srcLens.Pos, Range: srcLens.Range}}
	}
	if !udevgo.Has_guru {
		return onerr("Not installed: guru", "for more information, see: Zentient Main Menu / Tooling / guru.")
	}
	srcLens.EnsureSrcFull()
	bytepos := srcLens.ByteOffsetForFirstLineBeginningWith("package ")
	gd, err := udevgo.QueryDesc_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(bytepos))
	if err != nil {
		return onerr("Error running guru:", err.Error())
	} else if gd.Package == nil {
		return onerr("Error running guru:", "not in a Go package")
	}

	// no more early-returns, now get busy
	fpathok := func(fp string) bool { return (!curFileOnly) || fp == srcLens.FilePath }
	curpkgdir, numsyms, fallbackfilepath := filepath.Dir(srcLens.FilePath), len(gd.Package.Members), func() string { return srcLens.FilePath }
	allsyms = make([]*z.SrcLens, 0, numsyms)
	for _, pm := range gd.Package.Members {
		ispmlisted := false
		if srcref := udev.SrcMsgFromLn(pm.Pos); srcref != nil && fpathok(srcref.Ref) {
			pmtype, lens := pm.Type, &z.SrcLens{Str: pm.Kind + " " + pm.Name}
			pmuntyped := strings.HasPrefix(pmtype, "untyped ")
			if ispmlisted = true; pmuntyped {
				pmtype = pmtype[8:]
			} else {
				next := func() int { return strings.Index(pmtype, ` "json:\"`) }
				for ij1 := next(); ij1 > 0; ij1 = next() {
					if ij2 := strings.Index(pmtype[ij1+9:], `\""`); ij2 >= 0 {
						pref, suff := pmtype[:ij1], pmtype[ij1+9+ij2+3:]
						pmtype = pref + suff
					} else {
						break
					}
				}
			}
			pmtype = udevgo.PkgImpPathsToNamesInLn(pmtype, curpkgdir)
			lens.SetFrom(srcref, fallbackfilepath)
			switch pm.Kind {
			case "const":
				if lens.Flag, lens.Txt = int(z.SYM_CONSTANT), pmtype+" = "+pm.Value; !pmuntyped {
					lens.Str, lens.Flag = "▶   "+pm.Name, int(z.SYM_NUMBER)
				}
			case "var":
				lens.Flag, lens.Txt = int(z.SYM_VARIABLE), pmtype
			case "func":
				fnargs, fnret := goSrcFuncDeclBreak(pmtype)
				lens.Flag, lens.Txt = int(z.SYM_FUNCTION), udevgo.PkgImpPathsToNamesInLn(fnret, curpkgdir)
				lens.Str += "  " + udevgo.PkgImpPathsToNamesInLn(strings.TrimPrefix(fnargs, "func"), curpkgdir)
			case "type":
				lens.Txt = pmtype
				e := z.SYM_CLASS
				switch pmtype {
				case "float32", "float64", "float", "complex", "int64", "uint64":
					e = z.SYM_NUMBER
				case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
					e = z.SYM_ENUMMEMBER
				case "rune", "string":
					e = z.SYM_STRING
				case "bool":
					e = z.SYM_BOOLEAN
				default:
					found := false
					for pref, enum := range map[string]z.Symbol{
						"*":          z.SYM_NULL,
						"map[":       z.SYM_OBJECT,
						"[]":         z.SYM_ARRAY,
						"func(":      z.SYM_EVENT,
						"interface{": z.SYM_INTERFACE,
						"struct{":    z.SYM_STRUCT,
					} {
						if found = strings.HasPrefix(pmtype, pref); found {
							e = enum
							break
						}
					}
					if !found {
						println(pmtype)
					}
				}
				lens.Flag = int(e)
			default:
				z.BadPanic("guru.DescribeMember.Kind", pm.Kind)
			}
			allsyms = append(allsyms, lens)
		}
		for _, method := range pm.Methods {
			if isok, srcref := ispmlisted || !strings.HasPrefix(pm.Type, "interface{"), udev.SrcMsgFromLn(method.Pos); isok && srcref != nil && fpathok(srcref.Ref) {
				p1, p2 := strings.Index(method.Name, " ("), strings.Index(method.Name, ") ")
				methodtype, methodtitle := method.Name[:p2][p1+2:], method.Name[p2+2:]
				methodname := strings.TrimSpace(methodtitle[:strings.Index(methodtitle, "(")])
				if strings.HasPrefix(pm.Type, "interface{") && !(strings.Contains(pm.Type, "{"+methodname+"(") || strings.Contains(pm.Type, "; "+methodname+"(")) {
					continue
				}
				if strings.HasPrefix(pm.Type, "struct{") {
					if srcln := srcLens.Ln(srcref.Pos1Ln); !strings.Contains(srcln, methodtype+") "+methodname+"(") {
						continue
					}
				}
				lens := &z.SrcLens{Flag: int(z.SYM_METHOD), Str: "▶   " + methodtitle}
				lens.SetFrom(srcref, fallbackfilepath)
				if !ispmlisted {
					lens.Str = methodtype + " " + lens.Str
				}

				lens.Str, lens.Txt = goSrcFuncDeclBreak(lens.Str)
				lens.Str, lens.Txt = udevgo.PkgImpPathsToNamesInLn(lens.Str, curpkgdir), udevgo.PkgImpPathsToNamesInLn(lens.Txt, curpkgdir)
				if i := strings.Index(lens.Str, "("); i > 0 {
					lens.Str = lens.Str[:i] + "  " + lens.Str[i:]
				}
				allsyms = append(allsyms, lens)
			}
		}
	}
	return
}

func goSrcFuncDeclBreak(fndecl string) (fnargs string, fnret string) {
	if i := strings.Index(fndecl, ") "); i <= 0 {
		fnargs, fnret = fndecl, ""
	} else {
		fnargs, fnret = fndecl[:i+1], fndecl[i+2:]
	}
	return
}
