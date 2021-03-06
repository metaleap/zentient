package zgo

import (
	"path/filepath"
	"strings"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/fs"
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

	goformat  *z.Tool
	gofmt     *z.Tool
	goimports *z.Tool
	goreturns *z.Tool

	godoc    *z.Tool
	go_doc   *z.Tool
	gogetdoc *z.Tool
	godef    *z.Tool
	guru     *z.Tool
	gocode   *z.Tool

	gorename  *z.Tool
	godocdown *z.Tool

	govet        *z.Tool
	golint       *z.Tool
	checkvar     *z.Tool
	checkalign   *z.Tool
	checkstruct  *z.Tool
	errcheck     *z.Tool
	ineffassign  *z.Tool
	interfacer   *z.Tool
	unparam      *z.Tool
	unindent     *z.Tool
	unconvert    *z.Tool
	maligned     *z.Tool
	goconst      *z.Tool
	gosimple     *z.Tool
	unused       *z.Tool
	staticcheck  *z.Tool
	structlayout *z.Tool
	deadcode     *z.Tool
}

func (me *goTooling) onPreInit() {
	me.goformat = &z.Tool{Name: "go/format", Website: "http://golang.org/pkg/go/format", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.gofmt = &z.Tool{Name: "gofmt", Website: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.goimports = &z.Tool{Name: "goimports", Website: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	me.goreturns = &z.Tool{Name: "goreturns", Website: "http://github.com/sqs/goreturns#readme", Installed: udevgo.Has_goreturns, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}

	me.go_doc = &z.Tool{Name: "go doc", Website: "http://golang.org/cmd/doc/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_EXTRAS_QUERY}}
	me.godoc = &z.Tool{Name: "godoc", Website: "http://golang.org/x/tools/cmd/godoc", Installed: udevgo.Has_godoc}
	me.gogetdoc = &z.Tool{Name: "gogetdoc", Website: "http://github.com/zmb3/gogetdoc#readme", Installed: udevgo.Has_gogetdoc, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_CMPL, z.TOOLS_CAT_INTEL_NAV}}
	me.godef = &z.Tool{Name: "godef", Website: "http://github.com/rogpeppe/godef#readme", Installed: udevgo.Has_godef, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_CMPL, z.TOOLS_CAT_INTEL_NAV}}
	me.guru = &z.Tool{Name: "guru", Website: "http://golang.org/x/tools/cmd/guru", Installed: udevgo.Has_guru, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_SYMS, z.TOOLS_CAT_INTEL_HIGH, z.TOOLS_CAT_INTEL_NAV}}
	me.gocode = &z.Tool{Name: "gocode", Website: "http://github.com/nsf/gocode#readme", Installed: udevgo.Has_gocode, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_CMPL}}

	me.gorename = &z.Tool{Name: "gorename", Website: "http://golang.org/x/tools/cmd/gorename", Installed: udevgo.Has_gorename, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_REN}}
	me.godocdown = &z.Tool{Name: "godocdown", Website: "http://github.com/robertkrimen/godocdown", Installed: udevgo.Has_godocdown, Cats: []z.ToolCats{z.TOOLS_CAT_RUNONSAVE}}

	me.govet = &z.Tool{Name: "go vet", Website: "http://golang.org/cmd/vet/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_WARN}
	me.ineffassign = &z.Tool{Name: "ineffassign", Website: "http://github.com/gordonklaus/ineffassign#readme", Installed: udevgo.Has_ineffassign, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_WARN}
	me.golint = &z.Tool{Name: "golint", Website: "http://github.com/golang/lint#readme", Installed: udevgo.Has_golint, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.checkvar = &z.Tool{Name: "varcheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkvar, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.checkalign = &z.Tool{Name: "aligncheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkalign, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.checkstruct = &z.Tool{Name: "structcheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkstruct, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.errcheck = &z.Tool{Name: "errcheck", Website: "http://github.com/kisielk/errcheck#readme", Installed: udevgo.Has_errcheck, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.interfacer = &z.Tool{Name: "interfacer", Website: "http://github.com/mvdan/interfacer#readme", Installed: udevgo.Has_interfacer, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.unparam = &z.Tool{Name: "unparam", Website: "http://github.com/mvdan/unparam#readme", Installed: udevgo.Has_unparam, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.unindent = &z.Tool{Name: "unindent", Website: "http://github.com/mvdan/unindent#readme", Installed: udevgo.Has_unindent, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.unconvert = &z.Tool{Name: "unconvert", Website: "http://github.com/mdempsky/unconvert#readme", Installed: udevgo.Has_unconvert, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.maligned = &z.Tool{Name: "maligned", Website: "http://github.com/mdempsky/maligned#readme", Installed: udevgo.Has_maligned, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.goconst = &z.Tool{Name: "goconst", Website: "http://github.com/jgautheron/goconst#readme", Installed: udevgo.Has_goconst, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.gosimple = &z.Tool{Name: "gosimple", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_gosimple, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.unused = &z.Tool{Name: "unused", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_unused, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.staticcheck = &z.Tool{Name: "staticcheck", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_staticcheck, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	me.structlayout = &z.Tool{Name: "structlayout", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_structlayout, Cats: []z.ToolCats{z.TOOLS_CAT_EXTRAS_QUERY}}
	me.deadcode = &z.Tool{Name: "deadcode", Website: "http://github.com/remyoudompheng/go-misc/tree/master/deadcode#readme", Installed: udevgo.Has_deadcode, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}

	me.all = z.Tools{
		me.goformat,
		me.gofmt,
		me.goimports,
		me.goreturns,
		me.gorename,
		me.gocode,
		me.gogetdoc,
		me.godef,
		me.guru,
		me.go_doc,
		me.govet,
		me.golint,
		me.checkvar,
		me.checkalign,
		me.checkstruct,
		me.errcheck,
		me.ineffassign,
		me.interfacer,
		me.unparam,
		me.unindent,
		me.unconvert,
		me.maligned,
		me.goconst,
		me.gosimple,
		me.unused,
		me.staticcheck,
		me.structlayout,
		me.deadcode,
	}
	me.numInst = me.CountNumInst(me.all)
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

func (me *goTooling) execGodocdown(pkg *udevgo.Pkg) {
	var proceed bool
	for _, imppathprefix := range settings.cfgGddGopaths.ValStrs() {
		if proceed = pkg.ImportPath == imppathprefix || strings.HasPrefix(pkg.ImportPath, imppathprefix+"/"); proceed {
			break
		}
	}
	if readmefilepath := filepath.Join(pkg.Dir, settings.cfgGddFileName.ValStr()); proceed && ufs.IsFile(readmefilepath) {
		go tools.godocdown.Exec(false, "", "godocdown", []string{"-output", readmefilepath, pkg.ImportPath})
	}
}
