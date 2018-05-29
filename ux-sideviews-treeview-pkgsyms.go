package z

type TreeDataProviderPkgSyms struct {
}

func (me *TreeDataProviderPkgSyms) id() string { return "pkgSyms" }

func (me *TreeDataProviderPkgSyms) getTreeItem([]string) *TreeItem {
	return nil
}

func (me *TreeDataProviderPkgSyms) getChildren([]string) []string {
	return nil
}
