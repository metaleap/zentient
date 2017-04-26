# zdbgvscp
--
    import "github.com/metaleap/zentient/zdbg-vsc/proto"

### VS Code Debug Protocol


A json schema for the VS Code Debug Protocol

Package codegen'd from
github.com/metaleap/zentient/zdbg-vsc/_notes_misc_etc/vscdbgprotocol.json with
github.com/metaleap/zentient/zdbg-vsc-proto-gen

## Usage

```go
var OnAttachRequest func(*AttachRequest, *AttachResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `AttachRequest`) to
further populate the given `AttachResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnCompletionsRequest func(*CompletionsRequest, *CompletionsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `CompletionsRequest`)
to further populate the given `CompletionsResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnConfigurationDoneRequest func(*ConfigurationDoneRequest, *ConfigurationDoneResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`ConfigurationDoneRequest`) to further populate the given
`ConfigurationDoneResponse` before returning it to its caller (in addition to
this handler's returned `error`).

```go
var OnContinueRequest func(*ContinueRequest, *ContinueResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `ContinueRequest`) to
further populate the given `ContinueResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnDisconnectRequest func(*DisconnectRequest, *DisconnectResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `DisconnectRequest`)
to further populate the given `DisconnectResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnEvaluateRequest func(*EvaluateRequest, *EvaluateResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `EvaluateRequest`) to
further populate the given `EvaluateResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnExceptionInfoRequest func(*ExceptionInfoRequest, *ExceptionInfoResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`ExceptionInfoRequest`) to further populate the given `ExceptionInfoResponse`
before returning it to its caller (in addition to this handler's returned
`error`).

```go
var OnGotoRequest func(*GotoRequest, *GotoResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `GotoRequest`) to
further populate the given `GotoResponse` before returning it to its caller (in
addition to this handler's returned `error`).

```go
var OnGotoTargetsRequest func(*GotoTargetsRequest, *GotoTargetsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `GotoTargetsRequest`)
to further populate the given `GotoTargetsResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnInitializeRequest func(*InitializeRequest, *InitializeResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `InitializeRequest`)
to further populate the given `InitializeResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnLaunchRequest func(*LaunchRequest, *LaunchResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `LaunchRequest`) to
further populate the given `LaunchResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnModulesRequest func(*ModulesRequest, *ModulesResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `ModulesRequest`) to
further populate the given `ModulesResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnNextRequest func(*NextRequest, *NextResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `NextRequest`) to
further populate the given `NextResponse` before returning it to its caller (in
addition to this handler's returned `error`).

```go
var OnPauseRequest func(*PauseRequest, *PauseResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `PauseRequest`) to
further populate the given `PauseResponse` before returning it to its caller (in
addition to this handler's returned `error`).

```go
var OnRestartFrameRequest func(*RestartFrameRequest, *RestartFrameResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `RestartFrameRequest`)
to further populate the given `RestartFrameResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnRestartRequest func(*RestartRequest, *RestartResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `RestartRequest`) to
further populate the given `RestartResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnReverseContinueRequest func(*ReverseContinueRequest, *ReverseContinueResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`ReverseContinueRequest`) to further populate the given
`ReverseContinueResponse` before returning it to its caller (in addition to this
handler's returned `error`).

```go
var OnRunInTerminalRequest func(*RunInTerminalRequest, *RunInTerminalResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`RunInTerminalRequest`) to further populate the given `RunInTerminalResponse`
before returning it to its caller (in addition to this handler's returned
`error`).

```go
var OnScopesRequest func(*ScopesRequest, *ScopesResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `ScopesRequest`) to
further populate the given `ScopesResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnSetBreakpointsRequest func(*SetBreakpointsRequest, *SetBreakpointsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`SetBreakpointsRequest`) to further populate the given `SetBreakpointsResponse`
before returning it to its caller (in addition to this handler's returned
`error`).

```go
var OnSetExceptionBreakpointsRequest func(*SetExceptionBreakpointsRequest, *SetExceptionBreakpointsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`SetExceptionBreakpointsRequest`) to further populate the given
`SetExceptionBreakpointsResponse` before returning it to its caller (in addition
to this handler's returned `error`).

```go
var OnSetFunctionBreakpointsRequest func(*SetFunctionBreakpointsRequest, *SetFunctionBreakpointsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`SetFunctionBreakpointsRequest`) to further populate the given
`SetFunctionBreakpointsResponse` before returning it to its caller (in addition
to this handler's returned `error`).

```go
var OnSetVariableRequest func(*SetVariableRequest, *SetVariableResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `SetVariableRequest`)
to further populate the given `SetVariableResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnSourceRequest func(*SourceRequest, *SourceResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `SourceRequest`) to
further populate the given `SourceResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnStackTraceRequest func(*StackTraceRequest, *StackTraceResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `StackTraceRequest`)
to further populate the given `StackTraceResponse` before returning it to its
caller (in addition to this handler's returned `error`).

```go
var OnStepBackRequest func(*StepBackRequest, *StepBackResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `StepBackRequest`) to
further populate the given `StepBackResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnStepInRequest func(*StepInRequest, *StepInResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `StepInRequest`) to
further populate the given `StepInResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnStepInTargetsRequest func(*StepInTargetsRequest, *StepInTargetsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given
`StepInTargetsRequest`) to further populate the given `StepInTargetsResponse`
before returning it to its caller (in addition to this handler's returned
`error`).

```go
var OnStepOutRequest func(*StepOutRequest, *StepOutResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `StepOutRequest`) to
further populate the given `StepOutResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnThreadsRequest func(*ThreadsRequest, *ThreadsResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `ThreadsRequest`) to
further populate the given `ThreadsResponse` before returning it to its caller
(in addition to this handler's returned `error`).

```go
var OnVariablesRequest func(*VariablesRequest, *VariablesResponse) error
```
Called by `HandleRequest` (after it unmarshaled the given `VariablesRequest`) to
further populate the given `VariablesResponse` before returning it to its caller
(in addition to this handler's returned `error`).

#### func  HandleRequest

```go
func HandleRequest(inRequest interface{}, initNewResponse func(*Request, *Response)) (outResponse interface{}, baseResponse *Response, err error)
```
If a type-switch on `inRequest` succeeds, `outResponse` points to a
`Response`-based `struct` value containing the `Response` initialized by the
specified `initNewResponse` and further populated by the `OnFooRequest` handler
corresponding to the concrete type of `inRequest` (if any). The only `err`
returned, if any, is that returned by the specialized `OnFooRequest` handler.

#### func  TryUnmarshalEvent

```go
func TryUnmarshalEvent(js string) (ptr interface{}, err error)
```
TryUnmarshalEvent attempts to unmarshal JSON string `js` (if it starts with a
`{` and ends with a `}`) into a `struct` based on `Event` as follows:

If `js` contains `"event":"initialized"`, attempts to unmarshal into a new
`InitializedEvent`.

If `js` contains `"event":"continued"`, attempts to unmarshal into a new
`ContinuedEvent`.

If `js` contains `"event":"output"`, attempts to unmarshal into a new
`OutputEvent`.

If `js` contains `"event":"exited"`, attempts to unmarshal into a new
`ExitedEvent`.

If `js` contains `"event":"terminated"`, attempts to unmarshal into a new
`TerminatedEvent`.

If `js` contains `"event":"thread"`, attempts to unmarshal into a new
`ThreadEvent`.

If `js` contains `"event":"module"`, attempts to unmarshal into a new
`ModuleEvent`.

If `js` contains `"event":"breakpoint"`, attempts to unmarshal into a new
`BreakpointEvent`.

If `js` contains `"event":"stopped"`, attempts to unmarshal into a new
`StoppedEvent`.

Otherwise, `err`'s message will be: `Event: encountered unknown JSON value for
event: ` followed by the `event` value encountered.

In general: the `err` returned may be either `nil`, the above message, or an
`encoding/json.Unmarshal()` return value. `ptr` will be a pointer to the
unmarshaled `struct` value if that succeeded, else `nil`. Both `err` and `ptr`
will be `nil` if `js` doesn't: start with `{` and end with `}` and contain
`"event":"` followed by a subsequent `"`.

#### func  TryUnmarshalProtocolMessage

```go
func TryUnmarshalProtocolMessage(js string) (ptr interface{}, err error)
```
TryUnmarshalProtocolMessage attempts to unmarshal JSON string `js` (if it starts
with a `{` and ends with a `}`) into a `struct` based on `ProtocolMessage` as
follows:

If `js` contains `"type":"response"`, attempts to unmarshal via
`TryUnmarshalResponse`

If `js` contains `"type":"event"`, attempts to unmarshal via `TryUnmarshalEvent`

If `js` contains `"type":"request"`, attempts to unmarshal via
`TryUnmarshalRequest`

Otherwise, `err`'s message will be: `ProtocolMessage: encountered unknown JSON
value for type: ` followed by the `type` value encountered.

In general: the `err` returned may be either `nil`, the above message, or an
`encoding/json.Unmarshal()` return value. `ptr` will be a pointer to the
unmarshaled `struct` value if that succeeded, else `nil`. Both `err` and `ptr`
will be `nil` if `js` doesn't: start with `{` and end with `}` and contain
`"type":"` followed by a subsequent `"`.

#### func  TryUnmarshalRequest

```go
func TryUnmarshalRequest(js string) (ptr interface{}, err error)
```
TryUnmarshalRequest attempts to unmarshal JSON string `js` (if it starts with a
`{` and ends with a `}`) into a `struct` based on `Request` as follows:

If `js` contains `"command":"goto"`, attempts to unmarshal into a new
`GotoRequest`.

If `js` contains `"command":"scopes"`, attempts to unmarshal into a new
`ScopesRequest`.

If `js` contains `"command":"restartFrame"`, attempts to unmarshal into a new
`RestartFrameRequest`.

If `js` contains `"command":"stepIn"`, attempts to unmarshal into a new
`StepInRequest`.

If `js` contains `"command":"setVariable"`, attempts to unmarshal into a new
`SetVariableRequest`.

If `js` contains `"command":"configurationDone"`, attempts to unmarshal into a
new `ConfigurationDoneRequest`.

If `js` contains `"command":"stepOut"`, attempts to unmarshal into a new
`StepOutRequest`.

If `js` contains `"command":"modules"`, attempts to unmarshal into a new
`ModulesRequest`.

If `js` contains `"command":"next"`, attempts to unmarshal into a new
`NextRequest`.

If `js` contains `"command":"variables"`, attempts to unmarshal into a new
`VariablesRequest`.

If `js` contains `"command":"setFunctionBreakpoints"`, attempts to unmarshal
into a new `SetFunctionBreakpointsRequest`.

If `js` contains `"command":"initialize"`, attempts to unmarshal into a new
`InitializeRequest`.

If `js` contains `"command":"continue"`, attempts to unmarshal into a new
`ContinueRequest`.

If `js` contains `"command":"reverseContinue"`, attempts to unmarshal into a new
`ReverseContinueRequest`.

If `js` contains `"command":"disconnect"`, attempts to unmarshal into a new
`DisconnectRequest`.

If `js` contains `"command":"stackTrace"`, attempts to unmarshal into a new
`StackTraceRequest`.

If `js` contains `"command":"attach"`, attempts to unmarshal into a new
`AttachRequest`.

If `js` contains `"command":"launch"`, attempts to unmarshal into a new
`LaunchRequest`.

If `js` contains `"command":"source"`, attempts to unmarshal into a new
`SourceRequest`.

If `js` contains `"command":"evaluate"`, attempts to unmarshal into a new
`EvaluateRequest`.

If `js` contains `"command":"setBreakpoints"`, attempts to unmarshal into a new
`SetBreakpointsRequest`.

If `js` contains `"command":"gotoTargets"`, attempts to unmarshal into a new
`GotoTargetsRequest`.

If `js` contains `"command":"stepInTargets"`, attempts to unmarshal into a new
`StepInTargetsRequest`.

If `js` contains `"command":"runInTerminal"`, attempts to unmarshal into a new
`RunInTerminalRequest`.

If `js` contains `"command":"pause"`, attempts to unmarshal into a new
`PauseRequest`.

If `js` contains `"command":"exceptionInfo"`, attempts to unmarshal into a new
`ExceptionInfoRequest`.

If `js` contains `"command":"threads"`, attempts to unmarshal into a new
`ThreadsRequest`.

If `js` contains `"command":"setExceptionBreakpoints"`, attempts to unmarshal
into a new `SetExceptionBreakpointsRequest`.

If `js` contains `"command":"completions"`, attempts to unmarshal into a new
`CompletionsRequest`.

If `js` contains `"command":"restart"`, attempts to unmarshal into a new
`RestartRequest`.

If `js` contains `"command":"stepBack"`, attempts to unmarshal into a new
`StepBackRequest`.

Otherwise, `err`'s message will be: `Request: encountered unknown JSON value for
command: ` followed by the `command` value encountered.

In general: the `err` returned may be either `nil`, the above message, or an
`encoding/json.Unmarshal()` return value. `ptr` will be a pointer to the
unmarshaled `struct` value if that succeeded, else `nil`. Both `err` and `ptr`
will be `nil` if `js` doesn't: start with `{` and end with `}` and contain
`"command":"` followed by a subsequent `"`.

#### func  TryUnmarshalResponse

```go
func TryUnmarshalResponse(js string) (ptr interface{}, err error)
```
TryUnmarshalResponse attempts to unmarshal JSON string `js` (if it starts with a
`{` and ends with a `}`) into a `struct` based on `Response` as follows:

If `js` contains `"command":"configurationDone"`, attempts to unmarshal into a
new `ConfigurationDoneResponse`.

If `js` contains `"command":"gotoTargets"`, attempts to unmarshal into a new
`GotoTargetsResponse`.

If `js` contains `"command":"continue"`, attempts to unmarshal into a new
`ContinueResponse`.

If `js` contains `"command":"threads"`, attempts to unmarshal into a new
`ThreadsResponse`.

If `js` contains `"command":"disconnect"`, attempts to unmarshal into a new
`DisconnectResponse`.

If `js` contains `"command":"setFunctionBreakpoints"`, attempts to unmarshal
into a new `SetFunctionBreakpointsResponse`.

If `js` contains `"command":"stepBack"`, attempts to unmarshal into a new
`StepBackResponse`.

If `js` contains `"command":"setVariable"`, attempts to unmarshal into a new
`SetVariableResponse`.

If `js` contains `"command":"variables"`, attempts to unmarshal into a new
`VariablesResponse`.

If `js` contains `"command":"setExceptionBreakpoints"`, attempts to unmarshal
into a new `SetExceptionBreakpointsResponse`.

If `js` contains `"command":"attach"`, attempts to unmarshal into a new
`AttachResponse`.

If `js` contains `"command":"runInTerminal"`, attempts to unmarshal into a new
`RunInTerminalResponse`.

If `js` contains `"command":"goto"`, attempts to unmarshal into a new
`GotoResponse`.

If `js` contains `"command":"stepInTargets"`, attempts to unmarshal into a new
`StepInTargetsResponse`.

If `js` contains `"command":"modules"`, attempts to unmarshal into a new
`ModulesResponse`.

If `js` contains `"command":"source"`, attempts to unmarshal into a new
`SourceResponse`.

If `js` contains `"command":"next"`, attempts to unmarshal into a new
`NextResponse`.

If `js` contains `"command":"initialize"`, attempts to unmarshal into a new
`InitializeResponse`.

If `js` contains `"command":"completions"`, attempts to unmarshal into a new
`CompletionsResponse`.

If `js` contains `"command":"setBreakpoints"`, attempts to unmarshal into a new
`SetBreakpointsResponse`.

If `js` contains `"command":"restartFrame"`, attempts to unmarshal into a new
`RestartFrameResponse`.

If `js` contains `"command":"launch"`, attempts to unmarshal into a new
`LaunchResponse`.

If `js` contains `"command":"reverseContinue"`, attempts to unmarshal into a new
`ReverseContinueResponse`.

If `js` contains `"command":"exceptionInfo"`, attempts to unmarshal into a new
`ExceptionInfoResponse`.

If `js` contains `"command":"stepIn"`, attempts to unmarshal into a new
`StepInResponse`.

If `js` contains `"command":"evaluate"`, attempts to unmarshal into a new
`EvaluateResponse`.

If `js` contains `"command":"stepOut"`, attempts to unmarshal into a new
`StepOutResponse`.

If `js` contains `"command":"restart"`, attempts to unmarshal into a new
`RestartResponse`.

If `js` contains `"command":"stackTrace"`, attempts to unmarshal into a new
`StackTraceResponse`.

If `js` contains `"command":"pause"`, attempts to unmarshal into a new
`PauseResponse`.

If `js` contains `"command":"scopes"`, attempts to unmarshal into a new
`ScopesResponse`.

Otherwise, `err`'s message will be: `Response: encountered unknown JSON value
for command: ` followed by the `command` value encountered.

In general: the `err` returned may be either `nil`, the above message, or an
`encoding/json.Unmarshal()` return value. `ptr` will be a pointer to the
unmarshaled `struct` value if that succeeded, else `nil`. Both `err` and `ptr`
will be `nil` if `js` doesn't: start with `{` and end with `}` and contain
`"command":"` followed by a subsequent `"`.

#### type AttachRequest

```go
type AttachRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `attach`
	Command string `json:"command"`

	Arguments AttachRequestArguments `json:"arguments"`
} // struct AttachRequest

```

Attach request; value of command field is 'attach'.

#### func  NewAttachRequest

```go
func NewAttachRequest() *AttachRequest
```
Returns a new `AttachRequest` with the following fields set: `Command`, `Type`

#### type AttachRequestArguments

```go
type AttachRequestArguments struct {
} // struct AttachRequestArguments

```

Arguments for 'attach' request. The attach request has no standardized
attributes.

#### type AttachResponse

```go
type AttachResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `attach`
	Command string `json:"command,omitempty"`
} // struct AttachResponse

```

Response to 'attach' request. This is just an acknowledgement, so no body field
is required.

#### func  NewAttachResponse

```go
func NewAttachResponse() *AttachResponse
```
Returns a new `AttachResponse` with the following fields set: `Command`, `Type`

#### type Breakpoint

```go
type Breakpoint struct {

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

	// The start line of the actual range covered by the breakpoint.
	Line int `json:"line,omitempty"`

	// An optional start column of the actual range covered by the breakpoint.
	Column int `json:"column,omitempty"`
} // struct Breakpoint

```

Information about a Breakpoint created in setBreakpoints or
setFunctionBreakpoints.

#### type BreakpointEvent

```go
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

```

Event message for 'breakpoint' event type. The event indicates that some
information about a breakpoint has changed.

#### func  NewBreakpointEvent

```go
func NewBreakpointEvent() *BreakpointEvent
```
Returns a new `BreakpointEvent` with the following fields set: `Event_`, `Type`

#### type Capabilities

```go
type Capabilities struct {

	// Available filters or options for the setExceptionBreakpoints request.
	ExceptionBreakpointFilters []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`

	// The debug adapter supports a 'format' attribute on the stackTraceRequest, variablesRequest, and evaluateRequest.
	SupportsValueFormattingOptions bool `json:"supportsValueFormattingOptions,omitempty"`

	// The debug adapter supports the configurationDoneRequest.
	SupportsConfigurationDoneRequest bool `json:"supportsConfigurationDoneRequest,omitempty"`

	// The debug adapter supports function breakpoints.
	SupportsFunctionBreakpoints bool `json:"supportsFunctionBreakpoints,omitempty"`

	// The debug adapter supports stepping back via the stepBack and reverseContinue requests.
	SupportsStepBack bool `json:"supportsStepBack,omitempty"`

	// The debug adapter supports the modules request.
	SupportsModulesRequest bool `json:"supportsModulesRequest,omitempty"`

	// The set of additional module information exposed by the debug adapter.
	AdditionalModuleColumns []ColumnDescriptor `json:"additionalModuleColumns,omitempty"`

	// The debug adapter supports breakpoints that break execution after a specified number of hits.
	SupportsHitConditionalBreakpoints bool `json:"supportsHitConditionalBreakpoints,omitempty"`

	// The debug adapter supports a (side effect free) evaluate request for data hovers.
	SupportsEvaluateForHovers bool `json:"supportsEvaluateForHovers,omitempty"`

	// The debug adapter supports setting a variable to a value.
	SupportsSetVariable bool `json:"supportsSetVariable,omitempty"`

	// The debug adapter supports restarting a frame.
	SupportsRestartFrame bool `json:"supportsRestartFrame,omitempty"`

	// The debug adapter supports the gotoTargetsRequest.
	SupportsGotoTargetsRequest bool `json:"supportsGotoTargetsRequest,omitempty"`

	// The debug adapter supports the stepInTargetsRequest.
	SupportsStepInTargetsRequest bool `json:"supportsStepInTargetsRequest,omitempty"`

	// The debug adapter supports the RestartRequest. In this case a client should not implement 'restart' by terminating and relaunching the adapter but by calling the RestartRequest.
	SupportsRestartRequest bool `json:"supportsRestartRequest,omitempty"`

	// The debug adapter supports the exceptionInfo request.
	SupportsExceptionInfoRequest bool `json:"supportsExceptionInfoRequest,omitempty"`

	// The debug adapter supports conditional breakpoints.
	SupportsConditionalBreakpoints bool `json:"supportsConditionalBreakpoints,omitempty"`

	// The debug adapter supports the completionsRequest.
	SupportsCompletionsRequest bool `json:"supportsCompletionsRequest,omitempty"`

	// Checksum algorithms supported by the debug adapter.
	SupportedChecksumAlgorithms []ChecksumAlgorithm `json:"supportedChecksumAlgorithms,omitempty"`

	// The debug adapter supports 'exceptionOptions' on the setExceptionBreakpoints request.
	SupportsExceptionOptions bool `json:"supportsExceptionOptions,omitempty"`

	// The debug adapter supports the 'terminateDebuggee' attribute on the 'disconnect' request.
	SupportTerminateDebuggee bool `json:"supportTerminateDebuggee,omitempty"`
} // struct Capabilities

```

Information about the capabilities of a debug adapter.

#### type Checksum

```go
type Checksum struct {

	// The algorithm used to calculate this checksum.
	Algorithm ChecksumAlgorithm `json:"algorithm"`

	// Value of the checksum.
	Checksum string `json:"checksum"`
} // struct Checksum

```

The checksum of an item calculated by the specified algorithm.

#### type ChecksumAlgorithm

```go
type ChecksumAlgorithm string
```

Names of checksum algorithms that may be supported by a debug adapter.

POSSIBLE VALUES: `MD5`, `SHA1`, `SHA256`, `timestamp`

#### type ColumnDescriptor

```go
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

```

A ColumnDescriptor specifies what module attribute to show in a column of the
ModulesView, how to format it, and what the column's label should be. It is only
used if the underlying UI actually supports this level of customization.

#### type CompletionItem

```go
type CompletionItem struct {

	// This value determines the location (in the CompletionsRequest's 'text' attribute) where the completion text is added.
	// If missing the text is added at the location specified by the CompletionsRequest's 'column' attribute.
	Start int `json:"start,omitempty"`

	// This value determines how many characters are overwritten by the completion text.
	// If missing the value 0 is assumed which results in the completion text being inserted.
	Length int `json:"length,omitempty"`

	// The label of this completion item. By default this is also the text that is inserted when selecting this completion.
	Label string `json:"label"`

	// If text is not falsy then it is inserted instead of the label.
	Text string `json:"text,omitempty"`

	// The item's type. Typically the client uses this information to render the item in the UI with an icon.
	Type CompletionItemType `json:"type,omitempty"`
} // struct CompletionItem

```

CompletionItems are the suggestions returned from the CompletionsRequest.

#### type CompletionItemType

```go
type CompletionItemType string
```

Some predefined types for the CompletionItem. Please note that not all clients
have specific icons for all of them.

POSSIBLE VALUES: `method`, `function`, `constructor`, `field`, `variable`,
`class`, `interface`, `module`, `property`, `unit`, `value`, `enum`, `keyword`,
`snippet`, `text`, `color`, `file`, `reference`, `customcolor`

#### type CompletionsArguments

```go
type CompletionsArguments struct {

	// One or more source lines. Typically this is the text a user has typed into the debug console before he asked for completion.
	Text string `json:"text"`

	// The character position for which to determine the completion proposals.
	Column int `json:"column"`

	// An optional line for which to determine the completion proposals. If missing the first line of the text is assumed.
	Line int `json:"line,omitempty"`

	// Returns completions in the scope of this stack frame. If not specified, the completions are returned for the global scope.
	FrameId int `json:"frameId,omitempty"`
} // struct CompletionsArguments

```

Arguments for 'completions' request.

#### type CompletionsRequest

```go
type CompletionsRequest struct {
	// A client or server-initiated request.
	Request

	Arguments CompletionsArguments `json:"arguments"`

	// POSSIBLE VALUES: `completions`
	Command string `json:"command"`
} // struct CompletionsRequest

```

CompletionsRequest request; value of command field is 'completions'. Returns a
list of possible completions for a given caret position and text. The
CompletionsRequest may only be called if the 'supportsCompletionsRequest'
capability exists and is true.

#### func  NewCompletionsRequest

```go
func NewCompletionsRequest() *CompletionsRequest
```
Returns a new `CompletionsRequest` with the following fields set: `Command`,
`Type`

#### type CompletionsResponse

```go
type CompletionsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible completions for .
		Targets []CompletionItem `json:"targets"`
	} `json:"body"`

	// POSSIBLE VALUES: `completions`
	Command string `json:"command,omitempty"`
} // struct CompletionsResponse

```

Response to 'completions' request.

#### func  NewCompletionsResponse

```go
func NewCompletionsResponse() *CompletionsResponse
```
Returns a new `CompletionsResponse` with the following fields set: `Command`,
`Type`

#### type ConfigurationDoneArguments

```go
type ConfigurationDoneArguments struct {
} // struct ConfigurationDoneArguments

```

Arguments for 'configurationDone' request. The configurationDone request has no
standardized attributes.

#### type ConfigurationDoneRequest

```go
type ConfigurationDoneRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `configurationDone`
	Command string `json:"command"`

	Arguments ConfigurationDoneArguments `json:"arguments,omitempty"`
} // struct ConfigurationDoneRequest

```

ConfigurationDone request; value of command field is 'configurationDone'. The
client of the debug protocol must send this request at the end of the sequence
of configuration requests (which was started by the InitializedEvent).

#### func  NewConfigurationDoneRequest

```go
func NewConfigurationDoneRequest() *ConfigurationDoneRequest
```
Returns a new `ConfigurationDoneRequest` with the following fields set:
`Command`, `Type`

#### type ConfigurationDoneResponse

```go
type ConfigurationDoneResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `configurationDone`
	Command string `json:"command,omitempty"`
} // struct ConfigurationDoneResponse

```

Response to 'configurationDone' request. This is just an acknowledgement, so no
body field is required.

#### func  NewConfigurationDoneResponse

```go
func NewConfigurationDoneResponse() *ConfigurationDoneResponse
```
Returns a new `ConfigurationDoneResponse` with the following fields set:
`Command`, `Type`

#### type ContinueArguments

```go
type ContinueArguments struct {

	// Continue execution for the specified thread (if possible). If the backend cannot continue on a single thread but will continue on all threads, it should set the allThreadsContinued attribute in the response to true.
	ThreadId int `json:"threadId"`
} // struct ContinueArguments

```

Arguments for 'continue' request.

#### type ContinueRequest

```go
type ContinueRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `continue`
	Command string `json:"command"`

	Arguments ContinueArguments `json:"arguments"`
} // struct ContinueRequest

```

Continue request; value of command field is 'continue'. The request starts the
debuggee to run again.

#### func  NewContinueRequest

```go
func NewContinueRequest() *ContinueRequest
```
Returns a new `ContinueRequest` with the following fields set: `Command`, `Type`

#### type ContinueResponse

```go
type ContinueResponse struct {
	// Response to a request.
	Response

	Body struct {

		// If true, the continue request has ignored the specified thread and continued all threads instead. If this attribute is missing a value of 'true' is assumed for backward compatibility.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`
	} `json:"body"`

	// POSSIBLE VALUES: `continue`
	Command string `json:"command,omitempty"`
} // struct ContinueResponse

```

Response to 'continue' request.

#### func  NewContinueResponse

```go
func NewContinueResponse() *ContinueResponse
```
Returns a new `ContinueResponse` with the following fields set: `Command`,
`Type`

#### type ContinuedEvent

```go
type ContinuedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `continued`
	Event_ string `json:"event"`

	Body struct {

		// The thread which was continued.
		ThreadId int `json:"threadId"`

		// If allThreadsContinued is true, a debug adapter can announce that all threads have continued.
		AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`
	} `json:"body"`
} // struct ContinuedEvent

```

Event message for 'continued' event type. The event indicates that the execution
of the debuggee has continued. Please note: a debug adapter is not expected to
send this event in response to a request that implies that execution continues,
e.g. 'launch' or 'continue'. It is only necessary to send a ContinuedEvent if
there was no previous request that implied this.

#### func  NewContinuedEvent

```go
func NewContinuedEvent() *ContinuedEvent
```
Returns a new `ContinuedEvent` with the following fields set: `Event_`, `Type`

#### type DisconnectArguments

```go
type DisconnectArguments struct {

	// restart
	Restart bool `json:"restart,omitempty"`

	// Indicates whether the debuggee should be terminated when the debugger is disconnected.
	// If unspecified, the debug adapter is free to do whatever it thinks is best.
	// A client can only rely on this attribute being properly honored if a debug adapter returns true for the 'supportTerminateDebuggee' capability.
	TerminateDebuggee bool `json:"terminateDebuggee,omitempty"`
} // struct DisconnectArguments

```

Arguments for 'disconnect' request.

#### type DisconnectRequest

```go
type DisconnectRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `disconnect`
	Command string `json:"command"`

	Arguments DisconnectArguments `json:"arguments,omitempty"`
} // struct DisconnectRequest

```

Disconnect request; value of command field is 'disconnect'.

#### func  NewDisconnectRequest

```go
func NewDisconnectRequest() *DisconnectRequest
```
Returns a new `DisconnectRequest` with the following fields set: `Command`,
`Type`

#### type DisconnectResponse

```go
type DisconnectResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `disconnect`
	Command string `json:"command,omitempty"`
} // struct DisconnectResponse

```

Response to 'disconnect' request. This is just an acknowledgement, so no body
field is required.

#### func  NewDisconnectResponse

```go
func NewDisconnectResponse() *DisconnectResponse
```
Returns a new `DisconnectResponse` with the following fields set: `Command`,
`Type`

#### type ErrorResponse

```go
type ErrorResponse struct {
	// Response to a request.
	Response

	Body struct {

		// An optional, structured error message.
		Error Message `json:"error,omitempty"`
	} `json:"body"`
} // struct ErrorResponse

```

On error that is whenever 'success' is false, the body can provide more details.

#### func  NewErrorResponse

```go
func NewErrorResponse() *ErrorResponse
```
Returns a new `ErrorResponse` with the following fields set: `Type`

#### type EvaluateArguments

```go
type EvaluateArguments struct {

	// The expression to evaluate.
	Expression string `json:"expression"`

	// Evaluate the expression in the scope of this stack frame. If not specified, the expression is evaluated in the global scope.
	FrameId int `json:"frameId,omitempty"`

	// The context in which the evaluate request is run. Possible values are 'watch' if evaluate is run in a watch, 'repl' if run from the REPL console, or 'hover' if run from a data hover.
	//
	// POSSIBLE VALUES: `watch`, `repl`, `hover`
	Context string `json:"context,omitempty"`

	// Specifies details on how to format the Evaluate result.
	Format ValueFormat `json:"format,omitempty"`
} // struct EvaluateArguments

```

Arguments for 'evaluate' request.

#### type EvaluateRequest

```go
type EvaluateRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `evaluate`
	Command string `json:"command"`

	Arguments EvaluateArguments `json:"arguments"`
} // struct EvaluateRequest

```

Evaluate request; value of command field is 'evaluate'. Evaluates the given
expression in the context of the top most stack frame. The expression has access
to any variables and arguments that are in scope.

#### func  NewEvaluateRequest

```go
func NewEvaluateRequest() *EvaluateRequest
```
Returns a new `EvaluateRequest` with the following fields set: `Command`, `Type`

#### type EvaluateResponse

```go
type EvaluateResponse struct {
	// Response to a request.
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

		// The result of the evaluate request.
		Result string `json:"result"`

		// The optional type of the evaluate result.
		Type string `json:"type,omitempty"`
	} `json:"body"`
} // struct EvaluateResponse

```

Response to 'evaluate' request.

#### func  NewEvaluateResponse

```go
func NewEvaluateResponse() *EvaluateResponse
```
Returns a new `EvaluateResponse` with the following fields set: `Command`,
`Type`

#### type Event

```go
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

```

Server-initiated event.

#### func  BaseEvent

```go
func BaseEvent(someEvent interface{}) (baseEvent *Event)
```

#### type ExceptionBreakMode

```go
type ExceptionBreakMode string
```

This enumeration defines all possible conditions when a thrown exception should
result in a break. never: never breaks, always: always breaks, unhandled: breaks
when excpetion unhandled, userUnhandled: breaks if the exception is not handled
by user code.

POSSIBLE VALUES: `never`, `always`, `unhandled`, `userUnhandled`

#### type ExceptionBreakpointsFilter

```go
type ExceptionBreakpointsFilter struct {

	// The internal ID of the filter. This value is passed to the setExceptionBreakpoints request.
	Filter string `json:"filter"`

	// The name of the filter. This will be shown in the UI.
	Label string `json:"label"`

	// Initial value of the filter. If not specified a value 'false' is assumed.
	Default bool `json:"default,omitempty"`
} // struct ExceptionBreakpointsFilter

```

An ExceptionBreakpointsFilter is shown in the UI as an option for configuring
how exceptions are dealt with.

#### type ExceptionDetails

```go
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

```

Detailed information about an exception that has occurred.

#### type ExceptionInfoArguments

```go
type ExceptionInfoArguments struct {

	// Thread for which exception information should be retrieved.
	ThreadId int `json:"threadId"`
} // struct ExceptionInfoArguments

```

Arguments for 'exceptionInfo' request.

#### type ExceptionInfoRequest

```go
type ExceptionInfoRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `exceptionInfo`
	Command string `json:"command"`

	Arguments ExceptionInfoArguments `json:"arguments"`
} // struct ExceptionInfoRequest

```

ExceptionInfoRequest request; value of command field is 'exceptionInfo'.
Retrieves the details of the exception that caused the StoppedEvent to be
raised.

#### func  NewExceptionInfoRequest

```go
func NewExceptionInfoRequest() *ExceptionInfoRequest
```
Returns a new `ExceptionInfoRequest` with the following fields set: `Command`,
`Type`

#### type ExceptionInfoResponse

```go
type ExceptionInfoResponse struct {
	// Response to a request.
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

```

Response to 'exceptionInfo' request.

#### func  NewExceptionInfoResponse

```go
func NewExceptionInfoResponse() *ExceptionInfoResponse
```
Returns a new `ExceptionInfoResponse` with the following fields set: `Command`,
`Type`

#### type ExceptionOptions

```go
type ExceptionOptions struct {

	// A path that selects a single or multiple exceptions in a tree. If 'path' is missing, the whole tree is selected. By convention the first segment of the path is a category that is used to group exceptions in the UI.
	Path []ExceptionPathSegment `json:"path,omitempty"`

	// Condition when a thrown exception should result in a break.
	BreakMode ExceptionBreakMode `json:"breakMode"`
} // struct ExceptionOptions

```

An ExceptionOptions assigns configuration options to a set of exceptions.

#### type ExceptionPathSegment

```go
type ExceptionPathSegment struct {

	// If false or missing this segment matches the names provided, otherwise it matches anything except the names provided.
	Negate bool `json:"negate,omitempty"`

	// Depending on the value of 'negate' the names that should match or not match.
	Names []string `json:"names"`
} // struct ExceptionPathSegment

```

An ExceptionPathSegment represents a segment in a path that is used to match
leafs or nodes in a tree of exceptions. If a segment consists of more than one
name, it matches the names provided if 'negate' is false or missing or it
matches anything except the names provided if 'negate' is true.

#### type ExitedEvent

```go
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

```

Event message for 'exited' event type. The event indicates that the debuggee has
exited.

#### func  NewExitedEvent

```go
func NewExitedEvent() *ExitedEvent
```
Returns a new `ExitedEvent` with the following fields set: `Event_`, `Type`

#### type FunctionBreakpoint

```go
type FunctionBreakpoint struct {

	// The name of the function.
	Name string `json:"name"`

	// An optional expression for conditional breakpoints.
	Condition string `json:"condition,omitempty"`

	// An optional expression that controls how many hits of the breakpoint are ignored. The backend is expected to interpret the expression as needed.
	HitCondition string `json:"hitCondition,omitempty"`
} // struct FunctionBreakpoint

```

Properties of a breakpoint passed to the setFunctionBreakpoints request.

#### type GotoArguments

```go
type GotoArguments struct {

	// Set the goto target for this thread.
	ThreadId int `json:"threadId"`

	// The location where the debuggee will continue to run.
	TargetId int `json:"targetId"`
} // struct GotoArguments

```

Arguments for 'goto' request.

#### type GotoRequest

```go
type GotoRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `goto`
	Command string `json:"command"`

	Arguments GotoArguments `json:"arguments"`
} // struct GotoRequest

```

Goto request; value of command field is 'goto'. The request sets the location
where the debuggee will continue to run. This makes it possible to skip the
execution of code or to executed code again. The code between the current
location and the goto target is not executed but skipped. The debug adapter
first sends the GotoResponse and then a StoppedEvent (event type 'goto').

#### func  NewGotoRequest

```go
func NewGotoRequest() *GotoRequest
```
Returns a new `GotoRequest` with the following fields set: `Command`, `Type`

#### type GotoResponse

```go
type GotoResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `goto`
	Command string `json:"command,omitempty"`
} // struct GotoResponse

```

Response to 'goto' request. This is just an acknowledgement, so no body field is
required.

#### func  NewGotoResponse

```go
func NewGotoResponse() *GotoResponse
```
Returns a new `GotoResponse` with the following fields set: `Command`, `Type`

#### type GotoTarget

```go
type GotoTarget struct {

	// An optional end line of the range covered by the goto target.
	EndLine int `json:"endLine,omitempty"`

	// An optional end column of the range covered by the goto target.
	EndColumn int `json:"endColumn,omitempty"`

	// Unique identifier for a goto target. This is used in the goto request.
	Id int `json:"id"`

	// The name of the goto target (shown in the UI).
	Label string `json:"label"`

	// The line of the goto target.
	Line int `json:"line"`

	// An optional column of the goto target.
	Column int `json:"column,omitempty"`
} // struct GotoTarget

```

A GotoTarget describes a code location that can be used as a target in the
'goto' request. The possible goto targets can be determined via the
'gotoTargets' request.

#### type GotoTargetsArguments

```go
type GotoTargetsArguments struct {

	// The source location for which the goto targets are determined.
	Source Source `json:"source"`

	// The line location for which the goto targets are determined.
	Line int `json:"line"`

	// An optional column location for which the goto targets are determined.
	Column int `json:"column,omitempty"`
} // struct GotoTargetsArguments

```

Arguments for 'gotoTargets' request.

#### type GotoTargetsRequest

```go
type GotoTargetsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `gotoTargets`
	Command string `json:"command"`

	Arguments GotoTargetsArguments `json:"arguments"`
} // struct GotoTargetsRequest

```

GotoTargets request; value of command field is 'gotoTargets'. This request
retrieves the possible goto targets for the specified source location. These
targets can be used in the 'goto' request. The GotoTargets request may only be
called if the 'supportsGotoTargetsRequest' capability exists and is true.

#### func  NewGotoTargetsRequest

```go
func NewGotoTargetsRequest() *GotoTargetsRequest
```
Returns a new `GotoTargetsRequest` with the following fields set: `Command`,
`Type`

#### type GotoTargetsResponse

```go
type GotoTargetsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible goto targets of the specified location.
		Targets []GotoTarget `json:"targets"`
	} `json:"body"`

	// POSSIBLE VALUES: `gotoTargets`
	Command string `json:"command,omitempty"`
} // struct GotoTargetsResponse

```

Response to 'gotoTargets' request.

#### func  NewGotoTargetsResponse

```go
func NewGotoTargetsResponse() *GotoTargetsResponse
```
Returns a new `GotoTargetsResponse` with the following fields set: `Command`,
`Type`

#### type InitializeRequest

```go
type InitializeRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `initialize`
	Command string `json:"command"`

	Arguments InitializeRequestArguments `json:"arguments"`
} // struct InitializeRequest

```

Initialize request; value of command field is 'initialize'.

#### func  NewInitializeRequest

```go
func NewInitializeRequest() *InitializeRequest
```
Returns a new `InitializeRequest` with the following fields set: `Command`,
`Type`

#### type InitializeRequestArguments

```go
type InitializeRequestArguments struct {

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

	// The ID of the (frontend) client using this adapter.
	ClientID string `json:"clientID,omitempty"`
} // struct InitializeRequestArguments

```

Arguments for 'initialize' request.

#### type InitializeResponse

```go
type InitializeResponse struct {
	// Response to a request.
	Response

	// The capabilities of this debug adapter.
	Body Capabilities `json:"body,omitempty"`

	// POSSIBLE VALUES: `initialize`
	Command string `json:"command,omitempty"`
} // struct InitializeResponse

```

Response to 'initialize' request.

#### func  NewInitializeResponse

```go
func NewInitializeResponse() *InitializeResponse
```
Returns a new `InitializeResponse` with the following fields set: `Command`,
`Type`

#### type InitializedEvent

```go
type InitializedEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `initialized`
	Event_ string `json:"event"`
} // struct InitializedEvent

```

Event message for 'initialized' event type. This event indicates that the debug
adapter is ready to accept configuration requests (e.g. SetBreakpointsRequest,
SetExceptionBreakpointsRequest). A debug adapter is expected to send this event
when it is ready to accept configuration requests (but not before the
InitializeRequest has finished). The sequence of events/requests is as follows:
- adapters sends InitializedEvent (after the InitializeRequest has returned) -
frontend sends zero or more SetBreakpointsRequest - frontend sends one
SetFunctionBreakpointsRequest - frontend sends a SetExceptionBreakpointsRequest
if one or more exceptionBreakpointFilters have been defined (or if
supportsConfigurationDoneRequest is not defined or false) - frontend sends other
future configuration requests - frontend sends one ConfigurationDoneRequest to
indicate the end of the configuration

#### func  NewInitializedEvent

```go
func NewInitializedEvent() *InitializedEvent
```
Returns a new `InitializedEvent` with the following fields set: `Event_`, `Type`

#### type LaunchRequest

```go
type LaunchRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `launch`
	Command string `json:"command"`

	Arguments LaunchRequestArguments `json:"arguments"`
} // struct LaunchRequest

```

Launch request; value of command field is 'launch'.

#### func  NewLaunchRequest

```go
func NewLaunchRequest() *LaunchRequest
```
Returns a new `LaunchRequest` with the following fields set: `Command`, `Type`

#### type LaunchRequestArguments

```go
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

```

Arguments for 'launch' request.

#### type LaunchResponse

```go
type LaunchResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `launch`
	Command string `json:"command,omitempty"`
} // struct LaunchResponse

```

Response to 'launch' request. This is just an acknowledgement, so no body field
is required.

#### func  NewLaunchResponse

```go
func NewLaunchResponse() *LaunchResponse
```
Returns a new `LaunchResponse` with the following fields set: `Command`, `Type`

#### type Message

```go
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

```

A structured message object. Used to return errors from requests.

#### type Module

```go
type Module struct {

	// True if the module is optimized.
	IsOptimized bool `json:"isOptimized,omitempty"`

	// Version of Module.
	Version string `json:"version,omitempty"`

	// Logical full path to the symbol file. The exact definition is implementation defined.
	SymbolFilePath string `json:"symbolFilePath,omitempty"`

	// A name of the module.
	Name string `json:"name"`

	// optional but recommended attributes.
	// always try to use these first before introducing additional attributes.
	//
	// Logical full path to the module. The exact definition is implementation defined, but usually this would be a full path to the on-disk file for the module.
	Path string `json:"path,omitempty"`

	// True if the module is considered 'user code' by a debugger that supports 'Just My Code'.
	IsUserCode bool `json:"isUserCode,omitempty"`

	// User understandable description of if symbols were found for the module (ex: 'Symbols Loaded', 'Symbols not found', etc.
	SymbolStatus string `json:"symbolStatus,omitempty"`

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
} // struct Module

```

A Module object represents a row in the modules view. Two attributes are
mandatory: an id identifies a module in the modules view and is used in a
ModuleEvent for identifying a module for adding, updating or deleting. The name
is used to minimally render the module in the UI.

Additional attributes can be added to the module. They will show up in the
module View if they have a corresponding ColumnDescriptor.

To avoid an unnecessary proliferation of additional attributes with similar
semantics but different names we recommend to re-use attributes from the
'recommended' list below first, and only introduce new attributes if nothing
appropriate could be found.

#### type ModuleEvent

```go
type ModuleEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `module`
	Event_ string `json:"event"`

	Body struct {

		// The new, changed, or removed module. In case of 'removed' only the module id is used.
		Module Module `json:"module"`

		// The reason for the event.
		//
		// POSSIBLE VALUES: `new`, `changed`, `removed`
		Reason string `json:"reason"`
	} `json:"body"`
} // struct ModuleEvent

```

Event message for 'module' event type. The event indicates that some information
about a module has changed.

#### func  NewModuleEvent

```go
func NewModuleEvent() *ModuleEvent
```
Returns a new `ModuleEvent` with the following fields set: `Event_`, `Type`

#### type ModulesArguments

```go
type ModulesArguments struct {

	// The index of the first module to return; if omitted modules start at 0.
	StartModule int `json:"startModule,omitempty"`

	// The number of modules to return. If moduleCount is not specified or 0, all modules are returned.
	ModuleCount int `json:"moduleCount,omitempty"`
} // struct ModulesArguments

```

Arguments for 'modules' request.

#### type ModulesRequest

```go
type ModulesRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `modules`
	Command string `json:"command"`

	Arguments ModulesArguments `json:"arguments"`
} // struct ModulesRequest

```

Modules can be retrieved from the debug adapter with the ModulesRequest which
can either return all modules or a range of modules to support paging.

#### func  NewModulesRequest

```go
func NewModulesRequest() *ModulesRequest
```
Returns a new `ModulesRequest` with the following fields set: `Command`, `Type`

#### type ModulesResponse

```go
type ModulesResponse struct {
	// Response to a request.
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

```

Response to 'modules' request.

#### func  NewModulesResponse

```go
func NewModulesResponse() *ModulesResponse
```
Returns a new `ModulesResponse` with the following fields set: `Command`, `Type`

#### type ModulesViewDescriptor

```go
type ModulesViewDescriptor struct {
	Columns []ColumnDescriptor `json:"columns"`
} // struct ModulesViewDescriptor

```

The ModulesViewDescriptor is the container for all declarative configuration
options of a ModuleView. For now it only specifies the columns to be shown in
the modules view.

#### type NextArguments

```go
type NextArguments struct {

	// Execute 'next' for this thread.
	ThreadId int `json:"threadId"`
} // struct NextArguments

```

Arguments for 'next' request.

#### type NextRequest

```go
type NextRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `next`
	Command string `json:"command"`

	Arguments NextArguments `json:"arguments"`
} // struct NextRequest

```

Next request; value of command field is 'next'. The request starts the debuggee
to run again for one step. The debug adapter first sends the NextResponse and
then a StoppedEvent (event type 'step') after the step has completed.

#### func  NewNextRequest

```go
func NewNextRequest() *NextRequest
```
Returns a new `NextRequest` with the following fields set: `Command`, `Type`

#### type NextResponse

```go
type NextResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `next`
	Command string `json:"command,omitempty"`
} // struct NextResponse

```

Response to 'next' request. This is just an acknowledgement, so no body field is
required.

#### func  NewNextResponse

```go
func NewNextResponse() *NextResponse
```
Returns a new `NextResponse` with the following fields set: `Command`, `Type`

#### type OutputEvent

```go
type OutputEvent struct {
	// Server-initiated event.
	Event

	// POSSIBLE VALUES: `output`
	Event_ string `json:"event"`

	Body struct {

		// The category of output (such as: 'console', 'stdout', 'stderr', 'telemetry'). If not specified, 'console' is assumed.
		//
		// POSSIBLE VALUES: `console`, `stdout`, `stderr`, `telemetry`
		Category string `json:"category,omitempty"`

		// The output to report.
		Output string `json:"output"`

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
	} `json:"body"`
} // struct OutputEvent

```

Event message for 'output' event type. The event indicates that the target has
produced some output.

#### func  NewOutputEvent

```go
func NewOutputEvent() *OutputEvent
```
Returns a new `OutputEvent` with the following fields set: `Event_`, `Type`

#### type PauseArguments

```go
type PauseArguments struct {

	// Pause execution for this thread.
	ThreadId int `json:"threadId"`
} // struct PauseArguments

```

Arguments for 'pause' request.

#### type PauseRequest

```go
type PauseRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `pause`
	Command string `json:"command"`

	Arguments PauseArguments `json:"arguments"`
} // struct PauseRequest

```

Pause request; value of command field is 'pause'. The request suspenses the
debuggee. The debug adapter first sends the PauseResponse and then a
StoppedEvent (event type 'pause') after the thread has been paused successfully.

#### func  NewPauseRequest

```go
func NewPauseRequest() *PauseRequest
```
Returns a new `PauseRequest` with the following fields set: `Command`, `Type`

#### type PauseResponse

```go
type PauseResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `pause`
	Command string `json:"command,omitempty"`
} // struct PauseResponse

```

Response to 'pause' request. This is just an acknowledgement, so no body field
is required.

#### func  NewPauseResponse

```go
func NewPauseResponse() *PauseResponse
```
Returns a new `PauseResponse` with the following fields set: `Command`, `Type`

#### type ProtocolMessage

```go
type ProtocolMessage struct {

	// Sequence number.
	Seq int `json:"seq"`

	// One of 'request', 'response', or 'event'.
	//
	// POSSIBLE VALUES: `request`, `response`, `event`
	Type string `json:"type"`
} // struct ProtocolMessage

```

Base class of requests, responses, and events.

#### type Request

```go
type Request struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

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

	// POSSIBLE VALUES: `request`
	Type string `json:"type"`

	// The command to execute.
	Command string `json:"command"`
} // struct Request

```

A client or server-initiated request.

#### func  BaseRequest

```go
func BaseRequest(someRequest interface{}) (baseRequest *Request)
```

#### type Response

```go
type Response struct {
	// Base class of requests, responses, and events.
	ProtocolMessage

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

	// POSSIBLE VALUES: `response`
	Type string `json:"type"`

	// Sequence number of the corresponding request.
	Request_seq int `json:"request_seq"`
} // struct Response

```

Response to a request.

#### func  BaseResponse

```go
func BaseResponse(someResponse interface{}) (baseResponse *Response)
```

#### type RestartArguments

```go
type RestartArguments struct {
} // struct RestartArguments

```

Arguments for 'restart' request. The restart request has no standardized
attributes.

#### type RestartFrameArguments

```go
type RestartFrameArguments struct {

	// Restart this stackframe.
	FrameId int `json:"frameId"`
} // struct RestartFrameArguments

```

Arguments for 'restartFrame' request.

#### type RestartFrameRequest

```go
type RestartFrameRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `restartFrame`
	Command string `json:"command"`

	Arguments RestartFrameArguments `json:"arguments"`
} // struct RestartFrameRequest

```

RestartFrame request; value of command field is 'restartFrame'. The request
restarts execution of the specified stackframe. The debug adapter first sends
the RestartFrameResponse and then a StoppedEvent (event type 'restart') after
the restart has completed.

#### func  NewRestartFrameRequest

```go
func NewRestartFrameRequest() *RestartFrameRequest
```
Returns a new `RestartFrameRequest` with the following fields set: `Command`,
`Type`

#### type RestartFrameResponse

```go
type RestartFrameResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `restartFrame`
	Command string `json:"command,omitempty"`
} // struct RestartFrameResponse

```

Response to 'restartFrame' request. This is just an acknowledgement, so no body
field is required.

#### func  NewRestartFrameResponse

```go
func NewRestartFrameResponse() *RestartFrameResponse
```
Returns a new `RestartFrameResponse` with the following fields set: `Command`,
`Type`

#### type RestartRequest

```go
type RestartRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `restart`
	Command string `json:"command"`

	Arguments RestartArguments `json:"arguments,omitempty"`
} // struct RestartRequest

```

Restart request; value of command field is 'restart'. Restarts a debug session.
If the capability 'supportsRestartRequest' is missing or has the value false,
the client will implement 'restart' by terminating the debug adapter first and
then launching it anew. A debug adapter can override this default behaviour by
implementing a restart request and setting the capability
'supportsRestartRequest' to true.

#### func  NewRestartRequest

```go
func NewRestartRequest() *RestartRequest
```
Returns a new `RestartRequest` with the following fields set: `Command`, `Type`

#### type RestartResponse

```go
type RestartResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `restart`
	Command string `json:"command,omitempty"`
} // struct RestartResponse

```

Response to 'restart' request. This is just an acknowledgement, so no body field
is required.

#### func  NewRestartResponse

```go
func NewRestartResponse() *RestartResponse
```
Returns a new `RestartResponse` with the following fields set: `Command`, `Type`

#### type ReverseContinueArguments

```go
type ReverseContinueArguments struct {

	// Exceute 'reverseContinue' for this thread.
	ThreadId int `json:"threadId"`
} // struct ReverseContinueArguments

```

Arguments for 'reverseContinue' request.

#### type ReverseContinueRequest

```go
type ReverseContinueRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `reverseContinue`
	Command string `json:"command"`

	Arguments ReverseContinueArguments `json:"arguments"`
} // struct ReverseContinueRequest

```

ReverseContinue request; value of command field is 'reverseContinue'. The
request starts the debuggee to run backward. Clients should only call this
request if the capability supportsStepBack is true.

#### func  NewReverseContinueRequest

```go
func NewReverseContinueRequest() *ReverseContinueRequest
```
Returns a new `ReverseContinueRequest` with the following fields set: `Command`,
`Type`

#### type ReverseContinueResponse

```go
type ReverseContinueResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `reverseContinue`
	Command string `json:"command,omitempty"`
} // struct ReverseContinueResponse

```

Response to 'reverseContinue' request. This is just an acknowledgement, so no
body field is required.

#### func  NewReverseContinueResponse

```go
func NewReverseContinueResponse() *ReverseContinueResponse
```
Returns a new `ReverseContinueResponse` with the following fields set:
`Command`, `Type`

#### type RunInTerminalRequest

```go
type RunInTerminalRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `runInTerminal`
	Command string `json:"command"`

	Arguments RunInTerminalRequestArguments `json:"arguments"`
} // struct RunInTerminalRequest

```

runInTerminal request; value of command field is 'runInTerminal'. With this
request a debug adapter can run a command in a terminal.

#### func  NewRunInTerminalRequest

```go
func NewRunInTerminalRequest() *RunInTerminalRequest
```
Returns a new `RunInTerminalRequest` with the following fields set: `Command`,
`Type`

#### type RunInTerminalRequestArguments

```go
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

```

Arguments for 'runInTerminal' request.

#### type RunInTerminalResponse

```go
type RunInTerminalResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The process ID.
		ProcessId int64 `json:"processId,omitempty"`
	} `json:"body"`

	// POSSIBLE VALUES: `runInTerminal`
	Command string `json:"command,omitempty"`
} // struct RunInTerminalResponse

```

Response to Initialize request.

#### func  NewRunInTerminalResponse

```go
func NewRunInTerminalResponse() *RunInTerminalResponse
```
Returns a new `RunInTerminalResponse` with the following fields set: `Command`,
`Type`

#### type Scope

```go
type Scope struct {

	// Optional source for this scope.
	Source Source `json:"source,omitempty"`

	// Optional start line of the range covered by this scope.
	Line int `json:"line,omitempty"`

	// Optional end column of the range covered by this scope.
	EndColumn int `json:"endColumn,omitempty"`

	// Name of the scope such as 'Arguments', 'Locals'.
	Name string `json:"name"`

	// The variables of this scope can be retrieved by passing the value of variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`

	// The number of named variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	NamedVariables int `json:"namedVariables,omitempty"`

	// Optional end line of the range covered by this scope.
	EndLine int `json:"endLine,omitempty"`

	// The number of indexed variables in this scope.
	// The client can use this optional information to present the variables in a paged UI and fetch them in chunks.
	IndexedVariables int `json:"indexedVariables,omitempty"`

	// If true, the number of variables in this scope is large or expensive to retrieve.
	Expensive bool `json:"expensive"`

	// Optional start column of the range covered by this scope.
	Column int `json:"column,omitempty"`
} // struct Scope

```

A Scope is a named container for variables. Optionally a scope can map to a
source or a range within a source.

#### type ScopesArguments

```go
type ScopesArguments struct {

	// Retrieve the scopes for this stackframe.
	FrameId int `json:"frameId"`
} // struct ScopesArguments

```

Arguments for 'scopes' request.

#### type ScopesRequest

```go
type ScopesRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `scopes`
	Command string `json:"command"`

	Arguments ScopesArguments `json:"arguments"`
} // struct ScopesRequest

```

Scopes request; value of command field is 'scopes'. The request returns the
variable scopes for a given stackframe ID.

#### func  NewScopesRequest

```go
func NewScopesRequest() *ScopesRequest
```
Returns a new `ScopesRequest` with the following fields set: `Command`, `Type`

#### type ScopesResponse

```go
type ScopesResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The scopes of the stackframe. If the array has length zero, there are no scopes available.
		Scopes []Scope `json:"scopes"`
	} `json:"body"`

	// POSSIBLE VALUES: `scopes`
	Command string `json:"command,omitempty"`
} // struct ScopesResponse

