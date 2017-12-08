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
	RunIntelExtra(*SrcLens, string, string, *ExtrasResp)
	RunQueryExtra(*SrcLens, string, string, *ExtrasResp)
}

type ExtrasItem struct {
	ID          string     `json:"id"`
	Kind        ExtrasKind `json:"kind"`
	Label       string     `json:"label"`
	Description string     `json:"description"`
	Detail      string     `json:"detail,omitempty"`
	QueryArg    string     `json:"arg,omitempty"`
}

type ExtrasResp struct {
	SrcIntels
	Items []ExtrasItem `json:"items,omitempty"`
}

type ExtrasBase struct {
	Impl iExtras
}

func (*ExtrasBase) Init() {
}

func (me *ExtrasBase) dispatch(req *msgReq, resp *msgResp) bool {
	resp.Extras = &ExtrasResp{}
	switch req.MsgID {
	case MSGID_EXTRAS_INTEL_LIST:
		me.onList(req, resp, EXTRAS_INTEL)
	case MSGID_EXTRAS_QUERY_LIST:
		me.onList(req, resp, EXTRAS_QUERY)
	case MSGID_EXTRAS_INTEL_RUN:
		me.onRun(req, resp, EXTRAS_INTEL)
	case MSGID_EXTRAS_QUERY_RUN:
		me.onRun(req, resp, EXTRAS_QUERY)
	default:
		resp.Extras = nil
		return false
	}
	return true
}

func (me *ExtrasBase) onList(req *msgReq, resp *msgResp, kind ExtrasKind) {
	list := me.Impl.ListIntelExtras
	if kind == EXTRAS_QUERY {
		list = me.Impl.ListQueryExtras
	}
	resp.Extras.Items = list()
}

func (me *ExtrasBase) onRun(req *msgReq, resp *msgResp, kind ExtrasKind) {
	msgargs := req.MsgArgs.([]interface{})
	id, _ := msgargs[0].(string)
	arg, _ := msgargs[1].(string)
	run := me.Impl.RunIntelExtra
	if kind == EXTRAS_QUERY {
		run = me.Impl.RunQueryExtra
	}
	run(req.SrcLens, id, arg, resp.Extras)
}
