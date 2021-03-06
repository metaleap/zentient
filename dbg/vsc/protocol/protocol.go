// Debug Adapter Protocol
// 
// The Debug Adapter Protocol defines the protocol used between an editor or IDE and a debugger or runtime.
// 
// Package codegen'd on 2019-06-01T20:46:00+02:00 from github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol/vscdbgprotocol.json with github.com/metaleap/zentient/cmd/zentient-dbg-vsc-genprotocol
package zdbgvscp
import "encoding/json"
import "errors"
import "strings"



// Arguments for 'setFunctionBreakpoints' request.
type SetFunctionBreakpointsArguments struct {

	// The function names of the breakpoints.
	Breakpoints []FunctionBreakpoint `json:"breakpoints"`

} // struct SetFunctionBreakpointsArguments

func (me *SetFunctionBreakpointsArguments) propagateFieldsToBase() {
}



// Arguments for 'disassemble' request.
type DisassembleArguments struct {

	// Memory reference to the base location containing the instructions to disassemble.
	MemoryReference string `json:"memoryReference"`

	// Optional offset (in bytes) to be applied to the reference location before disassembling. Can be negative.
	Offset int `json:"offset,omitempty"`

	// Optional offset (in instructions) to be applied after the byte offset (if any) before disassembling. Can be negative.
	InstructionOffset int `json:"instructionOffset,omitempty"`

	// Number of instructions to disassemble starting at the specified location and offset. An adapter must return exactly this number of instructions - any unavailable instructions should be replaced with an implementation-defined 'invalid instruction' value.
	InstructionCount int `json:"instructionCount"`

	// If true, the adapter should attempt to resolve memory addresses and other values to symbolic names.
	ResolveSymbols bool `json:"resolveSymbols,omitempty"`

} // struct DisassembleArguments

func (me *DisassembleArguments) propagateFieldsToBase() {
}



// The request returns a stacktrace from the current execution state.
type StackTraceRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command"`

	Arguments StackTraceArguments `json:"arguments"`

} // struct StackTraceRequest

func (me *StackTraceRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The request retrieves a list of all threads.
type ThreadsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `threads`
	Command string `json:"command"`

} // struct ThreadsRequest

func (me *ThreadsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// The request terminates the threads with the given ids.
type TerminateThreadsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `terminateThreads`
	Command string `json:"command"`

	Arguments TerminateThreadsArguments `json:"arguments"`

} // struct TerminateThreadsRequest

func (me *TerminateThreadsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'setDataBreakpoints' request.
type SetDataBreakpointsArguments struct {

	// The contents of this array replaces all existing data breakpoints. An empty array clears all data breakpoints.
	Breakpoints []DataBreakpoint `json:"breakpoints"`

} // struct SetDataBreakpointsArguments

func (me *SetDataBreakpointsArguments) propagateFieldsToBase() {
}



// The event indicates that a thread has started or exited.
type ThreadEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `thread`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `started`, `exited`
		Reason string `json:"reason"`

		// The identifier of the thread.
		ThreadId int `json:"threadId"`

	} `json:"body"`

} // struct ThreadEvent

func (me *ThreadEvent) propagateFieldsToBase() {
	me.Event.Body = me.Body
	me.Event.Event = me.Event_
	me.Event.propagateFieldsToBase()
}



// Information about the capabilities of a debug adapter.
type Capabilities struct {

	// The debug adapter supports the 'configurationDone' request.
	SupportsConfigurationDoneRequest bool `json:"supportsConfigurationDoneRequest,omitempty"`

	// The debug adapter supports breakpoints that break execution after a specified number of hits.
	SupportsHitConditionalBreakpoints bool `json:"supportsHitConditionalBreakpoints,omitempty"`

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

	// The debug adapter supports the 'modules' request.
	SupportsModulesRequest bool `json:"supportsModulesRequest,omitempty"`

	// The debug adapter supports the 'restart' request. In this case a client should not implement 'restart' by terminating and relaunching the adapter but by calling the RestartRequest.
	SupportsRestartRequest bool `json:"supportsRestartRequest,omitempty"`

	// The debug adapter supports the 'terminateThreads' request.
	SupportsTerminateThreadsRequest bool `json:"supportsTerminateThreadsRequest,omitempty"`

	// The debug adapter supports the 'gotoTargets' request.
	SupportsGotoTargetsRequest bool `json:"supportsGotoTargetsRequest,omitempty"`

	// The debug adapter supports a 'format' attribute on the stackTraceRequest, variablesRequest, and evaluateRequest.
	SupportsValueFormattingOptions bool `json:"supportsValueFormattingOptions,omitempty"`

	// The debug adapter supports the 'loadedSources' request.
	SupportsLoadedSourcesRequest bool `json:"supportsLoadedSourcesRequest,omitempty"`

	// The debug adapter supports the 'readMemory' request.
	SupportsReadMemoryRequest bool `json:"supportsReadMemoryRequest,omitempty"`

	// The debug adapter supports a (side effect free) evaluate request for data hovers.
	SupportsEvaluateForHovers bool `json:"supportsEvaluateForHovers,omitempty"`

	// The debug adapter supports setting a variable to a value.
	SupportsSetVariable bool `json:"supportsSetVariable,omitempty"`

	// The debug adapter supports restarting a frame.
	SupportsRestartFrame bool `json:"supportsRestartFrame,omitempty"`

	// The debug adapter supports the 'stepInTargets' request.
	SupportsStepInTargetsRequest bool `json:"supportsStepInTargetsRequest,omitempty"`

	// The debug adapter supports 'exceptionOptions' on the setExceptionBreakpoints request.
	SupportsExceptionOptions bool `json:"supportsExceptionOptions,omitempty"`

	// The debug adapter supports the delayed loading of parts of the stack, which requires that both the 'startFrame' and 'levels' arguments and the 'totalFrames' result of the 'StackTrace' request are supported.
	SupportsDelayedStackTraceLoading bool `json:"supportsDelayedStackTraceLoading,omitempty"`

	// The debug adapter supports stepping back via the 'stepBack' and 'reverseContinue' requests.
	SupportsStepBack bool `json:"supportsStepBack,omitempty"`

	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []ColumnDescriptor `json:"additionalModuleColumns,omitempty"`

	// The debug adapter supports logpoints by interpreting the 'logMessage' attribute of the SourceBreakpoint.
	SupportsLogPoints bool `json:"supportsLogPoints,omitempty"`

	// The debug adapter supports the 'disassemble' request.
	SupportsDisassembleRequest bool `json:"supportsDisassembleRequest,omitempty"`

	// The debug adapter supports the 'completions' request.
	SupportsCompletionsRequest bool `json:"supportsCompletionsRequest,omitempty"`

	// The debug adapter supports data breakpoints.
	SupportsDataBreakpoints bool `json:"supportsDataBreakpoints,omitempty"`

	// The debug adapter supports conditional breakpoints.
	SupportsConditionalBreakpoints bool `json:"supportsConditionalBreakpoints,omitempty"`

	// The debug adapter supports the 'setExpression' request.
	SupportsSetExpression bool `json:"supportsSetExpression,omitempty"`

	// The debug adapter supports the 'terminate' request.
	SupportsTerminateRequest bool `json:"supportsTerminateRequest,omitempty"`

	// The debug adapter supports function breakpoints.
	SupportsFunctionBreakpoints bool `json:"supportsFunctionBreakpoints,omitempty"`

	// Checksum algorithms supported by the debug adapter.
	SupportedChecksumAlgorithms []ChecksumAlgorithm `json:"supportedChecksumAlgorithms,omitempty"`

	// The debug adapter supports the 'exceptionInfo' request.
	SupportsExceptionInfoRequest bool `json:"supportsExceptionInfoRequest,omitempty"`

	// The debug adapter supports the 'terminateDebuggee' attribute on the 'disconnect' request.
	SupportTerminateDebuggee bool `json:"supportTerminateDebuggee,omitempty"`

} // struct Capabilities

func (me *Capabilities) propagateFieldsToBase() {
}



// This enumeration defines all possible conditions when a thrown exception should result in a break.
// never: never breaks,
// always: always breaks,
// unhandled: breaks when excpetion unhandled,
// userUnhandled: breaks if the exception is not handled by user code.
// 
// POSSIBLE VALUES: `never`, `always`, `unhandled`, `userUnhandled`
type ExceptionBreakMode string



// The attach request is sent from the client to the debug adapter to attach to a debuggee that is already running. Since attaching is debugger/runtime specific, the arguments for this request are not part of this specification.
type AttachRequest struct {
	// A client or debug adapter initiated request.
	Request

	Arguments AttachRequestArguments `json:"arguments"`

	// POSSIBLE VALUES: `attach`
	Command string `json:"command"`

} // struct AttachRequest

func (me *AttachRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Response to 'launch' request. This is just an acknowledgement, so no body field is required.
type LaunchResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `launch`
	Command string `json:"command,omitempty"`

} // struct LaunchResponse

func (me *LaunchResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The request starts the debuggee to run backward. Clients should only call this request if the capability 'supportsStepBack' is true.
type ReverseContinueRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `reverseContinue`
	Command string `json:"command"`

	Arguments ReverseContinueArguments `json:"arguments"`

} // struct ReverseContinueRequest

func (me *ReverseContinueRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The event indicates that some information about a breakpoint has changed.
type BreakpointEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `breakpoint`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `changed`, `new`, `removed`
		Reason string `json:"reason"`

		// The 'id' attribute is used to find the target breakpoint and the other attributes are used as the new values.
		Breakpoint Breakpoint `json:"breakpoint"`

	} `json:"body"`

} // struct BreakpointEvent

func (me *BreakpointEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Sets multiple breakpoints for a single source and clears all previous breakpoints in that source.
// To clear all breakpoint for a source, specify an empty array.
// When a breakpoint is hit, a 'stopped' event (with reason 'breakpoint') is generated.
type SetBreakpointsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `setBreakpoints`
	Command string `json:"command"`

	Arguments SetBreakpointsArguments `json:"arguments"`

} // struct SetBreakpointsRequest

func (me *SetBreakpointsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Some predefined types for the CompletionItem. Please note that not all clients have specific icons for all of them.
// 
// POSSIBLE VALUES: `method`, `function`, `constructor`, `field`, `variable`, `class`, `interface`, `module`, `property`, `unit`, `value`, `enum`, `keyword`, `snippet`, `text`, `color`, `file`, `reference`, `customcolor`
type CompletionItemType string



// Arguments for 'stepOut' request.
type StepOutArguments struct {

	// Execute 'stepOut' for this thread.
	ThreadId int `json:"threadId"`

} // struct StepOutArguments

func (me *StepOutArguments) propagateFieldsToBase() {
}



// A Module object represents a row in the modules view.
// Two attributes are mandatory: an id identifies a module in the modules view and is used in a ModuleEvent for identifying a module for adding, updating or deleting.
// The name is used to minimally render the module in the UI.
// 
// Additional attributes can be added to the module. They will show up in the module View if they have a corresponding ColumnDescriptor.
// 
// To avoid an unnecessary proliferation of additional attributes with similar semantics but different names
// we recommend to re-use attributes from the 'recommended' list below first, and only introduce new attributes if nothing appropriate could be found.
type Module struct {

	// optional but recommended attributes.
	// always try to use these first before introducing additional attributes.
	// 
	// Logical full path to the module. The exact definition is implementation defined, but usually this would be a full path to the on-disk file for the module.
	Path string `json:"path,omitempty"`

	// Version of Module.
	Version string `json:"version,omitempty"`

	// User understandable description of if symbols were found for the module (ex: 'Symbols Loaded', 'Symbols not found', etc.
	SymbolStatus string `json:"symbolStatus,omitempty"`

	// Logical full path to the symbol file. The exact definition is implementation defined.
	SymbolFilePath string `json:"symbolFilePath,omitempty"`

	// Module created or modified.
	DateTimeStamp string `json:"dateTimeStamp,omitempty"`

	// Address range covered by this module.
	AddressRange string `json:"addressRange,omitempty"`

	// Unique identifier for the module.
	// 
	// POSSIBLE TYPES:
	// - `int` (for JSON `integer`s)
	// - `string` (for JSON `string`s)
	Id interface{} `json:"id"`

	// A name of the module.
	Name string `json:"name"`

	// True if the module is optimized.
	IsOptimized bool `json:"isOptimized,omitempty"`

	// True if the module is considered 'user code' by a debugger that supports 'Just My Code'.
	IsUserCode bool `json:"isUserCode,omitempty"`

} // struct Module

func (me *Module) propagateFieldsToBase() {
}



// Response to 'scopes' request.
type ScopesResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The scopes of the stackframe. If the array has length zero, there are no scopes available.
		Scopes []Scope `json:"scopes"`

	} `json:"body"`

	// POSSIBLE VALUES: `scopes`
	Command string `json:"command,omitempty"`

} // struct ScopesResponse

func (me *ScopesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The request configures the debuggers response to thrown exceptions. If an exception is configured to break, a 'stopped' event is fired (with reason 'exception').
type SetExceptionBreakpointsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command"`

	Arguments SetExceptionBreakpointsArguments `json:"arguments"`

} // struct SetExceptionBreakpointsRequest

func (me *SetExceptionBreakpointsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'setExpression' request.
type SetExpressionResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The new value of the expression.
		Value string `json:"value"`

		// The optional type of the value.
		Type string `json:"type,omitempty"`

		// Properties of a value that can be used to determine how to render the result in the UI.
		PresentationHint VariablePresentationHint `json:"presentationHint,omitempty"`

		// If variablesReference is > 0, the value is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference,omitempty"`

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `setExpression`
	Command string `json:"command,omitempty"`

} // struct SetExpressionResponse

