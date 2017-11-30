package zgo

import (
	"github.com/metaleap/zentient"
)

type srcFormatting struct {
	z.SrcFormattingBase

	knownFormatters z.Tools
}

var (
	srcFmt srcFormatting
)

func init() {
	srcFmt.Self = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *srcFormatting) onPreInit() {
	me.knownFormatters = z.Tools{
		tools.gofmt, tools.goimports,
	}
}

func (me *srcFormatting) onPostInit() {
	if z.Prog.Cfg.FormatterName == "" && tools.gofmt.Installed {
		z.Prog.Cfg.FormatterName = "gofmt"
	}
}

func (me *srcFormatting) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *srcFormatting) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string, error) {
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

	return z.ExecTool(cmdName, cmdargs, src)
}
