package zdbg

import (
	"io"
	"os/exec"
	"sync"

	"github.com/metaleap/go-util/fs"
)

type IDbg interface {
	Dispose()
	Dequeue() string
	Enqueue(string)
	Init(string, string, string, func(bool, string)) error
	Kill() error
	PrintLn(bool, string)
	Start(io.Writer, io.Reader, io.Writer) error
	Wait() error
}

type Dbg struct {
	sync.Mutex

	Cmd struct {
		Dir  string
		Name string
		Args []string
	}

	cmd      *exec.Cmd
	cmdExprs []string
}

func (me *Dbg) Dispose() {
	if me.Cmd.Dir != "" {
		_ = ufs.ClearDirectory(me.Cmd.Dir)
	}
}

func (me *Dbg) Dequeue() (cmdEvalExpr string) {
	me.Lock()
	defer me.Unlock()
	if len(me.cmdExprs) > 0 {
		cmdEvalExpr = me.cmdExprs[0]
		me.cmdExprs = me.cmdExprs[1:]
	}
	return
}

func (me *Dbg) Enqueue(cmdEvalExpr string) {
	me.Lock()
	defer me.Unlock()
	me.cmdExprs = append(me.cmdExprs, cmdEvalExpr)
}

func (me *Dbg) Kill() (err error) {
	if me.cmd != nil && me.cmd.Process != nil {
		err = me.cmd.Process.Kill()
	}
	me.cmd = nil
	return
}

func (me *Dbg) Start(stdout io.Writer, stdin io.Reader, stderr io.Writer) (err error) {
	_ = me.Kill()
	me.cmd = exec.Command(me.Cmd.Name, me.Cmd.Args...)
	me.cmd.Stdout, me.cmd.Stdin, me.cmd.Stderr = stdout, stdin, stderr
	if err = me.cmd.Start(); err != nil {
		_ = me.Kill()
	}
	return
}

func (me *Dbg) Wait() (err error) {
	if me.cmd != nil && me.cmd.Process != nil {
		_, err = me.cmd.Process.Wait() // cmd.Wait() hung forever in this specific scenario --- impairing 'restart' functionality
		me.cmd = nil
	}
	return
}
