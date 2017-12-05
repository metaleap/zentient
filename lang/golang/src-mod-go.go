package zgo

import (
	"encoding/json"

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
		data, _ := json.Marshal(&fedit)
		println(string(data))
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
