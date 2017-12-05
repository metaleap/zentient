package z

import (
	"strings"

	"github.com/metaleap/go-util/dev"
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

func (_ *SrcIntelBase) Highlights(srcLens *SrcLens, curWord string) (all []SrcRange) {
	// bad implementation (will return buggy ranges with some exotic/unicode chars) but is meant to be overridden by a proper one anyway
	srcLens.ensureSrcFull()
	src := strings.ToUpper(srcLens.SrcFull)
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

func (_ *SrcIntelBase) Hovers(srcLens *SrcLens) (all []SrcIntelHover) {
	all = append(all,
		SrcIntelHover{Value: Strf("Hovers not yet implemented for **%s** by `%s`", Lang.Title, Prog.name)},
		SrcIntelHover{Value: Strf("File: %s", srcLens.FilePath), Language: "plaintext"},
		SrcIntelHover{Value: Strf("Line/Char/Offset: %v", *srcLens.Pos)},
	)
	return
}

func (*SrcIntelBase) Signature(srcLens *SrcLens) *SrcIntelSigHelp {
	var sig SrcIntelSigHelp
	sig.Signatures = []SrcIntelSigInfo{SrcIntelSigInfo{Label: "Signature", Documentation: SrcIntelDoc{IsTrusted: true, Value: "Helpful **doc** `comment`s.."}}}
	sig.Signatures[0].Parameters = []SrcIntelSigParam{SrcIntelSigParam{Label: "Parameter 1", Documentation: SrcIntelDoc{IsTrusted: true, Value: "Every argument gets a *helpful* `doc` comment."}}}
	return &sig
}

func (*SrcIntelBase) Symbols(srcLens *SrcLens, query string, curFileOnly bool) (all udev.SrcMsgs) {
	if curFileOnly {
		const symMinInvalid = int(SYM_MIN_INVALID)
		for i := 0; i < symMinInvalid; i++ {
			all = append(all,
				&udev.SrcMsg{Flag: i, Msg: Symbol(i).String(), Ref: srcLens.FilePath,
					Misc:   Strf("flag: %d", i),
					Pos1Ch: 1, Pos1Ln: i + 1, Pos2Ch: 1, Pos2Ln: i + 1,
				},
			)
		}
	}
	return
}
