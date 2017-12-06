package z

type ExtrasKind uint8

const (
	_ ExtrasKind = iota
	EXTRAS_INTEL
	EXTRAS_QUERY
)

type iExtras interface {
	iDispatcher
}

type extrasResp struct {
	Items []ExtrasItem
}

type ExtrasItem struct {
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

func (me *ExtrasBase) onList(req *msgReq, resp *msgResp, kind ExtrasKind) {
	resp.Extras.Items = append(resp.Extras.Items, ExtrasItem{Label: "moo label", Description: "moo desc", Detail: "moo detail"})
}
