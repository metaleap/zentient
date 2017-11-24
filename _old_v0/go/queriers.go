package zgo

import (
	"path/filepath"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient/z"
)

var queryTools = []*z.RespPick{
	&z.RespPick{Label: "go run", Detail: "any expression – attempts to evaluate the specified expression given the current source context", Desc: "go run"},
	&z.RespPick{Label: "go doc", Detail: "[package] [member name] – shows the specified item's summary description.", Desc: "go doc"},
}

func (me *zgo) QueryTools() []*z.RespPick {
	return queryTools
}

func (_ *zgo) QueryTool(req *z.ReqIntel) (resp *z.RespTxt) {
	resp = &z.RespTxt{Id: req.Id}
	switch req.Id {
	case "go run":
		onerr := func(errs ...error) {
			for _, e := range errs {
				if e != nil {
					resp.Warnings = append(resp.Warnings, e.Error())
				}
			}
		}
		pname := "tmp_repl_" + umisc.Str(time.Now().UnixNano())
		pdir := filepath.Join(z.Ctx.CacheDir, pname)
		srcmainify := func(src string) string {
			mln, lns := -1, ustr.Split(src, "\n")
			for i, ln := range lns {
				if mln < 0 && ustr.Pref(ln, "package ") {
					mln = i
					lns[i] = "package main"
				}
				if ustr.HasAnyPrefix(ln, "func main(", "func main ") { // good enough for us
					lns[i] = "func " + pname + ln[9:]
				}
			}
			if mln < 0 {
				mln = 0
				lns = append([]string{"package main"}, lns...)
			}
			src = ustr.Join(lns[mln:], "\n")
			return src
		}
		onerr(ufs.EnsureDirExists(pdir))
		srcfiles := map[string]string{}
		req.EnsureSrc()
		numainfunc := "\nfunc main() { fmt.Printf(\"%v\", " + req.Sym2 + ") }\n"
		isvoidcall := (len(req.Sym2) > 0 && req.Sym2[0] == ':')
		if isvoidcall {
			numainfunc = "\nfunc main() { " + req.Sym2[1:] + " } \n"
		}
		req.Src = srcmainify(req.Src) + numainfunc
		if !ufs.FileExists(req.Ffp) {
			srcfiles["main.go"] = req.Src
		} else {
			srcdir := filepath.Dir(req.Ffp)
			curfrp := filepath.Base(req.Ffp)
			onerr(ufs.WalkFilesIn(srcdir, func(ffp string) bool {
				if frp := filepath.Base(ffp); ustr.Suff(frp, ".go") && !ustr.Suff(frp, "_test.go") {
					magicval := pname + umisc.Str(time.Now().UnixNano()) + pname
					if frp == curfrp {
						srcfiles[frp] = req.Src
					} else if src := ufs.ReadTextFile(ffp, false, magicval); src != magicval {
						srcfiles[frp] = srcmainify(src)
					} else {
						onerr(umisc.E("Could not read: " + ffp))
					}
				}
				return true
			})...)
		}

		fixupmissingimport := func(msg string) string {
			//	./file.go:80: undefined: fmt in fmt.Printf
			if icol := ustr.Idx(msg, ".go:"); icol > 0 {
				if iundef := ustr.Idx(msg, ": undefined: "); iundef > 0 {
					if iin := ustr.Idx(msg, " in "); iin > (iundef + 13) {
						frp := msg[:icol+3]
						for frp[0] == '.' || frp[0] == '/' || frp[0] == '\\' {
							frp = frp[1:]
						}
						pkg := msg[:iin][iundef+13:]
						if _, ok := srcfiles[frp]; ok {
							if udevgo.PkgsByImP != nil {
								for _, p := range udevgo.PkgsByImP {
									if p.Name == pkg {
										pkg = p.ImportPath
										break
									}
								}
							}
							insert, src := "\nimport \""+pkg+"\"\n", srcfiles[frp]
							if iln := ustr.Idx(src, "\n"); iln < 0 {
								src += insert
							} else {
								src = src[:iln] + insert + src[iln:]
							}
							srcfiles[frp] = src
							return frp
						}
					}
				}
			}
			return ""
		}
		fixeruppers := []func(string) string{fixupmissingimport}
		frpmod := ""
		tryagain := true
		for tryagain {
			tryagain = false
			warns := []string{}
			cmdargs := []string{"run"}
			for frp, src := range srcfiles {
				if len(frpmod) == 0 || frp == frpmod {
					onerr(ufs.WriteTextFile(filepath.Join(pdir, frp), src))
				}
				cmdargs = append(cmdargs, frp)
			}
			frpmod = ""
			cmdout, cmderr, _ := urun.CmdExecStdin("", pdir, "go", cmdargs...)
			warns = append(warns, ustr.Split(cmderr, "\n")...)
			for _, warn := range warns {
				for _, fixup := range fixeruppers {
					if frpmod = fixup(warn); len(frpmod) > 0 {
						break
					}
				}
				if len(frpmod) > 0 {
					tryagain = true
					break
				}
			}
			if !tryagain {
				if len(warns) > 1 {
					warns = uslice.StrWithout(warns, true, "# command-line-arguments")
				}
				for i, w := range warns {
					if ustr.Suff(w, " used as value") {
						warns[i] = w + " ➜ use a single colon prefix ❬:❭ to evaluate a `void` expression"
					}
				}
				if resp.Result, resp.Warnings = cmdout, append(resp.Warnings, warns...); len(ustr.Trim(resp.Result)) == 0 && isvoidcall && len(warns) == 0 {
					resp.Result = "❬Code ran but did not produce any output❭"
				}
			}
		}
		ufs.ClearDirectory(pdir)
		ufs.ClearEmptyDirectories(filepath.Dir(pdir))
	case "go doc":
		req.Sym2 = ustr.Trim(req.Sym2)
		if i1, i2 := ustr.Idx(req.Sym2, "."), ustr.Idx(req.Sym2, " "); i1 > 0 && (i2 < 0 || i2 > i1) {
			req.Sym2 = req.Sym2[:i1] + " " + req.Sym2[i1+1:]
		}
		var cmd = ustr.Split(req.Sym2, " ")
		if udevgo.PkgsByImP != nil && ustr.IsLower(cmd[0][:1]) && udevgo.PkgsByImP[cmd[0]] == nil {
			if dp := filepath.Join(srcDir, cmd[0]); udevgo.PkgsByDir != nil && udevgo.PkgsByDir[dp] != nil {
				cmd[0] = udevgo.PkgsByDir[dp].ImportPath
			} else {
				for _, pkg := range udevgo.PkgsByImP {
					if pkg.Name == cmd[0] {
						cmd[0] = pkg.ImportPath
						break
					}
				}
			}
		}
		cmdout, cmderr, _ := urun.CmdExecStdin("", filepath.Dir(req.Ffp), "go", append([]string{"doc"}, cmd...)...)
		resp.Warnings = uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim)
		resp.Result = cmdout
	default:
		resp.Warnings = []string{"Unknown querier: " + req.Id}
	}
	if len(resp.Warnings) > 1 {
		resp.Warnings = uslice.StrWithout(resp.Warnings, true, "exit status 1")
	}
	return
}
