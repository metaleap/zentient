package z

import (
	"encoding/json"
)

type msgIDs uint8

const (
	_ msgIDs = iota

	msgID_metaCmds_ListAll

	msgID_codeFmt_ListAll
)

type msgReq struct {
	ReqID int64  `json:"i"`
	MsgID msgIDs `json:"m"`

	FilePath string                 `json:"fp"`
	SrcFull  string                 `json:"sf"`
	SrcSel   string                 `json:"ss"`
	PosOff   int                    `json:"po"`
	PosLn    int                    `json:"pl"`
	PosCol   int                    `json:"pc"`
	Args     map[string]interface{} `json:"a"`
}

func reqDecodeAndRespond(jsonreq string) *msgResp {
	var req msgReq
	var resp msgResp
	if err := json.Unmarshal([]byte(jsonreq), &req); err == nil {
		resp.ReqID = req.ReqID
		resp.to(&req)
	} else {
		resp.ErrMsg = err.Error()
	}
	return &resp
}
