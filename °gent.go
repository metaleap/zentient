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

func (me *IpcReq) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ri"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.ReqID = (int64)(v1.(float64))
		} else {
			me.ReqID = 0
		}
	}
	v4, o5 := v["ii"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.IpcID = (IpcIDs)(v4.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v7, o8 := v["ia"]
	if o8 {
		println(v7)
		if nil != v7 {
		} else {
			me.IpcArgs = nil
		}
	}
	v10, o11 := v["projUpd"]
	if o11 {
		println(v10)
		if nil != v10 {
			v13 := v10.(map[string]interface{})
			if v13 == nil {
				me.ProjUpd = nil
			} else {
				if nil == me.ProjUpd {
					me.ProjUpd = new(WorkspaceChanges)
				}
				me.ProjUpd.jsonUnmarshal_FromAny(v13)
			}
		} else {
			me.ProjUpd = nil
		}
	}
	v14, o15 := v["srcLens"]
	if o15 {
		println(v14)
		if nil != v14 {
			v17 := v14.(map[string]interface{})
			if v17 == nil {
				me.SrcLens = nil
			} else {
				if nil == me.SrcLens {
					me.SrcLens = new(SrcLens)
				}
				me.SrcLens.jsonUnmarshal_FromAny(v17)
			}
		} else {
			me.SrcLens = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcReq) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 5)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *IpcResp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si1 := len(r)
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
			ai21 := len(r)
			for i20 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i20].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai21] == 44 {
				r = append(r[:ai21], r[(ai21+1):]...)
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
			sl, e := __gent__jsonMarshal_interface____(me.Val)
			if e == nil {
				r = append(r, ",\"val\":"...)
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *IpcResp) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ii"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.IpcID = (IpcIDs)(v1.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v4, o5 := v["ri"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.ReqID = (int64)(v4.(float64))
		} else {
			me.ReqID = 0
		}
	}
	v7, o8 := v["err"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.ErrMsg = (string)(v7.(string))
		} else {
			me.ErrMsg = ""
		}
	}
	v10, o11 := v["sI"]
	if o11 {
		println(v10)
		if nil != v10 {
			v13 := v10.(map[string]interface{})
			if v13 == nil {
				me.SrcIntel = nil
			} else {
				if nil == me.SrcIntel {
					me.SrcIntel = new(SrcIntel)
				}
				me.SrcIntel.jsonUnmarshal_FromAny(v13)
			}
		} else {
			me.SrcIntel = nil
		}
	}
	v14, o15 := v["srcDiags"]
	if o15 {
		println(v14)
		if nil != v14 {
			v17 := v14.(map[string]interface{})
			if v17 == nil {
				me.SrcDiags = nil
			} else {
				if nil == me.SrcDiags {
					me.SrcDiags = new(Diags)
				}
				me.SrcDiags.jsonUnmarshal_FromAny(v17)
			}
		} else {
			me.SrcDiags = nil
		}
	}
	v18, o19 := v["srcMods"]
	if o19 {
		println(v18)
		if nil != v18 {
			v21 := v18.([]interface{})
			if v21 == nil {
				me.SrcMods = nil
			} else {
				if false {
				}
				me.SrcMods.jsonUnmarshal_FromAny(v21)
			}
		} else {
			me.SrcMods = nil
		}
	}
	v22, o23 := v["srcActions"]
	if o23 {
		println(v22)
		if nil != v22 {
			s25 := v22.([]interface{})
			if s25 == nil {
				me.SrcActions = nil
			} else {
				if len(me.SrcActions) >= len(s25) {
					me.SrcActions = me.SrcActions[0:len(s25)]
				} else {
					me.SrcActions = make([]EditorAction, len(s25))
				}
				for si26, sv27 := range s25 {
					println(sv27)
					if nil != sv27 {
						v29 := sv27.(map[string]interface{})
						if v29 == nil {
							var z28 EditorAction
							me.SrcActions[si26] = z28
						} else {
							if false {
							}
							me.SrcActions[si26].jsonUnmarshal_FromAny(v29)
						}
					} else {
						var z28 EditorAction
						me.SrcActions[si26] = z28
					}
				}
			}
		} else {
			me.SrcActions = nil
		}
	}
	v30, o31 := v["extras"]
	if o31 {
		println(v30)
		if nil != v30 {
			v33 := v30.(map[string]interface{})
			if v33 == nil {
				me.Extras = nil
			} else {
				if nil == me.Extras {
					me.Extras = new(Extras)
				}
				me.Extras.jsonUnmarshal_FromAny(v33)
			}
		} else {
			me.Extras = nil
		}
	}
	v34, o35 := v["menu"]
	if o35 {
		println(v34)
		if nil != v34 {
			v37 := v34.(map[string]interface{})
			if v37 == nil {
				me.Menu = nil
			} else {
				if nil == me.Menu {
					me.Menu = new(MenuResponse)
				}
				me.Menu.jsonUnmarshal_FromAny(v37)
			}
		} else {
			me.Menu = nil
		}
	}
	v38, o39 := v["caddy"]
	if o39 {
		println(v38)
		if nil != v38 {
			v41 := v38.(map[string]interface{})
			if v41 == nil {
				me.CaddyUpdate = nil
			} else {
				if nil == me.CaddyUpdate {
					me.CaddyUpdate = new(Caddy)
				}
				me.CaddyUpdate.jsonUnmarshal_FromAny(v41)
			}
		} else {
			me.CaddyUpdate = nil
		}
	}
	v42, o43 := v["val"]
	if o43 {
		println(v42)
		if nil != v42 {
		} else {
			me.Val = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcResp) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 11)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
			ai9 := len(r)
			for i8 := range me.FixUps {
				if nil != me.FixUps[i8] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.FixUps[i8].preview_MarshalJSON()
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
			if r[ai9] == 44 {
				r = append(r[:ai9], r[(ai9+1):]...)
			}
		}
		{
			r = append(r, ",\"LangID\":"...)
			r = append(r, pkg__strconv.Quote(me.LangID)...)
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *Diags) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["All"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				me.All = nil
			} else {
				if nil == me.All {
					me.All = make(DiagItemsBy, len(v4))
				}
				me.All.jsonUnmarshal_FromAny(v4)
			}
		} else {
			me.All = nil
		}
	}
	v5, o6 := v["FixUps"]
	if o6 {
		println(v5)
		if nil != v5 {
			s8 := v5.([]interface{})
			if s8 == nil {
				me.FixUps = nil
			} else {
				if len(me.FixUps) >= len(s8) {
					me.FixUps = me.FixUps[0:len(s8)]
				} else {
					me.FixUps = make([]*DiagFixUps, len(s8))
				}
				for si9, sv10 := range s8 {
					println(sv10)
					if nil != sv10 {
						v12 := sv10.(map[string]interface{})
						if v12 == nil {
							me.FixUps[si9] = nil
						} else {
							if nil == me.FixUps[si9] {
								me.FixUps[si9] = new(DiagFixUps)
							}
							me.FixUps[si9].jsonUnmarshal_FromAny(v12)
						}
					} else {
						me.FixUps[si9] = nil
					}
				}
			}
		} else {
			me.FixUps = nil
		}
	}
	v13, o14 := v["LangID"]
	if o14 {
		println(v13)
		if nil != v13 {
			me.LangID = (string)(v13.(string))
		} else {
			me.LangID = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Diags) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 3)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
		if len(me.SrcIntels.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai3 := len(r)
			for i2 := range me.SrcIntels.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcIntels.InfoTips[i2].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai3] == 44 {
				r = append(r[:ai3], r[(ai3+1):]...)
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
			ai17 := len(r)
			for i16 := range me.Items {
				if nil != me.Items[i16] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Items[i16].preview_MarshalJSON()
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
			if r[ai17] == 44 {
				r = append(r[:ai17], r[(ai17+1):]...)
			}
		}
		if len(me.Warns) != 0 {
			sl, e := __gent__jsonMarshal_s_string(me.Warns)
			if e == nil {
				r = append(r, ",\"Warns\":"...)
				r = append(r, sl...)
			} else {
				err = e
				return
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *Extras) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcIntels"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				var z3 SrcIntels
				me.SrcIntels = z3
			} else {
				if false {
				}
				me.SrcIntels.jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcIntels
			me.SrcIntels = z3
		}
	}
	v5, o6 := v["Items"]
	if o6 {
		println(v5)
		if nil != v5 {
			s8 := v5.([]interface{})
			if s8 == nil {
				me.Items = nil
			} else {
				if len(me.Items) >= len(s8) {
					me.Items = me.Items[0:len(s8)]
				} else {
					me.Items = make([]*ExtrasItem, len(s8))
				}
				for si9, sv10 := range s8 {
					println(sv10)
					if nil != sv10 {
						v12 := sv10.(map[string]interface{})
						if v12 == nil {
							me.Items[si9] = nil
						} else {
							if nil == me.Items[si9] {
								me.Items[si9] = new(ExtrasItem)
							}
							me.Items[si9].jsonUnmarshal_FromAny(v12)
						}
					} else {
						me.Items[si9] = nil
					}
				}
			}
		} else {
			me.Items = nil
		}
	}
	v13, o14 := v["Warns"]
	if o14 {
		println(v13)
		if nil != v13 {
			s16 := v13.([]interface{})
			if s16 == nil {
				me.Warns = nil
			} else {
				if len(me.Warns) >= len(s16) {
					me.Warns = me.Warns[0:len(s16)]
				} else {
					me.Warns = make([]string, len(s16))
				}
				for si17, sv18 := range s16 {
					println(sv18)
					if nil != sv18 {
						me.Warns[si17] = (string)(sv18.(string))
					} else {
						me.Warns[si17] = ""
					}
				}
			}
		} else {
			me.Warns = nil
		}
	}
	v20, o21 := v["Desc"]
	if o21 {
		println(v20)
		if nil != v20 {
			me.Desc = (string)(v20.(string))
		} else {
			me.Desc = ""
		}
	}
	v23, o24 := v["Url"]
	if o24 {
		println(v23)
		if nil != v23 {
			me.Url = (string)(v23.(string))
		} else {
			me.Url = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Extras) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 5)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *MenuResponse) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SubMenu"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				me.SubMenu = nil
			} else {
				if nil == me.SubMenu {
					me.SubMenu = new(Menu)
				}
				me.SubMenu.jsonUnmarshal_FromAny(v4)
			}
		} else {
			me.SubMenu = nil
		}
	}
	v5, o6 := v["WebsiteURL"]
	if o6 {
		println(v5)
		if nil != v5 {
			me.WebsiteURL = (string)(v5.(string))
		} else {
			me.WebsiteURL = ""
		}
	}
	v8, o9 := v["NoteInfo"]
	if o9 {
		println(v8)
		if nil != v8 {
			me.NoteInfo = (string)(v8.(string))
		} else {
			me.NoteInfo = ""
		}
	}
	v11, o12 := v["NoteWarn"]
	if o12 {
		println(v11)
		if nil != v11 {
			me.NoteWarn = (string)(v11.(string))
		} else {
			me.NoteWarn = ""
		}
	}
	v14, o15 := v["UxActionLabel"]
	if o15 {
		println(v14)
		if nil != v14 {
			me.UxActionLabel = (string)(v14.(string))
		} else {
			me.UxActionLabel = ""
		}
	}
	v17, o18 := v["Refs"]
	if o18 {
		println(v17)
		if nil != v17 {
			v20 := v17.([]interface{})
			if v20 == nil {
				me.Refs = nil
			} else {
				if false {
				}
				me.Refs.jsonUnmarshal_FromAny(v20)
			}
		} else {
			me.Refs = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuResponse) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 6)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
		if len(me.SrcIntels.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai3 := len(r)
			for i2 := range me.SrcIntels.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcIntels.InfoTips[i2].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai3] == 44 {
				r = append(r[:ai3], r[(ai3+1):]...)
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
			ai35 := len(r)
			for i34 := range me.Anns {
				if nil != me.Anns[i34] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Anns[i34].preview_MarshalJSON()
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
			if r[ai35] == 44 {
				r = append(r[:ai35], r[(ai35+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntel) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcIntels"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				var z3 SrcIntels
				me.SrcIntels = z3
			} else {
				if false {
				}
				me.SrcIntels.jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcIntels
			me.SrcIntels = z3
		}
	}
	v5, o6 := v["Sig"]
	if o6 {
		println(v5)
		if nil != v5 {
			v8 := v5.(map[string]interface{})
			if v8 == nil {
				me.Sig = nil
			} else {
				if nil == me.Sig {
					me.Sig = new(SrcIntelSigHelp)
				}
				me.Sig.jsonUnmarshal_FromAny(v8)
			}
		} else {
			me.Sig = nil
		}
	}
	v9, o10 := v["Cmpl"]
	if o10 {
		println(v9)
		if nil != v9 {
			v12 := v9.([]interface{})
			if v12 == nil {
				me.Cmpl = nil
			} else {
				if false {
				}
				me.Cmpl.jsonUnmarshal_FromAny(v12)
			}
		} else {
			me.Cmpl = nil
		}
	}
	v13, o14 := v["Syms"]
	if o14 {
		println(v13)
		if nil != v13 {
			v16 := v13.([]interface{})
			if v16 == nil {
				me.Syms = nil
			} else {
				if false {
				}
				me.Syms.jsonUnmarshal_FromAny(v16)
			}
		} else {
			me.Syms = nil
		}
	}
	v17, o18 := v["Anns"]
	if o18 {
		println(v17)
		if nil != v17 {
			s20 := v17.([]interface{})
			if s20 == nil {
				me.Anns = nil
			} else {
				if len(me.Anns) >= len(s20) {
					me.Anns = me.Anns[0:len(s20)]
				} else {
					me.Anns = make([]*SrcAnnotaction, len(s20))
				}
				for si21, sv22 := range s20 {
					println(sv22)
					if nil != sv22 {
						v24 := sv22.(map[string]interface{})
						if v24 == nil {
							me.Anns[si21] = nil
						} else {
							if nil == me.Anns[si21] {
								me.Anns[si21] = new(SrcAnnotaction)
							}
							me.Anns[si21].jsonUnmarshal_FromAny(v24)
						}
					} else {
						me.Anns[si21] = nil
					}
				}
			}
		} else {
			me.Anns = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntel) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 5)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		si2 := len(r)
		{
			r = append(r, ",\"Flag\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.Status.Flag), 10)...)
		}
		if len(me.Status.Desc) != 0 {
			r = append(r, ",\"Desc\":"...)
			r = append(r, pkg__strconv.Quote(me.Status.Desc)...)
		}
		r = append(r, 125)
		if r[si2] == 44 {
			r = append(r[:si2], r[(si2+1):]...)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *Caddy) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ID"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.ID = (string)(v1.(string))
		} else {
			me.ID = ""
		}
	}
	v4, o5 := v["LangID"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.LangID = (string)(v4.(string))
		} else {
			me.LangID = ""
		}
	}
	v7, o8 := v["Icon"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Icon = (string)(v7.(string))
		} else {
			me.Icon = ""
		}
	}
	v10, o11 := v["Title"]
	if o11 {
		println(v10)
		if nil != v10 {
			me.Title = (string)(v10.(string))
		} else {
			me.Title = ""
		}
	}
	v13, o14 := v["Status"]
	if o14 {
		println(v13)
		if nil != v13 {
			t16 := v13.(map[string]interface{})
			if nil == t16 {
				var z15 struct {
					Flag CaddyStatus
					Desc string `json:",omitempty"`
				}
				me.Status = z15
			} else {
				v17, o18 := t16["Flag"]
				if o18 {
					println(v17)
					if nil != v17 {
						me.Status.Flag = (CaddyStatus)(v17.(float64))
					} else {
						me.Status.Flag = 0
					}
				}
				v20, o21 := t16["Desc"]
				if o21 {
					println(v20)
					if nil != v20 {
						me.Status.Desc = (string)(v20.(string))
					} else {
						me.Status.Desc = ""
					}
				}
			}
		} else {
			var z15 struct {
				Flag CaddyStatus
				Desc string `json:",omitempty"`
			}
			me.Status = z15
		}
	}
	v23, o24 := v["Details"]
	if o24 {
		println(v23)
		if nil != v23 {
			me.Details = (string)(v23.(string))
		} else {
			me.Details = ""
		}
	}
	v26, o27 := v["UxActionID"]
	if o27 {
		println(v26)
		if nil != v26 {
			me.UxActionID = (string)(v26.(string))
		} else {
			me.UxActionID = ""
		}
	}
	v29, o30 := v["ShowTitle"]
	if o30 {
		println(v29)
		if nil != v29 {
			me.ShowTitle = (bool)(v29.(bool))
		} else {
			me.ShowTitle = false
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Caddy) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 11)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
			mi2 := len(r)
			for mk3, mv4 := range me.Desc {
				sl, e := __gent__jsonMarshal_s_string(mv4)
				if e == nil {
					{
						r = append(r, 44)
						r = append(r, pkg__strconv.Quote(mk3)...)
						r = append(r, 58)
					}
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
			r = append(r, 125)
			if r[mi2] == 44 {
				r = append(r[:mi2], r[(mi2+1):]...)
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
			ai12 := len(r)
			for i11 := range me.Dropped {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Dropped[i11].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai12] == 44 {
				r = append(r[:ai12], r[(ai12+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *DiagFixUps) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["FilePath"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.FilePath = (string)(v1.(string))
		} else {
			me.FilePath = ""
		}
	}
	v4, o5 := v["Desc"]
	if o5 {
		println(v4)
		if nil != v4 {
		} else {
			me.Desc = nil
		}
	}
	v7, o8 := v["Edits"]
	if o8 {
		println(v7)
		if nil != v7 {
			v10 := v7.([]interface{})
			if v10 == nil {
				me.Edits = nil
			} else {
				if false {
				}
				me.Edits.jsonUnmarshal_FromAny(v10)
			}
		} else {
			me.Edits = nil
		}
	}
	v11, o12 := v["Dropped"]
	if o12 {
		println(v11)
		if nil != v11 {
			s14 := v11.([]interface{})
			if s14 == nil {
				me.Dropped = nil
			} else {
				if len(me.Dropped) >= len(s14) {
					me.Dropped = me.Dropped[0:len(s14)]
				} else {
					me.Dropped = make([]SrcModEdit, len(s14))
				}
				for si15, sv16 := range s14 {
					println(sv16)
					if nil != sv16 {
						v18 := sv16.(map[string]interface{})
						if v18 == nil {
							var z17 SrcModEdit
							me.Dropped[si15] = z17
						} else {
							if false {
							}
							me.Dropped[si15].jsonUnmarshal_FromAny(v18)
						}
					} else {
						var z17 SrcModEdit
						me.Dropped[si15] = z17
					}
				}
			}
		} else {
			me.Dropped = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagFixUps) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 4)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if len(me.Rel) != 0 {
			r = append(r, ",\"Rel\":"...)
			r = append(r, 91)
			ai9 := len(r)
			for i8 := range me.Rel {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Rel[i8].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai9] == 44 {
				r = append(r[:ai9], r[(ai9+1):]...)
			}
		}
		if len(me.SrcActions) != 0 {
			r = append(r, ",\"SrcActions\":"...)
			r = append(r, 91)
			ai17 := len(r)
			for i16 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i16].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai17] == 44 {
				r = append(r[:ai17], r[(ai17+1):]...)
			}
		}
		if me.StickyAuto {
			r = append(r, ",\"Sticky\":"...)
			r = append(r, pkg__strconv.FormatBool(me.StickyAuto)...)
		}
		if len(me.Tags) != 0 {
			r = append(r, ",\"Tags\":"...)
			r = append(r, 91)
			ai25 := len(r)
			for i24 := range me.Tags {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.FormatInt((int64)(me.Tags[i24]), 10)...)
				}
			}
			r = append(r, 93)
			if r[ai25] == 44 {
				r = append(r[:ai25], r[(ai25+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *DiagItem) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["Cat"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Cat = (string)(v1.(string))
		} else {
			me.Cat = ""
		}
	}
	v4, o5 := v["Loc"]
	if o5 {
		println(v4)
		if nil != v4 {
			v7 := v4.(map[string]interface{})
			if v7 == nil {
				var z6 SrcLoc
				me.Loc = z6
			} else {
				if false {
				}
				me.Loc.jsonUnmarshal_FromAny(v7)
			}
		} else {
			var z6 SrcLoc
			me.Loc = z6
		}
	}
	v8, o9 := v["Msg"]
	if o9 {
		println(v8)
		if nil != v8 {
			me.Msg = (string)(v8.(string))
		} else {
			me.Msg = ""
		}
	}
	v11, o12 := v["Rel"]
	if o12 {
		println(v11)
		if nil != v11 {
			s14 := v11.([]interface{})
			if s14 == nil {
				me.Rel = nil
			} else {
				if len(me.Rel) >= len(s14) {
					me.Rel = me.Rel[0:len(s14)]
				} else {
					me.Rel = make([]SrcLens, len(s14))
				}
				for si15, sv16 := range s14 {
					println(sv16)
					if nil != sv16 {
						v18 := sv16.(map[string]interface{})
						if v18 == nil {
							var z17 SrcLens
							me.Rel[si15] = z17
						} else {
							if false {
							}
							me.Rel[si15].jsonUnmarshal_FromAny(v18)
						}
					} else {
						var z17 SrcLens
						me.Rel[si15] = z17
					}
				}
			}
		} else {
			me.Rel = nil
		}
	}
	v19, o20 := v["SrcActions"]
	if o20 {
		println(v19)
		if nil != v19 {
			s22 := v19.([]interface{})
			if s22 == nil {
				me.SrcActions = nil
			} else {
				if len(me.SrcActions) >= len(s22) {
					me.SrcActions = me.SrcActions[0:len(s22)]
				} else {
					me.SrcActions = make([]EditorAction, len(s22))
				}
				for si23, sv24 := range s22 {
					println(sv24)
					if nil != sv24 {
						v26 := sv24.(map[string]interface{})
						if v26 == nil {
							var z25 EditorAction
							me.SrcActions[si23] = z25
						} else {
							if false {
							}
							me.SrcActions[si23].jsonUnmarshal_FromAny(v26)
						}
					} else {
						var z25 EditorAction
						me.SrcActions[si23] = z25
					}
				}
			}
		} else {
			me.SrcActions = nil
		}
	}
	v27, o28 := v["Sticky"]
	if o28 {
		println(v27)
		if nil != v27 {
			me.StickyAuto = (bool)(v27.(bool))
		} else {
			me.StickyAuto = false
		}
	}
	v30, o31 := v["Tags"]
	if o31 {
		println(v30)
		if nil != v30 {
			s33 := v30.([]interface{})
			if s33 == nil {
				me.Tags = nil
			} else {
				if len(me.Tags) >= len(s33) {
					me.Tags = me.Tags[0:len(s33)]
				} else {
					me.Tags = make([]int, len(s33))
				}
				for si34, sv35 := range s33 {
					println(sv35)
					if nil != sv35 {
						me.Tags[si34] = (int)(sv35.(float64))
					} else {
						me.Tags[si34] = 0
					}
				}
			}
		} else {
			me.Tags = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItem) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 8)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		ai2 := len(r)
		for i1 := range me {
			if nil != me[i1] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i1].preview_MarshalJSON()
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
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me DiagItems) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me DiagItems) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		mi1 := len(r)
		for mk2, mv3 := range me {
			{
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(mk2)...)
					r = append(r, 58)
				}
				var e error
				var sl []byte
				sl, e = mv3.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 125)
		if r[mi1] == 44 {
			r = append(r[:mi1], r[(mi1+1):]...)
		}
	}
	return
}

