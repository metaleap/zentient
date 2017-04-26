// VS Code Debug Protocol
// 
// A json schema for the VS Code Debug Protocol
// 
// Package codegen'd from github.com/metaleap/zentient/zdbg-vsc/_notes_misc_etc/vscdbgprotocol.json with github.com/metaleap/zentient/zdbg-vsc-proto-gen
package zdbgvscp
import "encoding/json"
import "errors"
import "strings"



// This enumeration defines all possible conditions when a thrown exception should result in a break.
// never: never breaks,
// always: always breaks,
// unhandled: breaks when excpetion unhandled,
// userUnhandled: breaks if the exception is not handled by user code.
// 
// POSSIBLE VALUES: `never`, `always`, `unhandled`, `userUnhandled`
type ExceptionBreakMode string



// Response to 'scopes' request.
type ScopesResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The scopes of the stackframe. If the array has length zero, there are no scopes available.
		Scopes []Scope `json:"scopes"`

	} `json:"body"`

} // struct ScopesResponse
func (me *ScopesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'setExceptionBreakpoints' request.
type SetExceptionBreakpointsArguments struct {

	// Configuration options for selected exceptions.
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`

	// IDs of checked exception options. The set of IDs is returned via the 'exceptionBreakpointFilters' capability.
	Filters []string `json:"filters"`

} // struct SetExceptionBreakpointsArguments
func (me *SetExceptionBreakpointsArguments) propagateFieldsToBase() {
}



// On error that is whenever 'success' is false, the body can provide more details.
type ErrorResponse struct {
	// Response to a request.
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



// Arguments for 'disconnect' request.
type DisconnectArguments struct {

	// Indicates whether the debuggee should be terminated when the debugger is disconnected.
	// If unspecified, the debug adapter is free to do whatever it thinks is best.
	// A client can only rely on this attribute being properly honored if a debug adapter returns true for the 'supportTerminateDebuggee' capability.
	TerminateDebuggee bool `json:"terminateDebuggee,omitempty"`

	// restart
	Restart bool `json:"restart,omitempty"`

} // struct DisconnectArguments
func (me *DisconnectArguments) propagateFieldsToBase() {
}



// Arguments for 'stepBack' request.
type StepBackArguments struct {

	// Exceute 'stepBack' for this thread.
	ThreadId int `json:"threadId"`

} // struct StepBackArguments
func (me *StepBackArguments) propagateFieldsToBase() {
}



// A client or server-initiated request.
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



// Event message for 'exited' event type.
// The event indicates that the debuggee has exited.
type ExitedEvent struct {
	// Server-initiated event.
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



// Response to 'completions' request.
type CompletionsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible completions for .
		Targets []CompletionItem `json:"targets"`

	} `json:"body"`

} // struct CompletionsResponse
func (me *CompletionsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
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



// Response to 'stepInTargets' request.
type StepInTargetsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible stepIn targets of the specified source location.
		Targets []StepInTarget `json:"targets"`

	} `json:"body"`

} // struct StepInTargetsResponse
func (me *StepInTargetsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'setFunctionBreakpoints' request.
type SetFunctionBreakpointsArguments struct {

	// The function names of the breakpoints.
	Breakpoints []FunctionBreakpoint `json:"breakpoints"`

} // struct SetFunctionBreakpointsArguments
func (me *SetFunctionBreakpointsArguments) propagateFieldsToBase() {
}



// Event message for 'stopped' event type.
// The event indicates that the execution of the debuggee has stopped due to some condition.
// This can be caused by a break point previously set, a stepping action has completed, by executing a debugger statement etc.
type StoppedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `stopped`
	Event_ string `json:"event"`

	Body struct {

		// Additional information. E.g. if reason is 'exception', text contains the exception name. This string is shown in the UI.
		Text string `json:"text,omitempty"`

		// If allThreadsStopped is true, a debug adapter can announce that all threads have stopped.
		// *  The client should use this information to enable that all threads can be expanded to access their stacktraces.
		// *  If the attribute is missing or false, only the thread with the given threadId can be expanded.
		AllThreadsStopped bool `json:"allThreadsStopped,omitempty"`

		// The reason for the event (such as: 'step', 'breakpoint', 'exception', 'pause', 'entry').
		// For backward compatibility this string is shown in the UI if the 'description' attribute is missing (but it must not be translated).
		// 
		// POSSIBLE VALUES: `step`, `breakpoint`, `exception`, `pause`, `entry`
		Reason string `json:"reason"`

		// The full reason for the event, e.g. 'Paused on exception'. This string is shown in the UI as is.
		Description string `json:"description,omitempty"`

		// The thread which was stopped.
		ThreadId int `json:"threadId,omitempty"`

	} `json:"body"`

} // struct StoppedEvent
func (me *StoppedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Information about the capabilities of a debug adapter.
type Capabilities struct {

	// The debug adapter supports breakpoints that break execution after a specified number of hits.
	SupportsHitConditionalBreakpoints bool `json:"supportsHitConditionalBreakpoints,omitempty"`

	// The debug adapter supports the stepInTargetsRequest.
	SupportsStepInTargetsRequest bool `json:"supportsStepInTargetsRequest,omitempty"`

	// The debug adapter supports 'exceptionOptions' on the setExceptionBreakpoints request.
	SupportsExceptionOptions bool `json:"supportsExceptionOptions,omitempty"`

	// The debug adapter supports a 'format' attribute on the stackTraceRequest, variablesRequest, and evaluateRequest.
	SupportsValueFormattingOptions bool `json:"supportsValueFormattingOptions,omitempty"`

	// The debug adapter supports the 'terminateDebuggee' attribute on the 'disconnect' request.
	SupportTerminateDebuggee bool `json:"supportTerminateDebuggee,omitempty"`

	// The debug adapter supports the configurationDoneRequest.
	SupportsConfigurationDoneRequest bool `json:"supportsConfigurationDoneRequest,omitempty"`

	// The debug adapter supports function breakpoints.
	SupportsFunctionBreakpoints bool `json:"supportsFunctionBreakpoints,omitempty"`

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

	// The debug adapter supports the gotoTargetsRequest.
	SupportsGotoTargetsRequest bool `json:"supportsGotoTargetsRequest,omitempty"`

	// Checksum algorithms supported by the debug adapter.
	SupportedChecksumAlgorithms []ChecksumAlgorithm `json:"supportedChecksumAlgorithms,omitempty"`

	// The debug adapter supports a (side effect free) evaluate request for data hovers.
	SupportsEvaluateForHovers bool `json:"supportsEvaluateForHovers,omitempty"`

	// The debug adapter supports setting a variable to a value.
	SupportsSetVariable bool `json:"supportsSetVariable,omitempty"`

	// The debug adapter supports restarting a frame.
	SupportsRestartFrame bool `json:"supportsRestartFrame,omitempty"`

	// The debug adapter supports the completionsRequest.
	SupportsCompletionsRequest bool `json:"supportsCompletionsRequest,omitempty"`

	// The debug adapter supports the modules request.
	SupportsModulesRequest bool `json:"supportsModulesRequest,omitempty"`

	// The debug adapter supports the RestartRequest. In this case a client should not implement 'restart' by terminating and relaunching the adapter but by calling the RestartRequest.
	SupportsRestartRequest bool `json:"supportsRestartRequest,omitempty"`

	// The debug adapter supports the exceptionInfo request.
	SupportsExceptionInfoRequest bool `json:"supportsExceptionInfoRequest,omitempty"`

	// The debug adapter supports conditional breakpoints.
	SupportsConditionalBreakpoints bool `json:"supportsConditionalBreakpoints,omitempty"`

	// The debug adapter supports stepping back via the stepBack and reverseContinue requests.
	SupportsStepBack bool `json:"supportsStepBack,omitempty"`

	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []ColumnDescriptor `json:"additionalModuleColumns,omitempty"`

} // struct Capabilities
func (me *Capabilities) propagateFieldsToBase() {
}



// Arguments for 'initialize' request.
type InitializeRequestArguments struct {

	// The ID of the (frontend) client using this adapter.
	ClientID string `json:"clientID,omitempty"`

	// The ID of the debug adapter.
	AdapterID string `json:"adapterID"`

	// If true all line numbers are 1-based (default).
	LinesStartAt1 bool `json:"linesStartAt1,omitempty"`

	// If true all column numbers are 1-based (default).
	ColumnsStartAt1 bool `json:"columnsStartAt1,omitempty"`

	// Determines in what format paths are specified. Possible values are 'path' or 'uri'. The default is 'path', which is the native format.
	// 
	// POSSIBLE VALUES: `path`, `uri`
	PathFormat string `json:"pathFormat,omitempty"`

	// Client supports the optional type attribute for variables.
	SupportsVariableType bool `json:"supportsVariableType,omitempty"`

	// Client supports the paging of variables.
	SupportsVariablePaging bool `json:"supportsVariablePaging,omitempty"`

	// Client supports the runInTerminal request.
	SupportsRunInTerminalRequest bool `json:"supportsRunInTerminalRequest,omitempty"`

} // struct InitializeRequestArguments
func (me *InitializeRequestArguments) propagateFieldsToBase() {
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



// Event message for 'module' event type.
// The event indicates that some information about a module has changed.
type ModuleEvent struct {
	// Server-initiated event.
	Event

	Body struct {

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `new`, `changed`, `removed`
		Reason string `json:"reason"`

		// The new, changed, or removed module. In case of 'removed' only the module id is used.
		Module Module `json:"module"`

	} `json:"body"`

	// POSSIBLE VALUES: `module`
	Event_ string `json:"event"`

} // struct ModuleEvent
func (me *ModuleEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// StepInTargets request; value of command field is 'stepInTargets'.
// This request retrieves the possible stepIn targets for the specified stack frame.
// These targets can be used in the 'stepIn' request.
// The StepInTargets may only be called if the 'supportsStepInTargetsRequest' capability exists and is true.
type StepInTargetsRequest struct {
	// A client or server-initiated request.
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



// Server-initiated event.
type Event struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

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

	// POSSIBLE VALUES: `event`
	Type string `json:"type"`

} // struct Event
func (me *Event) propagateFieldsToBase() {
	me.ProtocolMessage.Type = me.Type
	me.ProtocolMessage.propagateFieldsToBase()
}



// Event message for 'breakpoint' event type.
// The event indicates that some information about a breakpoint has changed.
type BreakpointEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `breakpoint`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event (such as: 'changed', 'new').
		// 
		// POSSIBLE VALUES: `changed`, `new`
		Reason string `json:"reason"`

		// The breakpoint.
		Breakpoint Breakpoint `json:"breakpoint"`

	} `json:"body"`

} // struct BreakpointEvent
func (me *BreakpointEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// SetBreakpoints request; value of command field is 'setBreakpoints'.
// Sets multiple breakpoints for a single source and clears all previous breakpoints in that source.
// To clear all breakpoint for a source, specify an empty array.
// When a breakpoint is hit, a StoppedEvent (event type 'breakpoint') is generated.
type SetBreakpointsRequest struct {
	// A client or server-initiated request.
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



// Launch request; value of command field is 'launch'.
type LaunchRequest struct {
	// A client or server-initiated request.
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



// Arguments for 'exceptionInfo' request.
type ExceptionInfoArguments struct {

	// Thread for which exception information should be retrieved.
	ThreadId int `json:"threadId"`

} // struct ExceptionInfoArguments
func (me *ExceptionInfoArguments) propagateFieldsToBase() {
}



// Variables request; value of command field is 'variables'.
// Retrieves all child variables for the given variable reference.
// An optional filter can be used to limit the fetched children to either named or indexed children.
type VariablesRequest struct {
	// A client or server-initiated request.
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



// Response to 'modules' request.
type ModulesResponse struct {
	// Response to a request.
	Response

	Body struct {

		// All modules or range of modules.
		Modules []Module `json:"modules"`

		// The total number of modules available.
		TotalModules int `json:"totalModules,omitempty"`

	} `json:"body"`

} // struct ModulesResponse
func (me *ModulesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Response to 'stepBack' request. This is just an acknowledgement, so no body field is required.
type StepBackResponse struct {
	// Response to a request.
	Response

} // struct StepBackResponse
func (me *StepBackResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Properties of a breakpoint passed to the setFunctionBreakpoints request.
type FunctionBreakpoint struct {

	// The name of the function.
	Name string `json:"name"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

} // struct FunctionBreakpoint
func (me *FunctionBreakpoint) propagateFieldsToBase() {
}



// Arguments for 'runInTerminal' request.
type RunInTerminalRequestArguments struct {

	// Optional title of the terminal.
	Title string `json:"title,omitempty"`

	// Working directory of the command.
	Cwd string `json:"cwd"`

	// List of arguments. The first argument is the command to run.
	Args []string `json:"args"`

	// Environment key-value pairs that are added to the default environment.
	Env map[string]string `json:"env,omitempty"`

	// What kind of terminal to launch.
	// 
	// POSSIBLE VALUES: `integrated`, `external`
	Kind string `json:"kind,omitempty"`

} // struct RunInTerminalRequestArguments
func (me *RunInTerminalRequestArguments) propagateFieldsToBase() {
}



// Arguments for 'setBreakpoints' request.
type SetBreakpointsArguments struct {

	// The source location of the breakpoints; either source.path or source.reference must be specified.
	Source Source `json:"source"`

	// The code locations of the breakpoints.
	Breakpoints []SourceBreakpoint `json:"breakpoints,omitempty"`

	// Deprecated: The code locations of the breakpoints.
	Lines []int `json:"lines,omitempty"`

	// A value of true indicates that the underlying source has been modified which results in new breakpoint locations.
	SourceModified bool `json:"sourceModified,omitempty"`

} // struct SetBreakpointsArguments
func (me *SetBreakpointsArguments) propagateFieldsToBase() {
}



// SetFunctionBreakpoints request; value of command field is 'setFunctionBreakpoints'.
// Sets multiple function breakpoints and clears all previous function breakpoints.
// To clear all function breakpoint, specify an empty array.
// When a function breakpoint is hit, a StoppedEvent (event type 'function breakpoint') is generated.
type SetFunctionBreakpointsRequest struct {
	// A client or server-initiated request.
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



// Response to Initialize request.
type RunInTerminalResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The process ID.
		ProcessId int64 `json:"processId,omitempty"`

	} `json:"body"`

} // struct RunInTerminalResponse
func (me *RunInTerminalResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Response to 'launch' request. This is just an acknowledgement, so no body field is required.
type LaunchResponse struct {
	// Response to a request.
	Response

} // struct LaunchResponse
func (me *LaunchResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Response to 'initialize' request.
type InitializeResponse struct {
	// Response to a request.
	Response

	// The capabilities of this debug adapter.
	Body Capabilities `json:"body,omitempty"`

} // struct InitializeResponse
func (me *InitializeResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Response to 'gotoTargets' request.
type GotoTargetsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible goto targets of the specified location.
		Targets []GotoTarget `json:"targets"`

	} `json:"body"`

} // struct GotoTargetsResponse
func (me *GotoTargetsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Attach request; value of command field is 'attach'.
type AttachRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `attach`
	Command string `json:"command"`

	Arguments AttachRequestArguments `json:"arguments"`

} // struct AttachRequest
func (me *AttachRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Properties of a breakpoint passed to the setBreakpoints request.
type SourceBreakpoint struct {

	// An optional source column of the breakpoint.
	Column int `json:"column,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`

	// The source line of the breakpoint.
	Line int `json:"line"`

} // struct SourceBreakpoint
func (me *SourceBreakpoint) propagateFieldsToBase() {
}



// Restart request; value of command field is 'restart'.
// Restarts a debug session. If the capability 'supportsRestartRequest' is missing or has the value false,
// the client will implement 'restart' by terminating the debug adapter first and then launching it anew.
// A debug adapter can override this default behaviour by implementing a restart request
// and setting the capability 'supportsRestartRequest' to true.
type RestartRequest struct {
	// A client or server-initiated request.
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



// Continue request; value of command field is 'continue'.
// The request starts the debuggee to run again.
type ContinueRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `continue`
	Command string `json:"command"`

	Arguments ContinueArguments `json:"arguments"`

} // struct ContinueRequest
func (me *ContinueRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'goto' request. This is just an acknowledgement, so no body field is required.
type GotoResponse struct {
	// Response to a request.
	Response

} // struct GotoResponse
func (me *GotoResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// A GotoTarget describes a code location that can be used as a target in the 'goto' request.
// The possible goto targets can be determined via the 'gotoTargets' request.
type GotoTarget struct {

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

	// An optional end column of the range covered by the goto target.
	EndColumn int `json:"endColumn,omitempty"`

} // struct GotoTarget
func (me *GotoTarget) propagateFieldsToBase() {
}



// Pause request; value of command field is 'pause'.
// The request suspenses the debuggee.
// The debug adapter first sends the PauseResponse and then a StoppedEvent (event type 'pause') after the thread has been paused successfully.
type PauseRequest struct {
	// A client or server-initiated request.
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



// CompletionsRequest request; value of command field is 'completions'.
// Returns a list of possible completions for a given caret position and text.
// The CompletionsRequest may only be called if the 'supportsCompletionsRequest' capability exists and is true.
type CompletionsRequest struct {
	// A client or server-initiated request.
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



// Response to 'stepIn' request. This is just an acknowledgement, so no body field is required.
type StepInResponse struct {
	// Response to a request.
	Response

} // struct StepInResponse
func (me *StepInResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// StackTrace request; value of command field is 'stackTrace'. The request returns a stacktrace from the current execution state.
type StackTraceRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command"`

	Arguments StackTraceArguments `json:"arguments"`

} // struct StackTraceRequest
func (me *StackTraceRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// A Source is a descriptor for source code. It is returned from the debug adapter as part of a StackFrame and it is used by clients when specifying breakpoints.
type Source struct {

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

	// The path of the source to be shown in the UI. It is only used to locate and load the content of the source if no sourceReference is specified (or its vaule is 0).
	Path string `json:"path,omitempty"`

	// If sourceReference > 0 the contents of the source must be retrieved through the SourceRequest (even if a path is specified). A sourceReference is only valid for a session, so it must not be used to persist a source.
	SourceReference int64 `json:"sourceReference,omitempty"`

	// An optional hint for how to present the source in the UI. A value of 'deemphasize' can be used to indicate that the source is not available or that it is skipped on stepping.
	// 
	// POSSIBLE VALUES: `emphasize`, `deemphasize`
	PresentationHint string `json:"presentationHint,omitempty"`

	// The (optional) origin of this source: possible values 'internal module', 'inlined content from source map', etc.
	Origin string `json:"origin,omitempty"`

} // struct Source
func (me *Source) propagateFieldsToBase() {
}



// Response to 'continue' request.
type ContinueResponse struct {
	// Response to a request.
	Response

	Body struct {

		// If true, the continue request has ignored the specified thread and continued all threads instead. If this attribute is missing a value of 'true' is assumed for backward compatibility.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

	} `json:"body"`

} // struct ContinueResponse
func (me *ContinueResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'stepInTargets' request.
type StepInTargetsArguments struct {

	// The stack frame for which to retrieve the possible stepIn targets.
	FrameId int `json:"frameId"`

} // struct StepInTargetsArguments
func (me *StepInTargetsArguments) propagateFieldsToBase() {
}



// A structured message object. Used to return errors from requests.
type Message struct {

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

	// If true send to telemetry.
	SendTelemetry bool `json:"sendTelemetry,omitempty"`

} // struct Message
func (me *Message) propagateFieldsToBase() {
}



// Response to 'reverseContinue' request. This is just an acknowledgement, so no body field is required.
type ReverseContinueResponse struct {
	// Response to a request.
	Response

} // struct ReverseContinueResponse
func (me *ReverseContinueResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Modules can be retrieved from the debug adapter with the ModulesRequest which can either return all modules or a range of modules to support paging.
type ModulesRequest struct {
	// A client or server-initiated request.
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



// ExceptionInfoRequest request; value of command field is 'exceptionInfo'.
// Retrieves the details of the exception that caused the StoppedEvent to be raised.
type ExceptionInfoRequest struct {
	// A client or server-initiated request.
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



// Arguments for 'next' request.
type NextArguments struct {

	// Execute 'next' for this thread.
	ThreadId int `json:"threadId"`

} // struct NextArguments
func (me *NextArguments) propagateFieldsToBase() {
}



// Response to 'next' request. This is just an acknowledgement, so no body field is required.
type NextResponse struct {
	// Response to a request.
	Response

} // struct NextResponse
func (me *NextResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// StepIn request; value of command field is 'stepIn'.
// The request starts the debuggee to step into a function/method if possible.
// If it cannot step into a target, 'stepIn' behaves like 'next'.
// The debug adapter first sends the StepInResponse and then a StoppedEvent (event type 'step') after the step has completed.
// If there are multiple function/method calls (or other targets) on the source line,
// the optional argument 'targetId' can be used to control into which target the 'stepIn' should occur.
// The list of possible targets for a given source line can be retrieved via the 'stepInTargets' request.
type StepInRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command"`

	Arguments StepInArguments `json:"arguments"`

} // struct StepInRequest
func (me *StepInRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Arguments for 'pause' request.
type PauseArguments struct {

	// Pause execution for this thread.
	ThreadId int `json:"threadId"`

} // struct PauseArguments
func (me *PauseArguments) propagateFieldsToBase() {
}



// ConfigurationDone request; value of command field is 'configurationDone'.
// The client of the debug protocol must send this request at the end of the sequence of configuration requests (which was started by the InitializedEvent).
type ConfigurationDoneRequest struct {
	// A client or server-initiated request.
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



// Response to 'setExceptionBreakpoints' request. This is just an acknowledgement, so no body field is required.
type SetExceptionBreakpointsResponse struct {
	// Response to a request.
	Response

} // struct SetExceptionBreakpointsResponse
func (me *SetExceptionBreakpointsResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
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

	// A name of the module.
	Name string `json:"name"`

	// True if the module is optimized.
	IsOptimized bool `json:"isOptimized,omitempty"`

	// Version of Module.
	Version string `json:"version,omitempty"`

	// User understandable description of if symbols were found for the module (ex: 'Symbols Loaded', 'Symbols not found', etc.
	SymbolStatus string `json:"symbolStatus,omitempty"`

	// Unique identifier for the module.
	// 
	// POSSIBLE TYPES:
	// - `int` (for JSON `integer`s)
	// - `string` (for JSON `string`s)
	Id interface{} `json:"id"`

	// optional but recommended attributes.
	// always try to use these first before introducing additional attributes.
	// 
	// Logical full path to the module. The exact definition is implementation defined, but usually this would be a full path to the on-disk file for the module.
	Path string `json:"path,omitempty"`

	// True if the module is considered 'user code' by a debugger that supports 'Just My Code'.
	IsUserCode bool `json:"isUserCode,omitempty"`

	// Logical full path to the symbol file. The exact definition is implementation defined.
	SymbolFilePath string `json:"symbolFilePath,omitempty"`

	// Module created or modified.
	DateTimeStamp string `json:"dateTimeStamp,omitempty"`

	// Address range covered by this module.
	AddressRange string `json:"addressRange,omitempty"`

} // struct Module
func (me *Module) propagateFieldsToBase() {
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



// Initialize request; value of command field is 'initialize'.
type InitializeRequest struct {
	// A client or server-initiated request.
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



// The ModulesViewDescriptor is the container for all declarative configuration options of a ModuleView.
// For now it only specifies the columns to be shown in the modules view.
type ModulesViewDescriptor struct {

	Columns []ColumnDescriptor `json:"columns"`

} // struct ModulesViewDescriptor
func (me *ModulesViewDescriptor) propagateFieldsToBase() {
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



// Response to 'restart' request. This is just an acknowledgement, so no body field is required.
type RestartResponse struct {
	// Response to a request.
	Response

} // struct RestartResponse
func (me *RestartResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// runInTerminal request; value of command field is 'runInTerminal'.
// With this request a debug adapter can run a command in a terminal.
type RunInTerminalRequest struct {
	// A client or server-initiated request.
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



// Base class of requests, responses, and events.
type ProtocolMessage struct {

	// Sequence number.
	Seq int `json:"seq"`

	// One of 'request', 'response', or 'event'.
	// 
	// POSSIBLE VALUES: `request`, `response`, `event`
	Type string `json:"type"`

} // struct ProtocolMessage
func (me *ProtocolMessage) propagateFieldsToBase() {
}



// Arguments for 'scopes' request.
type ScopesArguments struct {

	// Retrieve the scopes for this stackframe.
	FrameId int `json:"frameId"`

} // struct ScopesArguments
func (me *ScopesArguments) propagateFieldsToBase() {
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



// Event message for 'continued' event type.
// The event indicates that the execution of the debuggee has continued.
// Please note: a debug adapter is not expected to send this event in response to a request that implies that execution continues, e.g. 'launch' or 'continue'.
// It is only necessary to send a ContinuedEvent if there was no previous request that implied this.
type ContinuedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `continued`
	Event_ string `json:"event"`

	Body struct {

		// If allThreadsContinued is true, a debug adapter can announce that all threads have continued.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

		// The thread which was continued.
		ThreadId int `json:"threadId"`

	} `json:"body"`

} // struct ContinuedEvent
func (me *ContinuedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Response to 'pause' request. This is just an acknowledgement, so no body field is required.
type PauseResponse struct {
	// Response to a request.
	Response

} // struct PauseResponse
func (me *PauseResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Some predefined types for the CompletionItem. Please note that not all clients have specific icons for all of them.
// 
// POSSIBLE VALUES: `method`, `function`, `constructor`, `field`, `variable`, `class`, `interface`, `module`, `property`, `unit`, `value`, `enum`, `keyword`, `snippet`, `text`, `color`, `file`, `reference`, `customcolor`
type CompletionItemType string



// Provides formatting information for a stack frame.
type StackFrameFormat struct {
	// Provides formatting information for a value.
	ValueFormat

	// Displays the types of parameters for the stack frame.
	ParameterTypes bool `json:"parameterTypes,omitempty"`

	// Displays the names of parameters for the stack frame.
	ParameterNames bool `json:"parameterNames,omitempty"`

	// Displays the values of parameters for the stack frame.
	ParameterValues bool `json:"parameterValues,omitempty"`

	// Displays the line number of the stack frame.
	Line bool `json:"line,omitempty"`

	// Displays the module of the stack frame.
	Module bool `json:"module,omitempty"`

	// Displays parameters for the stack frame.
	Parameters bool `json:"parameters,omitempty"`

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



// Arguments for 'stepIn' request.
type StepInArguments struct {

	// Execute 'stepIn' for this thread.
	ThreadId int `json:"threadId"`

	// Optional id of the target to step into.
	TargetId int `json:"targetId,omitempty"`

} // struct StepInArguments
func (me *StepInArguments) propagateFieldsToBase() {
}



// Arguments for 'restart' request.
// The restart request has no standardized attributes.
type RestartArguments struct {

} // struct RestartArguments
func (me *RestartArguments) propagateFieldsToBase() {
}



// Event message for 'initialized' event type.
// This event indicates that the debug adapter is ready to accept configuration requests (e.g. SetBreakpointsRequest, SetExceptionBreakpointsRequest).
// A debug adapter is expected to send this event when it is ready to accept configuration requests (but not before the InitializeRequest has finished).
// The sequence of events/requests is as follows:
// - adapters sends InitializedEvent (after the InitializeRequest has returned)
// - frontend sends zero or more SetBreakpointsRequest
// - frontend sends one SetFunctionBreakpointsRequest
// - frontend sends a SetExceptionBreakpointsRequest if one or more exceptionBreakpointFilters have been defined (or if supportsConfigurationDoneRequest is not defined or false)
// - frontend sends other future configuration requests
// - frontend sends one ConfigurationDoneRequest to indicate the end of the configuration
type InitializedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `initialized`
	Event_ string `json:"event"`

} // struct InitializedEvent
func (me *InitializedEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.propagateFieldsToBase()
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



// Response to a request.
type Response struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

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

	// POSSIBLE VALUES: `response`
	Type string `json:"type"`

	// Sequence number of the corresponding request.
	Request_seq int `json:"request_seq"`

	// Outcome of the request.
	Success bool `json:"success"`

} // struct Response
func (me *Response) propagateFieldsToBase() {
	me.ProtocolMessage.Type = me.Type
	me.ProtocolMessage.propagateFieldsToBase()
}



// Arguments for 'stackTrace' request.
type StackTraceArguments struct {

	// The maximum number of frames to return. If levels is not specified or 0, all frames are returned.
	Levels int `json:"levels,omitempty"`

	// Specifies details on how to format the stack frames.
	Format StackFrameFormat `json:"format,omitempty"`

	// Retrieve the stacktrace for this thread.
	ThreadId int `json:"threadId"`

	// The index of the first frame to return; if omitted frames start at 0.
	StartFrame int `json:"startFrame,omitempty"`

} // struct StackTraceArguments
func (me *StackTraceArguments) propagateFieldsToBase() {
}



// Source request; value of command field is 'source'.
// The request retrieves the source code for a given source reference.
type SourceRequest struct {
	// A client or server-initiated request.
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



// Response to 'threads' request.
type ThreadsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// All threads.
		Threads []Thread `json:"threads"`

	} `json:"body"`

} // struct ThreadsResponse
func (me *ThreadsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'configurationDone' request.
// The configurationDone request has no standardized attributes.
type ConfigurationDoneArguments struct {

} // struct ConfigurationDoneArguments
func (me *ConfigurationDoneArguments) propagateFieldsToBase() {
}



// Provides formatting information for a value.
type ValueFormat struct {

	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`

} // struct ValueFormat
func (me *ValueFormat) propagateFieldsToBase() {
}



// Arguments for 'launch' request.
type LaunchRequestArguments struct {

	// If noDebug is true the launch request should launch the program without enabling debugging.
	NoDebug bool `json:"noDebug,omitempty"`

	// w
	W string `json:"w,omitempty"`

	// c
	C string `json:"c,omitempty"`

	// f
	F string `json:"f,omitempty"`

} // struct LaunchRequestArguments
func (me *LaunchRequestArguments) propagateFieldsToBase() {
}



// Arguments for 'continue' request.
type ContinueArguments struct {

	// Continue execution for the specified thread (if possible). If the backend cannot continue on a single thread but will continue on all threads, it should set the allThreadsContinued attribute in the response to true.
	ThreadId int `json:"threadId"`

} // struct ContinueArguments
func (me *ContinueArguments) propagateFieldsToBase() {
}



// Arguments for 'setVariable' request.
type SetVariableArguments struct {

	// Specifies details on how to format the response value.
	Format ValueFormat `json:"format,omitempty"`

	// The reference of the variable container.
	VariablesReference int `json:"variablesReference"`

	// The name of the variable.
	Name string `json:"name"`

	// The value of the variable.
	Value string `json:"value"`

} // struct SetVariableArguments
func (me *SetVariableArguments) propagateFieldsToBase() {
}



// A Stackframe contains the source location.
type StackFrame struct {

	// An optional end column of the range covered by the stack frame.
	EndColumn int `json:"endColumn,omitempty"`

	// The module associated with this frame, if any.
	// 
	// POSSIBLE TYPES:
	// - `int` (for JSON `integer`s)
	// - `string` (for JSON `string`s)
	ModuleId interface{} `json:"moduleId,omitempty"`

	// The name of the stack frame, typically a method name.
	Name string `json:"name"`

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line int `json:"line"`

	// An optional end line of the range covered by the stack frame.
	EndLine int `json:"endLine,omitempty"`

	// An optional hint for how to present this frame in the UI. A value of 'label' can be used to indicate that the frame is an artificial frame that is used as a visual label or separator.
	// 
	// POSSIBLE VALUES: `normal`, `label`
	PresentationHint string `json:"presentationHint,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id int `json:"id"`

	// The optional source of the frame.
	Source Source `json:"source,omitempty"`

	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column int `json:"column"`

} // struct StackFrame
func (me *StackFrame) propagateFieldsToBase() {
}



// RestartFrame request; value of command field is 'restartFrame'.
// The request restarts execution of the specified stackframe.
// The debug adapter first sends the RestartFrameResponse and then a StoppedEvent (event type 'restart') after the restart has completed.
type RestartFrameRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `restartFrame`
	Command string `json:"command"`

	Arguments RestartFrameArguments `json:"arguments"`

} // struct RestartFrameRequest
func (me *RestartFrameRequest) propagateFieldsToBase() {
	me.Request.Arguments = me.Arguments
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// Response to 'evaluate' request.
type EvaluateResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

		// The result of the evaluate request.
		Result string `json:"result"`

		// The optional type of the evaluate result.
		Type string `json:"type,omitempty"`

		// If variablesReference is > 0, the evaluate result is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference"`

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

	} `json:"body"`

} // struct EvaluateResponse
func (me *EvaluateResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// A Scope is a named container for variables. Optionally a scope can map to a source or a range within a source.
type Scope struct {

	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive"`

	// Optional start line of the range covered by this scope.
	Line int `json:"line,omitempty"`

	// Optional start column of the range covered by this scope.
	Column int `json:"column,omitempty"`

	// Name of the scope such as 'Arguments', 'Locals'.
	Name string `json:"name"`

	// The number of indexed variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	IndexedVariables int `json:"indexedVariables,omitempty"`

	// Optional source for this scope.
	Source Source `json:"source,omitempty"`

	// Optional end line of the range covered by this scope.
	EndLine int `json:"endLine,omitempty"`

	// Optional end column of the range covered by this scope.
	EndColumn int `json:"endColumn,omitempty"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

	// The number of named variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	NamedVariables int `json:"namedVariables,omitempty"`

} // struct Scope
func (me *Scope) propagateFieldsToBase() {
}



// Disconnect request; value of command field is 'disconnect'.
type DisconnectRequest struct {
	// A client or server-initiated request.
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



// SetExceptionBreakpoints request; value of command field is 'setExceptionBreakpoints'.
// The request configures the debuggers response to thrown exceptions. If an exception is configured to break, a StoppedEvent is fired (event type 'exception').
type SetExceptionBreakpointsRequest struct {
	// A client or server-initiated request.
	Request

	Arguments SetExceptionBreakpointsArguments `json:"arguments"`

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command"`

} // struct SetExceptionBreakpointsRequest
func (me *SetExceptionBreakpointsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'restartFrame' request. This is just an acknowledgement, so no body field is required.
type RestartFrameResponse struct {
	// Response to a request.
	Response

} // struct RestartFrameResponse
func (me *RestartFrameResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Names of checksum algorithms that may be supported by a debug adapter.
// 
// POSSIBLE VALUES: `MD5`, `SHA1`, `SHA256`, `timestamp`
type ChecksumAlgorithm string



// Response to 'configurationDone' request. This is just an acknowledgement, so no body field is required.
type ConfigurationDoneResponse struct {
	// Response to a request.
	Response

} // struct ConfigurationDoneResponse
func (me *ConfigurationDoneResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Goto request; value of command field is 'goto'.
// The request sets the location where the debuggee will continue to run.
// This makes it possible to skip the execution of code or to executed code again.
// The code between the current location and the goto target is not executed but skipped.
// The debug adapter first sends the GotoResponse and then a StoppedEvent (event type 'goto').
type GotoRequest struct {
	// A client or server-initiated request.
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



// ReverseContinue request; value of command field is 'reverseContinue'.
// The request starts the debuggee to run backward. Clients should only call this request if the capability supportsStepBack is true.
type ReverseContinueRequest struct {
	// A client or server-initiated request.
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



// Information about a Breakpoint created in setBreakpoints or setFunctionBreakpoints.
type Breakpoint struct {

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

	// An optional end column of the actual range covered by the breakpoint. If no end line is given, then the end column is assumed to be in the start line.
	EndColumn int `json:"endColumn,omitempty"`

	// An optional unique identifier for the breakpoint.
	Id int `json:"id,omitempty"`

} // struct Breakpoint
func (me *Breakpoint) propagateFieldsToBase() {
}



// Event message for 'output' event type.
// The event indicates that the target has produced some output.
type OutputEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `output`
	Event_ string `json:"event"`

	Body struct {

		// If an attribute 'variablesReference' exists and its value is > 0, the output contains objects which can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference,omitempty"`

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

		// The category of output (such as: 'console', 'stdout', 'stderr', 'telemetry'). If not specified, 'console' is assumed.
		// 
		// POSSIBLE VALUES: `console`, `stdout`, `stderr`, `telemetry`
		Category string `json:"category,omitempty"`

		// The output to report.
		Output string `json:"output"`

	} `json:"body"`

} // struct OutputEvent
func (me *OutputEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
}



// Arguments for 'evaluate' request.
type EvaluateArguments struct {

	// The context in which the evaluate request is run. Possible values are 'watch' if evaluate is run in a watch, 'repl' if run from the REPL console, or 'hover' if run from a data hover.
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



// Response to 'setBreakpoints' request.
// Returned is information about each breakpoint created by this request.
// This includes the actual code location and whether the breakpoint could be verified.
// The breakpoints returned are in the same order as the elements of the 'breakpoints'
// (or the deprecated 'lines') in the SetBreakpointsArguments.
type SetBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements are in the same order as the elements of the 'breakpoints' (or the deprecated 'lines') in the SetBreakpointsArguments.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`

} // struct SetBreakpointsResponse
func (me *SetBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// An ExceptionBreakpointsFilter is shown in the UI as an option for configuring how exceptions are dealt with.
type ExceptionBreakpointsFilter struct {

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label"`

	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter"`

} // struct ExceptionBreakpointsFilter
func (me *ExceptionBreakpointsFilter) propagateFieldsToBase() {
}



// Arguments for 'attach' request.
// The attach request has no standardized attributes.
type AttachRequestArguments struct {

} // struct AttachRequestArguments
func (me *AttachRequestArguments) propagateFieldsToBase() {
}



// Arguments for 'restartFrame' request.
type RestartFrameArguments struct {

	// Restart this stackframe.
	FrameId int `json:"frameId"`

} // struct RestartFrameArguments
func (me *RestartFrameArguments) propagateFieldsToBase() {
}



// StepBack request; value of command field is 'stepBack'.
// The request starts the debuggee to run one step backwards.
// The debug adapter first sends the StepBackResponse and then a StoppedEvent (event type 'step') after the step has completed. Clients should only call this request if the capability supportsStepBack is true.
type StepBackRequest struct {
	// A client or server-initiated request.
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



// A Variable is a name/value pair.
// Optionally a variable can have a 'type' that is shown if space permits or when hovering over the variable's name.
// An optional 'kind' is used to render additional properties of the variable, e.g. different icons can be used to indicate that a variable is public or private.
// If the value is structured (has children), a handle is provided to retrieve the children with the VariablesRequest.
// If the number of named or indexed children is large, the numbers should be returned via the optional 'namedVariables' and 'indexedVariables' attributes.
// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
type Variable struct {

	// Properties of a variable that can be used to determine how to render the variable in the UI. Format of the string value: TBD.
	Kind string `json:"kind,omitempty"`

	// Optional evaluatable name of this variable which can be passed to the 'EvaluateRequest' to fetch the variable's value.
	EvaluateName string `json:"evaluateName,omitempty"`

	// If variablesReference is > 0, the variable is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

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

	// The type of the variable's value. Typically shown in the UI when hovering over the value.
	Type string `json:"type,omitempty"`

} // struct Variable
func (me *Variable) propagateFieldsToBase() {
}



// StepOut request; value of command field is 'stepOut'.
// The request starts the debuggee to run again for one step.
// The debug adapter first sends the StepOutResponse and then a StoppedEvent (event type 'step') after the step has completed.
type StepOutRequest struct {
	// A client or server-initiated request.
	Request

	Arguments StepOutArguments `json:"arguments"`

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command"`

} // struct StepOutRequest
func (me *StepOutRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'disconnect' request. This is just an acknowledgement, so no body field is required.
type DisconnectResponse struct {
	// Response to a request.
	Response

} // struct DisconnectResponse
func (me *DisconnectResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Arguments for 'reverseContinue' request.
type ReverseContinueArguments struct {

	// Exceute 'reverseContinue' for this thread.
	ThreadId int `json:"threadId"`

} // struct ReverseContinueArguments
func (me *ReverseContinueArguments) propagateFieldsToBase() {
}



// Scopes request; value of command field is 'scopes'.
// The request returns the variable scopes for a given stackframe ID.
type ScopesRequest struct {
	// A client or server-initiated request.
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



// Response to 'variables' request.
type VariablesResponse struct {
	// Response to a request.
	Response

	Body struct {

		// All (or a range) of variables for the given variable reference.
		Variables []Variable `json:"variables"`

	} `json:"body"`

} // struct VariablesResponse
func (me *VariablesResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Next request; value of command field is 'next'.
// The request starts the debuggee to run again for one step.
// The debug adapter first sends the NextResponse and then a StoppedEvent (event type 'step') after the step has completed.
type NextRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `next`
	Command string `json:"command"`

	Arguments NextArguments `json:"arguments"`

} // struct NextRequest
func (me *NextRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Response to 'exceptionInfo' request.
type ExceptionInfoResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Mode that caused the exception notification to be raised.
		BreakMode ExceptionBreakMode `json:"breakMode"`

		// Detailed information about the exception.
		Details ExceptionDetails `json:"details,omitempty"`

		// ID of the exception that was thrown.
		ExceptionId string `json:"exceptionId"`

		// Descriptive text for the exception provided by the debug adapter.
		Description string `json:"description,omitempty"`

	} `json:"body"`

} // struct ExceptionInfoResponse
func (me *ExceptionInfoResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Event message for 'thread' event type.
// The event indicates that a thread has started or exited.
type ThreadEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `thread`
	Event_ string `json:"event"`

	Body struct {

		// The reason for the event (such as: 'started', 'exited').
		// 
		// POSSIBLE VALUES: `started`, `exited`
		Reason string `json:"reason"`

		// The identifier of the thread.
		ThreadId int `json:"threadId"`

	} `json:"body"`

} // struct ThreadEvent
func (me *ThreadEvent) propagateFieldsToBase() {
	me.Event.Event = me.Event_
	me.Event.Body = me.Body
	me.Event.propagateFieldsToBase()
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



// Response to 'stepOut' request. This is just an acknowledgement, so no body field is required.
type StepOutResponse struct {
	// Response to a request.
	Response

} // struct StepOutResponse
func (me *StepOutResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// Arguments for 'stepOut' request.
type StepOutArguments struct {

	// Execute 'stepOut' for this thread.
	ThreadId int `json:"threadId"`

} // struct StepOutArguments
func (me *StepOutArguments) propagateFieldsToBase() {
}



// Evaluate request; value of command field is 'evaluate'.
// Evaluates the given expression in the context of the top most stack frame.
// The expression has access to any variables and arguments that are in scope.
type EvaluateRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `evaluate`
	Command string `json:"command"`

	Arguments EvaluateArguments `json:"arguments"`

} // struct EvaluateRequest
func (me *EvaluateRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.Arguments = me.Arguments
	me.Request.propagateFieldsToBase()
}



// Thread request; value of command field is 'threads'.
// The request retrieves a list of all threads.
type ThreadsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `threads`
	Command string `json:"command"`

} // struct ThreadsRequest
func (me *ThreadsRequest) propagateFieldsToBase() {
	me.Request.Command = me.Command
	me.Request.propagateFieldsToBase()
}



// setVariable request; value of command field is 'setVariable'.
// Set the variable with the given name in the variable container to a new value.
type SetVariableRequest struct {
	// A client or server-initiated request.
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



// Response to 'source' request.
type SourceResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Content of the source reference.
		Content string `json:"content"`

		// Optional content type (mime type) of the source.
		MimeType string `json:"mimeType,omitempty"`

	} `json:"body"`

} // struct SourceResponse
func (me *SourceResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// GotoTargets request; value of command field is 'gotoTargets'.
// This request retrieves the possible goto targets for the specified source location.
// These targets can be used in the 'goto' request.
// The GotoTargets request may only be called if the 'supportsGotoTargetsRequest' capability exists and is true.
type GotoTargetsRequest struct {
	// A client or server-initiated request.
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



// Response to 'setFunctionBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetFunctionBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements correspond to the elements of the 'breakpoints' array.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`

} // struct SetFunctionBreakpointsResponse
func (me *SetFunctionBreakpointsResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// Arguments for 'goto' request.
type GotoArguments struct {

	// Set the goto target for this thread.
	ThreadId int `json:"threadId"`

	// The location where the debuggee will continue to run.
	TargetId int `json:"targetId"`

} // struct GotoArguments
func (me *GotoArguments) propagateFieldsToBase() {
}



// An ExceptionOptions assigns configuration options to a set of exceptions.
type ExceptionOptions struct {

	// Condition when a thrown exception should result in a break.
	BreakMode ExceptionBreakMode `json:"breakMode"`

	// A path that selects a single or multiple exceptions in a tree. If 'path' is missing, the whole tree is selected. By convention the first segment of the path is a category that is used to group exceptions in the UI.
	Path []ExceptionPathSegment `json:"path,omitempty"`

} // struct ExceptionOptions
func (me *ExceptionOptions) propagateFieldsToBase() {
}



// Event message for 'terminated' event types.
// The event indicates that debugging of the debuggee has terminated.
type TerminatedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `terminated`
	Event_ string `json:"event"`

	Body struct {

		// A debug adapter may set 'restart' to true to request that the front end restarts the session.
		Restart bool `json:"restart,omitempty"`

	} `json:"body,omitempty"`

} // struct TerminatedEvent
func (me *TerminatedEvent) propagateFieldsToBase() {
	me.Event.Body = me.Body
	me.Event.Event = me.Event_
	me.Event.propagateFieldsToBase()
}



// Response to 'attach' request. This is just an acknowledgement, so no body field is required.
type AttachResponse struct {
	// Response to a request.
	Response

} // struct AttachResponse
func (me *AttachResponse) propagateFieldsToBase() {
	me.Response.propagateFieldsToBase()
}



// A ColumnDescriptor specifies what module attribute to show in a column of the ModulesView, how to format it, and what the column's label should be.
// It is only used if the underlying UI actually supports this level of customization.
type ColumnDescriptor struct {

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

	// Name of the attribute rendered in this column.
	AttributeName string `json:"attributeName"`

} // struct ColumnDescriptor
func (me *ColumnDescriptor) propagateFieldsToBase() {
}



// Response to 'setVariable' request.
type SetVariableResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

		// The new value of the variable.
		Value string `json:"value"`

		// The type of the new value. Typically shown in the UI when hovering over the value.
		Type string `json:"type,omitempty"`

		// If variablesReference is > 0, the new value is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference,omitempty"`

	} `json:"body"`

} // struct SetVariableResponse
func (me *SetVariableResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
	me.Response.propagateFieldsToBase()
}



// A StepInTarget can be used in the 'stepIn' request and determines into which single target the stepIn request should step.
type StepInTarget struct {

	// Unique identifier for a stepIn target.
	Id int `json:"id"`

	// The name of the stepIn target (shown in the UI).
	Label string `json:"label"`

} // struct StepInTarget
func (me *StepInTarget) propagateFieldsToBase() {
}



// Response to 'stackTrace' request.
type StackTraceResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The frames of the stackframe. If the array has length zero, there are no stackframes available.
		// This means that there is no location information available.
		StackFrames []StackFrame `json:"stackFrames"`

		// The total number of frames available.
		TotalFrames int `json:"totalFrames,omitempty"`

	} `json:"body"`

} // struct StackTraceResponse
func (me *StackTraceResponse) propagateFieldsToBase() {
	me.Response.Body = me.Body
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

func BaseEvent (someEvent interface{}) (baseEvent *Event) {
	switch me := someEvent.(type) {
	case *OutputEvent: baseEvent = &me.Event
	case *ThreadEvent: baseEvent = &me.Event
	case *TerminatedEvent: baseEvent = &me.Event
	case *ExitedEvent: baseEvent = &me.Event
	case *StoppedEvent: baseEvent = &me.Event
	case *ModuleEvent: baseEvent = &me.Event
	case *BreakpointEvent: baseEvent = &me.Event
	case *ContinuedEvent: baseEvent = &me.Event
	case *InitializedEvent: baseEvent = &me.Event
	}
	return
}

func BaseResponse (someResponse interface{}) (baseResponse *Response) {
	switch me := someResponse.(type) {
	case *ThreadsResponse: baseResponse = &me.Response
	case *EvaluateResponse: baseResponse = &me.Response
	case *ConfigurationDoneResponse: baseResponse = &me.Response
	case *RestartFrameResponse: baseResponse = &me.Response
	case *SetBreakpointsResponse: baseResponse = &me.Response
	case *DisconnectResponse: baseResponse = &me.Response
	case *VariablesResponse: baseResponse = &me.Response
	case *StepOutResponse: baseResponse = &me.Response
	case *ExceptionInfoResponse: baseResponse = &me.Response
	case *SetFunctionBreakpointsResponse: baseResponse = &me.Response
	case *AttachResponse: baseResponse = &me.Response
	case *SourceResponse: baseResponse = &me.Response
	case *SetVariableResponse: baseResponse = &me.Response
	case *StackTraceResponse: baseResponse = &me.Response
	case *ScopesResponse: baseResponse = &me.Response
	case *ErrorResponse: baseResponse = &me.Response
	case *CompletionsResponse: baseResponse = &me.Response
	case *StepInTargetsResponse: baseResponse = &me.Response
	case *StepBackResponse: baseResponse = &me.Response
	case *RunInTerminalResponse: baseResponse = &me.Response
	case *LaunchResponse: baseResponse = &me.Response
	case *ModulesResponse: baseResponse = &me.Response
	case *InitializeResponse: baseResponse = &me.Response
	case *GotoTargetsResponse: baseResponse = &me.Response
	case *GotoResponse: baseResponse = &me.Response
	case *StepInResponse: baseResponse = &me.Response
	case *ContinueResponse: baseResponse = &me.Response
	case *ReverseContinueResponse: baseResponse = &me.Response
	case *NextResponse: baseResponse = &me.Response
	case *SetExceptionBreakpointsResponse: baseResponse = &me.Response
	case *RestartResponse: baseResponse = &me.Response
	case *PauseResponse: baseResponse = &me.Response
	}
	return
}

func BaseRequest (someRequest interface{}) (baseRequest *Request) {
	switch me := someRequest.(type) {
	case *StepInTargetsRequest: baseRequest = &me.Request
	case *SetBreakpointsRequest: baseRequest = &me.Request
	case *LaunchRequest: baseRequest = &me.Request
	case *VariablesRequest: baseRequest = &me.Request
	case *SetFunctionBreakpointsRequest: baseRequest = &me.Request
	case *AttachRequest: baseRequest = &me.Request
	case *RestartRequest: baseRequest = &me.Request
	case *ContinueRequest: baseRequest = &me.Request
	case *PauseRequest: baseRequest = &me.Request
	case *CompletionsRequest: baseRequest = &me.Request
	case *StackTraceRequest: baseRequest = &me.Request
	case *ModulesRequest: baseRequest = &me.Request
	case *ExceptionInfoRequest: baseRequest = &me.Request
	case *StepInRequest: baseRequest = &me.Request
	case *ConfigurationDoneRequest: baseRequest = &me.Request
	case *InitializeRequest: baseRequest = &me.Request
	case *RunInTerminalRequest: baseRequest = &me.Request
	case *SourceRequest: baseRequest = &me.Request
	case *RestartFrameRequest: baseRequest = &me.Request
	case *DisconnectRequest: baseRequest = &me.Request
	case *SetExceptionBreakpointsRequest: baseRequest = &me.Request
	case *GotoRequest: baseRequest = &me.Request
	case *ReverseContinueRequest: baseRequest = &me.Request
	case *StepBackRequest: baseRequest = &me.Request
	case *StepOutRequest: baseRequest = &me.Request
	case *ScopesRequest: baseRequest = &me.Request
	case *NextRequest: baseRequest = &me.Request
	case *EvaluateRequest: baseRequest = &me.Request
	case *ThreadsRequest: baseRequest = &me.Request
	case *SetVariableRequest: baseRequest = &me.Request
	case *GotoTargetsRequest: baseRequest = &me.Request
	}
	return
}

// Returns a new `StepOutResponse` with the followings fields set: `type`
func NewStepOutResponse () *StepOutResponse {
	newStepOutResponse := StepOutResponse{}
	newStepOutResponse.Type = "response"
	newStepOutResponse.propagateFieldsToBase()
	return &newStepOutResponse
}

// Returns a new `StepInRequest` with the followings fields set: `command`, `type`
func NewStepInRequest () *StepInRequest {
	newStepInRequest := StepInRequest{}
	newStepInRequest.Command = "stepIn"
	newStepInRequest.Type = "request"
	newStepInRequest.propagateFieldsToBase()
	return &newStepInRequest
}

// Returns a new `SetExceptionBreakpointsRequest` with the followings fields set: `command`, `type`
func NewSetExceptionBreakpointsRequest () *SetExceptionBreakpointsRequest {
	newSetExceptionBreakpointsRequest := SetExceptionBreakpointsRequest{}
	newSetExceptionBreakpointsRequest.Command = "setExceptionBreakpoints"
	newSetExceptionBreakpointsRequest.Type = "request"
	newSetExceptionBreakpointsRequest.propagateFieldsToBase()
	return &newSetExceptionBreakpointsRequest
}

// Returns a new `StepBackRequest` with the followings fields set: `command`, `type`
func NewStepBackRequest () *StepBackRequest {
	newStepBackRequest := StepBackRequest{}
	newStepBackRequest.Command = "stepBack"
	newStepBackRequest.Type = "request"
	newStepBackRequest.propagateFieldsToBase()
	return &newStepBackRequest
}

// Returns a new `ScopesRequest` with the followings fields set: `command`, `type`
func NewScopesRequest () *ScopesRequest {
	newScopesRequest := ScopesRequest{}
	newScopesRequest.Command = "scopes"
	newScopesRequest.Type = "request"
	newScopesRequest.propagateFieldsToBase()
	return &newScopesRequest
}

// Returns a new `VariablesResponse` with the followings fields set: `type`
func NewVariablesResponse () *VariablesResponse {
	newVariablesResponse := VariablesResponse{}
	newVariablesResponse.Type = "response"
	newVariablesResponse.propagateFieldsToBase()
	return &newVariablesResponse
}

// Returns a new `ThreadsRequest` with the followings fields set: `command`, `type`
func NewThreadsRequest () *ThreadsRequest {
	newThreadsRequest := ThreadsRequest{}
	newThreadsRequest.Command = "threads"
	newThreadsRequest.Type = "request"
	newThreadsRequest.propagateFieldsToBase()
	return &newThreadsRequest
}

// Returns a new `StepInTargetsRequest` with the followings fields set: `command`, `type`
func NewStepInTargetsRequest () *StepInTargetsRequest {
	newStepInTargetsRequest := StepInTargetsRequest{}
	newStepInTargetsRequest.Command = "stepInTargets"
	newStepInTargetsRequest.Type = "request"
	newStepInTargetsRequest.propagateFieldsToBase()
	return &newStepInTargetsRequest
}

// Returns a new `VariablesRequest` with the followings fields set: `command`, `type`
func NewVariablesRequest () *VariablesRequest {
	newVariablesRequest := VariablesRequest{}
	newVariablesRequest.Command = "variables"
	newVariablesRequest.Type = "request"
	newVariablesRequest.propagateFieldsToBase()
	return &newVariablesRequest
}

// Returns a new `ExceptionInfoRequest` with the followings fields set: `command`, `type`
func NewExceptionInfoRequest () *ExceptionInfoRequest {
	newExceptionInfoRequest := ExceptionInfoRequest{}
	newExceptionInfoRequest.Command = "exceptionInfo"
	newExceptionInfoRequest.Type = "request"
	newExceptionInfoRequest.propagateFieldsToBase()
	return &newExceptionInfoRequest
}

// Returns a new `NextRequest` with the followings fields set: `command`, `type`
func NewNextRequest () *NextRequest {
	newNextRequest := NextRequest{}
	newNextRequest.Command = "next"
	newNextRequest.Type = "request"
	newNextRequest.propagateFieldsToBase()
	return &newNextRequest
}

// Returns a new `ErrorResponse` with the followings fields set: `type`
func NewErrorResponse () *ErrorResponse {
	newErrorResponse := ErrorResponse{}
	newErrorResponse.Type = "response"
	newErrorResponse.propagateFieldsToBase()
	return &newErrorResponse
}

// Returns a new `GotoTargetsResponse` with the followings fields set: `type`
func NewGotoTargetsResponse () *GotoTargetsResponse {
	newGotoTargetsResponse := GotoTargetsResponse{}
	newGotoTargetsResponse.Type = "response"
	newGotoTargetsResponse.propagateFieldsToBase()
	return &newGotoTargetsResponse
}

// Returns a new `GotoResponse` with the followings fields set: `type`
func NewGotoResponse () *GotoResponse {
	newGotoResponse := GotoResponse{}
	newGotoResponse.Type = "response"
	newGotoResponse.propagateFieldsToBase()
	return &newGotoResponse
}

// Returns a new `StepOutRequest` with the followings fields set: `command`, `type`
func NewStepOutRequest () *StepOutRequest {
	newStepOutRequest := StepOutRequest{}
	newStepOutRequest.Command = "stepOut"
	newStepOutRequest.Type = "request"
	newStepOutRequest.propagateFieldsToBase()
	return &newStepOutRequest
}

// Returns a new `StackTraceResponse` with the followings fields set: `type`
func NewStackTraceResponse () *StackTraceResponse {
	newStackTraceResponse := StackTraceResponse{}
	newStackTraceResponse.Type = "response"
	newStackTraceResponse.propagateFieldsToBase()
	return &newStackTraceResponse
}

// Returns a new `ScopesResponse` with the followings fields set: `type`
func NewScopesResponse () *ScopesResponse {
	newScopesResponse := ScopesResponse{}
	newScopesResponse.Type = "response"
	newScopesResponse.propagateFieldsToBase()
	return &newScopesResponse
}

// Returns a new `RestartRequest` with the followings fields set: `command`, `type`
func NewRestartRequest () *RestartRequest {
	newRestartRequest := RestartRequest{}
	newRestartRequest.Command = "restart"
	newRestartRequest.Type = "request"
	newRestartRequest.propagateFieldsToBase()
	return &newRestartRequest
}

// Returns a new `RestartFrameResponse` with the followings fields set: `type`
func NewRestartFrameResponse () *RestartFrameResponse {
	newRestartFrameResponse := RestartFrameResponse{}
	newRestartFrameResponse.Type = "response"
	newRestartFrameResponse.propagateFieldsToBase()
	return &newRestartFrameResponse
}

// Returns a new `TerminatedEvent` with the followings fields set: `event`, `type`
func NewTerminatedEvent () *TerminatedEvent {
	newTerminatedEvent := TerminatedEvent{}
	newTerminatedEvent.Event_ = "terminated"
	newTerminatedEvent.Type = "event"
	newTerminatedEvent.propagateFieldsToBase()
	return &newTerminatedEvent
}

// Returns a new `ModuleEvent` with the followings fields set: `event`, `type`
func NewModuleEvent () *ModuleEvent {
	newModuleEvent := ModuleEvent{}
	newModuleEvent.Event_ = "module"
	newModuleEvent.Type = "event"
	newModuleEvent.propagateFieldsToBase()
	return &newModuleEvent
}

// Returns a new `BreakpointEvent` with the followings fields set: `event`, `type`
func NewBreakpointEvent () *BreakpointEvent {
	newBreakpointEvent := BreakpointEvent{}
	newBreakpointEvent.Event_ = "breakpoint"
	newBreakpointEvent.Type = "event"
	newBreakpointEvent.propagateFieldsToBase()
	return &newBreakpointEvent
}

// Returns a new `EvaluateResponse` with the followings fields set: `type`
func NewEvaluateResponse () *EvaluateResponse {
	newEvaluateResponse := EvaluateResponse{}
	newEvaluateResponse.Type = "response"
	newEvaluateResponse.propagateFieldsToBase()
	return &newEvaluateResponse
}

// Returns a new `ExceptionInfoResponse` with the followings fields set: `type`
func NewExceptionInfoResponse () *ExceptionInfoResponse {
	newExceptionInfoResponse := ExceptionInfoResponse{}
	newExceptionInfoResponse.Type = "response"
	newExceptionInfoResponse.propagateFieldsToBase()
	return &newExceptionInfoResponse
}

// Returns a new `ExitedEvent` with the followings fields set: `event`, `type`
func NewExitedEvent () *ExitedEvent {
	newExitedEvent := ExitedEvent{}
	newExitedEvent.Event_ = "exited"
	newExitedEvent.Type = "event"
	newExitedEvent.propagateFieldsToBase()
	return &newExitedEvent
}

// Returns a new `LaunchResponse` with the followings fields set: `type`
func NewLaunchResponse () *LaunchResponse {
	newLaunchResponse := LaunchResponse{}
	newLaunchResponse.Type = "response"
	newLaunchResponse.propagateFieldsToBase()
	return &newLaunchResponse
}

// Returns a new `ConfigurationDoneRequest` with the followings fields set: `command`, `type`
func NewConfigurationDoneRequest () *ConfigurationDoneRequest {
	newConfigurationDoneRequest := ConfigurationDoneRequest{}
	newConfigurationDoneRequest.Command = "configurationDone"
	newConfigurationDoneRequest.Type = "request"
	newConfigurationDoneRequest.propagateFieldsToBase()
	return &newConfigurationDoneRequest
}

// Returns a new `RunInTerminalRequest` with the followings fields set: `command`, `type`
func NewRunInTerminalRequest () *RunInTerminalRequest {
	newRunInTerminalRequest := RunInTerminalRequest{}
	newRunInTerminalRequest.Command = "runInTerminal"
	newRunInTerminalRequest.Type = "request"
	newRunInTerminalRequest.propagateFieldsToBase()
	return &newRunInTerminalRequest
}

// Returns a new `ThreadEvent` with the followings fields set: `event`, `type`
func NewThreadEvent () *ThreadEvent {
	newThreadEvent := ThreadEvent{}
	newThreadEvent.Event_ = "thread"
	newThreadEvent.Type = "event"
	newThreadEvent.propagateFieldsToBase()
	return &newThreadEvent
}

// Returns a new `AttachResponse` with the followings fields set: `type`
func NewAttachResponse () *AttachResponse {
	newAttachResponse := AttachResponse{}
	newAttachResponse.Type = "response"
	newAttachResponse.propagateFieldsToBase()
	return &newAttachResponse
}

// Returns a new `RunInTerminalResponse` with the followings fields set: `type`
func NewRunInTerminalResponse () *RunInTerminalResponse {
	newRunInTerminalResponse := RunInTerminalResponse{}
	newRunInTerminalResponse.Type = "response"
	newRunInTerminalResponse.propagateFieldsToBase()
	return &newRunInTerminalResponse
}

// Returns a new `RestartResponse` with the followings fields set: `type`
func NewRestartResponse () *RestartResponse {
	newRestartResponse := RestartResponse{}
	newRestartResponse.Type = "response"
	newRestartResponse.propagateFieldsToBase()
	return &newRestartResponse
}

// Returns a new `CompletionsResponse` with the followings fields set: `type`
func NewCompletionsResponse () *CompletionsResponse {
	newCompletionsResponse := CompletionsResponse{}
	newCompletionsResponse.Type = "response"
	newCompletionsResponse.propagateFieldsToBase()
	return &newCompletionsResponse
}

// Returns a new `StepBackResponse` with the followings fields set: `type`
func NewStepBackResponse () *StepBackResponse {
	newStepBackResponse := StepBackResponse{}
	newStepBackResponse.Type = "response"
	newStepBackResponse.propagateFieldsToBase()
	return &newStepBackResponse
}

// Returns a new `StepInResponse` with the followings fields set: `type`
func NewStepInResponse () *StepInResponse {
	newStepInResponse := StepInResponse{}
	newStepInResponse.Type = "response"
	newStepInResponse.propagateFieldsToBase()
	return &newStepInResponse
}

// Returns a new `ReverseContinueResponse` with the followings fields set: `type`
func NewReverseContinueResponse () *ReverseContinueResponse {
	newReverseContinueResponse := ReverseContinueResponse{}
	newReverseContinueResponse.Type = "response"
	newReverseContinueResponse.propagateFieldsToBase()
	return &newReverseContinueResponse
}

// Returns a new `DisconnectResponse` with the followings fields set: `type`
func NewDisconnectResponse () *DisconnectResponse {
	newDisconnectResponse := DisconnectResponse{}
	newDisconnectResponse.Type = "response"
	newDisconnectResponse.propagateFieldsToBase()
	return &newDisconnectResponse
}

// Returns a new `GotoTargetsRequest` with the followings fields set: `command`, `type`
func NewGotoTargetsRequest () *GotoTargetsRequest {
	newGotoTargetsRequest := GotoTargetsRequest{}
	newGotoTargetsRequest.Command = "gotoTargets"
	newGotoTargetsRequest.Type = "request"
	newGotoTargetsRequest.propagateFieldsToBase()
	return &newGotoTargetsRequest
}

// Returns a new `ContinueRequest` with the followings fields set: `command`, `type`
func NewContinueRequest () *ContinueRequest {
	newContinueRequest := ContinueRequest{}
	newContinueRequest.Command = "continue"
	newContinueRequest.Type = "request"
	newContinueRequest.propagateFieldsToBase()
	return &newContinueRequest
}

// Returns a new `ContinueResponse` with the followings fields set: `type`
func NewContinueResponse () *ContinueResponse {
	newContinueResponse := ContinueResponse{}
	newContinueResponse.Type = "response"
	newContinueResponse.propagateFieldsToBase()
	return &newContinueResponse
}

// Returns a new `SetExceptionBreakpointsResponse` with the followings fields set: `type`
func NewSetExceptionBreakpointsResponse () *SetExceptionBreakpointsResponse {
	newSetExceptionBreakpointsResponse := SetExceptionBreakpointsResponse{}
	newSetExceptionBreakpointsResponse.Type = "response"
	newSetExceptionBreakpointsResponse.propagateFieldsToBase()
	return &newSetExceptionBreakpointsResponse
}

// Returns a new `InitializeRequest` with the followings fields set: `command`, `type`
func NewInitializeRequest () *InitializeRequest {
	newInitializeRequest := InitializeRequest{}
	newInitializeRequest.Command = "initialize"
	newInitializeRequest.Type = "request"
	newInitializeRequest.propagateFieldsToBase()
	return &newInitializeRequest
}

// Returns a new `ConfigurationDoneResponse` with the followings fields set: `type`
func NewConfigurationDoneResponse () *ConfigurationDoneResponse {
	newConfigurationDoneResponse := ConfigurationDoneResponse{}
	newConfigurationDoneResponse.Type = "response"
	newConfigurationDoneResponse.propagateFieldsToBase()
	return &newConfigurationDoneResponse
}

// Returns a new `SetBreakpointsResponse` with the followings fields set: `type`
func NewSetBreakpointsResponse () *SetBreakpointsResponse {
	newSetBreakpointsResponse := SetBreakpointsResponse{}
	newSetBreakpointsResponse.Type = "response"
	newSetBreakpointsResponse.propagateFieldsToBase()
	return &newSetBreakpointsResponse
}

// Returns a new `SetVariableRequest` with the followings fields set: `command`, `type`
func NewSetVariableRequest () *SetVariableRequest {
	newSetVariableRequest := SetVariableRequest{}
	newSetVariableRequest.Command = "setVariable"
	newSetVariableRequest.Type = "request"
	newSetVariableRequest.propagateFieldsToBase()
	return &newSetVariableRequest
}

// Returns a new `StepInTargetsResponse` with the followings fields set: `type`
func NewStepInTargetsResponse () *StepInTargetsResponse {
	newStepInTargetsResponse := StepInTargetsResponse{}
	newStepInTargetsResponse.Type = "response"
	newStepInTargetsResponse.propagateFieldsToBase()
	return &newStepInTargetsResponse
}

// Returns a new `ModulesResponse` with the followings fields set: `type`
func NewModulesResponse () *ModulesResponse {
	newModulesResponse := ModulesResponse{}
	newModulesResponse.Type = "response"
	newModulesResponse.propagateFieldsToBase()
	return &newModulesResponse
}

// Returns a new `InitializeResponse` with the followings fields set: `type`
func NewInitializeResponse () *InitializeResponse {
	newInitializeResponse := InitializeResponse{}
	newInitializeResponse.Type = "response"
	newInitializeResponse.propagateFieldsToBase()
	return &newInitializeResponse
}

// Returns a new `SourceRequest` with the followings fields set: `command`, `type`
func NewSourceRequest () *SourceRequest {
	newSourceRequest := SourceRequest{}
	newSourceRequest.Command = "source"
	newSourceRequest.Type = "request"
	newSourceRequest.propagateFieldsToBase()
	return &newSourceRequest
}

// Returns a new `DisconnectRequest` with the followings fields set: `command`, `type`
func NewDisconnectRequest () *DisconnectRequest {
	newDisconnectRequest := DisconnectRequest{}
	newDisconnectRequest.Command = "disconnect"
	newDisconnectRequest.Type = "request"
	newDisconnectRequest.propagateFieldsToBase()
	return &newDisconnectRequest
}

// Returns a new `SourceResponse` with the followings fields set: `type`
func NewSourceResponse () *SourceResponse {
	newSourceResponse := SourceResponse{}
	newSourceResponse.Type = "response"
	newSourceResponse.propagateFieldsToBase()
	return &newSourceResponse
}

// Returns a new `SetVariableResponse` with the followings fields set: `type`
func NewSetVariableResponse () *SetVariableResponse {
	newSetVariableResponse := SetVariableResponse{}
	newSetVariableResponse.Type = "response"
	newSetVariableResponse.propagateFieldsToBase()
	return &newSetVariableResponse
}

// Returns a new `ContinuedEvent` with the followings fields set: `event`, `type`
func NewContinuedEvent () *ContinuedEvent {
	newContinuedEvent := ContinuedEvent{}
	newContinuedEvent.Event_ = "continued"
	newContinuedEvent.Type = "event"
	newContinuedEvent.propagateFieldsToBase()
	return &newContinuedEvent
}

// Returns a new `LaunchRequest` with the followings fields set: `command`, `type`
func NewLaunchRequest () *LaunchRequest {
	newLaunchRequest := LaunchRequest{}
	newLaunchRequest.Command = "launch"
	newLaunchRequest.Type = "request"
	newLaunchRequest.propagateFieldsToBase()
	return &newLaunchRequest
}

// Returns a new `CompletionsRequest` with the followings fields set: `command`, `type`
func NewCompletionsRequest () *CompletionsRequest {
	newCompletionsRequest := CompletionsRequest{}
	newCompletionsRequest.Command = "completions"
	newCompletionsRequest.Type = "request"
	newCompletionsRequest.propagateFieldsToBase()
	return &newCompletionsRequest
}

// Returns a new `NextResponse` with the followings fields set: `type`
func NewNextResponse () *NextResponse {
	newNextResponse := NextResponse{}
	newNextResponse.Type = "response"
	newNextResponse.propagateFieldsToBase()
	return &newNextResponse
}

// Returns a new `InitializedEvent` with the followings fields set: `event`, `type`
func NewInitializedEvent () *InitializedEvent {
	newInitializedEvent := InitializedEvent{}
	newInitializedEvent.Event_ = "initialized"
	newInitializedEvent.Type = "event"
	newInitializedEvent.propagateFieldsToBase()
	return &newInitializedEvent
}

// Returns a new `ThreadsResponse` with the followings fields set: `type`
func NewThreadsResponse () *ThreadsResponse {
	newThreadsResponse := ThreadsResponse{}
	newThreadsResponse.Type = "response"
	newThreadsResponse.propagateFieldsToBase()
	return &newThreadsResponse
}

// Returns a new `ReverseContinueRequest` with the followings fields set: `command`, `type`
func NewReverseContinueRequest () *ReverseContinueRequest {
	newReverseContinueRequest := ReverseContinueRequest{}
	newReverseContinueRequest.Command = "reverseContinue"
	newReverseContinueRequest.Type = "request"
	newReverseContinueRequest.propagateFieldsToBase()
	return &newReverseContinueRequest
}

// Returns a new `AttachRequest` with the followings fields set: `command`, `type`
func NewAttachRequest () *AttachRequest {
	newAttachRequest := AttachRequest{}
	newAttachRequest.Command = "attach"
	newAttachRequest.Type = "request"
	newAttachRequest.propagateFieldsToBase()
	return &newAttachRequest
}

// Returns a new `RestartFrameRequest` with the followings fields set: `command`, `type`
func NewRestartFrameRequest () *RestartFrameRequest {
	newRestartFrameRequest := RestartFrameRequest{}
	newRestartFrameRequest.Command = "restartFrame"
	newRestartFrameRequest.Type = "request"
	newRestartFrameRequest.propagateFieldsToBase()
	return &newRestartFrameRequest
}

// Returns a new `GotoRequest` with the followings fields set: `command`, `type`
func NewGotoRequest () *GotoRequest {
	newGotoRequest := GotoRequest{}
	newGotoRequest.Command = "goto"
	newGotoRequest.Type = "request"
	newGotoRequest.propagateFieldsToBase()
	return &newGotoRequest
}

// Returns a new `EvaluateRequest` with the followings fields set: `command`, `type`
func NewEvaluateRequest () *EvaluateRequest {
	newEvaluateRequest := EvaluateRequest{}
	newEvaluateRequest.Command = "evaluate"
	newEvaluateRequest.Type = "request"
	newEvaluateRequest.propagateFieldsToBase()
	return &newEvaluateRequest
}

// Returns a new `StoppedEvent` with the followings fields set: `event`, `type`
func NewStoppedEvent () *StoppedEvent {
	newStoppedEvent := StoppedEvent{}
	newStoppedEvent.Event_ = "stopped"
	newStoppedEvent.Type = "event"
	newStoppedEvent.propagateFieldsToBase()
	return &newStoppedEvent
}

// Returns a new `SetFunctionBreakpointsRequest` with the followings fields set: `command`, `type`
func NewSetFunctionBreakpointsRequest () *SetFunctionBreakpointsRequest {
	newSetFunctionBreakpointsRequest := SetFunctionBreakpointsRequest{}
	newSetFunctionBreakpointsRequest.Command = "setFunctionBreakpoints"
	newSetFunctionBreakpointsRequest.Type = "request"
	newSetFunctionBreakpointsRequest.propagateFieldsToBase()
	return &newSetFunctionBreakpointsRequest
}

// Returns a new `StackTraceRequest` with the followings fields set: `command`, `type`
func NewStackTraceRequest () *StackTraceRequest {
	newStackTraceRequest := StackTraceRequest{}
	newStackTraceRequest.Command = "stackTrace"
	newStackTraceRequest.Type = "request"
	newStackTraceRequest.propagateFieldsToBase()
	return &newStackTraceRequest
}

// Returns a new `SetBreakpointsRequest` with the followings fields set: `command`, `type`
func NewSetBreakpointsRequest () *SetBreakpointsRequest {
	newSetBreakpointsRequest := SetBreakpointsRequest{}
	newSetBreakpointsRequest.Command = "setBreakpoints"
	newSetBreakpointsRequest.Type = "request"
	newSetBreakpointsRequest.propagateFieldsToBase()
	return &newSetBreakpointsRequest
}

// Returns a new `PauseRequest` with the followings fields set: `command`, `type`
func NewPauseRequest () *PauseRequest {
	newPauseRequest := PauseRequest{}
	newPauseRequest.Command = "pause"
	newPauseRequest.Type = "request"
	newPauseRequest.propagateFieldsToBase()
	return &newPauseRequest
}

// Returns a new `ModulesRequest` with the followings fields set: `command`, `type`
func NewModulesRequest () *ModulesRequest {
	newModulesRequest := ModulesRequest{}
	newModulesRequest.Command = "modules"
	newModulesRequest.Type = "request"
	newModulesRequest.propagateFieldsToBase()
	return &newModulesRequest
}

// Returns a new `PauseResponse` with the followings fields set: `type`
func NewPauseResponse () *PauseResponse {
	newPauseResponse := PauseResponse{}
	newPauseResponse.Type = "response"
	newPauseResponse.propagateFieldsToBase()
	return &newPauseResponse
}

// Returns a new `OutputEvent` with the followings fields set: `event`, `type`
func NewOutputEvent () *OutputEvent {
	newOutputEvent := OutputEvent{}
	newOutputEvent.Event_ = "output"
	newOutputEvent.Type = "event"
	newOutputEvent.propagateFieldsToBase()
	return &newOutputEvent
}

// Returns a new `SetFunctionBreakpointsResponse` with the followings fields set: `type`
func NewSetFunctionBreakpointsResponse () *SetFunctionBreakpointsResponse {
	newSetFunctionBreakpointsResponse := SetFunctionBreakpointsResponse{}
	newSetFunctionBreakpointsResponse.Type = "response"
	newSetFunctionBreakpointsResponse.propagateFieldsToBase()
	return &newSetFunctionBreakpointsResponse
}


// TryUnmarshalProtocolMessage attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `ProtocolMessage` as follows:
//
// If `js` contains `"type":"request"`, attempts to unmarshal via `TryUnmarshalRequest`
//
// If `js` contains `"type":"event"`, attempts to unmarshal via `TryUnmarshalEvent`
//
// If `js` contains `"type":"response"`, attempts to unmarshal into a new `Response`.
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
	case "response":  var val Response  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "request":  ptr,err = TryUnmarshalRequest(js)
	case "event":  ptr,err = TryUnmarshalEvent(js)
	default: err = errors.New("ProtocolMessage: encountered unknown JSON value for type: " + type_of_ProtocolMessage)
	}
	return
}


// TryUnmarshalEvent attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `Event` as follows:
//
// If `js` contains `"event":"terminated"`, attempts to unmarshal into a new `TerminatedEvent`.
//
// If `js` contains `"event":"module"`, attempts to unmarshal into a new `ModuleEvent`.
//
// If `js` contains `"event":"breakpoint"`, attempts to unmarshal into a new `BreakpointEvent`.
//
// If `js` contains `"event":"continued"`, attempts to unmarshal into a new `ContinuedEvent`.
//
// If `js` contains `"event":"output"`, attempts to unmarshal into a new `OutputEvent`.
//
// If `js` contains `"event":"thread"`, attempts to unmarshal into a new `ThreadEvent`.
//
// If `js` contains `"event":"exited"`, attempts to unmarshal into a new `ExitedEvent`.
//
// If `js` contains `"event":"stopped"`, attempts to unmarshal into a new `StoppedEvent`.
//
// If `js` contains `"event":"initialized"`, attempts to unmarshal into a new `InitializedEvent`.
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
	case "continued":  var val ContinuedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "terminated":  var val TerminatedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "module":  var val ModuleEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "exited":  var val ExitedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stopped":  var val StoppedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "initialized":  var val InitializedEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "output":  var val OutputEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "thread":  var val ThreadEvent  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	default: err = errors.New("Event: encountered unknown JSON value for event: " + event_of_Event)
	}
	return
}


// TryUnmarshalRequest attempts to unmarshal JSON string `js` (if it starts with a `{` and ends with a `}`) into a `struct` based on `Request` as follows:
//
// If `js` contains `"command":"stackTrace"`, attempts to unmarshal into a new `StackTraceRequest`.
//
// If `js` contains `"command":"stepIn"`, attempts to unmarshal into a new `StepInRequest`.
//
// If `js` contains `"command":"stepOut"`, attempts to unmarshal into a new `StepOutRequest`.
//
// If `js` contains `"command":"setBreakpoints"`, attempts to unmarshal into a new `SetBreakpointsRequest`.
//
// If `js` contains `"command":"goto"`, attempts to unmarshal into a new `GotoRequest`.
//
// If `js` contains `"command":"configurationDone"`, attempts to unmarshal into a new `ConfigurationDoneRequest`.
//
// If `js` contains `"command":"modules"`, attempts to unmarshal into a new `ModulesRequest`.
//
// If `js` contains `"command":"stepBack"`, attempts to unmarshal into a new `StepBackRequest`.
//
// If `js` contains `"command":"launch"`, attempts to unmarshal into a new `LaunchRequest`.
//
// If `js` contains `"command":"disconnect"`, attempts to unmarshal into a new `DisconnectRequest`.
//
// If `js` contains `"command":"scopes"`, attempts to unmarshal into a new `ScopesRequest`.
//
// If `js` contains `"command":"threads"`, attempts to unmarshal into a new `ThreadsRequest`.
//
// If `js` contains `"command":"restartFrame"`, attempts to unmarshal into a new `RestartFrameRequest`.
//
// If `js` contains `"command":"runInTerminal"`, attempts to unmarshal into a new `RunInTerminalRequest`.
//
// If `js` contains `"command":"attach"`, attempts to unmarshal into a new `AttachRequest`.
//
// If `js` contains `"command":"pause"`, attempts to unmarshal into a new `PauseRequest`.
//
// If `js` contains `"command":"exceptionInfo"`, attempts to unmarshal into a new `ExceptionInfoRequest`.
//
// If `js` contains `"command":"initialize"`, attempts to unmarshal into a new `InitializeRequest`.
//
// If `js` contains `"command":"source"`, attempts to unmarshal into a new `SourceRequest`.
//
// If `js` contains `"command":"next"`, attempts to unmarshal into a new `NextRequest`.
//
// If `js` contains `"command":"setVariable"`, attempts to unmarshal into a new `SetVariableRequest`.
//
// If `js` contains `"command":"restart"`, attempts to unmarshal into a new `RestartRequest`.
//
// If `js` contains `"command":"setExceptionBreakpoints"`, attempts to unmarshal into a new `SetExceptionBreakpointsRequest`.
//
// If `js` contains `"command":"reverseContinue"`, attempts to unmarshal into a new `ReverseContinueRequest`.
//
// If `js` contains `"command":"evaluate"`, attempts to unmarshal into a new `EvaluateRequest`.
//
// If `js` contains `"command":"gotoTargets"`, attempts to unmarshal into a new `GotoTargetsRequest`.
//
// If `js` contains `"command":"stepInTargets"`, attempts to unmarshal into a new `StepInTargetsRequest`.
//
// If `js` contains `"command":"setFunctionBreakpoints"`, attempts to unmarshal into a new `SetFunctionBreakpointsRequest`.
//
// If `js` contains `"command":"continue"`, attempts to unmarshal into a new `ContinueRequest`.
//
// If `js` contains `"command":"completions"`, attempts to unmarshal into a new `CompletionsRequest`.
//
// If `js` contains `"command":"variables"`, attempts to unmarshal into a new `VariablesRequest`.
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
	case "stepBack":  var val StepBackRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "launch":  var val LaunchRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "modules":  var val ModulesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "scopes":  var val ScopesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "threads":  var val ThreadsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restartFrame":  var val RestartFrameRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "disconnect":  var val DisconnectRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "attach":  var val AttachRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "runInTerminal":  var val RunInTerminalRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "exceptionInfo":  var val ExceptionInfoRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "initialize":  var val InitializeRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "source":  var val SourceRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "next":  var val NextRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setVariable":  var val SetVariableRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "restart":  var val RestartRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "pause":  var val PauseRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "reverseContinue":  var val ReverseContinueRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "evaluate":  var val EvaluateRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "gotoTargets":  var val GotoTargetsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepInTargets":  var val StepInTargetsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setExceptionBreakpoints":  var val SetExceptionBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "continue":  var val ContinueRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "completions":  var val CompletionsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "variables":  var val VariablesRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setFunctionBreakpoints":  var val SetFunctionBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepIn":  var val StepInRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stepOut":  var val StepOutRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "setBreakpoints":  var val SetBreakpointsRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "stackTrace":  var val StackTraceRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "configurationDone":  var val ConfigurationDoneRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	case "goto":  var val GotoRequest  ;  if err = json.Unmarshal([]byte(js), &val); err==nil { val.propagateFieldsToBase()  ;  ptr = &val }
	default: err = errors.New("Request: encountered unknown JSON value for command: " + command_of_Request)
	}
	return
}

// Called by `HandleRequest` (after it unmarshaled the given `RunInTerminalRequest`) to further populate the given `RunInTerminalResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRunInTerminalRequest func(*RunInTerminalRequest, *RunInTerminalResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `GotoTargetsRequest`) to further populate the given `GotoTargetsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnGotoTargetsRequest func(*GotoTargetsRequest, *GotoTargetsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetBreakpointsRequest`) to further populate the given `SetBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetBreakpointsRequest func(*SetBreakpointsRequest, *SetBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepOutRequest`) to further populate the given `StepOutResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepOutRequest func(*StepOutRequest, *StepOutResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetVariableRequest`) to further populate the given `SetVariableResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetVariableRequest func(*SetVariableRequest, *SetVariableResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `InitializeRequest`) to further populate the given `InitializeResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnInitializeRequest func(*InitializeRequest, *InitializeResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ContinueRequest`) to further populate the given `ContinueResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnContinueRequest func(*ContinueRequest, *ContinueResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ReverseContinueRequest`) to further populate the given `ReverseContinueResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnReverseContinueRequest func(*ReverseContinueRequest, *ReverseContinueResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `EvaluateRequest`) to further populate the given `EvaluateResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnEvaluateRequest func(*EvaluateRequest, *EvaluateResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `DisconnectRequest`) to further populate the given `DisconnectResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnDisconnectRequest func(*DisconnectRequest, *DisconnectResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `CompletionsRequest`) to further populate the given `CompletionsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnCompletionsRequest func(*CompletionsRequest, *CompletionsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepBackRequest`) to further populate the given `StepBackResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepBackRequest func(*StepBackRequest, *StepBackResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ExceptionInfoRequest`) to further populate the given `ExceptionInfoResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnExceptionInfoRequest func(*ExceptionInfoRequest, *ExceptionInfoResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ThreadsRequest`) to further populate the given `ThreadsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnThreadsRequest func(*ThreadsRequest, *ThreadsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ScopesRequest`) to further populate the given `ScopesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnScopesRequest func(*ScopesRequest, *ScopesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `GotoRequest`) to further populate the given `GotoResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnGotoRequest func(*GotoRequest, *GotoResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SourceRequest`) to further populate the given `SourceResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSourceRequest func(*SourceRequest, *SourceResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepInRequest`) to further populate the given `StepInResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepInRequest func(*StepInRequest, *StepInResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `VariablesRequest`) to further populate the given `VariablesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnVariablesRequest func(*VariablesRequest, *VariablesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StackTraceRequest`) to further populate the given `StackTraceResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStackTraceRequest func(*StackTraceRequest, *StackTraceResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `NextRequest`) to further populate the given `NextResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnNextRequest func(*NextRequest, *NextResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `RestartRequest`) to further populate the given `RestartResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRestartRequest func(*RestartRequest, *RestartResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `PauseRequest`) to further populate the given `PauseResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnPauseRequest func(*PauseRequest, *PauseResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `RestartFrameRequest`) to further populate the given `RestartFrameResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnRestartFrameRequest func(*RestartFrameRequest, *RestartFrameResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetFunctionBreakpointsRequest`) to further populate the given `SetFunctionBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetFunctionBreakpointsRequest func(*SetFunctionBreakpointsRequest, *SetFunctionBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `StepInTargetsRequest`) to further populate the given `StepInTargetsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnStepInTargetsRequest func(*StepInTargetsRequest, *StepInTargetsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ModulesRequest`) to further populate the given `ModulesResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnModulesRequest func(*ModulesRequest, *ModulesResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `SetExceptionBreakpointsRequest`) to further populate the given `SetExceptionBreakpointsResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnSetExceptionBreakpointsRequest func(*SetExceptionBreakpointsRequest, *SetExceptionBreakpointsResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `ConfigurationDoneRequest`) to further populate the given `ConfigurationDoneResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnConfigurationDoneRequest func(*ConfigurationDoneRequest, *ConfigurationDoneResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `AttachRequest`) to further populate the given `AttachResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnAttachRequest func(*AttachRequest, *AttachResponse)error

// Called by `HandleRequest` (after it unmarshaled the given `LaunchRequest`) to further populate the given `LaunchResponse` before returning it to its caller (in addition to this handler's returned `error`).
var OnLaunchRequest func(*LaunchRequest, *LaunchResponse)error

// If a type-switch on `inRequest` succeeds, `outResponse` points to a `Response`-based `struct` value containing the `Response` initialized by the specified `initNewResponse` and further populated by the `OnFooRequest` handler corresponding to the concrete type of `inRequest` (if any). The only `err` returned, if any, is that returned by the specialized `OnFooRequest` handler.
func HandleRequest (inRequest interface{}, initNewResponse func(*Request, *Response)) (outResponse interface{}, baseResponse *Response, err error) {
	switch input := inRequest.(type) {
	case *ScopesRequest:
		o := NewScopesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnScopesRequest!=nil { err = OnScopesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *GotoRequest:
		o := NewGotoResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnGotoRequest!=nil { err = OnGotoRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ThreadsRequest:
		o := NewThreadsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnThreadsRequest!=nil { err = OnThreadsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SourceRequest:
		o := NewSourceResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnSourceRequest!=nil { err = OnSourceRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepInRequest:
		o := NewStepInResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnStepInRequest!=nil { err = OnStepInRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *VariablesRequest:
		o := NewVariablesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnVariablesRequest!=nil { err = OnVariablesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StackTraceRequest:
		o := NewStackTraceResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnStackTraceRequest!=nil { err = OnStackTraceRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *PauseRequest:
		o := NewPauseResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnPauseRequest!=nil { err = OnPauseRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RestartFrameRequest:
		o := NewRestartFrameResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnRestartFrameRequest!=nil { err = OnRestartFrameRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetFunctionBreakpointsRequest:
		o := NewSetFunctionBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnSetFunctionBreakpointsRequest!=nil { err = OnSetFunctionBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepInTargetsRequest:
		o := NewStepInTargetsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnStepInTargetsRequest!=nil { err = OnStepInTargetsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *NextRequest:
		o := NewNextResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnNextRequest!=nil { err = OnNextRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RestartRequest:
		o := NewRestartResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnRestartRequest!=nil { err = OnRestartRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ConfigurationDoneRequest:
		o := NewConfigurationDoneResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnConfigurationDoneRequest!=nil { err = OnConfigurationDoneRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *AttachRequest:
		o := NewAttachResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnAttachRequest!=nil { err = OnAttachRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *LaunchRequest:
		o := NewLaunchResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnLaunchRequest!=nil { err = OnLaunchRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ModulesRequest:
		o := NewModulesResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnModulesRequest!=nil { err = OnModulesRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetExceptionBreakpointsRequest:
		o := NewSetExceptionBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnSetExceptionBreakpointsRequest!=nil { err = OnSetExceptionBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetBreakpointsRequest:
		o := NewSetBreakpointsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnSetBreakpointsRequest!=nil { err = OnSetBreakpointsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepOutRequest:
		o := NewStepOutResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnStepOutRequest!=nil { err = OnStepOutRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *SetVariableRequest:
		o := NewSetVariableResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnSetVariableRequest!=nil { err = OnSetVariableRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *RunInTerminalRequest:
		o := NewRunInTerminalResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnRunInTerminalRequest!=nil { err = OnRunInTerminalRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *GotoTargetsRequest:
		o := NewGotoTargetsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnGotoTargetsRequest!=nil { err = OnGotoTargetsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ReverseContinueRequest:
		o := NewReverseContinueResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnReverseContinueRequest!=nil { err = OnReverseContinueRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *EvaluateRequest:
		o := NewEvaluateResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnEvaluateRequest!=nil { err = OnEvaluateRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *DisconnectRequest:
		o := NewDisconnectResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnDisconnectRequest!=nil { err = OnDisconnectRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *CompletionsRequest:
		o := NewCompletionsResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnCompletionsRequest!=nil { err = OnCompletionsRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *StepBackRequest:
		o := NewStepBackResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnStepBackRequest!=nil { err = OnStepBackRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *InitializeRequest:
		o := NewInitializeResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnInitializeRequest!=nil { err = OnInitializeRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ContinueRequest:
		o := NewContinueResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnContinueRequest!=nil { err = OnContinueRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	case *ExceptionInfoRequest:
		o := NewExceptionInfoResponse()
		if initNewResponse!=nil { initNewResponse(&input.Request, &o.Response)  ;  o.propagateFieldsToBase() }
		if OnExceptionInfoRequest!=nil { err = OnExceptionInfoRequest(input, o)  ;  o.propagateFieldsToBase() }
		outResponse , baseResponse = o , &o.Response
	}
	return
}
