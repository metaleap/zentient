package z

import (
	"sort"
)

type ISrcIntel interface {
	iDispatcher

	ComplDetails(*SrcLens, string) *SrcIntelCompl
	ComplItems(*SrcLens) SrcIntelCompls
	DefSym(*SrcLens) SrcLocs
	DefType(*SrcLens) SrcLocs
	DefImpl(*SrcLens) SrcLocs
	Highlights(*SrcLens, string) SrcLocs
	Hovers(*SrcLens) []InfoTip
	References(*SrcLens, bool) SrcLocs
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) SrcLenses
}

type SrcIntels struct {
	Info []InfoTip `json:",omitempty"`
	Refs SrcLocs   `json:",omitempty"`
}

type srcIntelResp struct {
	SrcIntels
	Sig  *SrcIntelSigHelp `json:",omitempty"`
	Cmpl SrcIntelCompls   `json:",omitempty"`
	Syms SrcLenses        `json:",omitempty"`
}

type SrcIntelCompl struct {
	Kind          Completion   `json:"kind,omitempty"`
	Label         string       `json:"label"`
	Documentation *SrcIntelDoc `json:"documentation,omitempty"`
	Detail        string       `json:"detail,omitempty"`
	SortText      string       `json:"sortText,omitempty"`
	// FilterText    string       `json:"filterText,omitempty"`
	// InsertText    string       `json:"insertText,omitempty"`
	// CommitChars   []string     `json:"commitCharacters,omitempty"` // basically in all languages always operator/separator/punctuation (that is, "non-identifier") chars --- no need to send them for each item, for each language --- the client-side will do it
	SortPrio int `json:"-"`
}

type SrcIntelCompls []*SrcIntelCompl

func (me SrcIntelCompls) Len() int               { return len(me) }
func (me SrcIntelCompls) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me SrcIntelCompls) Less(i int, j int) bool { return me[i].SortPrio < me[j].SortPrio }

type SrcIntelDoc struct {
	Value     string `json:"value,omitempty"`
	IsTrusted bool   `json:"isTrusted,omitempty"`
}

type SrcIntelSigHelp struct {
	ActiveSignature int               `json:"activeSignature"`
	ActiveParameter int               `json:"activeParameter,omitempty"`
	Signatures      []SrcIntelSigInfo `json:"signatures,omitempty"`
}

type SrcIntelSigInfo struct {
	Label         string             `json:"label"`
	Documentation SrcIntelDoc        `json:"documentation,omitempty"`
	Parameters    []SrcIntelSigParam `json:"parameters"`
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
	switch req.IpcID {
	case IPCID_SRCINTEL_HOVER:
		me.onHover(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_SYMS_FILE, IPCID_SRCINTEL_SYMS_PROJ:
		me.onSyms(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_CMPL_ITEMS:
		me.onCmplItems(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_CMPL_DETAILS:
		me.onCmplDetails(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_HIGHLIGHTS:
		me.onHighlights(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_SIGNATURE:
		me.onSignature(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_REFERENCES:
		me.onReferences(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_DEFIMPL:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefImpl)
	case IPCID_SRCINTEL_DEFSYM:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefSym)
	case IPCID_SRCINTEL_DEFTYPE:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefType)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) onCmplItems(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.Cmpl = me.Impl.ComplItems(req.SrcLens)
	var shouldsort bool
	for _, c := range resp.SrcIntel.Cmpl {
		if shouldsort = c.SortPrio != 0; shouldsort {
			break
		}
	}
	if shouldsort {
		sort.Sort(resp.SrcIntel.Cmpl)
		for i, c := range resp.SrcIntel.Cmpl {
			c.SortText = Strf("%03d", i)
		}
	}
}

func (me *SrcIntelBase) onCmplDetails(req *ipcReq, resp *ipcResp) {
	itemtext, _ := req.IpcArgs.(string)
	if cmpl := me.Impl.ComplDetails(req.SrcLens, itemtext); cmpl != nil {
		resp.SrcIntel.Cmpl = SrcIntelCompls{cmpl}
	}
}

func (*SrcIntelBase) onDefinition(req *ipcReq, resp *ipcResp, def func(*SrcLens) SrcLocs) {
	resp.SrcIntel.Refs = def(req.SrcLens)
}

func (me *SrcIntelBase) onHighlights(req *ipcReq, resp *ipcResp) {
	curword, _ := req.IpcArgs.(string)
	resp.SrcIntel.Refs = me.Impl.Highlights(req.SrcLens, curword)
}

func (me *SrcIntelBase) onHover(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.Info = me.Impl.Hovers(req.SrcLens)
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
	if resp.SrcIntel.Sig = me.Impl.Signature(req.SrcLens); resp.SrcIntel.Sig != nil {
		for i := range resp.SrcIntel.Sig.Signatures { // vsc can't handle `null` for `parameters` but can handle `[]`
			if resp.SrcIntel.Sig.Signatures[i].Documentation.IsTrusted = true; resp.SrcIntel.Sig.Signatures[i].Parameters == nil {
				resp.SrcIntel.Sig.Signatures[i].Parameters = []SrcIntelSigParam{}
			}
		}
	}
}

func (me *SrcIntelBase) onSyms(req *ipcReq, resp *ipcResp) {
	var query string
	if req.IpcID == IPCID_SRCINTEL_SYMS_PROJ {
		query, _ = req.IpcArgs.(string)
	}
	resp.SrcIntel.Syms = me.Impl.Symbols(req.SrcLens, query, req.IpcID == IPCID_SRCINTEL_SYMS_FILE)
}

func (*SrcIntelBase) ComplItems(srcLens *SrcLens) SrcIntelCompls {
	return nil
}

func (*SrcIntelBase) ComplDetails(srcLens *SrcLens, itemText string) *SrcIntelCompl {
	return nil
}

func (*SrcIntelBase) DefImpl(srcLens *SrcLens) SrcLenses {
	return nil
}

func (*SrcIntelBase) DefSym(srcLens *SrcLens) SrcLenses {
	return nil
}

func (*SrcIntelBase) DefType(srcLens *SrcLens) SrcLenses {
	return nil
}

func (*SrcIntelBase) Highlights(srcLens *SrcLens, curWord string) SrcLenses {
	return nil
}

func (*SrcIntelBase) Hovers(srcLens *SrcLens) []InfoTip {
	return nil
}

func (*SrcIntelBase) References(srcLens *SrcLens, includeDeclaration bool) SrcLenses {
	return nil
}

func (*SrcIntelBase) Signature(srcLens *SrcLens) *SrcIntelSigHelp {
	return nil
}

func (*SrcIntelBase) Symbols(srcLens *SrcLens, query string, curFileOnly bool) SrcLenses {
	return nil
}
