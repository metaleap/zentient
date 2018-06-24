package z

// DO NOT EDIT: code generated with `demo-go-gent` using `github.com/metaleap/go-gent`

import (
	pkg__strconv "strconv"
)

// Valid returns whether the value of this `CaddyStatus` is between `CADDY_PENDING` (inclusive) and `CADDY_GOOD` (inclusive).
func (this CaddyStatus) Valid() (ret bool) {
	ret = (this >= CADDY_PENDING) && (this <= CADDY_GOOD)
	return
}

// String implements the `fmt.Stringer` interface.
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

// CaddyStatusFromString returns the `CaddyStatus` represented by `s` (as returned by `String`, and case-sensitively), or an `error` if none exists.
func CaddyStatusFromString(s string) (this CaddyStatus, err error) {
	switch {
	case s == "CADDY_PENDING":
		this = CADDY_PENDING
	case s == "CADDY_ERROR":
		this = CADDY_ERROR
	case s == "CADDY_BUSY":
		this = CADDY_BUSY
	case s == "CADDY_GOOD":
		this = CADDY_GOOD
	default:
		var thisuint8 uint64
		thisuint8, err = pkg__strconv.ParseUint(s, 10, 8)
		if err == nil {
			this = CaddyStatus(thisuint8)
		}
	}
	return
}

// CaddyStatusFromStringOr is like `CaddyStatusFromString` but returns `fallback` for bad inputs.
func CaddyStatusFromStringOr(s string, fallback CaddyStatus) (this CaddyStatus) {
	maybeCaddyStatus, err := CaddyStatusFromString(s)
	if err == nil {
		this = maybeCaddyStatus
	} else {
		this = fallback
	}
	return
}

// WellknownCaddyStatuses returns the `names` and `values` of all 4 well-known `CaddyStatus` enumerants.
func WellknownCaddyStatuses() (names []string, values []CaddyStatus) {
	names, values = []string{"CADDY_PENDING", "CADDY_ERROR", "CADDY_BUSY", "CADDY_GOOD"}, []CaddyStatus{CADDY_PENDING, CADDY_ERROR, CADDY_BUSY, CADDY_GOOD}
	return
}

func (this Settings) Indices(predicate func(Settings) bool) { return }

// Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN` (inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).
func (this IpcIDs) Valid() (ret bool) {
	ret = (this >= IPCID_MENUS_MAIN) && (this <= IPCID_EXTRAS_QUERY_RUN)
	return
}

// String implements the `fmt.Stringer` interface.
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

