// VS Code Debug Protocol
// 
// A json schema for the VS Code Debug Protocol
// 
// Package codegen'd from github.com/metaleap/zentient/_notes_misc_etc/vscdbgprotocol.json via github.com/metaleap/zentient/dbg/zentient-debug-protocol-gen
package zdbgproto



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



// Event message for 'breakpoint' event type.
// The event indicates that some information about a breakpoint has changed.
type BreakpointEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `breakpoint`
	Event_ string `json:"event"`

	Body struct{

		// The reason for the event (such as: 'changed', 'new').
		// 
		// POSSIBLE VALUES: ``
		Reason string `json:"reason"`

		// The breakpoint.
		Breakpoint Breakpoint `json:"breakpoint"`

	} `json:"body"`
} // struct BreakpointEvent



// Response to 'setExceptionBreakpoints' request. This is just an acknowledgement, so no body field is required.
type SetExceptionBreakpointsResponse struct {
	// Response to a request.
	Response
} // struct SetExceptionBreakpointsResponse



// Arguments for 'setBreakpoints' request.
type SetBreakpointsArguments struct {

	// The code locations of the breakpoints.
	Breakpoints []SourceBreakpoint `json:"breakpoints,omitempty"`

	// Deprecated: The code locations of the breakpoints.
	Lines []int `json:"lines,omitempty"`

	// A value of true indicates that the underlying source has been modified which results in new breakpoint locations.
	SourceModified bool `json:"sourceModified,omitempty"`

	// The source location of the breakpoints; either source.path or source.reference must be specified.
	Source Source `json:"source"`
} // struct SetBreakpointsArguments



// Disconnect request; value of command field is 'disconnect'.
type DisconnectRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `disconnect`
	Command string `json:"command"`

	Arguments DisconnectArguments `json:"arguments,omitempty"`
} // struct DisconnectRequest



// A Source is a descriptor for source code. It is returned from the debug adapter as part of a StackFrame and it is used by clients when specifying breakpoints.
type Source struct {

	// Optional data that a debug adapter might want to loop through the client. The client should leave the data intact and persist it across sessions. The client should not interpret the data.
	// 
	// POSSIBLE TYPES: `[]interface{}`, `bool`, `int`, `interface{/*nil*/}`, `int64`, `map[string]interface{}`, `string`
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



// Event message for 'exited' event type.
// The event indicates that the debuggee has exited.
type ExitedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `exited`
	Event_ string `json:"event"`

	Body struct{

		// The exit code returned from the debuggee.
		ExitCode int `json:"exitCode"`

	} `json:"body"`
} // struct ExitedEvent



// Arguments for 'gotoTargets' request.
type GotoTargetsArguments struct {

	// The source location for which the goto targets are determined.
	Source Source `json:"source"`

	// The line location for which the goto targets are determined.
	Line int `json:"line"`

	// An optional column location for which the goto targets are determined.
	Column int `json:"column,omitempty"`
} // struct GotoTargetsArguments



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



// Response to 'modules' request.
type ModulesResponse struct {
	// Response to a request.
	Response

	Body struct{

		// All modules or range of modules.
		Modules []Module `json:"modules"`

		// The total number of modules available.
		TotalModules int `json:"totalModules,omitempty"`

	} `json:"body"`
} // struct ModulesResponse



// Response to 'stepOut' request. This is just an acknowledgement, so no body field is required.
type StepOutResponse struct {
	// Response to a request.
	Response
} // struct StepOutResponse



// Provides formatting information for a stack frame.
type StackFrameFormat struct {
	// Provides formatting information for a value.
	ValueFormat

	// Displays parameters for the stack frame.
	Parameters bool `json:"parameters,omitempty"`

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
} // struct StackFrameFormat



// Event message for 'thread' event type.
// The event indicates that a thread has started or exited.
type ThreadEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `thread`
	Event_ string `json:"event"`

	Body struct{

		// The reason for the event (such as: 'started', 'exited').
		// 
		// POSSIBLE VALUES: ``
		Reason string `json:"reason"`

		// The identifier of the thread.
		ThreadId int `json:"threadId"`

	} `json:"body"`
} // struct ThreadEvent



// Initialize request; value of command field is 'initialize'.
type InitializeRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `initialize`
	Command string `json:"command"`

	Arguments InitializeRequestArguments `json:"arguments"`
} // struct InitializeRequest



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



