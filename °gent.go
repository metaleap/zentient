package z

// DO NOT EDIT: code generated with zentient-dbg-vsc-go-demo-go-gent using github.com/metaleap/go-gent

import (
	pkg__strconv "strconv"
)

// Valid returns whether the value of this `CaddyStatus` is between `CADDY_PENDING` (inclusive) and `CADDY_GOOD` (inclusive).
func (this CaddyStatus) Valid() (ret bool) {
	ret = ((this >= CADDY_PENDING) && (this <= CADDY_GOOD))
	return
}

// IsCADDY_PENDING returns whether the value of this `CaddyStatus` equals `CADDY_PENDING`.
func (this CaddyStatus) IsCADDY_PENDING() (ret bool) { ret = (this == CADDY_PENDING); return }

// IsCADDY_ERROR returns whether the value of this `CaddyStatus` equals `CADDY_ERROR`.
func (this CaddyStatus) IsCADDY_ERROR() (ret bool) { ret = (this == CADDY_ERROR); return }

// IsCADDY_BUSY returns whether the value of this `CaddyStatus` equals `CADDY_BUSY`.
func (this CaddyStatus) IsCADDY_BUSY() (ret bool) { ret = (this == CADDY_BUSY); return }

// IsCADDY_GOOD returns whether the value of this `CaddyStatus` equals `CADDY_GOOD`.
func (this CaddyStatus) IsCADDY_GOOD() (ret bool) { ret = (this == CADDY_GOOD); return }

func (this CaddyStatus) String() (ret string) {
	switch this {
	case CADDY_PENDING:
		ret = "CADDY_PENDING"
	case CADDY_ERROR:
		ret = "CADDY_ERROR"
	case CADDY_BUSY:
		ret = "CADDY_BUSY"
	case CADDY_GOOD:
		ret = "CADDY_GOOD"
	default:
		ret = pkg__strconv.FormatUint(uint64(this), 10)
	}
	return
}

func CaddyStatusFromString(s string) (this CaddyStatus, err error) { return }

// Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN` (inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).
func (this IpcIDs) Valid() (ret bool) {
	ret = ((this >= IPCID_MENUS_MAIN) && (this <= IPCID_EXTRAS_QUERY_RUN))
	return
}

// IsIPCID_MENUS_MAIN returns whether the value of this `IpcIDs` equals `IPCID_MENUS_MAIN`.
func (this IpcIDs) IsIPCID_MENUS_MAIN() (ret bool) { ret = (this == IPCID_MENUS_MAIN); return }

// IsIPCID_MENUS_PKGS returns whether the value of this `IpcIDs` equals `IPCID_MENUS_PKGS`.
func (this IpcIDs) IsIPCID_MENUS_PKGS() (ret bool) { ret = (this == IPCID_MENUS_PKGS); return }

// IsIPCID_MENUS_TOOLS returns whether the value of this `IpcIDs` equals `IPCID_MENUS_TOOLS`.
func (this IpcIDs) IsIPCID_MENUS_TOOLS() (ret bool) { ret = (this == IPCID_MENUS_TOOLS); return }

// IsIPCID_OBJ_SNAPSHOT returns whether the value of this `IpcIDs` equals `IPCID_OBJ_SNAPSHOT`.
func (this IpcIDs) IsIPCID_OBJ_SNAPSHOT() (ret bool) { ret = (this == IPCID_OBJ_SNAPSHOT); return }

// IsIPCID_PAGE_HTML returns whether the value of this `IpcIDs` equals `IPCID_PAGE_HTML`.
func (this IpcIDs) IsIPCID_PAGE_HTML() (ret bool) { ret = (this == IPCID_PAGE_HTML); return }

// IsIPCID_TREEVIEW_GETITEM returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_GETITEM`.
func (this IpcIDs) IsIPCID_TREEVIEW_GETITEM() (ret bool) {
	ret = (this == IPCID_TREEVIEW_GETITEM)
	return
}

// IsIPCID_TREEVIEW_CHILDREN returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_CHILDREN`.
func (this IpcIDs) IsIPCID_TREEVIEW_CHILDREN() (ret bool) {
	ret = (this == IPCID_TREEVIEW_CHILDREN)
	return
}

// IsIPCID_TREEVIEW_CHANGED returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_CHANGED`.
func (this IpcIDs) IsIPCID_TREEVIEW_CHANGED() (ret bool) {
	ret = (this == IPCID_TREEVIEW_CHANGED)
	return
}

// IsIPCID_CFG_RESETALL returns whether the value of this `IpcIDs` equals `IPCID_CFG_RESETALL`.
func (this IpcIDs) IsIPCID_CFG_RESETALL() (ret bool) { ret = (this == IPCID_CFG_RESETALL); return }

// IsIPCID_CFG_LIST returns whether the value of this `IpcIDs` equals `IPCID_CFG_LIST`.
func (this IpcIDs) IsIPCID_CFG_LIST() (ret bool) { ret = (this == IPCID_CFG_LIST); return }

// IsIPCID_CFG_SET returns whether the value of this `IpcIDs` equals `IPCID_CFG_SET`.
func (this IpcIDs) IsIPCID_CFG_SET() (ret bool) { ret = (this == IPCID_CFG_SET); return }

