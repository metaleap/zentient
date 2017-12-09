package z

type IPkgIntel interface {
	IListMenu
}

type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel
}

func init() {
	var dummy PkgIntelBase
	Lang.PkgIntel, dummy.Impl = &dummy, &dummy
}

func (me *PkgIntelBase) Init() {
	me.ListMenuBase.init(me.Impl, "Packages", "%s packages %s")
}

func (me *PkgIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	default:
		return false
	}
	return true
}