// Event message for 'stopped' event type.
// The event indicates that the execution of the debuggee has stopped due to some condition.
// This can be caused by a break point previously set, a stepping action has completed, by executing a debugger statement etc.
type StoppedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `stopped`
	Event_ string `json:"event"`

	Body struct{

		// The reason for the event (such as: 'step', 'breakpoint', 'exception', 'pause', 'entry').
		// For backward compatibility this string is shown in the UI if the 'description' attribute is missing (but it must not be translated).
		// 
		// POSSIBLE VALUES: ``
		Reason string `json:"reason"`

		// The full reason for the event, e.g. 'Paused on exception'. This string is shown in the UI as is.
		Description string `json:"description,omitempty"`

		// The thread which was stopped.
		ThreadId int `json:"threadId,omitempty"`

		// Additional information. E.g. if reason is 'exception', text contains the exception name. This string is shown in the UI.
		Text string `json:"text,omitempty"`

		// If allThreadsStopped is true, a debug adapter can announce that all threads have stopped.
		// *  The client should use this information to enable that all threads can be expanded to access their stacktraces.
		// *  If the attribute is missing or false, only the thread with the given threadId can be expanded.
		AllThreadsStopped bool `json:"allThreadsStopped,omitempty"`

	} `json:"body"`
} // struct StoppedEvent



// The checksum of an item calculated by the specified algorithm.
type Checksum struct {

	// The algorithm used to calculate this checksum.
	Algorithm ChecksumAlgorithm `json:"algorithm"`

	// Value of the checksum.
	Checksum string `json:"checksum"`
} // struct Checksum



// Arguments for 'stepBack' request.
type StepBackArguments struct {

	// Exceute 'stepBack' for this thread.
	ThreadId int `json:"threadId"`
} // struct StepBackArguments



// A Scope is a named container for variables. Optionally a scope can map to a source or a range within a source.
type Scope struct {

	// The number of named variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	NamedVariables int `json:"namedVariables,omitempty"`

	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive"`

	// Name of the scope such as 'Arguments', 'Locals'.
	Name string `json:"name"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

	// The number of indexed variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	IndexedVariables int `json:"indexedVariables,omitempty"`

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
} // struct Scope



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



// Response to 'setFunctionBreakpoints' request.
// Returned is information about each breakpoint created by this request.
type SetFunctionBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// Information about the breakpoints. The array elements correspond to the elements of the 'breakpoints' array.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`
} // struct SetFunctionBreakpointsResponse



// Launch request; value of command field is 'launch'.
type LaunchRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `launch`
	Command string `json:"command"`

	Arguments LaunchRequestArguments `json:"arguments"`
} // struct LaunchRequest



// Attach request; value of command field is 'attach'.
type AttachRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `attach`
	Command string `json:"command"`

	Arguments AttachRequestArguments `json:"arguments"`
} // struct AttachRequest



// Modules can be retrieved from the debug adapter with the ModulesRequest which can either return all modules or a range of modules to support paging.
type ModulesRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `modules`
	Command string `json:"command"`

	Arguments ModulesArguments `json:"arguments"`
} // struct ModulesRequest



// Response to 'goto' request. This is just an acknowledgement, so no body field is required.
type GotoResponse struct {
	// Response to a request.
	Response
} // struct GotoResponse



// Names of checksum algorithms that may be supported by a debug adapter.
// 
// POSSIBLE VALUES: `MD5`, `SHA1`, `SHA256`, `timestamp`
type ChecksumAlgorithm string 



// ReverseContinue request; value of command field is 'reverseContinue'.
// The request starts the debuggee to run backward. Clients should only call this request if the capability supportsStepBack is true.
type ReverseContinueRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `reverseContinue`
	Command string `json:"command"`

	Arguments ReverseContinueArguments `json:"arguments"`
} // struct ReverseContinueRequest



// Arguments for 'completions' request.
type CompletionsArguments struct {

	// Returns completions in the scope of this stack frame. If not specified, the completions are returned for the global scope.
	FrameId int `json:"frameId,omitempty"`

	// One or more source lines. Typically this is the text a user has typed into the debug console before he asked for completion.
	Text string `json:"text"`

	// The character position for which to determine the completion proposals.
	Column int `json:"column"`

	// An optional line for which to determine the completion proposals. If missing the first line of the text is assumed.
	Line int `json:"line,omitempty"`
} // struct CompletionsArguments



// Response to 'initialize' request.
type InitializeResponse struct {
	// Response to a request.
	Response

	// The capabilities of this debug adapter.
	Body Capabilities `json:"body,omitempty"`
} // struct InitializeResponse



// Event message for 'continued' event type.
// The event indicates that the execution of the debuggee has continued.
// Please note: a debug adapter is not expected to send this event in response to a request that implies that execution continues, e.g. 'launch' or 'continue'.
// It is only necessary to send a ContinuedEvent if there was no previous request that implied this.
type ContinuedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `continued`
	Event_ string `json:"event"`

	Body struct{

		// If allThreadsContinued is true, a debug adapter can announce that all threads have continued.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

		// The thread which was continued.
		ThreadId int `json:"threadId"`

	} `json:"body"`
} // struct ContinuedEvent



// Arguments for 'configurationDone' request.
// The configurationDone request has no standardized attributes.
type ConfigurationDoneArguments struct {
} // struct ConfigurationDoneArguments



// Thread request; value of command field is 'threads'.
// The request retrieves a list of all threads.
type ThreadsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `threads`
	Command string `json:"command"`
} // struct ThreadsRequest