// IpcIDsFromString returns the `IpcIDs` represented by `s` (as returned by `String`, and case-sensitively), or an `error` if none exists.
func IpcIDsFromString(s string) (this IpcIDs, err error) {
	switch {
	case s == "IPCID_MENUS_MAIN":
		this = IPCID_MENUS_MAIN
	case s == "IPCID_MENUS_PKGS":
		this = IPCID_MENUS_PKGS
	case s == "IPCID_MENUS_TOOLS":
		this = IPCID_MENUS_TOOLS
	case s == "IPCID_OBJ_SNAPSHOT":
		this = IPCID_OBJ_SNAPSHOT
	case s == "IPCID_PAGE_HTML":
		this = IPCID_PAGE_HTML
	case s == "IPCID_TREEVIEW_GETITEM":
		this = IPCID_TREEVIEW_GETITEM
	case s == "IPCID_TREEVIEW_CHILDREN":
		this = IPCID_TREEVIEW_CHILDREN
	case s == "IPCID_TREEVIEW_CHANGED":
		this = IPCID_TREEVIEW_CHANGED
	case s == "IPCID_CFG_RESETALL":
		this = IPCID_CFG_RESETALL
	case s == "IPCID_CFG_LIST":
		this = IPCID_CFG_LIST
	case s == "IPCID_CFG_SET":
		this = IPCID_CFG_SET
	case s == "IPCID_NOTIFY_INFO":
		this = IPCID_NOTIFY_INFO
	case s == "IPCID_NOTIFY_WARN":
		this = IPCID_NOTIFY_WARN
	case s == "IPCID_NOTIFY_ERR":
		this = IPCID_NOTIFY_ERR
	case s == "IPCID_PROJ_CHANGED":
		this = IPCID_PROJ_CHANGED
	case s == "IPCID_PROJ_POLLEVTS":
		this = IPCID_PROJ_POLLEVTS
	case s == "IPCID_SRCDIAG_LIST":
		this = IPCID_SRCDIAG_LIST
	case s == "IPCID_SRCDIAG_RUN_CURFILE":
		this = IPCID_SRCDIAG_RUN_CURFILE
	case s == "IPCID_SRCDIAG_RUN_OPENFILES":
		this = IPCID_SRCDIAG_RUN_OPENFILES
	case s == "IPCID_SRCDIAG_RUN_ALLFILES":
		this = IPCID_SRCDIAG_RUN_ALLFILES
	case s == "IPCID_SRCDIAG_FORGETALL":
		this = IPCID_SRCDIAG_FORGETALL
	case s == "IPCID_SRCDIAG_PEEKHIDDEN":
		this = IPCID_SRCDIAG_PEEKHIDDEN
	case s == "IPCID_SRCDIAG_PUB":
		this = IPCID_SRCDIAG_PUB
	case s == "IPCID_SRCDIAG_AUTO_TOGGLE":
		this = IPCID_SRCDIAG_AUTO_TOGGLE
	case s == "IPCID_SRCDIAG_AUTO_ALL":
		this = IPCID_SRCDIAG_AUTO_ALL
	case s == "IPCID_SRCDIAG_AUTO_NONE":
		this = IPCID_SRCDIAG_AUTO_NONE
	case s == "IPCID_SRCDIAG_STARTED":
		this = IPCID_SRCDIAG_STARTED
	case s == "IPCID_SRCDIAG_FINISHED":
		this = IPCID_SRCDIAG_FINISHED
	case s == "IPCID_SRCMOD_FMT_SETDEFMENU":
		this = IPCID_SRCMOD_FMT_SETDEFMENU
	case s == "IPCID_SRCMOD_FMT_SETDEFPICK":
		this = IPCID_SRCMOD_FMT_SETDEFPICK
	case s == "IPCID_SRCMOD_FMT_RUNONFILE":
		this = IPCID_SRCMOD_FMT_RUNONFILE
	case s == "IPCID_SRCMOD_FMT_RUNONSEL":
		this = IPCID_SRCMOD_FMT_RUNONSEL
	case s == "IPCID_SRCMOD_RENAME":
		this = IPCID_SRCMOD_RENAME
	case s == "IPCID_SRCMOD_ACTIONS":
		this = IPCID_SRCMOD_ACTIONS
	case s == "IPCID_SRCINTEL_HOVER":
		this = IPCID_SRCINTEL_HOVER
	case s == "IPCID_SRCINTEL_SYMS_FILE":
		this = IPCID_SRCINTEL_SYMS_FILE
	case s == "IPCID_SRCINTEL_SYMS_PROJ":
		this = IPCID_SRCINTEL_SYMS_PROJ
	case s == "IPCID_SRCINTEL_CMPL_ITEMS":
		this = IPCID_SRCINTEL_CMPL_ITEMS
	case s == "IPCID_SRCINTEL_CMPL_DETAILS":
		this = IPCID_SRCINTEL_CMPL_DETAILS
	case s == "IPCID_SRCINTEL_HIGHLIGHTS":
		this = IPCID_SRCINTEL_HIGHLIGHTS
	case s == "IPCID_SRCINTEL_SIGNATURE":
		this = IPCID_SRCINTEL_SIGNATURE
	case s == "IPCID_SRCINTEL_REFERENCES":
		this = IPCID_SRCINTEL_REFERENCES
	case s == "IPCID_SRCINTEL_DEFSYM":
		this = IPCID_SRCINTEL_DEFSYM
	case s == "IPCID_SRCINTEL_DEFTYPE":
		this = IPCID_SRCINTEL_DEFTYPE
	case s == "IPCID_SRCINTEL_DEFIMPL":
		this = IPCID_SRCINTEL_DEFIMPL
	case s == "IPCID_EXTRAS_INTEL_LIST":
		this = IPCID_EXTRAS_INTEL_LIST
	case s == "IPCID_EXTRAS_INTEL_RUN":
		this = IPCID_EXTRAS_INTEL_RUN
	case s == "IPCID_EXTRAS_QUERY_LIST":
		this = IPCID_EXTRAS_QUERY_LIST
	case s == "IPCID_EXTRAS_QUERY_RUN":
		this = IPCID_EXTRAS_QUERY_RUN
	default:
		var thisuint8 uint64
		thisuint8, err = pkg__strconv.ParseUint(s, 10, 8)
		if err == nil {
			this = IpcIDs(thisuint8)
		}
	}
	return
}