func (me DiagItemsBy) jsonUnmarshal_FromAny(v map[string]interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me DiagItemsBy) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
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
			ai3 := len(r)
			for i2 := range me.Arguments {
				sl, e := __gent__jsonMarshal_interface____(me.Arguments[i2])
				if e == nil {
					r = append(r, 44)
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
			r = append(r, 93)
			if r[ai3] == 44 {
				r = append(r[:ai3], r[(ai3+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *EditorAction) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["title"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Title = (string)(v1.(string))
		} else {
			me.Title = ""
		}
	}
	v4, o5 := v["command"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Cmd = (string)(v4.(string))
		} else {
			me.Cmd = ""
		}
	}
	v7, o8 := v["tooltip"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Hint = (string)(v7.(string))
		} else {
			me.Hint = ""
		}
	}
	v10, o11 := v["arguments"]
	if o11 {
		println(v10)
		if nil != v10 {
			s13 := v10.([]interface{})
			if s13 == nil {
				me.Arguments = nil
			} else {
				if len(me.Arguments) >= len(s13) {
					me.Arguments = me.Arguments[0:len(s13)]
				} else {
					me.Arguments = make([]interface{}, len(s13))
				}
				for si14, sv15 := range s13 {
					println(sv15)
					if nil != sv15 {
					} else {
						me.Arguments[si14] = nil
					}
				}
			}
		} else {
			me.Arguments = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *EditorAction) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 4)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *ExtrasItem) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["id"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.ID = (string)(v1.(string))
		} else {
			me.ID = ""
		}
	}
	v4, o5 := v["label"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Label = (string)(v4.(string))
		} else {
			me.Label = ""
		}
	}
	v7, o8 := v["description"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Desc = (string)(v7.(string))
		} else {
			me.Desc = ""
		}
	}
	v10, o11 := v["detail"]
	if o11 {
		println(v10)
		if nil != v10 {
			me.Detail = (string)(v10.(string))
		} else {
			me.Detail = ""
		}
	}
	v13, o14 := v["arg"]
	if o14 {
		println(v13)
		if nil != v13 {
			me.QueryArg = (string)(v13.(string))
		} else {
			me.QueryArg = ""
		}
	}
	v16, o17 := v["fPos"]
	if o17 {
		println(v16)
		if nil != v16 {
			me.FilePos = (string)(v16.(string))
		} else {
			me.FilePos = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *ExtrasItem) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 6)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *Menu) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["desc"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Desc = (string)(v1.(string))
		} else {
			me.Desc = ""
		}
	}
	v4, o5 := v["topLevel"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.TopLevel = (bool)(v4.(bool))
		} else {
			me.TopLevel = false
		}
	}
	v7, o8 := v["items"]
	if o8 {
		println(v7)
		if nil != v7 {
			v10 := v7.([]interface{})
			if v10 == nil {
				me.Items = nil
			} else {
				if false {
				}
				me.Items.jsonUnmarshal_FromAny(v10)
			}
		} else {
			me.Items = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Menu) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 3)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		ai2 := len(r)
		for i1 := range me {
			if nil != me[i1] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i1].preview_MarshalJSON()
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
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me MenuItems) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me MenuItems) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
		if me.IpcID != 0 {
			r = append(r, ",\"ii\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
		}
		if me.IpcArgs != nil {
			sl, e := __gent__jsonMarshal_interface____(me.IpcArgs)
			if e == nil {
				r = append(r, ",\"ia\":"...)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *MenuItem) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ii"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.IpcID = (IpcIDs)(v1.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v4, o5 := v["ia"]
	if o5 {
		println(v4)
		if nil != v4 {
		} else {
			me.IpcArgs = nil
		}
	}
	v7, o8 := v["c"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Category = (string)(v7.(string))
		} else {
			me.Category = ""
		}
	}
	v10, o11 := v["t"]
	if o11 {
		println(v10)
		if nil != v10 {
			me.Title = (string)(v10.(string))
		} else {
			me.Title = ""
		}
	}
	v13, o14 := v["d"]
	if o14 {
		println(v13)
		if nil != v13 {
			me.Desc = (string)(v13.(string))
		} else {
			me.Desc = ""
		}
	}
	v16, o17 := v["h"]
	if o17 {
		println(v16)
		if nil != v16 {
			me.Hint = (string)(v16.(string))
		} else {
			me.Hint = ""
		}
	}
	v19, o20 := v["q"]
	if o20 {
		println(v19)
		if nil != v19 {
			me.Confirm = (string)(v19.(string))
		} else {
			me.Confirm = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItem) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 8)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *MenuItemArgPrompt) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["prompt"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Prompt = (string)(v1.(string))
		} else {
			me.Prompt = ""
		}
	}
	v4, o5 := v["placeHolder"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Placeholder = (string)(v4.(string))
		} else {
			me.Placeholder = ""
		}
	}
	v7, o8 := v["value"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Value = (string)(v7.(string))
		} else {
			me.Value = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItemArgPrompt) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 3)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcAnnotaction) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["Range"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				var z3 SrcRange
				me.Range = z3
			} else {
				if false {
				}
				me.Range.jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcRange
			me.Range = z3
		}
	}
	v5, o6 := v["Title"]
	if o6 {
		println(v5)
		if nil != v5 {
			me.Title = (string)(v5.(string))
		} else {
			me.Title = ""
		}
	}
	v8, o9 := v["Desc"]
	if o9 {
		println(v8)
		if nil != v8 {
			me.Desc = (string)(v8.(string))
		} else {
			me.Desc = ""
		}
	}
	v11, o12 := v["CmdName"]
	if o12 {
		println(v11)
		if nil != v11 {
			me.CmdName = (string)(v11.(string))
		} else {
			me.CmdName = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcAnnotaction) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 4)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
		{
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if len(me.Language) != 0 {
			r = append(r, ",\"language\":"...)
			r = append(r, pkg__strconv.Quote(me.Language)...)
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcInfoTip) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["value"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Value = (string)(v1.(string))
		} else {
			me.Value = ""
		}
	}
	v4, o5 := v["language"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Language = (string)(v4.(string))
		} else {
			me.Language = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcInfoTip) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntelCompl) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["kind"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Kind = (Completion)(v1.(float64))
		} else {
			me.Kind = 0
		}
	}
	v4, o5 := v["label"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Label = (string)(v4.(string))
		} else {
			me.Label = ""
		}
	}
	v7, o8 := v["documentation"]
	if o8 {
		println(v7)
		if nil != v7 {
			v10 := v7.(map[string]interface{})
			if v10 == nil {
				me.Documentation = nil
			} else {
				if nil == me.Documentation {
					me.Documentation = new(SrcIntelDoc)
				}
				me.Documentation.jsonUnmarshal_FromAny(v10)
			}
		} else {
			me.Documentation = nil
		}
	}
	v11, o12 := v["detail"]
	if o12 {
		println(v11)
		if nil != v11 {
			me.Detail = (string)(v11.(string))
		} else {
			me.Detail = ""
		}
	}
	v14, o15 := v["sortText"]
	if o15 {
		println(v14)
		if nil != v14 {
			me.SortText = (string)(v14.(string))
		} else {
			me.SortText = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompl) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 6)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		ai2 := len(r)
		for i1 := range me {
			if nil != me[i1] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i1].preview_MarshalJSON()
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
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me SrcIntelCompls) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me SrcIntelCompls) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
		if len(me.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai3 := len(r)
			for i2 := range me.InfoTips {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.InfoTips[i2].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai3] == 44 {
				r = append(r[:ai3], r[(ai3+1):]...)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntels) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["InfoTips"]
	if o2 {
		println(v1)
		if nil != v1 {
			s4 := v1.([]interface{})
			if s4 == nil {
				me.InfoTips = nil
			} else {
				if len(me.InfoTips) >= len(s4) {
					me.InfoTips = me.InfoTips[0:len(s4)]
				} else {
					me.InfoTips = make([]SrcInfoTip, len(s4))
				}
				for si5, sv6 := range s4 {
					println(sv6)
					if nil != sv6 {
						v8 := sv6.(map[string]interface{})
						if v8 == nil {
							var z7 SrcInfoTip
							me.InfoTips[si5] = z7
						} else {
							if false {
							}
							me.InfoTips[si5].jsonUnmarshal_FromAny(v8)
						}
					} else {
						var z7 SrcInfoTip
						me.InfoTips[si5] = z7
					}
				}
			}
		} else {
			me.InfoTips = nil
		}
	}
	v9, o10 := v["Refs"]
	if o10 {
		println(v9)
		if nil != v9 {
			v12 := v9.([]interface{})
			if v12 == nil {
				me.Refs = nil
			} else {
				if false {
				}
				me.Refs.jsonUnmarshal_FromAny(v12)
			}
		} else {
			me.Refs = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntels) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
		if len(me.Value) != 0 {
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if me.IsTrusted {
			r = append(r, ",\"isTrusted\":"...)
			r = append(r, pkg__strconv.FormatBool(me.IsTrusted)...)
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntelDoc) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["value"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Value = (string)(v1.(string))
		} else {
			me.Value = ""
		}
	}
	v4, o5 := v["isTrusted"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.IsTrusted = (bool)(v4.(bool))
		} else {
			me.IsTrusted = false
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelDoc) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
			ai3 := len(r)
			for i2 := range me.Signatures {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Signatures[i2].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai3] == 44 {
				r = append(r[:ai3], r[(ai3+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntelSigHelp) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["activeSignature"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.ActiveSignature = (int)(v1.(float64))
		} else {
			me.ActiveSignature = 0
		}
	}
	v4, o5 := v["activeParameter"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.ActiveParameter = (int)(v4.(float64))
		} else {
			me.ActiveParameter = 0
		}
	}
	v7, o8 := v["signatures"]
	if o8 {
		println(v7)
		if nil != v7 {
			s10 := v7.([]interface{})
			if s10 == nil {
				me.Signatures = nil
			} else {
				if len(me.Signatures) >= len(s10) {
					me.Signatures = me.Signatures[0:len(s10)]
				} else {
					me.Signatures = make([]SrcIntelSigInfo, len(s10))
				}
				for si11, sv12 := range s10 {
					println(sv12)
					if nil != sv12 {
						v14 := sv12.(map[string]interface{})
						if v14 == nil {
							var z13 SrcIntelSigInfo
							me.Signatures[si11] = z13
						} else {
							if false {
							}
							me.Signatures[si11].jsonUnmarshal_FromAny(v14)
						}
					} else {
						var z13 SrcIntelSigInfo
						me.Signatures[si11] = z13
					}
				}
			}
		} else {
			me.Signatures = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigHelp) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 3)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
			ai9 := len(r)
			for i8 := range me.Parameters {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Parameters[i8].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai9] == 44 {
				r = append(r[:ai9], r[(ai9+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntelSigInfo) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["label"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Label = (string)(v1.(string))
		} else {
			me.Label = ""
		}
	}
	v4, o5 := v["documentation"]
	if o5 {
		println(v4)
		if nil != v4 {
			v7 := v4.(map[string]interface{})
			if v7 == nil {
				var z6 SrcIntelDoc
				me.Documentation = z6
			} else {
				if false {
				}
				me.Documentation.jsonUnmarshal_FromAny(v7)
			}
		} else {
			var z6 SrcIntelDoc
			me.Documentation = z6
		}
	}
	v8, o9 := v["parameters"]
	if o9 {
		println(v8)
		if nil != v8 {
			s11 := v8.([]interface{})
			if s11 == nil {
				me.Parameters = nil
			} else {
				if len(me.Parameters) >= len(s11) {
					me.Parameters = me.Parameters[0:len(s11)]
				} else {
					me.Parameters = make([]SrcIntelSigParam, len(s11))
				}
				for si12, sv13 := range s11 {
					println(sv13)
					if nil != sv13 {
						v15 := sv13.(map[string]interface{})
						if v15 == nil {
							var z14 SrcIntelSigParam
							me.Parameters[si12] = z14
						} else {
							if false {
							}
							me.Parameters[si12].jsonUnmarshal_FromAny(v15)
						}
					} else {
						var z14 SrcIntelSigParam
						me.Parameters[si12] = z14
					}
				}
			}
		} else {
			me.Parameters = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigInfo) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 3)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcIntelSigParam) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["label"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Label = (string)(v1.(string))
		} else {
			me.Label = ""
		}
	}
	v4, o5 := v["documentation"]
	if o5 {
		println(v4)
		if nil != v4 {
			v7 := v4.(map[string]interface{})
			if v7 == nil {
				var z6 SrcIntelDoc
				me.Documentation = z6
			} else {
				if false {
				}
				me.Documentation.jsonUnmarshal_FromAny(v7)
			}
		} else {
			var z6 SrcIntelDoc
			me.Documentation = z6
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigParam) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		ai2 := len(r)
		for i1 := range me {
			if nil != me[i1] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i1].preview_MarshalJSON()
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
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me SrcLenses) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me SrcLenses) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcLens) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcLoc"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				var z3 SrcLoc
				me.SrcLoc = z3
			} else {
				if false {
				}
				me.SrcLoc.jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcLoc
			me.SrcLoc = z3
		}
	}
	v5, o6 := v["t"]
	if o6 {
		println(v5)
		if nil != v5 {
			me.Txt = (string)(v5.(string))
		} else {
			me.Txt = ""
		}
	}
	v8, o9 := v["s"]
	if o9 {
		println(v8)
		if nil != v8 {
			me.Str = (string)(v8.(string))
		} else {
			me.Str = ""
		}
	}
	v11, o12 := v["l"]
	if o12 {
		println(v11)
		if nil != v11 {
			me.CrLf = (bool)(v11.(bool))
		} else {
			me.CrLf = false
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLens) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 4)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcLoc) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["e"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Flag = (int)(v1.(float64))
		} else {
			me.Flag = 0
		}
	}
	v4, o5 := v["f"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.FilePath = (string)(v4.(string))
		} else {
			me.FilePath = ""
		}
	}
	v7, o8 := v["p"]
	if o8 {
		println(v7)
		if nil != v7 {
			v10 := v7.(map[string]interface{})
			if v10 == nil {
				me.Pos = nil
			} else {
				if nil == me.Pos {
					me.Pos = new(SrcPos)
				}
				me.Pos.jsonUnmarshal_FromAny(v10)
			}
		} else {
			me.Pos = nil
		}
	}
	v11, o12 := v["r"]
	if o12 {
		println(v11)
		if nil != v11 {
			v14 := v11.(map[string]interface{})
			if v14 == nil {
				me.Range = nil
			} else {
				if nil == me.Range {
					me.Range = new(SrcRange)
				}
				me.Range.jsonUnmarshal_FromAny(v14)
			}
		} else {
			me.Range = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLoc) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 4)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcLocs) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if nil == me {
		r = append(r, "null"...)
	} else {
		r = append(r, 91)
		ai2 := len(r)
		for i1 := range me {
			if nil != me[i1] {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me[i1].preview_MarshalJSON()
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
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me SrcLocs) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me SrcLocs) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcModEdit) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["At"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				me.At = nil
			} else {
				if nil == me.At {
					me.At = new(SrcRange)
				}
				me.At.jsonUnmarshal_FromAny(v4)
			}
		} else {
			me.At = nil
		}
	}
	v5, o6 := v["Val"]
	if o6 {
		println(v5)
		if nil != v5 {
			me.Val = (string)(v5.(string))
		} else {
			me.Val = ""
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdit) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
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
		ai2 := len(r)
		for i1 := range me {
			{
				r = append(r, 44)
				var e error
				var sl []byte
				sl, e = me[i1].preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		r = append(r, 93)
		if r[ai2] == 44 {
			r = append(r[:ai2], r[(ai2+1):]...)
		}
	}
	return
}

func (me SrcModEdits) jsonUnmarshal_FromAny(v []interface{}) {}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me SrcModEdits) preview_UnmarshalJSON(v []byte) (err error) {
	{
		err = nil
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
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcPos) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["l"]
	if o2 {
		println(v1)
		if nil != v1 {
			me.Ln = (int)(v1.(float64))
		} else {
			me.Ln = 0
		}
	}
	v4, o5 := v["c"]
	if o5 {
		println(v4)
		if nil != v4 {
			me.Col = (int)(v4.(float64))
		} else {
			me.Col = 0
		}
	}
	v7, o8 := v["o"]
	if o8 {
		println(v7)
		if nil != v7 {
			me.Off = (int)(v7.(float64))
		} else {
			me.Off = 0
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcPos) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 5)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcRange) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si1 := len(r)
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
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