func (me *SetExpressionResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'modules' request.
type ModulesArguments struct {

	// The index of the first module to return; if omitted modules start at 0.
	StartModule int `json:"startModule,omitempty"`

	// The number of modules to return. If moduleCount is not specified or 0, all modules are returned.
	ModuleCount int `json:"moduleCount,omitempty"`

} // struct ModulesArguments

func (me *ModulesArguments) propagateFieldsToBase() {
}



// Arguments for 'stepIn' request.
type StepInArguments struct {

	// Execute 'stepIn' for this thread.
	ThreadId int `json:"threadId"`

	// Optional id of the target to step into.
	TargetId int `json:"targetId,omitempty"`

} // struct StepInArguments

func (me *StepInArguments) propagateFieldsToBase() {
}



// Response to 'dataBreakpointInfo' request.
type DataBreakpointInfoResponse struct {
	// Response for a request.
	Response

	Body struct {

		// An identifier for the data on which a data breakpoint can be registered with the setDataBreakpoints request or null if no data breakpoint is available.
		// 
		// POSSIBLE TYPES:
		// - `string` (for JSON `string`s)
		// - `interface{/*nil*/}` (for JSON `null`s)
		DataId interface{} `json:"dataId"`

		// UI string that describes on what data the breakpoint is set on or why a data breakpoint is not available.
		Description string `json:"description"`

		// Optional attribute listing the available access types for a potential data breakpoint. A UI frontend could surface this information.
		AccessTypes []DataBreakpointAccessType `json:"accessTypes,omitempty"`

		// Optional attribute indicating that a potential data breakpoint could be persisted across sessions.
		CanPersist bool `json:"canPersist,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `dataBreakpointInfo`
	Command string `json:"command,omitempty"`

} // struct DataBreakpointInfoResponse

func (me *DataBreakpointInfoResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The event indicates that the debuggee has exited and returns its exit code.
type ExitedEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `exited`
	Event_ string `json:"event"`

	Body struct {

		// The exit code returned from the debuggee.
		ExitCode int `json:"exitCode"`

	} `json:"body"`

} // struct ExitedEvent

func (me *ExitedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Response to 'stepBack' request. This is just an acknowledgement, so no body field is required.
type StepBackResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `stepBack`
	Command string `json:"command,omitempty"`

} // struct StepBackResponse

func (me *StepBackResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'exceptionInfo' request.
type ExceptionInfoArguments struct {

	// Thread for which exception information should be retrieved.
	ThreadId int `json:"threadId"`

} // struct ExceptionInfoArguments

func (me *ExceptionInfoArguments) propagateFieldsToBase() {
}



// The request starts the debuggee to run again.
type ContinueRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `continue`
	Command string `json:"command"`

	Arguments ContinueArguments `json:"arguments"`

} // struct ContinueRequest

func (me *ContinueRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Arguments for 'disconnect' request.
type DisconnectArguments struct {

	// A value of true indicates that this 'disconnect' request is part of a restart sequence.
	Restart bool `json:"restart,omitempty"`

	// Indicates whether the debuggee should be terminated when the debugger is disconnected.
	// If unspecified, the debug adapter is free to do whatever it thinks is best.
	// A client can only rely on this attribute being properly honored if a debug adapter returns true for the 'supportTerminateDebuggee' capability.
	TerminateDebuggee bool `json:"terminateDebuggee,omitempty"`

} // struct DisconnectArguments

func (me *DisconnectArguments) propagateFieldsToBase() {
}



// This event indicates that the debug adapter is ready to accept configuration requests (e.g. SetBreakpointsRequest, SetExceptionBreakpointsRequest).
// A debug adapter is expected to send this event when it is ready to accept configuration requests (but not before the 'initialize' request has finished).
// The sequence of events/requests is as follows:
// - adapters sends 'initialized' event (after the 'initialize' request has returned)
// - frontend sends zero or more 'setBreakpoints' requests
// - frontend sends one 'setFunctionBreakpoints' request
// - frontend sends a 'setExceptionBreakpoints' request if one or more 'exceptionBreakpointFilters' have been defined (or if 'supportsConfigurationDoneRequest' is not defined or false)
// - frontend sends other future configuration requests
// - frontend sends one 'configurationDone' request to indicate the end of the configuration.
type InitializedEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `initialized`
	Event_ string `json:"event"`

} // struct InitializedEvent

func (me *InitializedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.propagateFieldsToBase()
}



// Response to 'stackTrace' request.
type StackTraceResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The frames of the stackframe. If the array has length zero, there are no stackframes available.
		// This means that there is no location information available.
		StackFrames []StackFrame `json:"stackFrames"`

		// The total number of frames available.
		TotalFrames int `json:"totalFrames,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command,omitempty"`

} // struct StackTraceResponse

func (me *StackTraceResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Provides formatting information for a stack frame.
type StackFrameFormat struct {
	// Provides formatting information for a value.
	ValueFormat

	// Displays the names of parameters for the stack frame.
	ParameterNames bool `json:"parameterNames,omitempty"`

	// Displays the values of parameters for the stack frame.
	ParameterValues bool `json:"parameterValues,omitempty"`

	// Displays the line number of the stack frame.
	Line bool `json:"line,omitempty"`

	// Displays the module of the stack frame.
	Module bool `json:"module,omitempty"`

	// Includes all stack frames, including those the debug adapter might otherwise hide.
	IncludeAll bool `json:"includeAll,omitempty"`

	// Displays parameters for the stack frame.
	Parameters bool `json:"parameters,omitempty"`

	// Displays the types of parameters for the stack frame.
	ParameterTypes bool `json:"parameterTypes,omitempty"`

} // struct StackFrameFormat

func (me *StackFrameFormat) propagateFieldsToBase() {
	me.ValueFormat.propagateFieldsToBase()
}



// Arguments for 'variables' request.
type VariablesArguments struct {

	// Specifies details on how to format the Variable values.
	Format ValueFormat `json:"format,omitempty"`

	// The Variable reference.
	VariablesReference int `json:"variablesReference"`

	// Optional filter to limit the child variables to either named or indexed. If ommited, both types are fetched.
	// 
	// POSSIBLE VALUES: `indexed`, `named`
	Filter string `json:"filter,omitempty"`

	// The index of the first variable to return; if omitted children start at 0.
	Start int `json:"start,omitempty"`

	// The number of variables to return. If count is missing or 0, all variables are returned.
	Count int `json:"count,omitempty"`

} // struct VariablesArguments

func (me *VariablesArguments) propagateFieldsToBase() {
}



// Response to 'initialize' request.
type InitializeResponse struct {
	// Response for a request.
	Response

	// The capabilities of this debug adapter.
	Body Capabilities `json:"body,omitempty"`

	// POSSIBLE VALUES: `initialize`
	Command string `json:"command,omitempty"`

} // struct InitializeResponse

func (me *InitializeResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'setDataBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetDataBreakpointsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// Information about the data breakpoints. The array elements correspond to the elements of the input argument 'breakpoints' array.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`

	// POSSIBLE VALUES: `setDataBreakpoints`
	Command string `json:"command,omitempty"`

} // struct SetDataBreakpointsResponse

func (me *SetDataBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// A client or debug adapter initiated request.
type Request struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

	// POSSIBLE VALUES: `request`
	Type string `json:"type"`

	// The command to execute.
	Command string `json:"command"`

	// Object containing arguments for the command.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	Arguments interface{} `json:"arguments,omitempty"`

} // struct Request

func (me *Request) propagateFieldsToBase() {
	me.ProtocolMessage.Type = me.Type
	me.ProtocolMessage.propagateFieldsToBase()
}



// Arguments for 'gotoTargets' request.
type GotoTargetsArguments struct {

	// The source location for which the goto targets are determined.
	Source Source `json:"source"`

	// The line location for which the goto targets are determined.
	Line int `json:"line"`

	// An optional column location for which the goto targets are determined.
	Column int `json:"column,omitempty"`

} // struct GotoTargetsArguments

func (me *GotoTargetsArguments) propagateFieldsToBase() {
}



// Names of checksum algorithms that may be supported by a debug adapter.
// 
// POSSIBLE VALUES: `MD5`, `SHA1`, `SHA256`, `timestamp`
type ChecksumAlgorithm string



// Response to 'setFunctionBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetFunctionBreakpointsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements correspond to the elements of the 'breakpoints' array.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`

	// POSSIBLE VALUES: `setFunctionBreakpoints`
	Command string `json:"command,omitempty"`

} // struct SetFunctionBreakpointsResponse

func (me *SetFunctionBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'loadedSources' request.
type LoadedSourcesArguments map[string]interface{}



// Arguments for 'goto' request.
type GotoArguments struct {

	// The location where the debuggee will continue to run.
	TargetId int `json:"targetId"`

	// Set the goto target for this thread.
	ThreadId int `json:"threadId"`

} // struct GotoArguments

func (me *GotoArguments) propagateFieldsToBase() {
}



// CompletionItems are the suggestions returned from the CompletionsRequest.
type CompletionItem struct {

	// The label of this completion item. By default this is also the text that is inserted when selecting this completion.
	Label string `json:"label"`

	// If text is not falsy then it is inserted instead of the label.
	Text string `json:"text,omitempty"`

	// The item's type. Typically the client uses this information to render the item in the UI with an icon.
	Type CompletionItemType `json:"type,omitempty"`

	// This value determines the location (in the CompletionsRequest's 'text' attribute) where the completion text is added.
	// If missing the text is added at the location specified by the CompletionsRequest's 'column' attribute.
	Start int `json:"start,omitempty"`

	// This value determines how many characters are overwritten by the completion text.
	// If missing the value 0 is assumed which results in the completion text being inserted.
	Length int `json:"length,omitempty"`

} // struct CompletionItem

func (me *CompletionItem) propagateFieldsToBase() {
}



// A debug adapter initiated event.
type Event struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

	// POSSIBLE VALUES: `event`
	Type string `json:"type"`

	// Type of event.
	Event string `json:"event"`

	// Event-specific information.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	Body interface{} `json:"body,omitempty"`

} // struct Event

func (me *Event) propagateFieldsToBase() {
	me.ProtocolMessage.Type = me.Type
	me.ProtocolMessage.propagateFieldsToBase()
}



// Arguments for 'attach' request. Additional attributes are implementation specific.
type AttachRequestArguments struct {

	// Optional data from the previous, restarted session.
	// The data is sent as the 'restart' attribute of the 'terminated' event.
	// The client should leave the data intact.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	Restart__ interface{} `json:"__restart,omitempty"`

} // struct AttachRequestArguments

func (me *AttachRequestArguments) propagateFieldsToBase() {
}



// Arguments for 'setExceptionBreakpoints' request.
type SetExceptionBreakpointsArguments struct {

	// IDs of checked exception options. The set of IDs is returned via the 'exceptionBreakpointFilters' capability.
	Filters []string `json:"filters"`

	// Configuration options for selected exceptions.
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`

} // struct SetExceptionBreakpointsArguments

func (me *SetExceptionBreakpointsArguments) propagateFieldsToBase() {
}



// Arguments for 'restartFrame' request.
type RestartFrameArguments struct {

	// Restart this stackframe.
	FrameId int `json:"frameId"`

} // struct RestartFrameArguments

func (me *RestartFrameArguments) propagateFieldsToBase() {
}



// Detailed information about an exception that has occurred.
type ExceptionDetails struct {

	// Message contained in the exception.
	Message string `json:"message,omitempty"`

	// Short type name of the exception object.
	TypeName string `json:"typeName,omitempty"`

	// Fully-qualified type name of the exception object.
	FullTypeName string `json:"fullTypeName,omitempty"`

	// Optional expression that can be evaluated in the current scope to obtain the exception object.
	EvaluateName string `json:"evaluateName,omitempty"`

	// Stack trace at the time the exception was thrown.
	StackTrace string `json:"stackTrace,omitempty"`

	// Details of the exception contained by this exception, if any.
	InnerException []ExceptionDetails `json:"innerException,omitempty"`

} // struct ExceptionDetails

func (me *ExceptionDetails) propagateFieldsToBase() {
}



// The request starts the debuggee to run one step backwards.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed. Clients should only call this request if the capability 'supportsStepBack' is true.
type StepBackRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `stepBack`
	Command string `json:"command"`

	Arguments StepBackArguments `json:"arguments"`

} // struct StepBackRequest

func (me *StepBackRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// A structured message object. Used to return errors from requests.
type Message struct {

	// If true send to telemetry.
	SendTelemetry bool `json:"sendTelemetry,omitempty"`

	// If true show user.
	ShowUser bool `json:"showUser,omitempty"`

	// An optional url where additional information about this message can be found.
	Url string `json:"url,omitempty"`

	// An optional label that is presented to the user as the UI for opening the url.
	UrlLabel string `json:"urlLabel,omitempty"`

	// Unique identifier for the message.
	Id int `json:"id"`

	// A format string for the message. Embedded variables have the form '{name}'.
	// If variable name starts with an underscore character, the variable does not contain user data (PII) and can be safely used for telemetry purposes.
	Format string `json:"format"`

	// An object used as a dictionary for looking up the variables in the format string.
	Variables map[string]string `json:"variables,omitempty"`

} // struct Message

func (me *Message) propagateFieldsToBase() {
}



// Retrieves the details of the exception that caused this event to be raised.
type ExceptionInfoRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `exceptionInfo`
	Command string `json:"command"`

	Arguments ExceptionInfoArguments `json:"arguments"`

} // struct ExceptionInfoRequest

func (me *ExceptionInfoRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'goto' request. This is just an acknowledgement, so no body field is required.
type GotoResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `goto`
	Command string `json:"command,omitempty"`

} // struct GotoResponse

func (me *GotoResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'setExceptionBreakpoints' request. This is just an acknowledgement, so no body field is required.
type SetExceptionBreakpointsResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command,omitempty"`

} // struct SetExceptionBreakpointsResponse

func (me *SetExceptionBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Properties of a data breakpoint passed to the setDataBreakpoints request.
type DataBreakpoint struct {

	// An id representing the data. This id is returned from the dataBreakpointInfo request.
	DataId string `json:"dataId"`

	// The access type of the data.
	AccessType DataBreakpointAccessType `json:"accessType,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

} // struct DataBreakpoint

func (me *DataBreakpoint) propagateFieldsToBase() {
}



// A ColumnDescriptor specifies what module attribute to show in a column of the ModulesView, how to format it, and what the column's label should be.
// It is only used if the underlying UI actually supports this level of customization.
type ColumnDescriptor struct {

	// Name of the attribute rendered in this column.
	AttributeName string `json:"attributeName"`

	// Header UI label of column.
	Label string `json:"label"`

	// Format to use for the rendered values in this column. TBD how the format strings looks like.
	Format string `json:"format,omitempty"`

	// Datatype of values in this column.  Defaults to 'string' if not specified.
	// 
	// POSSIBLE VALUES: `string`, `number`, `boolean`, `unixTimestampUTC`
	Type string `json:"type,omitempty"`

	// Width of this column in characters (hint only).
	Width int `json:"width,omitempty"`

} // struct ColumnDescriptor

func (me *ColumnDescriptor) propagateFieldsToBase() {
}



// Response to 'terminateThreads' request. This is just an acknowledgement, so no body field is required.
type TerminateThreadsResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `terminateThreads`
	Command string `json:"command,omitempty"`

} // struct TerminateThreadsResponse

func (me *TerminateThreadsResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'gotoTargets' request.
type GotoTargetsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The possible goto targets of the specified location.
		Targets []GotoTarget `json:"targets"`

	} `json:"body"`

	// POSSIBLE VALUES: `gotoTargets`
	Command string `json:"command,omitempty"`

} // struct GotoTargetsResponse

func (me *GotoTargetsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The event indicates that one or more capabilities have changed.
// Since the capabilities are dependent on the frontend and its UI, it might not be possible to change that at random times (or too late).
// Consequently this event has a hint characteristic: a frontend can only be expected to make a 'best effort' in honouring individual capabilities but there are no guarantees.
// Only changed capabilities need to be included, all other capabilities keep their values.
type CapabilitiesEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `capabilities`
	Event_ string `json:"event"`

	Body struct {

		// The set of updated capabilities.
		Capabilities Capabilities `json:"capabilities"`

	} `json:"body"`

} // struct CapabilitiesEvent

func (me *CapabilitiesEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// The event indicates that the execution of the debuggee has stopped due to some condition.
// This can be caused by a break point previously set, a stepping action has completed, by executing a debugger statement etc.
type StoppedEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `stopped`
	Event_ string `json:"event"`

	Body struct {

		// If 'allThreadsStopped' is true, a debug adapter can announce that all threads have stopped.
		// - The client should use this information to enable that all threads can be expanded to access their stacktraces.
		// - If the attribute is missing or false, only the thread with the given threadId can be expanded.
		AllThreadsStopped bool `json:"allThreadsStopped,omitempty"`

		// The reason for the event.
		// For backward compatibility this string is shown in the UI if the 'description' attribute is missing (but it must not be translated).
		// 
		// POSSIBLE VALUES: `step`, `breakpoint`, `exception`, `pause`, `entry`, `goto`, `function breakpoint`, `data breakpoint`
		Reason string `json:"reason"`

		// The full reason for the event, e.g. 'Paused on exception'. This string is shown in the UI as is and must be translated.
		Description string `json:"description,omitempty"`

		// The thread which was stopped.
		ThreadId int `json:"threadId,omitempty"`

		// A value of true hints to the frontend that this event should not change the focus.
		PreserveFocusHint bool `json:"preserveFocusHint,omitempty"`

		// Additional information. E.g. if reason is 'exception', text contains the exception name. This string is shown in the UI.
		Text string `json:"text,omitempty"`

	} `json:"body"`

} // struct StoppedEvent

func (me *StoppedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Response to 'pause' request. This is just an acknowledgement, so no body field is required.
type PauseResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `pause`
	Command string `json:"command,omitempty"`

} // struct PauseResponse

func (me *PauseResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// This request retrieves the possible stepIn targets for the specified stack frame.
// These targets can be used in the 'stepIn' request.
// The StepInTargets may only be called if the 'supportsStepInTargetsRequest' capability exists and is true.
type StepInTargetsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `stepInTargets`
	Command string `json:"command"`

	Arguments StepInTargetsArguments `json:"arguments"`

} // struct StepInTargetsRequest

func (me *StepInTargetsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// A Variable is a name/value pair.
// Optionally a variable can have a 'type' that is shown if space permits or when hovering over the variable's name.
// An optional 'kind' is used to render additional properties of the variable, e.g. different icons can be used to indicate that a variable is public or private.
// If the value is structured (has children), a handle is provided to retrieve the children with the VariablesRequest.
// If the number of named or indexed children is large, the numbers should be returned via the optional 'namedVariables' and 'indexedVariables' attributes.
// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
type Variable struct {

	// Optional memory reference for the variable if the variable represents executable code, such as a function pointer.
	MemoryReference string `json:"memoryReference,omitempty"`

	// The type of the variable's value. Typically shown in the UI when hovering over the value.
	Type string `json:"type,omitempty"`

	// If variablesReference is > 0, the variable is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

	// Properties of a variable that can be used to determine how to render the variable in the UI.
	PresentationHint VariablePresentationHint `json:"presentationHint,omitempty"`

	// Optional evaluatable name of this variable which can be passed to the 'EvaluateRequest' to fetch the variable's value.
	EvaluateName string `json:"evaluateName,omitempty"`

	// The number of named child variables.
	// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
	NamedVariables int `json:"namedVariables,omitempty"`

	// The number of indexed child variables.
	// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
	IndexedVariables int `json:"indexedVariables,omitempty"`

	// The variable's name.
	Name string `json:"name"`

	// The variable's value. This can be a multi-line text, e.g. for a function the body of a function.
	Value string `json:"value"`

} // struct Variable

func (me *Variable) propagateFieldsToBase() {
}



// An ExceptionOptions assigns configuration options to a set of exceptions.
type ExceptionOptions struct {

	// A path that selects a single or multiple exceptions in a tree. If 'path' is missing, the whole tree is selected. By convention the first segment of the path is a category that is used to group exceptions in the UI.
	Path []ExceptionPathSegment `json:"path,omitempty"`

	// Condition when a thrown exception should result in a break.
	BreakMode ExceptionBreakMode `json:"breakMode"`

} // struct ExceptionOptions

func (me *ExceptionOptions) propagateFieldsToBase() {
}



// Arguments for 'runInTerminal' request.
type RunInTerminalRequestArguments struct {

	// What kind of terminal to launch.
	// 
	// POSSIBLE VALUES: `integrated`, `external`
	Kind string `json:"kind,omitempty"`

	// Optional title of the terminal.
	Title string `json:"title,omitempty"`

	// Working directory of the command.
	Cwd string `json:"cwd"`

	// List of arguments. The first argument is the command to run.
	Args []string `json:"args"`

	// Environment key-value pairs that are added to or removed from the default environment.
	Env map[string]interface{} `json:"env,omitempty"`

} // struct RunInTerminalRequestArguments

func (me *RunInTerminalRequestArguments) propagateFieldsToBase() {
}



// Response to 'readMemory' request.
type ReadMemoryResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The address of the first byte of data returned. Treated as a hex value if prefixed with '0x', or as a decimal value otherwise.
		Address string `json:"address"`

		// The number of unreadable bytes encountered after the last successfully read byte. This can be used to determine the number of bytes that must be skipped before a subsequent 'readMemory' request will succeed.
		UnreadableBytes int `json:"unreadableBytes,omitempty"`

		// The bytes read from memory, encoded using base64.
		Data string `json:"data,omitempty"`

	} `json:"body,omitempty"`

	// POSSIBLE VALUES: `readMemory`
	Command string `json:"command,omitempty"`

} // struct ReadMemoryResponse

func (me *ReadMemoryResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'dataBreakpointInfo' request.
type DataBreakpointInfoArguments struct {

	// Reference to the Variable container if the data breakpoint is requested for a child of the container.
	VariablesReference int `json:"variablesReference,omitempty"`

	// The name of the Variable's child to obtain data breakpoint information for. If variableReference isn’t provided, this can be an expression.
	Name string `json:"name"`

} // struct DataBreakpointInfoArguments

func (me *DataBreakpointInfoArguments) propagateFieldsToBase() {
}



// Arguments for 'initialize' request.
type InitializeRequestArguments struct {

	// Determines in what format paths are specified. The default is 'path', which is the native format.
	// 
	// POSSIBLE VALUES: `path`, `uri`
	PathFormat string `json:"pathFormat,omitempty"`

	// Client supports the optional type attribute for variables.
	SupportsVariableType bool `json:"supportsVariableType,omitempty"`

	// Client supports the paging of variables.
	SupportsVariablePaging bool `json:"supportsVariablePaging,omitempty"`

	// The ID of the debug adapter.
	AdapterID string `json:"adapterID"`

	// If true all line numbers are 1-based (default).
	LinesStartAt1 bool `json:"linesStartAt1,omitempty"`

	// The ISO-639 locale of the (frontend) client using this adapter, e.g. en-US or de-CH.
	Locale string `json:"locale,omitempty"`

	// If true all column numbers are 1-based (default).
	ColumnsStartAt1 bool `json:"columnsStartAt1,omitempty"`

	// Client supports the runInTerminal request.
	SupportsRunInTerminalRequest bool `json:"supportsRunInTerminalRequest,omitempty"`

	// Client supports memory references.
	SupportsMemoryReferences bool `json:"supportsMemoryReferences,omitempty"`

	// The ID of the (frontend) client using this adapter.
	ClientID string `json:"clientID,omitempty"`

	// The human readable name of the (frontend) client using this adapter.
	ClientName string `json:"clientName,omitempty"`

} // struct InitializeRequestArguments

func (me *InitializeRequestArguments) propagateFieldsToBase() {
}



// The request restarts execution of the specified stackframe.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'restart') after the restart has completed.
type RestartFrameRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `restartFrame`
	Command string `json:"command"`

	Arguments RestartFrameArguments `json:"arguments"`

} // struct RestartFrameRequest

func (me *RestartFrameRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// This request is sent from the debug adapter to the client to run a command in a terminal. This is typically used to launch the debuggee in a terminal provided by the client.
type RunInTerminalRequest struct {
	// A client or debug adapter initiated request.
	Request

	Arguments RunInTerminalRequestArguments `json:"arguments"`

	// POSSIBLE VALUES: `runInTerminal`
	Command string `json:"command"`

} // struct RunInTerminalRequest

func (me *RunInTerminalRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Returns a list of possible completions for a given caret position and text.
// The CompletionsRequest may only be called if the 'supportsCompletionsRequest' capability exists and is true.
type CompletionsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `completions`
	Command string `json:"command"`

	Arguments CompletionsArguments `json:"arguments"`

} // struct CompletionsRequest

func (me *CompletionsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The event indicates that the target has produced some output.
type OutputEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `output`
	Event_ string `json:"event"`

	Body struct {

		// The output category. If not specified, 'console' is assumed.
		// 
		// POSSIBLE VALUES: `console`, `stdout`, `stderr`, `telemetry`
		Category string `json:"category,omitempty"`

		// The output to report.
		Output string `json:"output"`

		// If an attribute 'variablesReference' exists and its value is > 0, the output contains objects which can be retrieved by passing 'variablesReference' to the 'variables' request.
		VariablesReference int64 `json:"variablesReference,omitempty"`

		// An optional source location where the output was produced.
		Source Source `json:"source,omitempty"`

		// An optional source location line where the output was produced.
		Line int `json:"line,omitempty"`

		// An optional source location column where the output was produced.
		Column int `json:"column,omitempty"`

		// Optional data to report. For the 'telemetry' category the data will be sent to telemetry, for the other categories the data is shown in JSON format.
		// 
		// POSSIBLE TYPES:
		// - `[]interface{}` (for JSON `array`s)
		// - `bool` (for JSON `boolean`s)
		// - `int` (for JSON `integer`s)
		// - `interface{/*nil*/}` (for JSON `null`s)
		// - `int64` (for JSON `number`s)
		// - `map[string]interface{}` (for JSON `object`s)
		// - `string` (for JSON `string`s)
		Data interface{} `json:"data,omitempty"`

	} `json:"body"`

} // struct OutputEvent

func (me *OutputEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Modules can be retrieved from the debug adapter with the ModulesRequest which can either return all modules or a range of modules to support paging.
type ModulesRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `modules`
	Command string `json:"command"`

	Arguments ModulesArguments `json:"arguments"`

} // struct ModulesRequest

func (me *ModulesRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'loadedSources' request.
type LoadedSourcesResponse struct {
	// Response for a request.
	Response

	Body struct {

		// Set of loaded sources.
		Sources []Source `json:"sources"`

	} `json:"body"`

	// POSSIBLE VALUES: `loadedSources`
	Command string `json:"command,omitempty"`

} // struct LoadedSourcesResponse

func (me *LoadedSourcesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The event indicates that some information about a module has changed.
type ModuleEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `module`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `new`, `changed`, `removed`
		Reason string `json:"reason"`

		// The new, changed, or removed module. In case of 'removed' only the module id is used.
		Module Module `json:"module"`

	} `json:"body"`

} // struct ModuleEvent

func (me *ModuleEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Response to 'continue' request.
type ContinueResponse struct {
	// Response for a request.
	Response

	Body struct {

		// If true, the 'continue' request has ignored the specified thread and continued all threads instead. If this attribute is missing a value of 'true' is assumed for backward compatibility.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `continue`
	Command string `json:"command,omitempty"`

} // struct ContinueResponse

func (me *ContinueResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// The request retrieves the source code for a given source reference.
type SourceRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `source`
	Command string `json:"command"`

	Arguments SourceArguments `json:"arguments"`

} // struct SourceRequest

func (me *SourceRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'source' request.
type SourceArguments struct {

	// Specifies the source content to load. Either source.path or source.sourceReference must be specified.
	Source Source `json:"source,omitempty"`

	// The reference to the source. This is the same as source.sourceReference. This is provided for backward compatibility since old backends do not understand the 'source' attribute.
	SourceReference int `json:"sourceReference"`

} // struct SourceArguments

func (me *SourceArguments) propagateFieldsToBase() {
}



// Information about a Breakpoint created in setBreakpoints or setFunctionBreakpoints.
type Breakpoint struct {

	// An optional end column of the actual range covered by the breakpoint. If no end line is given, then the end column is assumed to be in the start line.
	EndColumn int `json:"endColumn,omitempty"`

	// An optional identifier for the breakpoint. It is needed if breakpoint events are used to update or remove breakpoints.
	Id int `json:"id,omitempty"`

	// If true breakpoint could be set (but not necessarily at the desired location).
	Verified bool `json:"verified"`

	// An optional message about the state of the breakpoint. This is shown to the user and can be used to explain why a breakpoint could not be verified.
	Message string `json:"message,omitempty"`

	// The source where the breakpoint is located.
	Source Source `json:"source,omitempty"`

	// The start line of the actual range covered by the breakpoint.
	Line int `json:"line,omitempty"`

	// An optional start column of the actual range covered by the breakpoint.
	Column int `json:"column,omitempty"`

	// An optional end line of the actual range covered by the breakpoint.
	EndLine int `json:"endLine,omitempty"`

} // struct Breakpoint

func (me *Breakpoint) propagateFieldsToBase() {
}



// The request suspenses the debuggee.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'pause') after the thread has been paused successfully.
type PauseRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `pause`
	Command string `json:"command"`

	Arguments PauseArguments `json:"arguments"`

} // struct PauseRequest

func (me *PauseRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The client of the debug protocol must send this request at the end of the sequence of configuration requests (which was started by the 'initialized' event).
type ConfigurationDoneRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `configurationDone`
	Command string `json:"command"`

	Arguments ConfigurationDoneArguments `json:"arguments,omitempty"`

} // struct ConfigurationDoneRequest

func (me *ConfigurationDoneRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'stackTrace' request.
type StackTraceArguments struct {

	// Specifies details on how to format the stack frames.
	Format StackFrameFormat `json:"format,omitempty"`

	// Retrieve the stacktrace for this thread.
	ThreadId int `json:"threadId"`

	// The index of the first frame to return; if omitted frames start at 0.
	StartFrame int `json:"startFrame,omitempty"`

	// The maximum number of frames to return. If levels is not specified or 0, all frames are returned.
	Levels int `json:"levels,omitempty"`

} // struct StackTraceArguments

func (me *StackTraceArguments) propagateFieldsToBase() {
}



// Response to 'disassemble' request.
type DisassembleResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The list of disassembled instructions.
		Instructions []DisassembledInstruction `json:"instructions"`

	} `json:"body,omitempty"`

	// POSSIBLE VALUES: `disassemble`
	Command string `json:"command,omitempty"`

} // struct DisassembleResponse

func (me *DisassembleResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The 'terminate' request is sent from the client to the debug adapter in order to give the debuggee a chance for terminating itself.
type TerminateRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `terminate`
	Command string `json:"command"`

	Arguments TerminateArguments `json:"arguments,omitempty"`

} // struct TerminateRequest

func (me *TerminateRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The request starts the debuggee to run again for one step.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
type StepOutRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command"`

	Arguments StepOutArguments `json:"arguments"`

} // struct StepOutRequest

func (me *StepOutRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// A Thread
type Thread struct {

	// Unique identifier for the thread.
	Id int `json:"id"`

	// A name of the thread.
	Name string `json:"name"`

} // struct Thread

func (me *Thread) propagateFieldsToBase() {
}



// The event indicates that some source has been added, changed, or removed from the set of all loaded sources.
type LoadedSourceEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `loadedSource`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `new`, `changed`, `removed`
		Reason string `json:"reason"`

		// The new, changed, or removed source.
		Source Source `json:"source"`

	} `json:"body"`

} // struct LoadedSourceEvent

func (me *LoadedSourceEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Response to 'stepInTargets' request.
type StepInTargetsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The possible stepIn targets of the specified source location.
		Targets []StepInTarget `json:"targets"`

	} `json:"body"`

	// POSSIBLE VALUES: `stepInTargets`
	Command string `json:"command,omitempty"`

} // struct StepInTargetsResponse

