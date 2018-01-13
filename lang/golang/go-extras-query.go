package zgo

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
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
					resp.Items = append(resp.Items, &z.ExtrasItem{Label: ustr.FirstNonEmpty(sfname, "—"), Desc: sftype, Detail: ssize})
				}
			}
		}
	}
}

func (me *goExtras) runQuery_Godoc(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if isdocpath := strings.ContainsRune(arg, '/') || strings.ContainsRune(arg, '#'); !isdocpath {
		if isup := ustr.BeginsUpper(arg); isup && udevgo.PkgsByDir != nil {
			if pkg := udevgo.PkgsByDir[filepath.Dir(srcLens.FilePath)]; pkg != nil {
				arg = pkg.ImportPath + "#" + arg
			}
		} else if (!isup) && udevgo.PkgsByImP != nil && nil == udevgo.PkgsByImP[arg] {
			if pkgimppath := uslice.StrWithFewest(udevgo.PkgsByName(arg), "/", uslice.StrShortest); pkgimppath != "" {
				arg = pkgimppath
			}
		}
	}
	resp.Url = "zentient://" + z.Lang.ID + "/godoc/pkg/" + arg
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
	resp.Info = append(resp.Info, z.InfoTip{Value: ustr.Trim(cmdout)})
}

func (me *goExtras) runQuery_GoRun(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	pkgsrcdirpath := filepath.Dir(srcLens.FilePath)
	pkgtmpdirpath := filepath.Join(z.Prog.Dir.Cache, pkgsrcdirpath)
	var err error
	if ufs.DirExists(pkgtmpdirpath) {
		err = ufs.ClearDirectory(pkgtmpdirpath)
	} else {
		err = ufs.EnsureDirExists(pkgtmpdirpath)
	}
	if err == nil && udevgo.PkgsByDir == nil {
		err = udevgo.RefreshPkgs()
	}
	if err == nil {
		if pkg := udevgo.PkgsByDir[pkgsrcdirpath]; pkg == nil {
			err = z.Errf("Not (yet) a Go package: %s", pkgsrcdirpath)
		} else {
			for i, pos, src, gfps := 0, 0, "", pkg.GoFilePaths(); (err == nil) && (i < len(gfps)); i++ {
				iscursrc := gfps[i] == srcLens.FilePath
				if iscursrc && len(srcLens.Txt) > 0 {
					src = srcLens.Txt
				} else {
					src = ufs.ReadTextFile(gfps[i], true, "")
				}
				if strings.HasPrefix(src, "package ") {
					pos = 0
				} else if pos = 1 + strings.Index(src, "\npackage "); pos == 0 {
					err = z.Errf("Not a Go package source file: %s", gfps[i])
				}
				if posln := pos + strings.IndexRune(src[pos:], '\n'); err == nil {
					if oldpkgln := src[pos:posln]; oldpkgln != "package main" {
						src = src[:pos] + "package main" + src[posln:]
					}
					if pos = 1 + strings.Index(src, "\nfunc main() {"); pos > 0 {
						src = src[:pos] + z.Strf("func main%d() {", time.Now().UnixNano()) + src[pos+13:]
					} else if iscursrc {
						src += z.Strf("\n\nfunc main() { println(%s) }", arg)
					}
					err = ufs.WriteTextFile(filepath.Join(pkgtmpdirpath, filepath.Base(gfps[i])), src)
				}
			}
			if err == nil {
				if cmdout, cmderr, e := urun.CmdExecIn(pkgtmpdirpath, "go", "run", filepath.Base(srcLens.FilePath)); e != nil {
					err = e
				} else {
					resp.Desc = arg
					for _, ln := range ustr.Split(strings.TrimSpace(cmderr), "\n") {
						resp.Info = append(resp.Info, z.InfoTip{Value: ln})
					}
					resp.Warns = ustr.Split(cmdout, "\n")
				}
			}
		}
	}

	if err != nil {
		resp.Warns = []string{err.Error()}
	}
}
