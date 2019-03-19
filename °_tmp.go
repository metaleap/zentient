package z

import (
	"github.com/go-leap/str"
)

type fooResp struct {
	IpcID    IpcIDs        `json:"ii,omitempty"`
	ReqID    int64         `json:"ri,omitempty"`
	ErrMsg   string        `json:"err,omitempty"`
	SrcIntel *srcIntelResp `json:"sI,omitempty"`
	SrcDiags *diagResp     `json:"srcDiags,omitempty"`
	*ipcReq
	SrcMods    SrcLenses `json:"srcMods,omitempty"`
	muhPrivate int
	SrcActions []EditorAction `json:"srcActions,omitempty"`
	Extras     *ExtrasResp    `json:"extras,omitempty"`
	*SrcLens
	ustr.Pats
	Menu        *menuResp   `json:"menu,omitempty"`
	CaddyUpdate *Caddy      `json:"caddy,omitempty"`
	Val         interface{} `json:"val,omitempty"`
}

func (this *fooResp) tmp() {
	this.Pats = nil
}
