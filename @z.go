package z

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/sys"
)

var (
	Strf = fmt.Sprintf
	Lang struct {
		Enabled   bool
		ID        string
		Title     string
		SrcMod    ISrcMod
		SrcIntel  ISrcIntel
		Diag      IDiag
		Extras    IExtras
		PkgIntel  IPkgIntel
		Caddies   []*Caddy
		Workspace IWorkspace
		Tooling   ITooling
	}
	Prog struct {
		Cfg Config

		name string
		dir  struct {
			cache  string
			config string
		}
		menus       []IMenuItems
		dispatchers []iDispatcher
		objSnappers []IObjSnap
		pipeIO      struct {
			mutex         sync.Mutex
			stdoutEncoder *json.Encoder
			stdoutWriter  *bufio.Writer
			stdinReadLn   *bufio.Scanner
		}
	}
)

func BadMsg(what string, which string) string {
	return Strf("%s: invalid %s %s '%s'", Prog.name, Lang.Title, what, which)
}

func BadPanic(what string, which string) {
	panic(BadMsg(what, which))
}

func Init() (err error) {
	Prog.name = os.Args[0]
	Prog.dir.config = filepath.Join(usys.UserDataDirPath(false), Prog.name)
	Prog.dir.cache = filepath.Join(usys.UserDataDirPath(true), Prog.name)
	if err = ufs.EnsureDirExists(Prog.dir.config); err != nil {
		return
	} else if err = ufs.EnsureDirExists(Prog.dir.cache); err != nil {
		return
	}

	if Prog.Cfg.reload(); Prog.Cfg.err == nil {
		wellknowndispatchers := []iDispatcher{
			&mainMenu{}, Lang.Workspace, Lang.SrcIntel, Lang.Diag, Lang.SrcMod, Lang.Extras, Lang.PkgIntel, Lang.Tooling,
		}
		for _, disp := range wellknowndispatchers {
			if disp != nil {
				Prog.dispatchers = append(Prog.dispatchers, disp)
				disp.Init()
				if objsnp, _ := disp.(IObjSnap); objsnp != nil {
					Prog.objSnappers = append(Prog.objSnappers, objsnp)
				}
			}
		}

		wellknownmenus := []IMenuItems{
			Lang.Diag, Lang.SrcMod, Lang.PkgIntel, Lang.Tooling,
		}
		for _, menu := range wellknownmenus {
			if menu != nil {
				Prog.menus = append(Prog.menus, menu)
			}
		}
	}
	for _, c := range Lang.Caddies {
		c.onInit()
		c.LangID = Lang.ID
	}
	return
}

func InitAndServe(onPreInit func(), onPostInit func()) (err error) {
	// note to self: don't ADD any further logic in here, do it only in either Init() or Serve()
	onPreInit()
	if err = Init(); err == nil {
		onPostInit()
		err = Serve()
	}
	return
}

func InitAndServeOrPanic(onPreInit func(), onPostInit func()) {
	if err := InitAndServe(onPreInit, onPostInit); err != nil {
		panic(err)
	}
}
