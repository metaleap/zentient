package zgo

import (
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/zentient"
)

var srcMod goSrcMod

func init() {
	srcMod.Impl, z.Lang.SrcMod = &srcMod, &srcMod
}

type goSrcMod struct {
	z.SrcModBase

	knownFormatters z.Tools
}

func (me *goSrcMod) onPreInit() {
	me.knownFormatters = z.Tools{
		tools.gofmt, tools.goimports, tools.goreturns,
	}
}

func (me *goSrcMod) onPostInit() {
	if z.Prog.Cfg.FormatterName == "" && tools.gofmt.Installed {
		z.Prog.Cfg.FormatterName = "gofmt"
	}
}

func (*goSrcMod) DoesStdoutWithFilePathArg(*z.Tool) bool {
	return true
}

func (me *goSrcMod) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *goSrcMod) RunRenamer(srcLens *z.SrcLens, newName string) (srcMods z.SrcLenses) {
	if !tools.gorename.Installed {
		panic(tools.gorename.NotInstalledMessage())
	}
	eol, offset := "\n", srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	if srcLens.CrLf {
		eol = "\r\n"
	}
	fileedits, err := udevgo.Gorename(tools.gorename.Name, srcLens.FilePath, offset, newName, eol)
	if err != nil {
		panic(err)
	}
	for _, fedit := range fileedits {
		srcmod := z.SrcLens{Str: fedit.Msg, FilePath: fedit.Ref, Range: &z.SrcRange{}}
		srcmod.Range.Start.Col, srcmod.Range.Start.Ln = 1, fedit.Pos1Ln+1
		srcmod.Range.End.Col, srcmod.Range.End.Ln = 1, fedit.Pos2Ln+1
		srcMods = append(srcMods, &srcmod)
	}
	return
}

func (me *goSrcMod) RunFormatter(formatter *z.Tool, cmdName string, _ *z.SrcFormattingClientPrefs, srcFilePath string, src string) (string, string) {
	if formatter != tools.gofmt && formatter != tools.goimports && formatter != tools.goreturns {
		z.BadPanic("formatting tool", formatter.Name)
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
