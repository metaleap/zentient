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

func (me *goExtras) ListIntelExtras() (all []z.ExtrasItem) {
	return
}

func (me *goExtras) ListQueryExtras() (all []z.ExtrasItem) {
	all = append(all, querierGoDoc)
	return
}
