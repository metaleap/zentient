package z

type IpcIDs uint8

const (
	_ IpcIDs = iota

	IPCID_MENUS_MAIN
	IPCID_MENUS_PKGS
	IPCID_MENUS_TOOLS

	IPCID_OBJ_SNAPSHOT
	IPCID_PAGE_HTML
	IPCID_TREEVIEW_GETITEM
	IPCID_TREEVIEW_CHILDREN
	IPCID_TREEVIEW_CHANGED
	IPCID_CFG_RESETALL
	IPCID_CFG_LIST
	IPCID_CFG_SET
	IPCID_NOTIFY_INFO
	IPCID_NOTIFY_WARN
	IPCID_NOTIFY_ERR

	IPCID_PROJ_CHANGED
	IPCID_PROJ_POLLEVTS

	IPCID_SRCDIAG_LIST
	IPCID_SRCDIAG_RUN_CURFILE
	IPCID_SRCDIAG_RUN_OPENFILES
	IPCID_SRCDIAG_RUN_ALLFILES
	IPCID_SRCDIAG_FORGETALL
	IPCID_SRCDIAG_PEEKHIDDEN
	IPCID_SRCDIAG_PUB
	IPCID_SRCDIAG_AUTO_TOGGLE
	IPCID_SRCDIAG_AUTO_ALL
	IPCID_SRCDIAG_AUTO_NONE
	IPCID_SRCDIAG_STARTED
	IPCID_SRCDIAG_FINISHED

	IPCID_SRCMOD_FMT_SETDEFMENU
	IPCID_SRCMOD_FMT_SETDEFPICK
	IPCID_SRCMOD_FMT_RUNONFILE
	IPCID_SRCMOD_FMT_RUNONSEL
	IPCID_SRCMOD_RENAME
	IPCID_SRCMOD_ACTIONS

	IPCID_SRCINTEL_HOVER
	IPCID_SRCINTEL_SYMS_FILE
	IPCID_SRCINTEL_SYMS_PROJ
	IPCID_SRCINTEL_CMPL_ITEMS
	IPCID_SRCINTEL_CMPL_DETAILS
	IPCID_SRCINTEL_HIGHLIGHTS
	IPCID_SRCINTEL_ANNS
	IPCID_SRCINTEL_SIGNATURE
	IPCID_SRCINTEL_REFERENCES
	IPCID_SRCINTEL_DEFSYM
	IPCID_SRCINTEL_DEFTYPE
	IPCID_SRCINTEL_DEFIMPL

	IPCID_EXTRAS_INTEL_LIST
	IPCID_EXTRAS_INTEL_RUN
	IPCID_EXTRAS_QUERY_LIST
	IPCID_EXTRAS_QUERY_RUN
)

type DiagSeverity int

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

type CaddyStatus uint8

const (
	CADDY_PENDING CaddyStatus = iota
	CADDY_ERROR
	CADDY_BUSY
	CADDY_GOOD
)

type ipcReq struct {
	ReqID   int64       `json:"ri"`
	IpcID   IpcIDs      `json:"ii"`
	IpcArgs interface{} `json:"ia"`

	ProjUpd *WorkspaceChanges `json:"projUpd"`
	SrcLens *SrcLens          `json:"srcLens"`
}

type ipcResp struct {
	IpcID       IpcIDs           `json:"ii,omitempty"`
	ReqID       int64            `json:"ri,omitempty"`
	ErrMsg      string           `json:"err,omitempty"`
	SrcIntel    *ipcRespSrcIntel `json:"sI,omitempty"`
	SrcDiags    *ipcRespDiag     `json:"srcDiags,omitempty"`
	SrcMods     SrcLenses        `json:"srcMods,omitempty"`
	SrcActions  []EditorAction   `json:"srcActions,omitempty"`
	Extras      *IpcRespExtras   `json:"extras,omitempty"`
	Menu        *ipcRespMenu     `json:"menu,omitempty"`
	CaddyUpdate *Caddy           `json:"caddy,omitempty"`
	Val         interface{}      `json:"val,omitempty"`
}

type ipcRespDiag struct {
	All    diagItemsBy
	FixUps []*diagFixUps
	LangID string
}

type IpcRespExtras struct {
	SrcIntels
	Items []*ExtrasItem
	Warns []string `json:",omitempty"`
	Desc  string   `json:",omitempty"`
	Url   string   `json:",omitempty"`
}

type ipcRespMenu struct {
	SubMenu       *Menu   `json:",omitempty"`
	WebsiteURL    string  `json:",omitempty"`
	NoteInfo      string  `json:",omitempty"`
	NoteWarn      string  `json:",omitempty"`
	UxActionLabel string  `json:",omitempty"`
	Refs          SrcLocs `json:",omitempty"`
}

type ipcRespSrcIntel struct {
	SrcIntels
	Sig  *SrcIntelSigHelp  `json:",omitempty"`
	Cmpl SrcIntelCompls    `json:",omitempty"`
	Syms SrcLenses         `json:",omitempty"`
	Anns []*SrcAnnotaction `json:",omitempty"`
}

