package z

import (
	"encoding/json"
	"strings"
)

type iDispatcher interface {
	dispatch(*ipcReq, *ipcResp) bool
	Init()
}

type ipcReq struct {
	ReqID   int64       `json:"ri"`
	IpcID   ipcIDs      `json:"ii"`
	IpcArgs interface{} `json:"ia"`

	SrcLens *SrcLens `json:"sl"`
}

func ipcDecodeReqAndRespond(jsonreq string) *ipcResp {
	var req ipcReq
	var resp ipcResp
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
		resp.IpcID = req.IpcID
	}
	return &resp
}

type ipcResp struct {
	ReqID  int64  `json:"ri"`
	ErrMsg string `json:"e,omitempty"`

	IpcID       ipcIDs         `json:"ii,omitempty"`
	Menu        *MenuResp      `json:"menu,omitempty"`
	Extras      *ExtrasResp    `json:"extras,omitempty"`
	SrcIntel    *srcIntelResp  `json:"srcIntel,omitempty"`
	SrcMods     []*SrcLens     `json:"srcMods,omitempty"`
	SrcActions  []EditorAction `json:"srcActions,omitempty"`
	CaddyUpdate *Caddy         `json:"caddy,omitempty"`
}

func (me *ipcResp) onResponseReady() {
	if except := recover(); except != nil {
		me.ErrMsg = Strf("%v", except)
	}
	if me.ErrMsg != "" {
		me.ErrMsg = Strf("[%s] %s", Prog.name, me.ErrMsg)
		//	zero out almost-everything for a leaner response
		*me = ipcResp{ErrMsg: me.ErrMsg}
	}
}

func (me *ipcResp) to(req *ipcReq) {
	defer me.onResponseReady()
	for _, h := range dispatchers {
		if h.dispatch(req, me) {
			return
		}
	}
	if req.IpcID < IPCID_MENUS_MAIN || req.IpcID >= IPCID_MIN_INVALID {
		me.ErrMsg = Strf("Invalid IpcID %d", req.IpcID)
	} else {
		me.ErrMsg = Strf("The requested feature `%s` wasn't yet implemented for __%s__.", req.IpcID, Lang.Title)
	}
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
