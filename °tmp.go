package z

import (
	"encoding/json"
	"github.com/go-leap/str"
)

type fooResp struct {
	IpcID    IpcIDs `json:"ii"`
	ReqID    int64  `json:"ri,omitempty"`
	Flag     bool
	ErrMsg   string    `json:"err,omitempty"`
	SrcIntel *SrcIntel `json:"sI,omitempty"`
	SrcDiags *Diags    `json:"srcDiags,omitempty"`
	*IpcReq
	SrcMods    SrcLenses `json:"srcMods,omitempty"`
	muhPrivate int
	SrcActions []EditorAction `json:"srcActions,omitempty"`
	Extras     *Extras        `json:"extras,omitempty"`
	*SrcLens
	Fn   func()
	Link *fooResp
	Ch   chan bool
	ustr.Pats
	Menu        *MenuResponse `json:"menu,omitempty"`
	Nope        string        `json:"-"`
	CaddyUpdate *Caddy        `json:"caddy,omitempty"`
	Val         interface{}   `json:"valya"`
}

func (me *SrcModEdit) unm(_ []byte) (err error) {
	j := json.NewDecoder(nil)
	var k, v json.Token

	if _, err = j.Token(); err == nil {
		for j.More() {
			if k, err = j.Token(); err == nil {
				if v, err = j.Token(); err == nil && v != nil {
					if k == "At" {

					} else if k == "Val" {
						me.Val = v.(string)
					}
				}
			}
		}
		_, err = j.Token()
	}
	return
}

func tokIsDelim(t json.Token, err error) error {
	if err == nil {

	}
	return err
}
