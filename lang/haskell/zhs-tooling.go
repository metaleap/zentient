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

func (this *hsTooling) onPreInit() {
	this.hindent = &z.Tool{Name: "hindent", Website: "http://github.com/commercialhaskell/hindent#readme", Installed: udevhs.Has_hindent, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.stylishhaskell = &z.Tool{Name: "stylish-haskell", Website: "http://github.com/jaspervdj/stylish-haskell#readme", Installed: udevhs.Has_stylish_haskell, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.brittany = &z.Tool{Name: "brittany", Website: "http://github.com/lspitzner/brittany#readme", Installed: udevhs.Has_brittany, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.hlint = &z.Tool{Name: "hlint", Website: "http://github.com/ndmitchell/hlint#readme", Installed: udevhs.Has_hlint, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}

	this.all = z.Tools{
		this.hindent,
		this.stylishhaskell,
		this.brittany,
		this.hlint,
	}
	this.numInst = this.CountNumInst(this.all)
}

func (this *hsTooling) KnownTools() z.Tools {
	return this.all
}

func (this *hsTooling) NumInst() int {
	return this.numInst
}

func (this *hsTooling) NumTotal() int {
	return len(this.all)
}
