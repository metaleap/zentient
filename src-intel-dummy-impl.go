package z

func (me *SrcIntelBase) DefSym(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}

func (me *SrcIntelBase) DefType(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}

func (me *SrcIntelBase) DefImpl(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}
