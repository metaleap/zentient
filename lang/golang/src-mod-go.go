package zgo

import (
	"github.com/metaleap/zentient"
)

type goSrcMod struct {
	z.SrcModBase

	knownFormatters z.Tools
}

var srcMod goSrcMod

func init() {
	srcMod.Impl = &srcMod
	z.Lang.SrcMod = &srcMod
}

func (me *goSrcMod) onPreInit() {
	me.knownFormatters = z.Tools{
		tools.gofmt, tools.goimports,
	}
}

func (me *goSrcMod) onPostInit() {
	if z.Prog.Cfg.FormatterName == "" && tools.gofmt.Installed {
		z.Prog.Cfg.FormatterName = "gofmt"
	}
}

func (me *goSrcMod) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *goSrcMod) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string, error) {
	if formatter != tools.gofmt && formatter != tools.goimports {
		return "", "", z.Errf("Invalid tool: %s" + formatter.Name)
	}

	var cmdargs []string
	if formatter == tools.gofmt {
		cmdargs = append(cmdargs, "-s")
	}
	if srcFilePath != "" {
		cmdargs = append(cmdargs, srcFilePath)
	}

	return formatter.Exec(cmdName, cmdargs, src)
}
