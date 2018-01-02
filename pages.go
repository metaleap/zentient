package z

type IPages interface {
	iDispatcher

	PageBodyInnerHtml(string) string
}

type PagesBase struct {
	Impl IPages
}

func (*PagesBase) Init() {}

func (me *PagesBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_PAGE_HTML:
		resp.IpcID = req.IpcID
		resp.ObjSnapshot = me.Impl.PageBodyInnerHtml(req.IpcArgs.(string))
	default:
		return false
	}
	return true
}

func (*PagesBase) PageBodyInnerHtml(uriPath string) string {
	return Strf("<h1>Not Yet Implemented</h1><p>The Zentient %s provider <code>%s</code> has not implemented a custom <code>IPages.PageBodyInnerHtml(string) string</code> handler to serve this request with path: <code>%s</code></p>", Lang.Title, Prog.name, uriPath)
}
