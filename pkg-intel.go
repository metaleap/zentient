package z

import (
	"strings"
)

type IPkgIntel interface {
	IListMenu

	ObjSnap(string) ListItem
}

type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel
}

func (me *PkgIntelBase) Init() {
	me.ListMenuBase.init(me.Impl, Lang.Title+" "+"Packages", "Lists %s packages %s")
}

func (me *PkgIntelBase) IpcID(_ *ListFilter) ipcIDs {
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
		return true
	case IPCID_OBJ_SNAPSHOT:
		pref := me.ObjSnapPrefix()
		if path, _ := req.IpcArgs.(string); strings.HasPrefix(path, pref) {
			resp.Menu = &MenuResp{ObjSnapshot: me.Impl.ObjSnap(path[len(pref):])}
			return true
		}
	}
	return false
}