// Arguments for 'setExceptionBreakpoints' request.
type SetExceptionBreakpointsArguments struct {

	// IDs of checked exception options. The set of IDs is returned via the 'exceptionBreakpointFilters' capability.
	Filters []string `json:"filters"`

	// Configuration options for selected exceptions.
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`
} // struct SetExceptionBreakpointsArguments



// An ExceptionBreakpointsFilter is shown in the UI as an option for configuring how exceptions are dealt with.
type ExceptionBreakpointsFilter struct {

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter"`

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label"`

	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`
} // struct ExceptionBreakpointsFilter



// On error that is whenever 'success' is false, the body can provide more details.
type ErrorResponse struct {
	// Response to a request.
	Response

	Body struct{

		// An optional, structured error message.
		Error Message `json:"error,omitempty"`

	} `json:"body"`
} // struct ErrorResponse



// Response to 'setBreakpoints' request.
// Returned is information about each breakpoint created by this request.
// This includes the actual code location and whether the breakpoint could be verified.
// The breakpoints returned are in the same order as the elements of the 'breakpoints'
// (or the deprecated 'lines') in the SetBreakpointsArguments.
type SetBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// Information about the breakpoints. The array elements are in the same order as the elements of the 'breakpoints' (or the deprecated 'lines') in the SetBreakpointsArguments.
		Breakpoints []Breakpoint `json:"breakpoints"`

	} `json:"body"`
} // struct SetBreakpointsResponse



// Response to 'gotoTargets' request.
type GotoTargetsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The possible goto targets of the specified location.
		Targets []GotoTarget `json:"targets"`

	} `json:"body"`
} // struct GotoTargetsResponse



// This enumeration defines all possible conditions when a thrown exception should result in a break.
// never: never breaks,
// always: always breaks,
// unhandled: breaks when excpetion unhandled,
// userUnhandled: breaks if the exception is not handled by user code.
// 
// POSSIBLE VALUES: `never`, `always`, `unhandled`, `userUnhandled`
type ExceptionBreakMode string 



// Continue request; value of command field is 'continue'.
// The request starts the debuggee to run again.
type ContinueRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `continue`
	Command string `json:"command"`

	Arguments ContinueArguments `json:"arguments"`
} // struct ContinueRequest



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



// A structured message object. Used to return errors from requests.
type Message struct {

	// An object used as a dictionary for looking up the variables in the format string.
	Variables map[string]string `json:"variables,omitempty"`

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
} // struct Message



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
	// POSSIBLE TYPES: `[]interface{}`, `bool`, `int`, `interface{/*nil*/}`, `int64`, `map[string]interface{}`, `string`
	Body interface{} `json:"body,omitempty"`

	// POSSIBLE VALUES: `response`
	Type string `json:"type"`

	// Sequence number of the corresponding request.
	Request_seq int `json:"request_seq"`

	// Outcome of the request.
	Success bool `json:"success"`
} // struct Response



// Response to 'setVariable' request.
type SetVariableResponse struct {
	// Response to a request.
	Response

	Body struct{

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



// StackTrace request; value of command field is 'stackTrace'. The request returns a stacktrace from the current execution state.
type StackTraceRequest struct {
	// A client or server-initiated request.
	Request

	Arguments StackTraceArguments `json:"arguments"`

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command"`
} // struct StackTraceRequest



// Properties of a breakpoint passed to the setFunctionBreakpoints request.
type FunctionBreakpoint struct {

	// The name of the function.
	Name string `json:"name"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`
} // struct FunctionBreakpoint



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



// Information about a Breakpoint created in setBreakpoints or setFunctionBreakpoints.
type Breakpoint struct {

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

	// If true breakpoint could be set (but not necessarily at the desired location).
	Verified bool `json:"verified"`

	// An optional message about the state of the breakpoint. This is shown to the user and can be used to explain why a breakpoint could not be verified.
	Message string `json:"message,omitempty"`

	// The source where the breakpoint is located.
	Source Source `json:"source,omitempty"`
} // struct Breakpoint



// Arguments for 'evaluate' request.
type EvaluateArguments struct {

	// The expression to evaluate.
	Expression string `json:"expression"`

	// Evaluate the expression in the scope of this stack frame. If not specified, the expression is evaluated in the global scope.
	FrameId int `json:"frameId,omitempty"`

	// The context in which the evaluate request is run. Possible values are 'watch' if evaluate is run in a watch, 'repl' if run from the REPL console, or 'hover' if run from a data hover.
	// 
	// POSSIBLE VALUES: ``
	Context string `json:"context,omitempty"`

	// Specifies details on how to format the Evaluate result.
	Format ValueFormat `json:"format,omitempty"`
} // struct EvaluateArguments



// StepOut request; value of command field is 'stepOut'.
// The request starts the debuggee to run again for one step.
// The debug adapter first sends the StepOutResponse and then a StoppedEvent (event type 'step') after the step has completed.
type StepOutRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command"`

	Arguments StepOutArguments `json:"arguments"`
} // struct StepOutRequest



