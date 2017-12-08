package zgo

import (
	"strings"

	"github.com/metaleap/zentient"
)

type goExtras struct {
	z.ExtrasBase
}

var (
	extras goExtras

	querierGoDoc = z.ExtrasItem{ID: "go_doc", Kind: z.EXTRAS_QUERY, Label: "go doc",
		Description: "[package] [member name]", Detail: "➜ shows the specified item's summary description",
		QueryArg: "Query to `go doc`"}
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

func (me *goExtras) RunIntelExtra(srcLens *z.SrcLens, id string, arg string, resp *z.ExtrasResp) {
	switch id {
	default:
		z.Bad("CodeIntel Extras ID", id)
	}
}

func (me *goExtras) RunQueryExtra(srcLens *z.SrcLens, id string, arg string, resp *z.ExtrasResp) {
	switch id {
	case querierGoDoc.ID:
		me.runQueryGoDoc(srcLens, strings.TrimSpace(arg), resp)
	default:
		z.Bad("CodeQuery Extras ID", id)
	}
}
