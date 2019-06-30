package zgo

import (
	"path/filepath"
	"strings"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/run"
	"github.com/go-leap/str"
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/golang/dbg"
)

var (
	xQuerierGoRun = z.ExtrasItem{ID: "gorun", Label: "eval via `go run`",
		Desc: "foo", Detail: "➜ evaluates the specified expression in the context of this file's package",
		QueryArg: "q-args"}
	xQuerierGodoc = z.ExtrasItem{ID: "godoc", Label: "godoc",
		Desc: "package[/path][#Name]", Detail: "➜ opens the godoc page for the specified package",
		QueryArg: "package[/path][#Name]"}
	xQuerierGoDoc = z.ExtrasItem{ID: "go_doc", Label: "go doc",
		Desc: "[package] [member-name]", Detail: "➜ shows the specified item's summary desc",
		QueryArg: "[package] [member-name]"}
	xQuerierStructlayout = z.ExtrasItem{ID: "structlayout", Label: "structlayout",
		Desc: "[package] struct-name", Detail: "➜ shows the specified struct's memory layout",
		QueryArg: "Specify (optionally) a package and (always) a struct type definition's name"}
)

func (*goExtras) runQuery_StructLayout(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	args := ustr.Split(arg, " ")
	if pkgsbydir := udevgo.PkgsByDir; len(args) == 1 && pkgsbydir != nil && srcLens.FilePath != "" {
		if pkg := pkgsbydir[filepath.Dir(srcLens.FilePath)]; pkg != nil {
			args = append([]string{pkg.ImportPath}, args[0])
		}
	}
	if len(args) != 2 {
		z.BadPanic(tools.structlayout.Name+" args (need 1 or 2)", arg)
	}
	if cmdout, cmderr, err := urun.CmdExec(tools.structlayout.Name, args[0], args[1]); err != nil {
		panic(err)
	} else if cmdout = ustr.Trim(cmdout); cmdout != "" || cmderr != "" {
		scmddesc := z.Strf("structlayout %s %s", args[0], args[1])
		resp.Desc = z.Strf("Results of `%s`, sizes are in bytes:", scmddesc)
		if cmderr != "" {
			resp.Warns = ustr.Split(z.Strf("[%s]\n%s", scmddesc, cmderr), "\n")
		} else if cmdout != "" {
			for _, ln := range ustr.Split(cmdout, "\n") {
				if sfield, ssize := ustr.BreakOnLast(ln, ":"); sfield != "" {
					sfname, sftype := ustr.BreakOnFirstOrSuff(sfield, " ")
					resp.Items = append(resp.Items, &z.ExtrasItem{Label: ustr.FirstOf(sfname, "—"), Desc: sftype, Detail: ssize})
				}
			}
		}
	}
}

func (*goExtras) runQuery_Godoc(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if isdocpath := strings.ContainsRune(arg, '/') || strings.ContainsRune(arg, '#'); !isdocpath {
		if isup, pkgsbydir := ustr.BeginsUpper(arg), udevgo.PkgsByDir; isup && pkgsbydir != nil {
			if pkg := pkgsbydir[filepath.Dir(srcLens.FilePath)]; pkg != nil {
				arg = pkg.ImportPath + "#" + arg
			}
		} else if pkgsbyimp := udevgo.PkgsByImP; (!isup) && pkgsbyimp != nil && nil == pkgsbyimp[arg] {
			if pkgimppath := ustr.Fewest(udevgo.PkgsByName(arg), "/", ustr.Shortest); pkgimppath != "" {
				arg = pkgimppath
			}
		}
	}
	resp.Url = "zentient://" + z.Lang.ID + "/godoc/pkg/" + arg
}

func (*goExtras) runQuery_GoDoc(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if arg = ustr.Trim(arg); arg == "" {
		return
	}
	if i1, i2 := ustr.Pos(arg, "."), ustr.Pos(arg, " "); i1 > 0 && (i2 < 0 || i2 > i1) {
		arg = arg[:i1] + " " + arg[i1+1:]
	}
	var cmd = ustr.Split(arg, " ")
	if pkgsbyimp := udevgo.PkgsByImP; pkgsbyimp != nil && ustr.IsLower(cmd[0][:1]) && pkgsbyimp[cmd[0]] == nil {
		for _, pkg := range pkgsbyimp {
			if pkg.Name == cmd[0] {
				cmd[0] = pkg.ImportPath
				break
			}
		}
	}
	cmdout, cmderr, err := urun.CmdExecIn(filepath.Dir(srcLens.FilePath), "go", append([]string{"doc"}, cmd...)...)
	if err != nil {
		panic(err)
	}
	resp.Desc = z.Strf("Results of `go doc %s`:", ustr.Join(cmd, " "))
	resp.Warns = ustr.Filter(ustr.Map(ustr.Split(cmderr, "\n"), ustr.Trim),
		func(s string) bool { return !ustr.Pref(s, "exit status ") })
	resp.InfoTips = append(resp.InfoTips, z.InfoTip{Value: ustr.Trim(cmdout)})
}

func (*goExtras) runQuery_GoRun(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	evaloutandstderr, otherstdout, err := zgodbg.GoRunEval(z.Prog.Dir.Cache, srcLens.FilePath, srcLens.Txt, arg)
	if resp.Desc = arg; err != nil {
		panic(err)
	}
	for _, ln := range ustr.Split(strings.TrimSpace(evaloutandstderr), "\n") {
		resp.InfoTips = append(resp.InfoTips, z.InfoTip{Value: ln})
	}
	resp.Warns = ustr.Split(otherstdout, "\n")
}
