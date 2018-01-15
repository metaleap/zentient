package zgo

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/dbg"
)

type Dbg struct {
	zdbg.Dbg
}

func (me *Dbg) Init(tmpDirPath string, srcFilePath string, maybeSrcFull string) (err error) {
	var gorunargs []string
	if gorunargs, me.Cmd.Dir, err = goRunEvalPrepCmd(tmpDirPath, srcFilePath, maybeSrcFull, ""); err == nil {
		me.Cmd.Name = "zdbg-main-" + filepath.Base(tmpDirPath)
		gobuildargs := append([]string{"build", "-o", me.Cmd.Name}, gorunargs[1:]...)
		if _, cmderr, e := urun.CmdExecStdin("", me.Cmd.Dir, "go", gobuildargs...); e != nil {
			err = e
		} else if cmderr != "" {
			err = umisc.E(cmderr)
		}
	}
	return
}

func goRunEvalPrepCmd(tmpDirPath string, srcFilePath string, maybeSrcFull string, goEvalExpr string) (goRunArgs []string, goRunDir string, err error) {
	pkgsrcdirpath := filepath.Dir(srcFilePath)
	goRunDir = filepath.Join(tmpDirPath, pkgsrcdirpath)
	if ufs.DirExists(goRunDir) {
		err = ufs.ClearDirectory(goRunDir)
	} else {
		err = ufs.EnsureDirExists(goRunDir)
	}
	if err == nil && udevgo.PkgsByDir == nil {
		err = udevgo.RefreshPkgs()
	}
	if goRunArgs = []string{"run"}; err == nil {
		if pkg := udevgo.PkgsByDir[pkgsrcdirpath]; pkg == nil {
			err = umisc.E(z.BadMsg("Go package", pkgsrcdirpath))
		} else {
			for i, pos, src, gfps := 0, 0, "", pkg.GoFilePaths(false); (err == nil) && (i < len(gfps)); i++ {
				iscursrc := gfps[i] == srcFilePath
				if iscursrc && len(maybeSrcFull) > 0 {
					src = maybeSrcFull
				} else if err = ufs.ReadFileIntoStr(gfps[i], &src); err != nil {
					break
				}
				if strings.HasPrefix(src, "package ") {
					pos = 0
				} else if pos = 1 + strings.Index(src, "\npackage "); pos == 0 {
					err = umisc.E(z.BadMsg("Go package source file", gfps[i]))
				}
				if posln := pos + strings.IndexRune(src[pos:], '\n'); err == nil {
					if oldpkgln := src[pos:posln]; oldpkgln != "package main" {
						src = src[:pos] + "package main" + src[posln:]
					}
					if goEvalExpr != "" {
						if pos = 1 + strings.Index(src, "\nfunc main() {"); pos > 0 {
							src = src[:pos] + z.Strf("func main%d() {", time.Now().UnixNano()) + src[pos+13:]
						}
						if iscursrc {
							if strings.HasPrefix(goEvalExpr, ":") {
								src += z.Strf("\n\nfunc main() { %s }", goEvalExpr[1:])
							} else {
								src += z.Strf("\n\nfunc main() { println(%s) }", goEvalExpr)
							}
						}
					}
					goRunArgs = append(goRunArgs, filepath.Base(gfps[i]))
					err = ufs.WriteTextFile(filepath.Join(goRunDir, goRunArgs[len(goRunArgs)-1]), src)
				}
			}
		}
	}
	return
}

func goRunEval(srcFilePath string, maybeSrcFull string, goEvalExpr string) (evalOutAndStdErr string, otherStdOut string, err error) {
	gorunargs, gorundir, e := goRunEvalPrepCmd(z.Prog.Dir.Cache, srcFilePath, maybeSrcFull, goEvalExpr)
	defer ufs.ClearDirectory(gorundir)
	if err = e; err == nil {
		otherStdOut, evalOutAndStdErr, err = urun.CmdExecIn(gorundir, "go", gorunargs...)
	}
	if err != nil {
		panic(err)
	}
	return
}
