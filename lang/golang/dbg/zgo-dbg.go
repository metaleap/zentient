package zgodbg

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-leap/fs"
	"github.com/go-leap/run"
	"github.com/metaleap/zentient/dbg"
)

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
				ufs.WalkFilesIn(pkgdir, func(fullPath string, _ os.FileInfo) (keepWalking bool) {
					if keepWalking = true; strings.HasSuffix(fullPath, ".go") {
						if src := ufs.ReadTextFileOr(fullPath, ""); strings.HasPrefix(src, "package main\n") || strings.Index(src, "\npackage main\n") >= 2 {
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
				err = errors.New(cmderr)
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
				err = errors.New(cmderr)
			} else if cmdout != "" {
				err = errors.New(cmdout)
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
