package z

// DO NOT EDIT: code generated with `demo-go-gent` using `github.com/metaleap/go-gent`

import (
	pkg__strconv "strconv"
	pkg__strings "strings"
)

// IsCADDY_PENDING returns whether the value of this `CaddyStatus` equals `CADDY_PENDING`.
func (this CaddyStatus) IsCADDY_PENDING() (r bool) { r = this == CADDY_PENDING; return }

// IsCADDY_ERROR returns whether the value of this `CaddyStatus` equals `CADDY_ERROR`.
func (this CaddyStatus) IsCADDY_ERROR() (r bool) { r = this == CADDY_ERROR; return }

// IsCADDY_BUSY returns whether the value of this `CaddyStatus` equals `CADDY_BUSY`.
func (this CaddyStatus) IsCADDY_BUSY() (r bool) { r = this == CADDY_BUSY; return }

// IsCADDY_GOOD returns whether the value of this `CaddyStatus` equals `CADDY_GOOD`.
func (this CaddyStatus) IsCADDY_GOOD() (r bool) { r = this == CADDY_GOOD; return }

// Valid returns whether the value of this `CaddyStatus` is between `CADDY_PENDING` (inclusive) and `CADDY_GOOD` (inclusive).
func (this CaddyStatus) Valid() (r bool) { r = (this >= CADDY_PENDING) && (this <= CADDY_GOOD); return }

// WellknownCaddyStatuses returns the `names` and `values` of all 4 well-known `CaddyStatus` enumerants.
func WellknownCaddyStatuses() (names []string, values []CaddyStatus) {
	names, values = []string{"CADDY_PENDING", "CADDY_ERROR", "CADDY_BUSY", "CADDY_GOOD"}, []CaddyStatus{CADDY_PENDING, CADDY_ERROR, CADDY_BUSY, CADDY_GOOD}
	return
}

// String implements the `fmt.Stringer` interface.
func (this CaddyStatus) String() (r string) {
	if (this < CADDY_PENDING) || (this > CADDY_GOOD) {
		goto formatNum
	}
	switch this {
	case CADDY_PENDING:
		r = "Caddy·Pending"
	case CADDY_ERROR:
		r = "Caddy·Error"
	case CADDY_BUSY:
		r = "Caddy·Busy"
	case CADDY_GOOD:
		r = "Caddy·Good"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// CaddyStatusFromString returns the `CaddyStatus` represented by `s` (as returned by `CaddyStatus.String`, but case-insensitively), or an `error` if none exists.
func CaddyStatusFromString(s string) (this CaddyStatus, err error) {
	if (len(s) < 11) || (len(s) > 14) || (!pkg__strings.EqualFold(s[0:7], "Caddy·")) {
		goto tryParseNum
	}
	switch {
	case pkg__strings.EqualFold(s, "Caddy·Pending"):
		this = CADDY_PENDING
	case pkg__strings.EqualFold(s, "Caddy·Error"):
		this = CADDY_ERROR
	case pkg__strings.EqualFold(s, "Caddy·Busy"):
		this = CADDY_BUSY
	case pkg__strings.EqualFold(s, "Caddy·Good"):
		this = CADDY_GOOD
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (CaddyStatus)(v)
	}
	return
}

// CaddyStatusFromStringOr is like `CaddyStatusFromString` but returns `fallback` for unrecognized inputs.
func CaddyStatusFromStringOr(s string, fallback CaddyStatus) (this CaddyStatus) {
	maybeCaddyStatus, err := CaddyStatusFromString(s)
	if err == nil {
		this = maybeCaddyStatus
	} else {
		this = fallback
	}
	return
}

// GoString implements the `fmt.GoStringer` interface.
func (this CaddyStatus) GoString() (r string) {
	if (this < CADDY_PENDING) || (this > CADDY_GOOD) {
		goto formatNum
	}
	switch this {
	case CADDY_PENDING:
		r = "CADDY_PENDING"
	case CADDY_ERROR:
		r = "CADDY_ERROR"
	case CADDY_BUSY:
		r = "CADDY_BUSY"
	case CADDY_GOOD:
		r = "CADDY_GOOD"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// CaddyStatusFromGoString returns the `CaddyStatus` represented by `s` (as returned by `CaddyStatus.GoString`, and case-sensitively), or an `error` if none exists.
func CaddyStatusFromGoString(s string) (this CaddyStatus, err error) {
	if (len(s) < 10) || (len(s) > 13) || (s[0:6] != "CADDY_") {
		goto tryParseNum
	}
	switch s {
	case "CADDY_PENDING":
		this = CADDY_PENDING
	case "CADDY_ERROR":
		this = CADDY_ERROR
	case "CADDY_BUSY":
		this = CADDY_BUSY
	case "CADDY_GOOD":
		this = CADDY_GOOD
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (CaddyStatus)(v)
	}
	return
}

// CaddyStatusFromGoStringOr is like `CaddyStatusFromGoString` but returns `fallback` for unrecognized inputs.
func CaddyStatusFromGoStringOr(s string, fallback CaddyStatus) (this CaddyStatus) {
	maybeCaddyStatus, err := CaddyStatusFromGoString(s)
	if err == nil {
		this = maybeCaddyStatus
	} else {
		this = fallback
	}
	return
}

func (this Settings) Index(v *Setting) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Settings) IndexFunc(ok func(*Setting) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Settings) LastIndex(v *Setting) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Settings) LastIndexFunc(ok func(*Setting) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Settings) Indices(v *Setting) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this Settings) IndicesFunc(ok func(*Setting) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

// IsIPCID_MENUS_MAIN returns whether the value of this `IpcIDs` equals `IPCID_MENUS_MAIN`.
func (this IpcIDs) IsIPCID_MENUS_MAIN() (r bool) { r = this == IPCID_MENUS_MAIN; return }

// IsIPCID_MENUS_PKGS returns whether the value of this `IpcIDs` equals `IPCID_MENUS_PKGS`.
func (this IpcIDs) IsIPCID_MENUS_PKGS() (r bool) { r = this == IPCID_MENUS_PKGS; return }

// IsIPCID_MENUS_TOOLS returns whether the value of this `IpcIDs` equals `IPCID_MENUS_TOOLS`.
func (this IpcIDs) IsIPCID_MENUS_TOOLS() (r bool) { r = this == IPCID_MENUS_TOOLS; return }

// IsIPCID_OBJ_SNAPSHOT returns whether the value of this `IpcIDs` equals `IPCID_OBJ_SNAPSHOT`.
func (this IpcIDs) IsIPCID_OBJ_SNAPSHOT() (r bool) { r = this == IPCID_OBJ_SNAPSHOT; return }

// IsIPCID_PAGE_HTML returns whether the value of this `IpcIDs` equals `IPCID_PAGE_HTML`.
func (this IpcIDs) IsIPCID_PAGE_HTML() (r bool) { r = this == IPCID_PAGE_HTML; return }

// IsIPCID_TREEVIEW_GETITEM returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_GETITEM`.
func (this IpcIDs) IsIPCID_TREEVIEW_GETITEM() (r bool) { r = this == IPCID_TREEVIEW_GETITEM; return }

// IsIPCID_TREEVIEW_CHILDREN returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_CHILDREN`.
func (this IpcIDs) IsIPCID_TREEVIEW_CHILDREN() (r bool) { r = this == IPCID_TREEVIEW_CHILDREN; return }

// IsIPCID_TREEVIEW_CHANGED returns whether the value of this `IpcIDs` equals `IPCID_TREEVIEW_CHANGED`.
func (this IpcIDs) IsIPCID_TREEVIEW_CHANGED() (r bool) { r = this == IPCID_TREEVIEW_CHANGED; return }

// IsIPCID_CFG_RESETALL returns whether the value of this `IpcIDs` equals `IPCID_CFG_RESETALL`.
func (this IpcIDs) IsIPCID_CFG_RESETALL() (r bool) { r = this == IPCID_CFG_RESETALL; return }

// IsIPCID_CFG_LIST returns whether the value of this `IpcIDs` equals `IPCID_CFG_LIST`.
func (this IpcIDs) IsIPCID_CFG_LIST() (r bool) { r = this == IPCID_CFG_LIST; return }

// IsIPCID_CFG_SET returns whether the value of this `IpcIDs` equals `IPCID_CFG_SET`.
func (this IpcIDs) IsIPCID_CFG_SET() (r bool) { r = this == IPCID_CFG_SET; return }

// IsIPCID_NOTIFY_INFO returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_INFO`.
func (this IpcIDs) IsIPCID_NOTIFY_INFO() (r bool) { r = this == IPCID_NOTIFY_INFO; return }

// IsIPCID_NOTIFY_WARN returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_WARN`.
func (this IpcIDs) IsIPCID_NOTIFY_WARN() (r bool) { r = this == IPCID_NOTIFY_WARN; return }

// IsIPCID_NOTIFY_ERR returns whether the value of this `IpcIDs` equals `IPCID_NOTIFY_ERR`.
func (this IpcIDs) IsIPCID_NOTIFY_ERR() (r bool) { r = this == IPCID_NOTIFY_ERR; return }

// IsIPCID_PROJ_CHANGED returns whether the value of this `IpcIDs` equals `IPCID_PROJ_CHANGED`.
func (this IpcIDs) IsIPCID_PROJ_CHANGED() (r bool) { r = this == IPCID_PROJ_CHANGED; return }

// IsIPCID_PROJ_POLLEVTS returns whether the value of this `IpcIDs` equals `IPCID_PROJ_POLLEVTS`.
func (this IpcIDs) IsIPCID_PROJ_POLLEVTS() (r bool) { r = this == IPCID_PROJ_POLLEVTS; return }

// IsIPCID_SRCDIAG_LIST returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_LIST`.
func (this IpcIDs) IsIPCID_SRCDIAG_LIST() (r bool) { r = this == IPCID_SRCDIAG_LIST; return }

// IsIPCID_SRCDIAG_RUN_CURFILE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_CURFILE`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_CURFILE() (r bool) {
	r = this == IPCID_SRCDIAG_RUN_CURFILE
	return
}

// IsIPCID_SRCDIAG_RUN_OPENFILES returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_OPENFILES`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_OPENFILES() (r bool) {
	r = this == IPCID_SRCDIAG_RUN_OPENFILES
	return
}

// IsIPCID_SRCDIAG_RUN_ALLFILES returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_RUN_ALLFILES`.
func (this IpcIDs) IsIPCID_SRCDIAG_RUN_ALLFILES() (r bool) {
	r = this == IPCID_SRCDIAG_RUN_ALLFILES
	return
}

// IsIPCID_SRCDIAG_FORGETALL returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_FORGETALL`.
func (this IpcIDs) IsIPCID_SRCDIAG_FORGETALL() (r bool) { r = this == IPCID_SRCDIAG_FORGETALL; return }

// IsIPCID_SRCDIAG_PEEKHIDDEN returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_PEEKHIDDEN`.
func (this IpcIDs) IsIPCID_SRCDIAG_PEEKHIDDEN() (r bool) { r = this == IPCID_SRCDIAG_PEEKHIDDEN; return }

// IsIPCID_SRCDIAG_PUB returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_PUB`.
func (this IpcIDs) IsIPCID_SRCDIAG_PUB() (r bool) { r = this == IPCID_SRCDIAG_PUB; return }

// IsIPCID_SRCDIAG_AUTO_TOGGLE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_TOGGLE`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_TOGGLE() (r bool) {
	r = this == IPCID_SRCDIAG_AUTO_TOGGLE
	return
}

// IsIPCID_SRCDIAG_AUTO_ALL returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_ALL`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_ALL() (r bool) { r = this == IPCID_SRCDIAG_AUTO_ALL; return }

// IsIPCID_SRCDIAG_AUTO_NONE returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_AUTO_NONE`.
func (this IpcIDs) IsIPCID_SRCDIAG_AUTO_NONE() (r bool) { r = this == IPCID_SRCDIAG_AUTO_NONE; return }

// IsIPCID_SRCDIAG_STARTED returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_STARTED`.
func (this IpcIDs) IsIPCID_SRCDIAG_STARTED() (r bool) { r = this == IPCID_SRCDIAG_STARTED; return }

