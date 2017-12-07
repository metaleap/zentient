package zgo

import (
	"github.com/metaleap/zentient"
)

type goExtras struct {
	z.ExtrasBase
}

var (
	extras goExtras

	querierGoDoc = z.ExtrasItem{ID: "go_doc", Kind: z.EXTRAS_QUERY, Label: "go doc",
		Description: "[package] [member name]", Detail: "➜ shows the specified item's summary description"}
)

func init() {
	extras.Impl = &extras
	z.Lang.Extras = &extras
}

func (me *goExtras) KindOf(id string) z.ExtrasKind {
	switch id {
	case querierGoDoc.ID:
		return z.EXTRAS_QUERY
	default:
		return z.EXTRAS_INTEL
	}
}

func (me *goExtras) ListIntelExtras() (all []z.ExtrasItem) {
	return
}

func (me *goExtras) ListQueryExtras() (all []z.ExtrasItem) {
	all = append(all, querierGoDoc)
	return
}

func (me *goExtras) RunIntelExtra(srcLens *z.SrcLens, id string, arg string) *z.MenuResp {
	return &z.MenuResp{NoteInfo: "intel " + id}
}

func (me *goExtras) RunQueryExtra(srcLens *z.SrcLens, id string, arg string) *z.MenuResp {
	return &z.MenuResp{NoteWarn: "query " + id}
}
