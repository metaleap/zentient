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

func (me *goExtras) RunIntelExtra(srcLens *z.SrcLens, id string, arg string, resp *z.IpcRespExtras) {
	if strings.HasPrefix(id, "guru.") {
		me.runIntel_Guru(id[5:], srcLens, strings.TrimSpace(arg), resp)
	} else {
		z.BadPanic("CodeIntel ID", id)
	}
}

func (me *goExtras) RunQueryExtra(srcLens *z.SrcLens, id string, arg string, resp *z.IpcRespExtras) {
	var runner func(srcLens *z.SrcLens, arg string, resp *z.IpcRespExtras)
	switch id {
	case xQuerierGoRun.ID:
		runner = me.runQuery_GoRun
	case xQuerierGodoc.ID:
		runner = me.runQuery_Godoc
	case xQuerierGoDoc.ID:
		runner = me.runQuery_GoDoc
	case xQuerierStructlayout.ID:
		runner = me.runQuery_StructLayout
	default:
		z.BadPanic("CodeQuery ID", id)
	}
	if runner != nil {
		runner(srcLens, strings.TrimSpace(arg), resp)
	}
}
