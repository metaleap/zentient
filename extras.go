package z

type IExtras interface {
	iDispatcher

	ListIntelExtras() []*ExtrasItem
	ListQueryExtras() []*ExtrasItem
	RunIntelExtra(*SrcLens, string, string, *ExtrasResp)
	RunQueryExtra(*SrcLens, string, string, *ExtrasResp)
}

type ExtrasItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Desc     string `json:"description"`
	Detail   string `json:"detail,omitempty"`
	QueryArg string `json:"arg,omitempty"`
	FilePos  string `json:"fPos,omitempty"`
}

type ExtrasResp struct {
	SrcIntels
	Items []*ExtrasItem
	Warns []string `json:",omitempty"`
	Desc  string   `json:",omitempty"`
	Url   string   `json:",omitempty"`
}

type ExtrasBase struct {
	Impl IExtras
}

func (*ExtrasBase) Init() {
}

func (this *ExtrasBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_EXTRAS_INTEL_LIST:
		this.onList(req, resp.withExtras(), false)
	case IPCID_EXTRAS_QUERY_LIST:
		this.onList(req, resp.withExtras(), true)
	case IPCID_EXTRAS_INTEL_RUN:
		this.onRun(req, resp.withExtras(), false)
	case IPCID_EXTRAS_QUERY_RUN:
		this.onRun(req, resp.withExtras(), true)
	default:
		return false
	}
	resp.IpcID = req.IpcID
	return true
}

func (this *ExtrasBase) onList(req *ipcReq, resp *ipcResp, isQuery bool) {
	list := this.Impl.ListIntelExtras
	if isQuery {
		list = this.Impl.ListQueryExtras
	}
	resp.Extras.Items = list()
	for i := range resp.Extras.Items {
		if item := resp.Extras.Items[i]; req.SrcLens.Str != "" {
			item.Desc = "`" + req.SrcLens.Str + "`"
		} else {
			item.Desc = Strf("at %s in: ", req.SrcLens.Pos)
			if req.SrcLens.Txt != "" {
				item.Desc += "`" + req.SrcLens.Txt + "`"
			} else {
				item.Desc += Lang.Workspace.PrettyPath(req.SrcLens.FilePath)
			}
		}
	}
}

func (this *ExtrasBase) onRun(req *ipcReq, resp *ipcResp, isQuery bool) {
	ipcargs := req.IpcArgs.([]interface{})
	id, _ := ipcargs[0].(string)
	arg, _ := ipcargs[1].(string)
	run := this.Impl.RunIntelExtra
	if isQuery {
		run = this.Impl.RunQueryExtra
	}
	run(req.SrcLens, id, arg, resp.Extras)
}
