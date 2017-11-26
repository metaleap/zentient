package z

type iResponder interface {
	handle(*msgReq, *msgResp) bool
}

type msgResp struct {
	ReqID  int64  `json:"i"`
	ErrMsg string `json:"e,omitempty"`

	MetaCmdsMenu *metaCmdsMenu `json:"mcM,omitempty"`
}

func (me *msgResp) catch() {
	if except := recover(); except != nil {
		me.ErrMsg = strf("%v", except)
	}
}

func (me *msgResp) to(req *msgReq) {
	defer me.catch()
	h := false // handled?
	h = h || metaCmdsHandle(req, me)
	h = h || Lang.CodeFmt.handle(req, me)
	if !h {
		me.ErrMsg = strf("Invalid MsgID %d", req.MsgID)
	}
}