// IsIPCID_NOTIFY_INFO returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_INFO`.
func (this IpcIDs) IsIPCID_NOTIFY_INFO() (ret bool) { ret = (this == IPCID_NOTIFY_INFO); return }

// IsIPCID_NOTIFY_WARN returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_WARN`.
func (this IpcIDs) IsIPCID_NOTIFY_WARN() (ret bool) { ret = (this == IPCID_NOTIFY_WARN); return }

// IsIPCID_NOTIFY_ERR returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_ERR`.
func (this IpcIDs) IsIPCID_NOTIFY_ERR() (ret bool) { ret = (this == IPCID_NOTIFY_ERR); return }

// IsIPCID_PROJ_CHANGED returns whether the value of this `IpcIDs` equals `IPCID_PROJ_CHANGED`.
func (this IpcIDs) IsIPCID_PROJ_CHANGED() (ret bool) { ret = (this == IPCID_PROJ_CHANGED); return }

// IsIPCID_PROJ_POLLEVTS returns whether the value of this `IpcIDs` equals `IPCID_PROJ_POLLEVTS`.
func (this IpcIDs) IsIPCID_PROJ_POLLEVTS() (ret bool) { ret = (this == IPCID_PROJ_POLLEVTS); return }

// IsIPCID_SRCDIAG_LIST returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_LIST`.
func (this IpcIDs) IsIPCID_SRCDIAG_LIST() (ret bool) { ret = (this == IPCID_SRCDIAG_LIST); return }

// IsIPCID_SRCDIAG_RUN_CURFILE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_CURFILE`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_CURFILE() (ret bool) {
	ret = (this == IPCID_SRCDIAG_RUN_CURFILE)
	return
}

// IsIPCID_SRCDIAG_RUN_OPENFILES returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_OPENFILES`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_OPENFILES() (ret bool) {
	ret = (this == IPCID_SRCDIAG_RUN_OPENFILES)
	return
}

// IsIPCID_SRCDIAG_RUN_ALLFILES returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_ALLFILES`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_ALLFILES() (ret bool) {
	ret = (this == IPCID_SRCDIAG_RUN_ALLFILES)
	return
}

// IsIPCID_SRCDIAG_FORGETALL returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_FORGETALL`.
func (this IpcIDs) IsIPCID_SRCDIAG_FORGETALL() (ret bool) {
	ret = (this == IPCID_SRCDIAG_FORGETALL)
	return
}

// IsIPCID_SRCDIAG_PEEKHIDDEN returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_PEEKHIDDEN`.
func (this IpcIDs) IsIPCID_SRCDIAG_PEEKHIDDEN() (ret bool) {
	ret = (this == IPCID_SRCDIAG_PEEKHIDDEN)
	return
}

// IsIPCID_SRCDIAG_PUB returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_PUB`.
func (this IpcIDs) IsIPCID_SRCDIAG_PUB() (ret bool) { ret = (this == IPCID_SRCDIAG_PUB); return }

// IsIPCID_SRCDIAG_AUTO_TOGGLE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_TOGGLE`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_TOGGLE() (ret bool) {
	ret = (this == IPCID_SRCDIAG_AUTO_TOGGLE)
	return
}

// IsIPCID_SRCDIAG_AUTO_ALL returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_ALL`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_ALL() (ret bool) {
	ret = (this == IPCID_SRCDIAG_AUTO_ALL)
	return
}

// IsIPCID_SRCDIAG_AUTO_NONE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_NONE`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_NONE() (ret bool) {
	ret = (this == IPCID_SRCDIAG_AUTO_NONE)
	return
}

// IsIPCID_SRCDIAG_STARTED returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_STARTED`.
func (this IpcIDs) IsIPCID_SRCDIAG_STARTED() (ret bool) { ret = (this == IPCID_SRCDIAG_STARTED); return }

// IsIPCID_SRCDIAG_FINISHED returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_FINISHED`.
func (this IpcIDs) IsIPCID_SRCDIAG_FINISHED() (ret bool) {
	ret = (this == IPCID_SRCDIAG_FINISHED)
	return
}

// IsIPCID_SRCMOD_FMT_SETDEFMENU returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_SETDEFMENU`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_SETDEFMENU() (ret bool) {
	ret = (this == IPCID_SRCMOD_FMT_SETDEFMENU)
	return
}

// IsIPCID_SRCMOD_FMT_SETDEFPICK returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_SETDEFPICK`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_SETDEFPICK() (ret bool) {
	ret = (this == IPCID_SRCMOD_FMT_SETDEFPICK)
	return
}

// IsIPCID_SRCMOD_FMT_RUNONFILE returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_RUNONFILE`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_RUNONFILE() (ret bool) {
	ret = (this == IPCID_SRCMOD_FMT_RUNONFILE)
	return
}

// IsIPCID_SRCMOD_FMT_RUNONSEL returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_RUNONSEL`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_RUNONSEL() (ret bool) {
	ret = (this == IPCID_SRCMOD_FMT_RUNONSEL)
	return
}

// IsIPCID_SRCMOD_RENAME returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_RENAME`.
func (this IpcIDs) IsIPCID_SRCMOD_RENAME() (ret bool) { ret = (this == IPCID_SRCMOD_RENAME); return }

// IsIPCID_SRCMOD_ACTIONS returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_ACTIONS`.
func (this IpcIDs) IsIPCID_SRCMOD_ACTIONS() (ret bool) { ret = (this == IPCID_SRCMOD_ACTIONS); return }

