package z

type iHandler interface {
	handle(*msgReq, *msgResp) bool
	Init()
}

type msgResp struct {
	ReqID          int64  `json:"ri"`
	ErrMsg         string `json:"e,omitempty"`
	ErrMsgFromTool bool   `json:"et,omitempty"`

	MsgID        msgIDs        `json:"mi,omitempty"`
	CoreCmdsMenu *coreCmdsMenu `json:"menu,omitempty"`
	WebsiteURL   string        `json:"url,omitempty"`
	NoteInfo     string        `json:"info,omitempty"`
	NoteWarn     string        `json:"warn,omitempty"`
	MsgAction    string        `json:"action,omitempty"`
	SrcMod       *SrcLens      `json:"srcMod,omitempty"`
}

type msgArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

func (me *msgResp) onResponseReady(req *msgReq) {
	if except := recover(); except != nil {
		me.ErrMsg = Strf("%v", except)
	}
	if me.ErrMsg != "" {
		me.ErrMsg = Strf("[%s] %s", Prog.name, me.ErrMsg)
		//	zero out nearly-everything for a leaner response
		*me = msgResp{ErrMsg: me.ErrMsg, ErrMsgFromTool: me.ErrMsgFromTool, ReqID: me.ReqID, MsgID: req.MsgID}
	}
}

func (me *msgResp) to(req *msgReq) {
	defer me.onResponseReady(req)
	for _, h := range handlers {
		if h.handle(req, me) {
			return
		}
	}
	me.ErrMsg = Strf("Invalid MsgID %d", req.MsgID)
}
