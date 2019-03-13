package z

type xpcResp struct {
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

// func (this *xpcResp) MarshalJSON() ([]byte, error) {
// 	return []byte("{}"), nil
// }