// IsIPCID_SRCINTEL_HOVER returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_HOVER`.
func (this IpcIDs) IsIPCID_SRCINTEL_HOVER() (ret bool) { ret = (this == IPCID_SRCINTEL_HOVER); return }

// IsIPCID_SRCINTEL_SYMS_FILE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SYMS_FILE`.
func (this IpcIDs) IsIPCID_SRCINTEL_SYMS_FILE() (ret bool) {
	ret = (this == IPCID_SRCINTEL_SYMS_FILE)
	return
}

// IsIPCID_SRCINTEL_SYMS_PROJ returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SYMS_PROJ`.
func (this IpcIDs) IsIPCID_SRCINTEL_SYMS_PROJ() (ret bool) {
	ret = (this == IPCID_SRCINTEL_SYMS_PROJ)
	return
}

// IsIPCID_SRCINTEL_CMPL_ITEMS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_CMPL_ITEMS`.
func (this IpcIDs) IsIPCID_SRCINTEL_CMPL_ITEMS() (ret bool) {
	ret = (this == IPCID_SRCINTEL_CMPL_ITEMS)
	return
}

// IsIPCID_SRCINTEL_CMPL_DETAILS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_CMPL_DETAILS`.
func (this IpcIDs) IsIPCID_SRCINTEL_CMPL_DETAILS() (ret bool) {
	ret = (this == IPCID_SRCINTEL_CMPL_DETAILS)
	return
}

// IsIPCID_SRCINTEL_HIGHLIGHTS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_HIGHLIGHTS`.
func (this IpcIDs) IsIPCID_SRCINTEL_HIGHLIGHTS() (ret bool) {
	ret = (this == IPCID_SRCINTEL_HIGHLIGHTS)
	return
}

// IsIPCID_SRCINTEL_SIGNATURE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SIGNATURE`.
func (this IpcIDs) IsIPCID_SRCINTEL_SIGNATURE() (ret bool) {
	ret = (this == IPCID_SRCINTEL_SIGNATURE)
	return
}

// IsIPCID_SRCINTEL_REFERENCES returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_REFERENCES`.
func (this IpcIDs) IsIPCID_SRCINTEL_REFERENCES() (ret bool) {
	ret = (this == IPCID_SRCINTEL_REFERENCES)
	return
}

// IsIPCID_SRCINTEL_DEFSYM returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFSYM`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFSYM() (ret bool) { ret = (this == IPCID_SRCINTEL_DEFSYM); return }

// IsIPCID_SRCINTEL_DEFTYPE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFTYPE`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFTYPE() (ret bool) {
	ret = (this == IPCID_SRCINTEL_DEFTYPE)
	return
}

// IsIPCID_SRCINTEL_DEFIMPL returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFIMPL`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFIMPL() (ret bool) {
	ret = (this == IPCID_SRCINTEL_DEFIMPL)
	return
}

// IsIPCID_EXTRAS_INTEL_LIST returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_INTEL_LIST`.
func (this IpcIDs) IsIPCID_EXTRAS_INTEL_LIST() (ret bool) {
	ret = (this == IPCID_EXTRAS_INTEL_LIST)
	return
}

// IsIPCID_EXTRAS_INTEL_RUN returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_INTEL_RUN`.
func (this IpcIDs) IsIPCID_EXTRAS_INTEL_RUN() (ret bool) {
	ret = (this == IPCID_EXTRAS_INTEL_RUN)
	return
}

// IsIPCID_EXTRAS_QUERY_LIST returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_QUERY_LIST`.
func (this IpcIDs) IsIPCID_EXTRAS_QUERY_LIST() (ret bool) {
	ret = (this == IPCID_EXTRAS_QUERY_LIST)
	return
}

// IsIPCID_EXTRAS_QUERY_RUN returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_QUERY_RUN`.
func (this IpcIDs) IsIPCID_EXTRAS_QUERY_RUN() (ret bool) {
	ret = (this == IPCID_EXTRAS_QUERY_RUN)
	return
}

