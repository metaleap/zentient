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

func (*goSrcIntel) hoverShortenImpPaths(s string) string {
	if islash := ustr.Idx(s, "/"); islash > 0 {
		if idot := ustr.Idx(s[islash+1:], "."); idot > 0 && udevgo.ShortenImpPaths != nil {
			return udevgo.ShortenImpPaths.Replace(s)
		}
	}
	return s
}

func (me *goSrcIntel) Hovers(srcLens *z.SrcLens) (hovs []z.InfoTip) {
	var ggd *udevgo.Gogetdoc
	var decl *z.InfoTip
	offset := z.Strf("%d", srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos))

	if !tools.gogetdoc.Installed {
		hovs = append(hovs, z.InfoTip{Value: tools.gogetdoc.NotInstalledMessage()})
	} else {
		if ggd = udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.Txt, offset); ggd != nil {
			ispkglocal := ustr.Pref(ggd.Pos, filepath.Dir(srcLens.FilePath))
			if ggd.Err != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.Err})
			}
			if ggd.ErrMsgs != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.ErrMsgs})
			}
			if headline := ggd.ImpN; false && headline != "" && !ispkglocal {
				headline = me.hoverShortenImpPaths(headline)
				hovs = append(hovs, z.InfoTip{Value: "### " + headline})
			}
			if ggd.Decl = me.hoverDeclLineBreaks(ggd.Decl); ggd.Decl != "" {
				if ggd.ImpP != "" {
					ggd.Decl = strings.Replace(ggd.Decl, ggd.ImpP+".", "", -1)
				}
				if ustr.Pref(ggd.Decl, "field ") { // ensure syntax-highlighting:
					ggd.Decl = z.Strf("//ℤ/ struct field:\n{ %s }\n//ℤ/ field context (tags etc.) not shown", ggd.Decl[6:])
				}
				ggd.Decl = me.hoverShortenImpPaths(ggd.Decl)
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

func (me *goSrcIntel) Symbols(srcLens *z.SrcLens, query string, curFileOnly bool) (all []*z.SrcLens) {
	onerr := func(label string, detail string) []*z.SrcLens {
		return []*z.SrcLens{&z.SrcLens{Flag: int(z.SYM_EVENT), Str: label, Txt: detail, FilePath: srcLens.FilePath, Pos: srcLens.Pos, Range: srcLens.Range}}
	}
	if !udevgo.Has_guru {
		return onerr("Not installed: guru", "for more information, see: Zentient Main Menu / Tooling / guru.")
	}
	srcLens.EnsureSrcFull()
	bytepos := srcLens_IfSrcFull_BytePosOfPackageName(srcLens)
	gd, err := udevgo.QueryDesc_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(bytepos))
	if err != nil {
		return onerr("Error running guru:", err.Error())
	} else if gd.Package == nil {
		return onerr("Error running guru:", "not in a Go package")
	}
	fpathok := func(fp string) bool { return (!curFileOnly) || fp == srcLens.FilePath }
	curpkgdir, numsyms := filepath.Dir(srcLens.FilePath), len(gd.Package.Members)
	for _, pm := range gd.Package.Members {
		numsyms += len(pm.Methods)
	}
	all = make([]*z.SrcLens, 0, numsyms)
	for _, pm := range gd.Package.Members {
		if srcref := udev.SrcMsgFromLn(pm.Pos); srcref != nil && fpathok(srcref.Ref) {
			lens := &z.SrcLens{Flag: int(z.SYM_PACKAGE), Str: pm.Name, FilePath: srcref.Ref}
			pmtype := udevgo.PkgImpPathsToNamesIn(pm.Type, curpkgdir)
			switch pm.Kind {
			case "const":
				lens.Flag, lens.Txt = int(z.SYM_CONSTANT), "= "+pm.Value
			case "var":
				lens.Flag, lens.Txt = int(z.SYM_VARIABLE), pmtype
			case "func":
				lens.Flag, lens.Txt = int(z.SYM_FUNCTION), pmtype
			}
			all = append(all, lens)
		}
	}
	return
}
