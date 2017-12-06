package zgo

import (
	"github.com/metaleap/zentient"
)

type goExtras struct {
	z.ExtrasBase
}

var (
	extras goExtras

	querierRun = z.ExtrasItem{ID: "go_doc", Label: "go doc", Description: "[package] [member name]", Detail: "➜ shows the specified item's summary description"}
)

func init() {
	extras.Impl = &extras
	z.Lang.Extras = &extras
}

func (me *goExtras) ListIntelExtras() (all []z.ExtrasItem) {
	all = append(all, z.ExtrasItem{Label: "intel title", Description: "intel desc", Detail: "intel detail"})
	return
}

func (me *goExtras) ListQueryExtras() (all []z.ExtrasItem) {
	all = append(all, querierRun)
	return
}