func (this IpcIDs) String() (ret string) {
	switch this {
	case IPCID_MENUS_MAIN:
		ret = "IPCID_MENUS_MAIN"
	case IPCID_MENUS_PKGS:
		ret = "IPCID_MENUS_PKGS"
	case IPCID_MENUS_TOOLS:
		ret = "IPCID_MENUS_TOOLS"
	case IPCID_OBJ_SNAPSHOT:
		ret = "IPCID_OBJ_SNAPSHOT"
	case IPCID_PAGE_HTML:
		ret = "IPCID_PAGE_HTML"
	case IPCID_TREEVIEW_GETITEM:
		ret = "IPCID_TREEVIEW_GETITEM"
	case IPCID_TREEVIEW_CHILDREN:
		ret = "IPCID_TREEVIEW_CHILDREN"
	case IPCID_TREEVIEW_CHANGED:
		ret = "IPCID_TREEVIEW_CHANGED"
	case IPCID_CFG_RESETALL:
		ret = "IPCID_CFG_RESETALL"
	case IPCID_CFG_LIST:
		ret = "IPCID_CFG_LIST"
	case IPCID_CFG_SET:
		ret = "IPCID_CFG_SET"
	case IPCID_NOTIFY_INFO:
		ret = "IPCID_NOTIFY_INFO"
	case IPCID_NOTIFY_WARN:
		ret = "IPCID_NOTIFY_WARN"
	case IPCID_NOTIFY_ERR:
		ret = "IPCID_NOTIFY_ERR"
	case IPCID_PROJ_CHANGED:
		ret = "IPCID_PROJ_CHANGED"
	case IPCID_PROJ_POLLEVTS:
		ret = "IPCID_PROJ_POLLEVTS"
	case IPCID_SRCDIAG_LIST:
		ret = "IPCID_SRCDIAG_LIST"
	case IPCID_SRCDIAG_RUN_CURFILE:
		ret = "IPCID_SRCDIAG_RUN_CURFILE"
	case IPCID_SRCDIAG_RUN_OPENFILES:
		ret = "IPCID_SRCDIAG_RUN_OPENFILES"
	case IPCID_SRCDIAG_RUN_ALLFILES:
		ret = "IPCID_SRCDIAG_RUN_ALLFILES"
	case IPCID_SRCDIAG_FORGETALL:
		ret = "IPCID_SRCDIAG_FORGETALL"
	case IPCID_SRCDIAG_PEEKHIDDEN:
		ret = "IPCID_SRCDIAG_PEEKHIDDEN"
	case IPCID_SRCDIAG_PUB:
		ret = "IPCID_SRCDIAG_PUB"
	case IPCID_SRCDIAG_AUTO_TOGGLE:
		ret = "IPCID_SRCDIAG_AUTO_TOGGLE"
	case IPCID_SRCDIAG_AUTO_ALL:
		ret = "IPCID_SRCDIAG_AUTO_ALL"
	case IPCID_SRCDIAG_AUTO_NONE:
		ret = "IPCID_SRCDIAG_AUTO_NONE"
	case IPCID_SRCDIAG_STARTED:
		ret = "IPCID_SRCDIAG_STARTED"
	case IPCID_SRCDIAG_FINISHED:
		ret = "IPCID_SRCDIAG_FINISHED"
	case IPCID_SRCMOD_FMT_SETDEFMENU:
		ret = "IPCID_SRCMOD_FMT_SETDEFMENU"
	case IPCID_SRCMOD_FMT_SETDEFPICK:
		ret = "IPCID_SRCMOD_FMT_SETDEFPICK"
	case IPCID_SRCMOD_FMT_RUNONFILE:
		ret = "IPCID_SRCMOD_FMT_RUNONFILE"
	case IPCID_SRCMOD_FMT_RUNONSEL:
		ret = "IPCID_SRCMOD_FMT_RUNONSEL"
	case IPCID_SRCMOD_RENAME:
		ret = "IPCID_SRCMOD_RENAME"
	case IPCID_SRCMOD_ACTIONS:
		ret = "IPCID_SRCMOD_ACTIONS"
	case IPCID_SRCINTEL_HOVER:
		ret = "IPCID_SRCINTEL_HOVER"
	case IPCID_SRCINTEL_SYMS_FILE:
		ret = "IPCID_SRCINTEL_SYMS_FILE"
	case IPCID_SRCINTEL_SYMS_PROJ:
		ret = "IPCID_SRCINTEL_SYMS_PROJ"
	case IPCID_SRCINTEL_CMPL_ITEMS:
		ret = "IPCID_SRCINTEL_CMPL_ITEMS"
	case IPCID_SRCINTEL_CMPL_DETAILS:
		ret = "IPCID_SRCINTEL_CMPL_DETAILS"
	case IPCID_SRCINTEL_HIGHLIGHTS:
		ret = "IPCID_SRCINTEL_HIGHLIGHTS"
	case IPCID_SRCINTEL_SIGNATURE:
		ret = "IPCID_SRCINTEL_SIGNATURE"
	case IPCID_SRCINTEL_REFERENCES:
		ret = "IPCID_SRCINTEL_REFERENCES"
	case IPCID_SRCINTEL_DEFSYM:
		ret = "IPCID_SRCINTEL_DEFSYM"
	case IPCID_SRCINTEL_DEFTYPE:
		ret = "IPCID_SRCINTEL_DEFTYPE"
	case IPCID_SRCINTEL_DEFIMPL:
		ret = "IPCID_SRCINTEL_DEFIMPL"
	case IPCID_EXTRAS_INTEL_LIST:
		ret = "IPCID_EXTRAS_INTEL_LIST"
	case IPCID_EXTRAS_INTEL_RUN:
		ret = "IPCID_EXTRAS_INTEL_RUN"
	case IPCID_EXTRAS_QUERY_LIST:
		ret = "IPCID_EXTRAS_QUERY_LIST"
	case IPCID_EXTRAS_QUERY_RUN:
		ret = "IPCID_EXTRAS_QUERY_RUN"
	default:
		ret = pkg__strconv.FormatUint(uint64(this), 10)
	}
	return
}

func IpcIDsFromString(s string) (this IpcIDs, err error) { return }

// Valid returns whether the value of this `DiagSeverity` is between `DIAG_SEV_ERR` (inclusive) and `DIAG_SEV_HINT` (inclusive).
func (this DiagSeverity) Valid() (ret bool) {
	ret = ((this >= DIAG_SEV_ERR) && (this <= DIAG_SEV_HINT))
	return
}

