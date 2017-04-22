package zgo
import (
	"path/filepath"
	"time"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-slice"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)

var queryTools = []*z.RespPick {
		&z.RespPick{ Label: "go run", Detail: "any expression – attempts to evaluate the specified expression given the current source context", Desc: "go run" },
		&z.RespPick{ Label: "go doc", Detail: "[package] [member name] – shows the specified item's summary description.", Desc: "go doc" },
	}
func (me *zgo) QueryTools () []*z.RespPick {
	return queryTools
}


func (_ *zgo) QueryTool (req *z.ReqIntel) (resp *z.RespTxt) {
	resp = &z.RespTxt{ Id: req.Id }
	switch req.Id {
	case "go run":
		onerr := func (errs ...error) { for _,e := range errs { if e!=nil { resp.Warnings = append(resp.Warnings, e.Error()) } } }
		pname := "tmp_repl_" + ugo.SPr(time.Now().UnixNano())  ;  pdir := filepath.Join(z.Ctx.CacheDir, pname)
		srcmainify := func (src string) string {
			mln,lns := -1, ustr.Split(src, "\n")  ;  for i,ln := range lns {
				if mln<0 && ustr.Pref(ln, "package ") { mln = i  ;  lns[i] = "package main"  }
				if ustr.HasAnyPrefix(ln, "func main(", "func main ") { // good enough for us
					lns[i] = "func " + pname + ln[9:]
				}
			}
			if mln<0 { mln = 0  ;  lns = append([]string {"package main"}, lns...)  }
			src = ustr.Join(lns[mln:], "\n")
			return src
		}
		req.EnsureSrc()
		onerr(ufs.EnsureDirExists(pdir))
		srcfiles := map[string]string {}
		req.Src = srcmainify(req.Src) + "\n\nfunc main() {\n" +
											"\tfmt.Printf(\"%v\", "+ req.Sym2 +")\n" +
												"}"
		if !ufs.FileExists(req.Ffp) {
			srcfiles ["main.go"] = req.Src
		} else {
			srcdir  := filepath.Dir(req.Ffp) ;  curfrp := filepath.Base(req.Ffp)
			onerr(ufs.WalkFilesIn(srcdir, func(ffp string) bool {
				if frp := filepath.Base(ffp)  ;  ustr.Suff(frp, ".go") {
					magicval := pname + ugo.SPr(time.Now().UnixNano()) + pname
					if frp==curfrp {
						srcfiles[frp] = req.Src
					} else if src := ufs.ReadTextFile(ffp, false, magicval); src!=magicval {
						srcfiles[frp] = srcmainify(src)
					} else {
						onerr(ugo.E("Could not read: " + ffp))
					}
				}
				return true
			})...)
		}

		fixupmissingimport := func (msg string) bool {
			//	./stmt.go:80: undefined: fmt in fmt.Printf
			if icol := ustr.Idx(msg, ".go:")  ;  icol>0 {
				if iundef := ustr.Idx(msg, ": undefined: ")  ;  iundef>0 {
					if iin := ustr.Idx(msg, " in ")  ;  iin>(iundef+13) {
						frp := msg[:icol+3]  ;  for frp[0]=='.' || frp[0]=='/' || frp[0]=='\\' { frp = frp[1:] }
						pkg := msg[:iin][iundef+13:]  ;  if _,ok := srcfiles[frp]  ;  ok {
							if devgo.PkgsByImP!=nil {  for _,p := range devgo.PkgsByImP { if p.Name==pkg { pkg = p.ImportPath  ;  break } }  }
							insert,src := "\nimport \"" + pkg + "\"\n" , srcfiles[frp]  ;  if iln := ustr.Idx(src, "\n")  ;  iln<0 {
								src += insert
							} else {
								src = src[:iln] + insert + src[iln:]
							}
							srcfiles[frp] = src
							return true
						}
					}
				}
			}
			return false
		}
		tryagain := true  ;  for tryagain {
			warns := []string {}  ;  cmdargs := []string{ "run" }  ;  tryagain = false
			for frp,src := range srcfiles {
				onerr(ufs.WriteTextFile(filepath.Join(pdir, frp), src))
				cmdargs = append(cmdargs, frp)
			}
			cmdout,cmderr,_ := ugo.CmdExecStdin("", pdir, "go", cmdargs...)
			warns = append(warns, ustr.Split(cmderr, "\n")...)
			for _,warn := range warns {
				if fixupmissingimport(warn) {  tryagain = true  ;  break  }
			}
			if !tryagain {
				if len(warns)>1 {
					warns = uslice.StrWithout(warns, true, "# command-line-arguments")
				}
				resp.Result,resp.Warnings = cmdout , append(resp.Warnings, warns...)
			}
		}
		ufs.ClearDirectory(pdir)
		ufs.ClearEmptyDirectories(filepath.Dir(pdir))
	case "go doc":
		req.Sym2 = ustr.Trim(req.Sym2)
		if i1,i2 := ustr.Idx(req.Sym2, ".") , ustr.Idx(req.Sym2, " ")  ;  i1>0 && (i2<0 || i2>i1) { req.Sym2 = req.Sym2[:i1] + " " + req.Sym2[i1+1:] }
		var cmd = ustr.Split(req.Sym2, " ")  ;  if devgo.PkgsByImP!=nil && ustr.IsLower(cmd[0][:1]) && devgo.PkgsByImP[cmd[0]]==nil {
			if dp := filepath.Join(srcDir, cmd[0])  ;  devgo.PkgsByDir!=nil && devgo.PkgsByDir[dp]!=nil {
				cmd[0] = devgo.PkgsByDir[dp].ImportPath
			} else { for _,pkg := range devgo.PkgsByImP {
				if pkg.Name==cmd[0] { cmd[0] = pkg.ImportPath  ;  break }
			} }
		}
		cmdout,cmderr,_ := ugo.CmdExecStdin ("", filepath.Dir(req.Ffp), "go", append([]string{ "doc" }, cmd...)...)
		resp.Warnings = uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim)
		resp.Result = cmdout
	default:
		resp.Warnings = []string{ "Unknown querier: " + req.Id }
	}
	if len(resp.Warnings)>1 { resp.Warnings = uslice.StrWithout(resp.Warnings, true, "exit status 1") }
	return
}




		/*
[00:30 rox ~/.../metaleap/tmpprog]$ go run *.go
go run: cannot run non-main package
[00:30 rox ~/.../metaleap/tmpprog]$ go run *.go
go run: cannot run non-main package
[00:31 rox ~/.../metaleap/tmpprog]$ go run *.go
package main:
buf.go:1:1: expected 'package', found 'import'
buf.go:2:2: expected ';', found 'STRING' "bytes"
[00:31 rox ~/.../metaleap/tmpprog]$ go run *.go
102030
[00:31 rox ~/.../metaleap/tmpprog]$ go run *.go
# command-line-arguments
runtime.main: call to external function main.main
runtime.main: main.main: not defined
runtime.main: undefined: main.main
[00:32 rox ~/.../metaleap/tmpprog]$ go run *.go
102030
[00:32 rox ~/.../metaleap/tmpprog]$
		*/
