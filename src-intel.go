package z

type ISrcIntel interface {
	iDispatcher

	ComplDetails(*SrcLens, string, *SrcIntelCompl)
	ComplItems(*SrcLens) []SrcIntelCompl
	DefSym(*SrcLens) []*SrcLens
	DefType(*SrcLens) []*SrcLens
	DefImpl(*SrcLens) []*SrcLens
	Highlights(*SrcLens, string) []SrcRange
	Hovers(*SrcLens) []InfoTip
	References(*SrcLens, bool) []*SrcLens
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) []*SrcLens
}

type SrcIntels struct {
	InfoTips []InfoTip  `json:"tips,omitempty"`
	Refs     []*SrcLens `json:"refs,omitempty"`
}

type srcIntelResp struct {
	SrcIntels
	Signature  *SrcIntelSigHelp `json:"sig,omitempty"`
	Cmpl       []SrcIntelCompl  `json:"cmpl,omitempty"`
	Highlights []SrcRange       `json:"high,omitempty"`
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
	Impl ISrcIntel
}

func (*SrcIntelBase) Init() {
}

func (me *SrcIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	resp.SrcIntel = &srcIntelResp{}
	switch req.IpcID {
	case IPCID_SRCINTEL_HOVER:
		me.onHover(req, resp)
	case IPCID_SRCINTEL_SYMS_FILE, IPCID_SRCINTEL_SYMS_PROJ:
		me.onSyms(req, resp)
	case IPCID_SRCINTEL_CMPL_ITEMS:
		me.onCmplItems(req, resp)
	case IPCID_SRCINTEL_CMPL_DETAILS:
		me.onCmplDetails(req, resp)
	case IPCID_SRCINTEL_HIGHLIGHTS:
		me.onHighlights(req, resp)
	case IPCID_SRCINTEL_SIGNATURE:
		me.onSignature(req, resp)
	case IPCID_SRCINTEL_REFERENCES:
		me.onReferences(req, resp)
	case IPCID_SRCINTEL_DEFIMPL:
		me.onDefinition(req, resp, me.Impl.DefImpl)
	case IPCID_SRCINTEL_DEFSYM:
		me.onDefinition(req, resp, me.Impl.DefSym)
	case IPCID_SRCINTEL_DEFTYPE:
		me.onDefinition(req, resp, me.Impl.DefType)
	default:
		resp.SrcIntel = nil
		return false
	}
	return true
}

func (me *SrcIntelBase) onCmplItems(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.Cmpl = me.Impl.ComplItems(req.SrcLens)
}

func (me *SrcIntelBase) onCmplDetails(req *ipcReq, resp *ipcResp) {
	itemtext, _ := req.IpcArgs.(string)
	resp.SrcIntel.Cmpl = make([]SrcIntelCompl, 1)
	me.Impl.ComplDetails(req.SrcLens, itemtext, &(resp.SrcIntel.Cmpl[0]))
}

func (me *SrcIntelBase) onDefinition(req *ipcReq, resp *ipcResp, def func(*SrcLens) []*SrcLens) {
	resp.SrcIntel.Refs = def(req.SrcLens)
}

func (me *SrcIntelBase) onHighlights(req *ipcReq, resp *ipcResp) {
	curword, _ := req.IpcArgs.(string)
	resp.SrcIntel.Highlights = me.Impl.Highlights(req.SrcLens, curword)
}

func (me *SrcIntelBase) onHover(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.InfoTips = me.Impl.Hovers(req.SrcLens)
}

func (me *SrcIntelBase) onReferences(req *ipcReq, resp *ipcResp) {
	includeDeclaration := false
	if ctx, _ := req.IpcArgs.(map[string]interface{}); ctx != nil {
		if incldecl, ok := ctx["includeDeclaration"]; ok {
			includeDeclaration, _ = incldecl.(bool)
		}
	}
	resp.SrcIntel.Refs = me.Impl.References(req.SrcLens, includeDeclaration)
}

func (me *SrcIntelBase) onSignature(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.Signature = me.Impl.Signature(req.SrcLens)
}

func (me *SrcIntelBase) onSyms(req *ipcReq, resp *ipcResp) {
	var query string
	if req.IpcID == IPCID_SRCINTEL_SYMS_PROJ {
		query, _ = req.IpcArgs.(string)
	}
	resp.SrcIntel.Refs = me.Impl.Symbols(req.SrcLens, query, req.IpcID == IPCID_SRCINTEL_SYMS_FILE)
}
