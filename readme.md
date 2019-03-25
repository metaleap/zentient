# z
--
    import "github.com/metaleap/zentient"


## Usage

```go
var (
	Strf = fmt.Sprintf
	Lang struct {
		Enabled   bool
		ID        string
		Title     string
		SrcMod    ISrcMod
		SrcIntel  ISrcIntel
		Diag      IDiag
		Extras    IExtras
		PkgIntel  IPkgIntel
		Caddies   []*Caddy
		Settings  ISettings
		Tooling   ITooling
		Workspace IWorkspace
		Pages     IPages
	}
	Prog struct {
		Cfg Config

		Name string
		Dir  struct {
			Cache  string
			Config string
		}
	}
)
```

#### func  BadMsg

```go
func BadMsg(what string, which string) string
```

#### func  BadPanic

```go
func BadPanic(what string, which string)
```

#### func  CaddyBuildOnDone

```go
func CaddyBuildOnDone(failed map[string]bool, skipped map[string]bool, all []string, timeTaken time.Duration)
```

#### func  CaddyBuildOnRunning

```go
func CaddyBuildOnRunning(numJobs int, cur int, all string)
```

#### func  Init

```go
func Init() (err error)
```

#### func  InitAndServe

```go
func InitAndServe(onPreInit func(), onPostInit func()) (err error)
```

#### func  InitAndServeOrPanic

```go
func InitAndServeOrPanic(onPreInit func(), onPostInit func())
```

#### func  PrettifyPathsIn

```go
func PrettifyPathsIn(s string) string
```

#### func  SendNotificationMessageToClient

```go
func SendNotificationMessageToClient(level DiagSeverity, message string) (err error)
```

#### func  Serve

```go
func Serve() (err error)
```

#### func  ToolGonePanic

```go
func ToolGonePanic(missingToolName string)
```

#### func  ToolsMsgGone

```go
func ToolsMsgGone(missingToolName string) string
```

#### func  ToolsMsgMore

```go
func ToolsMsgMore(missingToolName string) string
```

#### type BuildProgress

```go
type BuildProgress struct {
	NumJobs   int
	StartTime time.Time
	Failed    map[string]bool
	Skipped   map[string]bool
	PkgNames  []string
}
```


#### func  NewBuildProgress

```go
func NewBuildProgress(numJobs int) *BuildProgress
```

#### func (*BuildProgress) AddPkgName

```go
func (this *BuildProgress) AddPkgName(pkgName string)
```

#### func (*BuildProgress) OnDone

```go
func (this *BuildProgress) OnDone()
```

#### func (*BuildProgress) OnJob

```go
func (this *BuildProgress) OnJob(i int)
```

#### func (*BuildProgress) String

```go
func (this *BuildProgress) String() string
```

#### type Caddy

```go
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

	OnReady func() `json:"-"`
}
```


#### func (*Caddy) IsPendingOrBusy

```go
func (this *Caddy) IsPendingOrBusy() bool
```

#### func (*Caddy) IsReady

```go
func (this *Caddy) IsReady() bool
```

#### func (*Caddy) OnStatusChanged

```go
func (this *Caddy) OnStatusChanged()
```

#### type CaddyStatus

```go
type CaddyStatus uint8
```


```go
const (
	CADDY_PENDING CaddyStatus = iota
	CADDY_ERROR
	CADDY_BUSY
	CADDY_GOOD
)
```

#### func (CaddyStatus) String

```go
func (me CaddyStatus) String() (r string)
```
String implements the Go standard library's `fmt.Stringer` interface.

#### type Completion

```go
type Completion uint8
```


```go
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
```

#### func (Completion) String

```go
func (me Completion) String() (r string)
```
String implements the Go standard library's `fmt.Stringer` interface.

#### type Config

```go
type Config struct {
	Internal      map[string]interface{} `json:",omitempty"`
	FormatterName string                 `json:",omitempty"`
	FormatterProg string                 `json:",omitempty"`
	AutoDiags     []string               `json:",omitempty"`
}
```


#### func (*Config) Save

```go
func (this *Config) Save() (err error)
```

#### type DiagBase

```go
type DiagBase struct {
	Impl IDiag
}
```


#### func (*DiagBase) FixerUppers

```go
func (*DiagBase) FixerUppers() []FixerUpper
```

#### func (*DiagBase) Init

