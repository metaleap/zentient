package z

import (
	"sort"
	"strings"

	"github.com/metaleap/go-util/run"
)

type ToolCats uint8

const (
	_ ToolCats = iota
	TOOLS_CAT_MOD_REN
	TOOLS_CAT_MOD_FMT
	TOOLS_CAT_INTEL_TIPS
	TOOLS_CAT_EXTRAS_QUERY
	TOOLS_CAT_INTEL_DIAG
)

func (me ToolCats) String() string {
	switch me {
	case TOOLS_CAT_MOD_FMT:
		return "Formatting"
	case TOOLS_CAT_MOD_REN:
		return "Symbol Renaming"
	case TOOLS_CAT_INTEL_TIPS:
		return "Info Tips"
	case TOOLS_CAT_INTEL_DIAG:
		return "Diagnostics"
	case TOOLS_CAT_EXTRAS_QUERY:
		return "CodeQuery"
	}
	return Strf("%d", me)
}

type ITooling interface {
	IMenuItems

	KnownTools() Tools
	NumInst() int
	NumTotal() int
}

type ToolingBase struct {
	Impl ITooling

	cmdListAll *MenuItem
}

func (me *ToolingBase) Init() {
	me.cmdListAll = &MenuItem{
		IpcID: IPCID_MENUS_TOOLS,
		Title: Strf("List Known & Supported %s Tools", Lang.Title),
		Desc:  Strf("All currently supported %s tools and their installation info", Lang.Title),
		Hint:  Strf("(%d of %d installed)", me.Impl.NumInst(), me.Impl.NumTotal()),
	}
}

func (me *ToolingBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_TOOLS:
		resp.Menu = &MenuResp{SubMenu: &Menu{Desc: me.cmdListAll.Desc, Items: me.onListAllTools()}}
	default:
		return false
	}
	return true
}

func (me *ToolingBase) onListAllTools() (menu []*MenuItem) {
	all := me.Impl.KnownTools()
	for _, t := range all {
		item := MenuItem{Title: t.Name, Desc: "➜ " + t.Website, Hint: "Installed", IpcArgs: t.Website}
		if !t.Installed {
			item.Hint = "Not Installed"
		}
		if cats := []string{}; len(t.Cats) > 0 {
			for _, c := range t.Cats {
				cats = append(cats, c.String())
			}
			item.Hint += "  ·  Used for: " + strings.Join(cats, ", ")
		}
		menu = append(menu, &item)
	}
	return
}

func (me *ToolingBase) MenuCategory() string {
	return "Tooling"
}

func (me *ToolingBase) MenuItems(srcLens *SrcLens) (menu []*MenuItem) {
	menu = append(menu, me.cmdListAll)
	return
}

func (me *ToolingBase) SortAndCountNumInst(all Tools) (numInst int) {
	sort.Sort(all)
	for _, t := range all {
		if t.Installed {
			numInst++
		}
	}
	return
}

type Tools []*Tool

func (me Tools) Len() int          { return len(me) }
func (me Tools) Swap(i int, j int) { me[i], me[j] = me[j], me[i] }
func (me Tools) Less(i1 int, i2 int) bool {
	one, two := me[i1], me[i2]
	if l := len(one.Cats); l != len(two.Cats) {
		return l < len(two.Cats)
	} else {
		for i := 0; i < l; i++ {
			if one.Cats[i] != two.Cats[i] {
				return one.Cats[i] < two.Cats[i]
			}
		}
	}
	return one.Name < two.Name
}

func (me Tools) ByName(name string) *Tool {
	if name != "" {
		for _, tool := range me {
			if tool.Name == name {
				return tool
			}
		}
	}
	return nil
}

type Tool struct {
	Cats      []ToolCats
	Name      string
	Installed bool
	Website   string
}

func (*Tool) Exec(cmdname string, cmdargs []string, stdin string) (string, string) {
	stdout, stderr, err := urun.CmdExecStdin(stdin, "", cmdname, cmdargs...)
	if err != nil {
		panic(err)
	}
	if stderr != "" {
		stderr = Strf("%s: %s", cmdname, stderr)
	}
	return stdout, stderr
}

func (me *Tool) NotInstalledMessage() string {
	return Strf("Not installed: `%s`, how-to at: %s", me.Name, me.Website)
}
