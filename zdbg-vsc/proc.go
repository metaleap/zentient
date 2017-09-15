package main
import (
	"os/exec"

	"github.com/metaleap/zentient/zdbg-vsc/proto"
)

type ProcInOut struct {
	outcat string
}

func (me *ProcInOut) Write(p []byte) (n int, err error) {
	n = len(p)
	onServerEvt_Output(me.outcat, string(p))
	return
}

func (me *ProcInOut) Read(p []byte) (n int, err error) {
	if len(cmdExprs)>0 {
		expr := []byte(cmdExprs[0] + "\n")
		onServerEvt_Output("stderr", "HM_" + cmdExprs[0] + "_OK")
		if cmdExprs,n = cmdExprs[1:] , len(expr)  ;  n>0 {
			for i := 0 ; i<n ; i++ {  p[i] = expr[i]  }
		}
	}
	return
}

var (
	cmd *exec.Cmd
	cmdExprs []string = []string {}
)

func launchProc (req *zdbgvscp.LaunchRequest) (err error) {
	/*
		{"command":"launch","arguments":{"type":"zdbg","name":"⟨ℤ⟩","request":"launch","w":"/home/rox/c/go/src/github.com/metaleap",
			"c":"/home/rox/c/go/src/github.com/metaleap", "f":"/home/rox/c/go/src/github.com/metaleap/go-util-fs/fs.go", "s":" "},
		"type":"request",
		"seq":3}
	*/
	cmd = exec.Command("go-stdinoutdummy")
	cmd.Stdout = &ProcInOut { outcat: "stdout" }
	cmd.Stderr = &ProcInOut { outcat: "stderr" }
	cmd.Stdin = &ProcInOut {}
	if err = cmd.Start()  ;  err==nil {
		go listenProc()
	}
	return
}

func listenProc () {
	if cmd!=nil {
		if err := cmd.Wait()  ;  err!=nil {
			onServerEvt_Output("stderr", "ERR:" + err.Error())
		}
		onServerEvt_Terminated()
	}
}

func terminateProc () (err error) {
	if cmd!=nil && cmd.Process!=nil {
		err = cmd.Process.Kill()
	}
	cmd = nil
	return
}
