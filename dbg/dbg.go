package zdbg

import (
	"io"
	"os/exec"
	"sync"

	"github.com/go-leap/fs"
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

func (this *Dbg) Dispose() {
	if this.Cmd.Dir != "" {
		_ = ufs.Del(this.Cmd.Dir)
	}
}

func (this *Dbg) Dequeue() (cmdEvalExpr string) {
	this.Lock()
	defer this.Unlock()
	if len(this.cmdExprs) > 0 {
		cmdEvalExpr = this.cmdExprs[0]
		this.cmdExprs = this.cmdExprs[1:]
	}
	return
}

func (this *Dbg) Enqueue(cmdEvalExpr string) {
	this.Lock()
	defer this.Unlock()
	this.cmdExprs = append(this.cmdExprs, cmdEvalExpr)
}

func (this *Dbg) Kill() (err error) {
	if this.cmd != nil && this.cmd.Process != nil {
		err = this.cmd.Process.Kill()
	}
	this.cmd = nil
	return
}

func (this *Dbg) Start(stdout io.Writer, stdin io.Reader, stderr io.Writer) (err error) {
	_ = this.Kill()
	this.cmd = exec.Command(this.Cmd.Name, this.Cmd.Args...)
	this.cmd.Stdout, this.cmd.Stdin, this.cmd.Stderr = stdout, stdin, stderr
	if err = this.cmd.Start(); err != nil {
		_ = this.Kill()
	}
	return
}

func (this *Dbg) Wait() (err error) {
	if this.cmd != nil && this.cmd.Process != nil {
		_, err = this.cmd.Process.Wait() // cmd.Wait() hung forever in this specific scenario --- impairing 'restart' functionality
		this.cmd = nil
	}
	return
}
