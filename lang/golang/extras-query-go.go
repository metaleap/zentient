package zgo

import (
	"path/filepath"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

func (me *goExtras) runQueryGoDoc(srcLens *z.SrcLens, arg string, resp *z.ExtrasResp) {
	if i1, i2 := ustr.Idx(arg, "."), ustr.Idx(arg, " "); i1 > 0 && (i2 < 0 || i2 > i1) {
		arg = arg[:i1] + " " + arg[i1+1:]
	}
	var cmd = ustr.Split(arg, " ")
	if udevgo.PkgsByImP != nil && ustr.IsLower(cmd[0][:1]) && udevgo.PkgsByImP[cmd[0]] == nil {
		for _, pkg := range udevgo.PkgsByImP {
			if pkg.Name == cmd[0] {
				cmd[0] = pkg.ImportPath
				break
			}
		}
	}
	cmdout, cmderr, err := urun.CmdExecIn(filepath.Dir(srcLens.FilePath), "go", append([]string{"doc"}, cmd...)...)
	if err != nil {
		panic(err)
	}
	resp.Warns = uslice.StrFiltered(uslice.StrMap(ustr.Split(cmderr, "\n"), ustr.Trim),
		func(s string) bool { return !ustr.Pref(s, "exit status ") })
	resp.InfoTips = append(resp.InfoTips, z.InfoTip{Value: ustr.Trim(cmdout)})
}