// IsDIAG_SEV_ERR returns whether the value of this `DiagSeverity` equals `DIAG_SEV_ERR`.
func (this DiagSeverity) IsDIAG_SEV_ERR() (ret bool) { ret = (this == DIAG_SEV_ERR); return }

// IsDIAG_SEV_WARN returns whether the value of this `DiagSeverity` equals `DIAG_SEV_WARN`.
func (this DiagSeverity) IsDIAG_SEV_WARN() (ret bool) { ret = (this == DIAG_SEV_WARN); return }

// IsDIAG_SEV_INFO returns whether the value of this `DiagSeverity` equals `DIAG_SEV_INFO`.
func (this DiagSeverity) IsDIAG_SEV_INFO() (ret bool) { ret = (this == DIAG_SEV_INFO); return }

// IsDIAG_SEV_HINT returns whether the value of this `DiagSeverity` equals `DIAG_SEV_HINT`.
func (this DiagSeverity) IsDIAG_SEV_HINT() (ret bool) { ret = (this == DIAG_SEV_HINT); return }

func (this DiagSeverity) String() (ret string) {
	switch this {
	case DIAG_SEV_ERR:
		ret = "DIAG_SEV_ERR"
	case DIAG_SEV_WARN:
		ret = "DIAG_SEV_WARN"
	case DIAG_SEV_INFO:
		ret = "DIAG_SEV_INFO"
	case DIAG_SEV_HINT:
		ret = "DIAG_SEV_HINT"
	default:
		ret = pkg__strconv.Itoa(int(this))
	}
	return
}

func DiagSeverityFromString(s string) (this DiagSeverity, err error) { return }

// Valid returns whether the value of this `Symbol` is between `SYM_FILE` (inclusive) and `SYM_TYPEPARAMETER` (inclusive).
func (this Symbol) Valid() (ret bool) {
	ret = ((this >= SYM_FILE) && (this <= SYM_TYPEPARAMETER))
	return
}

// IsSYM_FILE returns whether the value of this `Symbol` equals `SYM_FILE`.
func (this Symbol) IsSYM_FILE() (ret bool) { ret = (this == SYM_FILE); return }

// IsSYM_MODULE returns whether the value of this `Symbol` equals `SYM_MODULE`.
func (this Symbol) IsSYM_MODULE() (ret bool) { ret = (this == SYM_MODULE); return }

// IsSYM_NAMESPACE returns whether the value of this `Symbol` equals `SYM_NAMESPACE`.
func (this Symbol) IsSYM_NAMESPACE() (ret bool) { ret = (this == SYM_NAMESPACE); return }

// IsSYM_PACKAGE returns whether the value of this `Symbol` equals `SYM_PACKAGE`.
func (this Symbol) IsSYM_PACKAGE() (ret bool) { ret = (this == SYM_PACKAGE); return }

// IsSYM_CLASS returns whether the value of this `Symbol` equals `SYM_CLASS`.
func (this Symbol) IsSYM_CLASS() (ret bool) { ret = (this == SYM_CLASS); return }

// IsSYM_METHOD returns whether the value of this `Symbol` equals `SYM_METHOD`.
func (this Symbol) IsSYM_METHOD() (ret bool) { ret = (this == SYM_METHOD); return }

// IsSYM_PROPERTY returns whether the value of this `Symbol` equals `SYM_PROPERTY`.
func (this Symbol) IsSYM_PROPERTY() (ret bool) { ret = (this == SYM_PROPERTY); return }

// IsSYM_FIELD returns whether the value of this `Symbol` equals `SYM_FIELD`.
func (this Symbol) IsSYM_FIELD() (ret bool) { ret = (this == SYM_FIELD); return }

// IsSYM_CONSTRUCTOR returns whether the value of this `Symbol` equals `SYM_CONSTRUCTOR`.
func (this Symbol) IsSYM_CONSTRUCTOR() (ret bool) { ret = (this == SYM_CONSTRUCTOR); return }

// IsSYM_ENUM returns whether the value of this `Symbol` equals `SYM_ENUM`.
func (this Symbol) IsSYM_ENUM() (ret bool) { ret = (this == SYM_ENUM); return }

// IsSYM_INTERFACE returns whether the value of this `Symbol` equals `SYM_INTERFACE`.
func (this Symbol) IsSYM_INTERFACE() (ret bool) { ret = (this == SYM_INTERFACE); return }

// IsSYM_FUNCTION returns whether the value of this `Symbol` equals `SYM_FUNCTION`.
func (this Symbol) IsSYM_FUNCTION() (ret bool) { ret = (this == SYM_FUNCTION); return }

// IsSYM_VARIABLE returns whether the value of this `Symbol` equals `SYM_VARIABLE`.
func (this Symbol) IsSYM_VARIABLE() (ret bool) { ret = (this == SYM_VARIABLE); return }

// IsSYM_CONSTANT returns whether the value of this `Symbol` equals `SYM_CONSTANT`.
func (this Symbol) IsSYM_CONSTANT() (ret bool) { ret = (this == SYM_CONSTANT); return }

// IsSYM_STRING returns whether the value of this `Symbol` equals `SYM_STRING`.
func (this Symbol) IsSYM_STRING() (ret bool) { ret = (this == SYM_STRING); return }