```go
func (this *DiagBase) Init()
```

#### func (*DiagBase) MenuCategory

```go
func (this *DiagBase) MenuCategory() string
```

#### func (*DiagBase) NewDiagItemFrom

```go
func (this *DiagBase) NewDiagItemFrom(srcRef *udev.SrcMsg, toolName string, fallbackFilePath func() string) (di *DiagItem)
```

#### func (*DiagBase) UpdateBuildDiagsAsNeeded

```go
func (this *DiagBase) UpdateBuildDiagsAsNeeded(workspaceFiles WorkspaceFiles, writtenFiles []string)
```

#### func (*DiagBase) UpdateLintDiagsIfAndAsNeeded

```go
func (this *DiagBase) UpdateLintDiagsIfAndAsNeeded(workspaceFiles WorkspaceFiles, autos bool, onlyFilePaths ...string)
```

#### type DiagBuildJobs

```go
type DiagBuildJobs []*DiagJobBuild
```


#### func (DiagBuildJobs) Len

```go
func (this DiagBuildJobs) Len() int
```

#### func (DiagBuildJobs) Less

```go
func (this DiagBuildJobs) Less(i int, j int) bool
```

#### func (DiagBuildJobs) Swap

```go
func (this DiagBuildJobs) Swap(i int, j int)
```

#### type DiagItem

```go
type DiagItem struct {
	Cat         string `json:",omitempty"`
	Loc         SrcLoc
	Msg         string
	SrcActions  []EditorAction `json:",omitempty"`
	StickyForce bool           `json:"-"`
	StickyAuto  bool           `json:"Sticky,omitempty"`
}
```


#### type DiagItems

```go
type DiagItems []*DiagItem
```


#### type DiagJob

```go
type DiagJob struct {
	AffectedFilePaths []string
	Target            IDiagJobTarget
}
```


#### func (*DiagJob) String

```go
func (this *DiagJob) String() string
```

#### type DiagJobBuild

```go
type DiagJobBuild struct {
	DiagJob
	TargetCmp func(IDiagJobTarget, IDiagJobTarget) bool
	Succeeded bool
}
```


#### func (*DiagJobBuild) IsSortedPriorTo

```go
func (this *DiagJobBuild) IsSortedPriorTo(cmp interface{}) bool
```

#### type DiagJobLint

```go
type DiagJobLint struct {
	DiagJob
	Tool *Tool
}
```


#### func (*DiagJobLint) Yield

```go
func (this *DiagJobLint) Yield(diag *DiagItem)
```

#### type DiagLintJobs

```go
type DiagLintJobs []*DiagJobLint
```


#### type DiagSeverity

```go
type DiagSeverity int
```


```go
const (
	DIAG_SEV_ERR DiagSeverity = iota
	DIAG_SEV_WARN
	DIAG_SEV_INFO
	DIAG_SEV_HINT
)
```

#### func (DiagSeverity) String

```go
func (me DiagSeverity) String() (r string)
```
String implements the Go standard library's `fmt.Stringer` interface.

#### type EditorAction

```go
type EditorAction struct {
	Title     string        `json:"title"`
	Cmd       string        `json:"command"`
	Hint      string        `json:"tooltip,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
