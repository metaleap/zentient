package z

import (
	"encoding/json"
)

type MsgResp struct {
	ReqID  int64  `json:"i"`
	ErrMsg string `json:"e,omitempty"`

	MetaCmdsMenu *MetaCmdsMenu `json:"mcM,omitempty"`
}

func (me *MsgResp) encode() (jsonresp string, err error) {
	var data []byte
	if data, err = json.Marshal(me); err == nil {
		jsonresp = string(data)
	}
	return
}

func (resp *MsgResp) to(req *MsgReq) {
	switch req.MsgID {
	case REQ_META_CMDS_LISTALL:
		handleMetaCmdsListAll(req, resp)
	default:
		resp.ErrMsg = strf("Invalid MsgID %d", req.MsgID)
	}
}
