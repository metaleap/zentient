package z

import (
	"github.com/go-leap/str"
)

type treeDataProviderPkgDeps struct {
	onChanged func(string, []string) error
}

func (me *treeDataProviderPkgDeps) id() string { return "pkgDeps" }

func (me *treeDataProviderPkgDeps) getTreeItem(item []string) *TreeItem {
	foo := ustr.Join(item, sideViewsTreeItemSep)
	if len(item) == 1 && item[0] == "?" {
		foo = "(" + Prog.Name + " does not support the PkgIntel interface)"
	}
	return &TreeItem{ID: ustr.Lo(foo), Label: foo, Tooltip: ustr.Up(foo)}
}

func (me *treeDataProviderPkgDeps) getChildren(item []string) [][]string {
	if len(item) == 0 {
		if Lang.PkgIntel != nil {
			return [][]string{
				[]string{Lang.ID + "-Dep 1"},
				[]string{Lang.ID + "-Dep 2"},
				[]string{Lang.ID + "-Dep 3"},
			}
		}
	}
	return [][]string{[]string{"?"}}
}
