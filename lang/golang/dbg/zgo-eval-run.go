package zgodbg

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-leap/dev/go"
	"github.com/go-leap/fs"
	"github.com/go-leap/run"
)

func GoRunEval(userAppCacheDirPath string, srcFilePath string, maybeSrcFull string, goEvalExpr string) (evalOutAndStdErr string, otherStdOut string, err error) {
	gorunargs, gorundir, _, e := GoRunEvalPrepCmd(userAppCacheDirPath, srcFilePath, maybeSrcFull, goEvalExpr)
	defer ufs.ClearDir(gorundir)
	if err = e; err == nil {
		otherStdOut, evalOutAndStdErr, err = urun.CmdExecIn(gorundir, "go", gorunargs...)
	}
	return
}

func GoRunEvalPrepCmd(tmpDirPath string, srcFilePath string, maybeSrcFull string, goEvalExpr string) (goRunArgs []string, goRunDir string, hadMain bool, err error) {
	pkgsrcdirpath := filepath.Dir(srcFilePath)
	goRunDir = filepath.Join(tmpDirPath, pkgsrcdirpath)
	if ufs.IsDir(goRunDir) {
		err = ufs.ClearDir(goRunDir)
	} else {
		err = ufs.EnsureDir(goRunDir)
	}
	if err == nil && udevgo.PkgsByDir == nil {
		err = udevgo.RefreshPkgs()
	}
	if goRunArgs = []string{"run"}; err == nil {
		if pkg := udevgo.PkgsByDir[pkgsrcdirpath]; pkg == nil {
			err = errors.New("Bad Go package: " + pkgsrcdirpath)
		} else {
			for i, pos, src, gfps := 0, 0, "", pkg.GoFilePaths(false); (err == nil) && (i < len(gfps)); i++ {
				iscursrc := gfps[i] == srcFilePath
				if iscursrc && len(maybeSrcFull) > 0 {
					src = maybeSrcFull
				} else if src, err = ufs.ReadTextFile(gfps[i]); err != nil {
					break
				}
				if strings.HasPrefix(src, "package ") {
					pos = 0
				} else if pos = 1 + strings.Index(src, "\npackage "); pos == 0 {
					err = errors.New("Bad Go package source file: " + gfps[i])
				}
				if posln := pos + strings.IndexRune(src[pos:], '\n'); err == nil {
					if oldpkgln := src[pos:posln]; oldpkgln != "package main" {
						src = src[:pos] + "package main" + src[posln:]
					}
					if pos = 1 + strings.Index(src, "\nfunc main() {"); pos > 0 {
						if hadMain = true; goEvalExpr != "" {
							src = src[:pos] + fmt.Sprintf("func main%d() {", time.Now().UnixNano()) + src[pos+13:]
						}
					}
					if goEvalExpr != "" {
						if iscursrc {
							if strings.HasPrefix(goEvalExpr, ":") {
								src += fmt.Sprintf("\n\nfunc main() { %s }", goEvalExpr[1:])
							} else {
								src += fmt.Sprintf("\n\nfunc main() { println(%s) }", goEvalExpr)
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
