package z

// DON'T EDIT: code gen'd with `zentient-codegen` using `github.com/metaleap/go-gent`

import (
	pkg__encoding_json "encoding/json"
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
			r[si1] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelCompl) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntels) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si2 := len(r)
		if len(me.InfoTips) != 0 {
			r = append(r, ",\"InfoTips\":"...)
			r = append(r, 91)
			ai4 := len(r)
			for i3 := range me.InfoTips {
				r = append(r, 44)
				r = append(r, 123)
				si5 := len(r)
				{
					r = append(r, ",\"value\":"...)
					r = append(r, pkg__strconv.Quote(me.InfoTips[i3].Value)...)
				}
				if len(me.InfoTips[i3].Language) != 0 {
					r = append(r, ",\"language\":"...)
					r = append(r, pkg__strconv.Quote(me.InfoTips[i3].Language)...)
				}
				r = append(r, 125)
				if r[si5] == 44 {
					r[si5] = 32
				}
			}
			r = append(r, 93)
			if r[ai4] == 44 {
				r[ai4] = 32
			}
		}
		if len(me.Refs) != 0 {
			r = append(r, ",\"Refs\":"...)
			r = append(r, 91)
			ai7 := len(r)
			for i6 := range me.Refs {
				if nil != me.Refs[i6] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.Refs[i6].preview_MarshalJSON()
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
			if r[ai7] == 44 {
				r[ai7] = 32
			}
		}
		r = append(r, 125)
		if r[si2] == 44 {
			r[si2] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntels) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelDoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si8 := len(r)
		if len(me.Value) != 0 {
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if me.IsTrusted {
			r = append(r, ",\"isTrusted\":"...)
			r = append(r, pkg__strconv.FormatBool(me.IsTrusted)...)
		}
		r = append(r, 125)
		if r[si8] == 44 {
			r[si8] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelDoc) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigHelp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si9 := len(r)
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
			ai11 := len(r)
			for i10 := range me.Signatures {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Signatures[i10].preview_MarshalJSON()
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
				r[ai11] = 32
			}
		}
		r = append(r, 125)
		if r[si9] == 44 {
			r[si9] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigHelp) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigInfo) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si12 := len(r)
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
		{
			r = append(r, ",\"parameters\":"...)
			r = append(r, 91)
			ai14 := len(r)
			for i13 := range me.Parameters {
				{
					r = append(r, 44)
					var e error
					var sl []byte
					sl, e = me.Parameters[i13].preview_MarshalJSON()
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
				r[ai14] = 32
			}
		}
		r = append(r, 125)
		if r[si12] == 44 {
			r[si12] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigInfo) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcIntelSigParam) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si15 := len(r)
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
		if r[si15] == 44 {
			r[si15] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcIntelSigParam) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLens) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si16 := len(r)
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
		if r[si16] == 44 {
			r[si16] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLens) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcLoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si17 := len(r)
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
		if r[si17] == 44 {
			r[si17] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcLoc) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcModEdit) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si18 := len(r)
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
		if r[si18] == 44 {
			r[si18] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcModEdit) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcPos) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si19 := len(r)
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
		if r[si19] == 44 {
			r[si19] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcPos) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *SrcRange) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si20 := len(r)
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
		if r[si20] == 44 {
			r[si20] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *SrcRange) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *WorkspaceChanges) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si21 := len(r)
		{
			r = append(r, ",\"AddedDirs\":"...)
			r = append(r, 91)
			ai23 := len(r)
			for i22 := range me.AddedDirs {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.AddedDirs[i22])...)
				}
			}
			r = append(r, 93)
			if r[ai23] == 44 {
				r[ai23] = 32
			}
		}
		{
			r = append(r, ",\"RemovedDirs\":"...)
			r = append(r, 91)
			ai25 := len(r)
			for i24 := range me.RemovedDirs {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.RemovedDirs[i24])...)
				}
			}
			r = append(r, 93)
			if r[ai25] == 44 {
				r[ai25] = 32
			}
		}
		{
			r = append(r, ",\"OpenedFiles\":"...)
			r = append(r, 91)
			ai27 := len(r)
			for i26 := range me.OpenedFiles {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.OpenedFiles[i26])...)
				}
			}
			r = append(r, 93)
			if r[ai27] == 44 {
				r[ai27] = 32
			}
		}
		{
			r = append(r, ",\"ClosedFiles\":"...)
			r = append(r, 91)
			ai29 := len(r)
			for i28 := range me.ClosedFiles {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.ClosedFiles[i28])...)
				}
			}
			r = append(r, 93)
			if r[ai29] == 44 {
				r[ai29] = 32
			}
		}
		{
			r = append(r, ",\"WrittenFiles\":"...)
			r = append(r, 91)
			ai31 := len(r)
			for i30 := range me.WrittenFiles {
				{
					r = append(r, 44)
					r = append(r, pkg__strconv.Quote(me.WrittenFiles[i30])...)
				}
			}
			r = append(r, 93)
			if r[ai31] == 44 {
				r[ai31] = 32
			}
		}
		{
			r = append(r, ",\"LiveFiles\":"...)
			r = append(r, 123)
			mi32 := len(r)
			for mk33, mv34 := range me.LiveFiles {
				{
					{
						r = append(r, 44)
						r = append(r, pkg__strconv.Quote(mk33)...)
						r = append(r, 58)
					}
					r = append(r, pkg__strconv.Quote(mv34)...)
				}
			}
			r = append(r, 125)
			if r[mi32] == 44 {
				r[mi32] = 32
			}
		}
		r = append(r, 125)
		if r[si21] == 44 {
			r[si21] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *WorkspaceChanges) preview_UnmarshalJSON(v []byte) (err error) { return }

