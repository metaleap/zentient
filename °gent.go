package z

// DO NOT EDIT: code generated with zentient-dbg-vsc-go-demo-go-gent using github.com/metaleap/go-gent

// Valid returns whether the value of this `CaddyStatus` is between `CADDY_PENDING` (inclusive) and `CADDY_GOOD` (inclusive).
func (this CaddyStatus) Valid() (ret bool) {
	ret = ((this >= CADDY_PENDING) && (this <= CADDY_GOOD))
	return
}

// Valid returns whether the value of this `IpcIDs` is between `IPCID_MENUS_MAIN` (inclusive) and `IPCID_EXTRAS_QUERY_RUN` (inclusive).
func (this IpcIDs) Valid() (ret bool) {
	ret = ((this >= IPCID_MENUS_MAIN) && (this <= IPCID_EXTRAS_QUERY_RUN))
	return
}

// Valid returns whether the value of this `DiagSeverity` is between `DIAG_SEV_ERR` (inclusive) and `DIAG_SEV_HINT` (inclusive).
func (this DiagSeverity) Valid() (ret bool) {
	ret = ((this >= DIAG_SEV_ERR) && (this <= DIAG_SEV_HINT))
	return
}

// Valid returns whether the value of this `Symbol` is between `SYM_FILE` (inclusive) and `SYM_TYPEPARAMETER` (inclusive).
func (this Symbol) Valid() (ret bool) {
	ret = ((this >= SYM_FILE) && (this <= SYM_TYPEPARAMETER))
	return
}

// Valid returns whether the value of this `Completion` is between `CMPL_TEXT` (inclusive) and `CMPL_TYPEPARAMETER` (inclusive).
func (this Completion) Valid() (ret bool) {
	ret = ((this >= CMPL_TEXT) && (this <= CMPL_TYPEPARAMETER))
	return
}

// Valid returns whether the value of this `ToolCats` is between `TOOLS_CAT_MOD_REN` (inclusive) and `TOOLS_CAT_RUNONSAVE` (inclusive).
func (this ToolCats) Valid() (ret bool) {
	ret = ((this >= TOOLS_CAT_MOD_REN) && (this <= TOOLS_CAT_RUNONSAVE))
	return
}
