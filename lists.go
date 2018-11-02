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

func (this ListItems) Len() int               { return len(this) }
func (this ListItems) Less(i int, j int) bool { return this[i].IsSortedPriorTo(this[j]) }
func (this ListItems) Swap(i, j int)          { this[i], this[j] = this[j], this[i] }

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

func (this *ListBase) init(impl IList) {
	this.impl = impl
	this.listFilters = []*ListFilter{
		{Title: "All", Desc: this.impl.UnfilteredDesc()},
	}
	this.listFilters = append(this.listFilters, this.impl.Filters()...)
}

func (this *ListBase) Count(all ListFilters) int {
	return len(this.impl.List(all))
}

func (this *ListBase) FilterByID(id string) *ListFilter {
	for _, lf := range this.listFilters {
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

func (this *ListMenuBase) init(impl IListMenu, cat string, fdesc string) {
	this.fdesc = fdesc
	this.ListBase.init(impl)
	this.cat, this.impl = cat, impl

	for _, lf := range this.listFilters {
		item := &MenuItem{Title: lf.Title, Desc: this.itemDesc(nil, lf)}
		item.IpcID, item.IpcArgs = this.impl.ipcID(lf), lf.ID
		this.items = append(this.items, item)
	}
}

func (this *ListMenuBase) itemDesc(srcLens *SrcLens, lf *ListFilter) string {
	if lf.OnSrcLens != nil && srcLens != nil {
		lf.OnSrcLens(lf, srcLens)
	}
	return Strf(this.fdesc, Lang.Title, lf.Desc)
}

func (this *ListMenuBase) listItemsSubMenu(title string, desc string, filters ListFilters) *Menu {
	listitems := this.impl.List(filters)
	cat := this.cat
	if len(listitems) == 1 && strings.HasSuffix(cat, "s") {
		cat = cat[:len(cat)-1]
	}
	menu := &Menu{Desc: Strf("%d %s: %s (%s)", len(listitems), cat, title, desc)}
	for _, listitem := range listitems {
		if menuitem := this.impl.ListItemToMenuItem(listitem); menuitem != nil {
			menu.Items = append(menu.Items, menuitem)
		}
	}
	return menu
}

func (this *ListMenuBase) MenuCategory() string {
	return this.cat
}

func (this *ListMenuBase) menuItems(srcLens *SrcLens) MenuItems {
	const fhint = "(%v at last count)"
	for _, item := range this.items {
		fcount, filterid := "amount unknown", item.IpcArgs.(string)
		var filters ListFilters
		if filterid != "" {
			lf := this.impl.FilterByID(filterid)
			if lf.OnSrcLens != nil {
				item.Desc = this.itemDesc(srcLens, lf)
			}
			filters = ListFilters{lf: true}
		}
		count := this.impl.Count(filters)
		if count >= 0 {
			fcount = Strf("%d", count)
		}
		item.Hint = Strf(fhint, fcount)
	}
	return this.items
}
