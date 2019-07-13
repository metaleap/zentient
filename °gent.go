package z

// DON'T EDIT: code gen'd with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__encoding_json "encoding/json"
	pkg__github_com_go_leap_str "github.com/go-leap/str"
	pkg__strconv "strconv"
)

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
	case IPCID_SRCINTEL_ANNS:
		r = "IPCID_SRCINTEL_ANNS"
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

// StructFieldsTraverse calls `on` 20x: once for each field in this `fooResp` with its name, its pointer, `true` if name (or embed name) begins in upper-case (else `false`), and `true` if field is an embed (else `false`).
func (me *fooResp) StructFieldsTraverse(on func(name string, ptr interface{}, isNameUpperCase bool, isEmbed bool)) {
	on("IpcID", &me.IpcID, true, false)
	on("ReqID", &me.ReqID, true, false)
	on("Flagg", &me.Flagg, true, false)
	on("ErrMsg", &me.ErrMsg, true, false)
	on("SrcIntel", &me.SrcIntel, true, false)
	on("SrcDiags", &me.SrcDiags, true, false)
	on("IpcReq", &me.IpcReq, true, true)
	on("SrcMods", &me.SrcMods, true, false)
	on("muhPrivate", &me.muhPrivate, false, false)
	on("SrcActions", &me.SrcActions, true, false)
	on("Extras", &me.Extras, true, false)
	on("SrcLens", &me.SrcLens, true, true)
	on("Fn", &me.Fn, true, false)
	on("Link", &me.Link, true, false)
	on("Ch", &me.Ch, true, false)
	on("Pats", &me.Pats, true, true)
	on("Menu", &me.Menu, true, false)
	on("Nope", &me.Nope, true, false)
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
	case "Flagg":
		r = me.Flagg
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
	case "IpcReq":
		r = me.IpcReq
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
	case "Fn":
		r = me.Fn
		ok = true
	case "Link":
		r = me.Link
		ok = true
	case "Ch":
		r = me.Ch
		ok = true
	case "Pats":
		r = me.Pats
		ok = true
	case "Menu":
		r = me.Menu
		ok = true
	case "Nope":
		r = me.Nope
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
	case "Flagg":
		okName = true
		t, ok := v.(bool)
		if ok {
			okType = true
			me.Flagg = t
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
		t, ok := v.(*SrcIntel)
		if ok {
			okType = true
			me.SrcIntel = t
		}
	case "SrcDiags":
		okName = true
		t, ok := v.(*Diags)
		if ok {
			okType = true
			me.SrcDiags = t
		}
	case "IpcReq":
		okName = true
		t, ok := v.(*IpcReq)
		if ok {
			okType = true
			me.IpcReq = t
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
		t, ok := v.(*Extras)
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
	case "Fn":
		okName = true
		t, ok := v.(func())
		if ok {
			okType = true
			me.Fn = t
		}
	case "Link":
		okName = true
		t, ok := v.(*fooResp)
		if ok {
			okType = true
			me.Link = t
		}
	case "Ch":
		okName = true
		t, ok := v.(chan bool)
		if ok {
			okType = true
			me.Ch = t
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
		t, ok := v.(*MenuResponse)
		if ok {
			okType = true
			me.Menu = t
		}
	case "Nope":
		okName = true
		t, ok := v.(string)
		if ok {
			okType = true
			me.Nope = t
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

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *fooResp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 128)
	r = append(r, 123)
	{
		r = append(r, "\"ii\":"...)
		r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
	}
	if me.ReqID != 0 {
		r = append(r, ",\"ri\":"...)
		r = append(r, pkg__strconv.FormatInt((int64)(me.ReqID), 10)...)
	}
	{
		r = append(r, ",\"Flagg\":"...)
		r = append(r, pkg__strconv.FormatBool(me.Flagg)...)
	}
	if len(me.ErrMsg) != 0 {
		r = append(r, ",\"err\":"...)
		r = append(r, pkg__strconv.Quote(me.ErrMsg)...)
	}
	if nil != me.SrcIntel {
		r = append(r, ",\"sI\":"...)
		r = append(r, 123)
		if len(me.SrcIntel.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			for i1 := range me.SrcIntel.InfoTips {
				if i1 != 0 {
					r = append(r, 44)
				}
				r = append(r, 123)
				{
					r = append(r, "\"value\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcIntel.InfoTips[i1].Value)...)
				}
				if len(me.SrcIntel.InfoTips[i1].Language) != 0 {
					r = append(r, ",\"language\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcIntel.InfoTips[i1].Language)...)
				}
				r = append(r, 125)
			}
			r = append(r, 93)
		}
		if len(me.SrcIntel.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			r = append(r, 91)
			for i2 := range me.SrcIntel.Refs {
				if nil != me.SrcIntel.Refs[i2] {
					if i2 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"e\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Flag), 10)...)
					}
					if len(me.SrcIntel.Refs[i2].FilePath) != 0 {
						r = append(r, ",\"f\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Refs[i2].FilePath)...)
					}
					if nil != me.SrcIntel.Refs[i2].Pos {
						r = append(r, ",\"p\":"...)
						r = append(r, 123)
						if me.SrcIntel.Refs[i2].Pos.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Pos.Ln), 10)...)
						}
						if me.SrcIntel.Refs[i2].Pos.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Pos.Col), 10)...)
						}
						if me.SrcIntel.Refs[i2].Pos.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Pos.Off), 10)...)
						}
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"p\":"...)
						r = append(r, "null"...)
					}
					if nil != me.SrcIntel.Refs[i2].Range {
						r = append(r, ",\"r\":"...)
						r = append(r, 123)
						r = append(r, "\"s\":"...)
						r = append(r, 123)
						if me.SrcIntel.Refs[i2].Range.Start.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.Start.Ln), 10)...)
						}
						if me.SrcIntel.Refs[i2].Range.Start.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.Start.Col), 10)...)
						}
						if me.SrcIntel.Refs[i2].Range.Start.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.Start.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, ",\"e\":"...)
						r = append(r, 123)
						if me.SrcIntel.Refs[i2].Range.End.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.End.Ln), 10)...)
						}
						if me.SrcIntel.Refs[i2].Range.End.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.End.Col), 10)...)
						}
						if me.SrcIntel.Refs[i2].Range.End.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Refs[i2].Range.End.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"r\":"...)
						r = append(r, "null"...)
					}
					r = append(r, 125)
				} else if true {
					if i2 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		if nil != me.SrcIntel.Sig {
			r = append(r, ",\"Sig\":"...)
			r = append(r, 123)
			{
				r = append(r, "\"activeSignature\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Sig.ActiveSignature), 10)...)
			}
			if me.SrcIntel.Sig.ActiveParameter != 0 {
				r = append(r, ",\"activeParameter\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Sig.ActiveParameter), 10)...)
			}
			if len(me.SrcIntel.Sig.Signatures) != 0 {
				r = append(r, ",\"signatures\":"...)
				r = append(r, 91)
				for i3 := range me.SrcIntel.Sig.Signatures {
					if i3 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"label\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Sig.Signatures[i3].Label)...)
					}
					r = append(r, ",\"documentation\":"...)
					r = append(r, 123)
					if len(me.SrcIntel.Sig.Signatures[i3].Documentation.Value) != 0 {
						r = append(r, "\"value\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Sig.Signatures[i3].Documentation.Value)...)
					}
					if me.SrcIntel.Sig.Signatures[i3].Documentation.IsTrusted {
						r = append(r, ",\"isTrusted\":"...)
						r = append(r, pkg__strconv.FormatBool(me.SrcIntel.Sig.Signatures[i3].Documentation.IsTrusted)...)
					}
					r = append(r, 125)
					{
						r = append(r, ",\"parameters\":"...)
						r = append(r, 91)
						for i4 := range me.SrcIntel.Sig.Signatures[i3].Parameters {
							if i4 != 0 {
								r = append(r, 44)
							}
							r = append(r, 123)
							{
								r = append(r, "\"label\":"...)
								r = append(r, pkg__strconv.Quote(me.SrcIntel.Sig.Signatures[i3].Parameters[i4].Label)...)
							}
							r = append(r, ",\"documentation\":"...)
							r = append(r, 123)
							if len(me.SrcIntel.Sig.Signatures[i3].Parameters[i4].Documentation.Value) != 0 {
								r = append(r, "\"value\":"...)
								r = append(r, pkg__strconv.Quote(me.SrcIntel.Sig.Signatures[i3].Parameters[i4].Documentation.Value)...)
							}
							if me.SrcIntel.Sig.Signatures[i3].Parameters[i4].Documentation.IsTrusted {
								r = append(r, ",\"isTrusted\":"...)
								r = append(r, pkg__strconv.FormatBool(me.SrcIntel.Sig.Signatures[i3].Parameters[i4].Documentation.IsTrusted)...)
							}
							r = append(r, 125)
							r = append(r, 125)
						}
						r = append(r, 93)
					}
					r = append(r, 125)
				}
				r = append(r, 93)
			}
			r = append(r, 125)
		} else if false {
			r = append(r, ",\"Sig\":"...)
			r = append(r, "null"...)
		}
		if len(me.SrcIntel.Cmpl) != 0 {
			r = append(r, ",\"Cmpl\":"...)
			r = append(r, 91)
			for i5 := range me.SrcIntel.Cmpl {
				if nil != me.SrcIntel.Cmpl[i5] {
					if i5 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					if me.SrcIntel.Cmpl[i5].Kind != 0 {
						r = append(r, "\"kind\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Cmpl[i5].Kind), 10)...)
					}
					{
						r = append(r, ",\"label\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i5].Label)...)
					}
					if nil != me.SrcIntel.Cmpl[i5].Documentation {
						r = append(r, ",\"documentation\":"...)
						r = append(r, 123)
						if len(me.SrcIntel.Cmpl[i5].Documentation.Value) != 0 {
							r = append(r, "\"value\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i5].Documentation.Value)...)
						}
						if me.SrcIntel.Cmpl[i5].Documentation.IsTrusted {
							r = append(r, ",\"isTrusted\":"...)
							r = append(r, pkg__strconv.FormatBool(me.SrcIntel.Cmpl[i5].Documentation.IsTrusted)...)
						}
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"documentation\":"...)
						r = append(r, "null"...)
					}
					if len(me.SrcIntel.Cmpl[i5].Detail) != 0 {
						r = append(r, ",\"detail\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i5].Detail)...)
					}
					if len(me.SrcIntel.Cmpl[i5].SortText) != 0 {
						r = append(r, ",\"sortText\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i5].SortText)...)
					}
					r = append(r, 125)
				} else if true {
					if i5 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		if len(me.SrcIntel.Syms) != 0 {
			r = append(r, ",\"Syms\":"...)
			r = append(r, 91)
			for i6 := range me.SrcIntel.Syms {
				if nil != me.SrcIntel.Syms[i6] {
					if i6 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, ",\"e\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Flag), 10)...)
					}
					if len(me.SrcIntel.Syms[i6].FilePath) != 0 {
						r = append(r, ",\"f\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Syms[i6].FilePath)...)
					}
					if nil != me.SrcIntel.Syms[i6].Pos {
						r = append(r, ",\"p\":"...)
						r = append(r, 123)
						if me.SrcIntel.Syms[i6].Pos.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Pos.Ln), 10)...)
						}
						if me.SrcIntel.Syms[i6].Pos.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Pos.Col), 10)...)
						}
						if me.SrcIntel.Syms[i6].Pos.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Pos.Off), 10)...)
						}
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"p\":"...)
						r = append(r, "null"...)
					}
					if nil != me.SrcIntel.Syms[i6].Range {
						r = append(r, ",\"r\":"...)
						r = append(r, 123)
						r = append(r, "\"s\":"...)
						r = append(r, 123)
						if me.SrcIntel.Syms[i6].Range.Start.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.Start.Ln), 10)...)
						}
						if me.SrcIntel.Syms[i6].Range.Start.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.Start.Col), 10)...)
						}
						if me.SrcIntel.Syms[i6].Range.Start.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.Start.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, ",\"e\":"...)
						r = append(r, 123)
						if me.SrcIntel.Syms[i6].Range.End.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.End.Ln), 10)...)
						}
						if me.SrcIntel.Syms[i6].Range.End.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.End.Col), 10)...)
						}
						if me.SrcIntel.Syms[i6].Range.End.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Syms[i6].Range.End.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"r\":"...)
						r = append(r, "null"...)
					}
					if len(me.SrcIntel.Syms[i6].Txt) != 0 {
						r = append(r, ",\"t\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Syms[i6].Txt)...)
					}
					if len(me.SrcIntel.Syms[i6].Str) != 0 {
						r = append(r, ",\"s\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Syms[i6].Str)...)
					}
					if me.SrcIntel.Syms[i6].CrLf {
						r = append(r, ",\"l\":"...)
						r = append(r, pkg__strconv.FormatBool(me.SrcIntel.Syms[i6].CrLf)...)
					}
					r = append(r, 125)
				} else if true {
					if i6 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		if len(me.SrcIntel.Anns) != 0 {
			r = append(r, ",\"Anns\":"...)
			r = append(r, 91)
			for i7 := range me.SrcIntel.Anns {
				if nil != me.SrcIntel.Anns[i7] {
					if i7 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					r = append(r, "\"Range\":"...)
					r = append(r, 123)
					r = append(r, "\"s\":"...)
					r = append(r, 123)
					if me.SrcIntel.Anns[i7].Range.Start.Ln != 0 {
						r = append(r, "\"l\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.Start.Ln), 10)...)
					}
					if me.SrcIntel.Anns[i7].Range.Start.Col != 0 {
						r = append(r, ",\"c\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.Start.Col), 10)...)
					}
					if me.SrcIntel.Anns[i7].Range.Start.Off != 0 {
						r = append(r, ",\"o\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.Start.Off), 10)...)
					}
					r = append(r, 125)
					r = append(r, ",\"e\":"...)
					r = append(r, 123)
					if me.SrcIntel.Anns[i7].Range.End.Ln != 0 {
						r = append(r, "\"l\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.End.Ln), 10)...)
					}
					if me.SrcIntel.Anns[i7].Range.End.Col != 0 {
						r = append(r, ",\"c\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.End.Col), 10)...)
					}
					if me.SrcIntel.Anns[i7].Range.End.Off != 0 {
						r = append(r, ",\"o\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Anns[i7].Range.End.Off), 10)...)
					}
					r = append(r, 125)
					r = append(r, 125)
					{
						r = append(r, ",\"Title\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i7].Title)...)
					}
					if len(me.SrcIntel.Anns[i7].Desc) != 0 {
						r = append(r, ",\"Desc\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i7].Desc)...)
					}
					{
						r = append(r, ",\"CmdName\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i7].CmdName)...)
					}
					r = append(r, 125)
				} else if true {
					if i7 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"sI\":"...)
		r = append(r, "null"...)
	}
	if nil != me.SrcDiags {
		r = append(r, ",\"srcDiags\":"...)
		r = append(r, 123)
		r = append(r, "\"All\":"...)
		r = append(r, "{}"...)
		{
			r = append(r, ",\"FixUps\":"...)
			r = append(r, 91)
			for i8 := range me.SrcDiags.FixUps {
				if nil != me.SrcDiags.FixUps[i8] {
					if i8 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"FilePath\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcDiags.FixUps[i8].FilePath)...)
					}
					r = append(r, ",\"Desc\":"...)
					r = append(r, "{}"...)
					{
						r = append(r, ",\"Edits\":"...)
						r = append(r, 91)
						for i9 := range me.SrcDiags.FixUps[i8].Edits {
							if i9 != 0 {
								r = append(r, 44)
							}
							r = append(r, 123)
							if nil != me.SrcDiags.FixUps[i8].Edits[i9].At {
								r = append(r, "\"At\":"...)
								r = append(r, 123)
								r = append(r, "\"s\":"...)
								r = append(r, 123)
								if me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Ln != 0 {
									r = append(r, "\"l\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Ln), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Col != 0 {
									r = append(r, ",\"c\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Col), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Off != 0 {
									r = append(r, ",\"o\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.Start.Off), 10)...)
								}
								r = append(r, 125)
								r = append(r, ",\"e\":"...)
								r = append(r, 123)
								if me.SrcDiags.FixUps[i8].Edits[i9].At.End.Ln != 0 {
									r = append(r, "\"l\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.End.Ln), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Edits[i9].At.End.Col != 0 {
									r = append(r, ",\"c\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.End.Col), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Edits[i9].At.End.Off != 0 {
									r = append(r, ",\"o\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Edits[i9].At.End.Off), 10)...)
								}
								r = append(r, 125)
								r = append(r, 125)
							} else if true {
								r = append(r, "\"At\":"...)
								r = append(r, "null"...)
							}
							{
								r = append(r, ",\"Val\":"...)
								r = append(r, pkg__strconv.Quote(me.SrcDiags.FixUps[i8].Edits[i9].Val)...)
							}
							r = append(r, 125)
						}
						r = append(r, 93)
					}
					{
						r = append(r, ",\"Dropped\":"...)
						r = append(r, 91)
						for i10 := range me.SrcDiags.FixUps[i8].Dropped {
							if i10 != 0 {
								r = append(r, 44)
							}
							r = append(r, 123)
							if nil != me.SrcDiags.FixUps[i8].Dropped[i10].At {
								r = append(r, "\"At\":"...)
								r = append(r, 123)
								r = append(r, "\"s\":"...)
								r = append(r, 123)
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Ln != 0 {
									r = append(r, "\"l\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Ln), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Col != 0 {
									r = append(r, ",\"c\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Col), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Off != 0 {
									r = append(r, ",\"o\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.Start.Off), 10)...)
								}
								r = append(r, 125)
								r = append(r, ",\"e\":"...)
								r = append(r, 123)
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Ln != 0 {
									r = append(r, "\"l\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Ln), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Col != 0 {
									r = append(r, ",\"c\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Col), 10)...)
								}
								if me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Off != 0 {
									r = append(r, ",\"o\":"...)
									r = append(r, pkg__strconv.FormatInt((int64)(me.SrcDiags.FixUps[i8].Dropped[i10].At.End.Off), 10)...)
								}
								r = append(r, 125)
								r = append(r, 125)
							} else if true {
								r = append(r, "\"At\":"...)
								r = append(r, "null"...)
							}
							{
								r = append(r, ",\"Val\":"...)
								r = append(r, pkg__strconv.Quote(me.SrcDiags.FixUps[i8].Dropped[i10].Val)...)
							}
							r = append(r, 125)
						}
						r = append(r, 93)
					}
					r = append(r, 125)
				} else if true {
					if i8 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"LangID\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcDiags.LangID)...)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"srcDiags\":"...)
		r = append(r, "null"...)
	}
	{
		r = append(r, ",\"ri\":"...)
		r = append(r, pkg__strconv.FormatInt((int64)(me.ReqID), 10)...)
	}
	{
		r = append(r, ",\"ii\":"...)
		r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
	}
	if me.IpcArgs != nil {
		r = append(r, ",\"ia\":"...)
		var e error
		var sl []byte
		j, ok := me.IpcArgs.(pkg__encoding_json.Marshaler)
		if ok && (j != nil) {
			sl, e = j.MarshalJSON()
		} else {
			sl, e = pkg__encoding_json.Marshal(me.IpcArgs)
		}
		if e == nil {
			r = append(r, sl...)
		} else {
			err = e
			return
		}
	} else {
		r = append(r, ",\"ia\":"...)
		r = append(r, "null"...)
	}
	if nil != me.ProjUpd {
		r = append(r, ",\"projUpd\":"...)
		r = append(r, 123)
		{
			r = append(r, "\"AddedDirs\":"...)
			r = append(r, 91)
			for i11 := range me.ProjUpd.AddedDirs {
				{
					if i11 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ProjUpd.AddedDirs[i11])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"RemovedDirs\":"...)
			r = append(r, 91)
			for i12 := range me.ProjUpd.RemovedDirs {
				{
					if i12 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ProjUpd.RemovedDirs[i12])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"OpenedFiles\":"...)
			r = append(r, 91)
			for i13 := range me.ProjUpd.OpenedFiles {
				{
					if i13 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ProjUpd.OpenedFiles[i13])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"ClosedFiles\":"...)
			r = append(r, 91)
			for i14 := range me.ProjUpd.ClosedFiles {
				{
					if i14 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ProjUpd.ClosedFiles[i14])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"WrittenFiles\":"...)
			r = append(r, 91)
			for i15 := range me.ProjUpd.WrittenFiles {
				{
					if i15 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ProjUpd.WrittenFiles[i15])...)
				}
			}
			r = append(r, 93)
		}
		r = append(r, ",\"LiveFiles\":"...)
		r = append(r, "{}"...)
		r = append(r, 125)
	} else if true {
		r = append(r, ",\"projUpd\":"...)
		r = append(r, "null"...)
	}
	if nil != me.SrcLens {
		r = append(r, ",\"srcLens\":"...)
		r = append(r, 123)
		{
			r = append(r, ",\"e\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Flag), 10)...)
		}
		if len(me.SrcLens.FilePath) != 0 {
			r = append(r, ",\"f\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcLens.FilePath)...)
		}
		if nil != me.SrcLens.Pos {
			r = append(r, ",\"p\":"...)
			r = append(r, 123)
			if me.SrcLens.Pos.Ln != 0 {
				r = append(r, "\"l\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Pos.Ln), 10)...)
			}
			if me.SrcLens.Pos.Col != 0 {
				r = append(r, ",\"c\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Pos.Col), 10)...)
			}
			if me.SrcLens.Pos.Off != 0 {
				r = append(r, ",\"o\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Pos.Off), 10)...)
			}
			r = append(r, 125)
		} else if false {
			r = append(r, ",\"p\":"...)
			r = append(r, "null"...)
		}
		if nil != me.SrcLens.Range {
			r = append(r, ",\"r\":"...)
			r = append(r, 123)
			r = append(r, "\"s\":"...)
			r = append(r, 123)
			if me.SrcLens.Range.Start.Ln != 0 {
				r = append(r, "\"l\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.Start.Ln), 10)...)
			}
			if me.SrcLens.Range.Start.Col != 0 {
				r = append(r, ",\"c\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.Start.Col), 10)...)
			}
			if me.SrcLens.Range.Start.Off != 0 {
				r = append(r, ",\"o\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.Start.Off), 10)...)
			}
			r = append(r, 125)
			r = append(r, ",\"e\":"...)
			r = append(r, 123)
			if me.SrcLens.Range.End.Ln != 0 {
				r = append(r, "\"l\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.End.Ln), 10)...)
			}
			if me.SrcLens.Range.End.Col != 0 {
				r = append(r, ",\"c\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.End.Col), 10)...)
			}
			if me.SrcLens.Range.End.Off != 0 {
				r = append(r, ",\"o\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.Range.End.Off), 10)...)
			}
			r = append(r, 125)
			r = append(r, 125)
		} else if false {
			r = append(r, ",\"r\":"...)
			r = append(r, "null"...)
		}
		if len(me.SrcLens.Txt) != 0 {
			r = append(r, ",\"t\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcLens.Txt)...)
		}
		if len(me.SrcLens.Str) != 0 {
			r = append(r, ",\"s\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcLens.Str)...)
		}
		if me.SrcLens.CrLf {
			r = append(r, ",\"l\":"...)
			r = append(r, pkg__strconv.FormatBool(me.SrcLens.CrLf)...)
		}
		r = append(r, 125)
	} else if true {
		r = append(r, ",\"srcLens\":"...)
		r = append(r, "null"...)
	}
	if len(me.SrcMods) != 0 {
		r = append(r, ",\"srcMods\":"...)
		r = append(r, 91)
		for i16 := range me.SrcMods {
			if nil != me.SrcMods[i16] {
				if i16 != 0 {
					r = append(r, 44)
				}
				r = append(r, 123)
				{
					r = append(r, ",\"e\":"...)
					r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Flag), 10)...)
				}
				if len(me.SrcMods[i16].FilePath) != 0 {
					r = append(r, ",\"f\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcMods[i16].FilePath)...)
				}
				if nil != me.SrcMods[i16].Pos {
					r = append(r, ",\"p\":"...)
					r = append(r, 123)
					if me.SrcMods[i16].Pos.Ln != 0 {
						r = append(r, "\"l\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Pos.Ln), 10)...)
					}
					if me.SrcMods[i16].Pos.Col != 0 {
						r = append(r, ",\"c\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Pos.Col), 10)...)
					}
					if me.SrcMods[i16].Pos.Off != 0 {
						r = append(r, ",\"o\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Pos.Off), 10)...)
					}
					r = append(r, 125)
				} else if false {
					r = append(r, ",\"p\":"...)
					r = append(r, "null"...)
				}
				if nil != me.SrcMods[i16].Range {
					r = append(r, ",\"r\":"...)
					r = append(r, 123)
					r = append(r, "\"s\":"...)
					r = append(r, 123)
					if me.SrcMods[i16].Range.Start.Ln != 0 {
						r = append(r, "\"l\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.Start.Ln), 10)...)
					}
					if me.SrcMods[i16].Range.Start.Col != 0 {
						r = append(r, ",\"c\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.Start.Col), 10)...)
					}
					if me.SrcMods[i16].Range.Start.Off != 0 {
						r = append(r, ",\"o\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.Start.Off), 10)...)
					}
					r = append(r, 125)
					r = append(r, ",\"e\":"...)
					r = append(r, 123)
					if me.SrcMods[i16].Range.End.Ln != 0 {
						r = append(r, "\"l\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.End.Ln), 10)...)
					}
					if me.SrcMods[i16].Range.End.Col != 0 {
						r = append(r, ",\"c\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.End.Col), 10)...)
					}
					if me.SrcMods[i16].Range.End.Off != 0 {
						r = append(r, ",\"o\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.SrcMods[i16].Range.End.Off), 10)...)
					}
					r = append(r, 125)
					r = append(r, 125)
				} else if false {
					r = append(r, ",\"r\":"...)
					r = append(r, "null"...)
				}
				if len(me.SrcMods[i16].Txt) != 0 {
					r = append(r, ",\"t\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcMods[i16].Txt)...)
				}
				if len(me.SrcMods[i16].Str) != 0 {
					r = append(r, ",\"s\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcMods[i16].Str)...)
				}
				if me.SrcMods[i16].CrLf {
					r = append(r, ",\"l\":"...)
					r = append(r, pkg__strconv.FormatBool(me.SrcMods[i16].CrLf)...)
				}
				r = append(r, 125)
			} else if true {
				if i16 != 0 {
					r = append(r, 44)
				}
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
	}
	if len(me.SrcActions) != 0 {
		r = append(r, ",\"srcActions\":"...)
		r = append(r, 91)
		for i17 := range me.SrcActions {
			if i17 != 0 {
				r = append(r, 44)
			}
			r = append(r, 123)
			{
				r = append(r, "\"title\":"...)
				r = append(r, pkg__strconv.Quote(me.SrcActions[i17].Title)...)
			}
			{
				r = append(r, ",\"command\":"...)
				r = append(r, pkg__strconv.Quote(me.SrcActions[i17].Cmd)...)
			}
			if len(me.SrcActions[i17].Hint) != 0 {
				r = append(r, ",\"tooltip\":"...)
				r = append(r, pkg__strconv.Quote(me.SrcActions[i17].Hint)...)
			}
			if len(me.SrcActions[i17].Arguments) != 0 {
				r = append(r, ",\"arguments\":"...)
				r = append(r, 91)
				for i18 := range me.SrcActions[i17].Arguments {
					if me.SrcActions[i17].Arguments[i18] != nil {
						if i18 != 0 {
							r = append(r, 44)
						}
						var e error
						var sl []byte
						j, ok := me.SrcActions[i17].Arguments[i18].(pkg__encoding_json.Marshaler)
						if ok && (j != nil) {
							sl, e = j.MarshalJSON()
						} else {
							sl, e = pkg__encoding_json.Marshal(me.SrcActions[i17].Arguments[i18])
						}
						if e == nil {
							r = append(r, sl...)
						} else {
							err = e
							return
						}
					} else {
						if i18 != 0 {
							r = append(r, 44)
						}
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
			}
			r = append(r, 125)
		}
		r = append(r, 93)
	}
	if nil != me.Extras {
		r = append(r, ",\"extras\":"...)
		r = append(r, 123)
		if len(me.Extras.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			for i19 := range me.Extras.InfoTips {
				if i19 != 0 {
					r = append(r, 44)
				}
				r = append(r, 123)
				{
					r = append(r, "\"value\":"...)
					r = append(r, pkg__strconv.Quote(me.Extras.InfoTips[i19].Value)...)
				}
				if len(me.Extras.InfoTips[i19].Language) != 0 {
					r = append(r, ",\"language\":"...)
					r = append(r, pkg__strconv.Quote(me.Extras.InfoTips[i19].Language)...)
				}
				r = append(r, 125)
			}
			r = append(r, 93)
		}
		if len(me.Extras.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			r = append(r, 91)
			for i20 := range me.Extras.Refs {
				if nil != me.Extras.Refs[i20] {
					if i20 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"e\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Flag), 10)...)
					}
					if len(me.Extras.Refs[i20].FilePath) != 0 {
						r = append(r, ",\"f\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Refs[i20].FilePath)...)
					}
					if nil != me.Extras.Refs[i20].Pos {
						r = append(r, ",\"p\":"...)
						r = append(r, 123)
						if me.Extras.Refs[i20].Pos.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Pos.Ln), 10)...)
						}
						if me.Extras.Refs[i20].Pos.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Pos.Col), 10)...)
						}
						if me.Extras.Refs[i20].Pos.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Pos.Off), 10)...)
						}
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"p\":"...)
						r = append(r, "null"...)
					}
					if nil != me.Extras.Refs[i20].Range {
						r = append(r, ",\"r\":"...)
						r = append(r, 123)
						r = append(r, "\"s\":"...)
						r = append(r, 123)
						if me.Extras.Refs[i20].Range.Start.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.Start.Ln), 10)...)
						}
						if me.Extras.Refs[i20].Range.Start.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.Start.Col), 10)...)
						}
						if me.Extras.Refs[i20].Range.Start.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.Start.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, ",\"e\":"...)
						r = append(r, 123)
						if me.Extras.Refs[i20].Range.End.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.End.Ln), 10)...)
						}
						if me.Extras.Refs[i20].Range.End.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.End.Col), 10)...)
						}
						if me.Extras.Refs[i20].Range.End.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Extras.Refs[i20].Range.End.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"r\":"...)
						r = append(r, "null"...)
					}
					r = append(r, 125)
				} else if true {
					if i20 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"Items\":"...)
			r = append(r, 91)
			for i21 := range me.Extras.Items {
				if nil != me.Extras.Items[i21] {
					if i21 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"id\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].ID)...)
					}
					{
						r = append(r, ",\"label\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].Label)...)
					}
					{
						r = append(r, ",\"description\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].Desc)...)
					}
					if len(me.Extras.Items[i21].Detail) != 0 {
						r = append(r, ",\"detail\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].Detail)...)
					}
					if len(me.Extras.Items[i21].QueryArg) != 0 {
						r = append(r, ",\"arg\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].QueryArg)...)
					}
					if len(me.Extras.Items[i21].FilePos) != 0 {
						r = append(r, ",\"fPos\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.Items[i21].FilePos)...)
					}
					r = append(r, 125)
				} else if true {
					if i21 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		if len(me.Extras.Warns) != 0 {
			r = append(r, ",\"Warns\":"...)
			r = append(r, 91)
			for i22 := range me.Extras.Warns {
				{
					if i22 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.Extras.Warns[i22])...)
				}
			}
			r = append(r, 93)
		}
		if len(me.Extras.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Extras.Desc)...)
		}
		if len(me.Extras.Url) != 0 {
			r = append(r, ",\"Url\":"...)
			r = append(r, pkg__strconv.Quote(me.Extras.Url)...)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"extras\":"...)
		r = append(r, "null"...)
	}
	{
		r = append(r, ",\"e\":"...)
		r = append(r, pkg__strconv.FormatInt((int64)(me.Flag), 10)...)
	}
	if len(me.FilePath) != 0 {
		r = append(r, ",\"f\":"...)
		r = append(r, pkg__strconv.Quote(me.FilePath)...)
	}
	if nil != me.Pos {
		r = append(r, ",\"p\":"...)
		r = append(r, 123)
		if me.Pos.Ln != 0 {
			r = append(r, "\"l\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Pos.Ln), 10)...)
		}
		if me.Pos.Col != 0 {
			r = append(r, ",\"c\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Pos.Col), 10)...)
		}
		if me.Pos.Off != 0 {
			r = append(r, ",\"o\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Pos.Off), 10)...)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"p\":"...)
		r = append(r, "null"...)
	}
	if nil != me.Range {
		r = append(r, ",\"r\":"...)
		r = append(r, 123)
		r = append(r, "\"s\":"...)
		r = append(r, 123)
		if me.Range.Start.Ln != 0 {
			r = append(r, "\"l\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.Start.Ln), 10)...)
		}
		if me.Range.Start.Col != 0 {
			r = append(r, ",\"c\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.Start.Col), 10)...)
		}
		if me.Range.Start.Off != 0 {
			r = append(r, ",\"o\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.Start.Off), 10)...)
		}
		r = append(r, 125)
		r = append(r, ",\"e\":"...)
		r = append(r, 123)
		if me.Range.End.Ln != 0 {
			r = append(r, "\"l\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.End.Ln), 10)...)
		}
		if me.Range.End.Col != 0 {
			r = append(r, ",\"c\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.End.Col), 10)...)
		}
		if me.Range.End.Off != 0 {
			r = append(r, ",\"o\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Range.End.Off), 10)...)
		}
		r = append(r, 125)
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"r\":"...)
		r = append(r, "null"...)
	}
	if len(me.Txt) != 0 {
		r = append(r, ",\"t\":"...)
		r = append(r, pkg__strconv.Quote(me.Txt)...)
	}
	if len(me.Str) != 0 {
		r = append(r, ",\"s\":"...)
		r = append(r, pkg__strconv.Quote(me.Str)...)
	}
	if me.CrLf {
		r = append(r, ",\"l\":"...)
		r = append(r, pkg__strconv.FormatBool(me.CrLf)...)
	}
	if nil != me.Link {
		if me.Link != nil {
			r = append(r, ",\"Link\":"...)
			var e error
			var sl []byte
			sl, e = me.Link.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		} else {
			r = append(r, ",\"Link\":"...)
			r = append(r, "null"...)
		}
	} else if true {
		r = append(r, ",\"Link\":"...)
		r = append(r, "null"...)
	}
	if me.Pats != nil {
		r = append(r, ",\"Pats\":"...)
		var e error
		var sl []byte
		sl, e = pkg__encoding_json.Marshal(me.Pats)
		if e == nil {
			r = append(r, sl...)
		} else {
			err = e
			return
		}
	} else {
		r = append(r, ",\"Pats\":"...)
		r = append(r, "null"...)
	}
	if nil != me.Menu {
		r = append(r, ",\"menu\":"...)
		r = append(r, 123)
		if nil != me.Menu.SubMenu {
			r = append(r, "\"SubMenu\":"...)
			r = append(r, 123)
			if len(me.Menu.SubMenu.Desc) != 0 {
				r = append(r, "\"desc\":"...)
				r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Desc)...)
			}
			if me.Menu.SubMenu.TopLevel {
				r = append(r, ",\"topLevel\":"...)
				r = append(r, pkg__strconv.FormatBool(me.Menu.SubMenu.TopLevel)...)
			}
			{
				r = append(r, ",\"items\":"...)
				r = append(r, 91)
				for i23 := range me.Menu.SubMenu.Items {
					if nil != me.Menu.SubMenu.Items[i23] {
						if i23 != 0 {
							r = append(r, 44)
						}
						r = append(r, 123)
						if me.Menu.SubMenu.Items[i23].IpcID != 0 {
							r = append(r, "\"ii\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.SubMenu.Items[i23].IpcID), 10)...)
						}
						if me.Menu.SubMenu.Items[i23].IpcArgs != nil {
							r = append(r, ",\"ia\":"...)
							var e error
							var sl []byte
							j, ok := me.Menu.SubMenu.Items[i23].IpcArgs.(pkg__encoding_json.Marshaler)
							if ok && (j != nil) {
								sl, e = j.MarshalJSON()
							} else {
								sl, e = pkg__encoding_json.Marshal(me.Menu.SubMenu.Items[i23].IpcArgs)
							}
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
						if len(me.Menu.SubMenu.Items[i23].Category) != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i23].Category)...)
						}
						{
							r = append(r, ",\"t\":"...)
							r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i23].Title)...)
						}
						if len(me.Menu.SubMenu.Items[i23].Desc) != 0 {
							r = append(r, ",\"d\":"...)
							r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i23].Desc)...)
						}
						if len(me.Menu.SubMenu.Items[i23].Hint) != 0 {
							r = append(r, ",\"h\":"...)
							r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i23].Hint)...)
						}
						if len(me.Menu.SubMenu.Items[i23].Confirm) != 0 {
							r = append(r, ",\"q\":"...)
							r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i23].Confirm)...)
						}
						r = append(r, 125)
					} else if true {
						if i23 != 0 {
							r = append(r, 44)
						}
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
			}
			r = append(r, 125)
		} else if false {
			r = append(r, "\"SubMenu\":"...)
			r = append(r, "null"...)
		}
		if len(me.Menu.WebsiteURL) != 0 {
			r = append(r, ",\"WebsiteURL\":"...)
			r = append(r, pkg__strconv.Quote(me.Menu.WebsiteURL)...)
		}
		if len(me.Menu.NoteInfo) != 0 {
			r = append(r, ",\"NoteInfo\":"...)
			r = append(r, pkg__strconv.Quote(me.Menu.NoteInfo)...)
		}
		if len(me.Menu.NoteWarn) != 0 {
			r = append(r, ",\"NoteWarn\":"...)
			r = append(r, pkg__strconv.Quote(me.Menu.NoteWarn)...)
		}
		if len(me.Menu.UxActionLabel) != 0 {
			r = append(r, ",\"UxActionLabel\":"...)
			r = append(r, pkg__strconv.Quote(me.Menu.UxActionLabel)...)
		}
		if len(me.Menu.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			r = append(r, 91)
			for i24 := range me.Menu.Refs {
				if nil != me.Menu.Refs[i24] {
					if i24 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					{
						r = append(r, "\"e\":"...)
						r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Flag), 10)...)
					}
					if len(me.Menu.Refs[i24].FilePath) != 0 {
						r = append(r, ",\"f\":"...)
						r = append(r, pkg__strconv.Quote(me.Menu.Refs[i24].FilePath)...)
					}
					if nil != me.Menu.Refs[i24].Pos {
						r = append(r, ",\"p\":"...)
						r = append(r, 123)
						if me.Menu.Refs[i24].Pos.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Pos.Ln), 10)...)
						}
						if me.Menu.Refs[i24].Pos.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Pos.Col), 10)...)
						}
						if me.Menu.Refs[i24].Pos.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Pos.Off), 10)...)
						}
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"p\":"...)
						r = append(r, "null"...)
					}
					if nil != me.Menu.Refs[i24].Range {
						r = append(r, ",\"r\":"...)
						r = append(r, 123)
						r = append(r, "\"s\":"...)
						r = append(r, 123)
						if me.Menu.Refs[i24].Range.Start.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.Start.Ln), 10)...)
						}
						if me.Menu.Refs[i24].Range.Start.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.Start.Col), 10)...)
						}
						if me.Menu.Refs[i24].Range.Start.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.Start.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, ",\"e\":"...)
						r = append(r, 123)
						if me.Menu.Refs[i24].Range.End.Ln != 0 {
							r = append(r, "\"l\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.End.Ln), 10)...)
						}
						if me.Menu.Refs[i24].Range.End.Col != 0 {
							r = append(r, ",\"c\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.End.Col), 10)...)
						}
						if me.Menu.Refs[i24].Range.End.Off != 0 {
							r = append(r, ",\"o\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.Refs[i24].Range.End.Off), 10)...)
						}
						r = append(r, 125)
						r = append(r, 125)
					} else if false {
						r = append(r, ",\"r\":"...)
						r = append(r, "null"...)
					}
					r = append(r, 125)
				} else if true {
					if i24 != 0 {
						r = append(r, 44)
					}
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"menu\":"...)
		r = append(r, "null"...)
	}
	if nil != me.CaddyUpdate {
		r = append(r, ",\"caddy\":"...)
		r = append(r, 123)
		if len(me.CaddyUpdate.ID) != 0 {
			r = append(r, "\"ID\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.ID)...)
		}
		if len(me.CaddyUpdate.LangID) != 0 {
			r = append(r, ",\"LangID\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.LangID)...)
		}
		{
			r = append(r, ",\"Icon\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Icon)...)
		}
		if len(me.CaddyUpdate.Title) != 0 {
			r = append(r, ",\"Title\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Title)...)
		}
		r = append(r, ",\"Status\":"...)
		r = append(r, 123)
		{
			r = append(r, "\"Flag\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.CaddyUpdate.Status.Flag), 10)...)
		}
		if len(me.CaddyUpdate.Status.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Status.Desc)...)
		}
		r = append(r, 125)
		if len(me.CaddyUpdate.Details) != 0 {
			r = append(r, ",\"Details\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Details)...)
		}
		if len(me.CaddyUpdate.UxActionID) != 0 {
			r = append(r, ",\"UxActionID\":"...)
			r = append(r, pkg__strconv.Quote(me.CaddyUpdate.UxActionID)...)
		}
		if me.CaddyUpdate.ShowTitle {
			r = append(r, ",\"ShowTitle\":"...)
			r = append(r, pkg__strconv.FormatBool(me.CaddyUpdate.ShowTitle)...)
		}
		r = append(r, 125)
	} else if false {
		r = append(r, ",\"caddy\":"...)
		r = append(r, "null"...)
	}
	if me.Val != nil {
		r = append(r, ",\"valya\":"...)
		var e error
		var sl []byte
		j, ok := me.Val.(pkg__encoding_json.Marshaler)
		if ok && (j != nil) {
			sl, e = j.MarshalJSON()
		} else {
			sl, e = pkg__encoding_json.Marshal(me.Val)
		}
		if e == nil {
			r = append(r, sl...)
		} else {
			err = e
			return
		}
	} else {
		r = append(r, ",\"valya\":"...)
		r = append(r, "null"...)
	}
	r = append(r, 125)
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *fooResp) preview_UnmarshalJSON(v []byte) (err error) { return }
