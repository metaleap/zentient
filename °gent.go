package z

// DON'T EDIT: code gen'd with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__encoding_json "encoding/json"
	pkg__fmt "fmt"
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

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *IpcReq) preview_MarshalJSON() (r []byte, err error) { panic("IpcReq"); return }

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcReq) preview_UnmarshalJSON(v []byte) (err error) { ; return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *IpcResp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si31 := len(r)
		if me.IpcID != 0 {
			r = append(r, ",\"ii\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
		}
		if me.ReqID != 0 {
			r = append(r, ",\"ri\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.ReqID), 10)...)
		}
		if len(me.ErrMsg) != 0 {
			r = append(r, ",\"err\":"...)
			r = append(r, pkg__strconv.Quote(me.ErrMsg)...)
		}
		if nil != me.SrcIntel {
			{
				r = append(r, ",\"sI\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcIntel.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.SrcDiags {
			{
				r = append(r, ",\"srcDiags\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcDiags.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if len(me.SrcMods) != 0 {
			r = append(r, ",\"srcMods\":"...)
			var e error
			var sl []byte
			sl, e = me.SrcMods.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if len(me.SrcActions) != 0 {
			r = append(r, ",\"srcActions\":"...)
			r = append(r, 91)
			ai51 := len(r)
			for i50 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i50].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai51] == 44 {
				r = append(r[:ai51], r[(ai51+1):]...)
			}
		}
		if nil != me.Extras {
			{
				r = append(r, ",\"extras\":"...)
				var e error
				var sl []byte
				sl, e = me.Extras.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.Menu {
			{
				r = append(r, ",\"menu\":"...)
				var e error
				var sl []byte
				sl, e = me.Menu.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.CaddyUpdate {
			{
				r = append(r, ",\"caddy\":"...)
				var e error
				var sl []byte
				sl, e = me.CaddyUpdate.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if me.Val != nil {
			r = append(r, ",\"val\":"...)
			var e error
			var sl []byte
			j, _ := me.Val.(pkg__encoding_json.Marshaler)
			if j != nil {
				sl, e = j.MarshalJSON()
			} else {
				v80, ok81 := me.Val.(map[string]interface{})
				if ok81 {
					sl, e = __gent__jsonMarshal_mapsstring_interface____(v80)
				} else {
					v78, ok79 := me.Val.([]string)
					if ok79 {
						sl, e = __gent__jsonMarshal_s_string(v78)
					} else {
						v76, ok77 := me.Val.(string)
						if ok77 {
							sl, e = __gent__jsonMarshal_string(v76)
						} else {
							println("JSON.MARSHAL:", pkg__fmt.Sprintf("%T", me.Val))
							sl, e = pkg__encoding_json.Marshal(me.Val)
						}
					}
				}
			}
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si31] == 44 {
			r = append(r[:si31], r[(si31+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Diags) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si82 := len(r)
		{
			r = append(r, ",\"All\":"...)
			var e error
			var sl []byte
			sl, e = me.All.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if nil == me.FixUps {
			r = append(r, ",\"FixUps\":"...)
			r = append(r, "null"...)
		} else {
			r = append(r, ",\"FixUps\":"...)
			r = append(r, 91)
			ai90 := len(r)
			for i89 := range me.FixUps {
				if nil != me.FixUps[i89] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.FixUps[i89].preview_MarshalJSON()
						if e == nil {
							r = append(r, sl...)
						} else {
							err = e
							return
						}
					}
				} else {
					r = append(r, 44)
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
			if r[ai90] == 44 {
				r = append(r[:ai90], r[(ai90+1):]...)
			}
		}
		{
			r = append(r, ",\"LangID\":"...)
			r = append(r, pkg__strconv.Quote(me.LangID)...)
		}
		r = append(r, 125)
		if r[si82] == 44 {
			r = append(r[:si82], r[(si82+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Extras) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si97 := len(r)
		if len(me.SrcIntels.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai99 := len(r)
			for i98 := range me.SrcIntels.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcIntels.InfoTips[i98].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai99] == 44 {
				r = append(r[:ai99], r[(ai99+1):]...)
			}
		}
		if len(me.SrcIntels.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			var e error
			var sl []byte
			sl, e = me.SrcIntels.Refs.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if nil == me.Items {
			r = append(r, ",\"Items\":"...)
			r = append(r, "null"...)
		} else {
			r = append(r, ",\"Items\":"...)
			r = append(r, 91)
			ai113 := len(r)
			for i112 := range me.Items {
				if nil != me.Items[i112] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Items[i112].preview_MarshalJSON()
						if e == nil {
							r = append(r, sl...)
						} else {
							err = e
							return
						}
					}
				} else {
					r = append(r, 44)
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
			if r[ai113] == 44 {
				r = append(r[:ai113], r[(ai113+1):]...)
			}
		}
		if len(me.Warns) != 0 {
			r = append(r, ",\"Warns\":"...)
			r = append(r, 91)
			ai121 := len(r)
			for i120 := range me.Warns {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.Warns[i120])...)
				}
			}
			r = append(r, 93)
			if r[ai121] == 44 {
				r = append(r[:ai121], r[(ai121+1):]...)
			}
		}
		if len(me.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Desc)...)
		}
		if len(me.Url) != 0 {
			r = append(r, ",\"Url\":"...)
			r = append(r, pkg__strconv.Quote(me.Url)...)
		}
		r = append(r, 125)
		if r[si97] == 44 {
			r = append(r[:si97], r[(si97+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuResponse) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si122 := len(r)
		if nil != me.SubMenu {
			{
				r = append(r, ",\"SubMenu\":"...)
				var e error
				var sl []byte
				sl, e = me.SubMenu.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if len(me.WebsiteURL) != 0 {
			r = append(r, ",\"WebsiteURL\":"...)
			r = append(r, pkg__strconv.Quote(me.WebsiteURL)...)
		}
		if len(me.NoteInfo) != 0 {
			r = append(r, ",\"NoteInfo\":"...)
			r = append(r, pkg__strconv.Quote(me.NoteInfo)...)
		}
		if len(me.NoteWarn) != 0 {
			r = append(r, ",\"NoteWarn\":"...)
			r = append(r, pkg__strconv.Quote(me.NoteWarn)...)
		}
		if len(me.UxActionLabel) != 0 {
			r = append(r, ",\"UxActionLabel\":"...)
			r = append(r, pkg__strconv.Quote(me.UxActionLabel)...)
		}
		if len(me.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			var e error
			var sl []byte
			sl, e = me.Refs.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si122] == 44 {
			r = append(r[:si122], r[(si122+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntel) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si135 := len(r)
		if len(me.SrcIntels.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai137 := len(r)
			for i136 := range me.SrcIntels.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcIntels.InfoTips[i136].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai137] == 44 {
				r = append(r[:ai137], r[(ai137+1):]...)
			}
		}
		if len(me.SrcIntels.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			var e error
			var sl []byte
			sl, e = me.SrcIntels.Refs.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if nil != me.Sig {
			{
				r = append(r, ",\"Sig\":"...)
				var e error
				var sl []byte
				sl, e = me.Sig.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if len(me.Cmpl) != 0 {
			r = append(r, ",\"Cmpl\":"...)
			var e error
			var sl []byte
			sl, e = me.Cmpl.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if len(me.Syms) != 0 {
			r = append(r, ",\"Syms\":"...)
			var e error
			var sl []byte
			sl, e = me.Syms.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if len(me.Anns) != 0 {
			r = append(r, ",\"Anns\":"...)
			r = append(r, 91)
			ai169 := len(r)
			for i168 := range me.Anns {
				if nil != me.Anns[i168] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Anns[i168].preview_MarshalJSON()
						if e == nil {
							r = append(r, sl...)
						} else {
							err = e
							return
						}
					}
				} else {
					r = append(r, 44)
					r = append(r, "null"...)
				}
			}
			r = append(r, 93)
			if r[ai169] == 44 {
				r = append(r[:ai169], r[(ai169+1):]...)
			}
		}
		r = append(r, 125)
		if r[si135] == 44 {
			r = append(r[:si135], r[(si135+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Caddy) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si176 := len(r)
		if len(me.ID) != 0 {
			r = append(r, ",\"ID\":"...)
			r = append(r, pkg__strconv.Quote(me.ID)...)
		}
		if len(me.LangID) != 0 {
			r = append(r, ",\"LangID\":"...)
			r = append(r, pkg__strconv.Quote(me.LangID)...)
		}
		{
			r = append(r, ",\"Icon\":"...)
			r = append(r, pkg__strconv.Quote(me.Icon)...)
		}
		if len(me.Title) != 0 {
			r = append(r, ",\"Title\":"...)
			r = append(r, pkg__strconv.Quote(me.Title)...)
		}
		r = append(r, ",\"Status\":"...)
		r = append(r, 123)
		si177 := len(r)
		{
			r = append(r, ",\"Flag\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Status.Flag), 10)...)
		}
		if len(me.Status.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Status.Desc)...)
		}
		r = append(r, 125)
		if r[si177] == 44 {
			r = append(r[:si177], r[(si177+1):]...)
		}
		if len(me.Details) != 0 {
			r = append(r, ",\"Details\":"...)
			r = append(r, pkg__strconv.Quote(me.Details)...)
		}
		if len(me.UxActionID) != 0 {
			r = append(r, ",\"UxActionID\":"...)
			r = append(r, pkg__strconv.Quote(me.UxActionID)...)
		}
		if me.ShowTitle {
			r = append(r, ",\"ShowTitle\":"...)
			r = append(r, pkg__strconv.FormatBool(me.ShowTitle)...)
		}
		r = append(r, 125)
		if r[si176] == 44 {
			r = append(r[:si176], r[(si176+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *DiagFixUps) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si178 := len(r)
		{
			r = append(r, ",\"FilePath\":"...)
			r = append(r, pkg__strconv.Quote(me.FilePath)...)
		}
		if nil == me.Desc {
			r = append(r, ",\"Desc\":"...)
			r = append(r, "null"...)
		} else {
			r = append(r, ",\"Desc\":"...)
			r = append(r, 123)
			mi179 := len(r)
			for mk180, mv181 := range me.Desc {
				if nil == mv181 {
					{
						r = append(r, 44)
						r = append(r, pkg__strconv.Quote(mk180)...)
						r = append(r, 58)
					}
					r = append(r, "null"...)
				} else {
					{
						r = append(r, 44)
						r = append(r, pkg__strconv.Quote(mk180)...)
						r = append(r, 58)
					}
					r = append(r, 91)
					ai183 := len(r)
					for i182 := range mv181 {
						{
							r = append(r, 44)
							r = append(r, pkg__strconv.Quote(mv181[i182])...)
						}
					}
					r = append(r, 93)
					if r[ai183] == 44 {
						r = append(r[:ai183], r[(ai183+1):]...)
					}
				}
			}
			r = append(r, 125)
			if r[mi179] == 44 {
				r = append(r[:mi179], r[(mi179+1):]...)
			}
		}
		{
			r = append(r, ",\"Edits\":"...)
			var e error
			var sl []byte
			sl, e = me.Edits.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if nil == me.Dropped {
			r = append(r, ",\"Dropped\":"...)
			r = append(r, "null"...)
		} else {
			r = append(r, ",\"Dropped\":"...)
			r = append(r, 91)
			ai191 := len(r)
			for i190 := range me.Dropped {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Dropped[i190].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai191] == 44 {
				r = append(r[:ai191], r[(ai191+1):]...)
			}
		}
		r = append(r, 125)
		if r[si178] == 44 {
			r = append(r[:si178], r[(si178+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *DiagItem) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si198 := len(r)
		if len(me.Cat) != 0 {
			r = append(r, ",\"Cat\":"...)
			r = append(r, pkg__strconv.Quote(me.Cat)...)
		}
		{
			r = append(r, ",\"Loc\":"...)
			var e error
			var sl []byte
			sl, e = me.Loc.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		{
			r = append(r, ",\"Msg\":"...)
			r = append(r, pkg__strconv.Quote(me.Msg)...)
		}
		if len(me.SrcActions) != 0 {
			r = append(r, ",\"SrcActions\":"...)
			r = append(r, 91)
			ai206 := len(r)
			for i205 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i205].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai206] == 44 {
				r = append(r[:ai206], r[(ai206+1):]...)
			}
		}
		if me.StickyAuto {
			r = append(r, ",\"Sticky\":"...)
			r = append(r, pkg__strconv.FormatBool(me.StickyAuto)...)
		}
		if len(me.Tags) != 0 {
			r = append(r, ",\"Tags\":"...)
			r = append(r, 91)
			ai214 := len(r)
			for i213 := range me.Tags {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.FormatInt((int64)(me.Tags[i213]), 10)...)
				}
			}
			r = append(r, 93)
			if r[ai214] == 44 {
				r = append(r[:ai214], r[(ai214+1):]...)
			}
		}
		r = append(r, 125)
		if r[si198] == 44 {
			r = append(r[:si198], r[(si198+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me DiagItems) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai216 := len(r)
		for i215 := range me {
			if nil != me[i215] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i215].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			} else {
				r = append(r, 44)
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
		if r[ai216] == 44 {
			r = append(r[:ai216], r[(ai216+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me DiagItemsBy) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		mi223 := len(r)
		for mk224, mv225 := range me {
			{
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(mk224)...)
					r = append(r, 58)
				}
				var e error
				var sl []byte
				sl, e = mv225.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 125)
		if r[mi223] == 44 {
			r = append(r[:mi223], r[(mi223+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *EditorAction) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si232 := len(r)
		{
			r = append(r, ",\"title\":"...)
			r = append(r, pkg__strconv.Quote(me.Title)...)
		}
		{
			r = append(r, ",\"command\":"...)
			r = append(r, pkg__strconv.Quote(me.Cmd)...)
		}
		if len(me.Hint) != 0 {
			r = append(r, ",\"tooltip\":"...)
			r = append(r, pkg__strconv.Quote(me.Hint)...)
		}
		if len(me.Arguments) != 0 {
			r = append(r, ",\"arguments\":"...)
			r = append(r, 91)
			ai234 := len(r)
			for i233 := range me.Arguments {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					j, _ := me.Arguments[i233].(pkg__encoding_json.Marshaler)
					if j != nil {
						sl, e = j.MarshalJSON()
					} else {
						v239, ok240 := me.Arguments[i233].(map[string]interface{})
						if ok240 {
							sl, e = __gent__jsonMarshal_mapsstring_interface____(v239)
						} else {
							v237, ok238 := me.Arguments[i233].([]string)
							if ok238 {
								sl, e = __gent__jsonMarshal_s_string(v237)
							} else {
								v235, ok236 := me.Arguments[i233].(string)
								if ok236 {
									sl, e = __gent__jsonMarshal_string(v235)
								} else {
									println("JSON.MARSHAL:", pkg__fmt.Sprintf("%T", me.Arguments[i233]))
									sl, e = pkg__encoding_json.Marshal(me.Arguments[i233])
								}
							}
						}
					}
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai234] == 44 {
				r = append(r[:ai234], r[(ai234+1):]...)
			}
		}
		r = append(r, 125)
		if r[si232] == 44 {
			r = append(r[:si232], r[(si232+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *ExtrasItem) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si241 := len(r)
		{
			r = append(r, ",\"id\":"...)
			r = append(r, pkg__strconv.Quote(me.ID)...)
		}
		{
			r = append(r, ",\"label\":"...)
			r = append(r, pkg__strconv.Quote(me.Label)...)
		}
		{
			r = append(r, ",\"description\":"...)
			r = append(r, pkg__strconv.Quote(me.Desc)...)
		}
		if len(me.Detail) != 0 {
			r = append(r, ",\"detail\":"...)
			r = append(r, pkg__strconv.Quote(me.Detail)...)
		}
		if len(me.QueryArg) != 0 {
			r = append(r, ",\"arg\":"...)
			r = append(r, pkg__strconv.Quote(me.QueryArg)...)
		}
		if len(me.FilePos) != 0 {
			r = append(r, ",\"fPos\":"...)
			r = append(r, pkg__strconv.Quote(me.FilePos)...)
		}
		r = append(r, 125)
		if r[si241] == 44 {
			r = append(r[:si241], r[(si241+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Menu) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si242 := len(r)
		if len(me.Desc) != 0 {
			r = append(r, ",\"desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Desc)...)
		}
		if me.TopLevel {
			r = append(r, ",\"topLevel\":"...)
			r = append(r, pkg__strconv.FormatBool(me.TopLevel)...)
		}
		{
			r = append(r, ",\"items\":"...)
			var e error
			var sl []byte
			sl, e = me.Items.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si242] == 44 {
			r = append(r[:si242], r[(si242+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me MenuItems) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai250 := len(r)
		for i249 := range me {
			if nil != me[i249] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i249].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			} else {
				r = append(r, 44)
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
		if r[ai250] == 44 {
			r = append(r[:ai250], r[(ai250+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuItem) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si257 := len(r)
		if me.IpcID != 0 {
			r = append(r, ",\"ii\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
		}
		if me.IpcArgs != nil {
			r = append(r, ",\"ia\":"...)
			var e error
			var sl []byte
			j, _ := me.IpcArgs.(pkg__encoding_json.Marshaler)
			if j != nil {
				sl, e = j.MarshalJSON()
			} else {
				v262, ok263 := me.IpcArgs.(map[string]interface{})
				if ok263 {
					sl, e = __gent__jsonMarshal_mapsstring_interface____(v262)
				} else {
					v260, ok261 := me.IpcArgs.([]string)
					if ok261 {
						sl, e = __gent__jsonMarshal_s_string(v260)
					} else {
						v258, ok259 := me.IpcArgs.(string)
						if ok259 {
							sl, e = __gent__jsonMarshal_string(v258)
						} else {
							println("JSON.MARSHAL:", pkg__fmt.Sprintf("%T", me.IpcArgs))
							sl, e = pkg__encoding_json.Marshal(me.IpcArgs)
						}
					}
				}
			}
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if len(me.Category) != 0 {
			r = append(r, ",\"c\":"...)
			r = append(r, pkg__strconv.Quote(me.Category)...)
		}
		{
			r = append(r, ",\"t\":"...)
			r = append(r, pkg__strconv.Quote(me.Title)...)
		}
		if len(me.Desc) != 0 {
			r = append(r, ",\"d\":"...)
			r = append(r, pkg__strconv.Quote(me.Desc)...)
		}
		if len(me.Hint) != 0 {
			r = append(r, ",\"h\":"...)
			r = append(r, pkg__strconv.Quote(me.Hint)...)
		}
		if len(me.Confirm) != 0 {
			r = append(r, ",\"q\":"...)
			r = append(r, pkg__strconv.Quote(me.Confirm)...)
		}
		r = append(r, 125)
		if r[si257] == 44 {
			r = append(r[:si257], r[(si257+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuItemArgPrompt) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si264 := len(r)
		if len(me.Prompt) != 0 {
			r = append(r, ",\"prompt\":"...)
			r = append(r, pkg__strconv.Quote(me.Prompt)...)
		}
		if len(me.Placeholder) != 0 {
			r = append(r, ",\"placeHolder\":"...)
			r = append(r, pkg__strconv.Quote(me.Placeholder)...)
		}
		if len(me.Value) != 0 {
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		r = append(r, 125)
		if r[si264] == 44 {
			r = append(r[:si264], r[(si264+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcAnnotaction) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si265 := len(r)
		{
			r = append(r, ",\"Range\":"...)
			var e error
			var sl []byte
			sl, e = me.Range.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		{
			r = append(r, ",\"Title\":"...)
			r = append(r, pkg__strconv.Quote(me.Title)...)
		}
		if len(me.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Desc)...)
		}
		{
			r = append(r, ",\"CmdName\":"...)
			r = append(r, pkg__strconv.Quote(me.CmdName)...)
		}
		r = append(r, 125)
		if r[si265] == 44 {
			r = append(r[:si265], r[(si265+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcInfoTip) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si272 := len(r)
		{
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if len(me.Language) != 0 {
			r = append(r, ",\"language\":"...)
			r = append(r, pkg__strconv.Quote(me.Language)...)
		}
		r = append(r, 125)
		if r[si272] == 44 {
			r = append(r[:si272], r[(si272+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelCompl) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si273 := len(r)
		{
			r = append(r, ",\"kind\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Kind), 10)...)
		}
		{
			r = append(r, ",\"label\":"...)
			r = append(r, pkg__strconv.Quote(me.Label)...)
		}
		if nil != me.Documentation {
			{
				r = append(r, ",\"documentation\":"...)
				var e error
				var sl []byte
				sl, e = me.Documentation.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if len(me.Detail) != 0 {
			r = append(r, ",\"detail\":"...)
			r = append(r, pkg__strconv.Quote(me.Detail)...)
		}
		if len(me.SortText) != 0 {
			r = append(r, ",\"sortText\":"...)
			r = append(r, pkg__strconv.Quote(me.SortText)...)
		}
		r = append(r, 125)
		if r[si273] == 44 {
			r = append(r[:si273], r[(si273+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcIntelCompls) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai281 := len(r)
		for i280 := range me {
			if nil != me[i280] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i280].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			} else {
				r = append(r, 44)
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
		if r[ai281] == 44 {
			r = append(r[:ai281], r[(ai281+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntels) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si288 := len(r)
		if len(me.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai290 := len(r)
			for i289 := range me.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.InfoTips[i289].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai290] == 44 {
				r = append(r[:ai290], r[(ai290+1):]...)
			}
		}
		if len(me.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			var e error
			var sl []byte
			sl, e = me.Refs.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si288] == 44 {
			r = append(r[:si288], r[(si288+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelDoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si303 := len(r)
		if len(me.Value) != 0 {
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if me.IsTrusted {
			r = append(r, ",\"isTrusted\":"...)
			r = append(r, pkg__strconv.FormatBool(me.IsTrusted)...)
		}
		r = append(r, 125)
		if r[si303] == 44 {
			r = append(r[:si303], r[(si303+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigHelp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si304 := len(r)
		{
			r = append(r, ",\"activeSignature\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.ActiveSignature), 10)...)
		}
		if me.ActiveParameter != 0 {
			r = append(r, ",\"activeParameter\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.ActiveParameter), 10)...)
		}
		if len(me.Signatures) != 0 {
			r = append(r, ",\"signatures\":"...)
			r = append(r, 91)
			ai306 := len(r)
			for i305 := range me.Signatures {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Signatures[i305].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai306] == 44 {
				r = append(r[:ai306], r[(ai306+1):]...)
			}
		}
		r = append(r, 125)
		if r[si304] == 44 {
			r = append(r[:si304], r[(si304+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigInfo) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si313 := len(r)
		{
			r = append(r, ",\"label\":"...)
			r = append(r, pkg__strconv.Quote(me.Label)...)
		}
		{
			r = append(r, ",\"documentation\":"...)
			var e error
			var sl []byte
			sl, e = me.Documentation.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		if nil == me.Parameters {
			r = append(r, ",\"parameters\":"...)
			r = append(r, "null"...)
		} else {
			r = append(r, ",\"parameters\":"...)
			r = append(r, 91)
			ai321 := len(r)
			for i320 := range me.Parameters {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Parameters[i320].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai321] == 44 {
				r = append(r[:ai321], r[(ai321+1):]...)
			}
		}
		r = append(r, 125)
		if r[si313] == 44 {
			r = append(r[:si313], r[(si313+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigParam) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si328 := len(r)
		{
			r = append(r, ",\"label\":"...)
			r = append(r, pkg__strconv.Quote(me.Label)...)
		}
		{
			r = append(r, ",\"documentation\":"...)
			var e error
			var sl []byte
			sl, e = me.Documentation.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si328] == 44 {
			r = append(r[:si328], r[(si328+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcLenses) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai336 := len(r)
		for i335 := range me {
			if nil != me[i335] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i335].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			} else {
				r = append(r, 44)
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
		if r[ai336] == 44 {
			r = append(r[:ai336], r[(ai336+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLens) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si343 := len(r)
		{
			r = append(r, ",\"e\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLoc.Flag), 10)...)
		}
		if len(me.SrcLoc.FilePath) != 0 {
			r = append(r, ",\"f\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcLoc.FilePath)...)
		}
		if nil != me.SrcLoc.Pos {
			{
				r = append(r, ",\"p\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcLoc.Pos.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.SrcLoc.Range {
			{
				r = append(r, ",\"r\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcLoc.Range.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
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
		r = append(r, 125)
		if r[si343] == 44 {
			r = append(r[:si343], r[(si343+1):]...)
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLens) preview_UnmarshalJSON(v []byte) (err error) { ; return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si356 := len(r)
		{
			r = append(r, ",\"e\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Flag), 10)...)
		}
		if len(me.FilePath) != 0 {
			r = append(r, ",\"f\":"...)
			r = append(r, pkg__strconv.Quote(me.FilePath)...)
		}
		if nil != me.Pos {
			{
				r = append(r, ",\"p\":"...)
				var e error
				var sl []byte
				sl, e = me.Pos.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.Range {
			{
				r = append(r, ",\"r\":"...)
				var e error
				var sl []byte
				sl, e = me.Range.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 125)
		if r[si356] == 44 {
			r = append(r[:si356], r[(si356+1):]...)
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLoc) preview_UnmarshalJSON(v []byte) (err error) { ; return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcLocs) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai370 := len(r)
		for i369 := range me {
			if nil != me[i369] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i369].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			} else {
				r = append(r, 44)
				r = append(r, "null"...)
			}
		}
		r = append(r, 93)
		if r[ai370] == 44 {
			r = append(r[:ai370], r[(ai370+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcModEdit) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si377 := len(r)
		if nil != me.At {
			{
				r = append(r, ",\"At\":"...)
				var e error
				var sl []byte
				sl, e = me.At.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		} else {
			r = append(r, ",\"At\":"...)
			r = append(r, "null"...)
		}
		{
			r = append(r, ",\"Val\":"...)
			r = append(r, pkg__strconv.Quote(me.Val)...)
		}
		r = append(r, 125)
		if r[si377] == 44 {
			r = append(r[:si377], r[(si377+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcModEdits) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai385 := len(r)
		for i384 := range me {
			{
				r = append(r, 44)
				var e error
				var sl []byte
				sl, e = me[i384].preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 93)
		if r[ai385] == 44 {
			r = append(r[:ai385], r[(ai385+1):]...)
		}
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcPos) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si392 := len(r)
		if me.Ln != 0 {
			r = append(r, ",\"l\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Ln), 10)...)
		}
		if me.Col != 0 {
			r = append(r, ",\"c\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Col), 10)...)
		}
		if me.Off != 0 {
			r = append(r, ",\"o\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Off), 10)...)
		}
		r = append(r, 125)
		if r[si392] == 44 {
			r = append(r[:si392], r[(si392+1):]...)
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcPos) preview_UnmarshalJSON(v []byte) (err error) { ; return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcRange) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si393 := len(r)
		{
			r = append(r, ",\"s\":"...)
			var e error
			var sl []byte
			sl, e = me.Start.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		{
			r = append(r, ",\"e\":"...)
			var e error
			var sl []byte
			sl, e = me.End.preview_MarshalJSON()
			if e == nil {
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si393] == 44 {
			r = append(r[:si393], r[(si393+1):]...)
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) preview_UnmarshalJSON(v []byte) (err error) { ; return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) preview_MarshalJSON() (r []byte, err error) {
	panic("WorkspaceChanges")
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *WorkspaceChanges) preview_UnmarshalJSON(v []byte) (err error) { ; return }

func __gent__jsonMarshal_string(v string) (r []byte, err error) {
	{
		r = append(r, pkg__strconv.Quote(v)...)
	}
	return
}

func __gent__jsonMarshal_s_string(v []string) (r []byte, err error) {
	if nil == v {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai3 := len(r)
		for i2 := range v {
			{
				r = append(r, 44)
				r = append(r, pkg__strconv.Quote(v[i2])...)
			}
		}
		r = append(r, 93)
		if r[ai3] == 44 {
			r = append(r[:ai3], r[(ai3+1):]...)
		}
	}
	return
}

func __gent__jsonMarshal_mapsstring_interface____(v map[string]interface{}) (r []byte, err error) {
	if nil == v {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		mi4 := len(r)
		for mk5, mv6 := range v {
			{
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(mk5)...)
					r = append(r, 58)
				}
				var e error
				var sl []byte
				j, _ := mv6.(pkg__encoding_json.Marshaler)
				if j != nil {
					sl, e = j.MarshalJSON()
				} else {
					v11, ok12 := mv6.(map[string]interface{})
					if ok12 {
						sl, e = __gent__jsonMarshal_mapsstring_interface____(v11)
					} else {
						v9, ok10 := mv6.([]string)
						if ok10 {
							sl, e = __gent__jsonMarshal_s_string(v9)
						} else {
							v7, ok8 := mv6.(string)
							if ok8 {
								sl, e = __gent__jsonMarshal_string(v7)
							} else {
								println("JSON.MARSHAL:", pkg__fmt.Sprintf("%T", mv6))
								sl, e = pkg__encoding_json.Marshal(mv6)
							}
						}
					}
				}
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 125)
		if r[mi4] == 44 {
			r = append(r[:mi4], r[(mi4+1):]...)
		}
	}
	return
}
