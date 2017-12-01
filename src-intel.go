package z

type iSrcIntel interface {
	iHandler
}

type srcIntelResp struct {
	Hovers []srcHover `json:"h,omitempty"`
}

type srcHover struct {
	Value    string `json:"value"`
	Language string `json:"language,omitempty"`
}

type SrcIntelBase struct {
	Self iSrcIntel
}

func (me *SrcIntelBase) Init() {
	handlers = append(handlers, me.Self)
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

func (*SrcIntelBase) handle_Hover(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{}
	resp.SrcIntel.Hovers = append(resp.SrcIntel.Hovers,
		srcHover{Value: "test **one** is _live_"},
		srcHover{Value: "test **two** is a _go_", Language: "plaintext"},
		srcHover{Value: "func main() { println(123) }", Language: "go"},
	)
}
