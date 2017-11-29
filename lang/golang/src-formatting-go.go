package zgo

import (
	"github.com/metaleap/go-util/run"
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

func (me *srcFormatting) RunFormatter(formatter *z.Tool, customProgName string, srcFilePath string, src string) (string, string, error) {
	if formatter != tools.gofmt && formatter != tools.goimports {
		return "", "", z.Errf("Invalid tool: %s" + formatter.Name)
	}

	cmdname := formatter.Name
	if customProgName != "" {
		cmdname = customProgName
	}

	var cmdargs []string
	if formatter == tools.gofmt {
		cmdargs = append(cmdargs, "-s")
	}
	if srcFilePath != "" {
		cmdargs = append(cmdargs, srcFilePath)
	}

	stdout, stderr, err := urun.CmdExecStdin(src, "", cmdname, cmdargs...)
	if stderr != "" {
		stderr = z.Strf("%s: %s", cmdname, stderr)
	}
	return stdout, stderr, err
}
