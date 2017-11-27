package zhs

import (
	"github.com/metaleap/go-util/dev/hs"
	"github.com/metaleap/zentient"
)

var tools struct {
	hindent        *z.Tool
	brittany       *z.Tool
	stylishhaskell *z.Tool
}

func toolsInit() {
	t := &tools

	t.stylishhaskell = &z.Tool{Name: "stylish-haskell", Website: "http://github.com/jaspervdj/stylish-haskell#readme", Installed: udevhs.Has_stylish_haskell}
	t.hindent = &z.Tool{Name: "hindent", Website: "http://github.com/commercialhaskell/hindent#readme", Installed: udevhs.Has_hindent}
	t.brittany = &z.Tool{Name: "brittany", Website: "http://github.com/lspitzner/brittany#readme", Installed: udevhs.Has_brittany}
}
