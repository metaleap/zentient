package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var tools struct {
	gofmt     *z.Tool
	goimports *z.Tool
	goreturns *z.Tool

	gogetdoc *z.Tool
	godef    *z.Tool

	gorename *z.Tool
}

func toolsInit() {
	t := &tools

	t.gofmt = &z.Tool{Name: "gofmt", Website: "http://golang.org/cmd/gofmt", Installed: udevgo.Has_gofmt}
	t.goimports = &z.Tool{Name: "goimports", Website: "http://golang.org/x/tools/cmd/goimports", Installed: udevgo.Has_goimports}
	t.goreturns = &z.Tool{Name: "goreturns", Website: "http://github.com/sqs/goreturns", Installed: udevgo.Has_goreturns}

	t.gogetdoc = &z.Tool{Name: "gogetdoc", Website: "http://github.com/zmb3/gogetdoc#readme", Installed: udevgo.Has_gogetdoc}
	t.godef = &z.Tool{Name: "godef", Website: "http://github.com/rogpeppe/godef#readme", Installed: udevgo.Has_godef}

	t.gorename = &z.Tool{Name: "gorename", Website: "http://golang.org/x/tools/cmd/gorename", Installed: udevgo.Has_gorename}
}
