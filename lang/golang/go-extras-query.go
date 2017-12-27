package zgo

import (
	"path/filepath"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var (
	xQuerierGoDoc = z.ExtrasItem{ID: "go_doc", Label: "go doc",
		Desc: "[package] [member-name]", Detail: "➜ shows the specified item's summary desc",
		QueryArg: "Query to `go doc`"}
	xQuerierStructlayout = z.ExtrasItem{ID: "structlayout", Label: "structlayout",
		Desc: "[package] struct-name", Detail: "➜ shows the specified struct's memory layout",
		QueryArg: "Specify (optionally) a package and (always) a struct type definition's name"}
)

func (me *goExtras) runQuery_StructLayout(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	args := ustr.Split(arg, " ")
	if len(args) == 1 && udevgo.PkgsByDir != nil && srcLens.FilePath != "" {
		if pkg := udevgo.PkgsByDir[filepath.Dir(srcLens.FilePath)]; pkg != nil {
			args = append([]string{pkg.ImportPath}, args[0])
		}
	}
	if len(args) != 2 {
		z.BadPanic("structlayout args (need 1 or 2)", arg)
	}
	if cmdout, cmderr, err := urun.CmdExec("structlayout", args[0], args[1]); err != nil {
		panic(err)
	} else if cmdout = ustr.Trim(cmdout); cmdout != "" || cmderr != "" {
		scmddesc := z.Strf("structlayout %s %s", args[0], args[1])
		resp.Desc = z.Strf("Results of `%s`, sizes are in bytes:", scmddesc)
		if cmderr != "" {
			resp.Warns = ustr.Split(z.Strf("[%s]\n%s", scmddesc, cmderr), "\n")
		} else if cmdout != "" {
			for _, ln := range ustr.Split(cmdout, "\n") {
				if sfield, ssize := ustr.BreakOnLast(ln, ":"); sfield != "" {
					sfname, sftype := ustr.BreakOn(sfield, " ")
					resp.Items = append(resp.Items, z.ExtrasItem{Label: ustr.FirstNonEmpty(sfname, "—"), Desc: sftype, Detail: ssize})
				}
			}
		}
	}
}

func (me *goExtras) runQuery_GoDoc(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if arg = ustr.Trim(arg); arg == "" {
		return
	}
	if i1, i2 := ustr.Idx(arg, "."), ustr.Idx(arg, " "); i1 > 0 && (i2 < 0 || i2 > i1) {
		arg = arg[:i1] + " " + arg[i1+1:]
	}
	var cmd = ustr.Split(arg, " ")
	if udevgo.PkgsByImP != nil && ustr.IsLower(cmd[0][:1]) && udevgo.PkgsByImP[cmd[0]] == nil {
		for _, pkg := range udevgo.PkgsByImP {
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
	resp.Warns = uslice.StrFiltered(uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim),
		func(s string) bool { return !ustr.Pref(s, "exit status ") })
	resp.InfoTips = append(resp.InfoTips, z.InfoTip{Value: ustr.Trim(cmdout)})
}