// IpcIDsFromStringOr is like `IpcIDsFromString` but returns `fallback` for bad inputs.
func IpcIDsFromStringOr(s string, fallback IpcIDs) (this IpcIDs) {
	maybeIpcIDs, err := IpcIDsFromString(s)
	if err == nil {
		this = maybeIpcIDs
	} else {
		this = fallback
	}
	return
}

// WellknownIpcIDses returns the `names` and `values` of all 49 well-known `IpcIDs` enumerants.
func WellknownIpcIDses() (names []string, values []IpcIDs) {
	names, values = []string{"IPCID_MENUS_MAIN", "IPCID_MENUS_PKGS", "IPCID_MENUS_TOOLS", "IPCID_OBJ_SNAPSHOT", "IPCID_PAGE_HTML", "IPCID_TREEVIEW_GETITEM", "IPCID_TREEVIEW_CHILDREN", "IPCID_TREEVIEW_CHANGED", "IPCID_CFG_RESETALL", "IPCID_CFG_LIST", "IPCID_CFG_SET", "IPCID_NOTIFY_INFO", "IPCID_NOTIFY_WARN", "IPCID_NOTIFY_ERR", "IPCID_PROJ_CHANGED", "IPCID_PROJ_POLLEVTS", "IPCID_SRCDIAG_LIST", "IPCID_SRCDIAG_RUN_CURFILE", "IPCID_SRCDIAG_RUN_OPENFILES", "IPCID_SRCDIAG_RUN_ALLFILES", "IPCID_SRCDIAG_FORGETALL", "IPCID_SRCDIAG_PEEKHIDDEN", "IPCID_SRCDIAG_PUB", "IPCID_SRCDIAG_AUTO_TOGGLE", "IPCID_SRCDIAG_AUTO_ALL", "IPCID_SRCDIAG_AUTO_NONE", "IPCID_SRCDIAG_STARTED", "IPCID_SRCDIAG_FINISHED", "IPCID_SRCMOD_FMT_SETDEFMENU", "IPCID_SRCMOD_FMT_SETDEFPICK", "IPCID_SRCMOD_FMT_RUNONFILE", "IPCID_SRCMOD_FMT_RUNONSEL", "IPCID_SRCMOD_RENAME", "IPCID_SRCMOD_ACTIONS", "IPCID_SRCINTEL_HOVER", "IPCID_SRCINTEL_SYMS_FILE", "IPCID_SRCINTEL_SYMS_PROJ", "IPCID_SRCINTEL_CMPL_ITEMS", "IPCID_SRCINTEL_CMPL_DETAILS", "IPCID_SRCINTEL_HIGHLIGHTS", "IPCID_SRCINTEL_SIGNATURE", "IPCID_SRCINTEL_REFERENCES", "IPCID_SRCINTEL_DEFSYM", "IPCID_SRCINTEL_DEFTYPE", "IPCID_SRCINTEL_DEFIMPL", "IPCID_EXTRAS_INTEL_LIST", "IPCID_EXTRAS_INTEL_RUN", "IPCID_EXTRAS_QUERY_LIST", "IPCID_EXTRAS_QUERY_RUN"}, []IpcIDs{IPCID_MENUS_MAIN, IPCID_MENUS_PKGS, IPCID_MENUS_TOOLS, IPCID_OBJ_SNAPSHOT, IPCID_PAGE_HTML, IPCID_TREEVIEW_GETITEM, IPCID_TREEVIEW_CHILDREN, IPCID_TREEVIEW_CHANGED, IPCID_CFG_RESETALL, IPCID_CFG_LIST, IPCID_CFG_SET, IPCID_NOTIFY_INFO, IPCID_NOTIFY_WARN, IPCID_NOTIFY_ERR, IPCID_PROJ_CHANGED, IPCID_PROJ_POLLEVTS, IPCID_SRCDIAG_LIST, IPCID_SRCDIAG_RUN_CURFILE, IPCID_SRCDIAG_RUN_OPENFILES, IPCID_SRCDIAG_RUN_ALLFILES, IPCID_SRCDIAG_FORGETALL, IPCID_SRCDIAG_PEEKHIDDEN, IPCID_SRCDIAG_PUB, IPCID_SRCDIAG_AUTO_TOGGLE, IPCID_SRCDIAG_AUTO_ALL, IPCID_SRCDIAG_AUTO_NONE, IPCID_SRCDIAG_STARTED, IPCID_SRCDIAG_FINISHED, IPCID_SRCMOD_FMT_SETDEFMENU, IPCID_SRCMOD_FMT_SETDEFPICK, IPCID_SRCMOD_FMT_RUNONFILE, IPCID_SRCMOD_FMT_RUNONSEL, IPCID_SRCMOD_RENAME, IPCID_SRCMOD_ACTIONS, IPCID_SRCINTEL_HOVER, IPCID_SRCINTEL_SYMS_FILE, IPCID_SRCINTEL_SYMS_PROJ, IPCID_SRCINTEL_CMPL_ITEMS, IPCID_SRCINTEL_CMPL_DETAILS, IPCID_SRCINTEL_HIGHLIGHTS, IPCID_SRCINTEL_SIGNATURE, IPCID_SRCINTEL_REFERENCES, IPCID_SRCINTEL_DEFSYM, IPCID_SRCINTEL_DEFTYPE, IPCID_SRCINTEL_DEFIMPL, IPCID_EXTRAS_INTEL_LIST, IPCID_EXTRAS_INTEL_RUN, IPCID_EXTRAS_QUERY_LIST, IPCID_EXTRAS_QUERY_RUN}
	return
}

