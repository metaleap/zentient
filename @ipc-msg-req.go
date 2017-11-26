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
	msgID_srcFmt_SetDef
	msgID_srcFmt_RunOnFile
	msgID_srcFmt_RunOnSel
)

type msgReq struct {
	ReqID   int64       `json:"ri"`
	MsgID   msgIDs      `json:"mi"`
	MsgArgs interface{} `json:"ma"`

	SrcLoc *SrcLoc `json:"sl"`
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
