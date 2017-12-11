package zgo

import (
	"strings"

	"github.com/metaleap/zentient"
)

var (
	extras goExtras

	querierGoDoc = z.ExtrasItem{ID: "go_doc", Label: "go doc",
		Description: "[package] [member name]", Detail: "➜ shows the specified item's summary description",
		QueryArg: "Query to `go doc`"}
)

func init() {
	extras.Impl, z.Lang.Extras = &extras, &extras
}

type goExtras struct {
	z.ExtrasBase
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
		z.BadPanic("CodeIntel ID", id)
	}
}

func (me *goExtras) RunQueryExtra(srcLens *z.SrcLens, id string, arg string, resp *z.ExtrasResp) {
	switch id {
	case querierGoDoc.ID:
		me.runQueryGoDoc(srcLens, strings.TrimSpace(arg), resp)
	default:
		z.BadPanic("CodeQuery ID", id)
	}
}