func (me *StepInTargetsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// On error (whenever 'success' is false), the body can provide more details.
type ErrorResponse struct {
	// Response for a request.
	Response

	Body struct {

		// An optional, structured error message.
		Error Message `json:"error,omitempty"`

	} `json:"body"`

} // struct ErrorResponse

func (me *ErrorResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'pause' request.
type PauseArguments struct {

	// Pause execution for this thread.
	ThreadId int `json:"threadId"`

} // struct PauseArguments

func (me *PauseArguments) propagateFieldsToBase() {
}



// Arguments for 'launch' request. Additional attributes are implementation specific.
type LaunchRequestArguments struct {

	// If noDebug is true the launch request should launch the program without enabling debugging.
	NoDebug bool `json:"noDebug,omitempty"`

	// Optional data from the previous, restarted session.
	// The data is sent as the 'restart' attribute of the 'terminated' event.
	// The client should leave the data intact.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	Restart__ interface{} `json:"__restart,omitempty"`

} // struct LaunchRequestArguments

func (me *LaunchRequestArguments) propagateFieldsToBase() {
}



// Response to 'attach' request. This is just an acknowledgement, so no body field is required.
type AttachResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `attach`
	Command string `json:"command,omitempty"`

} // struct AttachResponse

func (me *AttachResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// This request retrieves the possible goto targets for the specified source location.
// These targets can be used in the 'goto' request.
// The GotoTargets request may only be called if the 'supportsGotoTargetsRequest' capability exists and is true.
type GotoTargetsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `gotoTargets`
	Command string `json:"command"`

	Arguments GotoTargetsArguments `json:"arguments"`

} // struct GotoTargetsRequest

func (me *GotoTargetsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The event indicates that the debugger has begun debugging a new process. Either one that it has launched, or one that it has attached to.
type ProcessEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `process`
	Event_ string `json:"event"`

	Body struct {

		// The size of a pointer or address for this process, in bits. This value may be used by clients when formatting addresses for display.
		PointerSize int `json:"pointerSize,omitempty"`

		// The logical name of the process. This is usually the full path to process's executable file. Example: /home/example/myproj/program.js.
		Name string `json:"name"`

		// The system process id of the debugged process. This property will be missing for non-system processes.
		SystemProcessId int `json:"systemProcessId,omitempty"`

		// If true, the process is running on the same computer as the debug adapter.
		IsLocalProcess bool `json:"isLocalProcess,omitempty"`

		// Describes how the debug engine started debugging this process.
		// 
		// POSSIBLE VALUES: `launch`, `attach`, `attachForSuspendedLaunch`
		StartMethod string `json:"startMethod,omitempty"`

	} `json:"body"`

} // struct ProcessEvent

func (me *ProcessEvent) propagateFieldsToBase() {
	me.Event.Body = me.Body
	me.Event.Event = me.Event_
	me.Event.propagateFieldsToBase()
}



// Arguments for 'restart' request.
type RestartArguments map[string]interface{}



// A StepInTarget can be used in the 'stepIn' request and determines into which single target the stepIn request should step.
type StepInTarget struct {

	// The name of the stepIn target (shown in the UI).
	Label string `json:"label"`

	// Unique identifier for a stepIn target.
	Id int `json:"id"`

} // struct StepInTarget

func (me *StepInTarget) propagateFieldsToBase() {
}



// Arguments for 'stepInTargets' request.
type StepInTargetsArguments struct {

	// The stack frame for which to retrieve the possible stepIn targets.
	FrameId int `json:"frameId"`

} // struct StepInTargetsArguments

func (me *StepInTargetsArguments) propagateFieldsToBase() {
}



// Replaces all existing data breakpoints with new data breakpoints.
// To clear all data breakpoints, specify an empty array.
// When a data breakpoint is hit, a 'stopped' event (with reason 'data breakpoint') is generated.
type SetDataBreakpointsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `setDataBreakpoints`
	Command string `json:"command"`

	Arguments SetDataBreakpointsArguments `json:"arguments"`

} // struct SetDataBreakpointsRequest

