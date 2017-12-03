package z

type iHandler interface {
	handle(*msgReq, *msgResp) bool
	Init()
}

type msgResp struct {
	ReqID          int64  `json:"ri"`
	ErrMsg         string `json:"e,omitempty"`
	ErrMsgFromTool bool   `json:"et,omitempty"`

	MsgID    msgIDs        `json:"mi,omitempty"`
	CoreCmd  *coreCmdResp  `json:"coreCmd,omitempty"`
	SrcIntel *srcIntelResp `json:"srcIntel,omitempty"`
	SrcMod   *SrcLens      `json:"srcMod,omitempty"`
}

type msgArgPrompt struct {
	Prompt      string `json:"prompt,omitempty"`
	Placeholder string `json:"placeHolder,omitempty"`
	Value       string `json:"value,omitempty"`
}

func (me *msgResp) onResponseReady() {
	if except := recover(); except != nil {
		me.ErrMsg = Strf("%v", except)
	}
	if me.ErrMsg != "" {
		me.ErrMsg = Strf("[%s] %s", Prog.name, me.ErrMsg)
		//	zero out nearly-everything for a leaner response
		*me = msgResp{ErrMsg: me.ErrMsg, ErrMsgFromTool: me.ErrMsgFromTool, ReqID: me.ReqID}
	}
}

func (me *msgResp) to(req *msgReq) {
	defer me.onResponseReady()
	for _, h := range handlers {
		if h.handle(req, me) {
			return
		}
	}
	if req.MsgID < MSGID_CORECMDS_PALETTE || req.MsgID >= MSGID_MIN_INVALID {
		me.ErrMsg = Strf("Invalid MsgID %s", req.MsgID)
	} else {
		me.ErrMsg = Strf("The requested feature (MsgID %s) wasn't implemented for **%s**.", req.MsgID, Lang.Title)
	}
}
