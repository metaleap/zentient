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
func (me *SrcIntelDoc) preview_MarshalJSON() (r []byte, err error) {
	r = make([]byte, 0, 64)
	if me == nil {
		r = append(r, "null"...)
	} else {
		r = append(r, 123)
		idx1 := len(r)
		if len(me.Value) != 0 {
			r = append(r, ",\"value\":"...)
			r = append(r, pkg__strconv.Quote(me.Value)...)
		}
		if me.IsTrusted {
			r = append(r, ",\"isTrusted\":"...)
			r = append(r, pkg__strconv.FormatBool(me.IsTrusted)...)
		}
		r = append(r, 125)
		if r[idx1] == 44 {
			r[idx1] = 32
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
		idx2 := len(r)
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
			for i3 := range me.Signatures {
				{
					if i3 != 0 {
						r = append(r, 44)
					}
					var e error
					var sl []byte
					sl, e = me.Signatures[i3].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
		}
		r = append(r, 125)
		if r[idx2] == 44 {
			r[idx2] = 32
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
		idx4 := len(r)
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
			for i5 := range me.Parameters {
				{
					if i5 != 0 {
						r = append(r, 44)
					}
					var e error
					var sl []byte
					sl, e = me.Parameters[i5].preview_MarshalJSON()
					if e == nil {
						r = append(r, sl...)
					} else {
						err = e
						return
					}
				}
			}
			r = append(r, 93)
		}
		r = append(r, 125)
		if r[idx4] == 44 {
			r[idx4] = 32
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
		idx6 := len(r)
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
		if r[idx6] == 44 {
			r[idx6] = 32
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
		idx7 := len(r)
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
		if r[idx7] == 44 {
			r[idx7] = 32
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
		idx8 := len(r)
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
		if r[idx8] == 44 {
			r[idx8] = 32
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
		idx9 := len(r)
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
		if r[idx9] == 44 {
			r[idx9] = 32
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
		idx10 := len(r)
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
		if r[idx10] == 44 {
			r[idx10] = 32
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
		idx11 := len(r)
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
		if r[idx11] == 44 {
			r[idx11] = 32
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
		idx12 := len(r)
		{
			r = append(r, ",\"AddedDirs\":"...)
			r = append(r, 91)
			for i13 := range me.AddedDirs {
				{
					if i13 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.AddedDirs[i13])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"RemovedDirs\":"...)
			r = append(r, 91)
			for i14 := range me.RemovedDirs {
				{
					if i14 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.RemovedDirs[i14])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"OpenedFiles\":"...)
			r = append(r, 91)
			for i15 := range me.OpenedFiles {
				{
					if i15 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.OpenedFiles[i15])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"ClosedFiles\":"...)
			r = append(r, 91)
			for i16 := range me.ClosedFiles {
				{
					if i16 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.ClosedFiles[i16])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"WrittenFiles\":"...)
			r = append(r, 91)
			for i17 := range me.WrittenFiles {
				{
					if i17 != 0 {
						r = append(r, 44)
					}
					r = append(r, pkg__strconv.Quote(me.WrittenFiles[i17])...)
				}
			}
			r = append(r, 93)
		}
		{
			r = append(r, ",\"LiveFiles\":"...)
			mf18 := true
			r = append(r, 123)
			for mk19, mv20 := range me.LiveFiles {
				{
					{
						if mf18 {
							mf18 = false
						} else {
							r = append(r, 44)
						}
						r = append(r, pkg__strconv.Quote(mk19)...)
						r = append(r, 58)
					}
					r = append(r, pkg__strconv.Quote(mv20)...)
				}
			}
			r = append(r, 125)
		}
		r = append(r, 125)
		if r[idx12] == 44 {
			r[idx12] = 32
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
		idx21 := len(r)
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
			idx22 := len(r)
			if len(me.SrcIntel.SrcIntels.InfoTips) != 0 {
				r = append(r, ",\"InfoTips\":"...)
				r = append(r, 91)
				for i23 := range me.SrcIntel.SrcIntels.InfoTips {
					if i23 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					idx24 := len(r)
					{
						r = append(r, ",\"value\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.SrcIntels.InfoTips[i23].Value)...)
					}
					if len(me.SrcIntel.SrcIntels.InfoTips[i23].Language) != 0 {
						r = append(r, ",\"language\":"...)
						r = append(r, pkg__strconv.Quote(me.SrcIntel.SrcIntels.InfoTips[i23].Language)...)
					}
					r = append(r, 125)
					if r[idx24] == 44 {
						r[idx24] = 32
					}
				}
				r = append(r, 93)
			}
			if len(me.SrcIntel.SrcIntels.Refs) != 0 {
				r = append(r, ",\"Refs\":"...)
				r = append(r, 91)
				for i25 := range me.SrcIntel.SrcIntels.Refs {
					if nil != me.SrcIntel.SrcIntels.Refs[i25] {
						{
							if i25 != 0 {
								r = append(r, 44)
							}
							var e error
							var sl []byte
							sl, e = me.SrcIntel.SrcIntels.Refs[i25].preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
					} else {
						if i25 != 0 {
							r = append(r, 44)
						}
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
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
				for i26 := range me.SrcIntel.Cmpl {
					if nil != me.SrcIntel.Cmpl[i26] {
						if i26 != 0 {
							r = append(r, 44)
						}
						r = append(r, 123)
						idx27 := len(r)
						if me.SrcIntel.Cmpl[i26].Kind != 0 {
							r = append(r, ",\"kind\":"...)
							r = append(r, pkg__strconv.FormatInt((int64)(me.SrcIntel.Cmpl[i26].Kind), 10)...)
						}
						{
							r = append(r, ",\"label\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i26].Label)...)
						}
						if nil != me.SrcIntel.Cmpl[i26].Documentation {
							{
								r = append(r, ",\"documentation\":"...)
								var e error
								var sl []byte
								sl, e = me.SrcIntel.Cmpl[i26].Documentation.preview_MarshalJSON()
								if e == nil {
									r = append(r, sl...)
								} else {
									err = e
									return
								}
							}
						}
						if len(me.SrcIntel.Cmpl[i26].Detail) != 0 {
							r = append(r, ",\"detail\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i26].Detail)...)
						}
						if len(me.SrcIntel.Cmpl[i26].SortText) != 0 {
							r = append(r, ",\"sortText\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Cmpl[i26].SortText)...)
						}
						r = append(r, 125)
						if r[idx27] == 44 {
							r[idx27] = 32
						}
					} else {
						if i26 != 0 {
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
				for i28 := range me.SrcIntel.Syms {
					if nil != me.SrcIntel.Syms[i28] {
						{
							if i28 != 0 {
								r = append(r, 44)
							}
							var e error
							var sl []byte
							sl, e = me.SrcIntel.Syms[i28].preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
					} else {
						if i28 != 0 {
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
				for i29 := range me.SrcIntel.Anns {
					if nil != me.SrcIntel.Anns[i29] {
						if i29 != 0 {
							r = append(r, 44)
						}
						r = append(r, 123)
						idx30 := len(r)
						{
							r = append(r, ",\"Range\":"...)
							var e error
							var sl []byte
							sl, e = me.SrcIntel.Anns[i29].Range.preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
						{
							r = append(r, ",\"Title\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i29].Title)...)
						}
						if len(me.SrcIntel.Anns[i29].Desc) != 0 {
							r = append(r, ",\"Desc\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i29].Desc)...)
						}
						{
							r = append(r, ",\"CmdName\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcIntel.Anns[i29].CmdName)...)
						}
						r = append(r, 125)
						if r[idx30] == 44 {
							r[idx30] = 32
						}
					} else {
						if i29 != 0 {
							r = append(r, 44)
						}
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
			}
			r = append(r, 125)
			if r[idx22] == 44 {
				r[idx22] = 32
			}
		}
		if nil != me.SrcDiags {
			r = append(r, ",\"srcDiags\":"...)
			r = append(r, 123)
			idx31 := len(r)
			{
				r = append(r, ",\"All\":"...)
				mf32 := true
				r = append(r, 123)
				for mk33, mv34 := range me.SrcDiags.All {
					{
						{
							if mf32 {
								mf32 = false
							} else {
								r = append(r, 44)
							}
							r = append(r, pkg__strconv.Quote(mk33)...)
							r = append(r, 58)
						}
						r = append(r, 91)
						for i35 := range mv34 {
							if nil != mv34[i35] {
								if i35 != 0 {
									r = append(r, 44)
								}
								r = append(r, 123)
								idx36 := len(r)
								if len(mv34[i35].Cat) != 0 {
									r = append(r, ",\"Cat\":"...)
									r = append(r, pkg__strconv.Quote(mv34[i35].Cat)...)
								}
								{
									r = append(r, ",\"Loc\":"...)
									var e error
									var sl []byte
									sl, e = mv34[i35].Loc.preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
								{
									r = append(r, ",\"Msg\":"...)
									r = append(r, pkg__strconv.Quote(mv34[i35].Msg)...)
								}
								if len(mv34[i35].SrcActions) != 0 {
									r = append(r, ",\"SrcActions\":"...)
									r = append(r, 91)
									for i37 := range mv34[i35].SrcActions {
										if i37 != 0 {
											r = append(r, 44)
										}
										r = append(r, 123)
										idx38 := len(r)
										{
											r = append(r, ",\"title\":"...)
											r = append(r, pkg__strconv.Quote(mv34[i35].SrcActions[i37].Title)...)
										}
										{
											r = append(r, ",\"command\":"...)
											r = append(r, pkg__strconv.Quote(mv34[i35].SrcActions[i37].Cmd)...)
										}
										if len(mv34[i35].SrcActions[i37].Hint) != 0 {
											r = append(r, ",\"tooltip\":"...)
											r = append(r, pkg__strconv.Quote(mv34[i35].SrcActions[i37].Hint)...)
										}
										if len(mv34[i35].SrcActions[i37].Arguments) != 0 {
											r = append(r, ",\"arguments\":"...)
											r = append(r, 91)
											for i39 := range mv34[i35].SrcActions[i37].Arguments {
												if mv34[i35].SrcActions[i37].Arguments[i39] != nil {
													if i39 != 0 {
														r = append(r, 44)
													}
													var e error
													var sl []byte
													j, ok := mv34[i35].SrcActions[i37].Arguments[i39].(pkg__encoding_json.Marshaler)
													if ok && (j != nil) {
														sl, e = j.MarshalJSON()
													} else {
														sl, e = pkg__encoding_json.Marshal(mv34[i35].SrcActions[i37].Arguments[i39])
													}
													if e == nil {
														r = append(r, sl...)
													} else {
														err = e
														return
													}
												} else {
													if i39 != 0 {
														r = append(r, 44)
													}
													r = append(r, "null"...)
												}
											}
											r = append(r, 93)
										}
										r = append(r, 125)
										if r[idx38] == 44 {
											r[idx38] = 32
										}
									}
									r = append(r, 93)
								}
								if mv34[i35].StickyAuto {
									r = append(r, ",\"Sticky\":"...)
									r = append(r, pkg__strconv.FormatBool(mv34[i35].StickyAuto)...)
								}
								if len(mv34[i35].Tags) != 0 {
									r = append(r, ",\"Tags\":"...)
									r = append(r, 91)
									for i40 := range mv34[i35].Tags {
										{
											if i40 != 0 {
												r = append(r, 44)
											}
											r = append(r, pkg__strconv.FormatInt((int64)(mv34[i35].Tags[i40]), 10)...)
										}
									}
									r = append(r, 93)
								}
								r = append(r, 125)
								if r[idx36] == 44 {
									r[idx36] = 32
								}
							} else {
								if i35 != 0 {
									r = append(r, 44)
								}
								r = append(r, "null"...)
							}
						}
						r = append(r, 93)
					}
				}
				r = append(r, 125)
			}
			{
				r = append(r, ",\"FixUps\":"...)
				r = append(r, 91)
				for i41 := range me.SrcDiags.FixUps {
					if nil != me.SrcDiags.FixUps[i41] {
						if i41 != 0 {
							r = append(r, 44)
						}
						r = append(r, 123)
						idx42 := len(r)
						{
							r = append(r, ",\"FilePath\":"...)
							r = append(r, pkg__strconv.Quote(me.SrcDiags.FixUps[i41].FilePath)...)
						}
						{
							r = append(r, ",\"Desc\":"...)
							mf43 := true
							r = append(r, 123)
							for mk44, mv45 := range me.SrcDiags.FixUps[i41].Desc {
								{
									{
										if mf43 {
											mf43 = false
										} else {
											r = append(r, 44)
										}
										r = append(r, pkg__strconv.Quote(mk44)...)
										r = append(r, 58)
									}
									r = append(r, 91)
									for i46 := range mv45 {
										{
											if i46 != 0 {
												r = append(r, 44)
											}
											r = append(r, pkg__strconv.Quote(mv45[i46])...)
										}
									}
									r = append(r, 93)
								}
							}
							r = append(r, 125)
						}
						{
							r = append(r, ",\"Edits\":"...)
							r = append(r, 91)
							for i47 := range me.SrcDiags.FixUps[i41].Edits {
								{
									if i47 != 0 {
										r = append(r, 44)
									}
									var e error
									var sl []byte
									sl, e = me.SrcDiags.FixUps[i41].Edits[i47].preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
							}
							r = append(r, 93)
						}
						{
							r = append(r, ",\"Dropped\":"...)
							r = append(r, 91)
							for i48 := range me.SrcDiags.FixUps[i41].Dropped {
								{
									if i48 != 0 {
										r = append(r, 44)
									}
									var e error
									var sl []byte
									sl, e = me.SrcDiags.FixUps[i41].Dropped[i48].preview_MarshalJSON()
									if e == nil {
										r = append(r, sl...)
									} else {
										err = e
										return
									}
								}
							}
							r = append(r, 93)
						}
						r = append(r, 125)
						if r[idx42] == 44 {
							r[idx42] = 32
						}
					} else {
						if i41 != 0 {
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
			if r[idx31] == 44 {
				r[idx31] = 32
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
			for i49 := range me.SrcMods {
				if nil != me.SrcMods[i49] {
					{
						if i49 != 0 {
							r = append(r, 44)
						}
						var e error
						var sl []byte
						sl, e = me.SrcMods[i49].preview_MarshalJSON()
						if e == nil {
							r = append(r, sl...)
						} else {
							err = e
							return
						}
					}
				} else {
					if i49 != 0 {
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
			for i50 := range me.SrcActions {
				if i50 != 0 {
					r = append(r, 44)
				}
				r = append(r, 123)
				idx51 := len(r)
				{
					r = append(r, ",\"title\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i50].Title)...)
				}
				{
					r = append(r, ",\"command\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i50].Cmd)...)
				}
				if len(me.SrcActions[i50].Hint) != 0 {
					r = append(r, ",\"tooltip\":"...)
					r = append(r, pkg__strconv.Quote(me.SrcActions[i50].Hint)...)
				}
				if len(me.SrcActions[i50].Arguments) != 0 {
					r = append(r, ",\"arguments\":"...)
					r = append(r, 91)
					for i52 := range me.SrcActions[i50].Arguments {
						if me.SrcActions[i50].Arguments[i52] != nil {
							if i52 != 0 {
								r = append(r, 44)
							}
							var e error
							var sl []byte
							j, ok := me.SrcActions[i50].Arguments[i52].(pkg__encoding_json.Marshaler)
							if ok && (j != nil) {
								sl, e = j.MarshalJSON()
							} else {
								sl, e = pkg__encoding_json.Marshal(me.SrcActions[i50].Arguments[i52])
							}
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						} else {
							if i52 != 0 {
								r = append(r, 44)
							}
							r = append(r, "null"...)
						}
					}
					r = append(r, 93)
				}
				r = append(r, 125)
				if r[idx51] == 44 {
					r[idx51] = 32
				}
			}
			r = append(r, 93)
		}
		if nil != me.Extras {
			r = append(r, ",\"extras\":"...)
			r = append(r, 123)
			idx53 := len(r)
			if len(me.Extras.SrcIntels.InfoTips) != 0 {
				r = append(r, ",\"InfoTips\":"...)
				r = append(r, 91)
				for i54 := range me.Extras.SrcIntels.InfoTips {
					if i54 != 0 {
						r = append(r, 44)
					}
					r = append(r, 123)
					idx55 := len(r)
					{
						r = append(r, ",\"value\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.SrcIntels.InfoTips[i54].Value)...)
					}
					if len(me.Extras.SrcIntels.InfoTips[i54].Language) != 0 {
						r = append(r, ",\"language\":"...)
						r = append(r, pkg__strconv.Quote(me.Extras.SrcIntels.InfoTips[i54].Language)...)
					}
					r = append(r, 125)
					if r[idx55] == 44 {
						r[idx55] = 32
					}
				}
				r = append(r, 93)
			}
			if len(me.Extras.SrcIntels.Refs) != 0 {
				r = append(r, ",\"Refs\":"...)
				r = append(r, 91)
				for i56 := range me.Extras.SrcIntels.Refs {
					if nil != me.Extras.SrcIntels.Refs[i56] {
						{
							if i56 != 0 {
								r = append(r, 44)
							}
							var e error
							var sl []byte
							sl, e = me.Extras.SrcIntels.Refs[i56].preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
					} else {
						if i56 != 0 {
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
				for i57 := range me.Extras.Items {
					if nil != me.Extras.Items[i57] {
						if i57 != 0 {
							r = append(r, 44)
						}
						r = append(r, 123)
						idx58 := len(r)
						{
							r = append(r, ",\"id\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].ID)...)
						}
						{
							r = append(r, ",\"label\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].Label)...)
						}
						{
							r = append(r, ",\"description\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].Desc)...)
						}
						if len(me.Extras.Items[i57].Detail) != 0 {
							r = append(r, ",\"detail\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].Detail)...)
						}
						if len(me.Extras.Items[i57].QueryArg) != 0 {
							r = append(r, ",\"arg\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].QueryArg)...)
						}
						if len(me.Extras.Items[i57].FilePos) != 0 {
							r = append(r, ",\"fPos\":"...)
							r = append(r, pkg__strconv.Quote(me.Extras.Items[i57].FilePos)...)
						}
						r = append(r, 125)
						if r[idx58] == 44 {
							r[idx58] = 32
						}
					} else {
						if i57 != 0 {
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
				for i59 := range me.Extras.Warns {
					{
						if i59 != 0 {
							r = append(r, 44)
						}
						r = append(r, pkg__strconv.Quote(me.Extras.Warns[i59])...)
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
			if r[idx53] == 44 {
				r[idx53] = 32
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
			idx60 := len(r)
			if nil != me.Menu.SubMenu {
				r = append(r, ",\"SubMenu\":"...)
				r = append(r, 123)
				idx61 := len(r)
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
					for i62 := range me.Menu.SubMenu.Items {
						if nil != me.Menu.SubMenu.Items[i62] {
							if i62 != 0 {
								r = append(r, 44)
							}
							r = append(r, 123)
							idx63 := len(r)
							if me.Menu.SubMenu.Items[i62].IpcID != 0 {
								r = append(r, ",\"ii\":"...)
								r = append(r, pkg__strconv.FormatInt((int64)(me.Menu.SubMenu.Items[i62].IpcID), 10)...)
							}
							if me.Menu.SubMenu.Items[i62].IpcArgs != nil {
								r = append(r, ",\"ia\":"...)
								var e error
								var sl []byte
								j, ok := me.Menu.SubMenu.Items[i62].IpcArgs.(pkg__encoding_json.Marshaler)
								if ok && (j != nil) {
									sl, e = j.MarshalJSON()
								} else {
									sl, e = pkg__encoding_json.Marshal(me.Menu.SubMenu.Items[i62].IpcArgs)
								}
								if e == nil {
									r = append(r, sl...)
								} else {
									err = e
									return
								}
							}
							if len(me.Menu.SubMenu.Items[i62].Category) != 0 {
								r = append(r, ",\"c\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i62].Category)...)
							}
							{
								r = append(r, ",\"t\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i62].Title)...)
							}
							if len(me.Menu.SubMenu.Items[i62].Desc) != 0 {
								r = append(r, ",\"d\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i62].Desc)...)
							}
							if len(me.Menu.SubMenu.Items[i62].Hint) != 0 {
								r = append(r, ",\"h\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i62].Hint)...)
							}
							if len(me.Menu.SubMenu.Items[i62].Confirm) != 0 {
								r = append(r, ",\"q\":"...)
								r = append(r, pkg__strconv.Quote(me.Menu.SubMenu.Items[i62].Confirm)...)
							}
							r = append(r, 125)
							if r[idx63] == 44 {
								r[idx63] = 32
							}
						} else {
							if i62 != 0 {
								r = append(r, 44)
							}
							r = append(r, "null"...)
						}
					}
					r = append(r, 93)
				}
				r = append(r, 125)
				if r[idx61] == 44 {
					r[idx61] = 32
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
				for i64 := range me.Menu.Refs {
					if nil != me.Menu.Refs[i64] {
						{
							if i64 != 0 {
								r = append(r, 44)
							}
							var e error
							var sl []byte
							sl, e = me.Menu.Refs[i64].preview_MarshalJSON()
							if e == nil {
								r = append(r, sl...)
							} else {
								err = e
								return
							}
						}
					} else {
						if i64 != 0 {
							r = append(r, 44)
						}
						r = append(r, "null"...)
					}
				}
				r = append(r, 93)
			}
			r = append(r, 125)
			if r[idx60] == 44 {
				r[idx60] = 32
			}
		}
		if nil != me.CaddyUpdate {
			r = append(r, ",\"caddy\":"...)
			r = append(r, 123)
			idx65 := len(r)
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
			idx66 := len(r)
			{
				r = append(r, ",\"Flag\":"...)
				r = append(r, pkg__strconv.FormatInt((int64)(me.CaddyUpdate.Status.Flag), 10)...)
			}
			if len(me.CaddyUpdate.Status.Desc) != 0 {
				r = append(r, ",\"Desc\":"...)
				r = append(r, pkg__strconv.Quote(me.CaddyUpdate.Status.Desc)...)
			}
			r = append(r, 125)
			if r[idx66] == 44 {
				r[idx66] = 32
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
			if r[idx65] == 44 {
				r[idx65] = 32
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
		if r[idx21] == 44 {
			r[idx21] = 32
		}
	}
	return
}

// preview_UnmarshalJSON implements the Go standard library's `encoding/json.Unmarshaler` interface.
func (me *fooResp) preview_UnmarshalJSON(v []byte) (err error) { return }
