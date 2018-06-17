package zhs

import (
	"github.com/metaleap/zentient"
)

var workspace hsWorkspace

func init() {
	workspace.Impl, z.Lang.Workspace = &workspace, &workspace
}

type hsWorkspace struct {
	z.WorkspaceBase
}
