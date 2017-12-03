package zhs

import (
	"github.com/metaleap/zentient"
)

type hsSrcFormatting struct {
	z.SrcFormattingBase

	knownFormatters z.Tools
}

var srcFmt hsSrcFormatting

func init() {
	srcFmt.Impl = &srcFmt
	z.Lang.SrcFmt = &srcFmt
}

func (me *hsSrcFormatting) onPreInit() {
	srcFmt.knownFormatters = z.Tools{
		tools.hindent, tools.stylishhaskell, tools.brittany,
	}
}

func (*hsSrcFormatting) DoesStdoutWithFilePathArg(tool *z.Tool) bool {
	return tool != tools.hindent
}

func (me *hsSrcFormatting) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *hsSrcFormatting) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string, error) {
	if formatter != tools.brittany && formatter != tools.hindent && formatter != tools.stylishhaskell {
		return "", "", z.Errf("Invalid tool: %s" + formatter.Name)
	}

	var cmdargs []string
	if formatter == tools.hindent {
		cmdargs = append(cmdargs, "--no-force-newline", "--sort-imports", "--indent-size", "4")
	} else if formatter == tools.brittany {
		cmdargs = append(cmdargs, "--indent", "4")
	}
	if srcFilePath != "" {
		cmdargs = append(cmdargs, srcFilePath)
	}

	return formatter.Exec(cmdName, cmdargs, src)
}
