package zhs

import (
	"github.com/go-leap/dev/hs"
	"github.com/metaleap/zentient"
)

var tools hsTooling

func init() {
	tools.Impl, z.Lang.Tooling = &tools, &tools
}

type hsTooling struct {
	z.ToolingBase

	all     z.Tools
	numInst int

	hindent        *z.Tool
	brittany       *z.Tool
	stylishhaskell *z.Tool
	hlint          *z.Tool
}

func (me *hsTooling) onPreInit() {
	me.hindent = &z.Tool{Name: "hindent", Website: "http://github.com/commercialhaskell/hindent#readme", Installed: udevhs.Has_hindent, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.stylishhaskell = &z.Tool{Name: "stylish-haskell", Website: "http://github.com/jaspervdj/stylish-haskell#readme", Installed: udevhs.Has_stylish_haskell, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.brittany = &z.Tool{Name: "brittany", Website: "http://github.com/lspitzner/brittany#readme", Installed: udevhs.Has_brittany, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.hlint = &z.Tool{Name: "hlint", Website: "http://github.com/ndmitchell/hlint#readme", Installed: udevhs.Has_hlint, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}

	me.all = z.Tools{
		me.hindent,
		me.stylishhaskell,
		me.brittany,
		me.hlint,
	}
	me.numInst = me.CountNumInst(me.all)
}

func (me *hsTooling) KnownTools() z.Tools {
	return me.all
}

func (me *hsTooling) NumInst() int {
	return me.numInst
}

func (me *hsTooling) NumTotal() int {
	return len(me.all)
}
