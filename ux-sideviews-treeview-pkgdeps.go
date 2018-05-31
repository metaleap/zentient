package z

import (
	"github.com/go-leap/str"
)

type TreeDataProviderPkgDeps struct {
	onChanged func(string, []string) error
}

func (me *TreeDataProviderPkgDeps) id() string { return "pkgDeps" }

func (me *TreeDataProviderPkgDeps) getTreeItem(item []string) *TreeItem {
	foo := ustr.Join(item, sideViewsTreeItemSep)
	return &TreeItem{ID: ustr.Lo(foo), Label: foo, Tooltip: ustr.Up(foo)}
}

func (me *TreeDataProviderPkgDeps) getChildren(item []string) [][]string {
	if len(item) == 0 {
		return [][]string{
			[]string{Lang.ID + "-Dep 1"},
			[]string{Lang.ID + "-Dep 2"},
			[]string{Lang.ID + "-Dep 3"},
		}
	}
	return nil
}
