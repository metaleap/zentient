package z

import (
	"github.com/go-leap/str"
)

type TreeDataProviderPkgSyms struct {
	onChanged func(string, []string) error
}

func (me *TreeDataProviderPkgSyms) id() string { return "pkgSyms" }

func (me *TreeDataProviderPkgSyms) getTreeItem(item []string) *TreeItem {
	foo := ustr.Join(item, sideViewsTreeItemSep)
	return &TreeItem{ID: ustr.Lo(foo), Label: foo, Tooltip: ustr.Up(foo)}
}

func (me *TreeDataProviderPkgSyms) getChildren(item []string) [][]string {
	if len(item) == 0 {
		return [][]string{
			[]string{Lang.ID + "-Sym 1"},
			[]string{Lang.ID + "-Sym 2"},
			[]string{Lang.ID + "-Sym 3"},
		}
	}
	return nil
}
