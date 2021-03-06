package z

type IPkgIntel interface {
	IListMenu
	IObjSnap
	Pkgs() PkgInfos
}

type PkgInfos []*PkgInfo

func (me *PkgInfos) Add(pkg *PkgInfo) {
	*me = append(*me, pkg)
}

func (me PkgInfos) ById(id string) *PkgInfo {
	for _, pkg := range me {
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

func (me *PkgIntelBase) Init() {
	me.ListMenuBase.init(me.Impl, "Packages", "Lists %s packages %s")
}

func (me *PkgIntelBase) ipcID(_ *ListFilter) IpcIDs {
	return IPCID_MENUS_PKGS
}

func (me *PkgIntelBase) ObjSnapPrefix() string {
	return Lang.ID + ".pkgIntel."
}

func (me *PkgIntelBase) dispatch(req *IpcReq, resp *IpcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_PKGS:
		filterid, _ := req.IpcArgs.(string)
		filter := me.Impl.FilterByID(filterid)
		var filters ListFilters
		if filterid != "" {
			filters = ListFilters{filter: true}
		}
		resp.Menu = &MenuResponse{
			SubMenu: me.Impl.listItemsSubMenu(filter.Title, filter.Desc, filters),
		}
	default:
		return false
	}
	return true
}

func (me *PkgIntelBase) PkgsAdd(pkg *PkgInfo) {
	me.pkgs.Add(pkg)
	// Lang.sideViews.sendOnChanged("", sideViewTreeItem{pkg.Id})
}

func (me *PkgIntelBase) Pkgs() PkgInfos {
	return me.pkgs
}
