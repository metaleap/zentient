package zgo

import (
	"github.com/metaleap/zentient"
)

type goSrcFormatting struct {
	z.SrcFormattingBase

	knownFormatters z.Tools
}

var srcFmt goSrcFormatting

func init() {
	srcFmt.Impl = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *goSrcFormatting) onPreInit() {
	me.knownFormatters = z.Tools{
		tools.gofmt, tools.goimports,
	}
}

func (me *goSrcFormatting) onPostInit() {
	if z.Prog.Cfg.FormatterName == "" && tools.gofmt.Installed {
		z.Prog.Cfg.FormatterName = "gofmt"
	}
}

func (me *goSrcFormatting) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *goSrcFormatting) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string, error) {
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
