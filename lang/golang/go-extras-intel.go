package zgo

import (
	"github.com/metaleap/zentient"
)

var (
	xIntelGuruCallees = z.ExtrasItem{ID: "guru.callees", Label: "Callees",
		Detail: "For this function / method call, lists possible implementations to which it might dispatch."}
	xIntelGuruCallers = z.ExtrasItem{ID: "guru.callers", Label: "Callers",
		Detail: "For this function / method implementation, lists possible callers."}
	xIntelGuruCallstack = z.ExtrasItem{ID: "guru.callstack", Label: "Call Stack",
		Detail: "For this function / method, shows an arbitrary path to the root of the call graph."}
	xIntelGuruFreevars = z.ExtrasItem{ID: "guru.freevars", Label: "Free Variables",
		Detail: "For this selection, lists variables referenced but not defined within it."}
	xIntelGuruErrtypes = z.ExtrasItem{ID: "guru.whicherrs", Label: "Error Types",
		Detail: "For this `error` value, lists its possible types."}
	xIntelGuruPointsto = z.ExtrasItem{ID: "guru.pointsto", Label: "Pointees",
		Detail: "For this pointer-typed / container-typed expression, lists possible associated types / symbols."}
	xIntelGuruChanpeers = z.ExtrasItem{ID: "guru.peers", Label: "Channel Peers",
		Detail: "For this `<-` operation's channel, lists associated allocations, sends, receives and closes."}
)