```

Response to 'scopes' request.

#### func  NewScopesResponse

```go
func NewScopesResponse() *ScopesResponse
```
Returns a new `ScopesResponse` with the following fields set: `Command`, `Type`

#### type SetBreakpointsArguments

```go
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

```

Arguments for 'setBreakpoints' request.

#### type SetBreakpointsRequest

```go
type SetBreakpointsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setBreakpoints`
	Command string `json:"command"`

	Arguments SetBreakpointsArguments `json:"arguments"`
} // struct SetBreakpointsRequest

```

SetBreakpoints request; value of command field is 'setBreakpoints'. Sets
multiple breakpoints for a single source and clears all previous breakpoints in
that source. To clear all breakpoint for a source, specify an empty array. When
a breakpoint is hit, a StoppedEvent (event type 'breakpoint') is generated.

#### func  NewSetBreakpointsRequest

```go
func NewSetBreakpointsRequest() *SetBreakpointsRequest
```
Returns a new `SetBreakpointsRequest` with the following fields set: `Command`,
`Type`

#### type SetBreakpointsResponse

```go
type SetBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements are in the same order as the elements of the 'breakpoints' (or the deprecated 'lines') in the SetBreakpointsArguments.
		Breakpoints []Breakpoint `json:"breakpoints"`
	} `json:"body"`

	// POSSIBLE VALUES: `setBreakpoints`
	Command string `json:"command,omitempty"`
} // struct SetBreakpointsResponse

```