// IsSYM_NUMBER returns whether the value of this `Symbol` equals `SYM_NUMBER`.
func (this Symbol) IsSYM_NUMBER() (ret bool) { ret = (this == SYM_NUMBER); return }

// IsSYM_BOOLEAN returns whether the value of this `Symbol` equals `SYM_BOOLEAN`.
func (this Symbol) IsSYM_BOOLEAN() (ret bool) { ret = (this == SYM_BOOLEAN); return }

// IsSYM_ARRAY returns whether the value of this `Symbol` equals `SYM_ARRAY`.
func (this Symbol) IsSYM_ARRAY() (ret bool) { ret = (this == SYM_ARRAY); return }

// IsSYM_OBJECT returns whether the value of this `Symbol` equals `SYM_OBJECT`.
func (this Symbol) IsSYM_OBJECT() (ret bool) { ret = (this == SYM_OBJECT); return }

// IsSYM_KEY returns whether the value of this `Symbol` equals `SYM_KEY`.
func (this Symbol) IsSYM_KEY() (ret bool) { ret = (this == SYM_KEY); return }

// IsSYM_NULL returns whether the value of this `Symbol` equals `SYM_NULL`.
func (this Symbol) IsSYM_NULL() (ret bool) { ret = (this == SYM_NULL); return }

// IsSYM_ENUMMEMBER returns whether the value of this `Symbol` equals `SYM_ENUMMEMBER`.
func (this Symbol) IsSYM_ENUMMEMBER() (ret bool) { ret = (this == SYM_ENUMMEMBER); return }

// IsSYM_STRUCT returns whether the value of this `Symbol` equals `SYM_STRUCT`.
func (this Symbol) IsSYM_STRUCT() (ret bool) { ret = (this == SYM_STRUCT); return }

// IsSYM_EVENT returns whether the value of this `Symbol` equals `SYM_EVENT`.
func (this Symbol) IsSYM_EVENT() (ret bool) { ret = (this == SYM_EVENT); return }

// IsSYM_OPERATOR returns whether the value of this `Symbol` equals `SYM_OPERATOR`.
func (this Symbol) IsSYM_OPERATOR() (ret bool) { ret = (this == SYM_OPERATOR); return }

// IsSYM_TYPEPARAMETER returns whether the value of this `Symbol` equals `SYM_TYPEPARAMETER`.
func (this Symbol) IsSYM_TYPEPARAMETER() (ret bool) { ret = (this == SYM_TYPEPARAMETER); return }

func (this Symbol) String() (ret string) {
	switch this {
	case SYM_FILE:
		ret = "SYM_FILE"
	case SYM_MODULE:
		ret = "SYM_MODULE"
	case SYM_NAMESPACE:
		ret = "SYM_NAMESPACE"
	case SYM_PACKAGE:
		ret = "SYM_PACKAGE"
	case SYM_CLASS:
		ret = "SYM_CLASS"
	case SYM_METHOD:
		ret = "SYM_METHOD"
	case SYM_PROPERTY:
		ret = "SYM_PROPERTY"
	case SYM_FIELD:
		ret = "SYM_FIELD"
	case SYM_CONSTRUCTOR:
		ret = "SYM_CONSTRUCTOR"
	case SYM_ENUM:
		ret = "SYM_ENUM"
	case SYM_INTERFACE:
		ret = "SYM_INTERFACE"
	case SYM_FUNCTION:
		ret = "SYM_FUNCTION"
	case SYM_VARIABLE:
		ret = "SYM_VARIABLE"
	case SYM_CONSTANT:
		ret = "SYM_CONSTANT"
	case SYM_STRING:
		ret = "SYM_STRING"
	case SYM_NUMBER:
		ret = "SYM_NUMBER"
	case SYM_BOOLEAN:
		ret = "SYM_BOOLEAN"
	case SYM_ARRAY:
		ret = "SYM_ARRAY"
	case SYM_OBJECT:
		ret = "SYM_OBJECT"
	case SYM_KEY:
		ret = "SYM_KEY"
	case SYM_NULL:
		ret = "SYM_NULL"
	case SYM_ENUMMEMBER:
		ret = "SYM_ENUMMEMBER"
	case SYM_STRUCT:
		ret = "SYM_STRUCT"
	case SYM_EVENT:
		ret = "SYM_EVENT"
	case SYM_OPERATOR:
		ret = "SYM_OPERATOR"
	case SYM_TYPEPARAMETER:
		ret = "SYM_TYPEPARAMETER"
	default:
		ret = pkg__strconv.FormatUint(uint64(this), 10)
	}
	return
}

func SymbolFromString(s string) (this Symbol, err error) { return }

// Valid returns whether the value of this `Completion` is between `CMPL_TEXT` (inclusive) and `CMPL_TYPEPARAMETER` (inclusive).
func (this Completion) Valid() (ret bool) {
	ret = ((this >= CMPL_TEXT) && (this <= CMPL_TYPEPARAMETER))
	return
}

// IsCMPL_TEXT returns whether the value of this `Completion` equals `CMPL_TEXT`.
func (this Completion) IsCMPL_TEXT() (ret bool) { ret = (this == CMPL_TEXT); return }

// IsCMPL_METHOD returns whether the value of this `Completion` equals `CMPL_METHOD`.
func (this Completion) IsCMPL_METHOD() (ret bool) { ret = (this == CMPL_METHOD); return }

