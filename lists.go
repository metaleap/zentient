package z

type IList interface {
	DescUnfiltered() string
	Count([]*ListFilter, map[string]bool) int
	Filters() []*ListFilter
	// List([]*ListFilter, map[string]bool) int
}

type IListMenu interface {
	IList
	IMenuItems
}

type ListFilter struct {
	ID    string
	Title string `json:"-"`
	Desc  string `json:"-"`
}

type ListBase struct {
	impl IList

	listFilters []*ListFilter
}

func (me *ListBase) init(impl IList) {
	me.impl = impl
	me.listFilters = []*ListFilter{
		&ListFilter{Title: "All", Desc: me.impl.DescUnfiltered()},
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
		item.IpcArgs = lf
		me.items = append(me.items, item)
	}
}

func (me *ListMenuBase) MenuCategory() string {
	return me.cat
}

func (me *ListMenuBase) MenuItems(*SrcLens) []*MenuItem {
	const fhint = "(%v at last count)"
	for _, item := range me.items {
		fcount, lf := "amount unknown", item.IpcArgs.(*ListFilter)
		count := me.impl.Count([]*ListFilter{lf}, nil)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return me.items
}
