package z

import (
	"runtime/debug"
	"strings"
)

type iDispatcher interface {
	dispatch(*IpcReq, *IpcResp) bool
	Init()
}

type IObjSnap interface {
	ObjSnapPrefix() string
	ObjSnap(string) interface{}
}

func ipcDecodeReqAndRespond(jsonreq []byte) *IpcResp {
	var req IpcReq
	var resp IpcResp
	if err := req.UnmarshalJSON(jsonreq); err != nil {
		resp.ErrMsg = err.Error()
	} else if Lang.InitErr != nil {
		resp.ErrMsg = Strf("%s does not appear to be installed correctly on this machine. (Re-)Install it or disable `"+Prog.Name+"` in your editor config to avoid repeats of this message. Init-time error message was:\n\n%s", Lang.Title, Lang.InitErr)
	} else if Prog.Cfg.err != nil {
		resp.ErrMsg = Strf("Your %s is currently broken: either fix it or delete it, then reload Zentient.", Prog.Cfg.filePath)
	} else if req.IpcID == IPCID_OBJ_SNAPSHOT {
		objpath, _ := req.IpcArgs.(string)
		found := false
		for _, objsnp := range Prog.objSnappers {
			if pref := objsnp.ObjSnapPrefix(); strings.HasPrefix(objpath, pref) {
				found, resp.Val = true, objsnp.ObjSnap(objpath[len(pref):])
				break
			}
		}
		if !found {
			resp.ErrMsg = BadMsg(req.IpcID.String()+" path", objpath)
		}
	} else {
		resp.to(&req)
	}
	if resp.ReqID = req.ReqID; resp.ErrMsg != "" {
		resp.IpcID = req.IpcID
	}
	return &resp
}

func (me *IpcResp) postProcess() {
	if me.Menu != nil && me.Menu.SubMenu != nil && me.Menu.SubMenu.Items == nil {
		// handles better on the client-side (and UX-wise) -- instead of a "silent nothing", show an empty menu ("nothing to choose from")
		me.Menu.SubMenu.Items = MenuItems{}
	}
}

func (me *IpcResp) onResponseReady() {
	if except := recover(); except != nil {
		me.ErrMsg = Strf("%v", except)
		println(me.ErrMsg)
		debug.PrintStack()
	}
	if me.ErrMsg != "" {
		me.ErrMsg = Strf("[%s] %s", Prog.Name, strings.TrimPrefix(me.ErrMsg, Prog.Name+": "))
		//	zero out almost-everything for a leaner response. req-ID is only added in afterwards anyways
		*me = IpcResp{ErrMsg: me.ErrMsg}
	}
}

func (me *IpcResp) to(req *IpcReq) {
	defer me.onResponseReady()
	for _, disp := range Prog.dispatchers {
		if disp.dispatch(req, me) {
			me.postProcess()
			return
		}
	}
	if !req.IpcID.Valid() {
		me.ErrMsg = BadMsg("IpcID", req.IpcID.String())
	} else {
		println(Strf("The requested feature `%s` wasn't yet implemented for %s.", req.IpcID, Lang.Title))
	}
}

func (me *IpcResp) withExtras() *IpcResp {
	me.Extras = &Extras{}
	return me
}
func (me *IpcResp) withMenu() *MenuResponse {
	me.Menu = &MenuResponse{}
	return me.Menu
}
func (me *IpcResp) withSrcIntel() *IpcResp {
	if me.SrcIntel == nil {
		me.SrcIntel = &SrcIntel{}
	}
	return me
}
