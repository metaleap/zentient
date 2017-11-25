package z

import (
	"encoding/json"
)

type MsgResp struct {
	ReqID  int64  `json:"i"`
	ErrMsg string `json:"e,omitempty"`

	Menu []*MsgRespPick `json:"m,omitempty"`
}

type MsgRespPick struct {
	ID    int    `json:"i"`
	Title string `json:"t"`
	Desc  string `json:"d,omitempty"`
}

func (me *MsgResp) encode() (jsonresp string, err error) {
	var data []byte
	if data, err = json.Marshal(me); err == nil {
		jsonresp = string(data)
	}
	return
}
