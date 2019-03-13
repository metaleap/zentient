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

func (this *goTooling) onPreInit() {
	this.goformat = &z.Tool{Name: "go/format", Website: "http://golang.org/pkg/go/format", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.gofmt = &z.Tool{Name: "gofmt", Website: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.goimports = &z.Tool{Name: "goimports", Website: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}
	this.goreturns = &z.Tool{Name: "goreturns", Website: "http://github.com/sqs/goreturns#readme", Installed: udevgo.Has_goreturns, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_FMT}}

	this.go_doc = &z.Tool{Name: "go doc", Website: "http://golang.org/cmd/doc/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_EXTRAS_QUERY}}
	this.godoc = &z.Tool{Name: "godoc", Website: "http://golang.org/x/tools/cmd/godoc", Installed: udevgo.Has_godoc}
	this.gogetdoc = &z.Tool{Name: "gogetdoc", Website: "http://github.com/zmb3/gogetdoc#readme", Installed: udevgo.Has_gogetdoc, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_CMPL, z.TOOLS_CAT_INTEL_NAV}}
	this.godef = &z.Tool{Name: "godef", Website: "http://github.com/rogpeppe/godef#readme", Installed: udevgo.Has_godef, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_CMPL, z.TOOLS_CAT_INTEL_NAV}}
	this.guru = &z.Tool{Name: "guru", Website: "http://golang.org/x/tools/cmd/guru", Installed: udevgo.Has_guru, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_TIPS, z.TOOLS_CAT_INTEL_SYMS, z.TOOLS_CAT_INTEL_HIGH, z.TOOLS_CAT_INTEL_NAV}}
	this.gocode = &z.Tool{Name: "gocode", Website: "http://github.com/nsf/gocode#readme", Installed: udevgo.Has_gocode, Cats: []z.ToolCats{z.TOOLS_CAT_INTEL_CMPL}}

	this.gorename = &z.Tool{Name: "gorename", Website: "http://golang.org/x/tools/cmd/gorename", Installed: udevgo.Has_gorename, Cats: []z.ToolCats{z.TOOLS_CAT_MOD_REN}}
	this.godocdown = &z.Tool{Name: "godocdown", Website: "http://github.com/robertkrimen/godocdown", Installed: udevgo.Has_godocdown, Cats: []z.ToolCats{z.TOOLS_CAT_RUNONSAVE}}

	this.govet = &z.Tool{Name: "go vet", Website: "http://golang.org/cmd/vet/", Installed: true, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_WARN}
	this.ineffassign = &z.Tool{Name: "ineffassign", Website: "http://github.com/gordonklaus/ineffassign#readme", Installed: udevgo.Has_ineffassign, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_WARN}
	this.golint = &z.Tool{Name: "golint", Website: "http://github.com/golang/lint#readme", Installed: udevgo.Has_golint, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.checkvar = &z.Tool{Name: "varcheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkvar, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.checkalign = &z.Tool{Name: "aligncheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkalign, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.checkstruct = &z.Tool{Name: "structcheck", Website: "http://gitlab.com/opennota/check#readme", Installed: udevgo.Has_checkstruct, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.errcheck = &z.Tool{Name: "errcheck", Website: "http://github.com/kisielk/errcheck#readme", Installed: udevgo.Has_errcheck, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.interfacer = &z.Tool{Name: "interfacer", Website: "http://github.com/mvdan/interfacer#readme", Installed: udevgo.Has_interfacer, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.unparam = &z.Tool{Name: "unparam", Website: "http://github.com/mvdan/unparam#readme", Installed: udevgo.Has_unparam, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.unindent = &z.Tool{Name: "unindent", Website: "http://github.com/mvdan/unindent#readme", Installed: udevgo.Has_unindent, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.unconvert = &z.Tool{Name: "unconvert", Website: "http://github.com/mdempsky/unconvert#readme", Installed: udevgo.Has_unconvert, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.maligned = &z.Tool{Name: "maligned", Website: "http://github.com/mdempsky/maligned#readme", Installed: udevgo.Has_maligned, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.goconst = &z.Tool{Name: "goconst", Website: "http://github.com/jgautheron/goconst#readme", Installed: udevgo.Has_goconst, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.gosimple = &z.Tool{Name: "gosimple", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_gosimple, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.unused = &z.Tool{Name: "unused", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_unused, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.staticcheck = &z.Tool{Name: "staticcheck", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_staticcheck, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}
	this.structlayout = &z.Tool{Name: "structlayout", Website: "http://github.com/dominikh/go-tools#readme", Installed: udevgo.Has_structlayout, Cats: []z.ToolCats{z.TOOLS_CAT_EXTRAS_QUERY}}
	this.deadcode = &z.Tool{Name: "deadcode", Website: "http://github.com/remyoudompheng/go-misc/tree/master/deadcode#readme", Installed: udevgo.Has_deadcode, Cats: []z.ToolCats{z.TOOLS_CAT_DIAGS}, DiagSev: z.DIAG_SEV_INFO}

	this.all = z.Tools{
		this.goformat,
		this.gofmt,
		this.goimports,
		this.goreturns,
		this.gorename,
		this.gocode,
		this.gogetdoc,
		this.godef,
		this.guru,
		this.go_doc,
		this.govet,
		this.golint,
		this.checkvar,
		this.checkalign,
		this.checkstruct,
		this.errcheck,
		this.ineffassign,
		this.interfacer,
		this.unparam,
		this.unindent,
		this.unconvert,
		this.maligned,
		this.goconst,
		this.gosimple,
		this.unused,
		this.staticcheck,
		this.structlayout,
		this.deadcode,
	}
	this.numInst = this.CountNumInst(this.all)
}

func (this *goTooling) KnownTools() z.Tools {
	return this.all
}

func (this *goTooling) NumInst() int {
	return this.numInst
}

func (this *goTooling) NumTotal() int {
	return len(this.all)
}

func (this *goTooling) execGodocdown(pkg *udevgo.Pkg) {
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
