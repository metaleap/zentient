package z

type ExtrasKind uint8

const (
	_ ExtrasKind = iota
	EXTRAS_INTEL
	EXTRAS_QUERY
)

type iExtras interface {
	iDispatcher

	ListIntelExtras() []ExtrasItem
	ListQueryExtras() []ExtrasItem
}

type ExtrasItem struct {
	ID          string     `json:"id"`
	Kind        ExtrasKind `json:"k"`
	Label       string     `json:"label"`
	Description string     `json:"description"`
	Detail      string     `json:"detail,omitempty"`
}

type ExtrasBase struct {
	Impl iExtras
}

func (*ExtrasBase) Init() {
}

func (me *ExtrasBase) dispatch(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_EXTRAS_QUERY_LIST:
		me.onList(req, resp, EXTRAS_QUERY)
	case MSGID_EXTRAS_INTEL_LIST:
		me.onList(req, resp, EXTRAS_INTEL)
	case MSGID_EXTRAS_INVOKE:
		me.onInvoke(req, resp)
	default:
		return false
	}
	return true
}

func (me *ExtrasBase) onList(req *msgReq, resp *msgResp, kind ExtrasKind) {
	switch kind {
	case EXTRAS_INTEL:
		resp.Extras = me.Impl.ListIntelExtras()
	case EXTRAS_QUERY:
		resp.Extras = me.Impl.ListQueryExtras()
	default:
		Bad("ExtrasKind", Strf("%v", kind))
	}
}

func (me *ExtrasBase) onInvoke(req *msgReq, resp *msgResp) {
	id, _ := req.MsgArgs.(string)
	resp.CoreCmd = &coreCmdResp{NoteInfo: "Invoked: " + id}
}