// Arguments for 'launch' request.
type LaunchRequestArguments struct {

	// If noDebug is true the launch request should launch the program without enabling debugging.
	NoDebug bool `json:"noDebug,omitempty"`
} // struct LaunchRequestArguments



// runInTerminal request; value of command field is 'runInTerminal'.
// With this request a debug adapter can run a command in a terminal.
type RunInTerminalRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `runInTerminal`
	Command string `json:"command"`

	Arguments RunInTerminalRequestArguments `json:"arguments"`
} // struct RunInTerminalRequest



// ConfigurationDone request; value of command field is 'configurationDone'.
// The client of the debug protocol must send this request at the end of the sequence of configuration requests (which was started by the InitializedEvent).
type ConfigurationDoneRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `configurationDone`
	Command string `json:"command"`

	Arguments ConfigurationDoneArguments `json:"arguments,omitempty"`
} // struct ConfigurationDoneRequest



// Arguments for 'restart' request.
// The restart request has no standardized attributes.
type RestartArguments struct {
} // struct RestartArguments



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



// Response to 'variables' request.
type VariablesResponse struct {
	// Response to a request.
	Response

	Body struct{

		// All (or a range) of variables for the given variable reference.
		Variables []Variable `json:"variables"`

	} `json:"body"`
} // struct VariablesResponse



// Response to 'launch' request. This is just an acknowledgement, so no body field is required.
type LaunchResponse struct {
	// Response to a request.
	Response
} // struct LaunchResponse



// Arguments for 'restartFrame' request.
type RestartFrameArguments struct {

	// Restart this stackframe.
	FrameId int `json:"frameId"`
} // struct RestartFrameArguments



// Arguments for 'variables' request.
type VariablesArguments struct {

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

	// Specifies details on how to format the Variable values.
	Format ValueFormat `json:"format,omitempty"`
} // struct VariablesArguments



// Arguments for 'stepInTargets' request.
type StepInTargetsArguments struct {

	// The stack frame for which to retrieve the possible stepIn targets.
	FrameId int `json:"frameId"`
} // struct StepInTargetsArguments



// Response to 'stackTrace' request.
type StackTraceResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The frames of the stackframe. If the array has length zero, there are no stackframes available.
		// This means that there is no location information available.
		StackFrames []StackFrame `json:"stackFrames"`

		// The total number of frames available.
		TotalFrames int `json:"totalFrames,omitempty"`

	} `json:"body"`
} // struct StackTraceResponse



// The ModulesViewDescriptor is the container for all declarative configuration options of a ModuleView.
// For now it only specifies the columns to be shown in the modules view.
type ModulesViewDescriptor struct {

	Columns []ColumnDescriptor `json:"columns"`
} // struct ModulesViewDescriptor



// Server-initiated event.
type Event struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

	// Type of event.
	Event string `json:"event"`

	// Event-specific information.
	// 
	// POSSIBLE TYPES: `[]interface{}`, `bool`, `int`, `interface{/*nil*/}`, `int64`, `map[string]interface{}`, `string`
	Body interface{} `json:"body,omitempty"`

	// POSSIBLE VALUES: `event`
	Type string `json:"type"`
} // struct Event



// Arguments for 'scopes' request.
type ScopesArguments struct {

	// Retrieve the scopes for this stackframe.
	FrameId int `json:"frameId"`
} // struct ScopesArguments



// Response to 'restart' request. This is just an acknowledgement, so no body field is required.
type RestartResponse struct {
	// Response to a request.
	Response
} // struct RestartResponse



// Event message for 'module' event type.
// The event indicates that some information about a module has changed.
type ModuleEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `module`
	Event_ string `json:"event"`

	Body struct{

		// The reason for the event.
		// 
		// POSSIBLE VALUES: `new`, `changed`, `removed`
		Reason string `json:"reason"`

		// The new, changed, or removed module. In case of 'removed' only the module id is used.
		Module Module `json:"module"`

	} `json:"body"`
} // struct ModuleEvent



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
	// POSSIBLE VALUES: ``
	PathFormat string `json:"pathFormat,omitempty"`

	// Client supports the optional type attribute for variables.
	SupportsVariableType bool `json:"supportsVariableType,omitempty"`

	// Client supports the paging of variables.
	SupportsVariablePaging bool `json:"supportsVariablePaging,omitempty"`

	// Client supports the runInTerminal request.
	SupportsRunInTerminalRequest bool `json:"supportsRunInTerminalRequest,omitempty"`
} // struct InitializeRequestArguments



