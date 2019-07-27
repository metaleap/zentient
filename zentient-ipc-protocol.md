(update: not in all detail up to speed with the codebase)

# zentient IPC protocol

Casual definitions of terms used in this doc:
- **"client"**: an editor plugin such as `zentient-vscode` or `zentient-textadept` (or other future / custom one)
- **"server"**: some backend program catering to just one specific language (or none, but not multiple), such as `./cmd/zentient-go` or `./cmd/zentient-hs` (or other future / custom one)
- the existing *client*s written so far start / restart / end the *server*s on their own, but the server shouldn't care or rely on how or by whom its lifetime will be managed: it must be ready to accept incoming request messages  immediately, it must expect to run for a long time, and it must be prepared to be killed any moment

## General protocol flow:

Line-based JSON message exchange:
- client sends messages to the server's `stdin` and receives messages from the server's `stdout`
- every line is a self-contained JSON message object, no matter how large
- both server and client can initiate messages to the counterparty at will, however generally most of the message flow tends to fit the _"client requests, server responds"_ paradigm
  - hence, *client*s keep track of pending **"Request-IDs"** (see just below) and generate / increment them, whereas *server*s don't and just reproduce them whenever responding to requests

### common JSON message fields:

- _all_ messages sent contain at least a number-typed **"IPC-ID"** (`ii`) field, denoting the type of message being sent
- messages that either request a response, or respond to such a request contain a number-typed **"Request-ID"** (`ri`) field.

## The _server_ perspective:

Modus operandi: long-running process that however could be killed without advance notice at any moment
  - long-running indicates: it's worth performing precompuations and cache things in memory (eg. AST-based-until-source-changes etc.), and it's feasible to have its own background tasks
  - server may for diagnostic purposes perform its own logging to its `stderr` in whatever way suits it
  - handling of fully-read incoming messages should happen concurrently to the `stdin` reading itself: incoming messages should be accepted as soon as they come in
    - obviously, any *server* implementation must guard its `stdout` against any concurrent response writes to it

### Caddies

are an abstraction that servers can keep around zero-or-more of. They change their own status (busy, ready etc.) and description / notices over time, and **tend to represent background tasks like build-on-save, lint-on-save, lint-on-open, package-tree refreshing,** or whatever specific concept might fit the abstraction. (In the real world, a caddy seems to be some sort of "an on-demand runner, otherwise on constant stand-by when idle", hence the term).

# Protocol message-types *("IPC-IDs")*

	IPCID_MENUS_MAIN              = 01
	IPCID_MENUS_PKGS              = 02
	IPCID_MENUS_TOOLS             = 03

	IPCID_OBJ_SNAPSHOT            = 04
	IPCID_PAGE_HTML               = 05
	IPCID_TREEVIEW_GETITEM        = 06
	IPCID_TREEVIEW_CHILDREN       = 07
	IPCID_TREEVIEW_CHANGED        = 08
	IPCID_CFG_RESETALL            = 09
	IPCID_CFG_LIST                = 10
	IPCID_CFG_SET                 = 11
	IPCID_NOTIFY_INFO             = 12
	IPCID_NOTIFY_WARN             = 13
	IPCID_NOTIFY_ERR              = 14

	IPCID_PROJ_CHANGED            = 15
	IPCID_PROJ_POLLEVTS           = 16

	IPCID_SRCDIAG_LIST            = 17
	IPCID_SRCDIAG_RUN_CURFILE     = 18
	IPCID_SRCDIAG_RUN_OPENFILES   = 19
	IPCID_SRCDIAG_RUN_ALLFILES    = 20
	IPCID_SRCDIAG_FORGETALL       = 21
	IPCID_SRCDIAG_PEEKHIDDEN      = 22
	IPCID_SRCDIAG_PUB             = 23
	IPCID_SRCDIAG_AUTO_TOGGLE     = 24
	IPCID_SRCDIAG_AUTO_ALL        = 25
	IPCID_SRCDIAG_AUTO_NONE       = 26
	IPCID_SRCDIAG_STARTED         = 27
	IPCID_SRCDIAG_FINISHED        = 28

	IPCID_SRCMOD_FMT_SETDEFMENU   = 29
	IPCID_SRCMOD_FMT_SETDEFPICK   = 30
	IPCID_SRCMOD_FMT_RUNONFILE    = 31
	IPCID_SRCMOD_FMT_RUNONSEL     = 32
	IPCID_SRCMOD_RENAME           = 33
	IPCID_SRCMOD_ACTIONS          = 34

	IPCID_SRCINTEL_HOVER          = 35
	IPCID_SRCINTEL_SYMS_FILE      = 36
	IPCID_SRCINTEL_SYMS_PROJ      = 37
	IPCID_SRCINTEL_CMPL_ITEMS     = 38
	IPCID_SRCINTEL_CMPL_DETAILS   = 39
	IPCID_SRCINTEL_HIGHLIGHTS     = 40
	IPCID_SRCINTEL_SIGNATURE      = 41
	IPCID_SRCINTEL_REFERENCES     = 42
	IPCID_SRCINTEL_DEFSYM         = 43
	IPCID_SRCINTEL_DEFTYPE        = 44
	IPCID_SRCINTEL_DEFIMPL        = 45

	IPCID_EXTRAS_INTEL_LIST       = 46
	IPCID_EXTRAS_INTEL_RUN        = 47
	IPCID_EXTRAS_QUERY_LIST       = 48
	IPCID_EXTRAS_QUERY_RUN        = 49
