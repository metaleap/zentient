package z

type msgResp struct {
	ReqID  int64  `json:"i"`
	ErrMsg string `json:"e,omitempty"`

	MetaCmdsMenu *metaCmdsMenu `json:"mcM,omitempty"`
}

func (me *msgResp) to(req *msgReq) {
	h := false // handled?
	h = h || metaCmdsHandle(req, me)
	h = h || codeFmtHandle(req, me)

	if !h {
		me.ErrMsg = strf("Invalid MsgID %d", req.MsgID)
	}
}
