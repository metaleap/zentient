package z

import (
	"github.com/metaleap/go-util/run"
)

type Tool struct {
	Name      string
	Installed bool
	Website   string
}

type Tools []*Tool

func (me Tools) ByName(name string) *Tool {
	if name != "" {
		for _, tool := range me {
			if tool.Name == name {
				return tool
			}
		}
	}
	return nil
}

func ExecTool(cmdname string, cmdargs []string, stdin string) (string, string, error) {
	stdout, stderr, err := urun.CmdExecStdin(stdin, "", cmdname, cmdargs...)
	if stderr != "" {
		stderr = Strf("%s: %s", cmdname, stderr)
	}
	return stdout, stderr, err
}
