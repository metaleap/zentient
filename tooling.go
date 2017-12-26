package z

import (
	"strings"

	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
)

type ToolCats uint8

const (
	_ ToolCats = iota
	TOOLS_CAT_MOD_REN
	TOOLS_CAT_MOD_FMT
	TOOLS_CAT_INTEL_TIPS
	TOOLS_CAT_INTEL_SYMS
	TOOLS_CAT_INTEL_HIGH
	TOOLS_CAT_INTEL_CMPL
	TOOLS_CAT_INTEL_NAV
	TOOLS_CAT_EXTRAS_QUERY
	TOOLS_CAT_DIAGS
)

func (me ToolCats) String() string {
	switch me {
	case TOOLS_CAT_MOD_FMT:
		return "Formatting"
	case TOOLS_CAT_MOD_REN:
		return "Symbol Renaming"
	case TOOLS_CAT_INTEL_TIPS:
		return "Info Tips"
	case TOOLS_CAT_INTEL_SYMS:
		return "Symbol Search"
	case TOOLS_CAT_INTEL_HIGH:
		return "Contextual Highlighting"
	case TOOLS_CAT_INTEL_CMPL:
		return "Completion Suggest"
	case TOOLS_CAT_INTEL_NAV:
		return "Lookups: References/Definition/Type/Interface/Implementations"
	case TOOLS_CAT_DIAGS:
		return "Linting / Diagnostics"
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

func (me *ToolingBase) onListAllTools() (menu MenuItems) {
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

func (me *ToolingBase) MenuItems(srcLens *SrcLens) (menu MenuItems) {
	menu = append(menu, me.cmdListAll)
	return
}

func (me *ToolingBase) CountNumInst(all Tools) (numInst int) {
	for _, t := range all {
		if t.Installed {
			numInst++
		}
	}
	return
}

func (me *ToolingBase) KnownToolsFor(cats ...ToolCats) (tools Tools) {
	alltools := me.Impl.KnownTools()
	if tools = alltools; len(cats) > 0 {
		tools = Tools{}
		added := false
		for _, t := range alltools {
			for _, tc := range t.Cats {
				for _, c := range cats {
					if added = (tc == c); added {
						tools = append(tools, t)
						break
					}
				}
				if added {
					break
				}
			}
		}
	}
	return
}

type Tools []*Tool

func (me Tools) Has(name string) bool {
	for _, t := range me {
		if t.Name == name {
			return true
		}
	}
	return false
}

func (me Tools) Len(inst bool) (num int) {
	for _, t := range me {
		if t.Installed == inst {
			num++
		}
	}
	return
}

// func (me Tools) Len() int          { return len(me) }
// func (me Tools) Swap(i int, j int) { me[i], me[j] = me[j], me[i] }
// func (me Tools) Less(i1 int, i2 int) bool {
// 	one, two := me[i1], me[i2]
// 	if l := len(one.Cats); l != len(two.Cats) {
// 		return l < len(two.Cats)
// 	} else {
// 		for i := 0; i < l; i++ {
// 			if one.Cats[i] != two.Cats[i] {
// 				return one.Cats[i] < two.Cats[i]
// 			}
// 		}
// 	}
// 	return one.Name < two.Name
// }

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
	DiagSev   DiagSeverity
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

func (me *Tool) IsInAutoDiags() bool {
	return uslice.StrHas(Prog.Cfg.AutoDiags, me.Name)
}

func (me *Tool) NotInstalledMessage() string {
	return Strf("Not installed: `%s`, how-to at: %s", me.Name, me.Website)
}

func (me *Tool) ToggleInAutoDiags() error {
	if me.IsInAutoDiags() {
		Prog.Cfg.AutoDiags = uslice.StrWithout(Prog.Cfg.AutoDiags, false, me.Name)
	} else {
		Prog.Cfg.AutoDiags = append(Prog.Cfg.AutoDiags, me.Name)
	}
	return Prog.Cfg.Save()
}
