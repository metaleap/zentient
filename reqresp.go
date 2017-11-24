package z

import (
	"encoding/json"
)

type MsgReq struct {
	ReqID int64  `json:"i"`
	MsgID string `json:"m"`
}

type MsgResp struct {
	ReqID int64  `json:"i"`
	Err   string `json:"e"`
}

func (me *MsgResp) encode() (jsonresp string, err error) {
	var data []byte
	if data, err = json.Marshal(me); err == nil {
		jsonresp = string(data)
	}
	return
}

const (
	REQ_DO_FMT = "DF:"
)

func handleReq(jsonreq string) (resp *MsgResp) {
	resp = &MsgResp{}
	if req, err := reqDecode(jsonreq); err == nil {
		resp.ReqID = req.ReqID
	} else {
		resp.Err = err.Error()
	}
	return
}

func reqDecode(jsonreq string) (*MsgReq, error) {
	var req MsgReq
	if err := json.Unmarshal([]byte(jsonreq), &req); err != nil {
		return nil, err
	}
	return &req, nil
}