// An ExceptionPathSegment represents a segment in a path that is used to match leafs or nodes in a tree of exceptions. If a segment consists of more than one name, it matches the names provided if 'negate' is false or missing or it matches anything except the names provided if 'negate' is true.
type ExceptionPathSegment struct {

	// If false or missing this segment matches the names provided, otherwise it matches anything except the names provided.
	Negate bool `json:"negate,omitempty"`

	// Depending on the value of 'negate' the names that should match or not match.
	Names []string `json:"names"`
} // struct ExceptionPathSegment



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

	// Environment key-value pairs that are added to the default environment.
	Env map[string]string `json:"env,omitempty"`
} // struct RunInTerminalRequestArguments



// Arguments for 'stepIn' request.
type StepInArguments struct {

	// Execute 'stepIn' for this thread.
	ThreadId int `json:"threadId"`

	// Optional id of the target to step into.
	TargetId int `json:"targetId,omitempty"`
} // struct StepInArguments



// Arguments for 'reverseContinue' request.
type ReverseContinueArguments struct {

	// Exceute 'reverseContinue' for this thread.
	ThreadId int `json:"threadId"`
} // struct ReverseContinueArguments



// Arguments for 'exceptionInfo' request.
type ExceptionInfoArguments struct {

	// Thread for which exception information should be retrieved.
	ThreadId int `json:"threadId"`
} // struct ExceptionInfoArguments



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



// Event message for 'output' event type.
// The event indicates that the target has produced some output.
type OutputEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `output`
	Event_ string `json:"event"`

	Body struct{

		// The category of output (such as: 'console', 'stdout', 'stderr', 'telemetry'). If not specified, 'console' is assumed.
		// 
		// POSSIBLE VALUES: ``
		Category string `json:"category,omitempty"`

		// The output to report.
		Output string `json:"output"`

		// If an attribute 'variablesReference' exists and its value is > 0, the output contains objects which can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference,omitempty"`

		// Optional data to report. For the 'telemetry' category the data will be sent to telemetry, for the other categories the data is shown in JSON format.
		// 
		// POSSIBLE TYPES: `[]interface{}`, `bool`, `int`, `interface{/*nil*/}`, `int64`, `map[string]interface{}`, `string`
		Data interface{} `json:"data,omitempty"`

	} `json:"body"`
} // struct OutputEvent



// Response to 'pause' request. This is just an acknowledgement, so no body field is required.
type PauseResponse struct {
	// Response to a request.
	Response
} // struct PauseResponse



// Arguments for 'pause' request.
type PauseArguments struct {

	// Pause execution for this thread.
	ThreadId int `json:"threadId"`
} // struct PauseArguments



// Response to 'attach' request. This is just an acknowledgement, so no body field is required.
type AttachResponse struct {
	// Response to a request.
	Response
} // struct AttachResponse



// Event message for 'terminated' event types.
// The event indicates that debugging of the debuggee has terminated.
type TerminatedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `terminated`
	Event_ string `json:"event"`

	Body struct{

		// A debug adapter may set 'restart' to true to request that the front end restarts the session.
		Restart bool `json:"restart,omitempty"`

	} `json:"body,omitempty"`
} // struct TerminatedEvent



// Response to 'stepInTargets' request.
type StepInTargetsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The possible stepIn targets of the specified source location.
		Targets []StepInTarget `json:"targets"`

	} `json:"body"`
} // struct StepInTargetsResponse



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
	// POSSIBLE TYPES: `[]interface{}`, `bool`, `int`, `interface{/*nil*/}`, `int64`, `map[string]interface{}`, `string`
	Arguments interface{} `json:"arguments,omitempty"`
} // struct Request



// Arguments for 'continue' request.
type ContinueArguments struct {

	// Continue execution for the specified thread (if possible). If the backend cannot continue on a single thread but will continue on all threads, it should set the allThreadsContinued attribute in the response to true.
	ThreadId int `json:"threadId"`
} // struct ContinueArguments



// Properties of a breakpoint passed to the setBreakpoints request.
type SourceBreakpoint struct {

	// The source line of the breakpoint.
	Line int `json:"line"`

	// An optional source column of the breakpoint.
	Column int `json:"column,omitempty"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`
} // struct SourceBreakpoint



// Arguments for 'next' request.
type NextArguments struct {

	// Execute 'next' for this thread.
	ThreadId int `json:"threadId"`
} // struct NextArguments



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



// Response to 'evaluate' request.
type EvaluateResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The optional type of the evaluate result.
		Type string `json:"type,omitempty"`

		// If variablesReference is > 0, the evaluate result is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
		VariablesReference int64 `json:"variablesReference"`

		// The number of named child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		NamedVariables int64 `json:"namedVariables,omitempty"`

		// The number of indexed child variables.
		// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
		IndexedVariables int64 `json:"indexedVariables,omitempty"`

		// The result of the evaluate request.
		Result string `json:"result"`

	} `json:"body"`
} // struct EvaluateResponse



