# zentient IPC protocol

Casual definitions of terms used in this doc:
- **"client"**: an editor plugin such as `zentient-vscode` or `zentient-textadept` (or other future / custom one)
- **"server"**: some backend program catering to just one specific language (or none, but not multiple), such as `zentient-go` or `zentient-hs` (or other future / custom one)
- the existing *client*s written so far start / restart / end the servers on their own, but the server shouldn't care or rely on how or by whom its lifetime will be managed: it must be ready to accept incoming request messages  immediately, it must expect to run for a long time, and it must be prepared to be killed any moment

## General protocol flow:

Line-based JSON message exchange via stdin/stdout
- every line is a self-contained JSON message object, no matter how large
- server may log to its stderr however and whenever desired

## The _server_ perspective:

M.O.: long-running process that however could be killed with no warning any moment
  - long-running means: it's worth doing precompuations and cachings (unlike other (non)designs where short-lived processes get run ad-hoc all the time for every auto-completion drop-down and every hover tip etc)