Response to 'setBreakpoints' request. Returned is information about each
breakpoint created by this request. This includes the actual code location and
whether the breakpoint could be verified. The breakpoints returned are in the
same order as the elements of the 'breakpoints' (or the deprecated 'lines') in
the SetBreakpointsArguments.

#### func  NewSetBreakpointsResponse

```go
func NewSetBreakpointsResponse() *SetBreakpointsResponse
```
Returns a new `SetBreakpointsResponse` with the following fields set: `Command`,
`Type`

#### type SetExceptionBreakpointsArguments

```go
type SetExceptionBreakpointsArguments struct {

	// IDs of checked exception options. The set of IDs is returned via the 'exceptionBreakpointFilters' capability.
	Filters []string `json:"filters"`

	// Configuration options for selected exceptions.
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`
} // struct SetExceptionBreakpointsArguments

```

Arguments for 'setExceptionBreakpoints' request.

#### type SetExceptionBreakpointsRequest

```go
type SetExceptionBreakpointsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command"`

	Arguments SetExceptionBreakpointsArguments `json:"arguments"`
} // struct SetExceptionBreakpointsRequest

```

SetExceptionBreakpoints request; value of command field is
'setExceptionBreakpoints'. The request configures the debuggers response to
thrown exceptions. If an exception is configured to break, a StoppedEvent is
fired (event type 'exception').

#### func  NewSetExceptionBreakpointsRequest

```go
func NewSetExceptionBreakpointsRequest() *SetExceptionBreakpointsRequest
```
Returns a new `SetExceptionBreakpointsRequest` with the following fields set:
`Command`, `Type`

#### type SetExceptionBreakpointsResponse

```go
type SetExceptionBreakpointsResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `setExceptionBreakpoints`
	Command string `json:"command,omitempty"`
} // struct SetExceptionBreakpointsResponse