// Information about the capabilities of a debug adapter.
type Capabilities struct {

	// The debug adapter supports restarting a frame.
	SupportsRestartFrame bool `json:"supportsRestartFrame,omitempty"`

	// The debug adapter supports the completionsRequest.
	SupportsCompletionsRequest bool `json:"supportsCompletionsRequest,omitempty"`

	// The debug adapter supports the RestartRequest. In this case a client should not implement 'restart' by terminating and relaunching the adapter but by calling the RestartRequest.
	SupportsRestartRequest bool `json:"supportsRestartRequest,omitempty"`

	// The debug adapter supports the stepInTargetsRequest.
	SupportsStepInTargetsRequest bool `json:"supportsStepInTargetsRequest,omitempty"`

	// The debug adapter supports 'exceptionOptions' on the setExceptionBreakpoints request.
	SupportsExceptionOptions bool `json:"supportsExceptionOptions,omitempty"`

	// The debug adapter supports a (side effect free) evaluate request for data hovers.
	SupportsEvaluateForHovers bool `json:"supportsEvaluateForHovers,omitempty"`

	// The debug adapter supports setting a variable to a value.
	SupportsSetVariable bool `json:"supportsSetVariable,omitempty"`

	// The debug adapter supports the gotoTargetsRequest.
	SupportsGotoTargetsRequest bool `json:"supportsGotoTargetsRequest,omitempty"`

	// The debug adapter supports the modules request.
	SupportsModulesRequest bool `json:"supportsModulesRequest,omitempty"`

	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []ColumnDescriptor `json:"additionalModuleColumns,omitempty"`

	// The debug adapter supports a 'format' attribute on the stackTraceRequest, variablesRequest, and evaluateRequest.
	SupportsValueFormattingOptions bool `json:"supportsValueFormattingOptions,omitempty"`

	// The debug adapter supports the exceptionInfo request.
	SupportsExceptionInfoRequest bool `json:"supportsExceptionInfoRequest,omitempty"`

	// The debug adapter supports the configurationDoneRequest.
	SupportsConfigurationDoneRequest bool `json:"supportsConfigurationDoneRequest,omitempty"`

	// The debug adapter supports function breakpoints.
	SupportsFunctionBreakpoints bool `json:"supportsFunctionBreakpoints,omitempty"`

	// The debug adapter supports conditional breakpoints.
	SupportsConditionalBreakpoints bool `json:"supportsConditionalBreakpoints,omitempty"`

	// Checksum algorithms supported by the debug adapter.
	SupportedChecksumAlgorithms []ChecksumAlgorithm `json:"supportedChecksumAlgorithms,omitempty"`

	// The debug adapter supports breakpoints that break execution after a specified number of hits.
	SupportsHitConditionalBreakpoints bool `json:"supportsHitConditionalBreakpoints,omitempty"`

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

	// The debug adapter supports stepping back via the stepBack and reverseContinue requests.
	SupportsStepBack bool `json:"supportsStepBack,omitempty"`
} // struct Capabilities



// Response to 'source' request.
type SourceResponse struct {
	// Response to a request.
	Response

	Body struct{

		// Content of the source reference.
		Content string `json:"content"`

		// Optional content type (mime type) of the source.
		MimeType string `json:"mimeType,omitempty"`

	} `json:"body"`
} // struct SourceResponse



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

	Arguments StepInArguments `json:"arguments"`

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command"`
} // struct StepInRequest



// An ExceptionOptions assigns configuration options to a set of exceptions.
type ExceptionOptions struct {

	// A path that selects a single or multiple exceptions in a tree. If 'path' is missing, the whole tree is selected. By convention the first segment of the path is a category that is used to group exceptions in the UI.
	Path []ExceptionPathSegment `json:"path,omitempty"`

	// Condition when a thrown exception should result in a break.
	BreakMode ExceptionBreakMode `json:"breakMode"`
} // struct ExceptionOptions



// Response to 'completions' request.
type CompletionsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The possible completions for .
		Targets []CompletionItem `json:"targets"`

	} `json:"body"`
} // struct CompletionsResponse



// SetExceptionBreakpoints request; value of command field is 'setExceptionBreakpoints'.
// The request configures the debuggers response to thrown exceptions. If an exception is configured to break, a StoppedEvent is fired (event type 'exception').
type SetExceptionBreakpointsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command"`

	Arguments SetExceptionBreakpointsArguments `json:"arguments"`
} // struct SetExceptionBreakpointsRequest



// Arguments for 'goto' request.
type GotoArguments struct {

	// The location where the debuggee will continue to run.
	TargetId int `json:"targetId"`

	// Set the goto target for this thread.
	ThreadId int `json:"threadId"`
} // struct GotoArguments



// Response to 'configurationDone' request. This is just an acknowledgement, so no body field is required.
type ConfigurationDoneResponse struct {
	// Response to a request.
	Response
} // struct ConfigurationDoneResponse



// ExceptionInfoRequest request; value of command field is 'exceptionInfo'.
// Retrieves the details of the exception that caused the StoppedEvent to be raised.
type ExceptionInfoRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `exceptionInfo`
	Command string `json:"command"`

	Arguments ExceptionInfoArguments `json:"arguments"`
} // struct ExceptionInfoRequest



