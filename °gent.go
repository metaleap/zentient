package z

// DO NOT EDIT: code generated with zentient-dbg-vsc-go-demo-go-gent using github.com/metaleap/go-gent

func (this CaddyStatus) Valid() (ret bool) {
	ret = ((this >= CADDY_PENDING) && (this <= CADDY_GOOD))
	return
}

func (this IpcIDs) Valid() (ret bool) {
	ret = ((this >= IPCID_MENUS_MAIN) && (this <= IPCID_EXTRAS_QUERY_RUN))
	return
}

func (this DiagSeverity) Valid() (ret bool) {
	ret = ((this >= DIAG_SEV_ERR) && (this <= DIAG_SEV_HINT))
	return
}

func (this Symbol) Valid() (ret bool) {
	ret = ((this >= SYM_FILE) && (this <= SYM_TYPEPARAMETER))
	return
}

func (this Completion) Valid() (ret bool) {
	ret = ((this >= CMPL_TEXT) && (this <= CMPL_TYPEPARAMETER))
	return
}

func (this ToolCats) Valid() (ret bool) {
	ret = ((this >= TOOLS_CAT_MOD_REN) && (this <= TOOLS_CAT_RUNONSAVE))
	return
}
