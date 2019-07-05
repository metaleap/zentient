package z

// TO MOT EDIT: code generated with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__github_com_go_leap_str "github.com/go-leap/str"
	pkg__strconv "strconv"
)

// String implements the Go standard library's `fmt.Stringer` interface.
func (me CaddyStatus) String() (r string) {
	switch me {
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
	r = pkg__strconv.FormatUint((uint64)(me), 10)
	return
}

// Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN` (inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).
func (me IpcIDs) Valid() (r bool) {
	r = ((me >= IPCID_MENUS_MAIN) && (me <= IPCID_EXTRAS_QUERY_RUN))
	return
}

// String implements the Go standard library's `fmt.Stringer` interface.
func (me IpcIDs) String() (r string) {
	switch me {
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
	case IPCID_SRCINTEL_INFOBITS:
		r = "IPCID_SRCINTEL_INFOBITS"
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
	r = pkg__strconv.FormatUint((uint64)(me), 10)
	return
}

// String implements the Go standard library's `fmt.Stringer` interface.
func (me DiagSeverity) String() (r string) {
	switch me {
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
	r = pkg__strconv.Itoa((int)(me))
	return
}

// String implements the Go standard library's `fmt.Stringer` interface.
func (me Symbol) String() (r string) {
	switch me {
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
	r = pkg__strconv.FormatUint((uint64)(me), 10)
	return
}

// String implements the Go standard library's `fmt.Stringer` interface.
func (me Completion) String() (r string) {
	switch me {
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
	r = pkg__strconv.FormatUint((uint64)(me), 10)
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *fooResp) MarshalJSON() (r []byte, err error) { return }

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *fooResp) UnmarshalJSON(v []byte) (err error) { return }

// StructFieldsTraverse calls `on` 15x: once for each field in this `fooResp` with its name, its pointer, `true` if name (or embed name) begins in upper-case (else `false`), and `true` if field is an embed (else `false`).
func (me *fooResp) StructFieldsTraverse(on func(name string, ptr interface{}, isNameUpperCase bool, isEmbed bool)) {
	on("IpcID", &me.IpcID, true, false)
	on("ReqID", &me.ReqID, true, false)
	on("ErrMsg", &me.ErrMsg, true, false)
	on("SrcIntel", &me.SrcIntel, true, false)
	on("SrcDiags", &me.SrcDiags, true, false)
	on("ipcReq", &me.ipcReq, false, true)
	on("SrcMods", &me.SrcMods, true, false)
	on("muhPrivate", &me.muhPrivate, false, false)
	on("SrcActions", &me.SrcActions, true, false)
	on("Extras", &me.Extras, true, false)
	on("SrcLens", &me.SrcLens, true, true)
	on("Pats", &me.Pats, true, true)
	on("Menu", &me.Menu, true, false)
	on("CaddyUpdate", &me.CaddyUpdate, true, false)
	on("Val", &me.Val, true, false)
}

func (me *fooResp) StructFieldsGet(name string, v interface{}) (r interface{}, ok bool) {
	switch name {
	case "IpcID":
		r = me.IpcID
		ok = true
	case "ReqID":
		r = me.ReqID
		ok = true
	case "ErrMsg":
		r = me.ErrMsg
		ok = true
	case "SrcIntel":
		r = me.SrcIntel
		ok = true
	case "SrcDiags":
		r = me.SrcDiags
		ok = true
	case "ipcReq":
		r = me.ipcReq
		ok = true
	case "SrcMods":
		r = me.SrcMods
		ok = true
	case "muhPrivate":
		r = me.muhPrivate
		ok = true
	case "SrcActions":
		r = me.SrcActions
		ok = true
	case "Extras":
		r = me.Extras
		ok = true
	case "SrcLens":
		r = me.SrcLens
		ok = true
	case "Pats":
		r = me.Pats
		ok = true
	case "Menu":
		r = me.Menu
		ok = true
	case "CaddyUpdate":
		r = me.CaddyUpdate
		ok = true
	case "Val":
		r = me.Val
		ok = true
	default:
		r = v
	}
	return
}

func (me *fooResp) StructFieldsSet(name string, v interface{}) (okName bool, okType bool) {
	switch name {
	case "IpcID":
		okName = true
		t, ok := v.(IpcIDs)
		if ok {
			okType = true
			me.IpcID = t
		}
	case "ReqID":
		okName = true
		t, ok := v.(int64)
		if ok {
			okType = true
			me.ReqID = t
		}
	case "ErrMsg":
		okName = true
		t, ok := v.(string)
		if ok {
			okType = true
			me.ErrMsg = t
		}
	case "SrcIntel":
		okName = true
		t, ok := v.(*ipcRespSrcIntel)
		if ok {
			okType = true
			me.SrcIntel = t
		}
	case "SrcDiags":
		okName = true
		t, ok := v.(*ipcRespDiag)
		if ok {
			okType = true
			me.SrcDiags = t
		}
	case "ipcReq":
		okName = true
		t, ok := v.(*ipcReq)
		if ok {
			okType = true
			me.ipcReq = t
		}
	case "SrcMods":
		okName = true
		t, ok := v.(SrcLenses)
		if ok {
			okType = true
			me.SrcMods = t
		}
	case "muhPrivate":
		okName = true
		t, ok := v.(int)
		if ok {
			okType = true
			me.muhPrivate = t
		}
	case "SrcActions":
		okName = true
		t, ok := v.([]EditorAction)
		if ok {
			okType = true
			me.SrcActions = t
		}
	case "Extras":
		okName = true
		t, ok := v.(*IpcRespExtras)
		if ok {
			okType = true
			me.Extras = t
		}
	case "SrcLens":
		okName = true
		t, ok := v.(*SrcLens)
		if ok {
			okType = true
			me.SrcLens = t
		}
	case "Pats":
		okName = true
		t, ok := v.(pkg__github_com_go_leap_str.Pats)
		if ok {
			okType = true
			me.Pats = t
		}
	case "Menu":
		okName = true
		t, ok := v.(*ipcRespMenu)
		if ok {
			okType = true
			me.Menu = t
		}
	case "CaddyUpdate":
		okName = true
		t, ok := v.(*Caddy)
		if ok {
			okType = true
			me.CaddyUpdate = t
		}
	case "Val":
		okName = true
		t, ok := v.(interface{})
		if ok {
			okType = true
			me.Val = t
		}
	}
	return
}