func (me *SrcRange) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["s"]
	if o2 {
		println(v1)
		if nil != v1 {
			v4 := v1.(map[string]interface{})
			if v4 == nil {
				var z3 SrcPos
				me.Start = z3
			} else {
				if false {
				}
				me.Start.jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcPos
			me.Start = z3
		}
	}
	v5, o6 := v["e"]
	if o6 {
		println(v5)
		if nil != v5 {
			v8 := v5.(map[string]interface{})
			if v8 == nil {
				var z7 SrcPos
				me.End = z7
			} else {
				if false {
				}
				me.End.jsonUnmarshal_FromAny(v8)
			}
		} else {
			var z7 SrcPos
			me.End = z7
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) preview_MarshalJSON() (r []byte, err error) {
	panic("WorkspaceChanges")
	return
}

func (me *WorkspaceChanges) jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["AddedDirs"]
	if o2 {
		println(v1)
		if nil != v1 {
			s4 := v1.([]interface{})
			if s4 == nil {
				me.AddedDirs = nil
			} else {
				if len(me.AddedDirs) >= len(s4) {
					me.AddedDirs = me.AddedDirs[0:len(s4)]
				} else {
					me.AddedDirs = make([]string, len(s4))
				}
				for si5, sv6 := range s4 {
					println(sv6)
					if nil != sv6 {
						me.AddedDirs[si5] = (string)(sv6.(string))
					} else {
						me.AddedDirs[si5] = ""
					}
				}
			}
		} else {
			me.AddedDirs = nil
		}
	}
	v8, o9 := v["RemovedDirs"]
	if o9 {
		println(v8)
		if nil != v8 {
			s11 := v8.([]interface{})
			if s11 == nil {
				me.RemovedDirs = nil
			} else {
				if len(me.RemovedDirs) >= len(s11) {
					me.RemovedDirs = me.RemovedDirs[0:len(s11)]
				} else {
					me.RemovedDirs = make([]string, len(s11))
				}
				for si12, sv13 := range s11 {
					println(sv13)
					if nil != sv13 {
						me.RemovedDirs[si12] = (string)(sv13.(string))
					} else {
						me.RemovedDirs[si12] = ""
					}
				}
			}
		} else {
			me.RemovedDirs = nil
		}
	}
	v15, o16 := v["OpenedFiles"]
	if o16 {
		println(v15)
		if nil != v15 {
			s18 := v15.([]interface{})
			if s18 == nil {
				me.OpenedFiles = nil
			} else {
				if len(me.OpenedFiles) >= len(s18) {
					me.OpenedFiles = me.OpenedFiles[0:len(s18)]
				} else {
					me.OpenedFiles = make([]string, len(s18))
				}
				for si19, sv20 := range s18 {
					println(sv20)
					if nil != sv20 {
						me.OpenedFiles[si19] = (string)(sv20.(string))
					} else {
						me.OpenedFiles[si19] = ""
					}
				}
			}
		} else {
			me.OpenedFiles = nil
		}
	}
	v22, o23 := v["ClosedFiles"]
	if o23 {
		println(v22)
		if nil != v22 {
			s25 := v22.([]interface{})
			if s25 == nil {
				me.ClosedFiles = nil
			} else {
				if len(me.ClosedFiles) >= len(s25) {
					me.ClosedFiles = me.ClosedFiles[0:len(s25)]
				} else {
					me.ClosedFiles = make([]string, len(s25))
				}
				for si26, sv27 := range s25 {
					println(sv27)
					if nil != sv27 {
						me.ClosedFiles[si26] = (string)(sv27.(string))
					} else {
						me.ClosedFiles[si26] = ""
					}
				}
			}
		} else {
			me.ClosedFiles = nil
		}
	}
	v29, o30 := v["WrittenFiles"]
	if o30 {
		println(v29)
		if nil != v29 {
			s32 := v29.([]interface{})
			if s32 == nil {
				me.WrittenFiles = nil
			} else {
				if len(me.WrittenFiles) >= len(s32) {
					me.WrittenFiles = me.WrittenFiles[0:len(s32)]
				} else {
					me.WrittenFiles = make([]string, len(s32))
				}
				for si33, sv34 := range s32 {
					println(sv34)
					if nil != sv34 {
						me.WrittenFiles[si33] = (string)(sv34.(string))
					} else {
						me.WrittenFiles[si33] = ""
					}
				}
			}
		} else {
			me.WrittenFiles = nil
		}
	}
	v36, o37 := v["LiveFiles"]
	if o37 {
		println(v36)
		if nil != v36 {
		} else {
			me.LiveFiles = nil
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *WorkspaceChanges) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 6)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.jsonUnmarshal_FromAny(kvs)
	}
	return
}

