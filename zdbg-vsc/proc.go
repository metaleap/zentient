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
	if n>0 {
		cmdeval := cmdEval
		if n<len(cmdEval) {
			cmdEval = cmdEval[n:]
		} else {
			cmdEval = []byte {}
		}
		for i:= 0; i<n && i<len(cmdeval); i++ {
			p[i] = cmdeval[i]
		}
	}
	return
}

var (
	cmd *exec.Cmd
	cmdEval []byte = []byte("Test Dis Shite\nWut Da Heck\n")
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
	err = cmd.Start()
	go listenProc()
	return
}

func listenProc () {
	for cmd!=nil && cmd.ProcessState!=nil && !cmd.ProcessState.Exited() {
	}
	onServerEvt_Terminated()
}
