package z

import (
	"encoding/json"
	"strings"
)

type msgIDs uint8

const (
	_ msgIDs = iota

	msgID_coreCmds_ListAll

	msgID_srcFmt_ListAll
	msgID_srcFmt_InfoLink
)

type msgReq struct {
	ReqID   int64       `json:"i"`
	MsgID   msgIDs      `json:"m"`
	MsgArgs interface{} `json:"a"`

	FilePath string `json:"fp"`
	SrcFull  string `json:"sf"`
	SrcSel   string `json:"ss"`
	PosOff   int    `json:"po"`
	PosLn    int    `json:"pl"`
	PosCol   int    `json:"pc"`
}

func reqDecodeAndRespond(jsonreq string) *msgResp {
	var req msgReq
	var resp msgResp
	if !Lang.Enabled {
		resp.ErrMsg = strf("%s does not appear to be installed on this machine.", Lang.Title)
	} else if err := json.NewDecoder(strings.NewReader(jsonreq)).Decode(&req); err == nil {
		resp.ReqID = req.ReqID
		resp.to(&req)
	} else {
		resp.ErrMsg = err.Error()
	}
	return &resp
}
