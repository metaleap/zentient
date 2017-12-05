package zgo

import (
	"github.com/metaleap/go-util/dev/go"
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

func (me *goSrcMod) RunRenamer(srcLens *z.SrcLens, newName string) (mods []*z.SrcLens) {
	if !tools.gorename.Installed {
		panic(tools.gorename.NotInstalledMessage())
	}
	offset := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	eol := "\n"
	if srcLens.CrLf {
		eol = "\r\n"
	}
	fileedits, err := udevgo.Gorename(tools.gorename.Name, srcLens.FilePath, offset, newName, eol)
	if err != nil {
		panic(err)
	}
	for _, fedit := range fileedits {
		// fedit.Ref fedit.Msg fedit.Pos1Ln fedit.Pos2Ln
		var mod z.SrcLens
		mod.SrcSel, mod.Range, mod.FilePath = fedit.Msg, &z.SrcRange{}, fedit.Ref
		mod.Range.Start.Col, mod.Range.Start.Ln = 1, fedit.Pos1Ln+1
		mod.Range.End.Col, mod.Range.End.Ln = 1, fedit.Pos2Ln+1
		mods = append(mods, &mod)
	}
	return
}

func (me *goSrcMod) RunFormatter(formatter *z.Tool, cmdName string, srcFilePath string, src string) (string, string) {
	if formatter != tools.gofmt && formatter != tools.goimports {
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
