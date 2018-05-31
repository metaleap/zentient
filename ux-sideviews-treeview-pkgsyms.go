package z

import (
	"github.com/go-leap/str"
)

type treeDataProviderPkgSyms struct {
	onChanged func(string, []string) error
}

func (me *treeDataProviderPkgSyms) id() string { return "pkgSyms" }

func (me *treeDataProviderPkgSyms) getTreeItem(item []string) *TreeItem {
	foo := ustr.Join(item, sideViewsTreeItemSep)
	if len(item) == 1 && item[0] == "?" {
		foo = "(" + Prog.Name + " does not support the PkgIntel interface)"
	}
	return &TreeItem{ID: ustr.Lo(foo), Label: foo, Tooltip: ustr.Up(foo)}
}

func (me *treeDataProviderPkgSyms) getChildren(item []string) [][]string {
	if len(item) == 0 {
		if Lang.PkgIntel != nil {
			return [][]string{
				[]string{Lang.ID + "-Sym 1"},
				[]string{Lang.ID + "-Sym 2"},
				[]string{Lang.ID + "-Sym 3"},
			}
		}
	}
	return [][]string{[]string{"?"}}
}
