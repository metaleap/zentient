package z

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