func __gent__jsonMarshal_interface____(v interface{}) (r []byte, err error) {
	if nil == v {
		r = append(r, "null"...)
	} else {
		{
			var e error
			var sl []byte
			j, ok := v.(pkg__encoding_json.Marshaler)
			if ok {
				if j != nil {
					sl, e = j.MarshalJSON()
				} else {
					sl = ([]byte)("null")
				}
			} else {
				v6, ok7 := v.(map[string]interface{})
				if ok7 {
					sl, e = __gent__jsonMarshal_mapsstring_interface____(v6)
				} else {
					v4, ok5 := v.([]string)
					if ok5 {
						sl, e = __gent__jsonMarshal_s_string(v4)
					} else {
						v2, ok3 := v.(string)
						if ok3 {
							sl, e = __gent__jsonMarshal_string(v2)
						} else {
							sl, e = pkg__encoding_json.Marshal(v)
							println("JSON.MARSHAL:", pkg__fmt.Sprintf("%T", v))
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
	return
}

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
		ai9 := len(r)
		for i8 := range v {
			{
				r = append(r, 44)
				r = append(r, pkg__strconv.Quote(v[i8])...)
			}
		}
		r = append(r, 93)
		if r[ai9] == 44 {
			r = append(r[:ai9], r[(ai9+1):]...)
		}
	}
	return
}

func __gent__jsonMarshal_mapsstring_interface____(v map[string]interface{}) (r []byte, err error) {
	if nil == v {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		mi10 := len(r)
		for mk11, mv12 := range v {
			sl, e := __gent__jsonMarshal_interface____(mv12)
			if e == nil {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(mk11)...)
					r = append(r, 58)
				}
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[mi10] == 44 {
			r = append(r[:mi10], r[(mi10+1):]...)
		}
	}
	return
}