```

Response to 'setExceptionBreakpoints' request. This is just an acknowledgement,
so no body field is required.

#### func  NewSetExceptionBreakpointsResponse

```go
func NewSetExceptionBreakpointsResponse() *SetExceptionBreakpointsResponse
```
Returns a new `SetExceptionBreakpointsResponse` with the following fields set:
`Command`, `Type`

#### type SetFunctionBreakpointsArguments

```go
type SetFunctionBreakpointsArguments struct {

	// The function names of the breakpoints.
	Breakpoints []FunctionBreakpoint `json:"breakpoints"`
} // struct SetFunctionBreakpointsArguments

```

Arguments for 'setFunctionBreakpoints' request.

#### type SetFunctionBreakpointsRequest

```go
type SetFunctionBreakpointsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setFunctionBreakpoints`
	Command string `json:"command"`

	Arguments SetFunctionBreakpointsArguments `json:"arguments"`
} // struct SetFunctionBreakpointsRequest

```

SetFunctionBreakpoints request; value of command field is
'setFunctionBreakpoints'. Sets multiple function breakpoints and clears all
previous function breakpoints. To clear all function breakpoint, specify an
empty array. When a function breakpoint is hit, a StoppedEvent (event type
'function breakpoint') is generated.

#### func  NewSetFunctionBreakpointsRequest

