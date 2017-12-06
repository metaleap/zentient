package z

type extrasKind uint8

const (
	_ extrasKind = iota
	EXTRAS_INTEL
	EXTRAS_QUERY
)

type iExtras interface {
	iDispatcher

	ListIntelExtras() []ExtrasItem
	ListQueryExtras() []ExtrasItem
}

type extrasResp struct {
	Items []ExtrasItem
}

type ExtrasItem struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Detail      string `json:"detail,omitempty"`
}

type ExtrasBase struct {
	Impl iExtras
}

func (*ExtrasBase) Init() {
}

func (me *ExtrasBase) dispatch(req *msgReq, resp *msgResp) bool {
	resp.Extras = &extrasResp{}
	switch req.MsgID {
	case MSGID_EXTRAS_QUERY_LIST:
		me.onList(req, resp, EXTRAS_QUERY)
	case MSGID_EXTRAS_INTEL_LIST:
		me.onList(req, resp, EXTRAS_INTEL)
	default:
		resp.Extras = nil
		return false
	}
	return true
}

func (me *ExtrasBase) onList(req *msgReq, resp *msgResp, kind extrasKind) {
	switch kind {
	case EXTRAS_INTEL:
		resp.Extras.Items = me.Impl.ListIntelExtras()
	case EXTRAS_QUERY:
		resp.Extras.Items = me.Impl.ListQueryExtras()
	default:
		Bad("extrasKind", Strf("%v", kind))
	}
}
