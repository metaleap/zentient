package z

import (
	"encoding/json"
	"strings"
)

type iDispatcher interface {
	dispatch(*msgReq, *msgResp) bool
	Init()
}

type msgReq struct {
	ReqID   int64       `json:"ri"`
	MsgID   msgIDs      `json:"mi"`
	MsgArgs interface{} `json:"ma"`

	SrcLens *SrcLens `json:"sl"`
}

func reqDecodeAndRespond(jsonreq string) *msgResp {
	var req msgReq
	var resp msgResp
	if !Lang.Enabled {
		resp.ErrMsg = Strf("%s does not appear to be installed on this machine.", Lang.Title)
	} else if Prog.Cfg.err != nil {
		resp.ErrMsg = Strf("Your %s is currently broken: either fix it or delete it, then reload Zentient.", Prog.Cfg.filePath)
	}
	if err := json.NewDecoder(strings.NewReader(jsonreq)).Decode(&req); err != nil {
		resp.ErrMsg = err.Error()
	} else if resp.ReqID = req.ReqID; resp.ErrMsg == "" {
		resp.to(&req)
	}
	if resp.ErrMsg != "" {
		resp.MsgID = req.MsgID
	}
	return &resp
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
	for _, h := range dispatchers {
		if h.dispatch(req, me) {
			return
		}
	}
	if req.MsgID < MSGID_CORECMDS_PALETTE || req.MsgID >= MSGID_MIN_INVALID {
		me.ErrMsg = Strf("Invalid MsgID %s", req.MsgID)
	} else {
		me.ErrMsg = Strf("The requested feature (MsgID %s) wasn't implemented for **%s**.", req.MsgID, Lang.Title)
	}
}