// preview_MarshalJSON implements the Go standard library's `encoding/json.Marshaler` interface.
func (me *fooResp) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		si35 := len(r)
		{
			r = append(r, ",\"ii\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcID), 10)...)
		}
		if me.ReqID != 0 {
			r = append(r, ",\"ri\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.ReqID), 10)...)
		}
		{
			r = append(r, ",\"Flag\":"...)
			r = append(r, pkg__strconv.FormatBool(me.Flag)...)
		}
		if len(me.ErrMsg) != 0 {
			r = append(r, ",\"err\":"...)
			r = append(r, pkg__strconv.Quote(me.ErrMsg)...)
		}
		if nil != me.SrcIntel {
			r = append(r, ",\"sI\":"...)
			r = append(r, 123)
			si36 := len(r)
			if len(me.SrcIntel.SrcIntels.InfoTips) != 0 {
				r = append(r, ",\"InfoTips\":"...)
				r = append(r, 91)
				ai38 := len(r)
				for i37 := range me.SrcIntel.SrcIntels.InfoTips {
					r = append(r, 44)
					r = append(r, 123)
					si39 := len(r)
					{
						r = append(r, ",\"value\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.SrcIntels.InfoTips[i37].Value)...)
					}
					if len(me.SrcIntel.SrcIntels.InfoTips[i37].Language) != 0 {
						r = append(r, ",\"language\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.SrcIntels.InfoTips[i37].Language)...)
					}
					r = append(r, 125)
					if r[si39] == 44 {
						r[si39] = 32
					}
				}
				r = append(r, 93)
				if r[ai38] == 44 {
					r[ai38] = 32
				}
			}
			if len(me.SrcIntel.SrcIntels.Refs) != 0 {
				r = append(r, ",\"Refs\":"...)
				r = append(r, 91)
				ai41 := len(r)
				for i40 := range me.SrcIntel.SrcIntels.Refs {
					if nil != me.SrcIntel.SrcIntels.Refs[i40] {
						{
							r = append(r, 44)
							var e error
							var sl []byte
							sl, e = me.SrcIntel.SrcIntels.Refs[i40].preview_MarshalJSON()
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
				if r[ai41] == 44 {
					r[ai41] = 32
				}
			}
			if nil != me.SrcIntel.Sig {
				{
					r = append(r, ",\"Sig\":"...)
					var e error
					var sl []byte
					sl, e = me.SrcIntel.Sig.preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			if len(me.SrcIntel.Cmpl) != 0 {
				r = append(r, ",\"Cmpl\":"...)
				r = append(r, 91)
				ai43 := len(r)
				for i42 := range me.SrcIntel.Cmpl {
					if nil != me.SrcIntel.Cmpl[i42] {
						{
							r = append(r, 44)
							var e error
							var sl []byte
							sl, e = me.SrcIntel.Cmpl[i42].preview_MarshalJSON()
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
				if r[ai43] == 44 {
					r[ai43] = 32
				}
			}
			if len(me.SrcIntel.Syms) != 0 {
				r = append(r, ",\"Syms\":"...)
				r = append(r, 91)
				ai45 := len(r)
				for i44 := range me.SrcIntel.Syms {
					if nil != me.SrcIntel.Syms[i44] {
						{
							r = append(r, 44)
							var e error
							var sl []byte
							sl, e = me.SrcIntel.Syms[i44].preview_MarshalJSON()
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
					r[ai45] = 32
				}
			}
			if len(me.SrcIntel.Anns) != 0 {
				r = append(r, ",\"Anns\":"...)
				r = append(r, 91)
				ai47 := len(r)
				for i46 := range me.SrcIntel.Anns {
					if nil != me.SrcIntel.Anns[i46] {
						r = append(r, 44)
						r = append(r, 123)
						si48 := len(r)
						{
							r = append(r, ",\"Range\":"...)
							var e error
							var sl []byte
							sl, e = me.SrcIntel.Anns[i46].Range.preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
						{
							r = append(r, ",\"Title\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i46].Title)...)
						}
						if len(me.SrcIntel.Anns[i46].Desc) != 0 {
							r = append(r, ",\"Desc\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i46].Desc)...)
						}
						{
							r = append(r, ",\"CmdName\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i46].CmdName)...)
						}
						r = append(r, 125)
						if r[si48] == 44 {
							r[si48] = 32
						}
					} else {
						r = append(r, 44)
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
				if r[ai47] == 44 {
					r[ai47] = 32
				}
			}
			r = append(r, 125)
			if r[si36] == 44 {
				r[si36] = 32
			}
		}
		if nil != me.SrcDiags {
			r = append(r, ",\"srcDiags\":"...)
			r = append(r, 123)
			si49 := len(r)
			{
				r = append(r, ",\"All\":"...)
				r = append(r, 123)
				mi50 := len(r)
				for mk51, mv52 := range me.SrcDiags.All {
					{
						{
							r = append(r, 44)
							r = append(r, pkg__strconv.Quote(mk51)...)
							r = append(r, 58)
						}
						r = append(r, 91)
						ai54 := len(r)
						for i53 := range mv52 {
							if nil != mv52[i53] {
								r = append(r, 44)
								r = append(r, 123)
								si55 := len(r)
								if len(mv52[i53].Cat) != 0 {
									r = append(r, ",\"Cat\":"...)
									r = append(r, pkg__strconv.Quote(mv52[i53].Cat)...)
								}
								{
									r = append(r, ",\"Loc\":"...)
									var e error
									var sl []byte
									sl, e = mv52[i53].Loc.preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
								{
									r = append(r, ",\"Msg\":"...)
									r = append(r, pkg__strconv.Quote(mv52[i53].Msg)...)
								}
								if len(mv52[i53].SrcActions) != 0 {
									r = append(r, ",\"SrcActions\":"...)
									r = append(r, 91)
									ai57 := len(r)
									for i56 := range mv52[i53].SrcActions {
										r = append(r, 44)
										r = append(r, 123)
										si58 := len(r)
										{
											r = append(r, ",\"title\":"...)
											r = append(r, pkg__strconv.Quote(mv52[i53].SrcActions[i56].Title)...)
										}
										{
											r = append(r, ",\"command\":"...)
											r = append(r, pkg__strconv.Quote(mv52[i53].SrcActions[i56].Cmd)...)
										}
										if len(mv52[i53].SrcActions[i56].Hint) != 0 {
											r = append(r, ",\"tooltip\":"...)
											r = append(r, pkg__strconv.Quote(mv52[i53].SrcActions[i56].Hint)...)
										}
										if len(mv52[i53].SrcActions[i56].Arguments) != 0 {
											r = append(r, ",\"arguments\":"...)
											r = append(r, 91)
											ai60 := len(r)
											for i59 := range mv52[i53].SrcActions[i56].Arguments {
												if mv52[i53].SrcActions[i56].Arguments[i59] != nil {
													r = append(r, 44)
													var e error
													var sl []byte
													j, ok := mv52[i53].SrcActions[i56].Arguments[i59].(pkg__encoding_json.Marshaler)
													if ok && (j != nil) {
														sl, e = j.MarshalJSON()
													} else {
														sl, e = pkg__encoding_json.Marshal(mv52[i53].SrcActions[i56].Arguments[i59])
													}
													if e == nil {
														r = append(r, sl...)
													} else {
														err = e
														return
													}
												} else {
													r = append(r, 44)
													r = append(r, "null"...)
												}
											}
											r = append(r, 93)
											if r[ai60] == 44 {
												r[ai60] = 32
											}
										}
										r = append(r, 125)
										if r[si58] == 44 {
											r[si58] = 32
										}
									}
									r = append(r, 93)
									if r[ai57] == 44 {
										r[ai57] = 32
									}
								}
								if mv52[i53].StickyAuto {
									r = append(r, ",\"Sticky\":"...)
									r = append(r, pkg__strconv.FormatBool(mv52[i53].StickyAuto)...)
								}
								if len(mv52[i53].Tags) != 0 {
									r = append(r, ",\"Tags\":"...)
									r = append(r, 91)
									ai62 := len(r)
									for i61 := range mv52[i53].Tags {
										{
											r = append(r, 44)
											r = append(r, pkg__strconv.FormatInt((int64)(mv52[i53].Tags[i61]), 10)...)
										}
									}
									r = append(r, 93)
									if r[ai62] == 44 {
										r[ai62] = 32
									}
								}
								r = append(r, 125)
								if r[si55] == 44 {
									r[si55] = 32
								}
							} else {
								r = append(r, 44)
								r = append(r, "null"...)
							}
						}
						r = append(r, 93)
						if r[ai54] == 44 {
							r[ai54] = 32
						}
					}
				}
				r = append(r, 125)
				if r[mi50] == 44 {
					r[mi50] = 32
				}
			}
			{
				r = append(r, ",\"FixUps\":"...)
				r = append(r, 91)
				ai64 := len(r)
				for i63 := range me.SrcDiags.FixUps {
					if nil != me.SrcDiags.FixUps[i63] {
						r = append(r, 44)
						r = append(r, 123)
						si65 := len(r)
						{
							r = append(r, ",\"FilePath\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcDiags.FixUps[i63].FilePath)...)
						}
						{
							r = append(r, ",\"Desc\":"...)
							r = append(r, 123)
							mi66 := len(r)
							for mk67, mv68 := range me.SrcDiags.FixUps[i63].Desc {
								{
									{
										r = append(r, 44)
										r = append(r, pkg__strconv.Quote(mk67)...)
										r = append(r, 58)
									}
									r = append(r, 91)
									ai70 := len(r)
									for i69 := range mv68 {
										{
											r = append(r, 44)
											r = append(r, pkg__strconv.Quote(mv68[i69])...)
										}
									}
									r = append(r, 93)
									if r[ai70] == 44 {
										r[ai70] = 32
									}
								}
							}
							r = append(r, 125)
							if r[mi66] == 44 {
								r[mi66] = 32
							}
						}
						{
							r = append(r, ",\"Edits\":"...)
							r = append(r, 91)
							ai72 := len(r)
							for i71 := range me.SrcDiags.FixUps[i63].Edits {
								{
									r = append(r, 44)
									var e error
									var sl []byte
									sl, e = me.SrcDiags.FixUps[i63].Edits[i71].preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
							}
							r = append(r, 93)
							if r[ai72] == 44 {
								r[ai72] = 32
							}
						}
						{
							r = append(r, ",\"Dropped\":"...)
							r = append(r, 91)
							ai74 := len(r)
							for i73 := range me.SrcDiags.FixUps[i63].Dropped {
								{
									r = append(r, 44)
									var e error
									var sl []byte
									sl, e = me.SrcDiags.FixUps[i63].Dropped[i73].preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
							}
							r = append(r, 93)
							if r[ai74] == 44 {
								r[ai74] = 32
							}
						}
						r = append(r, 125)
						if r[si65] == 44 {
							r[si65] = 32
						}
					} else {
						r = append(r, 44)
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
				if r[ai64] == 44 {
					r[ai64] = 32
				}
			}
			{
				r = append(r, ",\"LangID\":"...)
				r = append(r, pkg__strconv.Quote(me.SrcDiags.LangID)...)
			}
			r = append(r, 125)
			if r[si49] == 44 {
				r[si49] = 32
			}
		}
		{
			r = append(r, ",\"ri\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcReq.ReqID), 10)...)
		}
		{
			r = append(r, ",\"ii\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.IpcReq.IpcID), 10)...)
		}
		if me.IpcReq.IpcArgs != nil {
			r = append(r, ",\"ia\":"...)
			var e error
			var sl []byte
			j, ok := me.IpcReq.IpcArgs.(pkg__encoding_json.Marshaler)
			if ok && (j != nil) {
				sl, e = j.MarshalJSON()
			} else {
				sl, e = pkg__encoding_json.Marshal(me.IpcReq.IpcArgs)
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
		if nil != me.IpcReq.ProjUpd {
			{
				r = append(r, ",\"projUpd\":"...)
				var e error
				var sl []byte
				sl, e = me.IpcReq.ProjUpd.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		} else {
			r = append(r, ",\"projUpd\":"...)
			r = append(r, "null"...)
		}
		if nil != me.IpcReq.SrcLens {
			{
				r = append(r, ",\"srcLens\":"...)
				var e error
				var sl []byte
				sl, e = me.IpcReq.SrcLens.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		} else {
			r = append(r, ",\"srcLens\":"...)
			r = append(r, "null"...)
		}
		if len(me.SrcMods) != 0 {
			r = append(r, ",\"srcMods\":"...)
			r = append(r, 91)
			ai76 := len(r)
			for i75 := range me.SrcMods {
				if nil != me.SrcMods[i75] {
					{
						r = append(r, 44)
						var e error
						var sl []byte
						sl, e = me.SrcMods[i75].preview_MarshalJSON()
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
			if r[ai76] == 44 {
				r[ai76] = 32
			}
		}
		if len(me.SrcActions) != 0 {
			r = append(r, ",\"srcActions\":"...)
			r = append(r, 91)
			ai78 := len(r)
			for i77 := range me.SrcActions {
				r = append(r, 44)
				r = append(r, 123)
				si79 := len(r)
				{
					r = append(r, ",\"title\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i77].Title)...)
				}
				{
					r = append(r, ",\"command\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i77].Cmd)...)
				}
				if len(me.SrcActions[i77].Hint) != 0 {
					r = append(r, ",\"tooltip\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i77].Hint)...)
				}
				if len(me.SrcActions[i77].Arguments) != 0 {
					r = append(r, ",\"arguments\":"...)
					r = append(r, 91)
					ai81 := len(r)
					for i80 := range me.SrcActions[i77].Arguments {
						if me.SrcActions[i77].Arguments[i80] != nil {
							r = append(r, 44)
							var e error
							var sl []byte
							j, ok := me.SrcActions[i77].Arguments[i80].(pkg__encoding_json.Marshaler)
							if ok && (j != nil) {
								sl, e = j.MarshalJSON()
							} else {
								sl, e = pkg__encoding_json.Marshal(me.SrcActions[i77].Arguments[i80])
							}
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						} else {
							r = append(r, 44)
							r = append(r, "null"...)
						}
					}
					r = append(r, 93)
					if r[ai81] == 44 {
						r[ai81] = 32
					}
				}
				r = append(r, 125)
				if r[si79] == 44 {
					r[si79] = 32
				}
			}
			r = append(r, 93)
			if r[ai78] == 44 {
				r[ai78] = 32
			}
		}
		if nil != me.Extras {
			r = append(r, ",\"extras\":"...)
			r = append(r, 123)
			si82 := len(r)
			if len(me.Extras.SrcIntels.InfoTips) != 0 {
				r = append(r, ",\"InfoTips\":"...)
				r = append(r, 91)
				ai84 := len(r)
				for i83 := range me.Extras.SrcIntels.InfoTips {
					r = append(r, 44)
					r = append(r, 123)
					si85 := len(r)
					{
						r = append(r, ",\"value\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.SrcIntels.InfoTips[i83].Value)...)
					}
					if len(me.Extras.SrcIntels.InfoTips[i83].Language) != 0 {
						r = append(r, ",\"language\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.SrcIntels.InfoTips[i83].Language)...)
					}
					r = append(r, 125)
					if r[si85] == 44 {
						r[si85] = 32
					}
				}
				r = append(r, 93)
				if r[ai84] == 44 {
					r[ai84] = 32
				}
			}
			if len(me.Extras.SrcIntels.Refs) != 0 {
				r = append(r, ",\"Refs\":"...)
				r = append(r, 91)
				ai87 := len(r)
				for i86 := range me.Extras.SrcIntels.Refs {
					if nil != me.Extras.SrcIntels.Refs[i86] {
						{
							r = append(r, 44)
							var e error
							var sl []byte
							sl, e = me.Extras.SrcIntels.Refs[i86].preview_MarshalJSON()
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
				if r[ai87] == 44 {
					r[ai87] = 32
				}
			}
			{
				r = append(r, ",\"Items\":"...)
				r = append(r, 91)
				ai89 := len(r)
				for i88 := range me.Extras.Items {
					if nil != me.Extras.Items[i88] {
						r = append(r, 44)
						r = append(r, 123)
						si90 := len(r)
						{
							r = append(r, ",\"id\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].ID)...)
						}
						{
							r = append(r, ",\"label\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].Label)...)
						}
						{
							r = append(r, ",\"description\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].Desc)...)
						}
						if len(me.Extras.Items[i88].Detail) != 0 {
							r = append(r, ",\"detail\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].Detail)...)
						}
						if len(me.Extras.Items[i88].QueryArg) != 0 {
							r = append(r, ",\"arg\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].QueryArg)...)
						}
						if len(me.Extras.Items[i88].FilePos) != 0 {
							r = append(r, ",\"fPos\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i88].FilePos)...)
						}
						r = append(r, 125)
						if r[si90] == 44 {
							r[si90] = 32
						}
					} else {
						r = append(r, 44)
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
				if r[ai89] == 44 {
					r[ai89] = 32
				}
			}
			if len(me.Extras.Warns) != 0 {
				r = append(r, ",\"Warns\":"...)
				r = append(r, 91)
				ai92 := len(r)
				for i91 := range me.Extras.Warns {
					{
						r = append(r, 44)
						r = append(r, pkg__strconv.Quote(me.Extras.Warns[i91])...)
					}
				}
				r = append(r, 93)
				if r[ai92] == 44 {
					r[ai92] = 32
				}
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
			if r[si82] == 44 {
				r[si82] = 32
			}
		}
		{
			r = append(r, ",\"e\":"...)
			r = append(r, pkg__strconv.FormatInt((int64)(me.SrcLens.SrcLoc.Flag), 10)...)
		}
		if len(me.SrcLens.SrcLoc.FilePath) != 0 {
			r = append(r, ",\"f\":"...)
			r = append(r, pkg__strconv.Quote(me.SrcLens.SrcLoc.FilePath)...)
		}
		if nil != me.SrcLens.SrcLoc.Pos {
			{
				r = append(r, ",\"p\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcLens.SrcLoc.Pos.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
		}
		if nil != me.SrcLens.SrcLoc.Range {
			{
				r = append(r, ",\"r\":"...)
				var e error
				var sl []byte
				sl, e = me.SrcLens.SrcLoc.Range.preview_MarshalJSON()
				if e == nil {
					r = append(r, sl...)
				} else {
					err = e
					return
				}
			}
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
		if nil != me.Link {
			{
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
			}
		} else {
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
			si93 := len(r)
			if nil != me.Menu.SubMenu {
				r = append(r, ",\"SubMenu\":"...)
				r = append(r, 123)
				si94 := len(r)
				if len(me.Menu.SubMenu.Desc) != 0 {
					r = append(r, ",\"desc\":"...)
					r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Desc)...)
				}
				if me.Menu.SubMenu.TopLevel {
					r = append(r, ",\"topLevel\":"...)
					r = append(r, pkg__strconv.FormatBool(me.Menu.SubMenu.TopLevel)...)
				}
				{
					r = append(r, ",\"items\":"...)
					r = append(r, 91)
					ai96 := len(r)
					for i95 := range me.Menu.SubMenu.Items {
						if nil != me.Menu.SubMenu.Items[i95] {
							r = append(r, 44)
							r = append(r, 123)
							si97 := len(r)
							if me.Menu.SubMenu.Items[i95].IpcID != 0 {
								r = append(r, ",\"ii\":"...)
								r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.SubMenu.Items[i95].IpcID), 10)...)
							}
							if me.Menu.SubMenu.Items[i95].IpcArgs != nil {
								r = append(r, ",\"ia\":"...)
								var e error
								var sl []byte
								j, ok := me.Menu.SubMenu.Items[i95].IpcArgs.(pkg__encoding_json.Marshaler)
								if ok && (j != nil) {
									sl, e = j.MarshalJSON()
								} else {
									sl, e = pkg__encoding_json.Marshal(me.Menu.SubMenu.Items[i95].IpcArgs)
								}
								if e == nil {
									r = append(r, sl...)
								} else {
									err = e
									return
								}
							}
							if len(me.Menu.SubMenu.Items[i95].Category) != 0 {
								r = append(r, ",\"c\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i95].Category)...)
							}
							{
								r = append(r, ",\"t\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i95].Title)...)
							}
							if len(me.Menu.SubMenu.Items[i95].Desc) != 0 {
								r = append(r, ",\"d\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i95].Desc)...)
							}
							if len(me.Menu.SubMenu.Items[i95].Hint) != 0 {
								r = append(r, ",\"h\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i95].Hint)...)
							}
							if len(me.Menu.SubMenu.Items[i95].Confirm) != 0 {
								r = append(r, ",\"q\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i95].Confirm)...)
							}
							r = append(r, 125)
							if r[si97] == 44 {
								r[si97] = 32
							}
						} else {
							r = append(r, 44)
							r = append(r, "null"...)
						}
					}
					r = append(r, 93)
					if r[ai96] == 44 {
						r[ai96] = 32
					}
				}
				r = append(r, 125)
				if r[si94] == 44 {
					r[si94] = 32
				}
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
				ai99 := len(r)
				for i98 := range me.Menu.Refs {
					if nil != me.Menu.Refs[i98] {
						{
							r = append(r, 44)
							var e error
							var sl []byte
							sl, e = me.Menu.Refs[i98].preview_MarshalJSON()
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
				if r[ai99] == 44 {
					r[ai99] = 32
				}
			}
			r = append(r, 125)
			if r[si93] == 44 {
				r[si93] = 32
			}
		}
		if nil != me.CaddyUpdate {
			r = append(r, ",\"caddy\":"...)
			r = append(r, 123)
			si100 := len(r)
			if len(me.CaddyUpdate.ID) != 0 {
				r = append(r, ",\"ID\":"...)
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
			si101 := len(r)
			{
				r = append(r, ",\"Flag\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.CaddyUpdate.Status.Flag), 10)...)
			}
			if len(me.CaddyUpdate.Status.Desc) != 0 {
				r = append(r, ",\"Desc\":"...)
				r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Status.Desc)...)
			}
			r = append(r, 125)
			if r[si101] == 44 {
				r[si101] = 32
			}
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
			if r[si100] == 44 {
				r[si100] = 32
			}
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
		if r[si35] == 44 {
			r[si35] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *fooResp) preview_UnmarshalJSON(v []byte) (err error) { return }