```go
func NewSetFunctionBreakpointsRequest() *SetFunctionBreakpointsRequest
```
Returns a new `SetFunctionBreakpointsRequest` with the following fields set:
`Command`, `Type`

#### type SetFunctionBreakpointsResponse

```go
type SetFunctionBreakpointsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Information about the breakpoints. The array elements correspond to the elements of the 'breakpoints' array.
		Breakpoints []Breakpoint `json:"breakpoints"`
	} `json:"body"`

	// POSSIBLE VALUES: `setFunctionBreakpoints`
	Command string `json:"command,omitempty"`
} // struct SetFunctionBreakpointsResponse

```

Response to 'setFunctionBreakpoints' request. Returned is information about each
breakpoint created by this request.

#### func  NewSetFunctionBreakpointsResponse

```go
func NewSetFunctionBreakpointsResponse() *SetFunctionBreakpointsResponse
```
Returns a new `SetFunctionBreakpointsResponse` with the following fields set:
`Command`, `Type`

#### type SetVariableArguments

```go
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

```

Arguments for 'setVariable' request.

#### type SetVariableRequest

```go
type SetVariableRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `setVariable`
	Command string `json:"command"`

	Arguments SetVariableArguments `json:"arguments"`
} // struct SetVariableRequest

```

setVariable request; value of command field is 'setVariable'. Set the variable
with the given name in the variable container to a new value.

