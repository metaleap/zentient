package z

import (
	// "encoding/json"
	// "github.com/go-leap/std"
	"github.com/go-leap/str"
	// "strconv"
)

type fooResp struct {
	IpcID    IpcIDs        `json:"ii,omitempty"`
	ReqID    int64         `json:"ri,omitempty"`
	ErrMsg   string        `json:"err,omitempty"`
	SrcIntel *srcIntelResp `json:"sI,omitempty"`
	SrcDiags *diagResp     `json:"srcDiags,omitempty"`
	*ipcReq
	SrcMods    SrcLenses `json:"srcMods,omitempty"`
	muhPrivate int
	SrcActions []EditorAction `json:"srcActions,omitempty"`
	Extras     *ExtrasResp    `json:"extras,omitempty"`
	*SrcLens
	ustr.Pats
	Menu        *menuResp   `json:"menu,omitempty"`
	CaddyUpdate *Caddy      `json:"caddy,omitempty"`
	Val         interface{} `json:"val,omitempty"`
}

// func (this *fooResp) tmp() {
// 	this.Pats = nil
// }

// // MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
// func (this *fooResp) marshalJSON() (r []byte, err error) {

// 	var buf ustd.Buf

// 	buf.WriteString("{\"")
// 	buf.WriteString("ReqID")
// 	buf.WriteString("\":")
// 	buf.WriteString(strconv.FormatInt(this.ReqID, 10))

// 	buf.WriteString(",\"")
// 	buf.WriteString("Str")
// 	buf.WriteString("\":")
// 	buf.WriteString(strconv.Quote(this.Str))

// 	buf.WriteString(",\"")
// 	buf.WriteString("Flag")
// 	buf.WriteString("\":")
// 	buf.WriteString(strconv.FormatInt(int64(this.Flag), 10))

// 	buf.WriteString(",\"")
// 	buf.WriteString("IpcArgs")
// 	buf.WriteString("\":")
// 	enc := json.NewEncoder(&buf)
// 	enc.SetEscapeHTML(false)
// 	if err = enc.Encode(this.IpcArgs); err != nil {
// 		return
// 	}
// 	buf.TrimSuffix('\n')

// 	buf.WriteString(",\"")
// 	buf.WriteString("CrLf")
// 	buf.WriteString("\":")
// 	buf.WriteString(strconv.FormatBool(this.CrLf))

// 	buf.WriteByte('}')
// 	r = buf.Bytes()
// 	return
// }
