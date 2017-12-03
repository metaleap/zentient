package z

// REMINDER TO SELF, on the rare update also update corresponding String()ers at the end of the file

type msgIDs uint8

const (
	_ msgIDs = iota

	MSGID_CORECMDS_PALETTE

	MSGID_SRCFMT_SETDEFMENU
	MSGID_SRCFMT_SETDEFPICK
	MSGID_SRCFMT_RUNONFILE
	MSGID_SRCFMT_RUNONSEL

	MSGID_SRCINTEL_HOVER
	MSGID_SRCINTEL_SYMS_FILE
	MSGID_SRCINTEL_SYMS_PROJ

	MSGID_MIN_INVALID
)

type DiagSeverity uint8

const (
	DIAG_SEV_ERR DiagSeverity = iota
	DIAG_SEV_WARN
	DIAG_SEV_INFO
	DIAG_SEV_HINT
)

type Symbol uint8

const (
	SYM_FILE Symbol = iota
	SYM_MODULE
	SYM_NAMESPACE
	SYM_PACKAGE
	SYM_CLASS
	SYM_METHOD
	SYM_PROPERTY
	SYM_FIELD
	SYM_CONSTRUCTOR
	SYM_ENUM
	SYM_INTERFACE
	SYM_FUNCTION
	SYM_VARIABLE
	SYM_CONSTANT
	SYM_STRING
	SYM_NUMBER
	SYM_BOOLEAN
	SYM_ARRAY
	SYM_OBJECT
	SYM_KEY
	SYM_NULL
	SYM_ENUMMEMBER
	SYM_STRUCT
	SYM_EVENT
	SYM_OPERATOR
	SYM_TYPEPARAMETER
)

type Completion uint8

const (
	CMPL_TEXT Completion = iota
	CMPL_METHOD
	CMPL_FUNCTION
	CMPL_CONSTRUCTOR
	CMPL_FIELD
	CMPL_VARIABLE
	CMPL_CLASS
	CMPL_INTERFACE
	CMPL_MODULE
	CMPL_PROPERTY
	CMPL_UNIT
	CMPL_VALUE
	CMPL_ENUM
	CMPL_KEYWORD
	CMPL_SNIPPET
	CMPL_COLOR
	CMPL_FILE
	CMPL_REFERENCE
	CMPL_FOLDER
	CMPL_ENUMMEMBER
	CMPL_CONSTANT
	CMPL_STRUCT
	CMPL_EVENT
	CMPL_OPERATOR
	CMPL_TYPEPARAMETER
)

// Stringers below.
// Yes go generate exists. Not into it. One day I'll generate stuff like this with gomacro, for now manual werks as changes are rare

func (me msgIDs) String() string {
	switch me {
	case MSGID_CORECMDS_PALETTE:
		return "MSGID_CORECMDS_PALETTE"
	case MSGID_SRCFMT_SETDEFMENU:
		return "MSGID_SRCFMT_SETDEFMENU"
	case MSGID_SRCFMT_SETDEFPICK:
		return "MSGID_SRCFMT_SETDEFPICK"
	case MSGID_SRCFMT_RUNONFILE:
		return "MSGID_SRCFMT_RUNONFILE"
	case MSGID_SRCFMT_RUNONSEL:
		return "MSGID_SRCFMT_RUNONSEL"
	case MSGID_SRCINTEL_HOVER:
		return "MSGID_SRCINTEL_HOVER"
	case MSGID_SRCINTEL_SYMS_FILE:
		return "MSGID_SRCINTEL_SYMS_FILE"
	case MSGID_SRCINTEL_SYMS_PROJ:
		return "MSGID_SRCINTEL_SYMS_PROJ"
	}
	return Strf("%d", me)
}

func (me DiagSeverity) String() string {
	switch me {
	case DIAG_SEV_ERR:
		return "DIAG_SEV_ERR"
	case DIAG_SEV_WARN:
		return "DIAG_SEV_WARN"
	case DIAG_SEV_INFO:
		return "DIAG_SEV_INFO"
	case DIAG_SEV_HINT:
		return "DIAG_SEV_HINT"
	}
	return Strf("%d", me)
}