// Response to 'next' request. This is just an acknowledgement, so no body field is required.
type NextResponse struct {
	// Response to a request.
	Response
} // struct NextResponse



// Source request; value of command field is 'source'.
// The request retrieves the source code for a given source reference.
type SourceRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `source`
	Command string `json:"command"`

	Arguments SourceArguments `json:"arguments"`
} // struct SourceRequest



// Response to 'stepBack' request. This is just an acknowledgement, so no body field is required.
type StepBackResponse struct {
	// Response to a request.
	Response
} // struct StepBackResponse



// A Thread
type Thread struct {

	// Unique identifier for the thread.
	Id int `json:"id"`

	// A name of the thread.
	Name string `json:"name"`
} // struct Thread



// Response to 'continue' request.
type ContinueResponse struct {
	// Response to a request.
	Response

	Body struct{

		// If true, the continue request has ignored the specified thread and continued all threads instead. If this attribute is missing a value of 'true' is assumed for backward compatibility.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`

	} `json:"body"`
} // struct ContinueResponse



// Base class of requests, responses, and events.
type ProtocolMessage struct {

	// Sequence number.
	Seq int `json:"seq"`

	// One of 'request', 'response', or 'event'.
	// 
	// POSSIBLE VALUES: ``
	Type string `json:"type"`
} // struct ProtocolMessage



// Arguments for 'setFunctionBreakpoints' request.
type SetFunctionBreakpointsArguments struct {

	// The function names of the breakpoints.
	Breakpoints []FunctionBreakpoint `json:"breakpoints"`
} // struct SetFunctionBreakpointsArguments



// Restart request; value of command field is 'restart'.
// Restarts a debug session. If the capability 'supportsRestartRequest' is missing or has the value false,
// the client will implement 'restart' by terminating the debug adapter first and then launching it anew.
// A debug adapter can override this default behaviour by implementing a restart request
// and setting the capability 'supportsRestartRequest' to true.
type RestartRequest struct {
	// A client or server-initiated request.
	Request

	Arguments RestartArguments `json:"arguments,omitempty"`

	// POSSIBLE VALUES: `restart`
	Command string `json:"command"`
} // struct RestartRequest



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

	// True if the module is considered 'user code' by a debugger that supports 'Just My Code'.
	IsUserCode bool `json:"isUserCode,omitempty"`

	// Logical full path to the symbol file. The exact definition is implementation defined.
	SymbolFilePath string `json:"symbolFilePath,omitempty"`

	// Unique identifier for the module.
	// 
	// POSSIBLE TYPES: `int`, `string`
	Id interface{} `json:"id"`

	// A name of the module.
	Name string `json:"name"`

	// User understandable description of if symbols were found for the module (ex: 'Symbols Loaded', 'Symbols not found', etc.
	SymbolStatus string `json:"symbolStatus,omitempty"`

	// Module created or modified.
	DateTimeStamp string `json:"dateTimeStamp,omitempty"`

	// Address range covered by this module.
	AddressRange string `json:"addressRange,omitempty"`

	// True if the module is optimized.
	IsOptimized bool `json:"isOptimized,omitempty"`

	// Version of Module.
	Version string `json:"version,omitempty"`
} // struct Module



// Scopes request; value of command field is 'scopes'.
// The request returns the variable scopes for a given stackframe ID.
type ScopesRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `scopes`
	Command string `json:"command"`

	Arguments ScopesArguments `json:"arguments"`
} // struct ScopesRequest



// Provides formatting information for a value.
type ValueFormat struct {

	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`
} // struct ValueFormat



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



// Response to 'restartFrame' request. This is just an acknowledgement, so no body field is required.
type RestartFrameResponse struct {
	// Response to a request.
	Response
} // struct RestartFrameResponse



// Response to 'threads' request.
type ThreadsResponse struct {
	// Response to a request.
	Response

	Body struct{

		// All threads.
		Threads []Thread `json:"threads"`

	} `json:"body"`
} // struct ThreadsResponse



// Arguments for 'stepOut' request.
type StepOutArguments struct {

	// Execute 'stepOut' for this thread.
	ThreadId int `json:"threadId"`
} // struct StepOutArguments



// setVariable request; value of command field is 'setVariable'.
// Set the variable with the given name in the variable container to a new value.
type SetVariableRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setVariable`
	Command string `json:"command"`

	Arguments SetVariableArguments `json:"arguments"`
} // struct SetVariableRequest



// Response to 'reverseContinue' request. This is just an acknowledgement, so no body field is required.
type ReverseContinueResponse struct {
	// Response to a request.
	Response
} // struct ReverseContinueResponse



