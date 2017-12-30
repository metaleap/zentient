package z

type IPkgIntel interface {
	IListMenu
	IObjSnap
}

type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel
}

func (me *PkgIntelBase) Init() {
	me.ListMenuBase.init(me.Impl, "Packages", "Lists %s packages %s")
}

func (me *PkgIntelBase) IpcID(_ *ListFilter) IpcIDs {
	return IPCID_MENUS_PKGS
}

func (me *PkgIntelBase) ObjSnapPrefix() string {
	return Lang.ID + ".pkgIntel."
}

func (me *PkgIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_PKGS:
		filterid, _ := req.IpcArgs.(string)
		filter := me.Impl.FilterByID(filterid)
		var filters ListFilters
		if filterid != "" {
			filters = ListFilters{filter: true}
		}
		resp.Menu = &MenuResp{
			SubMenu: me.Impl.ListItemsSubMenu(filter.Title, filter.Desc, filters),
		}
	default:
		return false
	}
	return true
}
