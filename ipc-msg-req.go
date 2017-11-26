package z

import (
	"encoding/json"
)

type MsgIDs uint8

const (
	_ MsgIDs = iota
	REQ_META_CMDS_LISTALL
)

type MsgReq struct {
	ReqID int64  `json:"i"`
	MsgID MsgIDs `json:"m"`

	FilePath string                 `json:"fp"`
	SrcFull  string                 `json:"sf"`
	SrcSel   string                 `json:"ss"`
	PosOff   int                    `json:"po"`
	PosLn    int                    `json:"pl"`
	PosCol   int                    `json:"pc"`
	Args     map[string]interface{} `json:"a"`
}

func reqDecodeAndRespond(jsonreq string) *MsgResp {
	var req MsgReq
	var resp MsgResp
	if err := json.Unmarshal([]byte(jsonreq), &req); err == nil {
		resp.ReqID = req.ReqID
		resp.to(&req)
	} else {
		resp.ErrMsg = err.Error()
	}
	return &resp
}