// IsIPCID_SRCDIAG_FINISHED returns whether the value of this `IpcIDs` equals `IPCID_SRCDIAG_FINISHED`.
func (this IpcIDs) IsIPCID_SRCDIAG_FINISHED() (r bool) { r = this == IPCID_SRCDIAG_FINISHED; return }

// IsIPCID_SRCMOD_FMT_SETDEFMENU returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_SETDEFMENU`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_SETDEFMENU() (r bool) {
	r = this == IPCID_SRCMOD_FMT_SETDEFMENU
	return
}

// IsIPCID_SRCMOD_FMT_SETDEFPICK returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_SETDEFPICK`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_SETDEFPICK() (r bool) {
	r = this == IPCID_SRCMOD_FMT_SETDEFPICK
	return
}

// IsIPCID_SRCMOD_FMT_RUNONFILE returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_RUNONFILE`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_RUNONFILE() (r bool) {
	r = this == IPCID_SRCMOD_FMT_RUNONFILE
	return
}

// IsIPCID_SRCMOD_FMT_RUNONSEL returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_FMT_RUNONSEL`.
func (this IpcIDs) IsIPCID_SRCMOD_FMT_RUNONSEL() (r bool) {
	r = this == IPCID_SRCMOD_FMT_RUNONSEL
	return
}

// IsIPCID_SRCMOD_RENAME returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_RENAME`.
func (this IpcIDs) IsIPCID_SRCMOD_RENAME() (r bool) { r = this == IPCID_SRCMOD_RENAME; return }

// IsIPCID_SRCMOD_ACTIONS returns whether the value of this `IpcIDs` equals `IPCID_SRCMOD_ACTIONS`.
func (this IpcIDs) IsIPCID_SRCMOD_ACTIONS() (r bool) { r = this == IPCID_SRCMOD_ACTIONS; return }

// IsIPCID_SRCINTEL_HOVER returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_HOVER`.
func (this IpcIDs) IsIPCID_SRCINTEL_HOVER() (r bool) { r = this == IPCID_SRCINTEL_HOVER; return }

// IsIPCID_SRCINTEL_SYMS_FILE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SYMS_FILE`.
func (this IpcIDs) IsIPCID_SRCINTEL_SYMS_FILE() (r bool) { r = this == IPCID_SRCINTEL_SYMS_FILE; return }

// IsIPCID_SRCINTEL_SYMS_PROJ returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SYMS_PROJ`.
func (this IpcIDs) IsIPCID_SRCINTEL_SYMS_PROJ() (r bool) { r = this == IPCID_SRCINTEL_SYMS_PROJ; return }

// IsIPCID_SRCINTEL_CMPL_ITEMS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_CMPL_ITEMS`.
func (this IpcIDs) IsIPCID_SRCINTEL_CMPL_ITEMS() (r bool) {
	r = this == IPCID_SRCINTEL_CMPL_ITEMS
	return
}

// IsIPCID_SRCINTEL_CMPL_DETAILS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_CMPL_DETAILS`.
func (this IpcIDs) IsIPCID_SRCINTEL_CMPL_DETAILS() (r bool) {
	r = this == IPCID_SRCINTEL_CMPL_DETAILS
	return
}

// IsIPCID_SRCINTEL_HIGHLIGHTS returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_HIGHLIGHTS`.
func (this IpcIDs) IsIPCID_SRCINTEL_HIGHLIGHTS() (r bool) {
	r = this == IPCID_SRCINTEL_HIGHLIGHTS
	return
}

// IsIPCID_SRCINTEL_SIGNATURE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_SIGNATURE`.
func (this IpcIDs) IsIPCID_SRCINTEL_SIGNATURE() (r bool) { r = this == IPCID_SRCINTEL_SIGNATURE; return }

// IsIPCID_SRCINTEL_REFERENCES returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_REFERENCES`.
func (this IpcIDs) IsIPCID_SRCINTEL_REFERENCES() (r bool) {
	r = this == IPCID_SRCINTEL_REFERENCES
	return
}

// IsIPCID_SRCINTEL_DEFSYM returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFSYM`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFSYM() (r bool) { r = this == IPCID_SRCINTEL_DEFSYM; return }

// IsIPCID_SRCINTEL_DEFTYPE returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFTYPE`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFTYPE() (r bool) { r = this == IPCID_SRCINTEL_DEFTYPE; return }

// IsIPCID_SRCINTEL_DEFIMPL returns whether the value of this `IpcIDs` equals `IPCID_SRCINTEL_DEFIMPL`.
func (this IpcIDs) IsIPCID_SRCINTEL_DEFIMPL() (r bool) { r = this == IPCID_SRCINTEL_DEFIMPL; return }

// IsIPCID_EXTRAS_INTEL_LIST returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_INTEL_LIST`.
func (this IpcIDs) IsIPCID_EXTRAS_INTEL_LIST() (r bool) { r = this == IPCID_EXTRAS_INTEL_LIST; return }

// IsIPCID_EXTRAS_INTEL_RUN returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_INTEL_RUN`.
func (this IpcIDs) IsIPCID_EXTRAS_INTEL_RUN() (r bool) { r = this == IPCID_EXTRAS_INTEL_RUN; return }

// IsIPCID_EXTRAS_QUERY_LIST returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_QUERY_LIST`.
func (this IpcIDs) IsIPCID_EXTRAS_QUERY_LIST() (r bool) { r = this == IPCID_EXTRAS_QUERY_LIST; return }

// IsIPCID_EXTRAS_QUERY_RUN returns whether the value of this `IpcIDs` equals `IPCID_EXTRAS_QUERY_RUN`.
func (this IpcIDs) IsIPCID_EXTRAS_QUERY_RUN() (r bool) { r = this == IPCID_EXTRAS_QUERY_RUN; return }

// Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN` (inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).
func (this IpcIDs) Valid() (r bool) {
	r = (this >= IPCID_MENUS_MAIN) && (this <= IPCID_EXTRAS_QUERY_RUN)
	return
}

// WellknownIpcIDses returns the `names` and `values` of all 49 well-known `IpcIDs` enumerants.
func WellknownIpcIDses() (names []string, values []IpcIDs) {
	names, values = []string{"IPCID_MENUS_MAIN", "IPCID_MENUS_PKGS", "IPCID_MENUS_TOOLS", "IPCID_OBJ_SNAPSHOT", "IPCID_PAGE_HTML", "IPCID_TREEVIEW_GETITEM", "IPCID_TREEVIEW_CHILDREN", "IPCID_TREEVIEW_CHANGED", "IPCID_CFG_RESETALL", "IPCID_CFG_LIST", "IPCID_CFG_SET", "IPCID_NOTIFY_INFO", "IPCID_NOTIFY_WARN", "IPCID_NOTIFY_ERR", "IPCID_PROJ_CHANGED", "IPCID_PROJ_POLLEVTS", "IPCID_SRCDIAG_LIST", "IPCID_SRCDIAG_RUN_CURFILE", "IPCID_SRCDIAG_RUN_OPENFILES", "IPCID_SRCDIAG_RUN_ALLFILES", "IPCID_SRCDIAG_FORGETALL", "IPCID_SRCDIAG_PEEKHIDDEN", "IPCID_SRCDIAG_PUB", "IPCID_SRCDIAG_AUTO_TOGGLE", "IPCID_SRCDIAG_AUTO_ALL", "IPCID_SRCDIAG_AUTO_NONE", "IPCID_SRCDIAG_STARTED", "IPCID_SRCDIAG_FINISHED", "IPCID_SRCMOD_FMT_SETDEFMENU", "IPCID_SRCMOD_FMT_SETDEFPICK", "IPCID_SRCMOD_FMT_RUNONFILE", "IPCID_SRCMOD_FMT_RUNONSEL", "IPCID_SRCMOD_RENAME", "IPCID_SRCMOD_ACTIONS", "IPCID_SRCINTEL_HOVER", "IPCID_SRCINTEL_SYMS_FILE", "IPCID_SRCINTEL_SYMS_PROJ", "IPCID_SRCINTEL_CMPL_ITEMS", "IPCID_SRCINTEL_CMPL_DETAILS", "IPCID_SRCINTEL_HIGHLIGHTS", "IPCID_SRCINTEL_SIGNATURE", "IPCID_SRCINTEL_REFERENCES", "IPCID_SRCINTEL_DEFSYM", "IPCID_SRCINTEL_DEFTYPE", "IPCID_SRCINTEL_DEFIMPL", "IPCID_EXTRAS_INTEL_LIST", "IPCID_EXTRAS_INTEL_RUN", "IPCID_EXTRAS_QUERY_LIST", "IPCID_EXTRAS_QUERY_RUN"}, []IpcIDs{IPCID_MENUS_MAIN, IPCID_MENUS_PKGS, IPCID_MENUS_TOOLS, IPCID_OBJ_SNAPSHOT, IPCID_PAGE_HTML, IPCID_TREEVIEW_GETITEM, IPCID_TREEVIEW_CHILDREN, IPCID_TREEVIEW_CHANGED, IPCID_CFG_RESETALL, IPCID_CFG_LIST, IPCID_CFG_SET, IPCID_NOTIFY_INFO, IPCID_NOTIFY_WARN, IPCID_NOTIFY_ERR, IPCID_PROJ_CHANGED, IPCID_PROJ_POLLEVTS, IPCID_SRCDIAG_LIST, IPCID_SRCDIAG_RUN_CURFILE, IPCID_SRCDIAG_RUN_OPENFILES, IPCID_SRCDIAG_RUN_ALLFILES, IPCID_SRCDIAG_FORGETALL, IPCID_SRCDIAG_PEEKHIDDEN, IPCID_SRCDIAG_PUB, IPCID_SRCDIAG_AUTO_TOGGLE, IPCID_SRCDIAG_AUTO_ALL, IPCID_SRCDIAG_AUTO_NONE, IPCID_SRCDIAG_STARTED, IPCID_SRCDIAG_FINISHED, IPCID_SRCMOD_FMT_SETDEFMENU, IPCID_SRCMOD_FMT_SETDEFPICK, IPCID_SRCMOD_FMT_RUNONFILE, IPCID_SRCMOD_FMT_RUNONSEL, IPCID_SRCMOD_RENAME, IPCID_SRCMOD_ACTIONS, IPCID_SRCINTEL_HOVER, IPCID_SRCINTEL_SYMS_FILE, IPCID_SRCINTEL_SYMS_PROJ, IPCID_SRCINTEL_CMPL_ITEMS, IPCID_SRCINTEL_CMPL_DETAILS, IPCID_SRCINTEL_HIGHLIGHTS, IPCID_SRCINTEL_SIGNATURE, IPCID_SRCINTEL_REFERENCES, IPCID_SRCINTEL_DEFSYM, IPCID_SRCINTEL_DEFTYPE, IPCID_SRCINTEL_DEFIMPL, IPCID_EXTRAS_INTEL_LIST, IPCID_EXTRAS_INTEL_RUN, IPCID_EXTRAS_QUERY_LIST, IPCID_EXTRAS_QUERY_RUN}
	return
}

