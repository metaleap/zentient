package z

import (
	"strings"
)

type TreeItem struct {
	ID           string        `json:"id,omitempty"`
	Label        string        `json:"label,omitempty"`
	IconPath     string        `json:"iconPath,omitempty"`
	Tooltip      string        `json:"tooltip,omitempty"`
	Command      *EditorAction `json:"command,omitempty"`
	ContextValue string        `json:"contextValue,omitempty"`
}

type iTreeDataProvider interface {
	GetTreeItem(string) *TreeItem
	GetChildren(string) []string
}

type sideViews struct {
	treeDataProviders []iTreeDataProvider
}

func (me *sideViews) Init() {
	me.treeDataProviders = []iTreeDataProvider{&TreeDataProviderPkgSyms{}, &TreeDataProviderPkgDeps{}}
}

func (me *sideViews) dispatch(req *ipcReq, resp *ipcResp) bool {
	if req.IpcID == IPCID_TREEVIEW_GETITEM || req.IpcID == IPCID_TREEVIEW_CHILDREN {
		var dataprovider iTreeDataProvider
		treepath, _ := req.IpcArgs.(string)
		if treepathparts := strings.Split(treepath, ":"); len(treepathparts) == 0 {
			BadPanic(IPCID_TREEVIEW_GETITEM.String()+" arg", "")
		} else if dataprovider := me.treeDataProviders[treepathparts[0]]; dataprovider == nil {
			BadPanic("tree-data provider ID", treepathparts[0])
		}
		switch req.IpcID {
		case IPCID_TREEVIEW_GETITEM:
		case IPCID_TREEVIEW_CHILDREN:
		}
		return true
	}
	return false
}
