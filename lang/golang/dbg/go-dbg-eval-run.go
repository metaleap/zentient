package zgodbg

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/zentient/dbg"
)

var strf = fmt.Sprintf

type Dbg struct {
	zdbg.Dbg
	printLn func(bool, string)
	replish struct {
		is           bool
		killed       bool
		stdout       io.Writer
		stderr       io.Writer
		stdin        io.Reader
		tmpDirPath   string
		srcFilePath  string
		maybeSrcFull string
	}
}

func (me *Dbg) Init(tmpDirPath string, srcFilePath string, maybeSrcFull string, printLn func(bool, string)) (err error) {
	me.printLn = printLn
	var gorunargs []string
	var hadmain bool
	var jumpFilePath string
start:
	if jumpFilePath != "" {
		srcFilePath = jumpFilePath
	}
	if gorunargs, me.Cmd.Dir, hadmain, err = GoRunEvalPrepCmd(tmpDirPath, srcFilePath, maybeSrcFull, ""); err == nil {
		if (!hadmain) && jumpFilePath == "" {
			for pkgdir := filepath.Dir(filepath.Dir(srcFilePath)); len(pkgdir) > 3; pkgdir = filepath.Dir(pkgdir) {
				ufs.WalkFilesIn(pkgdir, func(fullPath string) (keepWalking bool) {
					if keepWalking = true; strings.HasSuffix(fullPath, ".go") {
						if src := ufs.ReadTextFile(fullPath, false, ""); strings.HasPrefix(src, "package main\n") || strings.Index(src, "\npackage main\n") >= 2 {
							jumpFilePath, keepWalking = fullPath, false
						}
					}
					return
				})
				if jumpFilePath != "" {
					break
				}
			}
			if jumpFilePath != "" {
				goto start
			}
		}
		if me.replish.is = !hadmain; me.replish.is {
			me.replish.tmpDirPath, me.replish.srcFilePath, me.replish.maybeSrcFull = tmpDirPath, srcFilePath, maybeSrcFull
		} else {
			me.Cmd.Name = filepath.Base(os.Args[0]) + "-" + filepath.Base(me.Cmd.Dir)
			gobuildargs := append([]string{"build", "-o", me.Cmd.Name}, gorunargs[1:]...)
			if _, cmderr, e := urun.CmdExecStdin("", me.Cmd.Dir, "go", gobuildargs...); e != nil {
				err = e
			} else if cmderr != "" {
				err = umisc.E(cmderr)
			}
			me.Cmd.Name = filepath.Join(me.Cmd.Dir, me.Cmd.Name)
		}
	}
	return
}

func (me *Dbg) Dequeue() string {
	if me.replish.is {
		return "Unexpected, but OK here's some faux stdin!\n"
	}
	return me.Dbg.Dequeue()
}

func (me *Dbg) Enqueue(goEvalExpr string) {
	if !me.replish.is {
		me.Dbg.Enqueue(goEvalExpr)
		return
	}
	gorunargs, gorundir, _, err := GoRunEvalPrepCmd(me.replish.tmpDirPath, me.replish.srcFilePath, me.replish.maybeSrcFull, goEvalExpr)
	if me.Cmd.Dir = gorundir; err == nil {
		cmdout, cmderr, cmdname := "", "", filepath.Base(os.Args[0])+"-"+filepath.Base(gorundir)
		gobuildargs := append([]string{"build", "-o", cmdname}, gorunargs[1:]...)
		if cmdout, cmderr, err = urun.CmdExecStdin("", gorundir, "go", gobuildargs...); err == nil {
			if cmderr != "" {
				err = umisc.E(cmderr)
			} else if cmdout != "" {
				err = umisc.E(cmdout)
			}
		}
		if err == nil {
			cmdname = filepath.Join(gorundir, cmdname)
			cmd := exec.Command(cmdname)
			cmd.Dir, cmd.Stdout, cmd.Stdin, cmd.Stderr = gorundir, me.replish.stdout, me.replish.stdin, me.replish.stderr
			if err = cmd.Start(); err == nil {
				err = cmd.Wait()
			}
		}
	}
	if err != nil {
		me.printLn(true, err.Error())
	}
}

func (me *Dbg) Kill() error {
	if me.replish.is {
		me.replish.killed, me.replish.stdout, me.replish.stdin, me.replish.stderr = true, nil, nil, nil
		return nil
	}
	return me.Dbg.Kill()
}

func (me *Dbg) PrintLn(isErr bool, ln string) {
	me.printLn(isErr, ln)
}

func (me *Dbg) Start(stdout io.Writer, stdin io.Reader, stderr io.Writer) error {
	if me.replish.is {
		me.replish.stdout, me.replish.stdin, me.replish.stderr = stdout, stdin, stderr
		return nil
	}
	return me.Dbg.Start(stdout, stdin, stderr)
}

func (me *Dbg) Wait() error {
	if me.replish.is {
		for !me.replish.killed {
			time.Sleep(time.Millisecond * 123)
		}
		return nil
	}
	return me.Dbg.Wait()
}

func GoRunEvalPrepCmd(tmpDirPath string, srcFilePath string, maybeSrcFull string, goEvalExpr string) (goRunArgs []string, goRunDir string, hadMain bool, err error) {
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
			err = umisc.E("Bad Go package: " + pkgsrcdirpath)
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
					err = umisc.E("Bad Go package source file: " + gfps[i])
				}
				if posln := pos + strings.IndexRune(src[pos:], '\n'); err == nil {
					if oldpkgln := src[pos:posln]; oldpkgln != "package main" {
						src = src[:pos] + "package main" + src[posln:]
					}
					if pos = 1 + strings.Index(src, "\nfunc main() {"); pos > 0 {
						if hadMain = true; goEvalExpr != "" {
							src = src[:pos] + strf("func main%d() {", time.Now().UnixNano()) + src[pos+13:]
						}
					}
					if goEvalExpr != "" {
						if iscursrc {
							if strings.HasPrefix(goEvalExpr, ":") {
								src += strf("\n\nfunc main() { %s }", goEvalExpr[1:])
							} else {
								src += strf("\n\nfunc main() { println(%s) }", goEvalExpr)
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