#### func  NewSetVariableRequest

```go
func NewSetVariableRequest() *SetVariableRequest
```
Returns a new `SetVariableRequest` with the following fields set: `Command`,
`Type`

#### type SetVariableResponse

```go
type SetVariableResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The new value of the variable.
		Value string `json:"value"`

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
	} `json:"body"`

	// POSSIBLE VALUES: `setVariable`
	Command string `json:"command,omitempty"`
} // struct SetVariableResponse

```

Response to 'setVariable' request.

#### func  NewSetVariableResponse

```go
func NewSetVariableResponse() *SetVariableResponse
```
Returns a new `SetVariableResponse` with the following fields set: `Command`,
`Type`

#### type Source

```go
type Source struct {

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
} // struct Source

```

A Source is a descriptor for source code. It is returned from the debug adapter
as part of a StackFrame and it is used by clients when specifying breakpoints.

#### type SourceArguments

```go
type SourceArguments struct {

	// Specifies the source content to load. Either source.path or source.sourceReference must be specified.
	Source Source `json:"source,omitempty"`

	// The reference to the source. This is the same as source.sourceReference. This is provided for backward compatibility since old backends do not understand the 'source' attribute.
	SourceReference int `json:"sourceReference"`
} // struct SourceArguments

```

Arguments for 'source' request.

