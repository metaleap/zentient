package z

type IPkgIntel interface {
	IListMenu
	IObjSnap
	Pkgs() PkgInfos
}

type PkgInfos []*PkgInfo

func (this *PkgInfos) Add(pkg *PkgInfo) {
	*this = append(*this, pkg)
}

func (this PkgInfos) ById(id string) *PkgInfo {
	for _, pkg := range this {
		if pkg.Id == id {
			return pkg
		}
	}
	return nil
}

type PkgInfo struct {
	Id        string
	ShortName string
	LongName  string
	Deps      func() PkgInfos
	Mems      func() []*PkgMemInfo
	Forget    func()
}

type PkgMemInfo struct {
	Kind Symbol
	Name string
	Desc string
	Subs func() []*PkgMemInfo
}

type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel

	pkgs PkgInfos
}

func (this *PkgIntelBase) Init() {
	this.ListMenuBase.init(this.Impl, "Packages", "Lists %s packages %s")
}

func (this *PkgIntelBase) ipcID(_ *ListFilter) IpcIDs {
	return IPCID_MENUS_PKGS
}

func (this *PkgIntelBase) ObjSnapPrefix() string {
	return Lang.ID + ".pkgIntel."
}

func (this *PkgIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_PKGS:
		filterid, _ := req.IpcArgs.(string)
		filter := this.Impl.FilterByID(filterid)
		var filters ListFilters
		if filterid != "" {
			filters = ListFilters{filter: true}
		}
		resp.Menu = &menuResp{
			SubMenu: this.Impl.listItemsSubMenu(filter.Title, filter.Desc, filters),
		}
	default:
		return false
	}
	return true
}

func (this *PkgIntelBase) PkgsAdd(pkg *PkgInfo) {
	this.pkgs.Add(pkg)
	// Lang.sideViews.sendOnChanged("", sideViewTreeItem{pkg.Id})
}

func (this *PkgIntelBase) Pkgs() PkgInfos {
	return this.pkgs
}
