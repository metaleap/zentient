package z

import (
	"strings"

	"github.com/go-leap/str"
	"github.com/metaleap/go-util/run"
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
	TOOLS_CAT_RUNONSAVE
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
	case TOOLS_CAT_RUNONSAVE:
		return "Run-on-Save"
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
		Title: Strf("Known %s Tools", Lang.Title),
		Desc:  Strf("All currently supported %s tools and their installation info", Lang.Title),
		Hint:  Strf("(%d of %d installed)", me.Impl.NumInst(), me.Impl.NumTotal()),
	}
}

func (me *ToolingBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_TOOLS:
		resp.Menu = &menuResp{SubMenu: &Menu{Desc: me.cmdListAll.Desc, Items: me.onListAllTools()}}
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

func (me *ToolingBase) menuItems(srcLens *SrcLens) (menu MenuItems) {
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
		for _, t := range alltools {
		__:
			for _, tc := range t.Cats {
				for _, c := range cats {
					if tc == c {
						tools = append(tools, t)
						break __
					}
				}
			}
		}
	}
	return
}

type Tools []*Tool

func (me Tools) has(name string) bool {
	for _, t := range me {
		if t.Name == name {
			return true
		}
	}
	return false
}

func (me Tools) len(inst bool) (num int) {
	for _, t := range me {
		if t.Installed == inst {
			num++
		}
	}
	return
}

func (me Tools) byName(name string) *Tool {
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

func (*Tool) Exec(panicOnErr bool, stdin string, cmdName string, cmdArgs []string) (string, string) {
	stdout, stderr, err := urun.CmdExecStdin(stdin, "", cmdName, cmdArgs...)
	if err != nil && panicOnErr {
		panic(err)
	}
	if stderr != "" {
		stderr = Strf("%s: %s", cmdName, stderr)
	}
	return stdout, stderr
}

func (me *Tool) isInAutoDiags() bool {
	return ustr.In(me.Name, Prog.Cfg.AutoDiags...)
}

func (me *Tool) NotInstalledMessage() string {
	return Strf("Not installed: `%s`, how-to at: %s", me.Name, me.Website)
}

func (me *Tool) toggleInAutoDiags() error {
	if me.isInAutoDiags() {
		Prog.Cfg.AutoDiags = ustr.Without(Prog.Cfg.AutoDiags, me.Name)
	} else {
		Prog.Cfg.AutoDiags = append(Prog.Cfg.AutoDiags, me.Name)
	}
	return Prog.Cfg.Save()
}
