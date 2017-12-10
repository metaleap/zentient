package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var tools goTooling

func init() {
	tools.Impl, z.Lang.Tooling = &tools, &tools
}

type goTooling struct {
	z.ToolingBase

	all     z.Tools
	numInst int

	gofmt     *z.Tool
	goimports *z.Tool
	goreturns *z.Tool

	go_doc   *z.Tool
	gogetdoc *z.Tool
	godef    *z.Tool
	guru     *z.Tool

	gorename *z.Tool

	govet       *z.Tool
	golint      *z.Tool
	checkvar    *z.Tool
	checkalign  *z.Tool
	checkstruct *z.Tool
	errcheck    *z.Tool
	ineffassign *z.Tool
	interfacer  *z.Tool
	unparam     *z.Tool
	unconvert   *z.Tool
	maligned    *z.Tool
	goconst     *z.Tool
	gosimple    *z.Tool
	unused      *z.Tool
	staticcheck *z.Tool
}

func (me *goTooling) onPreInit() {
	me.gofmt = &z.Tool{Name: "gofmt", Website: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.goimports = &z.Tool{Name: "goimports", Website: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.goreturns = &z.Tool{Name: "goreturns", Website: "http://github.com/sqs/goreturns#readme", Installed: udevgo.Has_goreturns, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}

	me.go_doc = &z.Tool{Name: "go doc", Website: "http://golang.org/cmd/doc/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_EXTRAS_QUERY}}
	// me.godoc = &z.Tool{Name: "godoc", Website: "http://golang.org/x/tools/cmd/godoc", Installed: udevgo.Has_godoc}
	me.gogetdoc = &z.Tool{Name: "gogetdoc", Website: "http://github.com/zmb3/gogetdoc#readme", Installed: udevgo.Has_gogetdoc, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS}}
	me.godef = &z.Tool{Name: "godef", Website: "http://github.com/rogpeppe/godef#readme", Installed: udevgo.Has_godef, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS}}

	me.gorename = &z.Tool{Name: "gorename", Website: "http://golang.org/x/tools/cmd/gorename", Installed: udevgo.Has_gorename, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_REN}}

	me.govet = &z.Tool{Name: "go vet", Website: "http://golang.org/cmd/vet/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.guru = &z.Tool{Name: "guru", Website: "http://golang.org/x/tools/cmd/guru", Installed: udevgo.Has_guru, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS}}
	me.golint = &z.Tool{Name: "golint", Website: "http://github.com/golang/lint#readme", Installed: udevgo.Has_golint, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.checkvar = &z.Tool{Name: "varcheck", Website: "http://github.com/opennota/check#readme", Installed: udevgo.Has_checkvar, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.checkalign = &z.Tool{Name: "aligncheck", Website: "http://github.com/opennota/check#readme", Installed: udevgo.Has_checkalign, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.checkstruct = &z.Tool{Name: "structcheck", Website: "http://github.com/opennota/check#readme", Installed: udevgo.Has_checkstruct, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.errcheck = &z.Tool{Name: "errcheck", Website: "http://github.com/kisielk/errcheck#readme", Installed: udevgo.Has_errcheck, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.ineffassign = &z.Tool{Name: "ineffassign", Website: "http://github.com/gordonklaus/ineffassign#readme", Installed: udevgo.Has_ineffassign, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.interfacer = &z.Tool{Name: "interfacer", Website: "http://github.com/mvdan/interfacer#readme", Installed: udevgo.Has_interfacer, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.unparam = &z.Tool{Name: "unparam", Website: "http://github.com/mvdan/unparam#readme", Installed: udevgo.Has_unparam, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.unconvert = &z.Tool{Name: "unconvert", Website: "http://github.com/mdempsky/unconvert#readme", Installed: udevgo.Has_unconvert, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.maligned = &z.Tool{Name: "maligned", Website: "http://github.com/mdempsky/maligned#readme", Installed: udevgo.Has_maligned, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.goconst = &z.Tool{Name: "goconst", Website: "http://github.com/jgautheron/goconst#readme", Installed: udevgo.Has_goconst, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.gosimple = &z.Tool{Name: "gosimple", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_gosimple, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.unused = &z.Tool{Name: "unused", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_unused, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}
	me.staticcheck = &z.Tool{Name: "staticcheck", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_staticcheck, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_DIAG}}

	me.all = z.Tools{
		me.gofmt,
		me.goimports,
		me.goreturns,
		me.go_doc,
		me.gogetdoc,
		me.godef,
		me.guru,
		me.gorename,
		me.govet,
		me.golint,
		me.checkvar,
		me.checkalign,
		me.checkstruct,
		me.errcheck,
		me.ineffassign,
		me.interfacer,
		me.unparam,
		me.unconvert,
		me.maligned,
		me.goconst,
		me.gosimple,
		me.unused,
		me.staticcheck,
	}
	me.numInst = me.SortAndCountNumInst(me.all)
}

func (me *goTooling) KnownTools() z.Tools {
	return me.all
}

func (me *goTooling) NumInst() int {
	return me.numInst
}

func (me *goTooling) NumTotal() int {
	return len(me.all)
}
