package z

import (
	"encoding/json"
)

type MsgIDs uint8

const (
	REQ_CMDS_LISTALL MsgIDs = 1
)

type MsgReq struct {
	ReqID int64  `json:"i"`
	MsgID MsgIDs `json:"m"`
}

type MsgResp struct {
	ReqID int64
	Err   string `json:"e"`
}

func (me *MsgResp) encode() (jsonresp string, err error) {
	var data []byte
	if data, err = json.Marshal(me); err == nil {
		jsonresp = string(data)
	}
	return
}

func handleReq(jsonreq string) *MsgResp {
	var resp MsgResp
	if req, err := reqDecode(jsonreq); err == nil {
		resp.ReqID = req.ReqID
		switch req.MsgID {
		case REQ_CMDS_LISTALL:
		default:
			resp.Err = strf("Invalid MsgID %d", req.MsgID)
		}
	} else {
		resp.Err = err.Error()
	}
	return &resp
}

func reqDecode(jsonreq string) (*MsgReq, error) {
	var req MsgReq
	if err := json.Unmarshal([]byte(jsonreq), &req); err != nil {
		return nil, err
	}
	return &req, nil
}
