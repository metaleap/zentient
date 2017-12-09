package z

type ListItem interface {
	LessThan(interface{}) bool
}

type ListItems []ListItem

func (me ListItems) Len() int               { return len(me) }
func (me ListItems) Less(i int, j int) bool { return me[i].LessThan(me[j]) }
func (me ListItems) Swap(i, j int)          { me[i], me[j] = me[j], me[i] }

type IList interface {
	UnfilteredDesc() string
	Count(ListFilters) int
	FilterByID(string) *ListFilter
	Filters() []*ListFilter
	List(ListFilters) ListItems
}

type IListMenu interface {
	IList
	IMenuItems

	IpcID(*ListFilter) ipcIDs
	ListItemsSubMenu(string, string, ListFilters) *Menu
	ListItemToMenuItem(ListItem) *MenuItem
}

type ListFilters map[*ListFilter]bool

type ListFilter struct {
	ID    string
	Title string `json:"-"`
	Desc  string `json:"-"`

	Pred func(ListItem) bool `json:"-"`
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

func (me *ListBase) Count(all ListFilters) int {
	return len(me.impl.List(all))
}

func (me *ListBase) FilterByID(id string) *ListFilter {
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
		item.IpcID, item.IpcArgs = me.impl.IpcID(lf), lf.ID
		me.items = append(me.items, item)
	}
}

func (me *ListMenuBase) ListItemsSubMenu(title string, desc string, filters ListFilters) *Menu {
	listitems := me.impl.List(filters)
	menu := &Menu{Desc: Strf("%d %s ➜ %s (%s)", len(listitems), me.cat, title, desc)}
	for _, listitem := range listitems {
		if menuitem := me.impl.ListItemToMenuItem(listitem); menuitem != nil {
			menu.Items = append(menu.Items, menuitem)
		}
	}
	return menu
}

func (me *ListMenuBase) MenuCategory() string {
	return me.cat
}

func (me *ListMenuBase) MenuItems(*SrcLens) []*MenuItem {
	const fhint = "(%v at last count)"
	for _, item := range me.items {
		fcount, filterid := "amount unknown", item.IpcArgs.(string)
		var filters ListFilters
		if filterid != "" {
			filters = ListFilters{me.impl.FilterByID(filterid): true}
		}
		count := me.impl.Count(filters)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return me.items
}
