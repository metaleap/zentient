package z

import (
	"strings"
)

type ISortable interface {
	IsSortedPriorTo(interface{}) bool
}

type IListItem interface {
	ISortable
}

type ListItemPredicate func(IListItem) bool

type ListItems []IListItem

func (me ListItems) Len() int               { return len(me) }
func (me ListItems) Less(i int, j int) bool { return me[i].IsSortedPriorTo(me[j]) }
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

	ipcID(*ListFilter) IpcIDs
	listItemsSubMenu(string, string, ListFilters) *Menu
	ListItemToMenuItem(IListItem) *MenuItem
}

type ListFilters map[*ListFilter]bool

type ListFilter struct {
	ID        string
	Disabled  bool
	Title     string
	Desc      string
	OnSrcLens func(*ListFilter, *SrcLens)
	Pred      ListItemPredicate
}

type ListBase struct {
	impl IList

	listFilters []*ListFilter
}

func (me *ListBase) init(impl IList) {
	me.impl = impl
	me.listFilters = []*ListFilter{
		{Title: "All", Desc: me.impl.UnfilteredDesc()},
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
	fdesc string
	items MenuItems
}

func (me *ListMenuBase) init(impl IListMenu, cat string, fdesc string) {
	me.fdesc = fdesc
	me.ListBase.init(impl)
	me.cat, me.impl = cat, impl

	for _, lf := range me.listFilters {
		item := &MenuItem{Title: lf.Title, Desc: me.itemDesc(nil, lf)}
		item.IpcID, item.IpcArgs = me.impl.ipcID(lf), lf.ID
		me.items = append(me.items, item)
	}
}

func (me *ListMenuBase) itemDesc(srcLens *SrcLens, lf *ListFilter) string {
	if lf.OnSrcLens != nil {
		lf.OnSrcLens(lf, srcLens)
	}
	return Strf(me.fdesc, Lang.Title, lf.Desc)
}

func (me *ListMenuBase) listItemsSubMenu(title string, desc string, filters ListFilters) *Menu {
	listitems := me.impl.List(filters)
	cat := me.cat
	if len(listitems) == 1 && strings.HasSuffix(cat, "s") {
		cat = cat[:len(cat)-1]
	}
	menu := &Menu{Desc: Strf("%d %s: %s (%s)", len(listitems), cat, title, desc)}
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

func (me *ListMenuBase) menuItems(srcLens *SrcLens) MenuItems {
	const fhint = "(%v at last count)"
	for _, item := range me.items {
		fcount, filterid := "amount unknown", item.IpcArgs.(string)
		var filters ListFilters
		if filterid != "" {
			lf := me.impl.FilterByID(filterid)
			if lf.OnSrcLens != nil {
				item.Desc = me.itemDesc(srcLens, lf)
			}
			filters = ListFilters{lf: true}
		}
		count := me.impl.Count(filters)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return me.items
}
