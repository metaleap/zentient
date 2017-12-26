package z

func (me *SrcIntelBase) DefImpl(srcLens *SrcLens) SrcLenses {
	return me.References(srcLens, true)
}
