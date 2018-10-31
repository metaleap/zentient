//
package z

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-leap/fs"
	"github.com/go-leap/sys"
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
		Settings  ISettings
		Tooling   ITooling
		Workspace IWorkspace
		Pages     IPages
		sideViews sideViews
	}
	Prog struct {
		Cfg Config

		Name string
		Dir  struct {
			Cache  string
			Config string
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
		// recall struct {
		// 	i64 map[string]int64
		// }
	}
)

func BadMsg(what string, which string) string {
	return Strf("%s: invalid %s %s '%s'", Prog.Name, Lang.Title, what, which)
}

func BadPanic(what string, which string) {
	panic(BadMsg(what, which))
}

func PrettifyPathsIn(s string) string {
	if mod := false; strings.ContainsRune(s, filepath.Separator) {
		words := strings.Split(s, " ")
		for i := range words {
			if strings.ContainsRune(words[i], filepath.Separator) {
				if w := Lang.Workspace.PrettyPath(words[i]); w != words[i] {
					mod, words[i] = true, w
				}
			}
		}
		if mod {
			s = strings.Join(words, " ")
		}
	}
	return s
}

func SendNotificationMessageToClient(level DiagSeverity, message string) (err error) {
	ipcid := IPCID_NOTIFY_INFO
	if level == DIAG_SEV_ERR {
		ipcid = IPCID_NOTIFY_ERR
	} else if level == DIAG_SEV_WARN {
		ipcid = IPCID_NOTIFY_WARN
	}
	err = send(&ipcResp{IpcID: ipcid, Val: message})
	return
}

func ToolsMsgGone(missingToolName string) string {
	return "Not installed: " + missingToolName
}

func ToolsMsgMore(missingToolName string) string {
	return "for more information, see: Zentient Main Menu / Tooling / " + missingToolName
}

func ToolGonePanic(missingToolName string) {
	panic(Strf("%s — %s", ToolsMsgGone(missingToolName), ToolsMsgMore(missingToolName)))
}

func Init() (err error) {
	Prog.Name = os.Args[0]
	Prog.Dir.Config = filepath.Join(usys.UserDataDirPath(false), Prog.Name)
	Prog.Dir.Cache = filepath.Join(usys.UserDataDirPath(true), Prog.Name)
	if err = ufs.EnsureDir(Prog.Dir.Config); err != nil {
		return
	} else if err = ufs.EnsureDir(Prog.Dir.Cache); err != nil {
		return
	}

	if Prog.Cfg.reload(); Prog.Cfg.err == nil {
		// Prog.Cfg.recall()
		wellknowndispatchers := []iDispatcher{
			Lang.SrcIntel, Lang.Workspace, Lang.Diag, Lang.SrcMod, Lang.Extras, Lang.PkgIntel, Lang.Tooling, Lang.Settings, &mainMenu{}, Lang.Pages, &Lang.sideViews,
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
			Lang.SrcMod, Lang.Diag, Lang.PkgIntel, Lang.Settings, Lang.Tooling,
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