// Valid returns whether the value of this `DiagSeverity` is between `DIAG_SEV_ERR` (inclusive) and `DIAG_SEV_HINT` (inclusive).
func (this DiagSeverity) Valid() (ret bool) {
	ret = (this >= DIAG_SEV_ERR) && (this <= DIAG_SEV_HINT)
	return
}

// String implements the `fmt.Stringer` interface.
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

// DiagSeverityFromString returns the `DiagSeverity` represented by `s` (as returned by `String`, and case-sensitively), or an `error` if none exists.
func DiagSeverityFromString(s string) (this DiagSeverity, err error) {
	switch {
	case s == "DIAG_SEV_ERR":
		this = DIAG_SEV_ERR
	case s == "DIAG_SEV_WARN":
		this = DIAG_SEV_WARN
	case s == "DIAG_SEV_INFO":
		this = DIAG_SEV_INFO
	case s == "DIAG_SEV_HINT":
		this = DIAG_SEV_HINT
	default:
		var thisint int
		thisint, err = pkg__strconv.Atoi(s)
		if err == nil {
			this = DiagSeverity(thisint)
		}
	}
	return
}

// DiagSeverityFromStringOr is like `DiagSeverityFromString` but returns `fallback` for bad inputs.
func DiagSeverityFromStringOr(s string, fallback DiagSeverity) (this DiagSeverity) {
	maybeDiagSeverity, err := DiagSeverityFromString(s)
	if err == nil {
		this = maybeDiagSeverity
	} else {
		this = fallback
	}
	return
}

// WellknownDiagSeverities returns the `names` and `values` of all 4 well-known `DiagSeverity` enumerants.
func WellknownDiagSeverities() (names []string, values []DiagSeverity) {
	names, values = []string{"DIAG_SEV_ERR", "DIAG_SEV_WARN", "DIAG_SEV_INFO", "DIAG_SEV_HINT"}, []DiagSeverity{DIAG_SEV_ERR, DIAG_SEV_WARN, DIAG_SEV_INFO, DIAG_SEV_HINT}
	return
}

