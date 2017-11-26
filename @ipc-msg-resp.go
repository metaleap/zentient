package z

type iResponder interface {
	handle(*msgReq, *msgResp) bool
	Init()
}

type msgResp struct {
	ReqID  int64  `json:"i"`
	ErrMsg string `json:"e,omitempty"`

	CoreCmdsMenu *coreCmdsMenu `json:"ccM,omitempty"`
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
