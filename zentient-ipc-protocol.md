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

M.O.: long-running process that however could be killed with no warning any moment
  - long-running indicates: it's worth performing precompuations and cache things in memory (eg. AST-based-until-source-changes etc.), and it's feasible to have its own background tasks
  - server may for diagnostic purposes perform its own logging to its `stderr` in whatever way suits it
  - server should handle incoming messages in (a) separate thread(s), such that its `stdin` stream is always reading and ready (therefore it must also guard its `stdout` against concurrent output writes)

### Caddies

are an abstraction server can keep around zero-or-more of. They change their own status (busy, ready etc.) and description / notices over time, and **tend to represent background tasks like build-on-save, lint-on-save, lint-on-open, package-tree refreshing,** or whatever specific concept might fit the abstraction. (In the real world, a caddy seems to be some sort of "an on-demand runner, otherwise on constant stand-by", hence the term).
