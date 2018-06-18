package z

import (
	"github.com/go-leap/run"
	"github.com/go-leap/str"
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

func (this ToolCats) String() string {
	switch this {
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
	return Strf("%d", this)
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

func (this *ToolingBase) Init() {
	this.cmdListAll = &MenuItem{
		IpcID: IPCID_MENUS_TOOLS,
		Title: Strf("Known %s Tools", Lang.Title),
		Desc:  Strf("All currently supported %s tools and their installation info", Lang.Title),
		Hint:  Strf("(%d of %d installed)", this.Impl.NumInst(), this.Impl.NumTotal()),
	}
}

func (this *ToolingBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_MENUS_TOOLS:
		resp.Menu = &menuResp{SubMenu: &Menu{Desc: this.cmdListAll.Desc, Items: this.onListAllTools()}}
	default:
		return false
	}
	return true
}

func (this *ToolingBase) onListAllTools() (menu MenuItems) {
	all := this.Impl.KnownTools()
	for _, t := range all {
		item := MenuItem{Title: t.Name, Desc: "➜ " + t.Website, Hint: "Installed", IpcArgs: t.Website}
		if !t.Installed {
			item.Hint = "Not Installed"
		}
		if cats := []string{}; len(t.Cats) > 0 {
			for _, c := range t.Cats {
				cats = append(cats, c.String())
			}
			item.Hint += "  ·  Used for: " + ustr.Join(cats, ", ")
		}
		menu = append(menu, &item)
	}
	return
}

func (this *ToolingBase) MenuCategory() string {
	return "Tooling"
}

func (this *ToolingBase) menuItems(srcLens *SrcLens) (menu MenuItems) {
	menu = append(menu, this.cmdListAll)
	return
}

func (this *ToolingBase) CountNumInst(all Tools) (numInst int) {
	for _, t := range all {
		if t.Installed {
			numInst++
		}
	}
	return
}

func (this *ToolingBase) KnownToolsFor(cats ...ToolCats) (tools Tools) {
	alltools := this.Impl.KnownTools()
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

func (this Tools) byName(name string) *Tool {
	if name != "" {
		for _, tool := range this {
			if tool.Name == name {
				return tool
			}
		}
	}
	return nil
}

func (this Tools) has(name string) bool {
	for _, t := range this {
		if t.Name == name {
			return true
		}
	}
	return false
}

func (this Tools) instOnly() (inst Tools) {
	inst = make(Tools, 0, len(this))
	for _, t := range this {
		if t.Installed {
			inst = append(inst, t)
		}
	}
	return
}

func (this Tools) len(inst bool) (num int) {
	for _, t := range this {
		if t.Installed == inst {
			num++
		}
	}
	return
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

func (this *Tool) isInAutoDiags() bool {
	return ustr.In(this.Name, Prog.Cfg.AutoDiags...)
}

func (this *Tool) NotInstalledMessage() string {
	return Strf("Not installed: `%s`, how-to at: %s", this.Name, this.Website)
}

func (this *Tool) toggleInAutoDiags() error {
	if this.isInAutoDiags() {
		Prog.Cfg.AutoDiags = ustr.Sans(Prog.Cfg.AutoDiags, this.Name)
	} else {
		Prog.Cfg.AutoDiags = append(Prog.Cfg.AutoDiags, this.Name)
	}
	return Prog.Cfg.Save()
}