// IsCMPL_FUNCTION returns whether the value of this `Completion` equals `CMPL_FUNCTION`.
func (this Completion) IsCMPL_FUNCTION() (ret bool) { ret = (this == CMPL_FUNCTION); return }

// IsCMPL_CONSTRUCTOR returns whether the value of this `Completion` equals `CMPL_CONSTRUCTOR`.
func (this Completion) IsCMPL_CONSTRUCTOR() (ret bool) { ret = (this == CMPL_CONSTRUCTOR); return }

// IsCMPL_FIELD returns whether the value of this `Completion` equals `CMPL_FIELD`.
func (this Completion) IsCMPL_FIELD() (ret bool) { ret = (this == CMPL_FIELD); return }

// IsCMPL_VARIABLE returns whether the value of this `Completion` equals `CMPL_VARIABLE`.
func (this Completion) IsCMPL_VARIABLE() (ret bool) { ret = (this == CMPL_VARIABLE); return }

// IsCMPL_CLASS returns whether the value of this `Completion` equals `CMPL_CLASS`.
func (this Completion) IsCMPL_CLASS() (ret bool) { ret = (this == CMPL_CLASS); return }

// IsCMPL_INTERFACE returns whether the value of this `Completion` equals `CMPL_INTERFACE`.
func (this Completion) IsCMPL_INTERFACE() (ret bool) { ret = (this == CMPL_INTERFACE); return }

// IsCMPL_MODULE returns whether the value of this `Completion` equals `CMPL_MODULE`.
func (this Completion) IsCMPL_MODULE() (ret bool) { ret = (this == CMPL_MODULE); return }

// IsCMPL_PROPERTY returns whether the value of this `Completion` equals `CMPL_PROPERTY`.
func (this Completion) IsCMPL_PROPERTY() (ret bool) { ret = (this == CMPL_PROPERTY); return }

// IsCMPL_UNIT returns whether the value of this `Completion` equals `CMPL_UNIT`.
func (this Completion) IsCMPL_UNIT() (ret bool) { ret = (this == CMPL_UNIT); return }

// IsCMPL_VALUE returns whether the value of this `Completion` equals `CMPL_VALUE`.
func (this Completion) IsCMPL_VALUE() (ret bool) { ret = (this == CMPL_VALUE); return }

// IsCMPL_ENUM returns whether the value of this `Completion` equals `CMPL_ENUM`.
func (this Completion) IsCMPL_ENUM() (ret bool) { ret = (this == CMPL_ENUM); return }

// IsCMPL_KEYWORD returns whether the value of this `Completion` equals `CMPL_KEYWORD`.
func (this Completion) IsCMPL_KEYWORD() (ret bool) { ret = (this == CMPL_KEYWORD); return }

// IsCMPL_SNIPPET returns whether the value of this `Completion` equals `CMPL_SNIPPET`.
func (this Completion) IsCMPL_SNIPPET() (ret bool) { ret = (this == CMPL_SNIPPET); return }

// IsCMPL_COLOR returns whether the value of this `Completion` equals `CMPL_COLOR`.
func (this Completion) IsCMPL_COLOR() (ret bool) { ret = (this == CMPL_COLOR); return }

// IsCMPL_FILE returns whether the value of this `Completion` equals `CMPL_FILE`.
func (this Completion) IsCMPL_FILE() (ret bool) { ret = (this == CMPL_FILE); return }

// IsCMPL_REFERENCE returns whether the value of this `Completion` equals `CMPL_REFERENCE`.
func (this Completion) IsCMPL_REFERENCE() (ret bool) { ret = (this == CMPL_REFERENCE); return }

// IsCMPL_FOLDER returns whether the value of this `Completion` equals `CMPL_FOLDER`.
func (this Completion) IsCMPL_FOLDER() (ret bool) { ret = (this == CMPL_FOLDER); return }

// IsCMPL_ENUMMEMBER returns whether the value of this `Completion` equals `CMPL_ENUMMEMBER`.
func (this Completion) IsCMPL_ENUMMEMBER() (ret bool) { ret = (this == CMPL_ENUMMEMBER); return }

// IsCMPL_CONSTANT returns whether the value of this `Completion` equals `CMPL_CONSTANT`.
func (this Completion) IsCMPL_CONSTANT() (ret bool) { ret = (this == CMPL_CONSTANT); return }

// IsCMPL_STRUCT returns whether the value of this `Completion` equals `CMPL_STRUCT`.
func (this Completion) IsCMPL_STRUCT() (ret bool) { ret = (this == CMPL_STRUCT); return }

// IsCMPL_EVENT returns whether the value of this `Completion` equals `CMPL_EVENT`.
func (this Completion) IsCMPL_EVENT() (ret bool) { ret = (this == CMPL_EVENT); return }

// IsCMPL_OPERATOR returns whether the value of this `Completion` equals `CMPL_OPERATOR`.
func (this Completion) IsCMPL_OPERATOR() (ret bool) { ret = (this == CMPL_OPERATOR); return }

// IsCMPL_TYPEPARAMETER returns whether the value of this `Completion` equals `CMPL_TYPEPARAMETER`.
func (this Completion) IsCMPL_TYPEPARAMETER() (ret bool) { ret = (this == CMPL_TYPEPARAMETER); return }

