package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var tools struct {
	gofmt     *z.Tool
	goimports *z.Tool
}

func toolsInit() {
	t := &tools

	t.gofmt = &z.Tool{Name: "gofmt", Website: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt}
	t.goimports = &z.Tool{Name: "goimports", Website: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports}
}
