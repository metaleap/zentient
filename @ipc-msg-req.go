package z

import (
	"encoding/json"
	"strings"
)

type msgIDs uint8

const (
	_ msgIDs = iota

	msgID_coreCmds_ListAll

	msgID_srcFmt_SetDefMenu
	msgID_srcFmt_SetDefPick
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
		resp.ErrMsg = Strf("%s does not appear to be installed on this machine.", Lang.Title)
	} else if Prog.Cfg.err != nil {
		resp.ErrMsg = Strf("Your %s is currently broken: either fix it or delete it, then reload Zentient.", Prog.Cfg.filePath)
	} else if err := json.NewDecoder(strings.NewReader(jsonreq)).Decode(&req); err == nil {
		resp.ReqID = req.ReqID
		resp.to(&req)
	} else {
		resp.ErrMsg = err.Error()
	}
	return &resp
}