func (this Completion) String() (ret string) {
	switch this {
	case CMPL_TEXT:
		ret = "CMPL_TEXT"
	case CMPL_METHOD:
		ret = "CMPL_METHOD"
	case CMPL_FUNCTION:
		ret = "CMPL_FUNCTION"
	case CMPL_CONSTRUCTOR:
		ret = "CMPL_CONSTRUCTOR"
	case CMPL_FIELD:
		ret = "CMPL_FIELD"
	case CMPL_VARIABLE:
		ret = "CMPL_VARIABLE"
	case CMPL_CLASS:
		ret = "CMPL_CLASS"
	case CMPL_INTERFACE:
		ret = "CMPL_INTERFACE"
	case CMPL_MODULE:
		ret = "CMPL_MODULE"
	case CMPL_PROPERTY:
		ret = "CMPL_PROPERTY"
	case CMPL_UNIT:
		ret = "CMPL_UNIT"
	case CMPL_VALUE:
		ret = "CMPL_VALUE"
	case CMPL_ENUM:
		ret = "CMPL_ENUM"
	case CMPL_KEYWORD:
		ret = "CMPL_KEYWORD"
	case CMPL_SNIPPET:
		ret = "CMPL_SNIPPET"
	case CMPL_COLOR:
		ret = "CMPL_COLOR"
	case CMPL_FILE:
		ret = "CMPL_FILE"
	case CMPL_REFERENCE:
		ret = "CMPL_REFERENCE"
	case CMPL_FOLDER:
		ret = "CMPL_FOLDER"
	case CMPL_ENUMMEMBER:
		ret = "CMPL_ENUMMEMBER"
	case CMPL_CONSTANT:
		ret = "CMPL_CONSTANT"
	case CMPL_STRUCT:
		ret = "CMPL_STRUCT"
	case CMPL_EVENT:
		ret = "CMPL_EVENT"
	case CMPL_OPERATOR:
		ret = "CMPL_OPERATOR"
	case CMPL_TYPEPARAMETER:
		ret = "CMPL_TYPEPARAMETER"
	default:
		ret = pkg__strconv.FormatUint(uint64(this), 10)
	}
	return
}

func CompletionFromString(s string) (this Completion, err error) { return }

// Valid returns whether the value of this `ToolCats` is between `TOOLS_CAT_MOD_REN` (inclusive) and `TOOLS_CAT_RUNONSAVE` (inclusive).
func (this ToolCats) Valid() (ret bool) {
	ret = ((this >= TOOLS_CAT_MOD_REN) && (this <= TOOLS_CAT_RUNONSAVE))
	return
}

// IsTOOLS_CAT_MOD_REN returns whether the value of this `ToolCats` equals `TOOLS_CAT_MOD_REN`.
func (this ToolCats) IsTOOLS_CAT_MOD_REN() (ret bool) { ret = (this == TOOLS_CAT_MOD_REN); return }

// IsTOOLS_CAT_MOD_FMT returns whether the value of this `ToolCats` equals `TOOLS_CAT_MOD_FMT`.
func (this ToolCats) IsTOOLS_CAT_MOD_FMT() (ret bool) { ret = (this == TOOLS_CAT_MOD_FMT); return }

// IsTOOLS_CAT_INTEL_TIPS returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_TIPS`.
func (this ToolCats) IsTOOLS_CAT_INTEL_TIPS() (ret bool) { ret = (this == TOOLS_CAT_INTEL_TIPS); return }

// IsTOOLS_CAT_INTEL_SYMS returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_SYMS`.
func (this ToolCats) IsTOOLS_CAT_INTEL_SYMS() (ret bool) { ret = (this == TOOLS_CAT_INTEL_SYMS); return }

// IsTOOLS_CAT_INTEL_HIGH returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_HIGH`.
func (this ToolCats) IsTOOLS_CAT_INTEL_HIGH() (ret bool) { ret = (this == TOOLS_CAT_INTEL_HIGH); return }

// IsTOOLS_CAT_INTEL_CMPL returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_CMPL`.
func (this ToolCats) IsTOOLS_CAT_INTEL_CMPL() (ret bool) { ret = (this == TOOLS_CAT_INTEL_CMPL); return }

// IsTOOLS_CAT_INTEL_NAV returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_NAV`.
func (this ToolCats) IsTOOLS_CAT_INTEL_NAV() (ret bool) { ret = (this == TOOLS_CAT_INTEL_NAV); return }

// IsTOOLS_CAT_EXTRAS_QUERY returns whether the value of this `ToolCats` equals `TOOLS_CAT_EXTRAS_QUERY`.
func (this ToolCats) IsTOOLS_CAT_EXTRAS_QUERY() (ret bool) {
	ret = (this == TOOLS_CAT_EXTRAS_QUERY)
	return
}

// IsTOOLS_CAT_DIAGS returns whether the value of this `ToolCats` equals `TOOLS_CAT_DIAGS`.
func (this ToolCats) IsTOOLS_CAT_DIAGS() (ret bool) { ret = (this == TOOLS_CAT_DIAGS); return }

// IsTOOLS_CAT_RUNONSAVE returns whether the value of this `ToolCats` equals `TOOLS_CAT_RUNONSAVE`.
func (this ToolCats) IsTOOLS_CAT_RUNONSAVE() (ret bool) { ret = (this == TOOLS_CAT_RUNONSAVE); return }