// Valid returns whether the value of this `Symbol` is between `SYM_FILE` (inclusive) and `SYM_TYPEPARAMETER` (inclusive).
func (this Symbol) Valid() (ret bool) { ret = (this >= SYM_FILE) && (this <= SYM_TYPEPARAMETER); return }

// String implements the `fmt.Stringer` interface.
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

// SymbolFromString returns the `Symbol` represented by `s` (as returned by `String`, and case-sensitively), or an `error` if none exists.
func SymbolFromString(s string) (this Symbol, err error) {
	switch {
	case s == "SYM_FILE":
		this = SYM_FILE
	case s == "SYM_MODULE":
		this = SYM_MODULE
	case s == "SYM_NAMESPACE":
		this = SYM_NAMESPACE
	case s == "SYM_PACKAGE":
		this = SYM_PACKAGE
	case s == "SYM_CLASS":
		this = SYM_CLASS
	case s == "SYM_METHOD":
		this = SYM_METHOD
	case s == "SYM_PROPERTY":
		this = SYM_PROPERTY
	case s == "SYM_FIELD":
		this = SYM_FIELD
	case s == "SYM_CONSTRUCTOR":
		this = SYM_CONSTRUCTOR
	case s == "SYM_ENUM":
		this = SYM_ENUM
	case s == "SYM_INTERFACE":
		this = SYM_INTERFACE
	case s == "SYM_FUNCTION":
		this = SYM_FUNCTION
	case s == "SYM_VARIABLE":
		this = SYM_VARIABLE
	case s == "SYM_CONSTANT":
		this = SYM_CONSTANT
	case s == "SYM_STRING":
		this = SYM_STRING
	case s == "SYM_NUMBER":
		this = SYM_NUMBER
	case s == "SYM_BOOLEAN":
		this = SYM_BOOLEAN
	case s == "SYM_ARRAY":
		this = SYM_ARRAY
	case s == "SYM_OBJECT":
		this = SYM_OBJECT
	case s == "SYM_KEY":
		this = SYM_KEY
	case s == "SYM_NULL":
		this = SYM_NULL
	case s == "SYM_ENUMMEMBER":
		this = SYM_ENUMMEMBER
	case s == "SYM_STRUCT":
		this = SYM_STRUCT
	case s == "SYM_EVENT":
		this = SYM_EVENT
	case s == "SYM_OPERATOR":
		this = SYM_OPERATOR
	case s == "SYM_TYPEPARAMETER":
		this = SYM_TYPEPARAMETER
	default:
		var thisuint8 uint64
		thisuint8, err = pkg__strconv.ParseUint(s, 10, 8)
		if err == nil {
			this = Symbol(thisuint8)
		}
	}
	return
}

// SymbolFromStringOr is like `SymbolFromString` but returns `fallback` for bad inputs.
func SymbolFromStringOr(s string, fallback Symbol) (this Symbol) {
	maybeSymbol, err := SymbolFromString(s)
	if err == nil {
		this = maybeSymbol
	} else {
		this = fallback
	}
	return
}

// WellknownSymbols returns the `names` and `values` of all 26 well-known `Symbol` enumerants.
func WellknownSymbols() (names []string, values []Symbol) {
	names, values = []string{"SYM_FILE", "SYM_MODULE", "SYM_NAMESPACE", "SYM_PACKAGE", "SYM_CLASS", "SYM_METHOD", "SYM_PROPERTY", "SYM_FIELD", "SYM_CONSTRUCTOR", "SYM_ENUM", "SYM_INTERFACE", "SYM_FUNCTION", "SYM_VARIABLE", "SYM_CONSTANT", "SYM_STRING", "SYM_NUMBER", "SYM_BOOLEAN", "SYM_ARRAY", "SYM_OBJECT", "SYM_KEY", "SYM_NULL", "SYM_ENUMMEMBER", "SYM_STRUCT", "SYM_EVENT", "SYM_OPERATOR", "SYM_TYPEPARAMETER"}, []Symbol{SYM_FILE, SYM_MODULE, SYM_NAMESPACE, SYM_PACKAGE, SYM_CLASS, SYM_METHOD, SYM_PROPERTY, SYM_FIELD, SYM_CONSTRUCTOR, SYM_ENUM, SYM_INTERFACE, SYM_FUNCTION, SYM_VARIABLE, SYM_CONSTANT, SYM_STRING, SYM_NUMBER, SYM_BOOLEAN, SYM_ARRAY, SYM_OBJECT, SYM_KEY, SYM_NULL, SYM_ENUMMEMBER, SYM_STRUCT, SYM_EVENT, SYM_OPERATOR, SYM_TYPEPARAMETER}
	return
}

