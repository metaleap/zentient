package z

type IPkgIntel interface {
	IListMenu
}

type PkgIntelBase struct {
	ListMenuBase

	Impl IPkgIntel
}

func (me *PkgIntelBase) Init() {
	me.ListMenuBase.init(me.Impl, "Packages", "Lists %s packages %s")
}

func (me *PkgIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	default:
		return false
	}
	return true
}
