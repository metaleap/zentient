package z

// DON'T EDIT: code gen'd with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__bytes "bytes"
	pkg__encoding_json "encoding/json"
	pkg__errors "errors"
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

func (me *IpcReq) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "ri":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.ReqID = (int64)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ii":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v6 int64
									v6, err = v.Int64()
									if err == nil {
										me.IpcID = (IpcIDs)(v6)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ia":
						case "projUpd":
							var pv7 WorkspaceChanges
							err = pv7.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.ProjUpd = &pv7
							}
						case "srcLens":
							var pv8 SrcLens
							err = pv8.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcLens = &pv8
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcReq) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *IpcResp) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t18 pkg__encoding_json.Token
	t18, err = j.Token()
	if (err == nil) && (t18 != nil) {
		switch d19 := t18.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d19 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "ii":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.IpcID = (IpcIDs)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ri":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v6 int64
									v6, err = v.Int64()
									if err == nil {
										me.ReqID = (int64)(v6)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "err":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.ErrMsg = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "sI":
							var pv9 SrcIntel
							err = pv9.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcIntel = &pv9
							}
						case "srcDiags":
							var pv10 Diags
							err = pv10.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SrcDiags = &pv10
							}
						case "srcMods":
							err = me.SrcMods.__gent__jsonUnmarshal_Decode(j)
						case "srcActions":
							var t13 pkg__encoding_json.Token
							t13, err = j.Token()
							if (err == nil) && (t13 != nil) {
								switch d14 := t13.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d14 {
										err = pkg__errors.New("expected [")
									} else {
										sl11 := make([]EditorAction, 0, 0)
										for (err == nil) && j.More() {
											var v12 EditorAction
											err = v12.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl11 = append(sl11, v12)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.SrcActions = sl11
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "extras":
							var pv15 Extras
							err = pv15.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Extras = &pv15
							}
						case "menu":
							var pv16 MenuResponse
							err = pv16.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Menu = &pv16
							}
						case "caddy":
							var pv17 Caddy
							err = pv17.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.CaddyUpdate = &pv17
							}
						case "val":
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *IpcResp) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *Diags) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t10 pkg__encoding_json.Token
	t10, err = j.Token()
	if (err == nil) && (t10 != nil) {
		switch d11 := t10.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d11 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "All":
							err = me.All.__gent__jsonUnmarshal_Decode(j)
						case "FixUps":
							var t6 pkg__encoding_json.Token
							t6, err = j.Token()
							if (err == nil) && (t6 != nil) {
								switch d7 := t6.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d7 {
										err = pkg__errors.New("expected [")
									} else {
										sl3 := make([]*DiagFixUps, 0, 0)
										for (err == nil) && j.More() {
											var v4 *DiagFixUps
											var pv5 DiagFixUps
											err = pv5.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												v4 = &pv5
											}
											if err == nil {
												sl3 = append(sl3, v4)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.FixUps = sl3
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "LangID":
							var jt8 pkg__encoding_json.Token
							jt8, err = j.Token()
							if (err == nil) && (jt8 != nil) {
								switch v := jt8.(type) {
								case nil:
								case string:
									me.LangID = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Diags) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *Extras) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t18 pkg__encoding_json.Token
	t18, err = j.Token()
	if (err == nil) && (t18 != nil) {
		switch d19 := t18.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d19 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "SrcIntels":
							err = me.SrcIntels.__gent__jsonUnmarshal_Decode(j)
						case "Items":
							var t6 pkg__encoding_json.Token
							t6, err = j.Token()
							if (err == nil) && (t6 != nil) {
								switch d7 := t6.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d7 {
										err = pkg__errors.New("expected [")
									} else {
										sl3 := make([]*ExtrasItem, 0, 0)
										for (err == nil) && j.More() {
											var v4 *ExtrasItem
											var pv5 ExtrasItem
											err = pv5.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												v4 = &pv5
											}
											if err == nil {
												sl3 = append(sl3, v4)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Items = sl3
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "Warns":
							var t12 pkg__encoding_json.Token
							t12, err = j.Token()
							if (err == nil) && (t12 != nil) {
								switch d13 := t12.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d13 {
										err = pkg__errors.New("expected [")
									} else {
										sl8 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v9 string
											var jt10 pkg__encoding_json.Token
											jt10, err = j.Token()
											if (err == nil) && (jt10 != nil) {
												switch v := jt10.(type) {
												case nil:
												case string:
													v9 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl8 = append(sl8, v9)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Warns = sl8
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "Desc":
							var jt14 pkg__encoding_json.Token
							jt14, err = j.Token()
							if (err == nil) && (jt14 != nil) {
								switch v := jt14.(type) {
								case nil:
								case string:
									me.Desc = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Url":
							var jt16 pkg__encoding_json.Token
							jt16, err = j.Token()
							if (err == nil) && (jt16 != nil) {
								switch v := jt16.(type) {
								case nil:
								case string:
									me.Url = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Extras) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *MenuResponse) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t12 pkg__encoding_json.Token
	t12, err = j.Token()
	if (err == nil) && (t12 != nil) {
		switch d13 := t12.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d13 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "SubMenu":
							var pv3 Menu
							err = pv3.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.SubMenu = &pv3
							}
						case "WebsiteURL":
							var jt4 pkg__encoding_json.Token
							jt4, err = j.Token()
							if (err == nil) && (jt4 != nil) {
								switch v := jt4.(type) {
								case nil:
								case string:
									me.WebsiteURL = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "NoteInfo":
							var jt6 pkg__encoding_json.Token
							jt6, err = j.Token()
							if (err == nil) && (jt6 != nil) {
								switch v := jt6.(type) {
								case nil:
								case string:
									me.NoteInfo = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "NoteWarn":
							var jt8 pkg__encoding_json.Token
							jt8, err = j.Token()
							if (err == nil) && (jt8 != nil) {
								switch v := jt8.(type) {
								case nil:
								case string:
									me.NoteWarn = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "UxActionLabel":
							var jt10 pkg__encoding_json.Token
							jt10, err = j.Token()
							if (err == nil) && (jt10 != nil) {
								switch v := jt10.(type) {
								case nil:
								case string:
									me.UxActionLabel = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Refs":
							err = me.Refs.__gent__jsonUnmarshal_Decode(j)
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuResponse) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntel) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "SrcIntels":
							err = me.SrcIntels.__gent__jsonUnmarshal_Decode(j)
						case "Sig":
							var pv3 SrcIntelSigHelp
							err = pv3.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Sig = &pv3
							}
						case "Cmpl":
							err = me.Cmpl.__gent__jsonUnmarshal_Decode(j)
						case "Syms":
							err = me.Syms.__gent__jsonUnmarshal_Decode(j)
						case "Anns":
							var t7 pkg__encoding_json.Token
							t7, err = j.Token()
							if (err == nil) && (t7 != nil) {
								switch d8 := t7.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d8 {
										err = pkg__errors.New("expected [")
									} else {
										sl4 := make([]*SrcAnnotaction, 0, 0)
										for (err == nil) && j.More() {
											var v5 *SrcAnnotaction
											var pv6 SrcAnnotaction
											err = pv6.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												v5 = &pv6
											}
											if err == nil {
												sl4 = append(sl4, v5)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Anns = sl4
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntel) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *Caddy) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t25 pkg__encoding_json.Token
	t25, err = j.Token()
	if (err == nil) && (t25 != nil) {
		switch d26 := t25.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d26 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "ID":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.ID = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "LangID":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.LangID = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Icon":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.Icon = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Title":
							var jt9 pkg__encoding_json.Token
							jt9, err = j.Token()
							if (err == nil) && (jt9 != nil) {
								switch v := jt9.(type) {
								case nil:
								case string:
									me.Title = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Status":
							var t17 pkg__encoding_json.Token
							t17, err = j.Token()
							if (err == nil) && (t17 != nil) {
								switch d18 := t17.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x7b != d18 {
										err = pkg__errors.New("expected {")
									} else {
										for (err == nil) && j.More() {
											var jk11 pkg__encoding_json.Token
											jk11, err = j.Token()
											if err == nil {
												fn12 := jk11.(string)
												switch fn12 {
												case "Flag":
													var jt13 pkg__encoding_json.Token
													jt13, err = j.Token()
													if (err == nil) && (jt13 != nil) {
														switch v := jt13.(type) {
														case nil:
														case pkg__encoding_json.Number:
															var v14 int64
															v14, err = v.Int64()
															if err == nil {
																me.Status.Flag = (CaddyStatus)(v14)
															}
														default:
															err = pkg__errors.New("expected pkg__encoding_json.Number")
														}
													}
												case "Desc":
													var jt15 pkg__encoding_json.Token
													jt15, err = j.Token()
													if (err == nil) && (jt15 != nil) {
														switch v := jt15.(type) {
														case nil:
														case string:
															me.Status.Desc = v
														default:
															err = pkg__errors.New("expected string")
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
						case "Details":
							var jt19 pkg__encoding_json.Token
							jt19, err = j.Token()
							if (err == nil) && (jt19 != nil) {
								switch v := jt19.(type) {
								case nil:
								case string:
									me.Details = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "UxActionID":
							var jt21 pkg__encoding_json.Token
							jt21, err = j.Token()
							if (err == nil) && (jt21 != nil) {
								switch v := jt21.(type) {
								case nil:
								case string:
									me.UxActionID = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "ShowTitle":
							var jt23 pkg__encoding_json.Token
							jt23, err = j.Token()
							if (err == nil) && (jt23 != nil) {
								switch v := jt23.(type) {
								case nil:
								case bool:
									me.ShowTitle = v
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Caddy) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *DiagFixUps) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t21 pkg__encoding_json.Token
	t21, err = j.Token()
	if (err == nil) && (t21 != nil) {
		switch d22 := t21.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d22 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "FilePath":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.FilePath = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Desc":
							var t15 pkg__encoding_json.Token
							t15, err = j.Token()
							if (err == nil) && (t15 != nil) {
								switch d16 := t15.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x7b != d16 {
										err = pkg__errors.New("expected {")
									} else {
										t8 := make(map[string][]string, 0)
										for (err == nil) && j.More() {
											var jk5 pkg__encoding_json.Token
											jk5, err = j.Token()
											if err == nil {
												mk6 := jk5.(string)
												var mv7 []string
												var t13 pkg__encoding_json.Token
												t13, err = j.Token()
												if (err == nil) && (t13 != nil) {
													switch d14 := t13.(type) {
													case nil:
													case pkg__encoding_json.Delim:
														if 0x5b != d14 {
															err = pkg__errors.New("expected [")
														} else {
															sl9 := make([]string, 0, 0)
															for (err == nil) && j.More() {
																var v10 string
																var jt11 pkg__encoding_json.Token
																jt11, err = j.Token()
																if (err == nil) && (jt11 != nil) {
																	switch v := jt11.(type) {
																	case nil:
																	case string:
																		v10 = v
																	default:
																		err = pkg__errors.New("expected string")
																	}
																}
																if err == nil {
																	sl9 = append(sl9, v10)
																}
															}
															if err == nil {
																_, err = j.Token()
															}
															if err == nil {
																mv7 = sl9
															}
														}
													default:
														err = pkg__errors.New("expected [")
													}
												}
												if err == nil {
													t8[mk6] = mv7
												}
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Desc = t8
										}
									}
								default:
									err = pkg__errors.New("expected {")
								}
							}
						case "Edits":
							err = me.Edits.__gent__jsonUnmarshal_Decode(j)
						case "Dropped":
							var t19 pkg__encoding_json.Token
							t19, err = j.Token()
							if (err == nil) && (t19 != nil) {
								switch d20 := t19.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d20 {
										err = pkg__errors.New("expected [")
									} else {
										sl17 := make([]SrcModEdit, 0, 0)
										for (err == nil) && j.More() {
											var v18 SrcModEdit
											err = v18.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl17 = append(sl17, v18)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Dropped = sl17
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagFixUps) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *DiagItem) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t23 pkg__encoding_json.Token
	t23, err = j.Token()
	if (err == nil) && (t23 != nil) {
		switch d24 := t23.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d24 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "Cat":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Cat = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Loc":
							err = me.Loc.__gent__jsonUnmarshal_Decode(j)
						case "Msg":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Msg = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Rel":
							var t9 pkg__encoding_json.Token
							t9, err = j.Token()
							if (err == nil) && (t9 != nil) {
								switch d10 := t9.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d10 {
										err = pkg__errors.New("expected [")
									} else {
										sl7 := make([]SrcLens, 0, 0)
										for (err == nil) && j.More() {
											var v8 SrcLens
											err = v8.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl7 = append(sl7, v8)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Rel = sl7
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "SrcActions":
							var t13 pkg__encoding_json.Token
							t13, err = j.Token()
							if (err == nil) && (t13 != nil) {
								switch d14 := t13.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d14 {
										err = pkg__errors.New("expected [")
									} else {
										sl11 := make([]EditorAction, 0, 0)
										for (err == nil) && j.More() {
											var v12 EditorAction
											err = v12.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl11 = append(sl11, v12)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.SrcActions = sl11
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "Sticky":
							var jt15 pkg__encoding_json.Token
							jt15, err = j.Token()
							if (err == nil) && (jt15 != nil) {
								switch v := jt15.(type) {
								case nil:
								case bool:
									me.StickyAuto = v
								default:
									err = pkg__errors.New("expected bool")
								}
							}
						case "Tags":
							var t21 pkg__encoding_json.Token
							t21, err = j.Token()
							if (err == nil) && (t21 != nil) {
								switch d22 := t21.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d22 {
										err = pkg__errors.New("expected [")
									} else {
										sl17 := make([]int, 0, 0)
										for (err == nil) && j.More() {
											var v18 int
											var jt19 pkg__encoding_json.Token
											jt19, err = j.Token()
											if (err == nil) && (jt19 != nil) {
												switch v := jt19.(type) {
												case nil:
												case pkg__encoding_json.Number:
													var v20 int64
													v20, err = v.Int64()
													if err == nil {
														v18 = (int)(v20)
													}
												default:
													err = pkg__errors.New("expected pkg__encoding_json.Number")
												}
											}
											if err == nil {
												sl17 = append(sl17, v18)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Tags = sl17
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItem) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *DiagItems) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t4 pkg__encoding_json.Token
	t4, err = j.Token()
	if (err == nil) && (t4 != nil) {
		switch d5 := t4.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d5 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]*DiagItem, 0, 0)
				for (err == nil) && j.More() {
					var v2 *DiagItem
					var pv3 DiagItem
					err = pv3.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						v2 = &pv3
					}
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItems) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *DiagItemsBy) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t5 pkg__encoding_json.Token
	t5, err = j.Token()
	if (err == nil) && (t5 != nil) {
		switch d6 := t5.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d6 {
				err = pkg__errors.New("expected {")
			} else {
				t4 := make(map[string]DiagItems, 0)
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						mk2 := jk1.(string)
						var mv3 DiagItems
						err = mv3.__gent__jsonUnmarshal_Decode(j)
						if err == nil {
							t4[mk2] = mv3
						}
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = t4
				}
			}
		default:
			err = pkg__errors.New("expected {")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *DiagItemsBy) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *EditorAction) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t13 pkg__encoding_json.Token
	t13, err = j.Token()
	if (err == nil) && (t13 != nil) {
		switch d14 := t13.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d14 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "title":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Title = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "command":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Cmd = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "tooltip":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.Hint = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "arguments":
							var t11 pkg__encoding_json.Token
							t11, err = j.Token()
							if (err == nil) && (t11 != nil) {
								switch d12 := t11.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d12 {
										err = pkg__errors.New("expected [")
									} else {
										sl9 := make([]interface{}, 0, 0)
										for (err == nil) && j.More() {
											var v10 interface{}
											if err == nil {
												sl9 = append(sl9, v10)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Arguments = sl9
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *EditorAction) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *ExtrasItem) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t15 pkg__encoding_json.Token
	t15, err = j.Token()
	if (err == nil) && (t15 != nil) {
		switch d16 := t15.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d16 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "id":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.ID = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "label":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Label = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "description":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.Desc = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "detail":
							var jt9 pkg__encoding_json.Token
							jt9, err = j.Token()
							if (err == nil) && (jt9 != nil) {
								switch v := jt9.(type) {
								case nil:
								case string:
									me.Detail = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "arg":
							var jt11 pkg__encoding_json.Token
							jt11, err = j.Token()
							if (err == nil) && (jt11 != nil) {
								switch v := jt11.(type) {
								case nil:
								case string:
									me.QueryArg = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "fPos":
							var jt13 pkg__encoding_json.Token
							jt13, err = j.Token()
							if (err == nil) && (jt13 != nil) {
								switch v := jt13.(type) {
								case nil:
								case string:
									me.FilePos = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *ExtrasItem) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *Menu) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "desc":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Desc = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "topLevel":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case bool:
									me.TopLevel = v
								default:
									err = pkg__errors.New("expected bool")
								}
							}
						case "items":
							err = me.Items.__gent__jsonUnmarshal_Decode(j)
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *Menu) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *MenuItems) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t4 pkg__encoding_json.Token
	t4, err = j.Token()
	if (err == nil) && (t4 != nil) {
		switch d5 := t4.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d5 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]*MenuItem, 0, 0)
				for (err == nil) && j.More() {
					var v2 *MenuItem
					var pv3 MenuItem
					err = pv3.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						v2 = &pv3
					}
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItems) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *MenuItem) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t15 pkg__encoding_json.Token
	t15, err = j.Token()
	if (err == nil) && (t15 != nil) {
		switch d16 := t15.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d16 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "ii":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.IpcID = (IpcIDs)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "ia":
						case "c":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Category = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "t":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.Title = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "d":
							var jt9 pkg__encoding_json.Token
							jt9, err = j.Token()
							if (err == nil) && (jt9 != nil) {
								switch v := jt9.(type) {
								case nil:
								case string:
									me.Desc = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "h":
							var jt11 pkg__encoding_json.Token
							jt11, err = j.Token()
							if (err == nil) && (jt11 != nil) {
								switch v := jt11.(type) {
								case nil:
								case string:
									me.Hint = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "q":
							var jt13 pkg__encoding_json.Token
							jt13, err = j.Token()
							if (err == nil) && (jt13 != nil) {
								switch v := jt13.(type) {
								case nil:
								case string:
									me.Confirm = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItem) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *MenuItemArgPrompt) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "prompt":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Prompt = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "placeHolder":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Placeholder = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "value":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.Value = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *MenuItemArgPrompt) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcAnnotaction) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "Range":
							err = me.Range.__gent__jsonUnmarshal_Decode(j)
						case "Title":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Title = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "Desc":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Desc = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "CmdName":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case string:
									me.CmdName = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcAnnotaction) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcInfoTip) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "value":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Value = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "language":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Language = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcInfoTip) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelCompl) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t12 pkg__encoding_json.Token
	t12, err = j.Token()
	if (err == nil) && (t12 != nil) {
		switch d13 := t12.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d13 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "kind":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.Kind = (Completion)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "label":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Label = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "documentation":
							var pv7 SrcIntelDoc
							err = pv7.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Documentation = &pv7
							}
						case "detail":
							var jt8 pkg__encoding_json.Token
							jt8, err = j.Token()
							if (err == nil) && (jt8 != nil) {
								switch v := jt8.(type) {
								case nil:
								case string:
									me.Detail = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "sortText":
							var jt10 pkg__encoding_json.Token
							jt10, err = j.Token()
							if (err == nil) && (jt10 != nil) {
								switch v := jt10.(type) {
								case nil:
								case string:
									me.SortText = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompl) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelCompls) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t4 pkg__encoding_json.Token
	t4, err = j.Token()
	if (err == nil) && (t4 != nil) {
		switch d5 := t4.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d5 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]*SrcIntelCompl, 0, 0)
				for (err == nil) && j.More() {
					var v2 *SrcIntelCompl
					var pv3 SrcIntelCompl
					err = pv3.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						v2 = &pv3
					}
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompls) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntels) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "InfoTips":
							var t5 pkg__encoding_json.Token
							t5, err = j.Token()
							if (err == nil) && (t5 != nil) {
								switch d6 := t5.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d6 {
										err = pkg__errors.New("expected [")
									} else {
										sl3 := make([]SrcInfoTip, 0, 0)
										for (err == nil) && j.More() {
											var v4 SrcInfoTip
											err = v4.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl3 = append(sl3, v4)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.InfoTips = sl3
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "Refs":
							err = me.Refs.__gent__jsonUnmarshal_Decode(j)
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntels) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelDoc) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "value":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Value = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "isTrusted":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case bool:
									me.IsTrusted = v
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelDoc) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelSigHelp) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t11 pkg__encoding_json.Token
	t11, err = j.Token()
	if (err == nil) && (t11 != nil) {
		switch d12 := t11.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d12 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "activeSignature":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.ActiveSignature = (int)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "activeParameter":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v6 int64
									v6, err = v.Int64()
									if err == nil {
										me.ActiveParameter = (int)(v6)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "signatures":
							var t9 pkg__encoding_json.Token
							t9, err = j.Token()
							if (err == nil) && (t9 != nil) {
								switch d10 := t9.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d10 {
										err = pkg__errors.New("expected [")
									} else {
										sl7 := make([]SrcIntelSigInfo, 0, 0)
										for (err == nil) && j.More() {
											var v8 SrcIntelSigInfo
											err = v8.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl7 = append(sl7, v8)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Signatures = sl7
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigHelp) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelSigInfo) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "label":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Label = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "documentation":
							err = me.Documentation.__gent__jsonUnmarshal_Decode(j)
						case "parameters":
							var t7 pkg__encoding_json.Token
							t7, err = j.Token()
							if (err == nil) && (t7 != nil) {
								switch d8 := t7.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d8 {
										err = pkg__errors.New("expected [")
									} else {
										sl5 := make([]SrcIntelSigParam, 0, 0)
										for (err == nil) && j.More() {
											var v6 SrcIntelSigParam
											err = v6.__gent__jsonUnmarshal_Decode(j)
											if err == nil {
												sl5 = append(sl5, v6)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.Parameters = sl5
										}
									}
								default:
									err = pkg__errors.New("expected [")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigInfo) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcIntelSigParam) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "label":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Label = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "documentation":
							err = me.Documentation.__gent__jsonUnmarshal_Decode(j)
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigParam) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcLenses) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t4 pkg__encoding_json.Token
	t4, err = j.Token()
	if (err == nil) && (t4 != nil) {
		switch d5 := t4.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d5 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]*SrcLens, 0, 0)
				for (err == nil) && j.More() {
					var v2 *SrcLens
					var pv3 SrcLens
					err = pv3.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						v2 = &pv3
					}
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLenses) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcLens) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "SrcLoc":
							err = me.SrcLoc.__gent__jsonUnmarshal_Decode(j)
						case "t":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case string:
									me.Txt = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "s":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.Str = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "l":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLens) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcLoc) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "e":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.Flag = (int)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "f":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case string:
									me.FilePath = v
								default:
									err = pkg__errors.New("expected string")
								}
							}
						case "p":
							var pv7 SrcPos
							err = pv7.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Pos = &pv7
							}
						case "r":
							var pv8 SrcRange
							err = pv8.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.Range = &pv8
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLoc) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcLocs) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t4 pkg__encoding_json.Token
	t4, err = j.Token()
	if (err == nil) && (t4 != nil) {
		switch d5 := t4.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d5 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]*SrcLoc, 0, 0)
				for (err == nil) && j.More() {
					var v2 *SrcLoc
					var pv3 SrcLoc
					err = pv3.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						v2 = &pv3
					}
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLocs) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcModEdit) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t6 pkg__encoding_json.Token
	t6, err = j.Token()
	if (err == nil) && (t6 != nil) {
		switch d7 := t6.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d7 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "At":
							var pv3 SrcRange
							err = pv3.__gent__jsonUnmarshal_Decode(j)
							if err == nil {
								me.At = &pv3
							}
						case "Val":
							var jt4 pkg__encoding_json.Token
							jt4, err = j.Token()
							if (err == nil) && (jt4 != nil) {
								switch v := jt4.(type) {
								case nil:
								case string:
									me.Val = v
								default:
									err = pkg__errors.New("expected string")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdit) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcModEdits) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t3 pkg__encoding_json.Token
	t3, err = j.Token()
	if (err == nil) && (t3 != nil) {
		switch d4 := t3.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x5b != d4 {
				err = pkg__errors.New("expected [")
			} else {
				sl1 := make([]SrcModEdit, 0, 0)
				for (err == nil) && j.More() {
					var v2 SrcModEdit
					err = v2.__gent__jsonUnmarshal_Decode(j)
					if err == nil {
						sl1 = append(sl1, v2)
					}
				}
				if err == nil {
					_, err = j.Token()
				}
				if err == nil {
					*me = sl1
				}
			}
		default:
			err = pkg__errors.New("expected [")
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdits) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

func (me *SrcPos) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
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
						case "l":
							var jt3 pkg__encoding_json.Token
							jt3, err = j.Token()
							if (err == nil) && (jt3 != nil) {
								switch v := jt3.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v4 int64
									v4, err = v.Int64()
									if err == nil {
										me.Ln = (int)(v4)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "c":
							var jt5 pkg__encoding_json.Token
							jt5, err = j.Token()
							if (err == nil) && (jt5 != nil) {
								switch v := jt5.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v6 int64
									v6, err = v.Int64()
									if err == nil {
										me.Col = (int)(v6)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
								}
							}
						case "o":
							var jt7 pkg__encoding_json.Token
							jt7, err = j.Token()
							if (err == nil) && (jt7 != nil) {
								switch v := jt7.(type) {
								case nil:
								case pkg__encoding_json.Number:
									var v8 int64
									v8, err = v.Int64()
									if err == nil {
										me.Off = (int)(v8)
									}
								default:
									err = pkg__errors.New("expected pkg__encoding_json.Number")
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcPos) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
	return
}

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) preview_MarshalJSON() (r []byte, err error) {
	panic("WorkspaceChanges")
	return
}

func (me *WorkspaceChanges) __gent__jsonUnmarshal_Decode(j *pkg__encoding_json.Decoder) (err error) {
	var t41 pkg__encoding_json.Token
	t41, err = j.Token()
	if (err == nil) && (t41 != nil) {
		switch d42 := t41.(type) {
		case nil:
		case pkg__encoding_json.Delim:
			if 0x7b != d42 {
				err = pkg__errors.New("expected {")
			} else {
				for (err == nil) && j.More() {
					var jk1 pkg__encoding_json.Token
					jk1, err = j.Token()
					if err == nil {
						fn2 := jk1.(string)
						switch fn2 {
						case "AddedDirs":
							var t7 pkg__encoding_json.Token
							t7, err = j.Token()
							if (err == nil) && (t7 != nil) {
								switch d8 := t7.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d8 {
										err = pkg__errors.New("expected [")
									} else {
										sl3 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v4 string
											var jt5 pkg__encoding_json.Token
											jt5, err = j.Token()
											if (err == nil) && (jt5 != nil) {
												switch v := jt5.(type) {
												case nil:
												case string:
													v4 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl3 = append(sl3, v4)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.AddedDirs = sl3
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "RemovedDirs":
							var t13 pkg__encoding_json.Token
							t13, err = j.Token()
							if (err == nil) && (t13 != nil) {
								switch d14 := t13.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d14 {
										err = pkg__errors.New("expected [")
									} else {
										sl9 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v10 string
											var jt11 pkg__encoding_json.Token
											jt11, err = j.Token()
											if (err == nil) && (jt11 != nil) {
												switch v := jt11.(type) {
												case nil:
												case string:
													v10 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl9 = append(sl9, v10)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.RemovedDirs = sl9
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "OpenedFiles":
							var t19 pkg__encoding_json.Token
							t19, err = j.Token()
							if (err == nil) && (t19 != nil) {
								switch d20 := t19.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d20 {
										err = pkg__errors.New("expected [")
									} else {
										sl15 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v16 string
											var jt17 pkg__encoding_json.Token
											jt17, err = j.Token()
											if (err == nil) && (jt17 != nil) {
												switch v := jt17.(type) {
												case nil:
												case string:
													v16 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl15 = append(sl15, v16)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.OpenedFiles = sl15
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "ClosedFiles":
							var t25 pkg__encoding_json.Token
							t25, err = j.Token()
							if (err == nil) && (t25 != nil) {
								switch d26 := t25.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d26 {
										err = pkg__errors.New("expected [")
									} else {
										sl21 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v22 string
											var jt23 pkg__encoding_json.Token
											jt23, err = j.Token()
											if (err == nil) && (jt23 != nil) {
												switch v := jt23.(type) {
												case nil:
												case string:
													v22 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl21 = append(sl21, v22)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.ClosedFiles = sl21
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "WrittenFiles":
							var t31 pkg__encoding_json.Token
							t31, err = j.Token()
							if (err == nil) && (t31 != nil) {
								switch d32 := t31.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x5b != d32 {
										err = pkg__errors.New("expected [")
									} else {
										sl27 := make([]string, 0, 0)
										for (err == nil) && j.More() {
											var v28 string
											var jt29 pkg__encoding_json.Token
											jt29, err = j.Token()
											if (err == nil) && (jt29 != nil) {
												switch v := jt29.(type) {
												case nil:
												case string:
													v28 = v
												default:
													err = pkg__errors.New("expected string")
												}
											}
											if err == nil {
												sl27 = append(sl27, v28)
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.WrittenFiles = sl27
										}
									}
								default:
									err = pkg__errors.New("expected [")
								}
							}
						case "LiveFiles":
							var t39 pkg__encoding_json.Token
							t39, err = j.Token()
							if (err == nil) && (t39 != nil) {
								switch d40 := t39.(type) {
								case nil:
								case pkg__encoding_json.Delim:
									if 0x7b != d40 {
										err = pkg__errors.New("expected {")
									} else {
										t36 := make(map[string]string, 0)
										for (err == nil) && j.More() {
											var jk33 pkg__encoding_json.Token
											jk33, err = j.Token()
											if err == nil {
												mk34 := jk33.(string)
												var mv35 string
												var jt37 pkg__encoding_json.Token
												jt37, err = j.Token()
												if (err == nil) && (jt37 != nil) {
													switch v := jt37.(type) {
													case nil:
													case string:
														mv35 = v
													default:
														err = pkg__errors.New("expected string")
													}
												}
												if err == nil {
													t36[mk34] = mv35
												}
											}
										}
										if err == nil {
											_, err = j.Token()
										}
										if err == nil {
											me.LiveFiles = t36
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

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *WorkspaceChanges) preview_UnmarshalJSON(b []byte) (err error) {
	j := pkg__encoding_json.NewDecoder(pkg__bytes.NewReader(b))
	j.UseNumber()
	err = me.__gent__jsonUnmarshal_Decode(j)
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
