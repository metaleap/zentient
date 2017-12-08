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
	if err := json.NewDecoder(strings.NewReader(jsonreq)).Decode(&req); err != nil {
		resp.ErrMsg = err.Error()
	} else if !Lang.Enabled {
		resp.ErrMsg = Strf("%s does not appear to be installed on this machine.", Lang.Title)
	} else if Prog.Cfg.err != nil {
		resp.ErrMsg = Strf("Your %s is currently broken: either fix it or delete it, then reload Zentient.", Prog.Cfg.filePath)
	} else {
		resp.to(&req)
	}
	if resp.ReqID = req.ReqID; resp.ErrMsg != "" {
		resp.MsgID = req.MsgID
	}
	return &resp
}

type msgResp struct {
	ReqID  int64  `json:"ri"`
	ErrMsg string `json:"e,omitempty"`

	MsgID      msgIDs         `json:"mi,omitempty"`
	Menu       *MenuResp      `json:"menu,omitempty"`
	Extras     *ExtrasResp    `json:"extras,omitempty"`
	SrcIntel   *srcIntelResp  `json:"srcIntel,omitempty"`
	SrcMods    []*SrcLens     `json:"srcMods,omitempty"`
	SrcActions []EditorAction `json:"srcActions,omitempty"`
}

type EditorAction struct {
	Title     string        `json:"title"`
	Cmd       string        `json:"command"`
	Hint      string        `json:"tooltip,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

type InfoTip struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}

func (me *msgResp) onResponseReady() {
	if except := recover(); except != nil {
		me.ErrMsg = Strf("%v", except)
	}
	if me.ErrMsg != "" {
		me.ErrMsg = Strf("[%s] %s", Prog.name, me.ErrMsg)
		//	zero out almost-everything for a leaner response
		*me = msgResp{ErrMsg: me.ErrMsg}
	}
}

func (me *msgResp) to(req *msgReq) {
	defer me.onResponseReady()
	for _, h := range dispatchers {
		if h.dispatch(req, me) {
			return
		}
	}
	if req.MsgID < MSGID_MENUS_MAIN || req.MsgID >= MSGID_MIN_INVALID {
		me.ErrMsg = Strf("Invalid MsgID %d", req.MsgID)
	} else {
		me.ErrMsg = Strf("The requested feature `%s` wasn't yet implemented for __%s__.", req.MsgID, Lang.Title)
	}
}