// Valid returns whether the value of this `Completion` is between `CMPL_TEXT` (inclusive) and `CMPL_TYPEPARAMETER` (inclusive).
func (this Completion) Valid() (ret bool) {
	ret = (this >= CMPL_TEXT) && (this <= CMPL_TYPEPARAMETER)
	return
}

// String implements the `fmt.Stringer` interface.
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

// CompletionFromString returns the `Completion` represented by `s` (as returned by `String`, and case-sensitively), or an `error` if none exists.
func CompletionFromString(s string) (this Completion, err error) {
	switch {
	case s == "CMPL_TEXT":
		this = CMPL_TEXT
	case s == "CMPL_METHOD":
		this = CMPL_METHOD
	case s == "CMPL_FUNCTION":
		this = CMPL_FUNCTION
	case s == "CMPL_CONSTRUCTOR":
		this = CMPL_CONSTRUCTOR
	case s == "CMPL_FIELD":
		this = CMPL_FIELD
	case s == "CMPL_VARIABLE":
		this = CMPL_VARIABLE
	case s == "CMPL_CLASS":
		this = CMPL_CLASS
	case s == "CMPL_INTERFACE":
		this = CMPL_INTERFACE
	case s == "CMPL_MODULE":
		this = CMPL_MODULE
	case s == "CMPL_PROPERTY":
		this = CMPL_PROPERTY
	case s == "CMPL_UNIT":
		this = CMPL_UNIT
	case s == "CMPL_VALUE":
		this = CMPL_VALUE
	case s == "CMPL_ENUM":
		this = CMPL_ENUM
	case s == "CMPL_KEYWORD":
		this = CMPL_KEYWORD
	case s == "CMPL_SNIPPET":
		this = CMPL_SNIPPET
	case s == "CMPL_COLOR":
		this = CMPL_COLOR
	case s == "CMPL_FILE":
		this = CMPL_FILE
	case s == "CMPL_REFERENCE":
		this = CMPL_REFERENCE
	case s == "CMPL_FOLDER":
		this = CMPL_FOLDER
	case s == "CMPL_ENUMMEMBER":
		this = CMPL_ENUMMEMBER
	case s == "CMPL_CONSTANT":
		this = CMPL_CONSTANT
	case s == "CMPL_STRUCT":
		this = CMPL_STRUCT
	case s == "CMPL_EVENT":
		this = CMPL_EVENT
	case s == "CMPL_OPERATOR":
		this = CMPL_OPERATOR
	case s == "CMPL_TYPEPARAMETER":
		this = CMPL_TYPEPARAMETER
	default:
		var thisuint8 uint64
		thisuint8, err = pkg__strconv.ParseUint(s, 10, 8)
		if err == nil {
			this = Completion(thisuint8)
		}
	}
	return
}

// CompletionFromStringOr is like `CompletionFromString` but returns `fallback` for bad inputs.
func CompletionFromStringOr(s string, fallback Completion) (this Completion) {
	maybeCompletion, err := CompletionFromString(s)
	if err == nil {
		this = maybeCompletion
	} else {
		this = fallback
	}
	return
}

