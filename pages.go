package z

import (
	"net/url"
	"strings"
)

type IPages interface {
	iDispatcher

	PageBodyInnerHtml(string, []string, url.Values, string) string
}

type PagesBase struct {
	Impl IPages
}

func (*PagesBase) Init() {}

func (me *PagesBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_PAGE_HTML:
		resp.IpcID = req.IpcID
		resp.ObjSnapshot = me.onPageHtml(req.IpcArgs.(string))
	default:
		return false
	}
	return true
}

func (me *PagesBase) onPageHtml(rawUri string) string {
	uri, err := url.Parse(rawUri)
	if err != nil {
		return err.Error()
	}
	return me.Impl.PageBodyInnerHtml(rawUri, strings.Split(strings.Trim(uri.Path, "/"), "/"), uri.Query(), uri.Fragment)

}

func (*PagesBase) PageBodyInnerHtml(rawUri string, path []string, query url.Values, fragment string) string {
	return Strf("<h1>Not Yet Implemented</h1><p>The Zentient %s provider <code>%s</code> has not implemented a custom <code>IPages.PageBodyInnerHtml(string) string</code> handler to serve this request with path: <code>%s</code></p>", Lang.Title, Prog.name, rawUri)
}
