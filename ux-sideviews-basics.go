package z

import (
	"github.com/go-leap/str"
)

const (
	sideViewsTreeItemSep = "/"
)

type sideViewTreeItem []string

func (me sideViewTreeItem) String() string {
	return ustr.Join(me, sideViewsTreeItemSep)
}

type TreeViewItem struct {
	ID               string        `json:"id,omitempty"`
	Label            string        `json:"label,omitempty"`
	IconPath         string        `json:"iconPath,omitempty"`
	Tooltip          string        `json:"tooltip,omitempty"`
	Command          *EditorAction `json:"command,omitempty"`
	ContextValue     string        `json:"contextValue,omitempty"`
	CollapsibleState int           `json:"collapsibleState"`
}

type iTreeDataProvider interface {
	getTreeViewItem(sideViewTreeItem) *TreeViewItem
	getChildren(sideViewTreeItem) []sideViewTreeItem
	id() string
}

type sideViews struct {
	treeDataProviders       []iTreeDataProvider
	treeDataProviderPkgDeps treeDataProviderPkgIntel
	treeDataProviderPkgSyms treeDataProviderPkgIntel
}

func (me *sideViews) Init() {
	me.treeDataProviderPkgDeps.onChanged, me.treeDataProviderPkgSyms.onChanged = me.sendOnChanged, me.sendOnChanged
	me.treeDataProviderPkgDeps.treeViewId, me.treeDataProviderPkgSyms.treeViewId = "pkgDeps", "pkgSyms"
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
			resp.Val = dataprovider.getTreeViewItem(treepathparts)
		case reqchildren:
			childitems := dataprovider.getChildren(treepathparts)
			items := make([]string, len(childitems))
			for i, item := range childitems {
				items[i] = item.String()
			}
			resp.Val = items
		}
		resp.IpcID = req.IpcID
		return true
	}
	return false
}

func (me *sideViews) sendOnChanged(treeViewId string, item sideViewTreeItem) error {
	return send(&ipcResp{IpcID: IPCID_TREEVIEW_CHANGED, Val: []string{treeViewId, item.String()}})
}

type treeDataProviderPkgIntel struct {
	treeViewId string
	onChanged  func(string, sideViewTreeItem) error
}

func (me *treeDataProviderPkgIntel) id() string { return me.treeViewId }

func (me *treeDataProviderPkgIntel) getTreeViewItem(item sideViewTreeItem) *TreeViewItem {
	println(len(item))
	foo := item.String()
	if len(item) == 1 && item[0] == "?" {
		foo = "(" + Prog.Name + " does not support the PkgIntel interface)"
	}
	return &TreeViewItem{ID: ustr.Lo(foo), Label: foo, Tooltip: ustr.Up(foo)}
}

func (me *treeDataProviderPkgIntel) getChildren(item sideViewTreeItem) []sideViewTreeItem {
	if len(item) == 0 {
		if Lang.PkgIntel == nil {
			return []sideViewTreeItem{
				sideViewTreeItem{"?"},
			}
		} else {
			return []sideViewTreeItem{
				sideViewTreeItem{Lang.ID + "-Dep 1"},
				sideViewTreeItem{Lang.ID + "-Dep 2"},
				sideViewTreeItem{Lang.ID + "-Dep 3"},
			}
		}
	}
	return nil
}
