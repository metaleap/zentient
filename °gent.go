package z

// DON'T EDIT: code gen'd with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__bytes "bytes"
	pkg__encoding_json "encoding/json"
	pkg__errors "errors"
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

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *IpcReq) MarshalJSON() (r []byte, err error) { panic("IpcReq"); return }

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcReq) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *IpcReq) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t34 pkg__encoding_json.Token
	t34, err = j.Token()
	if (err == nil) && (t34 != nil) {
		switch d35 := t34.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d35 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk26 pkg__encoding_json.Token
					jk26, err = j.Token()
					if err == nil {
						fn27 := jk26.(string)
						switch fn27 {
						case "ri":
							var jt28 pkg__encoding_json.Token
							jt28, err = j.Token()
							if (err == nil) && (jt28 != nil) {
								switch v := jt28.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v29 int64
									v29, err = v.Int64()
									if err == nil {
										me.ReqID = (int64)(v29)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ii":
							var jt30 pkg__encoding_json.Token
							jt30, err = j.Token()
							if (err == nil) && (jt30 != nil) {
								switch v := jt30.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v31 int64
									v31, err = v.Int64()
									if err == nil {
										me.IpcID = (IpcIDs)(v31)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ia":
							me.IpcArgs, err = __gent__jsonUnmarshal_interface____(j)
						case "projUpd":
							var pv32 WorkspaceChanges
							err = pv32.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.ProjUpd = &pv32
							}
						case "srcLens":
							var pv33 SrcLens
							err = pv33.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcLens = &pv33
							}
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *IpcResp) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.SrcIntel.MarshalJSON()
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
				sl, e = me.SrcDiags.MarshalJSON()
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
			sl, e = me.SrcMods.MarshalJSON()
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
			ai27 := len(r)
			for i26 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i26].MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai27] == 44 {
				r = append(r[:ai27], r[(ai27+1):]...)
			}
		}
		if nil != me.Extras {
			{
				r = append(r, ",\"extras\":"...)
				var e error
				var sl []byte
				sl, e = me.Extras.MarshalJSON()
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
				sl, e = me.Menu.MarshalJSON()
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
				sl, e = me.CaddyUpdate.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcResp) UnmarshalJSON(b []byte) (err error) { panic("IpcResp"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Diags) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.All.MarshalJSON()
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
			ai11 := len(r)
			for i10 := range me.FixUps {
				if nil != me.FixUps[i10] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.FixUps[i10].MarshalJSON()
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
			if r[ai11] == 44 {
				r = append(r[:ai11], r[(ai11+1):]...)
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Diags) UnmarshalJSON(b []byte) (err error) { panic("Diags"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Extras) MarshalJSON() (r []byte, err error) {
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
					sl, e = me.SrcIntels.InfoTips[i2].MarshalJSON()
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
			sl, e = me.SrcIntels.Refs.MarshalJSON()
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
			ai21 := len(r)
			for i20 := range me.Items {
				if nil != me.Items[i20] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Items[i20].MarshalJSON()
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
			if r[ai21] == 44 {
				r = append(r[:ai21], r[(ai21+1):]...)
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Extras) UnmarshalJSON(b []byte) (err error) { panic("Extras"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuResponse) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.SubMenu.MarshalJSON()
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
			sl, e = me.Refs.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuResponse) UnmarshalJSON(b []byte) (err error) { panic("MenuResponse"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntel) MarshalJSON() (r []byte, err error) {
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
					sl, e = me.SrcIntels.InfoTips[i2].MarshalJSON()
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
			sl, e = me.SrcIntels.Refs.MarshalJSON()
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
				sl, e = me.Sig.MarshalJSON()
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
			sl, e = me.Cmpl.MarshalJSON()
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
			sl, e = me.Syms.MarshalJSON()
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
			ai45 := len(r)
			for i44 := range me.Anns {
				if nil != me.Anns[i44] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Anns[i44].MarshalJSON()
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
			if r[ai45] == 44 {
				r = append(r[:ai45], r[(ai45+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntel) UnmarshalJSON(b []byte) (err error) { panic("SrcIntel"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Caddy) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Caddy) UnmarshalJSON(b []byte) (err error) { panic("Caddy"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *DiagFixUps) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Edits.MarshalJSON()
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
			ai14 := len(r)
			for i13 := range me.Dropped {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Dropped[i13].MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai14] == 44 {
				r = append(r[:ai14], r[(ai14+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagFixUps) UnmarshalJSON(b []byte) (err error) { panic("DiagFixUps"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *DiagItem) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Loc.MarshalJSON()
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
			ai11 := len(r)
			for i10 := range me.Rel {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Rel[i10].MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai11] == 44 {
				r = append(r[:ai11], r[(ai11+1):]...)
			}
		}
		if len(me.SrcActions) != 0 {
			r = append(r, ",\"SrcActions\":"...)
			r = append(r, 91)
			ai21 := len(r)
			for i20 := range me.SrcActions {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.SrcActions[i20].MarshalJSON()
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
		if me.StickyAuto {
			r = append(r, ",\"Sticky\":"...)
			r = append(r, pkg__strconv.FormatBool(me.StickyAuto)...)
		}
		if len(me.Tags) != 0 {
			r = append(r, ",\"Tags\":"...)
			r = append(r, 91)
			ai31 := len(r)
			for i30 := range me.Tags {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.FormatInt((int64)(me.Tags[i30]), 10)...)
				}
			}
			r = append(r, 93)
			if r[ai31] == 44 {
				r = append(r[:ai31], r[(ai31+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItem) UnmarshalJSON(b []byte) (err error) { panic("DiagItem"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me DiagItems) MarshalJSON() (r []byte, err error) {
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
					sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItems) UnmarshalJSON(b []byte) (err error) { panic("DiagItems"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me DiagItemsBy) MarshalJSON() (r []byte, err error) {
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
				sl, e = mv3.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItemsBy) UnmarshalJSON(b []byte) (err error) { panic("DiagItemsBy"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *EditorAction) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *EditorAction) UnmarshalJSON(b []byte) (err error) { panic("EditorAction"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *ExtrasItem) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *ExtrasItem) UnmarshalJSON(b []byte) (err error) { panic("ExtrasItem"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *Menu) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Items.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Menu) UnmarshalJSON(b []byte) (err error) { panic("Menu"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me MenuItems) MarshalJSON() (r []byte, err error) {
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
					sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItems) UnmarshalJSON(b []byte) (err error) { panic("MenuItems"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuItem) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItem) UnmarshalJSON(b []byte) (err error) { panic("MenuItem"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *MenuItemArgPrompt) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItemArgPrompt) UnmarshalJSON(b []byte) (err error) { panic("MenuItemArgPrompt"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcAnnotaction) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Range.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcAnnotaction) UnmarshalJSON(b []byte) (err error) { panic("SrcAnnotaction"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcInfoTip) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcInfoTip) UnmarshalJSON(b []byte) (err error) { panic("SrcInfoTip"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelCompl) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.Documentation.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompl) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelCompl"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcIntelCompls) MarshalJSON() (r []byte, err error) {
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
					sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompls) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelCompls"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntels) MarshalJSON() (r []byte, err error) {
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
					sl, e = me.InfoTips[i2].MarshalJSON()
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
			sl, e = me.Refs.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntels) UnmarshalJSON(b []byte) (err error) { panic("SrcIntels"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelDoc) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelDoc) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelDoc"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigHelp) MarshalJSON() (r []byte, err error) {
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
					sl, e = me.Signatures[i2].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigHelp) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelSigHelp"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigInfo) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Documentation.MarshalJSON()
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
			ai11 := len(r)
			for i10 := range me.Parameters {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Parameters[i10].MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
			if r[ai11] == 44 {
				r = append(r[:ai11], r[(ai11+1):]...)
			}
		}
		r = append(r, 125)
		if r[si1] == 44 {
			r = append(r[:si1], r[(si1+1):]...)
		}
	}
	return
}

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigInfo) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelSigInfo"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigParam) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Documentation.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigParam) UnmarshalJSON(b []byte) (err error) { panic("SrcIntelSigParam"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcLenses) MarshalJSON() (r []byte, err error) {
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
					sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLenses) UnmarshalJSON(b []byte) (err error) { panic("SrcLenses"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLens) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.SrcLoc.Pos.MarshalJSON()
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
				sl, e = me.SrcLoc.Range.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLens) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *SrcLens) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t7 pkg__encoding_json.Token
	t7, err = j.Token()
	if (err == nil) && (t7 != nil) {
		switch d8 := t7.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d8 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "e":
							me.SrcLoc.Flag, err = __gent__jsonUnmarshal_int(j)
						case "f":
							me.SrcLoc.FilePath, err = __gent__jsonUnmarshal_string(j)
						case "p":
							var pv3 SrcPos
							err = pv3.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcLoc.Pos = &pv3
							}
						case "r":
							var pv4 SrcRange
							err = pv4.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcLoc.Range = &pv4
							}
						case "t":
							me.Txt, err = __gent__jsonUnmarshal_string(j)
						case "s":
							me.Str, err = __gent__jsonUnmarshal_string(j)
						case "l":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case bool:
									me.CrLf = v
								default:
									err = pkg__errors.New("expected bool")
								}
							}
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLoc) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.Pos.MarshalJSON()
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
				sl, e = me.Range.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLoc) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *SrcLoc) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t5 pkg__encoding_json.Token
	t5, err = j.Token()
	if (err == nil) && (t5 != nil) {
		switch d6 := t5.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d6 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "e":
							me.Flag, err = __gent__jsonUnmarshal_int(j)
						case "f":
							me.FilePath, err = __gent__jsonUnmarshal_string(j)
						case "p":
							var pv3 SrcPos
							err = pv3.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Pos = &pv3
							}
						case "r":
							var pv4 SrcRange
							err = pv4.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Range = &pv4
							}
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcLocs) MarshalJSON() (r []byte, err error) {
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
					sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLocs) UnmarshalJSON(b []byte) (err error) { panic("SrcLocs"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcModEdit) MarshalJSON() (r []byte, err error) {
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
				sl, e = me.At.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdit) UnmarshalJSON(b []byte) (err error) { panic("SrcModEdit"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me SrcModEdits) MarshalJSON() (r []byte, err error) {
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
				sl, e = me[i1].MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdits) UnmarshalJSON(b []byte) (err error) { panic("SrcModEdits"); return }

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcPos) MarshalJSON() (r []byte, err error) {
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcPos) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *SrcPos) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t3 pkg__encoding_json.Token
	t3, err = j.Token()
	if (err == nil) && (t3 != nil) {
		switch d4 := t3.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d4 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "l":
							me.Ln, err = __gent__jsonUnmarshal_int(j)
						case "c":
							me.Col, err = __gent__jsonUnmarshal_int(j)
						case "o":
							me.Off, err = __gent__jsonUnmarshal_int(j)
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcRange) MarshalJSON() (r []byte, err error) {
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
			sl, e = me.Start.MarshalJSON()
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
			sl, e = me.End.MarshalJSON()
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

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *SrcRange) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t3 pkg__encoding_json.Token
	t3, err = j.Token()
	if (err == nil) && (t3 != nil) {
		switch d4 := t3.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d4 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "s":
							err = me.Start.__gent__jsonUnmarshal_Decode(j)
						case "e":
							err = me.End.__gent__jsonUnmarshal_Decode(j)
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) MarshalJSON() (r []byte, err error) { panic("WorkspaceChanges"); return }

// UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *WorkspaceChanges) UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

func (me *WorkspaceChanges) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t9 pkg__encoding_json.Token
	t9, err = j.Token()
	if (err == nil) && (t9 != nil) {
		switch d10 := t9.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d10 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "AddedDirs":
							me.AddedDirs, err = __gent__jsonUnmarshal_s_string(j)
						case "RemovedDirs":
							me.RemovedDirs, err = __gent__jsonUnmarshal_s_string(j)
						case "OpenedFiles":
							me.OpenedFiles, err = __gent__jsonUnmarshal_s_string(j)
						case "ClosedFiles":
							me.ClosedFiles, err = __gent__jsonUnmarshal_s_string(j)
						case "WrittenFiles":
							me.WrittenFiles, err = __gent__jsonUnmarshal_s_string(j)
						case "LiveFiles":
							var t7 pkg__encoding_json.Token
							t7, err = j.Token()
							if (err == nil) && (t7 != nil) {
								switch d8 := t7.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x7b != d8 {
										err = pkg__errors.New("expected {")
									} else {
										t6 := make(map[string]string, 0)
										for (err == nil) && j.More() {
											var jk3 pkg__encoding_json.Token
											jk3, err = j.Token()
											if err == nil {
												mk4 := jk3.(string)
												var mv5 string
												mv5, err = __gent__jsonUnmarshal_string(j)
												if err == nil {
													t6[mk4] = mv5
												}
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.LiveFiles = t6
										}
									}
								default:
									err = pkg__errors.New("expected {")
								}
							}
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
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
				v8, ok9 := v.(map[string]interface{})
				if ok9 {
					sl, e = __gent__jsonMarshal_mapsstring_interface____(v8)
				} else {
					v6, ok7 := v.(MenuItemArgPrompt)
					if ok7 {
						sl, e = __gent__jsonMarshal_MenuItemArgPrompt(v6)
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
								panic(v)
							}
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
		ai11 := len(r)
		for i10 := range v {
			{
				r = append(r, 44)
				r = append(r, pkg__strconv.Quote(v[i10])...)
			}
		}
		r = append(r, 93)
		if r[ai11] == 44 {
			r = append(r[:ai11], r[(ai11+1):]...)
		}
	}
	return
}

func __gent__jsonMarshal_MenuItemArgPrompt(v MenuItemArgPrompt) (r []byte, err error) {
	{
		var e error
		var sl []byte
		sl, e = v.MarshalJSON()
		if e == nil {
			r = append(r, sl...)
		} else {
			err = e
			return
		}
	}
	return
}

func __gent__jsonMarshal_mapsstring_interface____(v map[string]interface{}) (r []byte, err error) {
	if nil == v {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		mi20 := len(r)
		for mk21, mv22 := range v {
			sl, e := __gent__jsonMarshal_interface____(mv22)
			if e == nil {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(mk21)...)
					r = append(r, 58)
				}
				r = append(r, sl...)
			} else {
				err = e
				return
			}
		}
		r = append(r, 125)
		if r[mi20] == 44 {
			r = append(r[:mi20], r[(mi20+1):]...)
		}
	}
	return
}

func __gent__jsonUnmarshal_interface____(j *pkg__encoding_json.Decoder) (r interface{}, err error) {
	var ttt1 pkg__encoding_json.Token
	ttt1, err = j.Token()
	if (err == nil) && (ttt1 != nil) {
		switch v := ttt1.(type) {
		case nil:
		case string:
			r = v
		case bool:
			r = v
		case pkg__encoding_json.Number:
			r, err = v.Float64()
		case pkg__encoding_json.Delim:
			switch v {
			case 123:
				var mmm2 map[string]interface{}
				t7 := make(map[string]interface{}, 0)
				for (err == nil) && j.More() {
					var jk4 pkg__encoding_json.Token
					jk4, err = j.Token()
					if err == nil {
						mk5 := jk4.(string)
						var mv6 interface{}
						mv6, err = __gent__jsonUnmarshal_interface____(j)
						if err == nil {
							t7[mk5] = mv6
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					mmm2 = t7
				}
				if err == nil {
					r = mmm2
				}
			case 91:
				var slsl3 []interface{}
				sl10 := make([]interface{}, 0, 0)
				for (err == nil) && j.More() {
					var v11 interface{}
					v11, err = __gent__jsonUnmarshal_interface____(j)
					if err == nil {
						sl10 = append(sl10, v11)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					slsl3 = sl10
				}
				if err == nil {
					r = slsl3
				}
			}
		}
	}
	return
}

func __gent__jsonUnmarshal_s_string(j *pkg__encoding_json.Decoder) (r []string, err error) {
	var t16 pkg__encoding_json.Token
	t16, err = j.Token()
	if (err == nil) && (t16 != nil) {
		switch d17 := t16.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d17 {
				err = pkg__errors.New("expected [")
			} else {
				sl14 := make([]string, 0, 0)
				for (err == nil) && j.More() {
					var v15 string
					v15, err = __gent__jsonUnmarshal_string(j)
					if err == nil {
						sl14 = append(sl14, v15)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					r = sl14
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

func __gent__jsonUnmarshal_s_interface____(j *pkg__encoding_json.Decoder) (r []interface{}, err error) {
	var t20 pkg__encoding_json.Token
	t20, err = j.Token()
	if (err == nil) && (t20 != nil) {
		switch d21 := t20.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d21 {
				err = pkg__errors.New("expected [")
			} else {
				sl18 := make([]interface{}, 0, 0)
				for (err == nil) && j.More() {
					var v19 interface{}
					v19, err = __gent__jsonUnmarshal_interface____(j)
					if err == nil {
						sl18 = append(sl18, v19)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					r = sl18
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

func __gent__jsonUnmarshal_string(j *pkg__encoding_json.Decoder) (r string, err error) {
	var jt22 pkg__encoding_json.Token
	jt22, err = j.Token()
	if (err == nil) && (jt22 != nil) {
		switch v := jt22.(type) {
		case nil:
		case string:
			r = v
		default:
			err = pkg__errors.New("expected string")
		}
	}
	return
}

func __gent__jsonUnmarshal_int(j *pkg__encoding_json.Decoder) (r int, err error) {
	var jt24 pkg__encoding_json.Token
	jt24, err = j.Token()
	if (err == nil) && (jt24 != nil) {
		switch v := jt24.(type) {
		case nil:
		case pkg__encoding_json.Number:
			var v25 int64
			v25, err = v.Int64()
			if err == nil {
				r = (int)(v25)
			}
		default:
			err = pkg__errors.New("expected pkg__encoding_json.Number")
		}
	}
	return
}
