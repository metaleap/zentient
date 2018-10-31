package z

import (
	"encoding/json"
	"strings"
)

type iDispatcher interface {
	dispatch(*ipcReq, *ipcResp) bool
	Init()
}

type IObjSnap interface {
	ObjSnapPrefix() string
	ObjSnap(string) interface{}
}

type ipcReq struct {
	ReqID   int64       `json:"ri"`
	IpcID   IpcIDs      `json:"ii"`
	IpcArgs interface{} `json:"ia"`

	ProjUpd *WorkspaceChanges `json:"projUpd"`
	SrcLens *SrcLens          `json:"srcLens"`
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
	} else if req.IpcID == IPCID_OBJ_SNAPSHOT {
		objpath, _ := req.IpcArgs.(string)
		found := false
		for _, objsnp := range Prog.objSnappers {
			if pref := objsnp.ObjSnapPrefix(); strings.HasPrefix(objpath, pref) {
				found, resp.Val = true, objsnp.ObjSnap(objpath[len(pref):])
				break
			}
		}
		if !found {
			resp.ErrMsg = BadMsg(req.IpcID.String()+" path", objpath)
		}
	} else {
		resp.to(&req)
	}
	if resp.ReqID = req.ReqID; resp.ErrMsg != "" {
		resp.IpcID = req.IpcID
	}
	return &resp
}

type ipcResp struct {
	IpcID       IpcIDs         `json:"ii,omitempty"`
	ReqID       int64          `json:"ri,omitempty"`
	ErrMsg      string         `json:"err,omitempty"`
	SrcIntel    *srcIntelResp  `json:"sI,omitempty"`
	SrcDiags    *diagResp      `json:"srcDiags,omitempty"`
	SrcMods     SrcLenses      `json:"srcMods,omitempty"`
	SrcActions  []EditorAction `json:"srcActions,omitempty"`
	Extras      *ExtrasResp    `json:"extras,omitempty"`
	Menu        *menuResp      `json:"menu,omitempty"`
	CaddyUpdate *Caddy         `json:"caddy,omitempty"`
	Val         interface{}    `json:"val,omitempty"`
}

func (this *ipcResp) postProcess() {
	if this.Menu != nil && this.Menu.SubMenu != nil && this.Menu.SubMenu.Items == nil {
		// handles better on the client-side (and UX-wise) --- instead of a "silent nothing", show an empty menu ("nothing to choose from")
		this.Menu.SubMenu.Items = MenuItems{}
	}
}

func (this *ipcResp) onResponseReady() {
	if except := recover(); except != nil {
		this.ErrMsg = Strf("%v", except)
	}
	if this.ErrMsg != "" {
		this.ErrMsg = Strf("[%s] %s", Prog.Name, strings.TrimPrefix(this.ErrMsg, Prog.Name+": "))
		//	zero out almost-everything for a leaner response. req-ID is only added in afterwards anyways
		*this = ipcResp{ErrMsg: this.ErrMsg}
	}
}

func (this *ipcResp) to(req *ipcReq) {
	defer this.onResponseReady()
	for _, disp := range Prog.dispatchers {
		if disp.dispatch(req, this) {
			this.postProcess()
			return
		}
	}
	if !req.IpcID.Valid() {
		this.ErrMsg = BadMsg("IpcID", req.IpcID.String())
	} else {
		this.ErrMsg = Strf("The requested feature `%s` wasn't yet implemented for %s.", req.IpcID, Lang.Title)
	}
}

func (this *ipcResp) withExtras() *ipcResp {
	this.Extras = &ExtrasResp{}
	return this
}

func (this *ipcResp) withMenu() *menuResp {
	this.Menu = &menuResp{}
	return this.Menu
}
func (this *ipcResp) withSrcIntel() *ipcResp {
	this.SrcIntel = &srcIntelResp{}
	return this
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
