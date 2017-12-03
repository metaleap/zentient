package z

import (
	"encoding/json"
	"strings"
)

type msgReq struct {
	ReqID   int64       `json:"ri"`
	MsgID   msgIDs      `json:"mi"`
	MsgArgs interface{} `json:"ma"`

	SrcLens *SrcLens `json:"sl"`
}

func reqDecodeAndRespond(jsonreq string) *msgResp {
	var req msgReq
	var resp msgResp
	if !Lang.Enabled {
		resp.ErrMsg = Strf("%s does not appear to be installed on this machine.", Lang.Title)
	} else if Prog.Cfg.err != nil {
		resp.ErrMsg = Strf("Your %s is currently broken: either fix it or delete it, then reload Zentient.", Prog.Cfg.filePath)
	}
	if err := json.NewDecoder(strings.NewReader(jsonreq)).Decode(&req); err != nil {
		resp.ErrMsg = err.Error()
	} else if resp.ReqID = req.ReqID; resp.ErrMsg == "" {
		resp.to(&req)
	}
	if resp.ErrMsg != "" {
		resp.MsgID = req.MsgID
	}
	return &resp
}
