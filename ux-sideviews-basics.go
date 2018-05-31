package z

import (
	"github.com/go-leap/str"
)

const (
	sideViewsTreeItemSep = "/"
)

type TreeItem struct {
	ID               string        `json:"id,omitempty"`
	Label            string        `json:"label,omitempty"`
	IconPath         string        `json:"iconPath,omitempty"`
	Tooltip          string        `json:"tooltip,omitempty"`
	Command          *EditorAction `json:"command,omitempty"`
	ContextValue     string        `json:"contextValue,omitempty"`
	CollapsibleState int           `json:"collapsibleState"`
}

type iTreeDataProvider interface {
	getTreeItem([]string) *TreeItem
	getChildren([]string) [][]string
	id() string
}

type sideViews struct {
	treeDataProviders       []iTreeDataProvider
	treeDataProviderPkgDeps treeDataProviderPkgDeps
	treeDataProviderPkgSyms treeDataProviderPkgSyms
}

func (me *sideViews) Init() {
	me.treeDataProviderPkgDeps.onChanged = me.sendOnChanged
	me.treeDataProviderPkgSyms.onChanged = me.sendOnChanged
	me.treeDataProviders = []iTreeDataProvider{&me.treeDataProviderPkgDeps, &me.treeDataProviderPkgSyms}
}

func (me *sideViews) dispatch(req *ipcReq, resp *ipcResp) bool {
	if reqtreeitem, reqchildren := req.IpcID == IPCID_TREEVIEW_GETITEM, req.IpcID == IPCID_TREEVIEW_CHILDREN; reqtreeitem || reqchildren {
		var dataprovider iTreeDataProvider
		ipcargs := req.IpcArgs.([]interface{})
		treeviewid := ipcargs[0].(string)
		treeitem, _ := ipcargs[1].(string)
		for _, dp := range me.treeDataProviders {
			if dp.id() == treeviewid {
				dataprovider = dp
				break
			}
		}
		if dataprovider == nil {
			BadPanic("tree-data provider ID", treeviewid)
		}
		treepathparts := ustr.Split(treeitem, sideViewsTreeItemSep)
		switch {
		case reqtreeitem:
			resp.Val = dataprovider.getTreeItem(treepathparts)
		case reqchildren:
			childitems := dataprovider.getChildren(treepathparts)
			items := make([]string, len(childitems))
			for i, item := range childitems {
				items[i] = ustr.Join(item, sideViewsTreeItemSep)
			}
			resp.Val = items
		}
		resp.IpcID = req.IpcID
		return true
	}
	return false
}

func (me *sideViews) sendOnChanged(treeViewId string, item []string) error {
	return send(&ipcResp{IpcID: IPCID_TREEVIEW_CHANGED, Val: []string{treeViewId, ustr.Join(item, sideViewsTreeItemSep)}})
}
