package zgo

import (
	"path/filepath"
	"strings"

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
		if ggd = udevgo.Query_Gogetdoc(srcLens.FilePath, srcLens.SrcFull, offset); ggd != nil {
			ispkglocal := ustr.Pref(ggd.Pos, filepath.Dir(srcLens.FilePath))
			if ggd.Err != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.Err})
			}
			if ggd.ErrMsgs != "" {
				hovs = append(hovs, z.InfoTip{Language: "plaintext", Value: ggd.ErrMsgs})
			}
			if headline := ggd.ImpN; headline != "" && !ispkglocal {
				headline = me.hoverShortenImpPaths(headline)
				hovs = append(hovs, z.InfoTip{Value: "### " + headline})
			}
			if ggd.Decl != "" {
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
		if defdecl := udevgo.QueryDefDecl_GoDef(srcLens.FilePath, srcLens.SrcFull, offset); defdecl != "" {
			decl = &z.InfoTip{Language: z.Lang.ID, Value: defdecl}
			hovs = append([]z.InfoTip{*decl}, hovs...)
		}
	}
	return
}