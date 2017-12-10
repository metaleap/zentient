package zgo

import (
	"github.com/metaleap/zentient"
)

var workspace goWorkspace

func init() {
	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
}

type goWorkspace struct {
	z.WorkspaceBase
}