#### type SourceBreakpoint

```go
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

```

Properties of a breakpoint passed to the setBreakpoints request.

#### type SourceRequest

```go
type SourceRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `source`
	Command string `json:"command"`

	Arguments SourceArguments `json:"arguments"`
} // struct SourceRequest

```

Source request; value of command field is 'source'. The request retrieves the
source code for a given source reference.

#### func  NewSourceRequest

```go
func NewSourceRequest() *SourceRequest
```
Returns a new `SourceRequest` with the following fields set: `Command`, `Type`

#### type SourceResponse

```go
type SourceResponse struct {
	// Response to a request.
	Response

	Body struct {

		// Optional content type (mime type) of the source.
		MimeType string `json:"mimeType,omitempty"`

		// Content of the source reference.
		Content string `json:"content"`
	} `json:"body"`

	// POSSIBLE VALUES: `source`
	Command string `json:"command,omitempty"`
} // struct SourceResponse

```

Response to 'source' request.

#### func  NewSourceResponse

```go
func NewSourceResponse() *SourceResponse
```
Returns a new `SourceResponse` with the following fields set: `Command`, `Type`

#### type StackFrame

```go
type StackFrame struct {

	// The name of the stack frame, typically a method name.
	Name string `json:"name"`

	// The optional source of the frame.
	Source Source `json:"source,omitempty"`

	// The column within the line. If source is null or doesn't exist, column is 0 and must be ignored.
	Column int `json:"column"`

	// An optional end column of the range covered by the stack frame.
	EndColumn int `json:"endColumn,omitempty"`

	// The module associated with this frame, if any.
	//
	// POSSIBLE TYPES:
	// - `int` (for JSON `integer`s)
	// - `string` (for JSON `string`s)
	ModuleId interface{} `json:"moduleId,omitempty"`

	// An optional hint for how to present this frame in the UI. A value of 'label' can be used to indicate that the frame is an artificial frame that is used as a visual label or separator.
	//
	// POSSIBLE VALUES: `normal`, `label`
	PresentationHint string `json:"presentationHint,omitempty"`

	// An identifier for the stack frame. It must be unique across all threads. This id can be used to retrieve the scopes of the frame with the 'scopesRequest' or to restart the execution of a stackframe.
	Id int `json:"id"`

	// An optional end line of the range covered by the stack frame.
	EndLine int `json:"endLine,omitempty"`

	// The line within the file of the frame. If source is null or doesn't exist, line is 0 and must be ignored.
	Line int `json:"line"`
} // struct StackFrame

```

A Stackframe contains the source location.

#### type StackFrameFormat

```go
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

```

Provides formatting information for a stack frame.

#### type StackTraceArguments

```go
type StackTraceArguments struct {

	// Retrieve the stacktrace for this thread.
	ThreadId int `json:"threadId"`

	// The index of the first frame to return; if omitted frames start at 0.
	StartFrame int `json:"startFrame,omitempty"`

	// The maximum number of frames to return. If levels is not specified or 0, all frames are returned.
	Levels int `json:"levels,omitempty"`

	// Specifies details on how to format the stack frames.
	Format StackFrameFormat `json:"format,omitempty"`
} // struct StackTraceArguments

```

Arguments for 'stackTrace' request.

#### type StackTraceRequest

```go
type StackTraceRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command"`

	Arguments StackTraceArguments `json:"arguments"`
} // struct StackTraceRequest

```

StackTrace request; value of command field is 'stackTrace'. The request returns
a stacktrace from the current execution state.

#### func  NewStackTraceRequest

```go
func NewStackTraceRequest() *StackTraceRequest
```
Returns a new `StackTraceRequest` with the following fields set: `Command`,
`Type`

#### type StackTraceResponse

```go
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

	// POSSIBLE VALUES: `stackTrace`
	Command string `json:"command,omitempty"`
} // struct StackTraceResponse

```

Response to 'stackTrace' request.

#### func  NewStackTraceResponse

```go
func NewStackTraceResponse() *StackTraceResponse
```
Returns a new `StackTraceResponse` with the following fields set: `Command`,
`Type`

#### type StepBackArguments

```go
type StepBackArguments struct {

	// Exceute 'stepBack' for this thread.
	ThreadId int `json:"threadId"`
} // struct StepBackArguments

```

Arguments for 'stepBack' request.

#### type StepBackRequest

```go
type StepBackRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stepBack`
	Command string `json:"command"`

	Arguments StepBackArguments `json:"arguments"`
} // struct StepBackRequest

```

