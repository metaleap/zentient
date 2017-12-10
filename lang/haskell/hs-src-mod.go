package zhs

import (
	"github.com/metaleap/zentient"
)

var srcMod hsSrcMod

func init() {
	srcMod.Impl, z.Lang.SrcMod = &srcMod, &srcMod
}

type hsSrcMod struct {
	z.SrcModBase

	knownFormatters z.Tools
}

func (me *hsSrcMod) onPreInit() {
	srcMod.knownFormatters = z.Tools{
		tools.hindent, tools.stylishhaskell, tools.brittany,
	}
}

func (*hsSrcMod) DoesStdoutWithFilePathArg(tool *z.Tool) bool {
	return tool != tools.hindent
}

func (me *hsSrcMod) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *hsSrcMod) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string) {
	if formatter != tools.brittany && formatter != tools.hindent && formatter != tools.stylishhaskell {
		z.BadPanic("formatting tool", formatter.Name)
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
