package zdbg

import (
	"io"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

type IDbg interface {
	Dequeue() string
	Enqueue(string)
	Init(string, string, string) error
	Kill() error
	Start(io.Writer, io.Reader, io.Writer) error
	Wait() error
}

type Dbg struct {
	Cmd struct {
		Dir  string
		Name string
		Args []string
	}

	sync.Mutex

	cmd      *exec.Cmd
	cmdExprs []string
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
		me.cmd.Process.Signal(syscall.SIGQUIT)
		time.Sleep(time.Second)
		err = me.cmd.Process.Kill()
	}
	me.cmd = nil
	return
}

func (me *Dbg) Start(stdout io.Writer, stdin io.Reader, stderr io.Writer) (err error) {
	_ = me.Kill()
	me.cmd = exec.Command(me.Cmd.Name, me.Cmd.Args...)
	me.cmd.Dir, me.cmd.Stdout, me.cmd.Stdin, me.cmd.Stderr = me.Cmd.Dir, stdout, stdin, stderr
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