```


#### type ExtrasBase

```go
type ExtrasBase struct {
	Impl IExtras
}
```


#### func (*ExtrasBase) Init

```go
func (*ExtrasBase) Init()
```

#### type ExtrasItem

```go
type ExtrasItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Desc     string `json:"description"`
	Detail   string `json:"detail,omitempty"`
	QueryArg string `json:"arg,omitempty"`
	FilePos  string `json:"fPos,omitempty"`
}
```


#### type ExtrasResp

```go
type ExtrasResp struct {
	SrcIntels
	Items []*ExtrasItem
	Warns []string `json:",omitempty"`
	Desc  string   `json:",omitempty"`
	Url   string   `json:",omitempty"`
}
```


#### type FixUp

```go
type FixUp struct {
	Name  string
	Items []string
	Edits SrcModEdits
}
```


#### type FixerUpper

```go
type FixerUpper func(*DiagItem) *FixUp
```


#### type IDiag

```go
type IDiag interface {
	IDiagBuild
	IDiagLint
	IMenuItems
	// contains filtered or unexported methods
}
```


#### type IDiagBuild

```go
type IDiagBuild interface {
	FixerUppers() []FixerUpper
	OnUpdateBuildDiags([]string) DiagBuildJobs
	RunBuildJobs(DiagBuildJobs, WorkspaceFiles) DiagItems
	UpdateBuildDiagsAsNeeded(WorkspaceFiles, []string)
}
```


#### type IDiagJobTarget

```go
type IDiagJobTarget interface {
}
```


#### type IDiagLint

```go
type IDiagLint interface {
	KnownLinters() Tools
	OnUpdateLintDiags(WorkspaceFiles, Tools, []string) DiagLintJobs
	RunLintJob(*DiagJobLint, WorkspaceFiles)
	UpdateLintDiagsIfAndAsNeeded(WorkspaceFiles, bool, ...string)
}
```


#### type IExtras

```go
type IExtras interface {
	ListIntelExtras() []*ExtrasItem
	ListQueryExtras() []*ExtrasItem
	RunIntelExtra(*SrcLens, string, string, *ExtrasResp)
	RunQueryExtra(*SrcLens, string, string, *ExtrasResp)
	// contains filtered or unexported methods
}
```


#### type IList

```go
type IList interface {
	UnfilteredDesc() string
	Count(ListFilters) int
	FilterByID(string) *ListFilter
	Filters() []*ListFilter
	List(ListFilters) ListItems
}
```


#### type IListItem

```go
type IListItem interface {
	ISortable
}
```


#### type IListMenu

```go
type IListMenu interface {
	IList
	IMenuItems

	ListItemToMenuItem(IListItem) *MenuItem
	// contains filtered or unexported methods
}
```


#### type IMenuItems

```go
type IMenuItems interface {
	MenuCategory() string
	// contains filtered or unexported methods
}
```


#### type IObjSnap

```go
type IObjSnap interface {
	ObjSnapPrefix() string
	ObjSnap(string) interface{}
}
```


#### type IPages

```go
type IPages interface {
	PageBodyInnerHtml(string, []string, url.Values, string) string
	// contains filtered or unexported methods
}
```


#### type IPkgIntel

```go
type IPkgIntel interface {
	IListMenu
	IObjSnap
	Pkgs() PkgInfos
}
```


#### type ISettings

```go
type ISettings interface {
	IMenuItems

	KnownSettings() Settings
}
```


#### type ISortable

```go
type ISortable interface {
	IsSortedPriorTo(interface{}) bool
}
```


#### type ISrcIntel

```go
type ISrcIntel interface {
	ComplDetails(*SrcLens, string) *SrcIntelCompl
	ComplItems(*SrcLens) SrcIntelCompls
	ComplItemsShouldSort(*SrcLens) bool
	DefSym(*SrcLens) SrcLocs
	DefType(*SrcLens) SrcLocs
	DefImpl(*SrcLens) SrcLocs
	Highlights(*SrcLens, string) SrcLocs
	Hovers(*SrcLens) []InfoTip
	References(*SrcLens, bool) SrcLocs
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) SrcLenses
	// contains filtered or unexported methods
}
```


#### type ISrcMod

```go
type ISrcMod interface {
	IMenuItems

	CodeActions(*SrcLens) []EditorAction
	DoesStdoutWithFilePathArg(*Tool) bool
	KnownFormatters() Tools
	RunRenamer(*SrcLens, string) SrcLenses
	RunFormatter(*Tool, string, *SrcFormattingClientPrefs, string, string) (string, string)
}
```


#### type ITooling

```go
type ITooling interface {
	IMenuItems

	KnownTools() Tools
	NumInst() int
	NumTotal() int
}
```


#### type IWorkspace

```go
type IWorkspace interface {
	IObjSnap
	json.Marshaler
	sync.Locker

	Dirs() WorkspaceDirs
	Files() WorkspaceFiles
	PrettyPath(string, ...string) string
	// contains filtered or unexported methods
}
```


#### type InfoTip

```go
type InfoTip struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}
```


#### type IpcIDs

```go
type IpcIDs uint8
```


```go
const (
	IPCID_MENUS_MAIN IpcIDs
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
```

#### func (IpcIDs) String

```go
func (me IpcIDs) String() (r string)
```
String implements the Go standard library's `fmt.Stringer` interface.

#### func (IpcIDs) Valid

```go
func (me IpcIDs) Valid() (r bool)
```
Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN`
(inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).

#### type ListBase

```go
type ListBase struct {
}
```


#### func (*ListBase) Count

```go
func (this *ListBase) Count(all ListFilters) int
```

#### func (*ListBase) FilterByID

```go
func (this *ListBase) FilterByID(id string) *ListFilter
```

#### type ListFilter

```go
type ListFilter struct {
	ID        string
	Disabled  bool
	Title     string
	Desc      string
	OnSrcLens func(*ListFilter, *SrcLens)
	Pred      ListItemPredicate
}
```


#### type ListFilters

```go
type ListFilters map[*ListFilter]bool
```


#### type ListItemPredicate

```go
type ListItemPredicate func(IListItem) bool
```


#### type ListItems

```go
type ListItems []IListItem
```


#### func (ListItems) Len

```go
func (this ListItems) Len() int
```

#### func (ListItems) Less

```go
func (this ListItems) Less(i int, j int) bool
```

#### func (ListItems) Swap

```go
func (this ListItems) Swap(i, j int)
```

#### type ListMenuBase

```go
type ListMenuBase struct {
	ListBase
}
```


#### func (*ListMenuBase) MenuCategory

```go
func (this *ListMenuBase) MenuCategory() string
```

#### type Menu

```go
type Menu struct {
	Desc     string    `json:"desc,omitempty"`
	TopLevel bool      `json:"topLevel,omitempty"`
	Items    MenuItems `json:"items"`
}
```


#### type MenuItem

```go
type MenuItem struct {
	IpcID    IpcIDs      `json:"ii,omitempty"`
	IpcArgs  interface{} `json:"ia,omitempty"`
	Category string      `json:"c,omitempty"`
	Title    string      `json:"t"`
	Desc     string      `json:"d,omitempty"`
	Hint     string      `json:"h,omitempty"`
	Confirm  string      `json:"q,omitempty"`
}
```


#### type MenuItems

```go
type MenuItems []*MenuItem
```


#### type PagesBase

```go
type PagesBase struct {
	Impl IPages
}
```


#### func (*PagesBase) Init

```go
func (*PagesBase) Init()
```

#### func (*PagesBase) PageBodyInnerHtml

```go
func (*PagesBase) PageBodyInnerHtml(rawUri string, path []string, query url.Values, fragment string) string
```

#### type PkgInfo

```go
type PkgInfo struct {
	Id        string
	ShortName string
	LongName  string
	Deps      func() PkgInfos
	Mems      func() []*PkgMemInfo
	Forget    func()
}
```


#### type PkgInfos

```go
type PkgInfos []*PkgInfo
```


#### func (*PkgInfos) Add

```go
func (this *PkgInfos) Add(pkg *PkgInfo)
```

#### func (PkgInfos) ById

```go
func (this PkgInfos) ById(id string) *PkgInfo
```

#### type PkgIntelBase

```go
type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel
}
```


#### func (*PkgIntelBase) Init

```go
func (this *PkgIntelBase) Init()
```

#### func (*PkgIntelBase) ObjSnapPrefix

```go
func (this *PkgIntelBase) ObjSnapPrefix() string
```

#### func (*PkgIntelBase) Pkgs

```go
func (this *PkgIntelBase) Pkgs() PkgInfos
```

#### func (*PkgIntelBase) PkgsAdd

```go
func (this *PkgIntelBase) PkgsAdd(pkg *PkgInfo)
```

#### type PkgMemInfo

```go
type PkgMemInfo struct {
	Kind Symbol
	Name string
	Desc string
	Subs func() []*PkgMemInfo
}
```


#### type Setting

```go
type Setting struct {
	Id         string
	Title      string
	Desc       string
	ValCfg     interface{}
	ValDef     interface{}
	OnChanging func(newVal interface{}) `json:"-"`
	OnChanged  func(oldVal interface{}) `json:"-"`
	OnReloaded func()                   `json:"-"`
}
```


#### func (*Setting) Val

```go
func (this *Setting) Val() interface{}
```

#### func (*Setting) ValBool

```go
func (this *Setting) ValBool() (val bool)
```

#### func (*Setting) ValInt

```go
func (this *Setting) ValInt() (val int64)
```

#### func (*Setting) ValStr

```go
func (this *Setting) ValStr() (val string)
```

#### func (*Setting) ValStrs

```go
func (this *Setting) ValStrs() (val []string)
```

#### func (*Setting) ValUInt

```go
func (this *Setting) ValUInt() (val uint64)
```

#### type Settings

```go
type Settings []*Setting
```


#### type SettingsBase

```go
type SettingsBase struct {
	Impl ISettings
}
```


#### func (*SettingsBase) Init

```go
func (this *SettingsBase) Init()
```

#### func (*SettingsBase) KnownSettings

```go
func (this *SettingsBase) KnownSettings() Settings
```

#### func (*SettingsBase) MenuCategory

```go
func (*SettingsBase) MenuCategory() string
```

#### type SrcFormattingClientPrefs

```go
type SrcFormattingClientPrefs struct {
	InsertSpaces *bool
	TabSize      *int
}
```


#### type SrcIntelBase

```go
type SrcIntelBase struct {
	Impl ISrcIntel
}
```


#### func (*SrcIntelBase) ComplDetails

```go
func (*SrcIntelBase) ComplDetails(*SrcLens, string) *SrcIntelCompl
```

#### func (*SrcIntelBase) ComplItems

```go
func (*SrcIntelBase) ComplItems(*SrcLens) SrcIntelCompls
```

#### func (*SrcIntelBase) ComplItemsShouldSort

```go
func (*SrcIntelBase) ComplItemsShouldSort(*SrcLens) bool
```

#### func (*SrcIntelBase) DefImpl

```go
func (*SrcIntelBase) DefImpl(*SrcLens) SrcLocs
```

#### func (*SrcIntelBase) DefSym

```go
func (*SrcIntelBase) DefSym(*SrcLens) SrcLocs
```

#### func (*SrcIntelBase) DefType

```go
func (*SrcIntelBase) DefType(*SrcLens) SrcLocs
```

#### func (*SrcIntelBase) Highlights

```go
func (*SrcIntelBase) Highlights(*SrcLens, string) SrcLocs
```

#### func (*SrcIntelBase) Hovers

```go
func (*SrcIntelBase) Hovers(*SrcLens) []InfoTip
```

#### func (*SrcIntelBase) Init

```go
func (*SrcIntelBase) Init()
```

#### func (*SrcIntelBase) References

```go
func (*SrcIntelBase) References(*SrcLens, bool) SrcLocs
```

#### func (*SrcIntelBase) Signature

```go
func (*SrcIntelBase) Signature(*SrcLens) *SrcIntelSigHelp
```

#### func (*SrcIntelBase) Symbols

```go
func (*SrcIntelBase) Symbols(*SrcLens, string, bool) SrcLenses
```

#### type SrcIntelCompl

```go
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
```


#### type SrcIntelCompls

```go
type SrcIntelCompls []*SrcIntelCompl
```


#### func (SrcIntelCompls) Len

```go
func (this SrcIntelCompls) Len() int
```

#### func (SrcIntelCompls) Less

```go
func (this SrcIntelCompls) Less(i int, j int) bool
```

#### func (SrcIntelCompls) Swap

```go
func (this SrcIntelCompls) Swap(i int, j int)
```

#### type SrcIntelDoc

```go
type SrcIntelDoc struct {
	Value     string `json:"value,omitempty"`
	IsTrusted bool   `json:"isTrusted,omitempty"`
}
```


#### type SrcIntelSigHelp

```go
type SrcIntelSigHelp struct {
	ActiveSignature int               `json:"activeSignature"`
	ActiveParameter int               `json:"activeParameter,omitempty"`
	Signatures      []SrcIntelSigInfo `json:"signatures,omitempty"`
}
```


#### type SrcIntelSigInfo

```go
type SrcIntelSigInfo struct {
	Label         string             `json:"label"`
	Documentation SrcIntelDoc        `json:"documentation,omitempty"`
	Parameters    []SrcIntelSigParam `json:"parameters"`
}
```


#### type SrcIntelSigParam

```go
type SrcIntelSigParam struct {
	Label         string      `json:"label"`
	Documentation SrcIntelDoc `json:"documentation,omitempty"`
}
```


#### type SrcIntels

```go
type SrcIntels struct {
	Info []InfoTip `json:",omitempty"`
	Refs SrcLocs   `json:",omitempty"`
}
```


#### type SrcLens

```go
type SrcLens struct {
	SrcLoc
	Txt  string `json:"t,omitempty"`
	Str  string `json:"s,omitempty"`
	CrLf bool   `json:"l,omitempty"`
}
```


#### func (*SrcLens) ByteOffsetForFirstLineBeginningWith

```go
func (this *SrcLens) ByteOffsetForFirstLineBeginningWith(prefix string) int
```

#### func (*SrcLens) ByteOffsetForPos

```go
func (this *SrcLens) ByteOffsetForPos(pos *SrcPos) int
```

#### func (*SrcLens) EnsureSrcFull

```go
func (this *SrcLens) EnsureSrcFull()
```

#### func (*SrcLens) Rune1OffsetForByte0Offset

```go
func (this *SrcLens) Rune1OffsetForByte0Offset(byte0off int) (rune1off int)
```

#### type SrcLenses

```go
type SrcLenses []*SrcLens
```


#### func (*SrcLenses) AddFrom

```go
func (this *SrcLenses) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (lens *SrcLens)
```

#### type SrcLoc

```go
type SrcLoc struct {
	Flag     int       `json:"e"` // don't omitempty
	FilePath string    `json:"f,omitempty"`
	Pos      *SrcPos   `json:"p,omitempty"`
	Range    *SrcRange `json:"r,omitempty"`
}
```


#### func (*SrcLoc) SetFilePathAndPosOrRangeFrom

```go
func (this *SrcLoc) SetFilePathAndPosOrRangeFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string)
```

#### type SrcLocs

```go
type SrcLocs []*SrcLoc
```


#### func (*SrcLocs) AddFrom

```go
func (this *SrcLocs) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (loc *SrcLoc)
```

#### type SrcModBase

```go
type SrcModBase struct {
	Impl ISrcMod
}
```


#### func (*SrcModBase) CodeActions

```go
func (*SrcModBase) CodeActions(srcLens *SrcLens) (all []EditorAction)
```

#### func (*SrcModBase) Init

```go
func (this *SrcModBase) Init()
```

#### func (*SrcModBase) MenuCategory

```go
func (*SrcModBase) MenuCategory() string
```

#### func (*SrcModBase) RunRenamer

```go
func (*SrcModBase) RunRenamer(srcLens *SrcLens, newName string) (all SrcLenses)
```

#### type SrcModEdits

```go
type SrcModEdits []srcModEdit
```


#### func (*SrcModEdits) AddDeleteLine

```go
func (this *SrcModEdits) AddDeleteLine(srcFilePath string, lineAt *SrcPos)
```

#### func (*SrcModEdits) AddInsert

```go
func (this *SrcModEdits) AddInsert(srcFilePath string, atPos func(*SrcLens, *SrcPos) string)
```

#### func (SrcModEdits) Len

```go
func (this SrcModEdits) Len() int
```

#### func (SrcModEdits) Less

```go
func (this SrcModEdits) Less(i int, j int) bool
```

#### func (SrcModEdits) Swap

```go
func (this SrcModEdits) Swap(i int, j int)
```

#### type SrcPos

```go
type SrcPos struct {
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
	Off int `json:"o,omitempty"`
}
```

All public fields are 1-based (so 0 means 'missing') and rune-not-byte-based

#### func (*SrcPos) String

```go
func (this *SrcPos) String() string
```

#### type SrcRange

```go
type SrcRange struct {
	Start SrcPos `json:"s"`
	End   SrcPos `json:"e,omitempty"`
}
```


#### type Symbol

```go
type Symbol uint8
```


```go
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
```

#### func (Symbol) String

```go
func (me Symbol) String() (r string)
```
String implements the Go standard library's `fmt.Stringer` interface.

#### type Tool

```go
type Tool struct {
	Cats      []ToolCats
	Name      string
	Installed bool
	Website   string
	DiagSev   DiagSeverity
}
```


#### func (*Tool) Exec

```go
func (*Tool) Exec(panicOnErr bool, stdin string, cmdName string, cmdArgs []string) (string, string)
```

#### func (*Tool) NotInstalledMessage

```go
func (this *Tool) NotInstalledMessage() string
```

#### type ToolCats

```go
type ToolCats uint8
```


```go
const (
	TOOLS_CAT_MOD_REN ToolCats
	TOOLS_CAT_MOD_FMT
	TOOLS_CAT_INTEL_TIPS
	TOOLS_CAT_INTEL_SYMS
	TOOLS_CAT_INTEL_HIGH
	TOOLS_CAT_INTEL_CMPL
	TOOLS_CAT_INTEL_NAV
	TOOLS_CAT_EXTRAS_QUERY
	TOOLS_CAT_DIAGS
	TOOLS_CAT_RUNONSAVE
)
```

#### func (ToolCats) String

```go
func (this ToolCats) String() string
```

#### type ToolingBase

```go
type ToolingBase struct {
	Impl ITooling
}
```


#### func (*ToolingBase) CountNumInst

```go
func (this *ToolingBase) CountNumInst(all Tools) (numInst int)
```

#### func (*ToolingBase) Init

```go
func (this *ToolingBase) Init()
```

#### func (*ToolingBase) KnownToolsFor

```go
func (this *ToolingBase) KnownToolsFor(cats ...ToolCats) (tools Tools)
```

#### func (*ToolingBase) MenuCategory

```go
func (this *ToolingBase) MenuCategory() string
```

#### type Tools

```go
type Tools []*Tool
```


#### type TreeViewItem

```go
type TreeViewItem struct {
	ID               string        `json:"id,omitempty"`
	Label            string        `json:"label,omitempty"`
	IconPath         string        `json:"iconPath,omitempty"`
	Tooltip          string        `json:"tooltip,omitempty"`
	Command          *EditorAction `json:"command,omitempty"`
	ContextValue     string        `json:"contextValue,omitempty"`
	CollapsibleState int           `json:"collapsibleState"`
}
```


#### type WorkspaceBase

```go
type WorkspaceBase struct {
	Impl IWorkspace `json:"-"`

	OnBeforeChanges WorkspaceChangesBefore `json:"-"`
	OnAfterChanges  WorkspaceChangesAfter  `json:"-"`
}
```


#### func (*WorkspaceBase) Dirs

```go
func (this *WorkspaceBase) Dirs() (dirs WorkspaceDirs)
```

#### func (*WorkspaceBase) Files

```go
func (this *WorkspaceBase) Files() (files WorkspaceFiles)
```

#### func (*WorkspaceBase) Init

```go
func (this *WorkspaceBase) Init()
```

#### func (*WorkspaceBase) Lock

```go
func (this *WorkspaceBase) Lock()
```

#### func (*WorkspaceBase) MarshalJSON

```go
func (this *WorkspaceBase) MarshalJSON() ([]byte, error)
```

#### func (*WorkspaceBase) ObjSnap

```go
func (this *WorkspaceBase) ObjSnap(string) interface{}
```

#### func (*WorkspaceBase) ObjSnapPrefix

```go
func (*WorkspaceBase) ObjSnapPrefix() string
```

#### func (*WorkspaceBase) PrettyPath

```go
func (this *WorkspaceBase) PrettyPath(fsPath string, otherEnvs ...string) string
```

#### func (*WorkspaceBase) Unlock

```go
func (this *WorkspaceBase) Unlock()
```

#### type WorkspaceChanges

```go
type WorkspaceChanges struct {
	AddedDirs    []string
	RemovedDirs  []string
	OpenedFiles  []string
	ClosedFiles  []string
	WrittenFiles []string
}
```


#### func (*WorkspaceChanges) HasDirChanges

```go
func (this *WorkspaceChanges) HasDirChanges() bool
```

#### type WorkspaceChangesAfter

```go
type WorkspaceChangesAfter func(upd *WorkspaceChanges)
```


#### type WorkspaceChangesBefore

```go
type WorkspaceChangesBefore func(upd *WorkspaceChanges, freshFiles []string, willAutoLint bool)
```


#### type WorkspaceDir

```go
type WorkspaceDir struct {
	Path string
}
```


#### type WorkspaceDirs

```go
type WorkspaceDirs map[string]*WorkspaceDir
```


#### type WorkspaceFile

```go
type WorkspaceFile struct {
	Path   string
	IsOpen bool `json:",omitempty"`
	Diags  struct {
		AutoLintUpToDate bool
		Build            diags
		Lint             diags
	}
}
```


#### type WorkspaceFiles

```go
type WorkspaceFiles map[string]*WorkspaceFile
```


#### func (WorkspaceFiles) HasBuildDiags

```go
func (this WorkspaceFiles) HasBuildDiags(filePath string) (has bool)
```