func (me Symbol) String() string {
	switch me {
	case SYM_FILE:
		return "SYM_FILE Symbol"
	case SYM_MODULE:
		return "SYM_MODULE"
	case SYM_NAMESPACE:
		return "SYM_NAMESPACE"
	case SYM_PACKAGE:
		return "SYM_PACKAGE"
	case SYM_CLASS:
		return "SYM_CLASS"
	case SYM_METHOD:
		return "SYM_METHOD"
	case SYM_PROPERTY:
		return "SYM_PROPERTY"
	case SYM_FIELD:
		return "SYM_FIELD"
	case SYM_CONSTRUCTOR:
		return "SYM_CONSTRUCTOR"
	case SYM_ENUM:
		return "SYM_ENUM"
	case SYM_INTERFACE:
		return "SYM_INTERFACE"
	case SYM_FUNCTION:
		return "SYM_FUNCTION"
	case SYM_VARIABLE:
		return "SYM_VARIABLE"
	case SYM_CONSTANT:
		return "SYM_CONSTANT"
	case SYM_STRING:
		return "SYM_STRING"
	case SYM_NUMBER:
		return "SYM_NUMBER"
	case SYM_BOOLEAN:
		return "SYM_BOOLEAN"
	case SYM_ARRAY:
		return "SYM_ARRAY"
	case SYM_OBJECT:
		return "SYM_OBJECT"
	case SYM_KEY:
		return "SYM_KEY"
	case SYM_NULL:
		return "SYM_NULL"
	case SYM_ENUMMEMBER:
		return "SYM_ENUMMEMBER"
	case SYM_STRUCT:
		return "SYM_STRUCT"
	case SYM_EVENT:
		return "SYM_EVENT"
	case SYM_OPERATOR:
		return "SYM_OPERATOR"
	case SYM_TYPEPARAMETER:
		return "SYM_TYPEPARAMETER"
	}
	return Strf("%d", me)
}

func (me Completion) String() string {
	switch me {
	case CMPL_TEXT:
		return "CMPL_TEXT"
	case CMPL_METHOD:
		return "CMPL_METHOD"
	case CMPL_FUNCTION:
		return "CMPL_FUNCTION"
	case CMPL_CONSTRUCTOR:
		return "CMPL_CONSTRUCTOR"
	case CMPL_FIELD:
		return "CMPL_FIELD"
	case CMPL_VARIABLE:
		return "CMPL_VARIABLE"
	case CMPL_CLASS:
		return "CMPL_CLASS"
	case CMPL_INTERFACE:
		return "CMPL_INTERFACE"
	case CMPL_MODULE:
		return "CMPL_MODULE"
	case CMPL_PROPERTY:
		return "CMPL_PROPERTY"
	case CMPL_UNIT:
		return "CMPL_UNIT"
	case CMPL_VALUE:
		return "CMPL_VALUE"
	case CMPL_ENUM:
		return "CMPL_ENUM"
	case CMPL_KEYWORD:
		return "CMPL_KEYWORD"
	case CMPL_SNIPPET:
		return "CMPL_SNIPPET"
	case CMPL_COLOR:
		return "CMPL_COLOR"
	case CMPL_FILE:
		return "CMPL_FILE"
	case CMPL_REFERENCE:
		return "CMPL_REFERENCE"
	case CMPL_FOLDER:
		return "CMPL_FOLDER"
	case CMPL_ENUMMEMBER:
		return "CMPL_ENUMMEMBER"
	case CMPL_CONSTANT:
		return "CMPL_CONSTANT"
	case CMPL_STRUCT:
		return "CMPL_STRUCT"
	case CMPL_EVENT:
		return "CMPL_EVENT"
	case CMPL_OPERATOR:
		return "CMPL_OPERATOR"
	case CMPL_TYPEPARAMETER:
		return "CMPL_TYPEPARAMETER"
	}
	return Strf("%d", me)
}