// String implements the `fmt.Stringer` interface.
func (this IpcIDs) String() (r string) {
	if (this < IPCID_MENUS_MAIN) || (this > IPCID_EXTRAS_QUERY_RUN) {
		goto formatNum
	}
	switch this {
	case IPCID_MENUS_MAIN:
		r = "Ipcid·Menus·Main"
	case IPCID_MENUS_PKGS:
		r = "Ipcid·Menus·Pkgs"
	case IPCID_MENUS_TOOLS:
		r = "Ipcid·Menus·Tools"
	case IPCID_OBJ_SNAPSHOT:
		r = "Ipcid·Obj·Snapshot"
	case IPCID_PAGE_HTML:
		r = "Ipcid·Page·Html"
	case IPCID_TREEVIEW_GETITEM:
		r = "Ipcid·Treeview·Getitem"
	case IPCID_TREEVIEW_CHILDREN:
		r = "Ipcid·Treeview·Children"
	case IPCID_TREEVIEW_CHANGED:
		r = "Ipcid·Treeview·Changed"
	case IPCID_CFG_RESETALL:
		r = "Ipcid·Cfg·Resetall"
	case IPCID_CFG_LIST:
		r = "Ipcid·Cfg·List"
	case IPCID_CFG_SET:
		r = "Ipcid·Cfg·Set"
	case IPCID_NOTIFY_INFO:
		r = "Ipcid·Notify·Info"
	case IPCID_NOTIFY_WARN:
		r = "Ipcid·Notify·Warn"
	case IPCID_NOTIFY_ERR:
		r = "Ipcid·Notify·Err"
	case IPCID_PROJ_CHANGED:
		r = "Ipcid·Proj·Changed"
	case IPCID_PROJ_POLLEVTS:
		r = "Ipcid·Proj·Pollevts"
	case IPCID_SRCDIAG_LIST:
		r = "Ipcid·Srcdiag·List"
	case IPCID_SRCDIAG_RUN_CURFILE:
		r = "Ipcid·Srcdiag·Run·Curfile"
	case IPCID_SRCDIAG_RUN_OPENFILES:
		r = "Ipcid·Srcdiag·Run·Openfiles"
	case IPCID_SRCDIAG_RUN_ALLFILES:
		r = "Ipcid·Srcdiag·Run·Allfiles"
	case IPCID_SRCDIAG_FORGETALL:
		r = "Ipcid·Srcdiag·Forgetall"
	case IPCID_SRCDIAG_PEEKHIDDEN:
		r = "Ipcid·Srcdiag·Peekhidden"
	case IPCID_SRCDIAG_PUB:
		r = "Ipcid·Srcdiag·Pub"
	case IPCID_SRCDIAG_AUTO_TOGGLE:
		r = "Ipcid·Srcdiag·Auto·Toggle"
	case IPCID_SRCDIAG_AUTO_ALL:
		r = "Ipcid·Srcdiag·Auto·All"
	case IPCID_SRCDIAG_AUTO_NONE:
		r = "Ipcid·Srcdiag·Auto·None"
	case IPCID_SRCDIAG_STARTED:
		r = "Ipcid·Srcdiag·Started"
	case IPCID_SRCDIAG_FINISHED:
		r = "Ipcid·Srcdiag·Finished"
	case IPCID_SRCMOD_FMT_SETDEFMENU:
		r = "Ipcid·Srcmod·Fmt·Setdefmenu"
	case IPCID_SRCMOD_FMT_SETDEFPICK:
		r = "Ipcid·Srcmod·Fmt·Setdefpick"
	case IPCID_SRCMOD_FMT_RUNONFILE:
		r = "Ipcid·Srcmod·Fmt·Runonfile"
	case IPCID_SRCMOD_FMT_RUNONSEL:
		r = "Ipcid·Srcmod·Fmt·Runonsel"
	case IPCID_SRCMOD_RENAME:
		r = "Ipcid·Srcmod·Rename"
	case IPCID_SRCMOD_ACTIONS:
		r = "Ipcid·Srcmod·Actions"
	case IPCID_SRCINTEL_HOVER:
		r = "Ipcid·Srcintel·Hover"
	case IPCID_SRCINTEL_SYMS_FILE:
		r = "Ipcid·Srcintel·Syms·File"
	case IPCID_SRCINTEL_SYMS_PROJ:
		r = "Ipcid·Srcintel·Syms·Proj"
	case IPCID_SRCINTEL_CMPL_ITEMS:
		r = "Ipcid·Srcintel·Cmpl·Items"
	case IPCID_SRCINTEL_CMPL_DETAILS:
		r = "Ipcid·Srcintel·Cmpl·Details"
	case IPCID_SRCINTEL_HIGHLIGHTS:
		r = "Ipcid·Srcintel·Highlights"
	case IPCID_SRCINTEL_SIGNATURE:
		r = "Ipcid·Srcintel·Signature"
	case IPCID_SRCINTEL_REFERENCES:
		r = "Ipcid·Srcintel·References"
	case IPCID_SRCINTEL_DEFSYM:
		r = "Ipcid·Srcintel·Defsym"
	case IPCID_SRCINTEL_DEFTYPE:
		r = "Ipcid·Srcintel·Deftype"
	case IPCID_SRCINTEL_DEFIMPL:
		r = "Ipcid·Srcintel·Defimpl"
	case IPCID_EXTRAS_INTEL_LIST:
		r = "Ipcid·Extras·Intel·List"
	case IPCID_EXTRAS_INTEL_RUN:
		r = "Ipcid·Extras·Intel·Run"
	case IPCID_EXTRAS_QUERY_LIST:
		r = "Ipcid·Extras·Query·List"
	case IPCID_EXTRAS_QUERY_RUN:
		r = "Ipcid·Extras·Query·Run"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// IpcIDsFromString returns the `IpcIDs` represented by `s` (as returned by `IpcIDs.String`, but case-insensitively), or an `error` if none exists.
func IpcIDsFromString(s string) (this IpcIDs, err error) {
	if (len(s) < 15) || (len(s) > 30) || (!pkg__strings.EqualFold(s[0:7], "Ipcid·")) {
		goto tryParseNum
	}
	switch {
	case pkg__strings.EqualFold(s, "Ipcid·Menus·Main"):
		this = IPCID_MENUS_MAIN
	case pkg__strings.EqualFold(s, "Ipcid·Menus·Pkgs"):
		this = IPCID_MENUS_PKGS
	case pkg__strings.EqualFold(s, "Ipcid·Menus·Tools"):
		this = IPCID_MENUS_TOOLS
	case pkg__strings.EqualFold(s, "Ipcid·Obj·Snapshot"):
		this = IPCID_OBJ_SNAPSHOT
	case pkg__strings.EqualFold(s, "Ipcid·Page·Html"):
		this = IPCID_PAGE_HTML
	case pkg__strings.EqualFold(s, "Ipcid·Treeview·Getitem"):
		this = IPCID_TREEVIEW_GETITEM
	case pkg__strings.EqualFold(s, "Ipcid·Treeview·Children"):
		this = IPCID_TREEVIEW_CHILDREN
	case pkg__strings.EqualFold(s, "Ipcid·Treeview·Changed"):
		this = IPCID_TREEVIEW_CHANGED
	case pkg__strings.EqualFold(s, "Ipcid·Cfg·Resetall"):
		this = IPCID_CFG_RESETALL
	case pkg__strings.EqualFold(s, "Ipcid·Cfg·List"):
		this = IPCID_CFG_LIST
	case pkg__strings.EqualFold(s, "Ipcid·Cfg·Set"):
		this = IPCID_CFG_SET
	case pkg__strings.EqualFold(s, "Ipcid·Notify·Info"):
		this = IPCID_NOTIFY_INFO
	case pkg__strings.EqualFold(s, "Ipcid·Notify·Warn"):
		this = IPCID_NOTIFY_WARN
	case pkg__strings.EqualFold(s, "Ipcid·Notify·Err"):
		this = IPCID_NOTIFY_ERR
	case pkg__strings.EqualFold(s, "Ipcid·Proj·Changed"):
		this = IPCID_PROJ_CHANGED
	case pkg__strings.EqualFold(s, "Ipcid·Proj·Pollevts"):
		this = IPCID_PROJ_POLLEVTS
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·List"):
		this = IPCID_SRCDIAG_LIST
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Run·Curfile"):
		this = IPCID_SRCDIAG_RUN_CURFILE
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Run·Openfiles"):
		this = IPCID_SRCDIAG_RUN_OPENFILES
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Run·Allfiles"):
		this = IPCID_SRCDIAG_RUN_ALLFILES
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Forgetall"):
		this = IPCID_SRCDIAG_FORGETALL
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Peekhidden"):
		this = IPCID_SRCDIAG_PEEKHIDDEN
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Pub"):
		this = IPCID_SRCDIAG_PUB
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Auto·Toggle"):
		this = IPCID_SRCDIAG_AUTO_TOGGLE
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Auto·All"):
		this = IPCID_SRCDIAG_AUTO_ALL
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Auto·None"):
		this = IPCID_SRCDIAG_AUTO_NONE
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Started"):
		this = IPCID_SRCDIAG_STARTED
	case pkg__strings.EqualFold(s, "Ipcid·Srcdiag·Finished"):
		this = IPCID_SRCDIAG_FINISHED
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Fmt·Setdefmenu"):
		this = IPCID_SRCMOD_FMT_SETDEFMENU
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Fmt·Setdefpick"):
		this = IPCID_SRCMOD_FMT_SETDEFPICK
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Fmt·Runonfile"):
		this = IPCID_SRCMOD_FMT_RUNONFILE
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Fmt·Runonsel"):
		this = IPCID_SRCMOD_FMT_RUNONSEL
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Rename"):
		this = IPCID_SRCMOD_RENAME
	case pkg__strings.EqualFold(s, "Ipcid·Srcmod·Actions"):
		this = IPCID_SRCMOD_ACTIONS
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Hover"):
		this = IPCID_SRCINTEL_HOVER
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Syms·File"):
		this = IPCID_SRCINTEL_SYMS_FILE
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Syms·Proj"):
		this = IPCID_SRCINTEL_SYMS_PROJ
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Cmpl·Items"):
		this = IPCID_SRCINTEL_CMPL_ITEMS
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Cmpl·Details"):
		this = IPCID_SRCINTEL_CMPL_DETAILS
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Highlights"):
		this = IPCID_SRCINTEL_HIGHLIGHTS
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Signature"):
		this = IPCID_SRCINTEL_SIGNATURE
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·References"):
		this = IPCID_SRCINTEL_REFERENCES
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Defsym"):
		this = IPCID_SRCINTEL_DEFSYM
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Deftype"):
		this = IPCID_SRCINTEL_DEFTYPE
	case pkg__strings.EqualFold(s, "Ipcid·Srcintel·Defimpl"):
		this = IPCID_SRCINTEL_DEFIMPL
	case pkg__strings.EqualFold(s, "Ipcid·Extras·Intel·List"):
		this = IPCID_EXTRAS_INTEL_LIST
	case pkg__strings.EqualFold(s, "Ipcid·Extras·Intel·Run"):
		this = IPCID_EXTRAS_INTEL_RUN
	case pkg__strings.EqualFold(s, "Ipcid·Extras·Query·List"):
		this = IPCID_EXTRAS_QUERY_LIST
	case pkg__strings.EqualFold(s, "Ipcid·Extras·Query·Run"):
		this = IPCID_EXTRAS_QUERY_RUN
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (IpcIDs)(v)
	}
	return
}

// IpcIDsFromStringOr is like `IpcIDsFromString` but returns `fallback` for unrecognized inputs.
func IpcIDsFromStringOr(s string, fallback IpcIDs) (this IpcIDs) {
	maybeIpcIDs, err := IpcIDsFromString(s)
	if err == nil {
		this = maybeIpcIDs
	} else {
		this = fallback
	}
	return
}

// GoString implements the `fmt.GoStringer` interface.
func (this IpcIDs) GoString() (r string) {
	if (this < IPCID_MENUS_MAIN) || (this > IPCID_EXTRAS_QUERY_RUN) {
		goto formatNum
	}
	switch this {
	case IPCID_MENUS_MAIN:
		r = "IPCID_MENUS_MAIN"
	case IPCID_MENUS_PKGS:
		r = "IPCID_MENUS_PKGS"
	case IPCID_MENUS_TOOLS:
		r = "IPCID_MENUS_TOOLS"
	case IPCID_OBJ_SNAPSHOT:
		r = "IPCID_OBJ_SNAPSHOT"
	case IPCID_PAGE_HTML:
		r = "IPCID_PAGE_HTML"
	case IPCID_TREEVIEW_GETITEM:
		r = "IPCID_TREEVIEW_GETITEM"
	case IPCID_TREEVIEW_CHILDREN:
		r = "IPCID_TREEVIEW_CHILDREN"
	case IPCID_TREEVIEW_CHANGED:
		r = "IPCID_TREEVIEW_CHANGED"
	case IPCID_CFG_RESETALL:
		r = "IPCID_CFG_RESETALL"
	case IPCID_CFG_LIST:
		r = "IPCID_CFG_LIST"
	case IPCID_CFG_SET:
		r = "IPCID_CFG_SET"
	case IPCID_NOTIFY_INFO:
		r = "IPCID_NOTIFY_INFO"
	case IPCID_NOTIFY_WARN:
		r = "IPCID_NOTIFY_WARN"
	case IPCID_NOTIFY_ERR:
		r = "IPCID_NOTIFY_ERR"
	case IPCID_PROJ_CHANGED:
		r = "IPCID_PROJ_CHANGED"
	case IPCID_PROJ_POLLEVTS:
		r = "IPCID_PROJ_POLLEVTS"
	case IPCID_SRCDIAG_LIST:
		r = "IPCID_SRCDIAG_LIST"
	case IPCID_SRCDIAG_RUN_CURFILE:
		r = "IPCID_SRCDIAG_RUN_CURFILE"
	case IPCID_SRCDIAG_RUN_OPENFILES:
		r = "IPCID_SRCDIAG_RUN_OPENFILES"
	case IPCID_SRCDIAG_RUN_ALLFILES:
		r = "IPCID_SRCDIAG_RUN_ALLFILES"
	case IPCID_SRCDIAG_FORGETALL:
		r = "IPCID_SRCDIAG_FORGETALL"
	case IPCID_SRCDIAG_PEEKHIDDEN:
		r = "IPCID_SRCDIAG_PEEKHIDDEN"
	case IPCID_SRCDIAG_PUB:
		r = "IPCID_SRCDIAG_PUB"
	case IPCID_SRCDIAG_AUTO_TOGGLE:
		r = "IPCID_SRCDIAG_AUTO_TOGGLE"
	case IPCID_SRCDIAG_AUTO_ALL:
		r = "IPCID_SRCDIAG_AUTO_ALL"
	case IPCID_SRCDIAG_AUTO_NONE:
		r = "IPCID_SRCDIAG_AUTO_NONE"
	case IPCID_SRCDIAG_STARTED:
		r = "IPCID_SRCDIAG_STARTED"
	case IPCID_SRCDIAG_FINISHED:
		r = "IPCID_SRCDIAG_FINISHED"
	case IPCID_SRCMOD_FMT_SETDEFMENU:
		r = "IPCID_SRCMOD_FMT_SETDEFMENU"
	case IPCID_SRCMOD_FMT_SETDEFPICK:
		r = "IPCID_SRCMOD_FMT_SETDEFPICK"
	case IPCID_SRCMOD_FMT_RUNONFILE:
		r = "IPCID_SRCMOD_FMT_RUNONFILE"
	case IPCID_SRCMOD_FMT_RUNONSEL:
		r = "IPCID_SRCMOD_FMT_RUNONSEL"
	case IPCID_SRCMOD_RENAME:
		r = "IPCID_SRCMOD_RENAME"
	case IPCID_SRCMOD_ACTIONS:
		r = "IPCID_SRCMOD_ACTIONS"
	case IPCID_SRCINTEL_HOVER:
		r = "IPCID_SRCINTEL_HOVER"
	case IPCID_SRCINTEL_SYMS_FILE:
		r = "IPCID_SRCINTEL_SYMS_FILE"
	case IPCID_SRCINTEL_SYMS_PROJ:
		r = "IPCID_SRCINTEL_SYMS_PROJ"
	case IPCID_SRCINTEL_CMPL_ITEMS:
		r = "IPCID_SRCINTEL_CMPL_ITEMS"
	case IPCID_SRCINTEL_CMPL_DETAILS:
		r = "IPCID_SRCINTEL_CMPL_DETAILS"
	case IPCID_SRCINTEL_HIGHLIGHTS:
		r = "IPCID_SRCINTEL_HIGHLIGHTS"
	case IPCID_SRCINTEL_SIGNATURE:
		r = "IPCID_SRCINTEL_SIGNATURE"
	case IPCID_SRCINTEL_REFERENCES:
		r = "IPCID_SRCINTEL_REFERENCES"
	case IPCID_SRCINTEL_DEFSYM:
		r = "IPCID_SRCINTEL_DEFSYM"
	case IPCID_SRCINTEL_DEFTYPE:
		r = "IPCID_SRCINTEL_DEFTYPE"
	case IPCID_SRCINTEL_DEFIMPL:
		r = "IPCID_SRCINTEL_DEFIMPL"
	case IPCID_EXTRAS_INTEL_LIST:
		r = "IPCID_EXTRAS_INTEL_LIST"
	case IPCID_EXTRAS_INTEL_RUN:
		r = "IPCID_EXTRAS_INTEL_RUN"
	case IPCID_EXTRAS_QUERY_LIST:
		r = "IPCID_EXTRAS_QUERY_LIST"
	case IPCID_EXTRAS_QUERY_RUN:
		r = "IPCID_EXTRAS_QUERY_RUN"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// IpcIDsFromGoString returns the `IpcIDs` represented by `s` (as returned by `IpcIDs.GoString`, and case-sensitively), or an `error` if none exists.
func IpcIDsFromGoString(s string) (this IpcIDs, err error) {
	if (len(s) < 13) || (len(s) > 27) || (s[0:6] != "IPCID_") {
		goto tryParseNum
	}
	switch s {
	case "IPCID_MENUS_MAIN":
		this = IPCID_MENUS_MAIN
	case "IPCID_MENUS_PKGS":
		this = IPCID_MENUS_PKGS
	case "IPCID_MENUS_TOOLS":
		this = IPCID_MENUS_TOOLS
	case "IPCID_OBJ_SNAPSHOT":
		this = IPCID_OBJ_SNAPSHOT
	case "IPCID_PAGE_HTML":
		this = IPCID_PAGE_HTML
	case "IPCID_TREEVIEW_GETITEM":
		this = IPCID_TREEVIEW_GETITEM
	case "IPCID_TREEVIEW_CHILDREN":
		this = IPCID_TREEVIEW_CHILDREN
	case "IPCID_TREEVIEW_CHANGED":
		this = IPCID_TREEVIEW_CHANGED
	case "IPCID_CFG_RESETALL":
		this = IPCID_CFG_RESETALL
	case "IPCID_CFG_LIST":
		this = IPCID_CFG_LIST
	case "IPCID_CFG_SET":
		this = IPCID_CFG_SET
	case "IPCID_NOTIFY_INFO":
		this = IPCID_NOTIFY_INFO
	case "IPCID_NOTIFY_WARN":
		this = IPCID_NOTIFY_WARN
	case "IPCID_NOTIFY_ERR":
		this = IPCID_NOTIFY_ERR
	case "IPCID_PROJ_CHANGED":
		this = IPCID_PROJ_CHANGED
	case "IPCID_PROJ_POLLEVTS":
		this = IPCID_PROJ_POLLEVTS
	case "IPCID_SRCDIAG_LIST":
		this = IPCID_SRCDIAG_LIST
	case "IPCID_SRCDIAG_RUN_CURFILE":
		this = IPCID_SRCDIAG_RUN_CURFILE
	case "IPCID_SRCDIAG_RUN_OPENFILES":
		this = IPCID_SRCDIAG_RUN_OPENFILES
	case "IPCID_SRCDIAG_RUN_ALLFILES":
		this = IPCID_SRCDIAG_RUN_ALLFILES
	case "IPCID_SRCDIAG_FORGETALL":
		this = IPCID_SRCDIAG_FORGETALL
	case "IPCID_SRCDIAG_PEEKHIDDEN":
		this = IPCID_SRCDIAG_PEEKHIDDEN
	case "IPCID_SRCDIAG_PUB":
		this = IPCID_SRCDIAG_PUB
	case "IPCID_SRCDIAG_AUTO_TOGGLE":
		this = IPCID_SRCDIAG_AUTO_TOGGLE
	case "IPCID_SRCDIAG_AUTO_ALL":
		this = IPCID_SRCDIAG_AUTO_ALL
	case "IPCID_SRCDIAG_AUTO_NONE":
		this = IPCID_SRCDIAG_AUTO_NONE
	case "IPCID_SRCDIAG_STARTED":
		this = IPCID_SRCDIAG_STARTED
	case "IPCID_SRCDIAG_FINISHED":
		this = IPCID_SRCDIAG_FINISHED
	case "IPCID_SRCMOD_FMT_SETDEFMENU":
		this = IPCID_SRCMOD_FMT_SETDEFMENU
	case "IPCID_SRCMOD_FMT_SETDEFPICK":
		this = IPCID_SRCMOD_FMT_SETDEFPICK
	case "IPCID_SRCMOD_FMT_RUNONFILE":
		this = IPCID_SRCMOD_FMT_RUNONFILE
	case "IPCID_SRCMOD_FMT_RUNONSEL":
		this = IPCID_SRCMOD_FMT_RUNONSEL
	case "IPCID_SRCMOD_RENAME":
		this = IPCID_SRCMOD_RENAME
	case "IPCID_SRCMOD_ACTIONS":
		this = IPCID_SRCMOD_ACTIONS
	case "IPCID_SRCINTEL_HOVER":
		this = IPCID_SRCINTEL_HOVER
	case "IPCID_SRCINTEL_SYMS_FILE":
		this = IPCID_SRCINTEL_SYMS_FILE
	case "IPCID_SRCINTEL_SYMS_PROJ":
		this = IPCID_SRCINTEL_SYMS_PROJ
	case "IPCID_SRCINTEL_CMPL_ITEMS":
		this = IPCID_SRCINTEL_CMPL_ITEMS
	case "IPCID_SRCINTEL_CMPL_DETAILS":
		this = IPCID_SRCINTEL_CMPL_DETAILS
	case "IPCID_SRCINTEL_HIGHLIGHTS":
		this = IPCID_SRCINTEL_HIGHLIGHTS
	case "IPCID_SRCINTEL_SIGNATURE":
		this = IPCID_SRCINTEL_SIGNATURE
	case "IPCID_SRCINTEL_REFERENCES":
		this = IPCID_SRCINTEL_REFERENCES
	case "IPCID_SRCINTEL_DEFSYM":
		this = IPCID_SRCINTEL_DEFSYM
	case "IPCID_SRCINTEL_DEFTYPE":
		this = IPCID_SRCINTEL_DEFTYPE
	case "IPCID_SRCINTEL_DEFIMPL":
		this = IPCID_SRCINTEL_DEFIMPL
	case "IPCID_EXTRAS_INTEL_LIST":
		this = IPCID_EXTRAS_INTEL_LIST
	case "IPCID_EXTRAS_INTEL_RUN":
		this = IPCID_EXTRAS_INTEL_RUN
	case "IPCID_EXTRAS_QUERY_LIST":
		this = IPCID_EXTRAS_QUERY_LIST
	case "IPCID_EXTRAS_QUERY_RUN":
		this = IPCID_EXTRAS_QUERY_RUN
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (IpcIDs)(v)
	}
	return
}

// IpcIDsFromGoStringOr is like `IpcIDsFromGoString` but returns `fallback` for unrecognized inputs.
func IpcIDsFromGoStringOr(s string, fallback IpcIDs) (this IpcIDs) {
	maybeIpcIDs, err := IpcIDsFromGoString(s)
	if err == nil {
		this = maybeIpcIDs
	} else {
		this = fallback
	}
	return
}

// IsDIAG_SEV_ERR returns whether the value of this `DiagSeverity` equals `DIAG_SEV_ERR`.
func (this DiagSeverity) IsDIAG_SEV_ERR() (r bool) { r = this == DIAG_SEV_ERR; return }

// IsDIAG_SEV_WARN returns whether the value of this `DiagSeverity` equals `DIAG_SEV_WARN`.
func (this DiagSeverity) IsDIAG_SEV_WARN() (r bool) { r = this == DIAG_SEV_WARN; return }

// IsDIAG_SEV_INFO returns whether the value of this `DiagSeverity` equals `DIAG_SEV_INFO`.
func (this DiagSeverity) IsDIAG_SEV_INFO() (r bool) { r = this == DIAG_SEV_INFO; return }

// IsDIAG_SEV_HINT returns whether the value of this `DiagSeverity` equals `DIAG_SEV_HINT`.
func (this DiagSeverity) IsDIAG_SEV_HINT() (r bool) { r = this == DIAG_SEV_HINT; return }

// Valid returns whether the value of this `DiagSeverity` is between `DIAG_SEV_ERR` (inclusive) and `DIAG_SEV_HINT` (inclusive).
func (this DiagSeverity) Valid() (r bool) {
	r = (this >= DIAG_SEV_ERR) && (this <= DIAG_SEV_HINT)
	return
}

// WellknownDiagSeverities returns the `names` and `values` of all 4 well-known `DiagSeverity` enumerants.
func WellknownDiagSeverities() (names []string, values []DiagSeverity) {
	names, values = []string{"DIAG_SEV_ERR", "DIAG_SEV_WARN", "DIAG_SEV_INFO", "DIAG_SEV_HINT"}, []DiagSeverity{DIAG_SEV_ERR, DIAG_SEV_WARN, DIAG_SEV_INFO, DIAG_SEV_HINT}
	return
}

// String implements the `fmt.Stringer` interface.
func (this DiagSeverity) String() (r string) {
	if (this < DIAG_SEV_ERR) || (this > DIAG_SEV_HINT) {
		goto formatNum
	}
	switch this {
	case DIAG_SEV_ERR:
		r = "Diag·Sev·Err"
	case DIAG_SEV_WARN:
		r = "Diag·Sev·Warn"
	case DIAG_SEV_INFO:
		r = "Diag·Sev·Info"
	case DIAG_SEV_HINT:
		r = "Diag·Sev·Hint"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.Itoa((int)(this))
	return
}

// DiagSeverityFromString returns the `DiagSeverity` represented by `s` (as returned by `DiagSeverity.String`, but case-insensitively), or an `error` if none exists.
func DiagSeverityFromString(s string) (this DiagSeverity, err error) {
	if (len(s) < 14) || (len(s) > 15) || (!pkg__strings.EqualFold(s[0:11], "Diag·Sev·")) {
		goto tryParseNum
	}
	switch {
	case pkg__strings.EqualFold(s, "Diag·Sev·Err"):
		this = DIAG_SEV_ERR
	case pkg__strings.EqualFold(s, "Diag·Sev·Warn"):
		this = DIAG_SEV_WARN
	case pkg__strings.EqualFold(s, "Diag·Sev·Info"):
		this = DIAG_SEV_INFO
	case pkg__strings.EqualFold(s, "Diag·Sev·Hint"):
		this = DIAG_SEV_HINT
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v int
	v, err = pkg__strconv.Atoi(s)
	if err == nil {
		this = (DiagSeverity)(v)
	}
	return
}

// DiagSeverityFromStringOr is like `DiagSeverityFromString` but returns `fallback` for unrecognized inputs.
func DiagSeverityFromStringOr(s string, fallback DiagSeverity) (this DiagSeverity) {
	maybeDiagSeverity, err := DiagSeverityFromString(s)
	if err == nil {
		this = maybeDiagSeverity
	} else {
		this = fallback
	}
	return
}

// GoString implements the `fmt.GoStringer` interface.
func (this DiagSeverity) GoString() (r string) {
	if (this < DIAG_SEV_ERR) || (this > DIAG_SEV_HINT) {
		goto formatNum
	}
	switch this {
	case DIAG_SEV_ERR:
		r = "DIAG_SEV_ERR"
	case DIAG_SEV_WARN:
		r = "DIAG_SEV_WARN"
	case DIAG_SEV_INFO:
		r = "DIAG_SEV_INFO"
	case DIAG_SEV_HINT:
		r = "DIAG_SEV_HINT"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.Itoa((int)(this))
	return
}

// DiagSeverityFromGoString returns the `DiagSeverity` represented by `s` (as returned by `DiagSeverity.GoString`, and case-sensitively), or an `error` if none exists.
func DiagSeverityFromGoString(s string) (this DiagSeverity, err error) {
	if (len(s) < 12) || (len(s) > 13) || (s[0:9] != "DIAG_SEV_") {
		goto tryParseNum
	}
	switch s {
	case "DIAG_SEV_ERR":
		this = DIAG_SEV_ERR
	case "DIAG_SEV_WARN":
		this = DIAG_SEV_WARN
	case "DIAG_SEV_INFO":
		this = DIAG_SEV_INFO
	case "DIAG_SEV_HINT":
		this = DIAG_SEV_HINT
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v int
	v, err = pkg__strconv.Atoi(s)
	if err == nil {
		this = (DiagSeverity)(v)
	}
	return
}

// DiagSeverityFromGoStringOr is like `DiagSeverityFromGoString` but returns `fallback` for unrecognized inputs.
func DiagSeverityFromGoStringOr(s string, fallback DiagSeverity) (this DiagSeverity) {
	maybeDiagSeverity, err := DiagSeverityFromGoString(s)
	if err == nil {
		this = maybeDiagSeverity
	} else {
		this = fallback
	}
	return
}

// IsSYM_FILE returns whether the value of this `Symbol` equals `SYM_FILE`.
func (this Symbol) IsSYM_FILE() (r bool) { r = this == SYM_FILE; return }

// IsSYM_MODULE returns whether the value of this `Symbol` equals `SYM_MODULE`.
func (this Symbol) IsSYM_MODULE() (r bool) { r = this == SYM_MODULE; return }

// IsSYM_NAMESPACE returns whether the value of this `Symbol` equals `SYM_NAMESPACE`.
func (this Symbol) IsSYM_NAMESPACE() (r bool) { r = this == SYM_NAMESPACE; return }

// IsSYM_PACKAGE returns whether the value of this `Symbol` equals `SYM_PACKAGE`.
func (this Symbol) IsSYM_PACKAGE() (r bool) { r = this == SYM_PACKAGE; return }

// IsSYM_CLASS returns whether the value of this `Symbol` equals `SYM_CLASS`.
func (this Symbol) IsSYM_CLASS() (r bool) { r = this == SYM_CLASS; return }

// IsSYM_METHOD returns whether the value of this `Symbol` equals `SYM_METHOD`.
func (this Symbol) IsSYM_METHOD() (r bool) { r = this == SYM_METHOD; return }

// IsSYM_PROPERTY returns whether the value of this `Symbol` equals `SYM_PROPERTY`.
func (this Symbol) IsSYM_PROPERTY() (r bool) { r = this == SYM_PROPERTY; return }

// IsSYM_FIELD returns whether the value of this `Symbol` equals `SYM_FIELD`.
func (this Symbol) IsSYM_FIELD() (r bool) { r = this == SYM_FIELD; return }

// IsSYM_CONSTRUCTOR returns whether the value of this `Symbol` equals `SYM_CONSTRUCTOR`.
func (this Symbol) IsSYM_CONSTRUCTOR() (r bool) { r = this == SYM_CONSTRUCTOR; return }

// IsSYM_ENUM returns whether the value of this `Symbol` equals `SYM_ENUM`.
func (this Symbol) IsSYM_ENUM() (r bool) { r = this == SYM_ENUM; return }

// IsSYM_INTERFACE returns whether the value of this `Symbol` equals `SYM_INTERFACE`.
func (this Symbol) IsSYM_INTERFACE() (r bool) { r = this == SYM_INTERFACE; return }

// IsSYM_FUNCTION returns whether the value of this `Symbol` equals `SYM_FUNCTION`.
func (this Symbol) IsSYM_FUNCTION() (r bool) { r = this == SYM_FUNCTION; return }

// IsSYM_VARIABLE returns whether the value of this `Symbol` equals `SYM_VARIABLE`.
func (this Symbol) IsSYM_VARIABLE() (r bool) { r = this == SYM_VARIABLE; return }

// IsSYM_CONSTANT returns whether the value of this `Symbol` equals `SYM_CONSTANT`.
func (this Symbol) IsSYM_CONSTANT() (r bool) { r = this == SYM_CONSTANT; return }

// IsSYM_STRING returns whether the value of this `Symbol` equals `SYM_STRING`.
func (this Symbol) IsSYM_STRING() (r bool) { r = this == SYM_STRING; return }

// IsSYM_NUMBER returns whether the value of this `Symbol` equals `SYM_NUMBER`.
func (this Symbol) IsSYM_NUMBER() (r bool) { r = this == SYM_NUMBER; return }

// IsSYM_BOOLEAN returns whether the value of this `Symbol` equals `SYM_BOOLEAN`.
func (this Symbol) IsSYM_BOOLEAN() (r bool) { r = this == SYM_BOOLEAN; return }

// IsSYM_ARRAY returns whether the value of this `Symbol` equals `SYM_ARRAY`.
func (this Symbol) IsSYM_ARRAY() (r bool) { r = this == SYM_ARRAY; return }

// IsSYM_OBJECT returns whether the value of this `Symbol` equals `SYM_OBJECT`.
func (this Symbol) IsSYM_OBJECT() (r bool) { r = this == SYM_OBJECT; return }

// IsSYM_KEY returns whether the value of this `Symbol` equals `SYM_KEY`.
func (this Symbol) IsSYM_KEY() (r bool) { r = this == SYM_KEY; return }

// IsSYM_NULL returns whether the value of this `Symbol` equals `SYM_NULL`.
func (this Symbol) IsSYM_NULL() (r bool) { r = this == SYM_NULL; return }

// IsSYM_ENUMMEMBER returns whether the value of this `Symbol` equals `SYM_ENUMMEMBER`.
func (this Symbol) IsSYM_ENUMMEMBER() (r bool) { r = this == SYM_ENUMMEMBER; return }

// IsSYM_STRUCT returns whether the value of this `Symbol` equals `SYM_STRUCT`.
func (this Symbol) IsSYM_STRUCT() (r bool) { r = this == SYM_STRUCT; return }

// IsSYM_EVENT returns whether the value of this `Symbol` equals `SYM_EVENT`.
func (this Symbol) IsSYM_EVENT() (r bool) { r = this == SYM_EVENT; return }

// IsSYM_OPERATOR returns whether the value of this `Symbol` equals `SYM_OPERATOR`.
func (this Symbol) IsSYM_OPERATOR() (r bool) { r = this == SYM_OPERATOR; return }

// IsSYM_TYPEPARAMETER returns whether the value of this `Symbol` equals `SYM_TYPEPARAMETER`.
func (this Symbol) IsSYM_TYPEPARAMETER() (r bool) { r = this == SYM_TYPEPARAMETER; return }

// Valid returns whether the value of this `Symbol` is between `SYM_FILE` (inclusive) and `SYM_TYPEPARAMETER` (inclusive).
func (this Symbol) Valid() (r bool) { r = (this >= SYM_FILE) && (this <= SYM_TYPEPARAMETER); return }

// WellknownSymbols returns the `names` and `values` of all 26 well-known `Symbol` enumerants.
func WellknownSymbols() (names []string, values []Symbol) {
	names, values = []string{"SYM_FILE", "SYM_MODULE", "SYM_NAMESPACE", "SYM_PACKAGE", "SYM_CLASS", "SYM_METHOD", "SYM_PROPERTY", "SYM_FIELD", "SYM_CONSTRUCTOR", "SYM_ENUM", "SYM_INTERFACE", "SYM_FUNCTION", "SYM_VARIABLE", "SYM_CONSTANT", "SYM_STRING", "SYM_NUMBER", "SYM_BOOLEAN", "SYM_ARRAY", "SYM_OBJECT", "SYM_KEY", "SYM_NULL", "SYM_ENUMMEMBER", "SYM_STRUCT", "SYM_EVENT", "SYM_OPERATOR", "SYM_TYPEPARAMETER"}, []Symbol{SYM_FILE, SYM_MODULE, SYM_NAMESPACE, SYM_PACKAGE, SYM_CLASS, SYM_METHOD, SYM_PROPERTY, SYM_FIELD, SYM_CONSTRUCTOR, SYM_ENUM, SYM_INTERFACE, SYM_FUNCTION, SYM_VARIABLE, SYM_CONSTANT, SYM_STRING, SYM_NUMBER, SYM_BOOLEAN, SYM_ARRAY, SYM_OBJECT, SYM_KEY, SYM_NULL, SYM_ENUMMEMBER, SYM_STRUCT, SYM_EVENT, SYM_OPERATOR, SYM_TYPEPARAMETER}
	return
}

// String implements the `fmt.Stringer` interface.
func (this Symbol) String() (r string) {
	if (this < SYM_FILE) || (this > SYM_TYPEPARAMETER) {
		goto formatNum
	}
	switch this {
	case SYM_FILE:
		r = "Sym·File"
	case SYM_MODULE:
		r = "Sym·Module"
	case SYM_NAMESPACE:
		r = "Sym·Namespace"
	case SYM_PACKAGE:
		r = "Sym·Package"
	case SYM_CLASS:
		r = "Sym·Class"
	case SYM_METHOD:
		r = "Sym·Method"
	case SYM_PROPERTY:
		r = "Sym·Property"
	case SYM_FIELD:
		r = "Sym·Field"
	case SYM_CONSTRUCTOR:
		r = "Sym·Constructor"
	case SYM_ENUM:
		r = "Sym·Enum"
	case SYM_INTERFACE:
		r = "Sym·Interface"
	case SYM_FUNCTION:
		r = "Sym·Function"
	case SYM_VARIABLE:
		r = "Sym·Variable"
	case SYM_CONSTANT:
		r = "Sym·Constant"
	case SYM_STRING:
		r = "Sym·String"
	case SYM_NUMBER:
		r = "Sym·Number"
	case SYM_BOOLEAN:
		r = "Sym·Boolean"
	case SYM_ARRAY:
		r = "Sym·Array"
	case SYM_OBJECT:
		r = "Sym·Object"
	case SYM_KEY:
		r = "Sym·Key"
	case SYM_NULL:
		r = "Sym·Null"
	case SYM_ENUMMEMBER:
		r = "Sym·Enummember"
	case SYM_STRUCT:
		r = "Sym·Struct"
	case SYM_EVENT:
		r = "Sym·Event"
	case SYM_OPERATOR:
		r = "Sym·Operator"
	case SYM_TYPEPARAMETER:
		r = "Sym·Typeparameter"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// SymbolFromString returns the `Symbol` represented by `s` (as returned by `Symbol.String`, but case-insensitively), or an `error` if none exists.
func SymbolFromString(s string) (this Symbol, err error) {
	if (len(s) < 8) || (len(s) > 18) || (!pkg__strings.EqualFold(s[0:5], "Sym·")) {
		goto tryParseNum
	}
	switch {
	case pkg__strings.EqualFold(s, "Sym·File"):
		this = SYM_FILE
	case pkg__strings.EqualFold(s, "Sym·Module"):
		this = SYM_MODULE
	case pkg__strings.EqualFold(s, "Sym·Namespace"):
		this = SYM_NAMESPACE
	case pkg__strings.EqualFold(s, "Sym·Package"):
		this = SYM_PACKAGE
	case pkg__strings.EqualFold(s, "Sym·Class"):
		this = SYM_CLASS
	case pkg__strings.EqualFold(s, "Sym·Method"):
		this = SYM_METHOD
	case pkg__strings.EqualFold(s, "Sym·Property"):
		this = SYM_PROPERTY
	case pkg__strings.EqualFold(s, "Sym·Field"):
		this = SYM_FIELD
	case pkg__strings.EqualFold(s, "Sym·Constructor"):
		this = SYM_CONSTRUCTOR
	case pkg__strings.EqualFold(s, "Sym·Enum"):
		this = SYM_ENUM
	case pkg__strings.EqualFold(s, "Sym·Interface"):
		this = SYM_INTERFACE
	case pkg__strings.EqualFold(s, "Sym·Function"):
		this = SYM_FUNCTION
	case pkg__strings.EqualFold(s, "Sym·Variable"):
		this = SYM_VARIABLE
	case pkg__strings.EqualFold(s, "Sym·Constant"):
		this = SYM_CONSTANT
	case pkg__strings.EqualFold(s, "Sym·String"):
		this = SYM_STRING
	case pkg__strings.EqualFold(s, "Sym·Number"):
		this = SYM_NUMBER
	case pkg__strings.EqualFold(s, "Sym·Boolean"):
		this = SYM_BOOLEAN
	case pkg__strings.EqualFold(s, "Sym·Array"):
		this = SYM_ARRAY
	case pkg__strings.EqualFold(s, "Sym·Object"):
		this = SYM_OBJECT
	case pkg__strings.EqualFold(s, "Sym·Key"):
		this = SYM_KEY
	case pkg__strings.EqualFold(s, "Sym·Null"):
		this = SYM_NULL
	case pkg__strings.EqualFold(s, "Sym·Enummember"):
		this = SYM_ENUMMEMBER
	case pkg__strings.EqualFold(s, "Sym·Struct"):
		this = SYM_STRUCT
	case pkg__strings.EqualFold(s, "Sym·Event"):
		this = SYM_EVENT
	case pkg__strings.EqualFold(s, "Sym·Operator"):
		this = SYM_OPERATOR
	case pkg__strings.EqualFold(s, "Sym·Typeparameter"):
		this = SYM_TYPEPARAMETER
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (Symbol)(v)
	}
	return
}

// SymbolFromStringOr is like `SymbolFromString` but returns `fallback` for unrecognized inputs.
func SymbolFromStringOr(s string, fallback Symbol) (this Symbol) {
	maybeSymbol, err := SymbolFromString(s)
	if err == nil {
		this = maybeSymbol
	} else {
		this = fallback
	}
	return
}

// GoString implements the `fmt.GoStringer` interface.
func (this Symbol) GoString() (r string) {
	if (this < SYM_FILE) || (this > SYM_TYPEPARAMETER) {
		goto formatNum
	}
	switch this {
	case SYM_FILE:
		r = "SYM_FILE"
	case SYM_MODULE:
		r = "SYM_MODULE"
	case SYM_NAMESPACE:
		r = "SYM_NAMESPACE"
	case SYM_PACKAGE:
		r = "SYM_PACKAGE"
	case SYM_CLASS:
		r = "SYM_CLASS"
	case SYM_METHOD:
		r = "SYM_METHOD"
	case SYM_PROPERTY:
		r = "SYM_PROPERTY"
	case SYM_FIELD:
		r = "SYM_FIELD"
	case SYM_CONSTRUCTOR:
		r = "SYM_CONSTRUCTOR"
	case SYM_ENUM:
		r = "SYM_ENUM"
	case SYM_INTERFACE:
		r = "SYM_INTERFACE"
	case SYM_FUNCTION:
		r = "SYM_FUNCTION"
	case SYM_VARIABLE:
		r = "SYM_VARIABLE"
	case SYM_CONSTANT:
		r = "SYM_CONSTANT"
	case SYM_STRING:
		r = "SYM_STRING"
	case SYM_NUMBER:
		r = "SYM_NUMBER"
	case SYM_BOOLEAN:
		r = "SYM_BOOLEAN"
	case SYM_ARRAY:
		r = "SYM_ARRAY"
	case SYM_OBJECT:
		r = "SYM_OBJECT"
	case SYM_KEY:
		r = "SYM_KEY"
	case SYM_NULL:
		r = "SYM_NULL"
	case SYM_ENUMMEMBER:
		r = "SYM_ENUMMEMBER"
	case SYM_STRUCT:
		r = "SYM_STRUCT"
	case SYM_EVENT:
		r = "SYM_EVENT"
	case SYM_OPERATOR:
		r = "SYM_OPERATOR"
	case SYM_TYPEPARAMETER:
		r = "SYM_TYPEPARAMETER"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// SymbolFromGoString returns the `Symbol` represented by `s` (as returned by `Symbol.GoString`, and case-sensitively), or an `error` if none exists.
func SymbolFromGoString(s string) (this Symbol, err error) {
	if (len(s) < 7) || (len(s) > 17) || (s[0:4] != "SYM_") {
		goto tryParseNum
	}
	switch s {
	case "SYM_FILE":
		this = SYM_FILE
	case "SYM_MODULE":
		this = SYM_MODULE
	case "SYM_NAMESPACE":
		this = SYM_NAMESPACE
	case "SYM_PACKAGE":
		this = SYM_PACKAGE
	case "SYM_CLASS":
		this = SYM_CLASS
	case "SYM_METHOD":
		this = SYM_METHOD
	case "SYM_PROPERTY":
		this = SYM_PROPERTY
	case "SYM_FIELD":
		this = SYM_FIELD
	case "SYM_CONSTRUCTOR":
		this = SYM_CONSTRUCTOR
	case "SYM_ENUM":
		this = SYM_ENUM
	case "SYM_INTERFACE":
		this = SYM_INTERFACE
	case "SYM_FUNCTION":
		this = SYM_FUNCTION
	case "SYM_VARIABLE":
		this = SYM_VARIABLE
	case "SYM_CONSTANT":
		this = SYM_CONSTANT
	case "SYM_STRING":
		this = SYM_STRING
	case "SYM_NUMBER":
		this = SYM_NUMBER
	case "SYM_BOOLEAN":
		this = SYM_BOOLEAN
	case "SYM_ARRAY":
		this = SYM_ARRAY
	case "SYM_OBJECT":
		this = SYM_OBJECT
	case "SYM_KEY":
		this = SYM_KEY
	case "SYM_NULL":
		this = SYM_NULL
	case "SYM_ENUMMEMBER":
		this = SYM_ENUMMEMBER
	case "SYM_STRUCT":
		this = SYM_STRUCT
	case "SYM_EVENT":
		this = SYM_EVENT
	case "SYM_OPERATOR":
		this = SYM_OPERATOR
	case "SYM_TYPEPARAMETER":
		this = SYM_TYPEPARAMETER
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (Symbol)(v)
	}
	return
}

// SymbolFromGoStringOr is like `SymbolFromGoString` but returns `fallback` for unrecognized inputs.
func SymbolFromGoStringOr(s string, fallback Symbol) (this Symbol) {
	maybeSymbol, err := SymbolFromGoString(s)
	if err == nil {
		this = maybeSymbol
	} else {
		this = fallback
	}
	return
}

// IsCMPL_TEXT returns whether the value of this `Completion` equals `CMPL_TEXT`.
func (this Completion) IsCMPL_TEXT() (r bool) { r = this == CMPL_TEXT; return }

// IsCMPL_METHOD returns whether the value of this `Completion` equals `CMPL_METHOD`.
func (this Completion) IsCMPL_METHOD() (r bool) { r = this == CMPL_METHOD; return }

// IsCMPL_FUNCTION returns whether the value of this `Completion` equals `CMPL_FUNCTION`.
func (this Completion) IsCMPL_FUNCTION() (r bool) { r = this == CMPL_FUNCTION; return }

// IsCMPL_CONSTRUCTOR returns whether the value of this `Completion` equals `CMPL_CONSTRUCTOR`.
func (this Completion) IsCMPL_CONSTRUCTOR() (r bool) { r = this == CMPL_CONSTRUCTOR; return }

// IsCMPL_FIELD returns whether the value of this `Completion` equals `CMPL_FIELD`.
func (this Completion) IsCMPL_FIELD() (r bool) { r = this == CMPL_FIELD; return }

// IsCMPL_VARIABLE returns whether the value of this `Completion` equals `CMPL_VARIABLE`.
func (this Completion) IsCMPL_VARIABLE() (r bool) { r = this == CMPL_VARIABLE; return }

// IsCMPL_CLASS returns whether the value of this `Completion` equals `CMPL_CLASS`.
func (this Completion) IsCMPL_CLASS() (r bool) { r = this == CMPL_CLASS; return }

// IsCMPL_INTERFACE returns whether the value of this `Completion` equals `CMPL_INTERFACE`.
func (this Completion) IsCMPL_INTERFACE() (r bool) { r = this == CMPL_INTERFACE; return }

// IsCMPL_MODULE returns whether the value of this `Completion` equals `CMPL_MODULE`.
func (this Completion) IsCMPL_MODULE() (r bool) { r = this == CMPL_MODULE; return }

// IsCMPL_PROPERTY returns whether the value of this `Completion` equals `CMPL_PROPERTY`.
func (this Completion) IsCMPL_PROPERTY() (r bool) { r = this == CMPL_PROPERTY; return }

// IsCMPL_UNIT returns whether the value of this `Completion` equals `CMPL_UNIT`.
func (this Completion) IsCMPL_UNIT() (r bool) { r = this == CMPL_UNIT; return }

// IsCMPL_VALUE returns whether the value of this `Completion` equals `CMPL_VALUE`.
func (this Completion) IsCMPL_VALUE() (r bool) { r = this == CMPL_VALUE; return }

// IsCMPL_ENUM returns whether the value of this `Completion` equals `CMPL_ENUM`.
func (this Completion) IsCMPL_ENUM() (r bool) { r = this == CMPL_ENUM; return }

// IsCMPL_KEYWORD returns whether the value of this `Completion` equals `CMPL_KEYWORD`.
func (this Completion) IsCMPL_KEYWORD() (r bool) { r = this == CMPL_KEYWORD; return }

// IsCMPL_SNIPPET returns whether the value of this `Completion` equals `CMPL_SNIPPET`.
func (this Completion) IsCMPL_SNIPPET() (r bool) { r = this == CMPL_SNIPPET; return }

// IsCMPL_COLOR returns whether the value of this `Completion` equals `CMPL_COLOR`.
func (this Completion) IsCMPL_COLOR() (r bool) { r = this == CMPL_COLOR; return }

// IsCMPL_FILE returns whether the value of this `Completion` equals `CMPL_FILE`.
func (this Completion) IsCMPL_FILE() (r bool) { r = this == CMPL_FILE; return }

// IsCMPL_REFERENCE returns whether the value of this `Completion` equals `CMPL_REFERENCE`.
func (this Completion) IsCMPL_REFERENCE() (r bool) { r = this == CMPL_REFERENCE; return }

// IsCMPL_FOLDER returns whether the value of this `Completion` equals `CMPL_FOLDER`.
func (this Completion) IsCMPL_FOLDER() (r bool) { r = this == CMPL_FOLDER; return }

// IsCMPL_ENUMMEMBER returns whether the value of this `Completion` equals `CMPL_ENUMMEMBER`.
func (this Completion) IsCMPL_ENUMMEMBER() (r bool) { r = this == CMPL_ENUMMEMBER; return }

// IsCMPL_CONSTANT returns whether the value of this `Completion` equals `CMPL_CONSTANT`.
func (this Completion) IsCMPL_CONSTANT() (r bool) { r = this == CMPL_CONSTANT; return }

// IsCMPL_STRUCT returns whether the value of this `Completion` equals `CMPL_STRUCT`.
func (this Completion) IsCMPL_STRUCT() (r bool) { r = this == CMPL_STRUCT; return }

// IsCMPL_EVENT returns whether the value of this `Completion` equals `CMPL_EVENT`.
func (this Completion) IsCMPL_EVENT() (r bool) { r = this == CMPL_EVENT; return }

// IsCMPL_OPERATOR returns whether the value of this `Completion` equals `CMPL_OPERATOR`.
func (this Completion) IsCMPL_OPERATOR() (r bool) { r = this == CMPL_OPERATOR; return }

// IsCMPL_TYPEPARAMETER returns whether the value of this `Completion` equals `CMPL_TYPEPARAMETER`.
func (this Completion) IsCMPL_TYPEPARAMETER() (r bool) { r = this == CMPL_TYPEPARAMETER; return }

// Valid returns whether the value of this `Completion` is between `CMPL_TEXT` (inclusive) and `CMPL_TYPEPARAMETER` (inclusive).
func (this Completion) Valid() (r bool) {
	r = (this >= CMPL_TEXT) && (this <= CMPL_TYPEPARAMETER)
	return
}

// WellknownCompletions returns the `names` and `values` of all 25 well-known `Completion` enumerants.
func WellknownCompletions() (names []string, values []Completion) {
	names, values = []string{"CMPL_TEXT", "CMPL_METHOD", "CMPL_FUNCTION", "CMPL_CONSTRUCTOR", "CMPL_FIELD", "CMPL_VARIABLE", "CMPL_CLASS", "CMPL_INTERFACE", "CMPL_MODULE", "CMPL_PROPERTY", "CMPL_UNIT", "CMPL_VALUE", "CMPL_ENUM", "CMPL_KEYWORD", "CMPL_SNIPPET", "CMPL_COLOR", "CMPL_FILE", "CMPL_REFERENCE", "CMPL_FOLDER", "CMPL_ENUMMEMBER", "CMPL_CONSTANT", "CMPL_STRUCT", "CMPL_EVENT", "CMPL_OPERATOR", "CMPL_TYPEPARAMETER"}, []Completion{CMPL_TEXT, CMPL_METHOD, CMPL_FUNCTION, CMPL_CONSTRUCTOR, CMPL_FIELD, CMPL_VARIABLE, CMPL_CLASS, CMPL_INTERFACE, CMPL_MODULE, CMPL_PROPERTY, CMPL_UNIT, CMPL_VALUE, CMPL_ENUM, CMPL_KEYWORD, CMPL_SNIPPET, CMPL_COLOR, CMPL_FILE, CMPL_REFERENCE, CMPL_FOLDER, CMPL_ENUMMEMBER, CMPL_CONSTANT, CMPL_STRUCT, CMPL_EVENT, CMPL_OPERATOR, CMPL_TYPEPARAMETER}
	return
}

// String implements the `fmt.Stringer` interface.
func (this Completion) String() (r string) {
	if (this < CMPL_TEXT) || (this > CMPL_TYPEPARAMETER) {
		goto formatNum
	}
	switch this {
	case CMPL_TEXT:
		r = "Cmpl·Text"
	case CMPL_METHOD:
		r = "Cmpl·Method"
	case CMPL_FUNCTION:
		r = "Cmpl·Function"
	case CMPL_CONSTRUCTOR:
		r = "Cmpl·Constructor"
	case CMPL_FIELD:
		r = "Cmpl·Field"
	case CMPL_VARIABLE:
		r = "Cmpl·Variable"
	case CMPL_CLASS:
		r = "Cmpl·Class"
	case CMPL_INTERFACE:
		r = "Cmpl·Interface"
	case CMPL_MODULE:
		r = "Cmpl·Module"
	case CMPL_PROPERTY:
		r = "Cmpl·Property"
	case CMPL_UNIT:
		r = "Cmpl·Unit"
	case CMPL_VALUE:
		r = "Cmpl·Value"
	case CMPL_ENUM:
		r = "Cmpl·Enum"
	case CMPL_KEYWORD:
		r = "Cmpl·Keyword"
	case CMPL_SNIPPET:
		r = "Cmpl·Snippet"
	case CMPL_COLOR:
		r = "Cmpl·Color"
	case CMPL_FILE:
		r = "Cmpl·File"
	case CMPL_REFERENCE:
		r = "Cmpl·Reference"
	case CMPL_FOLDER:
		r = "Cmpl·Folder"
	case CMPL_ENUMMEMBER:
		r = "Cmpl·Enummember"
	case CMPL_CONSTANT:
		r = "Cmpl·Constant"
	case CMPL_STRUCT:
		r = "Cmpl·Struct"
	case CMPL_EVENT:
		r = "Cmpl·Event"
	case CMPL_OPERATOR:
		r = "Cmpl·Operator"
	case CMPL_TYPEPARAMETER:
		r = "Cmpl·Typeparameter"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// CompletionFromString returns the `Completion` represented by `s` (as returned by `Completion.String`, but case-insensitively), or an `error` if none exists.
func CompletionFromString(s string) (this Completion, err error) {
	if (len(s) < 10) || (len(s) > 19) || (!pkg__strings.EqualFold(s[0:6], "Cmpl·")) {
		goto tryParseNum
	}
	switch {
	case pkg__strings.EqualFold(s, "Cmpl·Text"):
		this = CMPL_TEXT
	case pkg__strings.EqualFold(s, "Cmpl·Method"):
		this = CMPL_METHOD
	case pkg__strings.EqualFold(s, "Cmpl·Function"):
		this = CMPL_FUNCTION
	case pkg__strings.EqualFold(s, "Cmpl·Constructor"):
		this = CMPL_CONSTRUCTOR
	case pkg__strings.EqualFold(s, "Cmpl·Field"):
		this = CMPL_FIELD
	case pkg__strings.EqualFold(s, "Cmpl·Variable"):
		this = CMPL_VARIABLE
	case pkg__strings.EqualFold(s, "Cmpl·Class"):
		this = CMPL_CLASS
	case pkg__strings.EqualFold(s, "Cmpl·Interface"):
		this = CMPL_INTERFACE
	case pkg__strings.EqualFold(s, "Cmpl·Module"):
		this = CMPL_MODULE
	case pkg__strings.EqualFold(s, "Cmpl·Property"):
		this = CMPL_PROPERTY
	case pkg__strings.EqualFold(s, "Cmpl·Unit"):
		this = CMPL_UNIT
	case pkg__strings.EqualFold(s, "Cmpl·Value"):
		this = CMPL_VALUE
	case pkg__strings.EqualFold(s, "Cmpl·Enum"):
		this = CMPL_ENUM
	case pkg__strings.EqualFold(s, "Cmpl·Keyword"):
		this = CMPL_KEYWORD
	case pkg__strings.EqualFold(s, "Cmpl·Snippet"):
		this = CMPL_SNIPPET
	case pkg__strings.EqualFold(s, "Cmpl·Color"):
		this = CMPL_COLOR
	case pkg__strings.EqualFold(s, "Cmpl·File"):
		this = CMPL_FILE
	case pkg__strings.EqualFold(s, "Cmpl·Reference"):
		this = CMPL_REFERENCE
	case pkg__strings.EqualFold(s, "Cmpl·Folder"):
		this = CMPL_FOLDER
	case pkg__strings.EqualFold(s, "Cmpl·Enummember"):
		this = CMPL_ENUMMEMBER
	case pkg__strings.EqualFold(s, "Cmpl·Constant"):
		this = CMPL_CONSTANT
	case pkg__strings.EqualFold(s, "Cmpl·Struct"):
		this = CMPL_STRUCT
	case pkg__strings.EqualFold(s, "Cmpl·Event"):
		this = CMPL_EVENT
	case pkg__strings.EqualFold(s, "Cmpl·Operator"):
		this = CMPL_OPERATOR
	case pkg__strings.EqualFold(s, "Cmpl·Typeparameter"):
		this = CMPL_TYPEPARAMETER
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (Completion)(v)
	}
	return
}

// CompletionFromStringOr is like `CompletionFromString` but returns `fallback` for unrecognized inputs.
func CompletionFromStringOr(s string, fallback Completion) (this Completion) {
	maybeCompletion, err := CompletionFromString(s)
	if err == nil {
		this = maybeCompletion
	} else {
		this = fallback
	}
	return
}

// GoString implements the `fmt.GoStringer` interface.
func (this Completion) GoString() (r string) {
	if (this < CMPL_TEXT) || (this > CMPL_TYPEPARAMETER) {
		goto formatNum
	}
	switch this {
	case CMPL_TEXT:
		r = "CMPL_TEXT"
	case CMPL_METHOD:
		r = "CMPL_METHOD"
	case CMPL_FUNCTION:
		r = "CMPL_FUNCTION"
	case CMPL_CONSTRUCTOR:
		r = "CMPL_CONSTRUCTOR"
	case CMPL_FIELD:
		r = "CMPL_FIELD"
	case CMPL_VARIABLE:
		r = "CMPL_VARIABLE"
	case CMPL_CLASS:
		r = "CMPL_CLASS"
	case CMPL_INTERFACE:
		r = "CMPL_INTERFACE"
	case CMPL_MODULE:
		r = "CMPL_MODULE"
	case CMPL_PROPERTY:
		r = "CMPL_PROPERTY"
	case CMPL_UNIT:
		r = "CMPL_UNIT"
	case CMPL_VALUE:
		r = "CMPL_VALUE"
	case CMPL_ENUM:
		r = "CMPL_ENUM"
	case CMPL_KEYWORD:
		r = "CMPL_KEYWORD"
	case CMPL_SNIPPET:
		r = "CMPL_SNIPPET"
	case CMPL_COLOR:
		r = "CMPL_COLOR"
	case CMPL_FILE:
		r = "CMPL_FILE"
	case CMPL_REFERENCE:
		r = "CMPL_REFERENCE"
	case CMPL_FOLDER:
		r = "CMPL_FOLDER"
	case CMPL_ENUMMEMBER:
		r = "CMPL_ENUMMEMBER"
	case CMPL_CONSTANT:
		r = "CMPL_CONSTANT"
	case CMPL_STRUCT:
		r = "CMPL_STRUCT"
	case CMPL_EVENT:
		r = "CMPL_EVENT"
	case CMPL_OPERATOR:
		r = "CMPL_OPERATOR"
	case CMPL_TYPEPARAMETER:
		r = "CMPL_TYPEPARAMETER"
	default:
		goto formatNum
	}
	return
formatNum:
	r = pkg__strconv.FormatUint((uint64)(this), 10)
	return
}

// CompletionFromGoString returns the `Completion` represented by `s` (as returned by `Completion.GoString`, and case-sensitively), or an `error` if none exists.
func CompletionFromGoString(s string) (this Completion, err error) {
	if (len(s) < 9) || (len(s) > 18) || (s[0:5] != "CMPL_") {
		goto tryParseNum
	}
	switch s {
	case "CMPL_TEXT":
		this = CMPL_TEXT
	case "CMPL_METHOD":
		this = CMPL_METHOD
	case "CMPL_FUNCTION":
		this = CMPL_FUNCTION
	case "CMPL_CONSTRUCTOR":
		this = CMPL_CONSTRUCTOR
	case "CMPL_FIELD":
		this = CMPL_FIELD
	case "CMPL_VARIABLE":
		this = CMPL_VARIABLE
	case "CMPL_CLASS":
		this = CMPL_CLASS
	case "CMPL_INTERFACE":
		this = CMPL_INTERFACE
	case "CMPL_MODULE":
		this = CMPL_MODULE
	case "CMPL_PROPERTY":
		this = CMPL_PROPERTY
	case "CMPL_UNIT":
		this = CMPL_UNIT
	case "CMPL_VALUE":
		this = CMPL_VALUE
	case "CMPL_ENUM":
		this = CMPL_ENUM
	case "CMPL_KEYWORD":
		this = CMPL_KEYWORD
	case "CMPL_SNIPPET":
		this = CMPL_SNIPPET
	case "CMPL_COLOR":
		this = CMPL_COLOR
	case "CMPL_FILE":
		this = CMPL_FILE
	case "CMPL_REFERENCE":
		this = CMPL_REFERENCE
	case "CMPL_FOLDER":
		this = CMPL_FOLDER
	case "CMPL_ENUMMEMBER":
		this = CMPL_ENUMMEMBER
	case "CMPL_CONSTANT":
		this = CMPL_CONSTANT
	case "CMPL_STRUCT":
		this = CMPL_STRUCT
	case "CMPL_EVENT":
		this = CMPL_EVENT
	case "CMPL_OPERATOR":
		this = CMPL_OPERATOR
	case "CMPL_TYPEPARAMETER":
		this = CMPL_TYPEPARAMETER
	default:
		goto tryParseNum
	}
	return
tryParseNum:
	var v uint64
	v, err = pkg__strconv.ParseUint(s, 10, 8)
	if err == nil {
		this = (Completion)(v)
	}
	return
}

// CompletionFromGoStringOr is like `CompletionFromGoString` but returns `fallback` for unrecognized inputs.
func CompletionFromGoStringOr(s string, fallback Completion) (this Completion) {
	maybeCompletion, err := CompletionFromGoString(s)
	if err == nil {
		this = maybeCompletion
	} else {
		this = fallback
	}
	return
}

func (this ListItems) Index(v IListItem) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this ListItems) IndexFunc(ok func(IListItem) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this ListItems) LastIndex(v IListItem) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this ListItems) LastIndexFunc(ok func(IListItem) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this ListItems) Indices(v IListItem) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this ListItems) IndicesFunc(ok func(IListItem) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this PkgInfos) Index(v *PkgInfo) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this PkgInfos) IndexFunc(ok func(*PkgInfo) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this PkgInfos) LastIndex(v *PkgInfo) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this PkgInfos) LastIndexFunc(ok func(*PkgInfo) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this PkgInfos) Indices(v *PkgInfo) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this PkgInfos) IndicesFunc(ok func(*PkgInfo) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this DiagItems) Index(v *DiagItem) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagItems) IndexFunc(ok func(*DiagItem) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagItems) LastIndex(v *DiagItem) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagItems) LastIndexFunc(ok func(*DiagItem) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagItems) Indices(v *DiagItem) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this DiagItems) IndicesFunc(ok func(*DiagItem) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this DiagBuildJobs) Index(v *DiagJobBuild) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagBuildJobs) IndexFunc(ok func(*DiagJobBuild) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagBuildJobs) LastIndex(v *DiagJobBuild) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagBuildJobs) LastIndexFunc(ok func(*DiagJobBuild) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagBuildJobs) Indices(v *DiagJobBuild) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this DiagBuildJobs) IndicesFunc(ok func(*DiagJobBuild) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this DiagLintJobs) Index(v *DiagJobLint) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagLintJobs) IndexFunc(ok func(*DiagJobLint) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagLintJobs) LastIndex(v *DiagJobLint) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagLintJobs) LastIndexFunc(ok func(*DiagJobLint) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this DiagLintJobs) Indices(v *DiagJobLint) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this DiagLintJobs) IndicesFunc(ok func(*DiagJobLint) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this SrcIntelCompls) Index(v *SrcIntelCompl) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcIntelCompls) IndexFunc(ok func(*SrcIntelCompl) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcIntelCompls) LastIndex(v *SrcIntelCompl) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcIntelCompls) LastIndexFunc(ok func(*SrcIntelCompl) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcIntelCompls) Indices(v *SrcIntelCompl) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this SrcIntelCompls) IndicesFunc(ok func(*SrcIntelCompl) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this SrcLocs) Index(v *SrcLoc) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLocs) IndexFunc(ok func(*SrcLoc) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLocs) LastIndex(v *SrcLoc) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLocs) LastIndexFunc(ok func(*SrcLoc) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLocs) Indices(v *SrcLoc) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this SrcLocs) IndicesFunc(ok func(*SrcLoc) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this SrcLenses) Index(v *SrcLens) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLenses) IndexFunc(ok func(*SrcLens) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLenses) LastIndex(v *SrcLens) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLenses) LastIndexFunc(ok func(*SrcLens) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcLenses) Indices(v *SrcLens) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this SrcLenses) IndicesFunc(ok func(*SrcLens) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this SrcModEdits) Index(v srcModEdit) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcModEdits) IndexFunc(ok func(srcModEdit) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcModEdits) LastIndex(v srcModEdit) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcModEdits) LastIndexFunc(ok func(srcModEdit) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this SrcModEdits) Indices(v srcModEdit) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this SrcModEdits) IndicesFunc(ok func(srcModEdit) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

// IsTOOLS_CAT_MOD_REN returns whether the value of this `ToolCats` equals `TOOLS_CAT_MOD_REN`.
func (this ToolCats) IsTOOLS_CAT_MOD_REN() (r bool) { r = this == TOOLS_CAT_MOD_REN; return }

// IsTOOLS_CAT_MOD_FMT returns whether the value of this `ToolCats` equals `TOOLS_CAT_MOD_FMT`.
func (this ToolCats) IsTOOLS_CAT_MOD_FMT() (r bool) { r = this == TOOLS_CAT_MOD_FMT; return }

// IsTOOLS_CAT_INTEL_TIPS returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_TIPS`.
func (this ToolCats) IsTOOLS_CAT_INTEL_TIPS() (r bool) { r = this == TOOLS_CAT_INTEL_TIPS; return }

// IsTOOLS_CAT_INTEL_SYMS returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_SYMS`.
func (this ToolCats) IsTOOLS_CAT_INTEL_SYMS() (r bool) { r = this == TOOLS_CAT_INTEL_SYMS; return }

// IsTOOLS_CAT_INTEL_HIGH returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_HIGH`.
func (this ToolCats) IsTOOLS_CAT_INTEL_HIGH() (r bool) { r = this == TOOLS_CAT_INTEL_HIGH; return }

// IsTOOLS_CAT_INTEL_CMPL returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_CMPL`.
func (this ToolCats) IsTOOLS_CAT_INTEL_CMPL() (r bool) { r = this == TOOLS_CAT_INTEL_CMPL; return }

// IsTOOLS_CAT_INTEL_NAV returns whether the value of this `ToolCats` equals `TOOLS_CAT_INTEL_NAV`.
func (this ToolCats) IsTOOLS_CAT_INTEL_NAV() (r bool) { r = this == TOOLS_CAT_INTEL_NAV; return }

// IsTOOLS_CAT_EXTRAS_QUERY returns whether the value of this `ToolCats` equals `TOOLS_CAT_EXTRAS_QUERY`.
func (this ToolCats) IsTOOLS_CAT_EXTRAS_QUERY() (r bool) { r = this == TOOLS_CAT_EXTRAS_QUERY; return }

// IsTOOLS_CAT_DIAGS returns whether the value of this `ToolCats` equals `TOOLS_CAT_DIAGS`.
func (this ToolCats) IsTOOLS_CAT_DIAGS() (r bool) { r = this == TOOLS_CAT_DIAGS; return }

// IsTOOLS_CAT_RUNONSAVE returns whether the value of this `ToolCats` equals `TOOLS_CAT_RUNONSAVE`.
func (this ToolCats) IsTOOLS_CAT_RUNONSAVE() (r bool) { r = this == TOOLS_CAT_RUNONSAVE; return }

// Valid returns whether the value of this `ToolCats` is between `TOOLS_CAT_MOD_REN` (inclusive) and `TOOLS_CAT_RUNONSAVE` (inclusive).
func (this ToolCats) Valid() (r bool) {
	r = (this >= TOOLS_CAT_MOD_REN) && (this <= TOOLS_CAT_RUNONSAVE)
	return
}

// WellknownToolCatses returns the `names` and `values` of all 10 well-known `ToolCats` enumerants.
func WellknownToolCatses() (names []string, values []ToolCats) {
	names, values = []string{"TOOLS_CAT_MOD_REN", "TOOLS_CAT_MOD_FMT", "TOOLS_CAT_INTEL_TIPS", "TOOLS_CAT_INTEL_SYMS", "TOOLS_CAT_INTEL_HIGH", "TOOLS_CAT_INTEL_CMPL", "TOOLS_CAT_INTEL_NAV", "TOOLS_CAT_EXTRAS_QUERY", "TOOLS_CAT_DIAGS", "TOOLS_CAT_RUNONSAVE"}, []ToolCats{TOOLS_CAT_MOD_REN, TOOLS_CAT_MOD_FMT, TOOLS_CAT_INTEL_TIPS, TOOLS_CAT_INTEL_SYMS, TOOLS_CAT_INTEL_HIGH, TOOLS_CAT_INTEL_CMPL, TOOLS_CAT_INTEL_NAV, TOOLS_CAT_EXTRAS_QUERY, TOOLS_CAT_DIAGS, TOOLS_CAT_RUNONSAVE}
	return
}

func (this Tools) Index(v *Tool) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Tools) IndexFunc(ok func(*Tool) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Tools) LastIndex(v *Tool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Tools) LastIndexFunc(ok func(*Tool) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this Tools) Indices(v *Tool) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this Tools) IndicesFunc(ok func(*Tool) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this MenuItems) Index(v *MenuItem) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this MenuItems) IndexFunc(ok func(*MenuItem) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this MenuItems) LastIndex(v *MenuItem) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this MenuItems) LastIndexFunc(ok func(*MenuItem) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this MenuItems) Indices(v *MenuItem) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this MenuItems) IndicesFunc(ok func(*MenuItem) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}

func (this sideViewTreeItem) Index(v string) (r int) {
	for i := range this {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this sideViewTreeItem) IndexFunc(ok func(string) bool) (r int) {
	for i := range this {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this sideViewTreeItem) LastIndex(v string) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if this[i] == v {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this sideViewTreeItem) LastIndexFunc(ok func(string) bool) (r int) {
	for i := len(this) - 1; i > -1; i-- {
		if ok(this[i]) {
			r = i
			return
		}
	}
	r = -1
	return
}

func (this sideViewTreeItem) Indices(v string) (r []int) {
	for i := range this {
		if this[i] == v {
			r = append(r, i)
		}
	}
	return
}

func (this sideViewTreeItem) IndicesFunc(ok func(string) bool) (r []int) {
	for i := range this {
		if ok(this[i]) {
			r = append(r, i)
		}
	}
	return
}
