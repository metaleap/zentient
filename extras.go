package z

type IExtras interface {
	iDispatcher

	ListIntelExtras() []ExtrasItem
	ListQueryExtras() []ExtrasItem
	RunIntelExtra(*SrcLens, string, string, *ExtrasResp)
	RunQueryExtra(*SrcLens, string, string, *ExtrasResp)
}

type ExtrasItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Desc     string `json:"description"`
	Detail   string `json:"detail,omitempty"`
	QueryArg string `json:"arg,omitempty"`
}

type ExtrasResp struct {
	SrcIntels
	Items []ExtrasItem `json:"items,omitempty"`
	Warns []string     `json:"warns,omitempty"`
	Desc  string       `json:"desc,omitempty"`
}

type ExtrasBase struct {
	Impl IExtras
}

func (*ExtrasBase) Init() {
}

func (me *ExtrasBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	resp.Extras = &ExtrasResp{}
	switch req.IpcID {
	case IPCID_EXTRAS_INTEL_LIST:
		me.onList(req, resp, false)
	case IPCID_EXTRAS_QUERY_LIST:
		me.onList(req, resp, true)
	case IPCID_EXTRAS_INTEL_RUN:
		me.onRun(req, resp, false)
	case IPCID_EXTRAS_QUERY_RUN:
		me.onRun(req, resp, true)
	default:
		resp.Extras = nil
		return false
	}
	resp.IpcID = req.IpcID
	return true
}

func (me *ExtrasBase) onList(req *ipcReq, resp *ipcResp, isQuery bool) {
	list := me.Impl.ListIntelExtras
	if isQuery {
		list = me.Impl.ListQueryExtras
	}
	resp.Extras.Items = list()
	for i := range resp.Extras.Items {
		if item := &resp.Extras.Items[i]; item.Desc == "" {
			if req.SrcLens.Str != "" {
				item.Desc = req.SrcLens.Str
			} else {
				item.Desc = Strf("at %s in: ", req.SrcLens.Pos)
				if req.SrcLens.Txt != "" {
					item.Desc += req.SrcLens.Txt
				} else {
					item.Desc += Lang.Workspace.PrettyPath(req.SrcLens.FilePath)
				}
			}
		}
	}
}

func (me *ExtrasBase) onRun(req *ipcReq, resp *ipcResp, isQuery bool) {
	ipcargs := req.IpcArgs.([]interface{})
	id, _ := ipcargs[0].(string)
	arg, _ := ipcargs[1].(string)
	run := me.Impl.RunIntelExtra
	if isQuery {
		run = me.Impl.RunQueryExtra
	}
	run(req.SrcLens, id, arg, resp.Extras)
}
