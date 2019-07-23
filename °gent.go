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

func (me *IpcReq) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ri"]
	if !o2 {
		me.ReqID = 0
	} else {
		println(v1)
		if nil != v1 {
			me.ReqID = (int64)(v1.(float64))
		} else {
			me.ReqID = 0
		}
	}
	v5, o6 := v["ii"]
	if !o6 {
		me.IpcID = 0
	} else {
		println(v5)
		if nil != v5 {
			me.IpcID = (IpcIDs)(v5.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v9, o10 := v["ia"]
	if !o10 {
		me.IpcArgs = nil
	} else {
		println(v9)
		if nil != v9 {
		} else {
			me.IpcArgs = nil
		}
	}
	v13, o14 := v["projUpd"]
	if !o14 {
		me.ProjUpd = nil
	} else {
		println(v13)
		if nil != v13 {
			v17 := v13.(map[string]interface{})
			if v17 == nil {
				me.ProjUpd = nil
			} else {
				if nil == me.ProjUpd {
					me.ProjUpd = new(WorkspaceChanges)
				}
				me.ProjUpd.__gent__jsonUnmarshal_FromAny(v17)
			}
		} else {
			me.ProjUpd = nil
		}
	}
	v18, o19 := v["srcLens"]
	if !o19 {
		me.SrcLens = nil
	} else {
		println(v18)
		if nil != v18 {
			v22 := v18.(map[string]interface{})
			if v22 == nil {
				me.SrcLens = nil
			} else {
				if nil == me.SrcLens {
					me.SrcLens = new(SrcLens)
				}
				me.SrcLens.__gent__jsonUnmarshal_FromAny(v22)
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *IpcResp) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ii"]
	if !o2 {
		me.IpcID = 0
	} else {
		println(v1)
		if nil != v1 {
			me.IpcID = (IpcIDs)(v1.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v5, o6 := v["ri"]
	if !o6 {
		me.ReqID = 0
	} else {
		println(v5)
		if nil != v5 {
			me.ReqID = (int64)(v5.(float64))
		} else {
			me.ReqID = 0
		}
	}
	v9, o10 := v["err"]
	if !o10 {
		me.ErrMsg = ""
	} else {
		println(v9)
		if nil != v9 {
			me.ErrMsg = (string)(v9.(string))
		} else {
			me.ErrMsg = ""
		}
	}
	v13, o14 := v["sI"]
	if !o14 {
		me.SrcIntel = nil
	} else {
		println(v13)
		if nil != v13 {
			v17 := v13.(map[string]interface{})
			if v17 == nil {
				me.SrcIntel = nil
			} else {
				if nil == me.SrcIntel {
					me.SrcIntel = new(SrcIntel)
				}
				me.SrcIntel.__gent__jsonUnmarshal_FromAny(v17)
			}
		} else {
			me.SrcIntel = nil
		}
	}
	v18, o19 := v["srcDiags"]
	if !o19 {
		me.SrcDiags = nil
	} else {
		println(v18)
		if nil != v18 {
			v22 := v18.(map[string]interface{})
			if v22 == nil {
				me.SrcDiags = nil
			} else {
				if nil == me.SrcDiags {
					me.SrcDiags = new(Diags)
				}
				me.SrcDiags.__gent__jsonUnmarshal_FromAny(v22)
			}
		} else {
			me.SrcDiags = nil
		}
	}
	v23, o24 := v["srcMods"]
	if !o24 {
		me.SrcMods = nil
	} else {
		println(v23)
		if nil != v23 {
			v27 := v23.([]interface{})
			if v27 == nil {
				me.SrcMods = nil
			} else {
				if false {
				}
				me.SrcMods.__gent__jsonUnmarshal_FromAny(v27)
			}
		} else {
			me.SrcMods = nil
		}
	}
	v28, o29 := v["srcActions"]
	if !o29 {
		me.SrcActions = nil
	} else {
		println(v28)
		if nil != v28 {
			s32 := v28.([]interface{})
			if s32 == nil {
				me.SrcActions = nil
			} else {
				if len(me.SrcActions) >= len(s32) {
					me.SrcActions = me.SrcActions[0:len(s32)]
				} else {
					me.SrcActions = make([]EditorAction, len(s32))
				}
				for si33, sv34 := range s32 {
					println(sv34)
					if nil != sv34 {
						v36 := sv34.(map[string]interface{})
						if v36 == nil {
							var z35 EditorAction
							me.SrcActions[si33] = z35
						} else {
							if false {
							}
							me.SrcActions[si33].__gent__jsonUnmarshal_FromAny(v36)
						}
					} else {
						var z35 EditorAction
						me.SrcActions[si33] = z35
					}
				}
			}
		} else {
			me.SrcActions = nil
		}
	}
	v37, o38 := v["extras"]
	if !o38 {
		me.Extras = nil
	} else {
		println(v37)
		if nil != v37 {
			v41 := v37.(map[string]interface{})
			if v41 == nil {
				me.Extras = nil
			} else {
				if nil == me.Extras {
					me.Extras = new(Extras)
				}
				me.Extras.__gent__jsonUnmarshal_FromAny(v41)
			}
		} else {
			me.Extras = nil
		}
	}
	v42, o43 := v["menu"]
	if !o43 {
		me.Menu = nil
	} else {
		println(v42)
		if nil != v42 {
			v46 := v42.(map[string]interface{})
			if v46 == nil {
				me.Menu = nil
			} else {
				if nil == me.Menu {
					me.Menu = new(MenuResponse)
				}
				me.Menu.__gent__jsonUnmarshal_FromAny(v46)
			}
		} else {
			me.Menu = nil
		}
	}
	v47, o48 := v["caddy"]
	if !o48 {
		me.CaddyUpdate = nil
	} else {
		println(v47)
		if nil != v47 {
			v51 := v47.(map[string]interface{})
			if v51 == nil {
				me.CaddyUpdate = nil
			} else {
				if nil == me.CaddyUpdate {
					me.CaddyUpdate = new(Caddy)
				}
				me.CaddyUpdate.__gent__jsonUnmarshal_FromAny(v51)
			}
		} else {
			me.CaddyUpdate = nil
		}
	}
	v52, o53 := v["val"]
	if !o53 {
		me.Val = nil
	} else {
		println(v52)
		if nil != v52 {
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *Diags) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["All"]
	if !o2 {
		me.All = nil
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				me.All = nil
			} else {
				if nil == me.All {
					me.All = make(DiagItemsBy, len(v5))
				}
				me.All.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			me.All = nil
		}
	}
	v6, o7 := v["FixUps"]
	if !o7 {
		me.FixUps = nil
	} else {
		println(v6)
		if nil != v6 {
			s10 := v6.([]interface{})
			if s10 == nil {
				me.FixUps = nil
			} else {
				if len(me.FixUps) >= len(s10) {
					me.FixUps = me.FixUps[0:len(s10)]
				} else {
					me.FixUps = make([]*DiagFixUps, len(s10))
				}
				for si11, sv12 := range s10 {
					println(sv12)
					if nil != sv12 {
						v14 := sv12.(map[string]interface{})
						if v14 == nil {
							me.FixUps[si11] = nil
						} else {
							if nil == me.FixUps[si11] {
								me.FixUps[si11] = new(DiagFixUps)
							}
							me.FixUps[si11].__gent__jsonUnmarshal_FromAny(v14)
						}
					} else {
						me.FixUps[si11] = nil
					}
				}
			}
		} else {
			me.FixUps = nil
		}
	}
	v15, o16 := v["LangID"]
	if !o16 {
		me.LangID = ""
	} else {
		println(v15)
		if nil != v15 {
			me.LangID = (string)(v15.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *Extras) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcIntels"]
	if !o2 {
		var z3 SrcIntels
		me.SrcIntels = z3
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				var z4 SrcIntels
				me.SrcIntels = z4
			} else {
				if false {
				}
				me.SrcIntels.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			var z4 SrcIntels
			me.SrcIntels = z4
		}
	}
	v6, o7 := v["Items"]
	if !o7 {
		me.Items = nil
	} else {
		println(v6)
		if nil != v6 {
			s10 := v6.([]interface{})
			if s10 == nil {
				me.Items = nil
			} else {
				if len(me.Items) >= len(s10) {
					me.Items = me.Items[0:len(s10)]
				} else {
					me.Items = make([]*ExtrasItem, len(s10))
				}
				for si11, sv12 := range s10 {
					println(sv12)
					if nil != sv12 {
						v14 := sv12.(map[string]interface{})
						if v14 == nil {
							me.Items[si11] = nil
						} else {
							if nil == me.Items[si11] {
								me.Items[si11] = new(ExtrasItem)
							}
							me.Items[si11].__gent__jsonUnmarshal_FromAny(v14)
						}
					} else {
						me.Items[si11] = nil
					}
				}
			}
		} else {
			me.Items = nil
		}
	}
	v15, o16 := v["Warns"]
	if !o16 {
		me.Warns = nil
	} else {
		println(v15)
		if nil != v15 {
			s19 := v15.([]interface{})
			if s19 == nil {
				me.Warns = nil
			} else {
				if len(me.Warns) >= len(s19) {
					me.Warns = me.Warns[0:len(s19)]
				} else {
					me.Warns = make([]string, len(s19))
				}
				for si20, sv21 := range s19 {
					println(sv21)
					if nil != sv21 {
						me.Warns[si20] = (string)(sv21.(string))
					} else {
						me.Warns[si20] = ""
					}
				}
			}
		} else {
			me.Warns = nil
		}
	}
	v23, o24 := v["Desc"]
	if !o24 {
		me.Desc = ""
	} else {
		println(v23)
		if nil != v23 {
			me.Desc = (string)(v23.(string))
		} else {
			me.Desc = ""
		}
	}
	v27, o28 := v["Url"]
	if !o28 {
		me.Url = ""
	} else {
		println(v27)
		if nil != v27 {
			me.Url = (string)(v27.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *MenuResponse) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SubMenu"]
	if !o2 {
		me.SubMenu = nil
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				me.SubMenu = nil
			} else {
				if nil == me.SubMenu {
					me.SubMenu = new(Menu)
				}
				me.SubMenu.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			me.SubMenu = nil
		}
	}
	v6, o7 := v["WebsiteURL"]
	if !o7 {
		me.WebsiteURL = ""
	} else {
		println(v6)
		if nil != v6 {
			me.WebsiteURL = (string)(v6.(string))
		} else {
			me.WebsiteURL = ""
		}
	}
	v10, o11 := v["NoteInfo"]
	if !o11 {
		me.NoteInfo = ""
	} else {
		println(v10)
		if nil != v10 {
			me.NoteInfo = (string)(v10.(string))
		} else {
			me.NoteInfo = ""
		}
	}
	v14, o15 := v["NoteWarn"]
	if !o15 {
		me.NoteWarn = ""
	} else {
		println(v14)
		if nil != v14 {
			me.NoteWarn = (string)(v14.(string))
		} else {
			me.NoteWarn = ""
		}
	}
	v18, o19 := v["UxActionLabel"]
	if !o19 {
		me.UxActionLabel = ""
	} else {
		println(v18)
		if nil != v18 {
			me.UxActionLabel = (string)(v18.(string))
		} else {
			me.UxActionLabel = ""
		}
	}
	v22, o23 := v["Refs"]
	if !o23 {
		me.Refs = nil
	} else {
		println(v22)
		if nil != v22 {
			v26 := v22.([]interface{})
			if v26 == nil {
				me.Refs = nil
			} else {
				if false {
				}
				me.Refs.__gent__jsonUnmarshal_FromAny(v26)
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntel) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcIntels"]
	if !o2 {
		var z3 SrcIntels
		me.SrcIntels = z3
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				var z4 SrcIntels
				me.SrcIntels = z4
			} else {
				if false {
				}
				me.SrcIntels.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			var z4 SrcIntels
			me.SrcIntels = z4
		}
	}
	v6, o7 := v["Sig"]
	if !o7 {
		me.Sig = nil
	} else {
		println(v6)
		if nil != v6 {
			v10 := v6.(map[string]interface{})
			if v10 == nil {
				me.Sig = nil
			} else {
				if nil == me.Sig {
					me.Sig = new(SrcIntelSigHelp)
				}
				me.Sig.__gent__jsonUnmarshal_FromAny(v10)
			}
		} else {
			me.Sig = nil
		}
	}
	v11, o12 := v["Cmpl"]
	if !o12 {
		me.Cmpl = nil
	} else {
		println(v11)
		if nil != v11 {
			v15 := v11.([]interface{})
			if v15 == nil {
				me.Cmpl = nil
			} else {
				if false {
				}
				me.Cmpl.__gent__jsonUnmarshal_FromAny(v15)
			}
		} else {
			me.Cmpl = nil
		}
	}
	v16, o17 := v["Syms"]
	if !o17 {
		me.Syms = nil
	} else {
		println(v16)
		if nil != v16 {
			v20 := v16.([]interface{})
			if v20 == nil {
				me.Syms = nil
			} else {
				if false {
				}
				me.Syms.__gent__jsonUnmarshal_FromAny(v20)
			}
		} else {
			me.Syms = nil
		}
	}
	v21, o22 := v["Anns"]
	if !o22 {
		me.Anns = nil
	} else {
		println(v21)
		if nil != v21 {
			s25 := v21.([]interface{})
			if s25 == nil {
				me.Anns = nil
			} else {
				if len(me.Anns) >= len(s25) {
					me.Anns = me.Anns[0:len(s25)]
				} else {
					me.Anns = make([]*SrcAnnotaction, len(s25))
				}
				for si26, sv27 := range s25 {
					println(sv27)
					if nil != sv27 {
						v29 := sv27.(map[string]interface{})
						if v29 == nil {
							me.Anns[si26] = nil
						} else {
							if nil == me.Anns[si26] {
								me.Anns[si26] = new(SrcAnnotaction)
							}
							me.Anns[si26].__gent__jsonUnmarshal_FromAny(v29)
						}
					} else {
						me.Anns[si26] = nil
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *Caddy) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ID"]
	if !o2 {
		me.ID = ""
	} else {
		println(v1)
		if nil != v1 {
			me.ID = (string)(v1.(string))
		} else {
			me.ID = ""
		}
	}
	v5, o6 := v["LangID"]
	if !o6 {
		me.LangID = ""
	} else {
		println(v5)
		if nil != v5 {
			me.LangID = (string)(v5.(string))
		} else {
			me.LangID = ""
		}
	}
	v9, o10 := v["Icon"]
	if !o10 {
		me.Icon = ""
	} else {
		println(v9)
		if nil != v9 {
			me.Icon = (string)(v9.(string))
		} else {
			me.Icon = ""
		}
	}
	v13, o14 := v["Title"]
	if !o14 {
		me.Title = ""
	} else {
		println(v13)
		if nil != v13 {
			me.Title = (string)(v13.(string))
		} else {
			me.Title = ""
		}
	}
	v17, o18 := v["Status"]
	if !o18 {
		var z19 struct {
			Flag CaddyStatus
			Desc string `json:",omitempty"`
		}
		me.Status = z19
	} else {
		println(v17)
		if nil != v17 {
			t21 := v17.(map[string]interface{})
			if nil == t21 {
				var z20 struct {
					Flag CaddyStatus
					Desc string `json:",omitempty"`
				}
				me.Status = z20
			} else {
				v22, o23 := t21["Flag"]
				if !o23 {
					me.Status.Flag = 0
				} else {
					println(v22)
					if nil != v22 {
						me.Status.Flag = (CaddyStatus)(v22.(float64))
					} else {
						me.Status.Flag = 0
					}
				}
				v26, o27 := t21["Desc"]
				if !o27 {
					me.Status.Desc = ""
				} else {
					println(v26)
					if nil != v26 {
						me.Status.Desc = (string)(v26.(string))
					} else {
						me.Status.Desc = ""
					}
				}
			}
		} else {
			var z20 struct {
				Flag CaddyStatus
				Desc string `json:",omitempty"`
			}
			me.Status = z20
		}
	}
	v30, o31 := v["Details"]
	if !o31 {
		me.Details = ""
	} else {
		println(v30)
		if nil != v30 {
			me.Details = (string)(v30.(string))
		} else {
			me.Details = ""
		}
	}
	v34, o35 := v["UxActionID"]
	if !o35 {
		me.UxActionID = ""
	} else {
		println(v34)
		if nil != v34 {
			me.UxActionID = (string)(v34.(string))
		} else {
			me.UxActionID = ""
		}
	}
	v38, o39 := v["ShowTitle"]
	if !o39 {
		me.ShowTitle = false
	} else {
		println(v38)
		if nil != v38 {
			me.ShowTitle = (bool)(v38.(bool))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *DiagFixUps) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["FilePath"]
	if !o2 {
		me.FilePath = ""
	} else {
		println(v1)
		if nil != v1 {
			me.FilePath = (string)(v1.(string))
		} else {
			me.FilePath = ""
		}
	}
	v5, o6 := v["Desc"]
	if !o6 {
		me.Desc = nil
	} else {
		println(v5)
		if nil != v5 {
			m9 := v5.(map[string]interface{})
			if m9 == nil {
				me.Desc = nil
			} else {
				me.Desc = make(map[string][]string, len(m9))
				for mk10, mv11 := range m9 {
					var t12 []string
					println(mv11)
					if nil != mv11 {
						s14 := mv11.([]interface{})
						if s14 == nil {
							t12 = nil
						} else {
							if len(t12) >= len(s14) {
								t12 = t12[0:len(s14)]
							} else {
								t12 = make([]string, len(s14))
							}
							for si15, sv16 := range s14 {
								println(sv16)
								if nil != sv16 {
									t12[si15] = (string)(sv16.(string))
								} else {
									t12[si15] = ""
								}
							}
						}
					} else {
						t12 = nil
					}
					me.Desc[mk10] = t12
				}
			}
		} else {
			me.Desc = nil
		}
	}
	v18, o19 := v["Edits"]
	if !o19 {
		me.Edits = nil
	} else {
		println(v18)
		if nil != v18 {
			v22 := v18.([]interface{})
			if v22 == nil {
				me.Edits = nil
			} else {
				if false {
				}
				me.Edits.__gent__jsonUnmarshal_FromAny(v22)
			}
		} else {
			me.Edits = nil
		}
	}
	v23, o24 := v["Dropped"]
	if !o24 {
		me.Dropped = nil
	} else {
		println(v23)
		if nil != v23 {
			s27 := v23.([]interface{})
			if s27 == nil {
				me.Dropped = nil
			} else {
				if len(me.Dropped) >= len(s27) {
					me.Dropped = me.Dropped[0:len(s27)]
				} else {
					me.Dropped = make([]SrcModEdit, len(s27))
				}
				for si28, sv29 := range s27 {
					println(sv29)
					if nil != sv29 {
						v31 := sv29.(map[string]interface{})
						if v31 == nil {
							var z30 SrcModEdit
							me.Dropped[si28] = z30
						} else {
							if false {
							}
							me.Dropped[si28].__gent__jsonUnmarshal_FromAny(v31)
						}
					} else {
						var z30 SrcModEdit
						me.Dropped[si28] = z30
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *DiagItem) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["Cat"]
	if !o2 {
		me.Cat = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Cat = (string)(v1.(string))
		} else {
			me.Cat = ""
		}
	}
	v5, o6 := v["Loc"]
	if !o6 {
		var z7 SrcLoc
		me.Loc = z7
	} else {
		println(v5)
		if nil != v5 {
			v9 := v5.(map[string]interface{})
			if v9 == nil {
				var z8 SrcLoc
				me.Loc = z8
			} else {
				if false {
				}
				me.Loc.__gent__jsonUnmarshal_FromAny(v9)
			}
		} else {
			var z8 SrcLoc
			me.Loc = z8
		}
	}
	v10, o11 := v["Msg"]
	if !o11 {
		me.Msg = ""
	} else {
		println(v10)
		if nil != v10 {
			me.Msg = (string)(v10.(string))
		} else {
			me.Msg = ""
		}
	}
	v14, o15 := v["Rel"]
	if !o15 {
		me.Rel = nil
	} else {
		println(v14)
		if nil != v14 {
			s18 := v14.([]interface{})
			if s18 == nil {
				me.Rel = nil
			} else {
				if len(me.Rel) >= len(s18) {
					me.Rel = me.Rel[0:len(s18)]
				} else {
					me.Rel = make([]SrcLens, len(s18))
				}
				for si19, sv20 := range s18 {
					println(sv20)
					if nil != sv20 {
						v22 := sv20.(map[string]interface{})
						if v22 == nil {
							var z21 SrcLens
							me.Rel[si19] = z21
						} else {
							if false {
							}
							me.Rel[si19].__gent__jsonUnmarshal_FromAny(v22)
						}
					} else {
						var z21 SrcLens
						me.Rel[si19] = z21
					}
				}
			}
		} else {
			me.Rel = nil
		}
	}
	v23, o24 := v["SrcActions"]
	if !o24 {
		me.SrcActions = nil
	} else {
		println(v23)
		if nil != v23 {
			s27 := v23.([]interface{})
			if s27 == nil {
				me.SrcActions = nil
			} else {
				if len(me.SrcActions) >= len(s27) {
					me.SrcActions = me.SrcActions[0:len(s27)]
				} else {
					me.SrcActions = make([]EditorAction, len(s27))
				}
				for si28, sv29 := range s27 {
					println(sv29)
					if nil != sv29 {
						v31 := sv29.(map[string]interface{})
						if v31 == nil {
							var z30 EditorAction
							me.SrcActions[si28] = z30
						} else {
							if false {
							}
							me.SrcActions[si28].__gent__jsonUnmarshal_FromAny(v31)
						}
					} else {
						var z30 EditorAction
						me.SrcActions[si28] = z30
					}
				}
			}
		} else {
			me.SrcActions = nil
		}
	}
	v32, o33 := v["Sticky"]
	if !o33 {
		me.StickyAuto = false
	} else {
		println(v32)
		if nil != v32 {
			me.StickyAuto = (bool)(v32.(bool))
		} else {
			me.StickyAuto = false
		}
	}
	v36, o37 := v["Tags"]
	if !o37 {
		me.Tags = nil
	} else {
		println(v36)
		if nil != v36 {
			s40 := v36.([]interface{})
			if s40 == nil {
				me.Tags = nil
			} else {
				if len(me.Tags) >= len(s40) {
					me.Tags = me.Tags[0:len(s40)]
				} else {
					me.Tags = make([]int, len(s40))
				}
				for si41, sv42 := range s40 {
					println(sv42)
					if nil != sv42 {
						me.Tags[si41] = (int)(sv42.(float64))
					} else {
						me.Tags[si41] = 0
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *DiagItems) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(DiagItems, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				sl[si1] = nil
			} else {
				if nil == sl[si1] {
					sl[si1] = new(DiagItem)
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			sl[si1] = nil
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItems) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *DiagItemsBy) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	var kvs DiagItemsBy
	kvs = make(DiagItemsBy, len(v))
	for mk1, mv2 := range v {
		var t3 DiagItems
		println(mv2)
		if nil != mv2 {
			v5 := mv2.([]interface{})
			if v5 == nil {
				t3 = nil
			} else {
				if false {
				}
				t3.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			t3 = nil
		}
		kvs[mk1] = t3
	}
	*me = kvs
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItemsBy) preview_UnmarshalJSON(v []byte) (err error) {
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

func (me *EditorAction) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["title"]
	if !o2 {
		me.Title = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Title = (string)(v1.(string))
		} else {
			me.Title = ""
		}
	}
	v5, o6 := v["command"]
	if !o6 {
		me.Cmd = ""
	} else {
		println(v5)
		if nil != v5 {
			me.Cmd = (string)(v5.(string))
		} else {
			me.Cmd = ""
		}
	}
	v9, o10 := v["tooltip"]
	if !o10 {
		me.Hint = ""
	} else {
		println(v9)
		if nil != v9 {
			me.Hint = (string)(v9.(string))
		} else {
			me.Hint = ""
		}
	}
	v13, o14 := v["arguments"]
	if !o14 {
		me.Arguments = nil
	} else {
		println(v13)
		if nil != v13 {
			s17 := v13.([]interface{})
			if s17 == nil {
				me.Arguments = nil
			} else {
				if len(me.Arguments) >= len(s17) {
					me.Arguments = me.Arguments[0:len(s17)]
				} else {
					me.Arguments = make([]interface{}, len(s17))
				}
				for si18, sv19 := range s17 {
					println(sv19)
					if nil != sv19 {
					} else {
						me.Arguments[si18] = nil
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *ExtrasItem) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["id"]
	if !o2 {
		me.ID = ""
	} else {
		println(v1)
		if nil != v1 {
			me.ID = (string)(v1.(string))
		} else {
			me.ID = ""
		}
	}
	v5, o6 := v["label"]
	if !o6 {
		me.Label = ""
	} else {
		println(v5)
		if nil != v5 {
			me.Label = (string)(v5.(string))
		} else {
			me.Label = ""
		}
	}
	v9, o10 := v["description"]
	if !o10 {
		me.Desc = ""
	} else {
		println(v9)
		if nil != v9 {
			me.Desc = (string)(v9.(string))
		} else {
			me.Desc = ""
		}
	}
	v13, o14 := v["detail"]
	if !o14 {
		me.Detail = ""
	} else {
		println(v13)
		if nil != v13 {
			me.Detail = (string)(v13.(string))
		} else {
			me.Detail = ""
		}
	}
	v17, o18 := v["arg"]
	if !o18 {
		me.QueryArg = ""
	} else {
		println(v17)
		if nil != v17 {
			me.QueryArg = (string)(v17.(string))
		} else {
			me.QueryArg = ""
		}
	}
	v21, o22 := v["fPos"]
	if !o22 {
		me.FilePos = ""
	} else {
		println(v21)
		if nil != v21 {
			me.FilePos = (string)(v21.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *Menu) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["desc"]
	if !o2 {
		me.Desc = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Desc = (string)(v1.(string))
		} else {
			me.Desc = ""
		}
	}
	v5, o6 := v["topLevel"]
	if !o6 {
		me.TopLevel = false
	} else {
		println(v5)
		if nil != v5 {
			me.TopLevel = (bool)(v5.(bool))
		} else {
			me.TopLevel = false
		}
	}
	v9, o10 := v["items"]
	if !o10 {
		me.Items = nil
	} else {
		println(v9)
		if nil != v9 {
			v13 := v9.([]interface{})
			if v13 == nil {
				me.Items = nil
			} else {
				if false {
				}
				me.Items.__gent__jsonUnmarshal_FromAny(v13)
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *MenuItems) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(MenuItems, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				sl[si1] = nil
			} else {
				if nil == sl[si1] {
					sl[si1] = new(MenuItem)
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			sl[si1] = nil
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItems) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *MenuItem) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["ii"]
	if !o2 {
		me.IpcID = 0
	} else {
		println(v1)
		if nil != v1 {
			me.IpcID = (IpcIDs)(v1.(float64))
		} else {
			me.IpcID = 0
		}
	}
	v5, o6 := v["ia"]
	if !o6 {
		me.IpcArgs = nil
	} else {
		println(v5)
		if nil != v5 {
		} else {
			me.IpcArgs = nil
		}
	}
	v9, o10 := v["c"]
	if !o10 {
		me.Category = ""
	} else {
		println(v9)
		if nil != v9 {
			me.Category = (string)(v9.(string))
		} else {
			me.Category = ""
		}
	}
	v13, o14 := v["t"]
	if !o14 {
		me.Title = ""
	} else {
		println(v13)
		if nil != v13 {
			me.Title = (string)(v13.(string))
		} else {
			me.Title = ""
		}
	}
	v17, o18 := v["d"]
	if !o18 {
		me.Desc = ""
	} else {
		println(v17)
		if nil != v17 {
			me.Desc = (string)(v17.(string))
		} else {
			me.Desc = ""
		}
	}
	v21, o22 := v["h"]
	if !o22 {
		me.Hint = ""
	} else {
		println(v21)
		if nil != v21 {
			me.Hint = (string)(v21.(string))
		} else {
			me.Hint = ""
		}
	}
	v25, o26 := v["q"]
	if !o26 {
		me.Confirm = ""
	} else {
		println(v25)
		if nil != v25 {
			me.Confirm = (string)(v25.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *MenuItemArgPrompt) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["prompt"]
	if !o2 {
		me.Prompt = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Prompt = (string)(v1.(string))
		} else {
			me.Prompt = ""
		}
	}
	v5, o6 := v["placeHolder"]
	if !o6 {
		me.Placeholder = ""
	} else {
		println(v5)
		if nil != v5 {
			me.Placeholder = (string)(v5.(string))
		} else {
			me.Placeholder = ""
		}
	}
	v9, o10 := v["value"]
	if !o10 {
		me.Value = ""
	} else {
		println(v9)
		if nil != v9 {
			me.Value = (string)(v9.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcAnnotaction) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["Range"]
	if !o2 {
		var z3 SrcRange
		me.Range = z3
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				var z4 SrcRange
				me.Range = z4
			} else {
				if false {
				}
				me.Range.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			var z4 SrcRange
			me.Range = z4
		}
	}
	v6, o7 := v["Title"]
	if !o7 {
		me.Title = ""
	} else {
		println(v6)
		if nil != v6 {
			me.Title = (string)(v6.(string))
		} else {
			me.Title = ""
		}
	}
	v10, o11 := v["Desc"]
	if !o11 {
		me.Desc = ""
	} else {
		println(v10)
		if nil != v10 {
			me.Desc = (string)(v10.(string))
		} else {
			me.Desc = ""
		}
	}
	v14, o15 := v["CmdName"]
	if !o15 {
		me.CmdName = ""
	} else {
		println(v14)
		if nil != v14 {
			me.CmdName = (string)(v14.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcInfoTip) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["value"]
	if !o2 {
		me.Value = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Value = (string)(v1.(string))
		} else {
			me.Value = ""
		}
	}
	v5, o6 := v["language"]
	if !o6 {
		me.Language = ""
	} else {
		println(v5)
		if nil != v5 {
			me.Language = (string)(v5.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelCompl) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["kind"]
	if !o2 {
		me.Kind = 0
	} else {
		println(v1)
		if nil != v1 {
			me.Kind = (Completion)(v1.(float64))
		} else {
			me.Kind = 0
		}
	}
	v5, o6 := v["label"]
	if !o6 {
		me.Label = ""
	} else {
		println(v5)
		if nil != v5 {
			me.Label = (string)(v5.(string))
		} else {
			me.Label = ""
		}
	}
	v9, o10 := v["documentation"]
	if !o10 {
		me.Documentation = nil
	} else {
		println(v9)
		if nil != v9 {
			v13 := v9.(map[string]interface{})
			if v13 == nil {
				me.Documentation = nil
			} else {
				if nil == me.Documentation {
					me.Documentation = new(SrcIntelDoc)
				}
				me.Documentation.__gent__jsonUnmarshal_FromAny(v13)
			}
		} else {
			me.Documentation = nil
		}
	}
	v14, o15 := v["detail"]
	if !o15 {
		me.Detail = ""
	} else {
		println(v14)
		if nil != v14 {
			me.Detail = (string)(v14.(string))
		} else {
			me.Detail = ""
		}
	}
	v18, o19 := v["sortText"]
	if !o19 {
		me.SortText = ""
	} else {
		println(v18)
		if nil != v18 {
			me.SortText = (string)(v18.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelCompls) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(SrcIntelCompls, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				sl[si1] = nil
			} else {
				if nil == sl[si1] {
					sl[si1] = new(SrcIntelCompl)
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			sl[si1] = nil
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompls) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *SrcIntels) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["InfoTips"]
	if !o2 {
		me.InfoTips = nil
	} else {
		println(v1)
		if nil != v1 {
			s5 := v1.([]interface{})
			if s5 == nil {
				me.InfoTips = nil
			} else {
				if len(me.InfoTips) >= len(s5) {
					me.InfoTips = me.InfoTips[0:len(s5)]
				} else {
					me.InfoTips = make([]SrcInfoTip, len(s5))
				}
				for si6, sv7 := range s5 {
					println(sv7)
					if nil != sv7 {
						v9 := sv7.(map[string]interface{})
						if v9 == nil {
							var z8 SrcInfoTip
							me.InfoTips[si6] = z8
						} else {
							if false {
							}
							me.InfoTips[si6].__gent__jsonUnmarshal_FromAny(v9)
						}
					} else {
						var z8 SrcInfoTip
						me.InfoTips[si6] = z8
					}
				}
			}
		} else {
			me.InfoTips = nil
		}
	}
	v10, o11 := v["Refs"]
	if !o11 {
		me.Refs = nil
	} else {
		println(v10)
		if nil != v10 {
			v14 := v10.([]interface{})
			if v14 == nil {
				me.Refs = nil
			} else {
				if false {
				}
				me.Refs.__gent__jsonUnmarshal_FromAny(v14)
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelDoc) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["value"]
	if !o2 {
		me.Value = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Value = (string)(v1.(string))
		} else {
			me.Value = ""
		}
	}
	v5, o6 := v["isTrusted"]
	if !o6 {
		me.IsTrusted = false
	} else {
		println(v5)
		if nil != v5 {
			me.IsTrusted = (bool)(v5.(bool))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelSigHelp) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["activeSignature"]
	if !o2 {
		me.ActiveSignature = 0
	} else {
		println(v1)
		if nil != v1 {
			me.ActiveSignature = (int)(v1.(float64))
		} else {
			me.ActiveSignature = 0
		}
	}
	v5, o6 := v["activeParameter"]
	if !o6 {
		me.ActiveParameter = 0
	} else {
		println(v5)
		if nil != v5 {
			me.ActiveParameter = (int)(v5.(float64))
		} else {
			me.ActiveParameter = 0
		}
	}
	v9, o10 := v["signatures"]
	if !o10 {
		me.Signatures = nil
	} else {
		println(v9)
		if nil != v9 {
			s13 := v9.([]interface{})
			if s13 == nil {
				me.Signatures = nil
			} else {
				if len(me.Signatures) >= len(s13) {
					me.Signatures = me.Signatures[0:len(s13)]
				} else {
					me.Signatures = make([]SrcIntelSigInfo, len(s13))
				}
				for si14, sv15 := range s13 {
					println(sv15)
					if nil != sv15 {
						v17 := sv15.(map[string]interface{})
						if v17 == nil {
							var z16 SrcIntelSigInfo
							me.Signatures[si14] = z16
						} else {
							if false {
							}
							me.Signatures[si14].__gent__jsonUnmarshal_FromAny(v17)
						}
					} else {
						var z16 SrcIntelSigInfo
						me.Signatures[si14] = z16
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelSigInfo) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["label"]
	if !o2 {
		me.Label = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Label = (string)(v1.(string))
		} else {
			me.Label = ""
		}
	}
	v5, o6 := v["documentation"]
	if !o6 {
		var z7 SrcIntelDoc
		me.Documentation = z7
	} else {
		println(v5)
		if nil != v5 {
			v9 := v5.(map[string]interface{})
			if v9 == nil {
				var z8 SrcIntelDoc
				me.Documentation = z8
			} else {
				if false {
				}
				me.Documentation.__gent__jsonUnmarshal_FromAny(v9)
			}
		} else {
			var z8 SrcIntelDoc
			me.Documentation = z8
		}
	}
	v10, o11 := v["parameters"]
	if !o11 {
		me.Parameters = nil
	} else {
		println(v10)
		if nil != v10 {
			s14 := v10.([]interface{})
			if s14 == nil {
				me.Parameters = nil
			} else {
				if len(me.Parameters) >= len(s14) {
					me.Parameters = me.Parameters[0:len(s14)]
				} else {
					me.Parameters = make([]SrcIntelSigParam, len(s14))
				}
				for si15, sv16 := range s14 {
					println(sv16)
					if nil != sv16 {
						v18 := sv16.(map[string]interface{})
						if v18 == nil {
							var z17 SrcIntelSigParam
							me.Parameters[si15] = z17
						} else {
							if false {
							}
							me.Parameters[si15].__gent__jsonUnmarshal_FromAny(v18)
						}
					} else {
						var z17 SrcIntelSigParam
						me.Parameters[si15] = z17
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcIntelSigParam) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["label"]
	if !o2 {
		me.Label = ""
	} else {
		println(v1)
		if nil != v1 {
			me.Label = (string)(v1.(string))
		} else {
			me.Label = ""
		}
	}
	v5, o6 := v["documentation"]
	if !o6 {
		var z7 SrcIntelDoc
		me.Documentation = z7
	} else {
		println(v5)
		if nil != v5 {
			v9 := v5.(map[string]interface{})
			if v9 == nil {
				var z8 SrcIntelDoc
				me.Documentation = z8
			} else {
				if false {
				}
				me.Documentation.__gent__jsonUnmarshal_FromAny(v9)
			}
		} else {
			var z8 SrcIntelDoc
			me.Documentation = z8
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigParam) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcLenses) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(SrcLenses, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				sl[si1] = nil
			} else {
				if nil == sl[si1] {
					sl[si1] = new(SrcLens)
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			sl[si1] = nil
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLenses) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *SrcLens) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["SrcLoc"]
	if !o2 {
		var z3 SrcLoc
		me.SrcLoc = z3
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				var z4 SrcLoc
				me.SrcLoc = z4
			} else {
				if false {
				}
				me.SrcLoc.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			var z4 SrcLoc
			me.SrcLoc = z4
		}
	}
	v6, o7 := v["t"]
	if !o7 {
		me.Txt = ""
	} else {
		println(v6)
		if nil != v6 {
			me.Txt = (string)(v6.(string))
		} else {
			me.Txt = ""
		}
	}
	v10, o11 := v["s"]
	if !o11 {
		me.Str = ""
	} else {
		println(v10)
		if nil != v10 {
			me.Str = (string)(v10.(string))
		} else {
			me.Str = ""
		}
	}
	v14, o15 := v["l"]
	if !o15 {
		me.CrLf = false
	} else {
		println(v14)
		if nil != v14 {
			me.CrLf = (bool)(v14.(bool))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcLoc) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["e"]
	if !o2 {
		me.Flag = 0
	} else {
		println(v1)
		if nil != v1 {
			me.Flag = (int)(v1.(float64))
		} else {
			me.Flag = 0
		}
	}
	v5, o6 := v["f"]
	if !o6 {
		me.FilePath = ""
	} else {
		println(v5)
		if nil != v5 {
			me.FilePath = (string)(v5.(string))
		} else {
			me.FilePath = ""
		}
	}
	v9, o10 := v["p"]
	if !o10 {
		me.Pos = nil
	} else {
		println(v9)
		if nil != v9 {
			v13 := v9.(map[string]interface{})
			if v13 == nil {
				me.Pos = nil
			} else {
				if nil == me.Pos {
					me.Pos = new(SrcPos)
				}
				me.Pos.__gent__jsonUnmarshal_FromAny(v13)
			}
		} else {
			me.Pos = nil
		}
	}
	v14, o15 := v["r"]
	if !o15 {
		me.Range = nil
	} else {
		println(v14)
		if nil != v14 {
			v18 := v14.(map[string]interface{})
			if v18 == nil {
				me.Range = nil
			} else {
				if nil == me.Range {
					me.Range = new(SrcRange)
				}
				me.Range.__gent__jsonUnmarshal_FromAny(v18)
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcLocs) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(SrcLocs, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				sl[si1] = nil
			} else {
				if nil == sl[si1] {
					sl[si1] = new(SrcLoc)
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			sl[si1] = nil
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLocs) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *SrcModEdit) unm(v []byte) (err error) {
	jd := pkg__encoding_json.NewDecoder(nil)
	var t pkg__encoding_json.Token
	if t, err = jd.Token(); err == nil && t != nil {
		for jd.More() {

		}

		_, err = jd.Token()
	}
	return
}

func (me *SrcModEdit) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["At"]
	if !o2 {
		me.At = nil
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				me.At = nil
			} else {
				if nil == me.At {
					me.At = new(SrcRange)
				}
				me.At.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			me.At = nil
		}
	}
	v6, o7 := v["Val"]
	if !o7 {
		me.Val = ""
	} else {
		println(v6)
		if nil != v6 {
			me.Val = (string)(v6.(string))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcModEdits) __gent__jsonUnmarshal_FromAny(v []interface{}) {
	sl := *me
	if len(sl) >= len(v) {
		sl = sl[0:len(v)]
	} else {
		sl = make(SrcModEdits, len(v))
	}
	for si1, sv2 := range v {
		println(sv2)
		if nil != sv2 {
			v4 := sv2.(map[string]interface{})
			if v4 == nil {
				var z3 SrcModEdit
				sl[si1] = z3
			} else {
				if false {
				}
				sl[si1].__gent__jsonUnmarshal_FromAny(v4)
			}
		} else {
			var z3 SrcModEdit
			sl[si1] = z3
		}
	}
	*me = sl
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdits) preview_UnmarshalJSON(v []byte) (err error) {
	var sl []interface{}
	err = pkg__encoding_json.Unmarshal(v, &sl)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(sl)
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

func (me *SrcPos) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["l"]
	if !o2 {
		me.Ln = 0
	} else {
		println(v1)
		if nil != v1 {
			me.Ln = (int)(v1.(float64))
		} else {
			me.Ln = 0
		}
	}
	v5, o6 := v["c"]
	if !o6 {
		me.Col = 0
	} else {
		println(v5)
		if nil != v5 {
			me.Col = (int)(v5.(float64))
		} else {
			me.Col = 0
		}
	}
	v9, o10 := v["o"]
	if !o10 {
		me.Off = 0
	} else {
		println(v9)
		if nil != v9 {
			me.Off = (int)(v9.(float64))
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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

func (me *SrcRange) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["s"]
	if !o2 {
		var z3 SrcPos
		me.Start = z3
	} else {
		println(v1)
		if nil != v1 {
			v5 := v1.(map[string]interface{})
			if v5 == nil {
				var z4 SrcPos
				me.Start = z4
			} else {
				if false {
				}
				me.Start.__gent__jsonUnmarshal_FromAny(v5)
			}
		} else {
			var z4 SrcPos
			me.Start = z4
		}
	}
	v6, o7 := v["e"]
	if !o7 {
		var z8 SrcPos
		me.End = z8
	} else {
		println(v6)
		if nil != v6 {
			v10 := v6.(map[string]interface{})
			if v10 == nil {
				var z9 SrcPos
				me.End = z9
			} else {
				if false {
				}
				me.End.__gent__jsonUnmarshal_FromAny(v10)
			}
		} else {
			var z9 SrcPos
			me.End = z9
		}
	}
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) preview_UnmarshalJSON(v []byte) (err error) {
	var kvs = make(map[string]interface{}, 2)
	err = pkg__encoding_json.Unmarshal(v, &kvs)
	if err == nil {
		me.__gent__jsonUnmarshal_FromAny(kvs)
	}
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) preview_MarshalJSON() (r []byte, err error) {
	panic("WorkspaceChanges")
	return
}

func (me *WorkspaceChanges) __gent__jsonUnmarshal_FromAny(v map[string]interface{}) {
	v1, o2 := v["AddedDirs"]
	if !o2 {
		me.AddedDirs = nil
	} else {
		println(v1)
		if nil != v1 {
			s5 := v1.([]interface{})
			if s5 == nil {
				me.AddedDirs = nil
			} else {
				if len(me.AddedDirs) >= len(s5) {
					me.AddedDirs = me.AddedDirs[0:len(s5)]
				} else {
					me.AddedDirs = make([]string, len(s5))
				}
				for si6, sv7 := range s5 {
					println(sv7)
					if nil != sv7 {
						me.AddedDirs[si6] = (string)(sv7.(string))
					} else {
						me.AddedDirs[si6] = ""
					}
				}
			}
		} else {
			me.AddedDirs = nil
		}
	}
	v9, o10 := v["RemovedDirs"]
	if !o10 {
		me.RemovedDirs = nil
	} else {
		println(v9)
		if nil != v9 {
			s13 := v9.([]interface{})
			if s13 == nil {
				me.RemovedDirs = nil
			} else {
				if len(me.RemovedDirs) >= len(s13) {
					me.RemovedDirs = me.RemovedDirs[0:len(s13)]
				} else {
					me.RemovedDirs = make([]string, len(s13))
				}
				for si14, sv15 := range s13 {
					println(sv15)
					if nil != sv15 {
						me.RemovedDirs[si14] = (string)(sv15.(string))
					} else {
						me.RemovedDirs[si14] = ""
					}
				}
			}
		} else {
			me.RemovedDirs = nil
		}
	}
	v17, o18 := v["OpenedFiles"]
	if !o18 {
		me.OpenedFiles = nil
	} else {
		println(v17)
		if nil != v17 {
			s21 := v17.([]interface{})
			if s21 == nil {
				me.OpenedFiles = nil
			} else {
				if len(me.OpenedFiles) >= len(s21) {
					me.OpenedFiles = me.OpenedFiles[0:len(s21)]
				} else {
					me.OpenedFiles = make([]string, len(s21))
				}
				for si22, sv23 := range s21 {
					println(sv23)
					if nil != sv23 {
						me.OpenedFiles[si22] = (string)(sv23.(string))
					} else {
						me.OpenedFiles[si22] = ""
					}
				}
			}
		} else {
			me.OpenedFiles = nil
		}
	}
	v25, o26 := v["ClosedFiles"]
	if !o26 {
		me.ClosedFiles = nil
	} else {
		println(v25)
		if nil != v25 {
			s29 := v25.([]interface{})
			if s29 == nil {
				me.ClosedFiles = nil
			} else {
				if len(me.ClosedFiles) >= len(s29) {
					me.ClosedFiles = me.ClosedFiles[0:len(s29)]
				} else {
					me.ClosedFiles = make([]string, len(s29))
				}
				for si30, sv31 := range s29 {
					println(sv31)
					if nil != sv31 {
						me.ClosedFiles[si30] = (string)(sv31.(string))
					} else {
						me.ClosedFiles[si30] = ""
					}
				}
			}
		} else {
			me.ClosedFiles = nil
		}
	}
	v33, o34 := v["WrittenFiles"]
	if !o34 {
		me.WrittenFiles = nil
	} else {
		println(v33)
		if nil != v33 {
			s37 := v33.([]interface{})
			if s37 == nil {
				me.WrittenFiles = nil
			} else {
				if len(me.WrittenFiles) >= len(s37) {
					me.WrittenFiles = me.WrittenFiles[0:len(s37)]
				} else {
					me.WrittenFiles = make([]string, len(s37))
				}
				for si38, sv39 := range s37 {
					println(sv39)
					if nil != sv39 {
						me.WrittenFiles[si38] = (string)(sv39.(string))
					} else {
						me.WrittenFiles[si38] = ""
					}
				}
			}
		} else {
			me.WrittenFiles = nil
		}
	}
	v41, o42 := v["LiveFiles"]
	if !o42 {
		me.LiveFiles = nil
	} else {
		println(v41)
		if nil != v41 {
			m45 := v41.(map[string]interface{})
			if m45 == nil {
				me.LiveFiles = nil
			} else {
				me.LiveFiles = make(map[string]string, len(m45))
				for mk46, mv47 := range m45 {
					var t48 string
					println(mv47)
					if nil != mv47 {
						t48 = (string)(mv47.(string))
					} else {
						t48 = ""
					}
					me.LiveFiles[mk46] = t48
				}
			}
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
		me.__gent__jsonUnmarshal_FromAny(kvs)
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
