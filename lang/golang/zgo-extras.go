package zgo

import (
	"strings"

	"github.com/metaleap/zentient"
)

var (
	extras goExtras
)

func init() {
	extras.Impl, z.Lang.Extras = &extras, &extras
}

type goExtras struct {
	z.ExtrasBase
}

func (*goExtras) ListIntelExtras() (all []*z.ExtrasItem) {
	all = []*z.ExtrasItem{&xIntelGuruCallers, &xIntelGuruCallees, &xIntelGuruCallstack, &xIntelGuruFreevars, &xIntelGuruErrtypes, &xIntelGuruPointeeTypes, &xIntelGuruPointeeVals, &xIntelGuruChanpeers}
	return
}

func (*goExtras) ListQueryExtras() (all []*z.ExtrasItem) {
	all = []*z.ExtrasItem{&xQuerierGoRun, &xQuerierGodoc, &xQuerierGoDoc, &xQuerierStructlayout}
	return
}

func (this *goExtras) RunIntelExtra(srcLens *z.SrcLens, id string, arg string, resp *z.ExtrasResp) {
	if strings.HasPrefix(id, "guru.") {
		this.runIntel_Guru(id[5:], srcLens, strings.TrimSpace(arg), resp)
	} else {
		z.BadPanic("CodeIntel ID", id)
	}
}

func (this *goExtras) RunQueryExtra(srcLens *z.SrcLens, id string, arg string, resp *z.ExtrasResp) {
	var runner func(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp)
	switch id {
	case xQuerierGoRun.ID:
		runner = this.runQuery_GoRun
	case xQuerierGodoc.ID:
		runner = this.runQuery_Godoc
	case xQuerierGoDoc.ID:
		runner = this.runQuery_GoDoc
	case xQuerierStructlayout.ID:
		runner = this.runQuery_StructLayout
	default:
		z.BadPanic("CodeQuery ID", id)
	}
	if runner != nil {
		runner(srcLens, strings.TrimSpace(arg), resp)
	}
}
