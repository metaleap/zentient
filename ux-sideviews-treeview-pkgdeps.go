package z

type TreeDataProviderPkgDeps struct {
}

func (me *TreeDataProviderPkgDeps) id() string { return "pkgDeps" }

func (me *TreeDataProviderPkgDeps) getTreeItem([]string) *TreeItem {
	return nil
}

func (me *TreeDataProviderPkgDeps) getChildren([]string) []string {
	return nil
}
