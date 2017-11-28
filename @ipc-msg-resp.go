package z

type iHandler interface {
	handle(*msgReq, *msgResp) bool
	Init()
}

type msgResp struct {
	ReqID  int64  `json:"ri"`
	ErrMsg string `json:"e,omitempty"`

	CoreCmdsMenu *coreCmdsMenu `json:"menu,omitempty"`
	WebsiteURL   string        `json:"url,omitempty"`
	Note         string        `json:"note,omitempty"`
}

type msgArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

func (me *msgResp) catch() {
	if except := recover(); except != nil {
		me.ErrMsg = strf("%v", except)
	}
}

func (me *msgResp) to(req *msgReq) {
	defer me.catch()
	for _, h := range handlers {
		if h.handle(req, me) {
			return
		}
	}
	me.ErrMsg = strf("Invalid MsgID %d", req.MsgID)
}
