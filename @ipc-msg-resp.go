package z

const (
	DIAG_SEV_ERR  = 0
	DIAG_SEV_WARN = 1
	DIAG_SEV_INFO = 2
	DIAG_SEV_HINT = 3
)

const (
	SYM_FILE          = 0
	SYM_MODULE        = 1
	SYM_NAMESPACE     = 2
	SYM_PACKAGE       = 3
	SYM_CLASS         = 4
	SYM_METHOD        = 5
	SYM_PROPERTY      = 6
	SYM_FIELD         = 7
	SYM_CONSTRUCTOR   = 8
	SYM_ENUM          = 9
	SYM_INTERFACE     = 10
	SYM_FUNCTION      = 11
	SYM_VARIABLE      = 12
	SYM_CONSTANT      = 13
	SYM_STRING        = 14
	SYM_NUMBER        = 15
	SYM_BOOLEAN       = 16
	SYM_ARRAY         = 17
	SYM_OBJECT        = 18
	SYM_KEY           = 19
	SYM_NULL          = 20
	SYM_ENUMMEMBER    = 21
	SYM_STRUCT        = 22
	SYM_EVENT         = 23
	SYM_OPERATOR      = 24
	SYM_TYPEPARAMETER = 25
)

const (
	CMPL_TEXT          = 0
	CMPL_METHOD        = 1
	CMPL_FUNCTION      = 2
	CMPL_CONSTRUCTOR   = 3
	CMPL_FIELD         = 4
	CMPL_VARIABLE      = 5
	CMPL_CLASS         = 6
	CMPL_INTERFACE     = 7
	CMPL_MODULE        = 8
	CMPL_PROPERTY      = 9
	CMPL_UNIT          = 10
	CMPL_VALUE         = 11
	CMPL_ENUM          = 12
	CMPL_KEYWORD       = 13
	CMPL_SNIPPET       = 14
	CMPL_COLOR         = 15
	CMPL_FILE          = 16
	CMPL_REFERENCE     = 17
	CMPL_FOLDER        = 18
	CMPL_ENUMMEMBER    = 19
	CMPL_CONSTANT      = 20
	CMPL_STRUCT        = 21
	CMPL_EVENT         = 22
	CMPL_OPERATOR      = 23
	CMPL_TYPEPARAMETER = 24
)

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
		me.ErrMsg = Strf("Invalid MsgID %d", req.MsgID)
	} else {
		me.ErrMsg = Strf("The requested feature (MsgID %d) wasn't implemented for **%s**.", req.MsgID, Lang.Title)
	}
}
