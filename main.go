package main

import (
	"bufio"
	"runtime"

	"github.com/metaleap/go-util/run"

	"github.com/metaleap/zentient/go"
	"github.com/metaleap/zentient/hs"
	"github.com/metaleap/zentient/z"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := z.Init(map[string]func() z.Zengine{"go": zgo.Init, "hs": zhs.Init})
	if err != nil {
		panic(err)
	}

	var stdin *bufio.Scanner
	stdin, z.RawOut, z.Out = urun.SetupJsonProtoPipes(1024*1024*4, false, true)
	for stdin.Scan() {
		if err = z.HandleRequest(stdin.Text()); err == nil {
			err = z.RawOut.Flush()
		}
		if err != nil {
			z.Out.Encode(err.Error())
			err = z.RawOut.Flush()
			break
		}
	}
}