// WellknownCompletions returns the `names` and `values` of all 25 well-known `Completion` enumerants.
func WellknownCompletions() (names []string, values []Completion) {
	names, values = []string{"CMPL_TEXT", "CMPL_METHOD", "CMPL_FUNCTION", "CMPL_CONSTRUCTOR", "CMPL_FIELD", "CMPL_VARIABLE", "CMPL_CLASS", "CMPL_INTERFACE", "CMPL_MODULE", "CMPL_PROPERTY", "CMPL_UNIT", "CMPL_VALUE", "CMPL_ENUM", "CMPL_KEYWORD", "CMPL_SNIPPET", "CMPL_COLOR", "CMPL_FILE", "CMPL_REFERENCE", "CMPL_FOLDER", "CMPL_ENUMMEMBER", "CMPL_CONSTANT", "CMPL_STRUCT", "CMPL_EVENT", "CMPL_OPERATOR", "CMPL_TYPEPARAMETER"}, []Completion{CMPL_TEXT, CMPL_METHOD, CMPL_FUNCTION, CMPL_CONSTRUCTOR, CMPL_FIELD, CMPL_VARIABLE, CMPL_CLASS, CMPL_INTERFACE, CMPL_MODULE, CMPL_PROPERTY, CMPL_UNIT, CMPL_VALUE, CMPL_ENUM, CMPL_KEYWORD, CMPL_SNIPPET, CMPL_COLOR, CMPL_FILE, CMPL_REFERENCE, CMPL_FOLDER, CMPL_ENUMMEMBER, CMPL_CONSTANT, CMPL_STRUCT, CMPL_EVENT, CMPL_OPERATOR, CMPL_TYPEPARAMETER}
	return
}

func (this ListItems) Indices(predicate func(ListItems) bool) { return }

func (this PkgInfos) Indices(predicate func(PkgInfos) bool) { return }

func (this DiagItems) Indices(predicate func(DiagItems) bool) { return }

func (this DiagBuildJobs) Indices(predicate func(DiagBuildJobs) bool) { return }

func (this DiagLintJobs) Indices(predicate func(DiagLintJobs) bool) { return }

func (this SrcIntelCompls) Indices(predicate func(SrcIntelCompls) bool) { return }

func (this SrcLocs) Indices(predicate func(SrcLocs) bool) { return }

func (this SrcLenses) Indices(predicate func(SrcLenses) bool) { return }

func (this SrcModEdits) Indices(predicate func(SrcModEdits) bool) { return }

// Valid returns whether the value of this `ToolCats` is between `TOOLS_CAT_MOD_REN` (inclusive) and `TOOLS_CAT_RUNONSAVE` (inclusive).
func (this ToolCats) Valid() (ret bool) {
	ret = (this >= TOOLS_CAT_MOD_REN) && (this <= TOOLS_CAT_RUNONSAVE)
	return
}

// WellknownToolCatses returns the `names` and `values` of all 10 well-known `ToolCats` enumerants.
func WellknownToolCatses() (names []string, values []ToolCats) {
	names, values = []string{"TOOLS_CAT_MOD_REN", "TOOLS_CAT_MOD_FMT", "TOOLS_CAT_INTEL_TIPS", "TOOLS_CAT_INTEL_SYMS", "TOOLS_CAT_INTEL_HIGH", "TOOLS_CAT_INTEL_CMPL", "TOOLS_CAT_INTEL_NAV", "TOOLS_CAT_EXTRAS_QUERY", "TOOLS_CAT_DIAGS", "TOOLS_CAT_RUNONSAVE"}, []ToolCats{TOOLS_CAT_MOD_REN, TOOLS_CAT_MOD_FMT, TOOLS_CAT_INTEL_TIPS, TOOLS_CAT_INTEL_SYMS, TOOLS_CAT_INTEL_HIGH, TOOLS_CAT_INTEL_CMPL, TOOLS_CAT_INTEL_NAV, TOOLS_CAT_EXTRAS_QUERY, TOOLS_CAT_DIAGS, TOOLS_CAT_RUNONSAVE}
	return
}

func (this Tools) Indices(predicate func(Tools) bool) { return }

func (this MenuItems) Indices(predicate func(MenuItems) bool) { return }

func (this sideViewTreeItem) Indices(predicate func(sideViewTreeItem) bool) { return }
