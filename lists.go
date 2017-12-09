package z

type IList interface {
	UnfilteredDesc() string
	Count(ListFilters) int
	Filters() []*ListFilter
	List(ListFilters) []interface{}
}

type IListMenu interface {
	IList
	IMenuItems
}

type ListFilters map[*ListFilter]bool

type ListFilter struct {
	ID    string
	Title string `json:"-"`
	Desc  string `json:"-"`

	Pred func(interface{}) bool `json:"-"`
}

type ListBase struct {
	impl IList

	listFilters []*ListFilter
}

func (me *ListBase) init(impl IList) {
	me.impl = impl
	me.listFilters = []*ListFilter{
		&ListFilter{Title: "All", Desc: me.impl.UnfilteredDesc()},
	}
	me.listFilters = append(me.listFilters, me.impl.Filters()...)
}

func (me *ListBase) filterWithID(id string) *ListFilter {
	for _, lf := range me.listFilters {
		if lf.ID == id {
			return lf
		}
	}
	return nil
}

type ListMenuBase struct {
	ListBase
	impl IListMenu

	cat   string
	items []*MenuItem
}

func (me *ListMenuBase) init(impl IListMenu, cat string, fdesc string) {
	me.ListBase.init(impl)
	me.cat, me.impl = cat, impl

	for _, lf := range me.listFilters {
		item := &MenuItem{Title: lf.Title, Desc: Strf(fdesc, Lang.Title, lf.Desc)}
		item.IpcArgs = lf.ID
		me.items = append(me.items, item)
	}
}

func (me *ListMenuBase) MenuCategory() string {
	return me.cat
}

func (me *ListMenuBase) MenuItems(*SrcLens) []*MenuItem {
	const fhint = "(%v at last count)"
	for _, item := range me.items {
		fcount, filterid := "amount unknown", item.IpcArgs.(string)
		filters := ListFilters{me.filterWithID(filterid): true}
		if filterid == "" {
			filters = nil
		}
		count := me.impl.Count(filters)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return me.items
}
