package zgo

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/zentient"
)

func goRunEvalOrPanic(srcFilePath string, maybeSrcFull string, goEvalExpr string) (evalOutAndStdErr string, otherStdOut string) {
	pkgsrcdirpath := filepath.Dir(srcFilePath)
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
		if cmdargs, pkg := []string{"run"}, udevgo.PkgsByDir[pkgsrcdirpath]; pkg == nil {
			z.BadPanic("Go package", pkgsrcdirpath)
		} else {
			defer ufs.ClearDirectory(pkgtmpdirpath)
			for i, pos, src, gfps := 0, 0, "", pkg.GoFilePaths(); (err == nil) && (i < len(gfps)); i++ {
				iscursrc := gfps[i] == srcFilePath
				if iscursrc && len(maybeSrcFull) > 0 {
					src = maybeSrcFull
				} else {
					src = ufs.ReadTextFile(gfps[i], true, "")
				}
				if strings.HasPrefix(src, "package ") {
					pos = 0
				} else if pos = 1 + strings.Index(src, "\npackage "); pos == 0 {
					z.BadPanic("Go package source file", gfps[i])
				}
				if posln := pos + strings.IndexRune(src[pos:], '\n'); err == nil {
					if oldpkgln := src[pos:posln]; oldpkgln != "package main" {
						src = src[:pos] + "package main" + src[posln:]
					}
					if pos = 1 + strings.Index(src, "\nfunc main() {"); pos > 0 {
						src = src[:pos] + z.Strf("func main%d() {", time.Now().UnixNano()) + src[pos+13:]
					} else if iscursrc {
						src += z.Strf("\n\nfunc main() { println(%s) }", goEvalExpr)
					}
					cmdargs = append(cmdargs, filepath.Base(gfps[i]))
					err = ufs.WriteTextFile(filepath.Join(pkgtmpdirpath, cmdargs[len(cmdargs)-1]), src)
				}
			}
			if err == nil {
				otherStdOut, evalOutAndStdErr, err = urun.CmdExecIn(pkgtmpdirpath, "go", cmdargs...)
			}
		}
	}
	if err != nil {
		panic(err)
	}
	return
}
