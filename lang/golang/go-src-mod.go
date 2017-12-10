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
func (me *goSrcMod) KnownFormatters() z.Tools {
	return me.knownFormatters
}

func (me *goSrcMod) RunRenamer(srcLens *z.SrcLens, newName string) (srcMods []*z.SrcLens) {
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
		srcmod := z.SrcLens{SrcSel: fedit.Msg, FilePath: fedit.Ref, Range: &z.SrcRange{}}
		srcmod.Range.Start.Col, srcmod.Range.Start.Ln = 1, fedit.Pos1Ln+1
		srcmod.Range.End.Col, srcmod.Range.End.Ln = 1, fedit.Pos2Ln+1
		srcMods = append(srcMods, &srcmod)
	}
	return
}

func (me *goSrcMod) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string) {
	if formatter != tools.gofmt && formatter != tools.goimports && formatter != tools.goreturns {
		z.Bad("formatting tool", formatter.Name)
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