type Caddy struct {
	ID     string `json:",omitempty"`
	LangID string `json:",omitempty"`
	Icon   string
	Title  string `json:",omitempty"`
	Status struct {
		Flag CaddyStatus
		Desc string `json:",omitempty"`
	}
	Details                 string `json:",omitempty"`
	UxActionID              string `json:",omitempty"`
	ShowTitle               bool   `json:",omitempty"`
	ShouldReRunWhenNextDone bool   `json:"-"`

	ready bool

	OnReady func() `json:"-"`
}

type diagFixUps struct {
	FilePath string
	Desc     map[string][]string
	Edits    SrcModEdits
	Dropped  []srcModEdit
}

type DiagItem struct {
	Cat         string `json:",omitempty"`
	Loc         SrcLoc
	Msg         string
	SrcActions  []EditorAction `json:",omitempty"`
	StickyForce bool           `json:"-"`
	StickyAuto  bool           `json:"Sticky,omitempty"`
	Tags        []int          `json:"Tags,omitempty"`
	Misc        []interface{}  `json:"-"`
}

type DiagItems []*DiagItem

type diagItemsBy map[string]DiagItems

type EditorAction struct {
	Title     string        `json:"title"`
	Cmd       string        `json:"command"`
	Hint      string        `json:"tooltip,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

type ExtrasItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Desc     string `json:"description"`
	Detail   string `json:"detail,omitempty"`
	QueryArg string `json:"arg,omitempty"`
	FilePos  string `json:"fPos,omitempty"`
}

type Menu struct {
	Desc     string    `json:"desc,omitempty"`
	TopLevel bool      `json:"topLevel,omitempty"`
	Items    MenuItems `json:"items"`
}

type MenuItems []*MenuItem

type MenuItem struct {
	IpcID    IpcIDs      `json:"ii,omitempty"`
	IpcArgs  interface{} `json:"ia,omitempty"`
	Category string      `json:"c,omitempty"`
	Title    string      `json:"t"`
	Desc     string      `json:"d,omitempty"`
	Hint     string      `json:"h,omitempty"`
	Confirm  string      `json:"q,omitempty"`

	tag string
}

type SrcAnnotaction struct {
	Range   SrcRange
	Title   string
	Desc    string `json:",omitempty"`
	CmdName string
}

type SrcInfoTip struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}

type SrcIntelCompl struct {
	Kind          Completion   `json:"kind,omitempty"`
	Label         string       `json:"label"`
	Documentation *SrcIntelDoc `json:"documentation,omitempty"`
	Detail        string       `json:"detail,omitempty"`
	SortText      string       `json:"sortText,omitempty"`
	// FilterText    string       `json:"filterText,omitempty"`
	// InsertText    string       `json:"insertText,omitempty"`
	// CommitChars   []string     `json:"commitCharacters,omitempty"` // basically in all languages always operator/separator/punctuation (that is, "non-identifier") chars --- no need to send them for each item, for each language --- the client-side will do it
	SortPrio int `json:"-"`
}

type SrcIntelCompls []*SrcIntelCompl

type SrcIntels struct {
	InfoTips []SrcInfoTip `json:",omitempty"`
	Refs     SrcLocs      `json:",omitempty"`
}

type SrcIntelDoc struct {
	Value     string `json:"value,omitempty"`
	IsTrusted bool   `json:"isTrusted,omitempty"`
}

type SrcIntelSigHelp struct {
	ActiveSignature int               `json:"activeSignature"`
	ActiveParameter int               `json:"activeParameter,omitempty"`
	Signatures      []SrcIntelSigInfo `json:"signatures,omitempty"`
}

type SrcIntelSigInfo struct {
	Label         string             `json:"label"`
	Documentation SrcIntelDoc        `json:"documentation,omitempty"`
	Parameters    []SrcIntelSigParam `json:"parameters"`
}

type SrcIntelSigParam struct {
	Label         string      `json:"label"`
	Documentation SrcIntelDoc `json:"documentation,omitempty"`
}

type SrcLenses []*SrcLens

type SrcLens struct {
	SrcLoc
	Txt  string `json:"t,omitempty"`
	Str  string `json:"s,omitempty"`
	CrLf bool   `json:"l,omitempty"`
}

type SrcLoc struct {
	Flag     int       `json:"e"` // don't omitempty
	FilePath string    `json:"f,omitempty"`
	Pos      *SrcPos   `json:"p,omitempty"`
	Range    *SrcRange `json:"r,omitempty"`
}

type SrcLocs []*SrcLoc

type srcModEdit struct {
	At  *SrcRange
	Val string // if not empty: inserts if At is pos, replaces if At is range. if empty: deletes if At is range, errors if At is pos.
}

type SrcModEdits []srcModEdit

// All public fields are 1-based (so 0 means 'missing') and rune-not-byte-based
type SrcPos struct {
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
	// rune1 not byte0 offset!
	Off int `json:"o,omitempty"`

	// if & when this is computed, it'll be 0-based
	byteOff int
	byteoff bool
}

type SrcRange struct {
	Start SrcPos `json:"s"`
	End   SrcPos `json:"e,omitempty"`
}

type WorkspaceChanges struct {
	AddedDirs    []string
	RemovedDirs  []string
	OpenedFiles  []string
	ClosedFiles  []string
	WrittenFiles []string
	LiveFiles    map[string]string
}