// A Stackframe contains the source location.
type StackFrame struct {

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line int `json:"line"`

	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column int `json:"column"`

	// An optional end column of the range covered by the stack frame.
	EndColumn int `json:"endColumn,omitempty"`

	// The name of the stack frame, typically a method name.
	Name string `json:"name"`

	// The optional source of the frame.
	Source Source `json:"source,omitempty"`

	// The module associated with this frame, if any.
	// 
	// POSSIBLE TYPES: `int`, `string`
	ModuleId interface{} `json:"moduleId,omitempty"`

	// An optional hint for how to present this frame in the UI. A value of 'label' can be used to indicate that the frame is an artificial frame that is used as a visual label or separator.
	// 
	// POSSIBLE VALUES: `normal`, `label`
	PresentationHint string `json:"presentationHint,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id int `json:"id"`

	// An optional end line of the range covered by the stack frame.
	EndLine int `json:"endLine,omitempty"`
} // struct StackFrame



// Response to 'scopes' request.
type ScopesResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The scopes of the stackframe. If the array has length zero, there are no scopes available.
		Scopes []Scope `json:"scopes"`

	} `json:"body"`
} // struct ScopesResponse



// Response to 'stepIn' request. This is just an acknowledgement, so no body field is required.
type StepInResponse struct {
	// Response to a request.
	Response
} // struct StepInResponse



// Response to 'exceptionInfo' request.
type ExceptionInfoResponse struct {
	// Response to a request.
	Response

	Body struct{

		// ID of the exception that was thrown.
		ExceptionId string `json:"exceptionId"`

		// Descriptive text for the exception provided by the debug adapter.
		Description string `json:"description,omitempty"`

		// Mode that caused the exception notification to be raised.
		BreakMode ExceptionBreakMode `json:"breakMode"`

		// Detailed information about the exception.
		Details ExceptionDetails `json:"details,omitempty"`

	} `json:"body"`
} // struct ExceptionInfoResponse



// Response to 'disconnect' request. This is just an acknowledgement, so no body field is required.
type DisconnectResponse struct {
	// Response to a request.
	Response
} // struct DisconnectResponse



// Some predefined types for the CompletionItem. Please note that not all clients have specific icons for all of them.
// 
// POSSIBLE VALUES: `method`, `function`, `constructor`, `field`, `variable`, `class`, `interface`, `module`, `property`, `unit`, `value`, `enum`, `keyword`, `snippet`, `text`, `color`, `file`, `reference`, `customcolor`
type CompletionItemType string 



// Arguments for 'source' request.
type SourceArguments struct {

	// Specifies the source content to load. Either source.path or source.sourceReference must be specified.
	Source Source `json:"source,omitempty"`

	// The reference to the source. This is the same as source.sourceReference. This is provided for backward compatibility since old backends do not understand the 'source' attribute.
	SourceReference int `json:"sourceReference"`
} // struct SourceArguments



// A StepInTarget can be used in the 'stepIn' request and determines into which single target the stepIn request should step.
type StepInTarget struct {

	// Unique identifier for a stepIn target.
	Id int `json:"id"`

	// The name of the stepIn target (shown in the UI).
	Label string `json:"label"`
} // struct StepInTarget



// Arguments for 'disconnect' request.
// The disconnect request has no standardized attributes.
type DisconnectArguments struct {
} // struct DisconnectArguments



// Response to Initialize request.
type RunInTerminalResponse struct {
	// Response to a request.
	Response

	Body struct{

		// The process ID.
		ProcessId int64 `json:"processId,omitempty"`

	} `json:"body"`
} // struct RunInTerminalResponse



// A Variable is a name/value pair.
// Optionally a variable can have a 'type' that is shown if space permits or when hovering over the variable's name.
// An optional 'kind' is used to render additional properties of the variable, e.g. different icons can be used to indicate that a variable is public or private.
// If the value is structured (has children), a handle is provided to retrieve the children with the VariablesRequest.
// If the number of named or indexed children is large, the numbers should be returned via the optional 'namedVariables' and 'indexedVariables' attributes.
// The client can use this optional information to present the children in a paged UI and fetch them in chunks.
type Variable struct {

	// The variable's value. This can be a multi-line text, e.g. for a function the body of a function.
	Value string `json:"value"`

	// The type of the variable's value. Typically shown in the UI when hovering over the value.
	Type string `json:"type,omitempty"`

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
} // struct Variable



// Arguments for 'setVariable' request.
type SetVariableArguments struct {

	// The reference of the variable container.
	VariablesReference int `json:"variablesReference"`

	// The name of the variable.
	Name string `json:"name"`

	// The value of the variable.
	Value string `json:"value"`

	// Specifies details on how to format the response value.
	Format ValueFormat `json:"format,omitempty"`
} // struct SetVariableArguments



// Arguments for 'modules' request.
type ModulesArguments struct {

	// The index of the first module to return; if omitted modules start at 0.
	StartModule int `json:"startModule,omitempty"`

	// The number of modules to return. If moduleCount is not specified or 0, all modules are returned.
	ModuleCount int `json:"moduleCount,omitempty"`
} // struct ModulesArguments



// Arguments for 'attach' request.
// The attach request has no standardized attributes.
type AttachRequestArguments struct {
} // struct AttachRequestArguments



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