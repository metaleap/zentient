package z

type IPkgIntel interface {
	IMenuProvider

	ListAllDesc() string
	ListCount([]*PkgIntelFilter, map[string]bool) int
	ListFilters() []*PkgIntelFilter
}

type PkgIntelFilter struct {
	ID    string
	Title string `json:"-"`
	Desc  string `json:"-"`
}

type PkgIntelBase struct {
	Impl IPkgIntel

	listFilters   []*PkgIntelFilter
	listMenuItems []*MenuItem
}

func init() {
	dummy := &PkgIntelBase{}
	Lang.PkgIntel = dummy
	dummy.Impl = dummy
}

func (me *PkgIntelBase) Init() {
	me.listFilters = []*PkgIntelFilter{
		&PkgIntelFilter{Title: "All", Desc: me.Impl.ListAllDesc()},
	}
	me.listFilters = append(me.listFilters, me.Impl.ListFilters()...)

	const fdesc = "Lists locally-known %s packages %s"
	for _, lf := range me.listFilters {
		item := &MenuItem{Title: lf.Title, Desc: Strf(fdesc, Lang.Title, lf.Desc)}
		item.IpcArgs = lf
		me.listMenuItems = append(me.listMenuItems, item)
	}
}

func (me *PkgIntelBase) MenuItems(*SrcLens) []*MenuItem {
	const fhint = "(%v at last count)"
	for _, item := range me.listMenuItems {
		fcount, lf := "amount unknown", item.IpcArgs.(*PkgIntelFilter)
		count := me.Impl.ListCount([]*PkgIntelFilter{lf}, nil)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return me.listMenuItems
}

func (me *PkgIntelBase) MenuCategory() string {
	return "Packages"
}

func (me *PkgIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	default:
		return false
	}
	return true
}

func (me *PkgIntelBase) FilterByID(id string) *PkgIntelFilter {
	for _, lf := range me.listFilters {
		if lf.ID == id {
			return lf
		}
	}
	return nil
}
