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
	getTreeItem([]string) *TreeItem
	getChildren([]string) []string
	id() string
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
		treepathparts := strings.Split(treepath, ":")
		if len(treepathparts) == 0 {
			BadPanic(IPCID_TREEVIEW_GETITEM.String()+" arg", "")
		} else {
			for _, dp := range me.treeDataProviders {
				if dp.id() == treepathparts[0] {
					dataprovider = dp
					break
				}
			}
			if dataprovider == nil {
				BadPanic("tree-data provider ID", treepathparts[0])
			}
		}
		switch req.IpcID {
		case IPCID_TREEVIEW_GETITEM:
			resp.Val = dataprovider.getTreeItem(treepathparts)
		case IPCID_TREEVIEW_CHILDREN:
			resp.Val = dataprovider.getChildren(treepathparts)
		}
		return true
	}
	return false
}
