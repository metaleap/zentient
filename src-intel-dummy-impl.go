package z

import (
	"strings"
)

func (_ *SrcIntelBase) ComplDetails(srcLens *SrcLens, itemText string, into *SrcIntelCompl) {
	into.Detail = "Details for " + itemText
	into.Documentation.IsTrusted, into.Documentation.Value = true, "**Docs** for: `"+itemText+"`"
}

func (_ *SrcIntelBase) ComplItems(srcLens *SrcLens) (all []SrcIntelCompl) {
	all = make([]SrcIntelCompl, CMPL_MIN_INVALID)
	for i := 0; i < len(all); i++ {
		cmplkind := Completion(i)
		all[i].Label = cmplkind.String()
		all[i].Kind = cmplkind
	}
	return
}

func (me *SrcIntelBase) DefSym(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}

func (me *SrcIntelBase) DefType(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}

func (me *SrcIntelBase) DefImpl(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}

func (_ *SrcIntelBase) Highlights(srcLens *SrcLens, curWord string) (all []SrcRange) {
	// bad implementation (will return buggy ranges with some exotic/unicode chars) but is meant to be overridden by a proper one anyway
	srcLens.EnsureSrcFull()
	src := strings.ToUpper(srcLens.Txt)
	if curWord == "" && srcLens.Range != nil {
		curWord = src[:srcLens.Range.End.Off-1][srcLens.Range.Start.Off-1:]
	}
	if curWord != "" {
		curWord = strings.ToUpper(curWord)
		pos, idx := 0, -1
		for true {
			if idx = strings.Index(src[pos:], curWord); idx < 0 {
				break
			}
			sr := SrcRange{Start: SrcPos{Off: idx + pos + 1}, End: SrcPos{}}
			sr.End.Off = sr.Start.Off + len(curWord)
			all = append(all, sr)
			if pos += idx + 1; pos >= len(src) {
				break
			}
		}
	}
	return
}

func (_ *SrcIntelBase) Hovers(srcLens *SrcLens) (all []InfoTip) {
	all = append(all,
		InfoTip{Value: Strf("Hovers not yet implemented for **%s** by `%s`", Lang.Title, Prog.name)},
		InfoTip{Value: Strf("File: %s", srcLens.FilePath), Language: "plaintext"},
		InfoTip{Value: Strf("Line/Char/Offset: %v", *srcLens.Pos)},
	)
	return
}

func (*SrcIntelBase) References(srcLens *SrcLens, includeDeclaration bool) (all SrcLenses) {
	all = append(all,
		&SrcLens{FilePath: srcLens.FilePath, Pos: &SrcPos{Col: 1, Ln: 3}},
		&SrcLens{FilePath: srcLens.FilePath, Pos: &SrcPos{Col: 2, Ln: 5}},
		&SrcLens{FilePath: srcLens.FilePath, Pos: &SrcPos{Col: 4, Ln: 8}},
		&SrcLens{FilePath: srcLens.FilePath, Pos: &SrcPos{Col: 7, Ln: 12}})
	if includeDeclaration {
		all = append(all, srcLens)
	}
	return
}

func (*SrcIntelBase) Signature(srcLens *SrcLens) *SrcIntelSigHelp {
	var sig SrcIntelSigHelp
	sig.Signatures = []SrcIntelSigInfo{{Label: "Signature", Documentation: SrcIntelDoc{IsTrusted: true, Value: "Helpful **doc** `comment`s.."}}}
	sig.Signatures[0].Parameters = []SrcIntelSigParam{{Label: "Parameter 1", Documentation: SrcIntelDoc{IsTrusted: true, Value: "Every argument gets a *helpful* `doc` comment."}}}
	return &sig
}

func (*SrcIntelBase) Symbols(srcLens *SrcLens, query string, curFileOnly bool) (all SrcLenses) {
	if curFileOnly {
		const symMinInvalid = int(SYM_MIN_INVALID)
		for i := 0; i < symMinInvalid; i++ {
			all = append(all,
				&SrcLens{Flag: i, Str: "Str-Name:" + Symbol(i).String(), FilePath: srcLens.FilePath,
					Txt: Strf("Txt-ContainerName: %d", i), Pos: &SrcPos{Col: 1, Ln: i + 1},
				},
			)
		}
	}
	return
}