func (me *SetDataBreakpointsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'setVariable' request.
type SetVariableArguments struct {

	// The reference of the variable container.
	VariablesReference int `json:"variablesReference"`

	// The name of the variable in the container.
	Name string `json:"name"`

	// The value of the variable.
	Value string `json:"value"`

	// Specifies details on how to format the response value.
	Format ValueFormat `json:"format,omitempty"`

} // struct SetVariableArguments

func (me *SetVariableArguments) propagateFieldsToBase() {
}



// Response to 'modules' request.
type ModulesResponse struct {
	// Response for a request.
	Response

	Body struct {

		// All modules or range of modules.
		Modules []Module `json:"modules"`

		// The total number of modules available.
		TotalModules int `json:"totalModules,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `modules`
	Command string `json:"command,omitempty"`

} // struct ModulesResponse

func (me *ModulesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'stepIn' request. This is just an acknowledgement, so no body field is required.
type StepInResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command,omitempty"`

} // struct StepInResponse

func (me *StepInResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// An ExceptionBreakpointsFilter is shown in the UI as an option for configuring how exceptions are dealt with.
type ExceptionBreakpointsFilter struct {

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter"`

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label"`

	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`

} // struct ExceptionBreakpointsFilter

func (me *ExceptionBreakpointsFilter) propagateFieldsToBase() {
}



// Represents a single disassembled instruction.
type DisassembledInstruction struct {

	// The address of the instruction. Treated as a hex value if prefixed with '0x', or as a decimal value otherwise.
	Address string `json:"address"`

	// Text representing the instruction and its operands, in an implementation-defined format.
	Instruction string `json:"instruction"`

	// Source location that corresponds to this instruction, if any. Should always be set (if available) on the first instruction returned, but can be omitted afterwards if this instruction maps to the same source file as the previous instruction.
	Location Source `json:"location,omitempty"`

	// The column within the line that corresponds to this instruction, if any.
	Column int `json:"column,omitempty"`

	// The end line of the range that corresponds to this instruction, if any.
	EndLine int `json:"endLine,omitempty"`

	// The end column of the range that corresponds to this instruction, if any.
	EndColumn int `json:"endColumn,omitempty"`

	// Optional raw bytes representing the instruction and its operands, in an implementation-defined format.
	InstructionBytes string `json:"instructionBytes,omitempty"`

	// Name of the symbol that correponds with the location of this instruction, if any.
	Symbol string `json:"symbol,omitempty"`

	// The line within the source location that corresponds to this instruction, if any.
	Line int `json:"line,omitempty"`

} // struct DisassembledInstruction

func (me *DisassembledInstruction) propagateFieldsToBase() {
}



// Response to 'restartFrame' request. This is just an acknowledgement, so no body field is required.
type RestartFrameResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `restartFrame`
	Command string `json:"command,omitempty"`

} // struct RestartFrameResponse

func (me *RestartFrameResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'configurationDone' request.
type ConfigurationDoneArguments map[string]interface{}



// A Scope is a named container for variables. Optionally a scope can map to a source or a range within a source.
type Scope struct {

	// The number of named variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	NamedVariables int `json:"namedVariables,omitempty"`

	// The number of indexed variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	IndexedVariables int `json:"indexedVariables,omitempty"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

	// An optional hint for how to present this scope in the UI. If this attribute is missing, the scope is shown with a generic UI.
	// 
	// POSSIBLE VALUES: `arguments`, `locals`, `registers`
	PresentationHint string `json:"presentationHint,omitempty"`

	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive"`

	// Optional source for this scope.
	Source Source `json:"source,omitempty"`

	// Optional start line of the range covered by this scope.
	Line int `json:"line,omitempty"`

	// Optional start column of the range covered by this scope.
	Column int `json:"column,omitempty"`

	// Optional end line of the range covered by this scope.
	EndLine int `json:"endLine,omitempty"`

	// Optional end column of the range covered by this scope.
	EndColumn int `json:"endColumn,omitempty"`

	// Name of the scope such as 'Arguments', 'Locals', or 'Registers'. This string is shown in the UI as is and can be translated.
	Name string `json:"name"`

} // struct Scope

func (me *Scope) propagateFieldsToBase() {
}



// Response to 'exceptionInfo' request.
type ExceptionInfoResponse struct {
	// Response for a request.
	Response

	Body struct {

		// ID of the exception that was thrown.
		ExceptionId string `json:"exceptionId"`

		// Descriptive text for the exception provided by the debug adapter.
		Description string `json:"description,omitempty"`

		// Mode that caused the exception notification to be raised.
		BreakMode ExceptionBreakMode `json:"breakMode"`

		// Detailed information about the exception.
		Details ExceptionDetails `json:"details,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `exceptionInfo`
	Command string `json:"command,omitempty"`

} // struct ExceptionInfoResponse

func (me *ExceptionInfoResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'runInTerminal' request.
type RunInTerminalResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The process ID of the terminal shell.
		ShellProcessId int64 `json:"shellProcessId,omitempty"`

		// The process ID.
		ProcessId int64 `json:"processId,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `runInTerminal`
	Command string `json:"command,omitempty"`

} // struct RunInTerminalResponse

func (me *RunInTerminalResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// The ModulesViewDescriptor is the container for all declarative configuration options of a ModuleView.
// For now it only specifies the columns to be shown in the modules view.
type ModulesViewDescriptor struct {

	Columns []ColumnDescriptor `json:"columns"`

} // struct ModulesViewDescriptor

func (me *ModulesViewDescriptor) propagateFieldsToBase() {
}



// The event indicates that debugging of the debuggee has terminated. This does **not** mean that the debuggee itself has exited.
type TerminatedEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `terminated`
	Event_ string `json:"event"`

	Body struct {

		// A debug adapter may set 'restart' to true (or to an arbitrary object) to request that the front end restarts the session.
		// The value is not interpreted by the client and passed unmodified as an attribute '__restart' to the 'launch' and 'attach' requests.
		// 
		// POSSIBLE TYPES:
		// - `[]interface{}` (for JSON `array`s)
		// - `bool` (for JSON `boolean`s)
		// - `int` (for JSON `integer`s)
		// - `interface{/*nil*/}` (for JSON `null`s)
		// - `int64` (for JSON `number`s)
		// - `map[string]interface{}` (for JSON `object`s)
		// - `string` (for JSON `string`s)
		Restart interface{} `json:"restart,omitempty"`

	} `json:"body,omitempty"`

} // struct TerminatedEvent

func (me *TerminatedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// The checksum of an item calculated by the specified algorithm.
type Checksum struct {

	// The algorithm used to calculate this checksum.
	Algorithm ChecksumAlgorithm `json:"algorithm"`

	// Value of the checksum.
	Checksum string `json:"checksum"`

} // struct Checksum

func (me *Checksum) propagateFieldsToBase() {
}



// Response to 'variables' request.
type VariablesResponse struct {
	// Response for a request.
	Response

	Body struct {

		// All (or a range) of variables for the given variable reference.
		Variables []Variable `json:"variables"`

	} `json:"body"`

	// POSSIBLE VALUES: `variables`
	Command string `json:"command,omitempty"`

} // struct VariablesResponse

func (me *VariablesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'source' request.
type SourceResponse struct {
	// Response for a request.
	Response

	Body struct {

		// Content of the source reference.
		Content string `json:"content"`

		// Optional content type (mime type) of the source.
		MimeType string `json:"mimeType,omitempty"`

	} `json:"body"`

	// POSSIBLE VALUES: `source`
	Command string `json:"command,omitempty"`

} // struct SourceResponse

func (me *SourceResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The 'disconnect' request is sent from the client to the debug adapter in order to stop debugging. It asks the debug adapter to disconnect from the debuggee and to terminate the debug adapter. If the debuggee has been started with the 'launch' request, the 'disconnect' request terminates the debuggee. If the 'attach' request was used to connect to the debuggee, 'disconnect' does not terminate the debuggee. This behavior can be controlled with the 'terminateDebuggee' argument (if supported by the debug adapter).
type DisconnectRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `disconnect`
	Command string `json:"command"`

	Arguments DisconnectArguments `json:"arguments,omitempty"`

} // struct DisconnectRequest

func (me *DisconnectRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'completions' request.
type CompletionsArguments struct {

	// The character position for which to determine the completion proposals.
	Column int `json:"column"`

	// An optional line for which to determine the completion proposals. If missing the first line of the text is assumed.
	Line int `json:"line,omitempty"`

	// Returns completions in the scope of this stack frame. If not specified, the completions are returned for the global scope.
	FrameId int `json:"frameId,omitempty"`

	// One or more source lines. Typically this is the text a user has typed into the debug console before he asked for completion.
	Text string `json:"text"`

} // struct CompletionsArguments

func (me *CompletionsArguments) propagateFieldsToBase() {
}



// Arguments for 'continue' request.
type ContinueArguments struct {

	// Continue execution for the specified thread (if possible). If the backend cannot continue on a single thread but will continue on all threads, it should set the 'allThreadsContinued' attribute in the response to true.
	ThreadId int `json:"threadId"`

} // struct ContinueArguments

func (me *ContinueArguments) propagateFieldsToBase() {
}



// Arguments for 'evaluate' request.
type EvaluateArguments struct {

	// The context in which the evaluate request is run.
	// 
	// POSSIBLE VALUES: `watch`, `repl`, `hover`
	Context string `json:"context,omitempty"`

	// Specifies details on how to format the Evaluate result.
	Format ValueFormat `json:"format,omitempty"`

	// The expression to evaluate.
	Expression string `json:"expression"`

	// Evaluate the expression in the scope of this stack frame. If not specified, the expression is evaluated in the global scope.
	FrameId int `json:"frameId,omitempty"`

} // struct EvaluateArguments

func (me *EvaluateArguments) propagateFieldsToBase() {
}



// Arguments for 'setBreakpoints' request.
type SetBreakpointsArguments struct {

	// Deprecated: The code locations of the breakpoints.
	Lines []int `json:"lines,omitempty"`

	// A value of true indicates that the underlying source has been modified which results in new breakpoint locations.
	SourceModified bool `json:"sourceModified,omitempty"`

	// The source location of the breakpoints; either 'source.path' or 'source.reference' must be specified.
	Source Source `json:"source"`

	// The code locations of the breakpoints.
	Breakpoints []SourceBreakpoint `json:"breakpoints,omitempty"`

} // struct SetBreakpointsArguments

func (me *SetBreakpointsArguments) propagateFieldsToBase() {
}



// The 'initialize' request is sent as the first request from the client to the debug adapter in order to configure it with client capabilities and to retrieve capabilities from the debug adapter.
// Until the debug adapter has responded to with an 'initialize' response, the client must not send any additional requests or events to the debug adapter. In addition the debug adapter is not allowed to send any requests or events to the client until it has responded with an 'initialize' response.
// The 'initialize' request may only be sent once.
type InitializeRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `initialize`
	Command string `json:"command"`

	Arguments InitializeRequestArguments `json:"arguments"`

} // struct InitializeRequest

func (me *InitializeRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'setExpression' request.
type SetExpressionArguments struct {

	// Specifies how the resulting value should be formatted.
	Format ValueFormat `json:"format,omitempty"`

	// The l-value expression to assign to.
	Expression string `json:"expression"`

	// The value expression to assign to the l-value expression.
	Value string `json:"value"`

	// Evaluate the expressions in the scope of this stack frame. If not specified, the expressions are evaluated in the global scope.
	FrameId int `json:"frameId,omitempty"`

} // struct SetExpressionArguments

func (me *SetExpressionArguments) propagateFieldsToBase() {
}



// Arguments for 'terminateThreads' request.
type TerminateThreadsArguments struct {

	// Ids of threads to be terminated.
	ThreadIds []int `json:"threadIds,omitempty"`

} // struct TerminateThreadsArguments

func (me *TerminateThreadsArguments) propagateFieldsToBase() {
}



// Properties of a breakpoint passed to the setFunctionBreakpoints request.
type FunctionBreakpoint struct {

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

	// The name of the function.
	Name string `json:"name"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

} // struct FunctionBreakpoint

func (me *FunctionBreakpoint) propagateFieldsToBase() {
}



// Base class of requests, responses, and events.
type ProtocolMessage struct {

	// Sequence number.
	Seq int `json:"seq"`

	// Message type.
	// 
	// POSSIBLE VALUES: `request`, `response`, `event`
	Type string `json:"type"`

} // struct ProtocolMessage

func (me *ProtocolMessage) propagateFieldsToBase() {
}



// Response to 'configurationDone' request. This is just an acknowledgement, so no body field is required.
type ConfigurationDoneResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `configurationDone`
	Command string `json:"command,omitempty"`

} // struct ConfigurationDoneResponse

func (me *ConfigurationDoneResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Provides formatting information for a value.
type ValueFormat struct {

	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`

} // struct ValueFormat

func (me *ValueFormat) propagateFieldsToBase() {
}



// Disassembles code stored at the provided location.
type DisassembleRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `disassemble`
	Command string `json:"command"`

	Arguments DisassembleArguments `json:"arguments"`

} // struct DisassembleRequest

func (me *DisassembleRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The request sets the location where the debuggee will continue to run.
// This makes it possible to skip the execution of code or to executed code again.
// The code between the current location and the goto target is not executed but skipped.
// The debug adapter first sends the response and then a 'stopped' event with reason 'goto'.
type GotoRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `goto`
	Command string `json:"command"`

	Arguments GotoArguments `json:"arguments"`

} // struct GotoRequest

func (me *GotoRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'stepOut' request. This is just an acknowledgement, so no body field is required.
type StepOutResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command,omitempty"`

} // struct StepOutResponse

func (me *StepOutResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Properties of a breakpoint or logpoint passed to the setBreakpoints request.
type SourceBreakpoint struct {

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

	// If this attribute exists and is non-empty, the backend must not 'break' (stop) but log the message instead. Expressions within {} are interpolated.
	LogMessage string `json:"logMessage,omitempty"`

	// The source line of the breakpoint or logpoint.
	Line int `json:"line"`

	// An optional source column of the breakpoint.
	Column int `json:"column,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

} // struct SourceBreakpoint

func (me *SourceBreakpoint) propagateFieldsToBase() {
}



// The request starts the debuggee to step into a function/method if possible.
// If it cannot step into a target, 'stepIn' behaves like 'next'.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
// If there are multiple function/method calls (or other targets) on the source line,
// the optional argument 'targetId' can be used to control into which target the 'stepIn' should occur.
// The list of possible targets for a given source line can be retrieved via the 'stepInTargets' request.
type StepInRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command"`

	Arguments StepInArguments `json:"arguments"`

} // struct StepInRequest

func (me *StepInRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'terminate' request. This is just an acknowledgement, so no body field is required.
type TerminateResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `terminate`
	Command string `json:"command,omitempty"`

} // struct TerminateResponse

func (me *TerminateResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The event indicates that the execution of the debuggee has continued.
// Please note: a debug adapter is not expected to send this event in response to a request that implies that execution continues, e.g. 'launch' or 'continue'.
// It is only necessary to send a 'continued' event if there was no previous request that implied this.
type ContinuedEvent struct {
	// A debug adapter initiated event.
	Event

	// POSSIBLE VALUES: `continued`
	Event_ string `json:"event"`

	Body struct {

		// The thread which was continued.
		ThreadId int `json:"threadId"`

		// If 'allThreadsContinued' is true, a debug adapter can announce that all threads have continued.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

	} `json:"body"`

} // struct ContinuedEvent

func (me *ContinuedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Obtains information on a possible data breakpoint that could be set on an expression or variable.
type DataBreakpointInfoRequest struct {
	// A client or debug adapter initiated request.
	Request

	Arguments DataBreakpointInfoArguments `json:"arguments"`

	// POSSIBLE VALUES: `dataBreakpointInfo`
	Command string `json:"command"`

} // struct DataBreakpointInfoRequest

func (me *DataBreakpointInfoRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Response to 'completions' request.
type CompletionsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The possible completions for .
		Targets []CompletionItem `json:"targets"`

	} `json:"body"`

	// POSSIBLE VALUES: `completions`
	Command string `json:"command,omitempty"`

} // struct CompletionsResponse

func (me *CompletionsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'scopes' request.
type ScopesArguments struct {

	// Retrieve the scopes for this stackframe.
	FrameId int `json:"frameId"`

} // struct ScopesArguments

func (me *ScopesArguments) propagateFieldsToBase() {
}



// Reads bytes from memory at the provided location.
type ReadMemoryRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `readMemory`
	Command string `json:"command"`

	Arguments ReadMemoryArguments `json:"arguments"`

} // struct ReadMemoryRequest

func (me *ReadMemoryRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'setBreakpoints' request.
// Returned is information about each breakpoint created by this request.
// This includes the actual code location and whether the breakpoint could be verified.
// The breakpoints returned are in the same order as the elements of the 'breakpoints'
// (or the deprecated 'lines') array in the arguments.
type SetBreakpointsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements are in the same order as the elements of the 'breakpoints' (or the deprecated 'lines') array in the arguments.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`

	// POSSIBLE VALUES: `setBreakpoints`
	Command string `json:"command,omitempty"`

} // struct SetBreakpointsResponse

func (me *SetBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'reverseContinue' request. This is just an acknowledgement, so no body field is required.
type ReverseContinueResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `reverseContinue`
	Command string `json:"command,omitempty"`

} // struct ReverseContinueResponse

func (me *ReverseContinueResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'next' request. This is just an acknowledgement, so no body field is required.
type NextResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `next`
	Command string `json:"command,omitempty"`

} // struct NextResponse

func (me *NextResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Evaluates the given expression in the context of the top most stack frame.
// The expression has access to any variables and arguments that are in scope.
type EvaluateRequest struct {
	// A client or debug adapter initiated request.
	Request

	Arguments EvaluateArguments `json:"arguments"`

	// POSSIBLE VALUES: `evaluate`
	Command string `json:"command"`

} // struct EvaluateRequest

func (me *EvaluateRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// This enumeration defines all possible access types for data breakpoints.
// 
// POSSIBLE VALUES: `read`, `write`, `readWrite`
type DataBreakpointAccessType string



// Response to 'restart' request. This is just an acknowledgement, so no body field is required.
type RestartResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `restart`
	Command string `json:"command,omitempty"`

} // struct RestartResponse

func (me *RestartResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Response to 'evaluate' request.
type EvaluateResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `evaluate`
	Command string `json:"command,omitempty"`

	Body struct {

		// If variablesReference is > 0, the evaluate result is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference"`

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

		// Memory reference to a location appropriate for this result. For pointer type eval results, this is generally a reference to the memory address contained in the pointer.
		MemoryReference string `json:"memoryReference,omitempty"`

		// The result of the evaluate request.
		Result string `json:"result"`

		// The optional type of the evaluate result.
		Type string `json:"type,omitempty"`

		// Properties of a evaluate result that can be used to determine how to render the result in the UI.
		PresentationHint VariablePresentationHint `json:"presentationHint,omitempty"`

	} `json:"body"`

} // struct EvaluateResponse

func (me *EvaluateResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Retrieves all child variables for the given variable reference.
// An optional filter can be used to limit the fetched children to either named or indexed children.
type VariablesRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `variables`
	Command string `json:"command"`

	Arguments VariablesArguments `json:"arguments"`

} // struct VariablesRequest

func (me *VariablesRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// A GotoTarget describes a code location that can be used as a target in the 'goto' request.
// The possible goto targets can be determined via the 'gotoTargets' request.
type GotoTarget struct {

	// An optional end column of the range covered by the goto target.
	EndColumn int `json:"endColumn,omitempty"`

	// Optional memory reference for the instruction pointer value represented by this target.
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`

	// Unique identifier for a goto target. This is used in the goto request.
	Id int `json:"id"`

	// The name of the goto target (shown in the UI).
	Label string `json:"label"`

	// The line of the goto target.
	Line int `json:"line"`

	// An optional column of the goto target.
	Column int `json:"column,omitempty"`

	// An optional end line of the range covered by the goto target.
	EndLine int `json:"endLine,omitempty"`

} // struct GotoTarget

func (me *GotoTarget) propagateFieldsToBase() {
}



// Response for a request.
type Response struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

	// POSSIBLE VALUES: `response`
	Type string `json:"type"`

	// Sequence number of the corresponding request.
	Request_seq int `json:"request_seq"`

	// Outcome of the request.
	Success bool `json:"success"`

	// The command requested.
	Command string `json:"command"`

	// Contains error message if success == false.
	Message string `json:"message,omitempty"`

	// Contains request result if success is true and optional error details if success is false.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	Body interface{} `json:"body,omitempty"`

} // struct Response

func (me *Response) propagateFieldsToBase() {
	me.ProtocolMessage.Type = me.Type
	me.ProtocolMessage.propagateFieldsToBase()
}



// Replaces all existing function breakpoints with new function breakpoints.
// To clear all function breakpoints, specify an empty array.
// When a function breakpoint is hit, a 'stopped' event (with reason 'function breakpoint') is generated.
type SetFunctionBreakpointsRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `setFunctionBreakpoints`
	Command string `json:"command"`

	Arguments SetFunctionBreakpointsArguments `json:"arguments"`

} // struct SetFunctionBreakpointsRequest

func (me *SetFunctionBreakpointsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Evaluates the given 'value' expression and assigns it to the 'expression' which must be a modifiable l-value.
// The expressions have access to any variables and arguments that are in scope of the specified frame.
type SetExpressionRequest struct {
	// A client or debug adapter initiated request.
	Request

	Arguments SetExpressionArguments `json:"arguments"`

	// POSSIBLE VALUES: `setExpression`
	Command string `json:"command"`

} // struct SetExpressionRequest

func (me *SetExpressionRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// The request starts the debuggee to run again for one step.
// The debug adapter first sends the response and then a 'stopped' event (with reason 'step') after the step has completed.
type NextRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `next`
	Command string `json:"command"`

	Arguments NextArguments `json:"arguments"`

} // struct NextRequest

func (me *NextRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// A Stackframe contains the source location.
type StackFrame struct {

	// The name of the stack frame, typically a method name.
	Name string `json:"name"`

	// The optional source of the frame.
	Source Source `json:"source,omitempty"`

	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column int `json:"column"`

	// Optional memory reference for the current instruction pointer in this frame.
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`

	// An optional hint for how to present this frame in the UI. A value of 'label' can be used to indicate that the frame is an artificial frame that is used as a visual label or separator. A value of 'subtle' can be used to change the appearance of a frame in a 'subtle' way.
	// 
	// POSSIBLE VALUES: `normal`, `label`, `subtle`
	PresentationHint string `json:"presentationHint,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id int `json:"id"`

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line int `json:"line"`

	// An optional end line of the range covered by the stack frame.
	EndLine int `json:"endLine,omitempty"`

	// An optional end column of the range covered by the stack frame.
	EndColumn int `json:"endColumn,omitempty"`

	// The module associated with this frame, if any.
	// 
	// POSSIBLE TYPES:
	// - `int` (for JSON `integer`s)
	// - `string` (for JSON `string`s)
	ModuleId interface{} `json:"moduleId,omitempty"`

} // struct StackFrame

func (me *StackFrame) propagateFieldsToBase() {
}



// Response to 'disconnect' request. This is just an acknowledgement, so no body field is required.
type DisconnectResponse struct {
	// Response for a request.
	Response

	// POSSIBLE VALUES: `disconnect`
	Command string `json:"command,omitempty"`

} // struct DisconnectResponse

func (me *DisconnectResponse) propagateFieldsToBase() {
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The launch request is sent from the client to the debug adapter to start the debuggee with or without debugging (if 'noDebug' is true). Since launching is debugger/runtime specific, the arguments for this request are not part of this specification.
type LaunchRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `launch`
	Command string `json:"command"`

	Arguments LaunchRequestArguments `json:"arguments"`

} // struct LaunchRequest

func (me *LaunchRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// An ExceptionPathSegment represents a segment in a path that is used to match leafs or nodes in a tree of exceptions. If a segment consists of more than one name, it matches the names provided if 'negate' is false or missing or it matches anything except the names provided if 'negate' is true.
type ExceptionPathSegment struct {

	// If false or missing this segment matches the names provided, otherwise it matches anything except the names provided.
	Negate bool `json:"negate,omitempty"`

	// Depending on the value of 'negate' the names that should match or not match.
	Names []string `json:"names"`

} // struct ExceptionPathSegment

func (me *ExceptionPathSegment) propagateFieldsToBase() {
}



// Arguments for 'stepBack' request.
type StepBackArguments struct {

	// Execute 'stepBack' for this thread.
	ThreadId int `json:"threadId"`

} // struct StepBackArguments

func (me *StepBackArguments) propagateFieldsToBase() {
}



// Set the variable with the given name in the variable container to a new value.
type SetVariableRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `setVariable`
	Command string `json:"command"`

	Arguments SetVariableArguments `json:"arguments"`

} // struct SetVariableRequest

func (me *SetVariableRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Retrieves the set of all sources currently loaded by the debugged process.
type LoadedSourcesRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `loadedSources`
	Command string `json:"command"`

	Arguments LoadedSourcesArguments `json:"arguments,omitempty"`

} // struct LoadedSourcesRequest

func (me *LoadedSourcesRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Restarts a debug session. If the capability 'supportsRestartRequest' is missing or has the value false,
// the client will implement 'restart' by terminating the debug adapter first and then launching it anew.
// A debug adapter can override this default behaviour by implementing a restart request
// and setting the capability 'supportsRestartRequest' to true.
type RestartRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `restart`
	Command string `json:"command"`

	Arguments RestartArguments `json:"arguments,omitempty"`

} // struct RestartRequest

func (me *RestartRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// A Source is a descriptor for source code. It is returned from the debug adapter as part of a StackFrame and it is used by clients when specifying breakpoints.
type Source struct {

	// The (optional) origin of this source: possible values 'internal module', 'inlined content from source map', etc.
	Origin string `json:"origin,omitempty"`

	// An optional list of sources that are related to this source. These may be the source that generated this source.
	Sources []Source `json:"sources,omitempty"`

	// Optional data that a debug adapter might want to loop through the client. The client should leave the data intact and persist it across sessions. The client should not interpret the data.
	// 
	// POSSIBLE TYPES:
	// - `[]interface{}` (for JSON `array`s)
	// - `bool` (for JSON `boolean`s)
	// - `int` (for JSON `integer`s)
	// - `interface{/*nil*/}` (for JSON `null`s)
	// - `int64` (for JSON `number`s)
	// - `map[string]interface{}` (for JSON `object`s)
	// - `string` (for JSON `string`s)
	AdapterData interface{} `json:"adapterData,omitempty"`

	// The checksums associated with this file.
	Checksums []Checksum `json:"checksums,omitempty"`

	// The short name of the source. Every source returned from the debug adapter has a name. When sending a source to the debug adapter this name is optional.
	Name string `json:"name,omitempty"`

	// The path of the source to be shown in the UI. It is only used to locate and load the content of the source if no sourceReference is specified (or its value is 0).
	Path string `json:"path,omitempty"`

	// If sourceReference > 0 the contents of the source must be retrieved through the SourceRequest (even if a path is specified). A sourceReference is only valid for a session, so it must not be used to persist a source.
	SourceReference int64 `json:"sourceReference,omitempty"`

	// An optional hint for how to present the source in the UI. A value of 'deemphasize' can be used to indicate that the source is not available or that it is skipped on stepping.
	// 
	// POSSIBLE VALUES: `normal`, `emphasize`, `deemphasize`
	PresentationHint string `json:"presentationHint,omitempty"`

} // struct Source

func (me *Source) propagateFieldsToBase() {
}



// Response to 'setVariable' request.
type SetVariableResponse struct {
	// Response for a request.
	Response

	Body struct {

		// The type of the new value. Typically shown in the UI when hovering over the value.
		Type string `json:"type,omitempty"`

		// If variablesReference is > 0, the new value is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference,omitempty"`

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

		// The new value of the variable.
		Value string `json:"value"`

	} `json:"body"`

	// POSSIBLE VALUES: `setVariable`
	Command string `json:"command,omitempty"`

} // struct SetVariableResponse

func (me *SetVariableResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// Arguments for 'next' request.
type NextArguments struct {

	// Execute 'next' for this thread.
	ThreadId int `json:"threadId"`

} // struct NextArguments

func (me *NextArguments) propagateFieldsToBase() {
}



// Response to 'threads' request.
type ThreadsResponse struct {
	// Response for a request.
	Response

	Body struct {

		// All threads.
		Threads []Thread `json:"threads"`

	} `json:"body"`

	// POSSIBLE VALUES: `threads`
	Command string `json:"command,omitempty"`

} // struct ThreadsResponse

func (me *ThreadsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.Command = me.Command
	me.Response.propagateFieldsToBase()
}



// The request returns the variable scopes for a given stackframe ID.
type ScopesRequest struct {
	// A client or debug adapter initiated request.
	Request

	// POSSIBLE VALUES: `scopes`
	Command string `json:"command"`

	Arguments ScopesArguments `json:"arguments"`

} // struct ScopesRequest

func (me *ScopesRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Arguments for 'terminate' request.
type TerminateArguments struct {

	// A value of true indicates that this 'terminate' request is part of a restart sequence.
	Restart bool `json:"restart,omitempty"`

} // struct TerminateArguments

func (me *TerminateArguments) propagateFieldsToBase() {
}



// Arguments for 'reverseContinue' request.
type ReverseContinueArguments struct {

	// Execute 'reverseContinue' for this thread.
	ThreadId int `json:"threadId"`

} // struct ReverseContinueArguments

func (me *ReverseContinueArguments) propagateFieldsToBase() {
}



// Arguments for 'readMemory' request.
type ReadMemoryArguments struct {

	// Memory reference to the base location from which data should be read.
	MemoryReference string `json:"memoryReference"`

	// Optional offset (in bytes) to be applied to the reference location before reading data. Can be negative.
	Offset int `json:"offset,omitempty"`

	// Number of bytes to read at the specified location and offset.
	Count int `json:"count"`

} // struct ReadMemoryArguments

func (me *ReadMemoryArguments) propagateFieldsToBase() {
}



// Optional properties of a variable that can be used to determine how to render the variable in the UI.
type VariablePresentationHint struct {

	// Visibility of variable. Before introducing additional values, try to use the listed values.
	// 
	// POSSIBLE VALUES: `public`, `private`, `protected`, `internal`, `final`
	Visibility string `json:"visibility,omitempty"`

	// The kind of variable. Before introducing additional values, try to use the listed values.
	// 
	// POSSIBLE VALUES: `property`, `method`, `class`, `data`, `event`, `baseClass`, `innerClass`, `interface`, `mostDerivedClass`, `virtual`, `dataBreakpoint`
	Kind string `json:"kind,omitempty"`

	// Set of attributes represented as an array of strings. Before introducing additional values, try to use the listed values.
	Attributes []string `json:"attributes,omitempty"`

} // struct VariablePresentationHint

func (me *VariablePresentationHint) propagateFieldsToBase() {
}

func BaseEvent (someEvent interface{}) (baseEvent *Event) {
	switch me := someEvent.(type) {
	case *ExitedEvent: baseEvent = &me.Event
	case *InitializedEvent: baseEvent = &me.Event
	case *CapabilitiesEvent: baseEvent = &me.Event
	case *StoppedEvent: baseEvent = &me.Event
	case *OutputEvent: baseEvent = &me.Event
	case *ModuleEvent: baseEvent = &me.Event
	case *LoadedSourceEvent: baseEvent = &me.Event
	case *ProcessEvent: baseEvent = &me.Event
	case *TerminatedEvent: baseEvent = &me.Event
	case *ContinuedEvent: baseEvent = &me.Event
	case *ThreadEvent: baseEvent = &me.Event
	case *BreakpointEvent: baseEvent = &me.Event
	}
	return
}

func BaseResponse (someResponse interface{}) (baseResponse *Response) {
	switch me := someResponse.(type) {
	case *RestartResponse: baseResponse = &me.Response
	case *EvaluateResponse: baseResponse = &me.Response
	case *DisconnectResponse: baseResponse = &me.Response
	case *SetVariableResponse: baseResponse = &me.Response
	case *ThreadsResponse: baseResponse = &me.Response
	case *LaunchResponse: baseResponse = &me.Response
	case *ScopesResponse: baseResponse = &me.Response
	case *SetExpressionResponse: baseResponse = &me.Response
	case *DataBreakpointInfoResponse: baseResponse = &me.Response
	case *StepBackResponse: baseResponse = &me.Response
	case *InitializeResponse: baseResponse = &me.Response
	case *SetDataBreakpointsResponse: baseResponse = &me.Response
	case *StackTraceResponse: baseResponse = &me.Response
	case *SetFunctionBreakpointsResponse: baseResponse = &me.Response
	case *GotoResponse: baseResponse = &me.Response
	case *SetExceptionBreakpointsResponse: baseResponse = &me.Response
	case *TerminateThreadsResponse: baseResponse = &me.Response
	case *GotoTargetsResponse: baseResponse = &me.Response
	case *PauseResponse: baseResponse = &me.Response
	case *ReadMemoryResponse: baseResponse = &me.Response
	case *LoadedSourcesResponse: baseResponse = &me.Response
	case *ContinueResponse: baseResponse = &me.Response
	case *StepInTargetsResponse: baseResponse = &me.Response
	case *DisassembleResponse: baseResponse = &me.Response
	case *ErrorResponse: baseResponse = &me.Response
	case *AttachResponse: baseResponse = &me.Response
	case *ModulesResponse: baseResponse = &me.Response
	case *StepInResponse: baseResponse = &me.Response
	case *RestartFrameResponse: baseResponse = &me.Response
	case *VariablesResponse: baseResponse = &me.Response
	case *SourceResponse: baseResponse = &me.Response
	case *ExceptionInfoResponse: baseResponse = &me.Response
	case *RunInTerminalResponse: baseResponse = &me.Response
	case *ConfigurationDoneResponse: baseResponse = &me.Response
	case *StepOutResponse: baseResponse = &me.Response
	case *TerminateResponse: baseResponse = &me.Response
	case *ReverseContinueResponse: baseResponse = &me.Response
	case *NextResponse: baseResponse = &me.Response
	case *CompletionsResponse: baseResponse = &me.Response
	case *SetBreakpointsResponse: baseResponse = &me.Response
	}
	return
}

func BaseRequest (someRequest interface{}) (baseRequest *Request) {
	switch me := someRequest.(type) {
	case *StepBackRequest: baseRequest = &me.Request
	case *ExceptionInfoRequest: baseRequest = &me.Request
	case *StepInTargetsRequest: baseRequest = &me.Request
	case *RestartFrameRequest: baseRequest = &me.Request
	case *RunInTerminalRequest: baseRequest = &me.Request
	case *CompletionsRequest: baseRequest = &me.Request
	case *ModulesRequest: baseRequest = &me.Request
	case *SourceRequest: baseRequest = &me.Request
	case *PauseRequest: baseRequest = &me.Request
	case *ConfigurationDoneRequest: baseRequest = &me.Request
	case *TerminateRequest: baseRequest = &me.Request
	case *StepOutRequest: baseRequest = &me.Request
	case *GotoTargetsRequest: baseRequest = &me.Request
	case *SetDataBreakpointsRequest: baseRequest = &me.Request
	case *DisconnectRequest: baseRequest = &me.Request
	case *InitializeRequest: baseRequest = &me.Request
	case *DisassembleRequest: baseRequest = &me.Request
	case *GotoRequest: baseRequest = &me.Request
	case *StepInRequest: baseRequest = &me.Request
	case *DataBreakpointInfoRequest: baseRequest = &me.Request
	case *ReadMemoryRequest: baseRequest = &me.Request
	case *EvaluateRequest: baseRequest = &me.Request
	case *VariablesRequest: baseRequest = &me.Request
	case *SetFunctionBreakpointsRequest: baseRequest = &me.Request
	case *SetExpressionRequest: baseRequest = &me.Request
	case *NextRequest: baseRequest = &me.Request
	case *LaunchRequest: baseRequest = &me.Request
	case *SetVariableRequest: baseRequest = &me.Request
	case *LoadedSourcesRequest: baseRequest = &me.Request
	case *RestartRequest: baseRequest = &me.Request
	case *ScopesRequest: baseRequest = &me.Request
	case *StackTraceRequest: baseRequest = &me.Request
	case *ThreadsRequest: baseRequest = &me.Request
	case *TerminateThreadsRequest: baseRequest = &me.Request
	case *AttachRequest: baseRequest = &me.Request
	case *ReverseContinueRequest: baseRequest = &me.Request
	case *SetBreakpointsRequest: baseRequest = &me.Request
	case *SetExceptionBreakpointsRequest: baseRequest = &me.Request
	case *ContinueRequest: baseRequest = &me.Request
	}
	return
}

// Returns a new `StepInTargetsResponse` with the following fields set: `Command`, `Type`
func NewStepInTargetsResponse() *StepInTargetsResponse {
	newStepInTargetsResponse := StepInTargetsResponse{}
	newStepInTargetsResponse.Command = "stepInTargets"
	newStepInTargetsResponse.Type = "response"
	newStepInTargetsResponse.propagateFieldsToBase()
	return &newStepInTargetsResponse
}

// Returns a new `SetBreakpointsResponse` with the following fields set: `Command`, `Type`
func NewSetBreakpointsResponse() *SetBreakpointsResponse {
	newSetBreakpointsResponse := SetBreakpointsResponse{}
	newSetBreakpointsResponse.Command = "setBreakpoints"
	newSetBreakpointsResponse.Type = "response"
	newSetBreakpointsResponse.propagateFieldsToBase()
	return &newSetBreakpointsResponse
}

// Returns a new `TerminateResponse` with the following fields set: `Command`, `Type`
func NewTerminateResponse() *TerminateResponse {
	newTerminateResponse := TerminateResponse{}
	newTerminateResponse.Command = "terminate"
	newTerminateResponse.Type = "response"
	newTerminateResponse.propagateFieldsToBase()
	return &newTerminateResponse
}

// Returns a new `ContinuedEvent` with the following fields set: `Event_`, `Type`
func NewContinuedEvent() *ContinuedEvent {
	newContinuedEvent := ContinuedEvent{}
	newContinuedEvent.Event_ = "continued"
	newContinuedEvent.Type = "event"
	newContinuedEvent.propagateFieldsToBase()
	return &newContinuedEvent
}

// Returns a new `BreakpointEvent` with the following fields set: `Event_`, `Type`
func NewBreakpointEvent() *BreakpointEvent {
	newBreakpointEvent := BreakpointEvent{}
	newBreakpointEvent.Event_ = "breakpoint"
	newBreakpointEvent.Type = "event"
	newBreakpointEvent.propagateFieldsToBase()
	return &newBreakpointEvent
}

// Returns a new `SetFunctionBreakpointsResponse` with the following fields set: `Command`, `Type`
func NewSetFunctionBreakpointsResponse() *SetFunctionBreakpointsResponse {
	newSetFunctionBreakpointsResponse := SetFunctionBreakpointsResponse{}
	newSetFunctionBreakpointsResponse.Command = "setFunctionBreakpoints"
	newSetFunctionBreakpointsResponse.Type = "response"
	newSetFunctionBreakpointsResponse.propagateFieldsToBase()
	return &newSetFunctionBreakpointsResponse
}

// Returns a new `CapabilitiesEvent` with the following fields set: `Event_`, `Type`
func NewCapabilitiesEvent() *CapabilitiesEvent {
	newCapabilitiesEvent := CapabilitiesEvent{}
	newCapabilitiesEvent.Event_ = "capabilities"
	newCapabilitiesEvent.Type = "event"
	newCapabilitiesEvent.propagateFieldsToBase()
	return &newCapabilitiesEvent
}

// Returns a new `ConfigurationDoneRequest` with the following fields set: `Command`, `Type`
func NewConfigurationDoneRequest() *ConfigurationDoneRequest {
	newConfigurationDoneRequest := ConfigurationDoneRequest{}
	newConfigurationDoneRequest.Command = "configurationDone"
	newConfigurationDoneRequest.Type = "request"
	newConfigurationDoneRequest.propagateFieldsToBase()
	return &newConfigurationDoneRequest
}

// Returns a new `SetDataBreakpointsRequest` with the following fields set: `Command`, `Type`
func NewSetDataBreakpointsRequest() *SetDataBreakpointsRequest {
	newSetDataBreakpointsRequest := SetDataBreakpointsRequest{}
	newSetDataBreakpointsRequest.Command = "setDataBreakpoints"
	newSetDataBreakpointsRequest.Type = "request"
	newSetDataBreakpointsRequest.propagateFieldsToBase()
	return &newSetDataBreakpointsRequest
}

// Returns a new `ThreadsRequest` with the following fields set: `Command`, `Type`
func NewThreadsRequest() *ThreadsRequest {
	newThreadsRequest := ThreadsRequest{}
	newThreadsRequest.Command = "threads"
	newThreadsRequest.Type = "request"
	newThreadsRequest.propagateFieldsToBase()
	return &newThreadsRequest
}

// Returns a new `ExitedEvent` with the following fields set: `Event_`, `Type`
func NewExitedEvent() *ExitedEvent {
	newExitedEvent := ExitedEvent{}
	newExitedEvent.Event_ = "exited"
	newExitedEvent.Type = "event"
	newExitedEvent.propagateFieldsToBase()
	return &newExitedEvent
}

// Returns a new `VariablesResponse` with the following fields set: `Command`, `Type`
func NewVariablesResponse() *VariablesResponse {
	newVariablesResponse := VariablesResponse{}
	newVariablesResponse.Command = "variables"
	newVariablesResponse.Type = "response"
	newVariablesResponse.propagateFieldsToBase()
	return &newVariablesResponse
}

// Returns a new `DataBreakpointInfoRequest` with the following fields set: `Command`, `Type`
func NewDataBreakpointInfoRequest() *DataBreakpointInfoRequest {
	newDataBreakpointInfoRequest := DataBreakpointInfoRequest{}
	newDataBreakpointInfoRequest.Command = "dataBreakpointInfo"
	newDataBreakpointInfoRequest.Type = "request"
	newDataBreakpointInfoRequest.propagateFieldsToBase()
	return &newDataBreakpointInfoRequest
}

// Returns a new `SetExpressionRequest` with the following fields set: `Command`, `Type`
func NewSetExpressionRequest() *SetExpressionRequest {
	newSetExpressionRequest := SetExpressionRequest{}
	newSetExpressionRequest.Command = "setExpression"
	newSetExpressionRequest.Type = "request"
	newSetExpressionRequest.propagateFieldsToBase()
	return &newSetExpressionRequest
}

// Returns a new `ThreadEvent` with the following fields set: `Event_`, `Type`
func NewThreadEvent() *ThreadEvent {
	newThreadEvent := ThreadEvent{}
	newThreadEvent.Event_ = "thread"
	newThreadEvent.Type = "event"
	newThreadEvent.propagateFieldsToBase()
	return &newThreadEvent
}

// Returns a new `TerminatedEvent` with the following fields set: `Event_`, `Type`
func NewTerminatedEvent() *TerminatedEvent {
	newTerminatedEvent := TerminatedEvent{}
	newTerminatedEvent.Event_ = "terminated"
	newTerminatedEvent.Type = "event"
	newTerminatedEvent.propagateFieldsToBase()
	return &newTerminatedEvent
}

// Returns a new `SetVariableResponse` with the following fields set: `Command`, `Type`
func NewSetVariableResponse() *SetVariableResponse {
	newSetVariableResponse := SetVariableResponse{}
	newSetVariableResponse.Command = "setVariable"
	newSetVariableResponse.Type = "response"
	newSetVariableResponse.propagateFieldsToBase()
	return &newSetVariableResponse
}

// Returns a new `ThreadsResponse` with the following fields set: `Command`, `Type`
func NewThreadsResponse() *ThreadsResponse {
	newThreadsResponse := ThreadsResponse{}
	newThreadsResponse.Command = "threads"
	newThreadsResponse.Type = "response"
	newThreadsResponse.propagateFieldsToBase()
	return &newThreadsResponse
}

// Returns a new `SetExceptionBreakpointsRequest` with the following fields set: `Command`, `Type`
func NewSetExceptionBreakpointsRequest() *SetExceptionBreakpointsRequest {
	newSetExceptionBreakpointsRequest := SetExceptionBreakpointsRequest{}
	newSetExceptionBreakpointsRequest.Command = "setExceptionBreakpoints"
	newSetExceptionBreakpointsRequest.Type = "request"
	newSetExceptionBreakpointsRequest.propagateFieldsToBase()
	return &newSetExceptionBreakpointsRequest
}

// Returns a new `TerminateThreadsResponse` with the following fields set: `Command`, `Type`
func NewTerminateThreadsResponse() *TerminateThreadsResponse {
	newTerminateThreadsResponse := TerminateThreadsResponse{}
	newTerminateThreadsResponse.Command = "terminateThreads"
	newTerminateThreadsResponse.Type = "response"
	newTerminateThreadsResponse.propagateFieldsToBase()
	return &newTerminateThreadsResponse
}

// Returns a new `RestartFrameRequest` with the following fields set: `Command`, `Type`
func NewRestartFrameRequest() *RestartFrameRequest {
	newRestartFrameRequest := RestartFrameRequest{}
	newRestartFrameRequest.Command = "restartFrame"
	newRestartFrameRequest.Type = "request"
	newRestartFrameRequest.propagateFieldsToBase()
	return &newRestartFrameRequest
}

// Returns a new `SourceRequest` with the following fields set: `Command`, `Type`
func NewSourceRequest() *SourceRequest {
	newSourceRequest := SourceRequest{}
	newSourceRequest.Command = "source"
	newSourceRequest.Type = "request"
	newSourceRequest.propagateFieldsToBase()
	return &newSourceRequest
}

// Returns a new `GotoTargetsRequest` with the following fields set: `Command`, `Type`
func NewGotoTargetsRequest() *GotoTargetsRequest {
	newGotoTargetsRequest := GotoTargetsRequest{}
	newGotoTargetsRequest.Command = "gotoTargets"
	newGotoTargetsRequest.Type = "request"
	newGotoTargetsRequest.propagateFieldsToBase()
	return &newGotoTargetsRequest
}

// Returns a new `SetDataBreakpointsResponse` with the following fields set: `Command`, `Type`
func NewSetDataBreakpointsResponse() *SetDataBreakpointsResponse {
	newSetDataBreakpointsResponse := SetDataBreakpointsResponse{}
	newSetDataBreakpointsResponse.Command = "setDataBreakpoints"
	newSetDataBreakpointsResponse.Type = "response"
	newSetDataBreakpointsResponse.propagateFieldsToBase()
	return &newSetDataBreakpointsResponse
}

// Returns a new `AttachResponse` with the following fields set: `Command`, `Type`
func NewAttachResponse() *AttachResponse {
	newAttachResponse := AttachResponse{}
	newAttachResponse.Command = "attach"
	newAttachResponse.Type = "response"
	newAttachResponse.propagateFieldsToBase()
	return &newAttachResponse
}

// Returns a new `EvaluateRequest` with the following fields set: `Command`, `Type`
func NewEvaluateRequest() *EvaluateRequest {
	newEvaluateRequest := EvaluateRequest{}
	newEvaluateRequest.Command = "evaluate"
	newEvaluateRequest.Type = "request"
	newEvaluateRequest.propagateFieldsToBase()
	return &newEvaluateRequest
}

// Returns a new `VariablesRequest` with the following fields set: `Command`, `Type`
func NewVariablesRequest() *VariablesRequest {
	newVariablesRequest := VariablesRequest{}
	newVariablesRequest.Command = "variables"
	newVariablesRequest.Type = "request"
	newVariablesRequest.propagateFieldsToBase()
	return &newVariablesRequest
}

// Returns a new `LaunchRequest` with the following fields set: `Command`, `Type`
func NewLaunchRequest() *LaunchRequest {
	newLaunchRequest := LaunchRequest{}
	newLaunchRequest.Command = "launch"
	newLaunchRequest.Type = "request"
	newLaunchRequest.propagateFieldsToBase()
	return &newLaunchRequest
}

// Returns a new `AttachRequest` with the following fields set: `Command`, `Type`
func NewAttachRequest() *AttachRequest {
	newAttachRequest := AttachRequest{}
	newAttachRequest.Command = "attach"
	newAttachRequest.Type = "request"
	newAttachRequest.propagateFieldsToBase()
	return &newAttachRequest
}

// Returns a new `ContinueRequest` with the following fields set: `Command`, `Type`
func NewContinueRequest() *ContinueRequest {
	newContinueRequest := ContinueRequest{}
	newContinueRequest.Command = "continue"
	newContinueRequest.Type = "request"
	newContinueRequest.propagateFieldsToBase()
	return &newContinueRequest
}

// Returns a new `LoadedSourceEvent` with the following fields set: `Event_`, `Type`
func NewLoadedSourceEvent() *LoadedSourceEvent {
	newLoadedSourceEvent := LoadedSourceEvent{}
	newLoadedSourceEvent.Event_ = "loadedSource"
	newLoadedSourceEvent.Type = "event"
	newLoadedSourceEvent.propagateFieldsToBase()
	return &newLoadedSourceEvent
}

// Returns a new `StepOutResponse` with the following fields set: `Command`, `Type`
func NewStepOutResponse() *StepOutResponse {
	newStepOutResponse := StepOutResponse{}
	newStepOutResponse.Command = "stepOut"
	newStepOutResponse.Type = "response"
	newStepOutResponse.propagateFieldsToBase()
	return &newStepOutResponse
}

// Returns a new `DisconnectResponse` with the following fields set: `Command`, `Type`
func NewDisconnectResponse() *DisconnectResponse {
	newDisconnectResponse := DisconnectResponse{}
	newDisconnectResponse.Command = "disconnect"
	newDisconnectResponse.Type = "response"
	newDisconnectResponse.propagateFieldsToBase()
	return &newDisconnectResponse
}

// Returns a new `ReverseContinueResponse` with the following fields set: `Command`, `Type`
func NewReverseContinueResponse() *ReverseContinueResponse {
	newReverseContinueResponse := ReverseContinueResponse{}
	newReverseContinueResponse.Command = "reverseContinue"
	newReverseContinueResponse.Type = "response"
	newReverseContinueResponse.propagateFieldsToBase()
	return &newReverseContinueResponse
}

// Returns a new `LaunchResponse` with the following fields set: `Command`, `Type`
func NewLaunchResponse() *LaunchResponse {
	newLaunchResponse := LaunchResponse{}
	newLaunchResponse.Command = "launch"
	newLaunchResponse.Type = "response"
	newLaunchResponse.propagateFieldsToBase()
	return &newLaunchResponse
}

// Returns a new `ScopesResponse` with the following fields set: `Command`, `Type`
func NewScopesResponse() *ScopesResponse {
	newScopesResponse := ScopesResponse{}
	newScopesResponse.Command = "scopes"
	newScopesResponse.Type = "response"
	newScopesResponse.propagateFieldsToBase()
	return &newScopesResponse
}

// Returns a new `InitializedEvent` with the following fields set: `Event_`, `Type`
func NewInitializedEvent() *InitializedEvent {
	newInitializedEvent := InitializedEvent{}
	newInitializedEvent.Event_ = "initialized"
	newInitializedEvent.Type = "event"
	newInitializedEvent.propagateFieldsToBase()
	return &newInitializedEvent
}

// Returns a new `StepInTargetsRequest` with the following fields set: `Command`, `Type`
func NewStepInTargetsRequest() *StepInTargetsRequest {
	newStepInTargetsRequest := StepInTargetsRequest{}
	newStepInTargetsRequest.Command = "stepInTargets"
	newStepInTargetsRequest.Type = "request"
	newStepInTargetsRequest.propagateFieldsToBase()
	return &newStepInTargetsRequest
}

// Returns a new `RestartFrameResponse` with the following fields set: `Command`, `Type`
func NewRestartFrameResponse() *RestartFrameResponse {
	newRestartFrameResponse := RestartFrameResponse{}
	newRestartFrameResponse.Command = "restartFrame"
	newRestartFrameResponse.Type = "response"
	newRestartFrameResponse.propagateFieldsToBase()
	return &newRestartFrameResponse
}

// Returns a new `SetExpressionResponse` with the following fields set: `Command`, `Type`
func NewSetExpressionResponse() *SetExpressionResponse {
	newSetExpressionResponse := SetExpressionResponse{}
	newSetExpressionResponse.Command = "setExpression"
	newSetExpressionResponse.Type = "response"
	newSetExpressionResponse.propagateFieldsToBase()
	return &newSetExpressionResponse
}

// Returns a new `NextRequest` with the following fields set: `Command`, `Type`
func NewNextRequest() *NextRequest {
	newNextRequest := NextRequest{}
	newNextRequest.Command = "next"
	newNextRequest.Type = "request"
	newNextRequest.propagateFieldsToBase()
	return &newNextRequest
}

// Returns a new `ScopesRequest` with the following fields set: `Command`, `Type`
func NewScopesRequest() *ScopesRequest {
	newScopesRequest := ScopesRequest{}
	newScopesRequest.Command = "scopes"
	newScopesRequest.Type = "request"
	newScopesRequest.propagateFieldsToBase()
	return &newScopesRequest
}

// Returns a new `ReadMemoryRequest` with the following fields set: `Command`, `Type`
func NewReadMemoryRequest() *ReadMemoryRequest {
	newReadMemoryRequest := ReadMemoryRequest{}
	newReadMemoryRequest.Command = "readMemory"
	newReadMemoryRequest.Type = "request"
	newReadMemoryRequest.propagateFieldsToBase()
	return &newReadMemoryRequest
}

// Returns a new `RestartResponse` with the following fields set: `Command`, `Type`
func NewRestartResponse() *RestartResponse {
	newRestartResponse := RestartResponse{}
	newRestartResponse.Command = "restart"
	newRestartResponse.Type = "response"
	newRestartResponse.propagateFieldsToBase()
	return &newRestartResponse
}

// Returns a new `StackTraceRequest` with the following fields set: `Command`, `Type`
func NewStackTraceRequest() *StackTraceRequest {
	newStackTraceRequest := StackTraceRequest{}
	newStackTraceRequest.Command = "stackTrace"
	newStackTraceRequest.Type = "request"
	newStackTraceRequest.propagateFieldsToBase()
	return &newStackTraceRequest
}

// Returns a new `StepBackRequest` with the following fields set: `Command`, `Type`
func NewStepBackRequest() *StepBackRequest {
	newStepBackRequest := StepBackRequest{}
	newStepBackRequest.Command = "stepBack"
	newStepBackRequest.Type = "request"
	newStepBackRequest.propagateFieldsToBase()
	return &newStepBackRequest
}

// Returns a new `RunInTerminalRequest` with the following fields set: `Command`, `Type`
func NewRunInTerminalRequest() *RunInTerminalRequest {
	newRunInTerminalRequest := RunInTerminalRequest{}
	newRunInTerminalRequest.Command = "runInTerminal"
	newRunInTerminalRequest.Type = "request"
	newRunInTerminalRequest.propagateFieldsToBase()
	return &newRunInTerminalRequest
}

// Returns a new `LoadedSourcesResponse` with the following fields set: `Command`, `Type`
func NewLoadedSourcesResponse() *LoadedSourcesResponse {
	newLoadedSourcesResponse := LoadedSourcesResponse{}
	newLoadedSourcesResponse.Command = "loadedSources"
	newLoadedSourcesResponse.Type = "response"
	newLoadedSourcesResponse.propagateFieldsToBase()
	return &newLoadedSourcesResponse
}

// Returns a new `ErrorResponse` with the following fields set: `Type`
func NewErrorResponse() *ErrorResponse {
	newErrorResponse := ErrorResponse{}
	newErrorResponse.Type = "response"
	newErrorResponse.propagateFieldsToBase()
	return &newErrorResponse
}

// Returns a new `ConfigurationDoneResponse` with the following fields set: `Command`, `Type`
func NewConfigurationDoneResponse() *ConfigurationDoneResponse {
	newConfigurationDoneResponse := ConfigurationDoneResponse{}
	newConfigurationDoneResponse.Command = "configurationDone"
	newConfigurationDoneResponse.Type = "response"
	newConfigurationDoneResponse.propagateFieldsToBase()
	return &newConfigurationDoneResponse
}

// Returns a new `LoadedSourcesRequest` with the following fields set: `Command`, `Type`
func NewLoadedSourcesRequest() *LoadedSourcesRequest {
	newLoadedSourcesRequest := LoadedSourcesRequest{}
	newLoadedSourcesRequest.Command = "loadedSources"
	newLoadedSourcesRequest.Type = "request"
	newLoadedSourcesRequest.propagateFieldsToBase()
	return &newLoadedSourcesRequest
}

// Returns a new `StackTraceResponse` with the following fields set: `Command`, `Type`
func NewStackTraceResponse() *StackTraceResponse {
	newStackTraceResponse := StackTraceResponse{}
	newStackTraceResponse.Command = "stackTrace"
	newStackTraceResponse.Type = "response"
	newStackTraceResponse.propagateFieldsToBase()
	return &newStackTraceResponse
}

// Returns a new `GotoResponse` with the following fields set: `Command`, `Type`
func NewGotoResponse() *GotoResponse {
	newGotoResponse := GotoResponse{}
	newGotoResponse.Command = "goto"
	newGotoResponse.Type = "response"
	newGotoResponse.propagateFieldsToBase()
	return &newGotoResponse
}

// Returns a new `ModuleEvent` with the following fields set: `Event_`, `Type`
func NewModuleEvent() *ModuleEvent {
	newModuleEvent := ModuleEvent{}
	newModuleEvent.Event_ = "module"
	newModuleEvent.Type = "event"
	newModuleEvent.propagateFieldsToBase()
	return &newModuleEvent
}

// Returns a new `ProcessEvent` with the following fields set: `Event_`, `Type`
func NewProcessEvent() *ProcessEvent {
	newProcessEvent := ProcessEvent{}
	newProcessEvent.Event_ = "process"
	newProcessEvent.Type = "event"
	newProcessEvent.propagateFieldsToBase()
	return &newProcessEvent
}

// Returns a new `StepInResponse` with the following fields set: `Command`, `Type`
func NewStepInResponse() *StepInResponse {
	newStepInResponse := StepInResponse{}
	newStepInResponse.Command = "stepIn"
	newStepInResponse.Type = "response"
	newStepInResponse.propagateFieldsToBase()
	return &newStepInResponse
}

// Returns a new `DataBreakpointInfoResponse` with the following fields set: `Command`, `Type`
func NewDataBreakpointInfoResponse() *DataBreakpointInfoResponse {
	newDataBreakpointInfoResponse := DataBreakpointInfoResponse{}
	newDataBreakpointInfoResponse.Command = "dataBreakpointInfo"
	newDataBreakpointInfoResponse.Type = "response"
	newDataBreakpointInfoResponse.propagateFieldsToBase()
	return &newDataBreakpointInfoResponse
}

// Returns a new `TerminateRequest` with the following fields set: `Command`, `Type`
func NewTerminateRequest() *TerminateRequest {
	newTerminateRequest := TerminateRequest{}
	newTerminateRequest.Command = "terminate"
	newTerminateRequest.Type = "request"
	newTerminateRequest.propagateFieldsToBase()
	return &newTerminateRequest
}

// Returns a new `InitializeRequest` with the following fields set: `Command`, `Type`
func NewInitializeRequest() *InitializeRequest {
	newInitializeRequest := InitializeRequest{}
	newInitializeRequest.Command = "initialize"
	newInitializeRequest.Type = "request"
	newInitializeRequest.propagateFieldsToBase()
	return &newInitializeRequest
}

// Returns a new `DisassembleRequest` with the following fields set: `Command`, `Type`
func NewDisassembleRequest() *DisassembleRequest {
	newDisassembleRequest := DisassembleRequest{}
	newDisassembleRequest.Command = "disassemble"
	newDisassembleRequest.Type = "request"
	newDisassembleRequest.propagateFieldsToBase()
	return &newDisassembleRequest
}

// Returns a new `StepInRequest` with the following fields set: `Command`, `Type`
func NewStepInRequest() *StepInRequest {
	newStepInRequest := StepInRequest{}
	newStepInRequest.Command = "stepIn"
	newStepInRequest.Type = "request"
	newStepInRequest.propagateFieldsToBase()
	return &newStepInRequest
}

// Returns a new `RunInTerminalResponse` with the following fields set: `Command`, `Type`
func NewRunInTerminalResponse() *RunInTerminalResponse {
	newRunInTerminalResponse := RunInTerminalResponse{}
	newRunInTerminalResponse.Command = "runInTerminal"
	newRunInTerminalResponse.Type = "response"
	newRunInTerminalResponse.propagateFieldsToBase()
	return &newRunInTerminalResponse
}

// Returns a new `GotoRequest` with the following fields set: `Command`, `Type`
func NewGotoRequest() *GotoRequest {
	newGotoRequest := GotoRequest{}
	newGotoRequest.Command = "goto"
	newGotoRequest.Type = "request"
	newGotoRequest.propagateFieldsToBase()
	return &newGotoRequest
}

// Returns a new `ReverseContinueRequest` with the following fields set: `Command`, `Type`
func NewReverseContinueRequest() *ReverseContinueRequest {
	newReverseContinueRequest := ReverseContinueRequest{}
	newReverseContinueRequest.Command = "reverseContinue"
	newReverseContinueRequest.Type = "request"
	newReverseContinueRequest.propagateFieldsToBase()
	return &newReverseContinueRequest
}

// Returns a new `SetBreakpointsRequest` with the following fields set: `Command`, `Type`
func NewSetBreakpointsRequest() *SetBreakpointsRequest {
	newSetBreakpointsRequest := SetBreakpointsRequest{}
	newSetBreakpointsRequest.Command = "setBreakpoints"
	newSetBreakpointsRequest.Type = "request"
	newSetBreakpointsRequest.propagateFieldsToBase()
	return &newSetBreakpointsRequest
}

// Returns a new `StoppedEvent` with the following fields set: `Event_`, `Type`
func NewStoppedEvent() *StoppedEvent {
	newStoppedEvent := StoppedEvent{}
	newStoppedEvent.Event_ = "stopped"
	newStoppedEvent.Type = "event"
	newStoppedEvent.propagateFieldsToBase()
	return &newStoppedEvent
}

// Returns a new `PauseRequest` with the following fields set: `Command`, `Type`
func NewPauseRequest() *PauseRequest {
	newPauseRequest := PauseRequest{}
	newPauseRequest.Command = "pause"
	newPauseRequest.Type = "request"
	newPauseRequest.propagateFieldsToBase()
	return &newPauseRequest
}

// Returns a new `DisassembleResponse` with the following fields set: `Command`, `Type`
func NewDisassembleResponse() *DisassembleResponse {
	newDisassembleResponse := DisassembleResponse{}
	newDisassembleResponse.Command = "disassemble"
	newDisassembleResponse.Type = "response"
	newDisassembleResponse.propagateFieldsToBase()
	return &newDisassembleResponse
}

// Returns a new `ContinueResponse` with the following fields set: `Command`, `Type`
func NewContinueResponse() *ContinueResponse {
	newContinueResponse := ContinueResponse{}
	newContinueResponse.Command = "continue"
	newContinueResponse.Type = "response"
	newContinueResponse.propagateFieldsToBase()
	return &newContinueResponse
}

// Returns a new `ExceptionInfoResponse` with the following fields set: `Command`, `Type`
func NewExceptionInfoResponse() *ExceptionInfoResponse {
	newExceptionInfoResponse := ExceptionInfoResponse{}
	newExceptionInfoResponse.Command = "exceptionInfo"
	newExceptionInfoResponse.Type = "response"
	newExceptionInfoResponse.propagateFieldsToBase()
	return &newExceptionInfoResponse
}

// Returns a new `DisconnectRequest` with the following fields set: `Command`, `Type`
func NewDisconnectRequest() *DisconnectRequest {
	newDisconnectRequest := DisconnectRequest{}
	newDisconnectRequest.Command = "disconnect"
	newDisconnectRequest.Type = "request"
	newDisconnectRequest.propagateFieldsToBase()
	return &newDisconnectRequest
}

// Returns a new `ExceptionInfoRequest` with the following fields set: `Command`, `Type`
func NewExceptionInfoRequest() *ExceptionInfoRequest {
	newExceptionInfoRequest := ExceptionInfoRequest{}
	newExceptionInfoRequest.Command = "exceptionInfo"
	newExceptionInfoRequest.Type = "request"
	newExceptionInfoRequest.propagateFieldsToBase()
	return &newExceptionInfoRequest
}

// Returns a new `GotoTargetsResponse` with the following fields set: `Command`, `Type`
func NewGotoTargetsResponse() *GotoTargetsResponse {
	newGotoTargetsResponse := GotoTargetsResponse{}
	newGotoTargetsResponse.Command = "gotoTargets"
	newGotoTargetsResponse.Type = "response"
	newGotoTargetsResponse.propagateFieldsToBase()
	return &newGotoTargetsResponse
}

// Returns a new `PauseResponse` with the following fields set: `Command`, `Type`
func NewPauseResponse() *PauseResponse {
	newPauseResponse := PauseResponse{}
	newPauseResponse.Command = "pause"
	newPauseResponse.Type = "response"
	newPauseResponse.propagateFieldsToBase()
	return &newPauseResponse
}

// Returns a new `CompletionsRequest` with the following fields set: `Command`, `Type`
func NewCompletionsRequest() *CompletionsRequest {
	newCompletionsRequest := CompletionsRequest{}
	newCompletionsRequest.Command = "completions"
	newCompletionsRequest.Type = "request"
	newCompletionsRequest.propagateFieldsToBase()
	return &newCompletionsRequest
}

// Returns a new `OutputEvent` with the following fields set: `Event_`, `Type`
func NewOutputEvent() *OutputEvent {
	newOutputEvent := OutputEvent{}
	newOutputEvent.Event_ = "output"
	newOutputEvent.Type = "event"
	newOutputEvent.propagateFieldsToBase()
	return &newOutputEvent
}

// Returns a new `NextResponse` with the following fields set: `Command`, `Type`
func NewNextResponse() *NextResponse {
	newNextResponse := NextResponse{}
	newNextResponse.Command = "next"
	newNextResponse.Type = "response"
	newNextResponse.propagateFieldsToBase()
	return &newNextResponse
}

// Returns a new `EvaluateResponse` with the following fields set: `Command`, `Type`
func NewEvaluateResponse() *EvaluateResponse {
	newEvaluateResponse := EvaluateResponse{}
	newEvaluateResponse.Command = "evaluate"
	newEvaluateResponse.Type = "response"
	newEvaluateResponse.propagateFieldsToBase()
	return &newEvaluateResponse
}

// Returns a new `RestartRequest` with the following fields set: `Command`, `Type`
func NewRestartRequest() *RestartRequest {
	newRestartRequest := RestartRequest{}
	newRestartRequest.Command = "restart"
	newRestartRequest.Type = "request"
	newRestartRequest.propagateFieldsToBase()
	return &newRestartRequest
}

// Returns a new `TerminateThreadsRequest` with the following fields set: `Command`, `Type`
func NewTerminateThreadsRequest() *TerminateThreadsRequest {
	newTerminateThreadsRequest := TerminateThreadsRequest{}
	newTerminateThreadsRequest.Command = "terminateThreads"
	newTerminateThreadsRequest.Type = "request"
	newTerminateThreadsRequest.propagateFieldsToBase()
	return &newTerminateThreadsRequest
}

// Returns a new `StepBackResponse` with the following fields set: `Command`, `Type`
func NewStepBackResponse() *StepBackResponse {
	newStepBackResponse := StepBackResponse{}
	newStepBackResponse.Command = "stepBack"
	newStepBackResponse.Type = "response"
	newStepBackResponse.propagateFieldsToBase()
	return &newStepBackResponse
}

// Returns a new `InitializeResponse` with the following fields set: `Command`, `Type`
func NewInitializeResponse() *InitializeResponse {
	newInitializeResponse := InitializeResponse{}
	newInitializeResponse.Command = "initialize"
	newInitializeResponse.Type = "response"
	newInitializeResponse.propagateFieldsToBase()
	return &newInitializeResponse
}

// Returns a new `SourceResponse` with the following fields set: `Command`, `Type`
func NewSourceResponse() *SourceResponse {
	newSourceResponse := SourceResponse{}
	newSourceResponse.Command = "source"
	newSourceResponse.Type = "response"
	newSourceResponse.propagateFieldsToBase()
	return &newSourceResponse
}

// Returns a new `CompletionsResponse` with the following fields set: `Command`, `Type`
func NewCompletionsResponse() *CompletionsResponse {
	newCompletionsResponse := CompletionsResponse{}
	newCompletionsResponse.Command = "completions"
	newCompletionsResponse.Type = "response"
	newCompletionsResponse.propagateFieldsToBase()
	return &newCompletionsResponse
}

// Returns a new `SetFunctionBreakpointsRequest` with the following fields set: `Command`, `Type`
func NewSetFunctionBreakpointsRequest() *SetFunctionBreakpointsRequest {
	newSetFunctionBreakpointsRequest := SetFunctionBreakpointsRequest{}
	newSetFunctionBreakpointsRequest.Command = "setFunctionBreakpoints"
	newSetFunctionBreakpointsRequest.Type = "request"
	newSetFunctionBreakpointsRequest.propagateFieldsToBase()
	return &newSetFunctionBreakpointsRequest
}

// Returns a new `SetExceptionBreakpointsResponse` with the following fields set: `Command`, `Type`
func NewSetExceptionBreakpointsResponse() *SetExceptionBreakpointsResponse {
	newSetExceptionBreakpointsResponse := SetExceptionBreakpointsResponse{}
	newSetExceptionBreakpointsResponse.Command = "setExceptionBreakpoints"
	newSetExceptionBreakpointsResponse.Type = "response"
	newSetExceptionBreakpointsResponse.propagateFieldsToBase()
	return &newSetExceptionBreakpointsResponse
}

// Returns a new `ReadMemoryResponse` with the following fields set: `Command`, `Type`
func NewReadMemoryResponse() *ReadMemoryResponse {
	newReadMemoryResponse := ReadMemoryResponse{}
	newReadMemoryResponse.Command = "readMemory"
	newReadMemoryResponse.Type = "response"
	newReadMemoryResponse.propagateFieldsToBase()
	return &newReadMemoryResponse
}

// Returns a new `ModulesRequest` with the following fields set: `Command`, `Type`
func NewModulesRequest() *ModulesRequest {
	newModulesRequest := ModulesRequest{}
	newModulesRequest.Command = "modules"
	newModulesRequest.Type = "request"
	newModulesRequest.propagateFieldsToBase()
	return &newModulesRequest
}

// Returns a new `StepOutRequest` with the following fields set: `Command`, `Type`
func NewStepOutRequest() *StepOutRequest {
	newStepOutRequest := StepOutRequest{}
	newStepOutRequest.Command = "stepOut"
	newStepOutRequest.Type = "request"
	newStepOutRequest.propagateFieldsToBase()
	return &newStepOutRequest
}

// Returns a new `ModulesResponse` with the following fields set: `Command`, `Type`
func NewModulesResponse() *ModulesResponse {
	newModulesResponse := ModulesResponse{}
	newModulesResponse.Command = "modules"
	newModulesResponse.Type = "response"
	newModulesResponse.propagateFieldsToBase()
	return &newModulesResponse
}

// Returns a new `SetVariableRequest` with the following fields set: `Command`, `Type`
func NewSetVariableRequest() *SetVariableRequest {
	newSetVariableRequest := SetVariableRequest{}
	newSetVariableRequest.Command = "setVariable"
	newSetVariableRequest.Type = "request"
	newSetVariableRequest.propagateFieldsToBase()
	return &newSetVariableRequest
}


// TryUnmarshalProtocolMessage attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `ProtocolMessage` as follows:
// 
// If `js` contains `"type":"request"`, attempts to unmarshal via `TryUnmarshalRequest`
// 
// If `js` contains `"type":"event"`, attempts to unmarshal via `TryUnmarshalEvent`
// 
// If `js` contains `"type":"response"`, attempts to unmarshal via `TryUnmarshalResponse`
// 
// Otherwise, `err`'s message will be: `ProtocolMessage: encountered unknown JSON value for type: ` followed by the `type` value encountered.
// 
// In general: the `err` returned may be either `nil`, the above message, or an `encoding/json.Unmarshal()` return value.
// `ptr` will be a pointer to the unmarshaled `struct` value if that succeeded, else `nil`.
// Both `err` and `ptr` will be `nil` if `js` doesn't: start with `{` and end with `}` and contain `"type":"` followed by a subsequent `"`.
func TryUnmarshalProtocolMessage (js string) (ptr interface{}, err error) {
	if len(js)==0 || js[0]!='{' || js[len(js)-1]!='}' { return }
	i1 := strings.Index(js, "\"type\":\"")  ;  if i1<1 { return }
	subjs := js[i1+4+4:]
	i2 := strings.Index(subjs, "\"")  ;  if i2<1 { return }
	type_of_ProtocolMessage := subjs[:i2]  ;  switch type_of_ProtocolMessage {
	case "request":  ptr,err = TryUnmarshalRequest(js)
	case "event":  ptr,err = TryUnmarshalEvent(js)
	case "response":  ptr,err = TryUnmarshalResponse(js)
	default: err = errors.New("ProtocolMessage: encountered unknown JSON value for type: " + type_of_ProtocolMessage)
	}
	return
}


// TryUnmarshalEvent attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `Event` as follows:
// 
// If `js` contains `"event":"capabilities"`, attempts to unmarshal into a new `CapabilitiesEvent`.
// 
// If `js` contains `"event":"output"`, attempts to unmarshal into a new `OutputEvent`.
// 
// If `js` contains `"event":"loadedSource"`, attempts to unmarshal into a new `LoadedSourceEvent`.
// 
// If `js` contains `"event":"process"`, attempts to unmarshal into a new `ProcessEvent`.
// 
// If `js` contains `"event":"terminated"`, attempts to unmarshal into a new `TerminatedEvent`.
// 
// If `js` contains `"event":"thread"`, attempts to unmarshal into a new `ThreadEvent`.
// 
// If `js` contains `"event":"exited"`, attempts to unmarshal into a new `ExitedEvent`.
// 
// If `js` contains `"event":"initialized"`, attempts to unmarshal into a new `InitializedEvent`.
// 
// If `js` contains `"event":"continued"`, attempts to unmarshal into a new `ContinuedEvent`.
// 
// If `js` contains `"event":"breakpoint"`, attempts to unmarshal into a new `BreakpointEvent`.
// 
// If `js` contains `"event":"stopped"`, attempts to unmarshal into a new `StoppedEvent`.
// 
// If `js` contains `"event":"module"`, attempts to unmarshal into a new `ModuleEvent`.
// 
// Otherwise, `err`'s message will be: `Event: encountered unknown JSON value for event: ` followed by the `event` value encountered.
// 
// In general: the `err` returned may be either `nil`, the above message, or an `encoding/json.Unmarshal()` return value.
// `ptr` will be a pointer to the unmarshaled `struct` value if that succeeded, else `nil`.
// Both `err` and `ptr` will be `nil` if `js` doesn't: start with `{` and end with `}` and contain `"event":"` followed by a subsequent `"`.
func TryUnmarshalEvent (js string) (ptr interface{}, err error) {
	if len(js)==0 || js[0]!='{' || js[len(js)-1]!='}' { return }
	i1 := strings.Index(js, "\"event\":\"")  ;  if i1<1 { return }
	subjs := js[i1+4+5:]
	i2 := strings.Index(subjs, "\"")  ;  if i2<1 { return }
	event_of_Event := subjs[:i2]  ;  switch event_of_Event {
	case "breakpoint":  var val BreakpointEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stopped":  var val StoppedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "module":  var val ModuleEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "continued":  var val ContinuedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "thread":  var val ThreadEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "exited":  var val ExitedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "initialized":  var val InitializedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "capabilities":  var val CapabilitiesEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "output":  var val OutputEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "loadedSource":  var val LoadedSourceEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "process":  var val ProcessEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminated":  var val TerminatedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	default: err = errors.New("Event: encountered unknown JSON value for event: " + event_of_Event)
	}
	return
}


// TryUnmarshalRequest attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `Request` as follows:
// 
// If `js` contains `"command":"setFunctionBreakpoints"`, attempts to unmarshal into a new `SetFunctionBreakpointsRequest`.
// 
// If `js` contains `"command":"loadedSources"`, attempts to unmarshal into a new `LoadedSourcesRequest`.
// 
// If `js` contains `"command":"configurationDone"`, attempts to unmarshal into a new `ConfigurationDoneRequest`.
// 
// If `js` contains `"command":"disconnect"`, attempts to unmarshal into a new `DisconnectRequest`.
// 
// If `js` contains `"command":"gotoTargets"`, attempts to unmarshal into a new `GotoTargetsRequest`.
// 
// If `js` contains `"command":"setDataBreakpoints"`, attempts to unmarshal into a new `SetDataBreakpointsRequest`.
// 
// If `js` contains `"command":"initialize"`, attempts to unmarshal into a new `InitializeRequest`.
// 
// If `js` contains `"command":"modules"`, attempts to unmarshal into a new `ModulesRequest`.
// 
// If `js` contains `"command":"terminate"`, attempts to unmarshal into a new `TerminateRequest`.
// 
// If `js` contains `"command":"dataBreakpointInfo"`, attempts to unmarshal into a new `DataBreakpointInfoRequest`.
// 
// If `js` contains `"command":"variables"`, attempts to unmarshal into a new `VariablesRequest`.
// 
// If `js` contains `"command":"launch"`, attempts to unmarshal into a new `LaunchRequest`.
// 
// If `js` contains `"command":"scopes"`, attempts to unmarshal into a new `ScopesRequest`.
// 
// If `js` contains `"command":"stepBack"`, attempts to unmarshal into a new `StepBackRequest`.
// 
// If `js` contains `"command":"completions"`, attempts to unmarshal into a new `CompletionsRequest`.
// 
// If `js` contains `"command":"continue"`, attempts to unmarshal into a new `ContinueRequest`.
// 
// If `js` contains `"command":"setExpression"`, attempts to unmarshal into a new `SetExpressionRequest`.
// 
// If `js` contains `"command":"exceptionInfo"`, attempts to unmarshal into a new `ExceptionInfoRequest`.
// 
// If `js` contains `"command":"stepInTargets"`, attempts to unmarshal into a new `StepInTargetsRequest`.
// 
// If `js` contains `"command":"goto"`, attempts to unmarshal into a new `GotoRequest`.
// 
// If `js` contains `"command":"setVariable"`, attempts to unmarshal into a new `SetVariableRequest`.
// 
// If `js` contains `"command":"terminateThreads"`, attempts to unmarshal into a new `TerminateThreadsRequest`.
// 
// If `js` contains `"command":"setExceptionBreakpoints"`, attempts to unmarshal into a new `SetExceptionBreakpointsRequest`.
// 
// If `js` contains `"command":"source"`, attempts to unmarshal into a new `SourceRequest`.
// 
// If `js` contains `"command":"evaluate"`, attempts to unmarshal into a new `EvaluateRequest`.
// 
// If `js` contains `"command":"next"`, attempts to unmarshal into a new `NextRequest`.
// 
// If `js` contains `"command":"stackTrace"`, attempts to unmarshal into a new `StackTraceRequest`.
// 
// If `js` contains `"command":"runInTerminal"`, attempts to unmarshal into a new `RunInTerminalRequest`.
// 
// If `js` contains `"command":"restartFrame"`, attempts to unmarshal into a new `RestartFrameRequest`.
// 
// If `js` contains `"command":"readMemory"`, attempts to unmarshal into a new `ReadMemoryRequest`.
// 
// If `js` contains `"command":"attach"`, attempts to unmarshal into a new `AttachRequest`.
// 
// If `js` contains `"command":"setBreakpoints"`, attempts to unmarshal into a new `SetBreakpointsRequest`.
// 
// If `js` contains `"command":"pause"`, attempts to unmarshal into a new `PauseRequest`.
// 
// If `js` contains `"command":"stepOut"`, attempts to unmarshal into a new `StepOutRequest`.
// 
// If `js` contains `"command":"disassemble"`, attempts to unmarshal into a new `DisassembleRequest`.
// 
// If `js` contains `"command":"stepIn"`, attempts to unmarshal into a new `StepInRequest`.
// 
// If `js` contains `"command":"restart"`, attempts to unmarshal into a new `RestartRequest`.
// 
// If `js` contains `"command":"threads"`, attempts to unmarshal into a new `ThreadsRequest`.
// 
// If `js` contains `"command":"reverseContinue"`, attempts to unmarshal into a new `ReverseContinueRequest`.
// 
// Otherwise, `err`'s message will be: `Request: encountered unknown JSON value for command: ` followed by the `command` value encountered.
// 
// In general: the `err` returned may be either `nil`, the above message, or an `encoding/json.Unmarshal()` return value.
// `ptr` will be a pointer to the unmarshaled `struct` value if that succeeded, else `nil`.
// Both `err` and `ptr` will be `nil` if `js` doesn't: start with `{` and end with `}` and contain `"command":"` followed by a subsequent `"`.
func TryUnmarshalRequest (js string) (ptr interface{}, err error) {
	if len(js)==0 || js[0]!='{' || js[len(js)-1]!='}' { return }
	i1 := strings.Index(js, "\"command\":\"")  ;  if i1<1 { return }
	subjs := js[i1+4+7:]
	i2 := strings.Index(subjs, "\"")  ;  if i2<1 { return }
	command_of_Request := subjs[:i2]  ;  switch command_of_Request {
	case "goto":  var val GotoRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setVariable":  var val SetVariableRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminateThreads":  var val TerminateThreadsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setExceptionBreakpoints":  var val SetExceptionBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "exceptionInfo":  var val ExceptionInfoRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepInTargets":  var val StepInTargetsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "next":  var val NextRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stackTrace":  var val StackTraceRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "runInTerminal":  var val RunInTerminalRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "source":  var val SourceRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "evaluate":  var val EvaluateRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "attach":  var val AttachRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setBreakpoints":  var val SetBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restartFrame":  var val RestartFrameRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "readMemory":  var val ReadMemoryRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "disassemble":  var val DisassembleRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepIn":  var val StepInRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restart":  var val RestartRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "threads":  var val ThreadsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "reverseContinue":  var val ReverseContinueRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "pause":  var val PauseRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepOut":  var val StepOutRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "configurationDone":  var val ConfigurationDoneRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "disconnect":  var val DisconnectRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setFunctionBreakpoints":  var val SetFunctionBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "loadedSources":  var val LoadedSourcesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "initialize":  var val InitializeRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "modules":  var val ModulesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminate":  var val TerminateRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "gotoTargets":  var val GotoTargetsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setDataBreakpoints":  var val SetDataBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "launch":  var val LaunchRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "scopes":  var val ScopesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepBack":  var val StepBackRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "completions":  var val CompletionsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "dataBreakpointInfo":  var val DataBreakpointInfoRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "variables":  var val VariablesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "continue":  var val ContinueRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setExpression":  var val SetExpressionRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	default: err = errors.New("Request: encountered unknown JSON value for command: " + command_of_Request)
	}
	return
}


// TryUnmarshalResponse attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `Response` as follows:
// 
// If `js` contains `"command":"scopes"`, attempts to unmarshal into a new `ScopesResponse`.
// 
// If `js` contains `"command":"setExpression"`, attempts to unmarshal into a new `SetExpressionResponse`.
// 
// If `js` contains `"command":"stepBack"`, attempts to unmarshal into a new `StepBackResponse`.
// 
// If `js` contains `"command":"goto"`, attempts to unmarshal into a new `GotoResponse`.
// 
// If `js` contains `"command":"loadedSources"`, attempts to unmarshal into a new `LoadedSourcesResponse`.
// 
// If `js` contains `"command":"disassemble"`, attempts to unmarshal into a new `DisassembleResponse`.
// 
// If `js` contains `"command":"variables"`, attempts to unmarshal into a new `VariablesResponse`.
// 
// If `js` contains `"command":"disconnect"`, attempts to unmarshal into a new `DisconnectResponse`.
// 
// If `js` contains `"command":"restart"`, attempts to unmarshal into a new `RestartResponse`.
// 
// If `js` contains `"command":"stepOut"`, attempts to unmarshal into a new `StepOutResponse`.
// 
// If `js` contains `"command":"pause"`, attempts to unmarshal into a new `PauseResponse`.
// 
// If `js` contains `"command":"continue"`, attempts to unmarshal into a new `ContinueResponse`.
// 
// If `js` contains `"command":"stepIn"`, attempts to unmarshal into a new `StepInResponse`.
// 
// If `js` contains `"command":"source"`, attempts to unmarshal into a new `SourceResponse`.
// 
// If `js` contains `"command":"completions"`, attempts to unmarshal into a new `CompletionsResponse`.
// 
// If `js` contains `"command":"reverseContinue"`, attempts to unmarshal into a new `ReverseContinueResponse`.
// 
// If `js` contains `"command":"next"`, attempts to unmarshal into a new `NextResponse`.
// 
// If `js` contains `"command":"stackTrace"`, attempts to unmarshal into a new `StackTraceResponse`.
// 
// If `js` contains `"command":"setDataBreakpoints"`, attempts to unmarshal into a new `SetDataBreakpointsResponse`.
// 
// If `js` contains `"command":"dataBreakpointInfo"`, attempts to unmarshal into a new `DataBreakpointInfoResponse`.
// 
// If `js` contains `"command":"gotoTargets"`, attempts to unmarshal into a new `GotoTargetsResponse`.
// 
// If `js` contains `"command":"restartFrame"`, attempts to unmarshal into a new `RestartFrameResponse`.
// 
// If `js` contains `"command":"terminate"`, attempts to unmarshal into a new `TerminateResponse`.
// 
// If `js` contains `"command":"initialize"`, attempts to unmarshal into a new `InitializeResponse`.
// 
// If `js` contains `"command":"configurationDone"`, attempts to unmarshal into a new `ConfigurationDoneResponse`.
// 
// If `js` contains `"command":"setExceptionBreakpoints"`, attempts to unmarshal into a new `SetExceptionBreakpointsResponse`.
// 
// If `js` contains `"command":"terminateThreads"`, attempts to unmarshal into a new `TerminateThreadsResponse`.
// 
// If `js` contains `"command":"modules"`, attempts to unmarshal into a new `ModulesResponse`.
// 
// If `js` contains `"command":"exceptionInfo"`, attempts to unmarshal into a new `ExceptionInfoResponse`.
// 
// If `js` contains `"command":"evaluate"`, attempts to unmarshal into a new `EvaluateResponse`.
// 
// If `js` contains `"command":"launch"`, attempts to unmarshal into a new `LaunchResponse`.
// 
// If `js` contains `"command":"runInTerminal"`, attempts to unmarshal into a new `RunInTerminalResponse`.
// 
// If `js` contains `"command":"stepInTargets"`, attempts to unmarshal into a new `StepInTargetsResponse`.
// 
// If `js` contains `"command":"threads"`, attempts to unmarshal into a new `ThreadsResponse`.
// 
// If `js` contains `"command":"setFunctionBreakpoints"`, attempts to unmarshal into a new `SetFunctionBreakpointsResponse`.
// 
// If `js` contains `"command":"readMemory"`, attempts to unmarshal into a new `ReadMemoryResponse`.
// 
// If `js` contains `"command":"attach"`, attempts to unmarshal into a new `AttachResponse`.
// 
// If `js` contains `"command":"setBreakpoints"`, attempts to unmarshal into a new `SetBreakpointsResponse`.
// 
// If `js` contains `"command":"setVariable"`, attempts to unmarshal into a new `SetVariableResponse`.
// 
// Otherwise, `err`'s message will be: `Response: encountered unknown JSON value for command: ` followed by the `command` value encountered.
// 
// In general: the `err` returned may be either `nil`, the above message, or an `encoding/json.Unmarshal()` return value.
// `ptr` will be a pointer to the unmarshaled `struct` value if that succeeded, else `nil`.
// Both `err` and `ptr` will be `nil` if `js` doesn't: start with `{` and end with `}` and contain `"command":"` followed by a subsequent `"`.
func TryUnmarshalResponse (js string) (ptr interface{}, err error) {
	if len(js)==0 || js[0]!='{' || js[len(js)-1]!='}' { return }
	i1 := strings.Index(js, "\"command\":\"")  ;  if i1<1 { return }
	subjs := js[i1+4+7:]
	i2 := strings.Index(subjs, "\"")  ;  if i2<1 { return }
	command_of_Response := subjs[:i2]  ;  switch command_of_Response {
	case "setFunctionBreakpoints":  var val SetFunctionBreakpointsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "readMemory":  var val ReadMemoryResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "attach":  var val AttachResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setBreakpoints":  var val SetBreakpointsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setVariable":  var val SetVariableResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "threads":  var val ThreadsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setExpression":  var val SetExpressionResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepBack":  var val StepBackResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "goto":  var val GotoResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "loadedSources":  var val LoadedSourcesResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "disassemble":  var val DisassembleResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "variables":  var val VariablesResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "disconnect":  var val DisconnectResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "scopes":  var val ScopesResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepOut":  var val StepOutResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restart":  var val RestartResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "continue":  var val ContinueResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepIn":  var val StepInResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "source":  var val SourceResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "completions":  var val CompletionsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "reverseContinue":  var val ReverseContinueResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "next":  var val NextResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stackTrace":  var val StackTraceResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "pause":  var val PauseResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "dataBreakpointInfo":  var val DataBreakpointInfoResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setDataBreakpoints":  var val SetDataBreakpointsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restartFrame":  var val RestartFrameResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminate":  var val TerminateResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "initialize":  var val InitializeResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "gotoTargets":  var val GotoTargetsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "configurationDone":  var val ConfigurationDoneResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminateThreads":  var val TerminateThreadsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "modules":  var val ModulesResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "exceptionInfo":  var val ExceptionInfoResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "evaluate":  var val EvaluateResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "launch":  var val LaunchResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setExceptionBreakpoints":  var val SetExceptionBreakpointsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepInTargets":  var val StepInTargetsResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "runInTerminal":  var val RunInTerminalResponse  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	default: err = errors.New("Response: encountered unknown JSON value for command: " + command_of_Response)
	}
	return
}

// Called by `HandleRequest` (after it unmarshaled the given `SetVariableRequest`) to further populate the given `SetVariableResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetVariableRequest func(*SetVariableRequest, *SetVariableResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `RestartFrameRequest`) to further populate the given `RestartFrameResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRestartFrameRequest func(*RestartFrameRequest, *RestartFrameResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ExceptionInfoRequest`) to further populate the given `ExceptionInfoResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnExceptionInfoRequest func(*ExceptionInfoRequest, *ExceptionInfoResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `RunInTerminalRequest`) to further populate the given `RunInTerminalResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRunInTerminalRequest func(*RunInTerminalRequest, *RunInTerminalResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SourceRequest`) to further populate the given `SourceResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSourceRequest func(*SourceRequest, *SourceResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `GotoRequest`) to further populate the given `GotoResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnGotoRequest func(*GotoRequest, *GotoResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `TerminateThreadsRequest`) to further populate the given `TerminateThreadsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnTerminateThreadsRequest func(*TerminateThreadsRequest, *TerminateThreadsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetFunctionBreakpointsRequest`) to further populate the given `SetFunctionBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetFunctionBreakpointsRequest func(*SetFunctionBreakpointsRequest, *SetFunctionBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `DisassembleRequest`) to further populate the given `DisassembleResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnDisassembleRequest func(*DisassembleRequest, *DisassembleResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `AttachRequest`) to further populate the given `AttachResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnAttachRequest func(*AttachRequest, *AttachResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ConfigurationDoneRequest`) to further populate the given `ConfigurationDoneResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnConfigurationDoneRequest func(*ConfigurationDoneRequest, *ConfigurationDoneResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ReverseContinueRequest`) to further populate the given `ReverseContinueResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnReverseContinueRequest func(*ReverseContinueRequest, *ReverseContinueResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `NextRequest`) to further populate the given `NextResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnNextRequest func(*NextRequest, *NextResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `RestartRequest`) to further populate the given `RestartResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRestartRequest func(*RestartRequest, *RestartResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ThreadsRequest`) to further populate the given `ThreadsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnThreadsRequest func(*ThreadsRequest, *ThreadsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `DataBreakpointInfoRequest`) to further populate the given `DataBreakpointInfoResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnDataBreakpointInfoRequest func(*DataBreakpointInfoRequest, *DataBreakpointInfoResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `GotoTargetsRequest`) to further populate the given `GotoTargetsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnGotoTargetsRequest func(*GotoTargetsRequest, *GotoTargetsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepInTargetsRequest`) to further populate the given `StepInTargetsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepInTargetsRequest func(*StepInTargetsRequest, *StepInTargetsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepOutRequest`) to further populate the given `StepOutResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepOutRequest func(*StepOutRequest, *StepOutResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetBreakpointsRequest`) to further populate the given `SetBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetBreakpointsRequest func(*SetBreakpointsRequest, *SetBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `EvaluateRequest`) to further populate the given `EvaluateResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnEvaluateRequest func(*EvaluateRequest, *EvaluateResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ScopesRequest`) to further populate the given `ScopesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnScopesRequest func(*ScopesRequest, *ScopesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ReadMemoryRequest`) to further populate the given `ReadMemoryResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnReadMemoryRequest func(*ReadMemoryRequest, *ReadMemoryResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `LoadedSourcesRequest`) to further populate the given `LoadedSourcesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnLoadedSourcesRequest func(*LoadedSourcesRequest, *LoadedSourcesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ContinueRequest`) to further populate the given `ContinueResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnContinueRequest func(*ContinueRequest, *ContinueResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `LaunchRequest`) to further populate the given `LaunchResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnLaunchRequest func(*LaunchRequest, *LaunchResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetExceptionBreakpointsRequest`) to further populate the given `SetExceptionBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetExceptionBreakpointsRequest func(*SetExceptionBreakpointsRequest, *SetExceptionBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepInRequest`) to further populate the given `StepInResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepInRequest func(*StepInRequest, *StepInResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `TerminateRequest`) to further populate the given `TerminateResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnTerminateRequest func(*TerminateRequest, *TerminateResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `DisconnectRequest`) to further populate the given `DisconnectResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnDisconnectRequest func(*DisconnectRequest, *DisconnectResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetExpressionRequest`) to further populate the given `SetExpressionResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetExpressionRequest func(*SetExpressionRequest, *SetExpressionResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepBackRequest`) to further populate the given `StepBackResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepBackRequest func(*StepBackRequest, *StepBackResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StackTraceRequest`) to further populate the given `StackTraceResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStackTraceRequest func(*StackTraceRequest, *StackTraceResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `InitializeRequest`) to further populate the given `InitializeResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnInitializeRequest func(*InitializeRequest, *InitializeResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `PauseRequest`) to further populate the given `PauseResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnPauseRequest func(*PauseRequest, *PauseResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ModulesRequest`) to further populate the given `ModulesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnModulesRequest func(*ModulesRequest, *ModulesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `VariablesRequest`) to further populate the given `VariablesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnVariablesRequest func(*VariablesRequest, *VariablesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `CompletionsRequest`) to further populate the given `CompletionsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnCompletionsRequest func(*CompletionsRequest, *CompletionsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetDataBreakpointsRequest`) to further populate the given `SetDataBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetDataBreakpointsRequest func(*SetDataBreakpointsRequest, *SetDataBreakpointsResponse)error

// If a type-switch on `inRequest` succeeds, `outResponse` points to a `Response`-based `struct` value containing the `Response` initialized by the specified `initNewResponse` and further populated by the `OnFooRequest` handler corresponding to the concrete type of `inRequest` (if any). The only `err` returned, if any, is that returned by the specialized `OnFooRequest` handler.
func HandleRequest(inRequest interface{}, initNewResponse func(*Request, *Response)) (outResponse interface{}, baseResponse *Response, handled bool, err error) {
	switch input := inRequest.(type) {
	case *ModulesRequest:
		o := NewModulesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnModulesRequest!=nil; handled { err = OnModulesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *VariablesRequest:
		o := NewVariablesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnVariablesRequest!=nil; handled { err = OnVariablesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *CompletionsRequest:
		o := NewCompletionsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnCompletionsRequest!=nil; handled { err = OnCompletionsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetDataBreakpointsRequest:
		o := NewSetDataBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetDataBreakpointsRequest!=nil; handled { err = OnSetDataBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetVariableRequest:
		o := NewSetVariableResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetVariableRequest!=nil; handled { err = OnSetVariableRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RestartFrameRequest:
		o := NewRestartFrameResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnRestartFrameRequest!=nil; handled { err = OnRestartFrameRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ExceptionInfoRequest:
		o := NewExceptionInfoResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnExceptionInfoRequest!=nil; handled { err = OnExceptionInfoRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RunInTerminalRequest:
		o := NewRunInTerminalResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnRunInTerminalRequest!=nil; handled { err = OnRunInTerminalRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SourceRequest:
		o := NewSourceResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSourceRequest!=nil; handled { err = OnSourceRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *GotoRequest:
		o := NewGotoResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnGotoRequest!=nil; handled { err = OnGotoRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *TerminateThreadsRequest:
		o := NewTerminateThreadsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnTerminateThreadsRequest!=nil; handled { err = OnTerminateThreadsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetFunctionBreakpointsRequest:
		o := NewSetFunctionBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetFunctionBreakpointsRequest!=nil; handled { err = OnSetFunctionBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *DisassembleRequest:
		o := NewDisassembleResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnDisassembleRequest!=nil; handled { err = OnDisassembleRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *GotoTargetsRequest:
		o := NewGotoTargetsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnGotoTargetsRequest!=nil; handled { err = OnGotoTargetsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *AttachRequest:
		o := NewAttachResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnAttachRequest!=nil; handled { err = OnAttachRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ConfigurationDoneRequest:
		o := NewConfigurationDoneResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnConfigurationDoneRequest!=nil; handled { err = OnConfigurationDoneRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ReverseContinueRequest:
		o := NewReverseContinueResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnReverseContinueRequest!=nil; handled { err = OnReverseContinueRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *NextRequest:
		o := NewNextResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnNextRequest!=nil; handled { err = OnNextRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RestartRequest:
		o := NewRestartResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnRestartRequest!=nil; handled { err = OnRestartRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ThreadsRequest:
		o := NewThreadsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnThreadsRequest!=nil; handled { err = OnThreadsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *DataBreakpointInfoRequest:
		o := NewDataBreakpointInfoResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnDataBreakpointInfoRequest!=nil; handled { err = OnDataBreakpointInfoRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ContinueRequest:
		o := NewContinueResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnContinueRequest!=nil; handled { err = OnContinueRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepInTargetsRequest:
		o := NewStepInTargetsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnStepInTargetsRequest!=nil; handled { err = OnStepInTargetsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepOutRequest:
		o := NewStepOutResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnStepOutRequest!=nil; handled { err = OnStepOutRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetBreakpointsRequest:
		o := NewSetBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetBreakpointsRequest!=nil; handled { err = OnSetBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *EvaluateRequest:
		o := NewEvaluateResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnEvaluateRequest!=nil; handled { err = OnEvaluateRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ScopesRequest:
		o := NewScopesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnScopesRequest!=nil; handled { err = OnScopesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ReadMemoryRequest:
		o := NewReadMemoryResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnReadMemoryRequest!=nil; handled { err = OnReadMemoryRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *LoadedSourcesRequest:
		o := NewLoadedSourcesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnLoadedSourcesRequest!=nil; handled { err = OnLoadedSourcesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *LaunchRequest:
		o := NewLaunchResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnLaunchRequest!=nil; handled { err = OnLaunchRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *InitializeRequest:
		o := NewInitializeResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnInitializeRequest!=nil; handled { err = OnInitializeRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetExceptionBreakpointsRequest:
		o := NewSetExceptionBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetExceptionBreakpointsRequest!=nil; handled { err = OnSetExceptionBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepInRequest:
		o := NewStepInResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnStepInRequest!=nil; handled { err = OnStepInRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *TerminateRequest:
		o := NewTerminateResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnTerminateRequest!=nil; handled { err = OnTerminateRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *DisconnectRequest:
		o := NewDisconnectResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnDisconnectRequest!=nil; handled { err = OnDisconnectRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetExpressionRequest:
		o := NewSetExpressionResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnSetExpressionRequest!=nil; handled { err = OnSetExpressionRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepBackRequest:
		o := NewStepBackResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnStepBackRequest!=nil; handled { err = OnStepBackRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StackTraceRequest:
		o := NewStackTraceResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnStackTraceRequest!=nil; handled { err = OnStackTraceRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *PauseRequest:
		o := NewPauseResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if handled = OnPauseRequest!=nil; handled { err = OnPauseRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	}
	return
}
