package z

import (
	"github.com/metaleap/go-util/dev"
)

type iSrcIntel interface {
	iDispatcher

	ComplDetails(*SrcLens, string, *SrcIntelCompl)
	ComplItems(*SrcLens) []SrcIntelCompl
	Highlights(*SrcLens, string) []SrcRange
	Hovers(*SrcLens) []SrcIntelHover
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) udev.SrcMsgs
}

type srcIntelResp struct {
	Cmpl       []SrcIntelCompl  `json:"cmpl,omitempty"`
	Hovers     []SrcIntelHover  `json:"hovs,omitempty"`
	Symbols    udev.SrcMsgs     `json:"syms,omitempty"`
	Highlights []SrcRange       `json:"high,omitempty"`
	Signature  *SrcIntelSigHelp `json:"sig,omitempty"`
}

type SrcIntelCompl struct {
	Label         string      `json:"label"`
	Kind          Completion  `json:"kind,omitempty"`
	Detail        string      `json:"detail,omitempty"`
	Documentation SrcIntelDoc `json:"documentation,omitempty"`
	SortText      string      `json:"sortText,omitempty"`
	FilterText    string      `json:"filterText,omitempty"`
	InsertText    string      `json:"insertText,omitempty"`
	CommitChars   []string    `json:"commitCharacters,omitempty"`
	// Range               Range      `json:"Range,omitempty"`
	// AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
	// Command             Command    `json:"command,omitempty"`
}

type SrcIntelDoc struct {
	Value     string `json:"value"`
	IsTrusted bool   `json:"isTrusted"`
}

type SrcIntelHover struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}

type SrcIntelSigHelp struct {
	ActiveSignature int               `json:"activeSignature"`
	ActiveParameter int               `json:"activeParameter"`
	Signatures      []SrcIntelSigInfo `json:"signatures"`
}

type SrcIntelSigInfo struct {
	Label         string             `json:"label"`
	Documentation SrcIntelDoc        `json:"documentation,omitempty"`
	Parameters    []SrcIntelSigParam `json:"parameters,omitempty"`
}

type SrcIntelSigParam struct {
	Label         string      `json:"label"`
	Documentation SrcIntelDoc `json:"documentation,omitempty"`
}

type SrcIntelBase struct {
	Impl iSrcIntel
}

func (me *SrcIntelBase) Init() {
	dispatchers = append(dispatchers, me.Impl)
}

func (me *SrcIntelBase) dispatch(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_SRCINTEL_HOVER:
		me.onHover(req, resp)
	case MSGID_SRCINTEL_SYMS_FILE, MSGID_SRCINTEL_SYMS_PROJ:
		me.onSyms(req, resp)
	case MSGID_SRCINTEL_CMPL_ITEMS:
		me.onCmplItems(req, resp)
	case MSGID_SRCINTEL_CMPL_DETAILS:
		me.onCmplDetails(req, resp)
	case MSGID_SRCINTEL_HIGHLIGHTS:
		me.onHighlights(req, resp)
	case MSGID_SRCINTEL_SIGNATURE:
		me.onSignature(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) onCmplItems(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Cmpl: me.Impl.ComplItems(req.SrcLens)}
}

func (me *SrcIntelBase) onCmplDetails(req *msgReq, resp *msgResp) {
	itemtext, _ := req.MsgArgs.(string)
	resp.SrcIntel = &srcIntelResp{Cmpl: make([]SrcIntelCompl, 1, 1)}
	me.Impl.ComplDetails(req.SrcLens, itemtext, &(resp.SrcIntel.Cmpl[0]))
}

func (me *SrcIntelBase) onHighlights(req *msgReq, resp *msgResp) {
	curword, _ := req.MsgArgs.(string)
	resp.SrcIntel = &srcIntelResp{Highlights: me.Impl.Highlights(req.SrcLens, curword)}
}

func (me *SrcIntelBase) onHover(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Hovers: me.Impl.Hovers(req.SrcLens)}
}

func (me *SrcIntelBase) onSignature(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Signature: me.Impl.Signature(req.SrcLens)}
}

func (me *SrcIntelBase) onSyms(req *msgReq, resp *msgResp) {
	var query string
	if req.MsgID == MSGID_SRCINTEL_SYMS_PROJ {
		query, _ = req.MsgArgs.(string)
	}
	resp.SrcIntel = &srcIntelResp{Symbols: me.Impl.Symbols(req.SrcLens, query, req.MsgID == MSGID_SRCINTEL_SYMS_FILE)}
}