StepBack request; value of command field is 'stepBack'. The request starts the
debuggee to run one step backwards. The debug adapter first sends the
StepBackResponse and then a StoppedEvent (event type 'step') after the step has
completed. Clients should only call this request if the capability
supportsStepBack is true.

#### func  NewStepBackRequest

```go
func NewStepBackRequest() *StepBackRequest
```
Returns a new `StepBackRequest` with the following fields set: `Command`, `Type`

#### type StepBackResponse

```go
type StepBackResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `stepBack`
	Command string `json:"command,omitempty"`
} // struct StepBackResponse

```

Response to 'stepBack' request. This is just an acknowledgement, so no body
field is required.

#### func  NewStepBackResponse

```go
func NewStepBackResponse() *StepBackResponse
```
Returns a new `StepBackResponse` with the following fields set: `Command`,
`Type`

#### type StepInArguments

```go
type StepInArguments struct {

	// Execute 'stepIn' for this thread.
	ThreadId int `json:"threadId"`

	// Optional id of the target to step into.
	TargetId int `json:"targetId,omitempty"`
} // struct StepInArguments

```

Arguments for 'stepIn' request.

#### type StepInRequest

```go
type StepInRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command"`

	Arguments StepInArguments `json:"arguments"`
} // struct StepInRequest

```

StepIn request; value of command field is 'stepIn'. The request starts the
debuggee to step into a function/method if possible. If it cannot step into a
target, 'stepIn' behaves like 'next'. The debug adapter first sends the
StepInResponse and then a StoppedEvent (event type 'step') after the step has
completed. If there are multiple function/method calls (or other targets) on the
source line, the optional argument 'targetId' can be used to control into which
target the 'stepIn' should occur. The list of possible targets for a given
source line can be retrieved via the 'stepInTargets' request.

#### func  NewStepInRequest

```go
func NewStepInRequest() *StepInRequest
```
Returns a new `StepInRequest` with the following fields set: `Command`, `Type`

#### type StepInResponse

```go
type StepInResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `stepIn`
	Command string `json:"command,omitempty"`
} // struct StepInResponse

```

Response to 'stepIn' request. This is just an acknowledgement, so no body field
is required.

#### func  NewStepInResponse

```go
func NewStepInResponse() *StepInResponse
```
Returns a new `StepInResponse` with the following fields set: `Command`, `Type`

#### type StepInTarget

```go
type StepInTarget struct {

	// Unique identifier for a stepIn target.
	Id int `json:"id"`

	// The name of the stepIn target (shown in the UI).
	Label string `json:"label"`
} // struct StepInTarget

```

A StepInTarget can be used in the 'stepIn' request and determines into which
single target the stepIn request should step.

#### type StepInTargetsArguments

```go
type StepInTargetsArguments struct {

	// The stack frame for which to retrieve the possible stepIn targets.
	FrameId int `json:"frameId"`
} // struct StepInTargetsArguments

```

Arguments for 'stepInTargets' request.

#### type StepInTargetsRequest

```go
type StepInTargetsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `stepInTargets`
	Command string `json:"command"`

	Arguments StepInTargetsArguments `json:"arguments"`
} // struct StepInTargetsRequest

```

StepInTargets request; value of command field is 'stepInTargets'. This request
retrieves the possible stepIn targets for the specified stack frame. These
targets can be used in the 'stepIn' request. The StepInTargets may only be
called if the 'supportsStepInTargetsRequest' capability exists and is true.

#### func  NewStepInTargetsRequest

```go
func NewStepInTargetsRequest() *StepInTargetsRequest
```
Returns a new `StepInTargetsRequest` with the following fields set: `Command`,
`Type`

#### type StepInTargetsResponse

```go
type StepInTargetsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// The possible stepIn targets of the specified source location.
		Targets []StepInTarget `json:"targets"`
	} `json:"body"`

	// POSSIBLE VALUES: `stepInTargets`
	Command string `json:"command,omitempty"`
} // struct StepInTargetsResponse

```

Response to 'stepInTargets' request.

#### func  NewStepInTargetsResponse

```go
func NewStepInTargetsResponse() *StepInTargetsResponse
```
Returns a new `StepInTargetsResponse` with the following fields set: `Command`,
`Type`

#### type StepOutArguments

```go
type StepOutArguments struct {

	// Execute 'stepOut' for this thread.
	ThreadId int `json:"threadId"`
} // struct StepOutArguments

```

Arguments for 'stepOut' request.

#### type StepOutRequest

```go
type StepOutRequest struct {
	// A client or server-initiated request.
	Request

	Arguments StepOutArguments `json:"arguments"`

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command"`
} // struct StepOutRequest

```

StepOut request; value of command field is 'stepOut'. The request starts the
debuggee to run again for one step. The debug adapter first sends the
StepOutResponse and then a StoppedEvent (event type 'step') after the step has
completed.

#### func  NewStepOutRequest

```go
func NewStepOutRequest() *StepOutRequest
```
Returns a new `StepOutRequest` with the following fields set: `Command`, `Type`

#### type StepOutResponse

```go
type StepOutResponse struct {
	// Response to a request.
	Response

	// POSSIBLE VALUES: `stepOut`
	Command string `json:"command,omitempty"`
} // struct StepOutResponse

```

Response to 'stepOut' request. This is just an acknowledgement, so no body field
is required.

#### func  NewStepOutResponse

```go
func NewStepOutResponse() *StepOutResponse
```
Returns a new `StepOutResponse` with the following fields set: `Command`, `Type`

#### type StoppedEvent

```go
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

```

Event message for 'stopped' event type. The event indicates that the execution
of the debuggee has stopped due to some condition. This can be caused by a break
point previously set, a stepping action has completed, by executing a debugger
statement etc.

#### func  NewStoppedEvent

```go
func NewStoppedEvent() *StoppedEvent
```
Returns a new `StoppedEvent` with the following fields set: `Event_`, `Type`

#### type TerminatedEvent

```go
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

```

Event message for 'terminated' event types. The event indicates that debugging
of the debuggee has terminated.

#### func  NewTerminatedEvent

```go
func NewTerminatedEvent() *TerminatedEvent
```
Returns a new `TerminatedEvent` with the following fields set: `Event_`, `Type`

#### type Thread

```go
type Thread struct {

	// Unique identifier for the thread.
	Id int `json:"id"`

	// A name of the thread.
	Name string `json:"name"`
} // struct Thread

```

A Thread

#### type ThreadEvent

```go
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

```

Event message for 'thread' event type. The event indicates that a thread has
started or exited.

#### func  NewThreadEvent

```go
func NewThreadEvent() *ThreadEvent
```
Returns a new `ThreadEvent` with the following fields set: `Event_`, `Type`

#### type ThreadsRequest

```go
type ThreadsRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `threads`
	Command string `json:"command"`
} // struct ThreadsRequest

```

Thread request; value of command field is 'threads'. The request retrieves a
list of all threads.

#### func  NewThreadsRequest

```go
func NewThreadsRequest() *ThreadsRequest
```
Returns a new `ThreadsRequest` with the following fields set: `Command`, `Type`

#### type ThreadsResponse

```go
type ThreadsResponse struct {
	// Response to a request.
	Response

	Body struct {

		// All threads.
		Threads []Thread `json:"threads"`
	} `json:"body"`

	// POSSIBLE VALUES: `threads`
	Command string `json:"command,omitempty"`
} // struct ThreadsResponse

```

Response to 'threads' request.

#### func  NewThreadsResponse

```go
func NewThreadsResponse() *ThreadsResponse
```
Returns a new `ThreadsResponse` with the following fields set: `Command`, `Type`

#### type ValueFormat

```go
type ValueFormat struct {

	// Display the value in hex.
	Hex bool `json:"hex,omitempty"`
} // struct ValueFormat

```

Provides formatting information for a value.

#### type Variable

```go
type Variable struct {

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

	// Properties of a variable that can be used to determine how to render the variable in the UI. Format of the string value: TBD.
	Kind string `json:"kind,omitempty"`

	// Optional evaluatable name of this variable which can be passed to the 'EvaluateRequest' to fetch the variable's value.
	EvaluateName string `json:"evaluateName,omitempty"`

	// If variablesReference is > 0, the variable is structured and its children can be retrieved by passing variablesReference to the VariablesRequest.
	VariablesReference int `json:"variablesReference"`
} // struct Variable

```

A Variable is a name/value pair. Optionally a variable can have a 'type' that is
shown if space permits or when hovering over the variable's name. An optional
'kind' is used to render additional properties of the variable, e.g. different
icons can be used to indicate that a variable is public or private. If the value
is structured (has children), a handle is provided to retrieve the children with
the VariablesRequest. If the number of named or indexed children is large, the
numbers should be returned via the optional 'namedVariables' and
'indexedVariables' attributes. The client can use this optional information to
present the children in a paged UI and fetch them in chunks.

#### type VariablesArguments

```go
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

```

Arguments for 'variables' request.

#### type VariablesRequest

```go
type VariablesRequest struct {
	// A client or server-initiated request.
	Request

	// POSSIBLE VALUES: `variables`
	Command string `json:"command"`

	Arguments VariablesArguments `json:"arguments"`
} // struct VariablesRequest

```

Variables request; value of command field is 'variables'. Retrieves all child
variables for the given variable reference. An optional filter can be used to
limit the fetched children to either named or indexed children.

#### func  NewVariablesRequest

```go
func NewVariablesRequest() *VariablesRequest
```
Returns a new `VariablesRequest` with the following fields set: `Command`,
`Type`

#### type VariablesResponse

```go
type VariablesResponse struct {
	// Response to a request.
	Response

	Body struct {

		// All (or a range) of variables for the given variable reference.
		Variables []Variable `json:"variables"`
	} `json:"body"`

	// POSSIBLE VALUES: `variables`
	Command string `json:"command,omitempty"`
} // struct VariablesResponse

```

Response to 'variables' request.

#### func  NewVariablesResponse

```go
func NewVariablesResponse() *VariablesResponse
```
Returns a new `VariablesResponse` with the following fields set: `Command`,
`Type`
