package z

type iSrcIntel interface {
	iHandler

	Hovers(*SrcLens) []SrcIntelHover
}

type srcIntelResp struct {
	Hovers []SrcIntelHover `json:"h,omitempty"`
}

type SrcIntelHover struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}

type SrcIntelBase struct {
	Self iSrcIntel
}

func (me *SrcIntelBase) Init() {
	handlers = append(handlers, me.Self)
}

func (_ *SrcIntelBase) Hovers(srcLens *SrcLens) (hovs []SrcIntelHover) {
	hovs = append(hovs,
		SrcIntelHover{Value: Strf("Hovers not yet implemented for **%s** by `%s`", Lang.Title, Prog.name)},
		SrcIntelHover{Value: Strf("File: %s", srcLens.FilePath), Language: "plaintext"},
		SrcIntelHover{Value: Strf("Line/Char/Offset: %v", *srcLens.Pos)},
	)
	return
}

func (me *SrcIntelBase) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case msgID_srcIntel_Hover:
		me.handle_Hover(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) handle_Hover(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Hovers: me.Self.Hovers(req.SrcLens)}
}
