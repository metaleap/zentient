package z

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/str"
)

const (
	REQ_ZEN_STATUS    = "ZS:"
	REQ_ZEN_LANGS     = "ZL:"
	REQ_ZEN_CONFIG    = "ZC:"
	REQ_QUERY_CAPS    = "QC:"
	REQ_QUERY_DIAGS   = "QD:"
	REQ_QUERY_TOOL    = "Qt:"
	REQ_INTEL_DEFLOC  = "IL:"
	REQ_INTEL_TDEFLOC = "IT:"
	REQ_INTEL_IMPLS   = "IM:"
	REQ_INTEL_REFS    = "IR:"
	REQ_INTEL_HOVER   = "IH:"
	REQ_INTEL_CMPL    = "IC:"
	REQ_INTEL_CMPLDOC = "ID:"
	REQ_INTEL_HILITES = "II:"
	REQ_INTEL_SYM     = "IS:"
	REQ_INTEL_WSYM    = "IW:"
	REQ_TOOLS_INTEL   = "Ti:"
	REQ_TOOL_INTEL    = "TI:"
	REQ_TOOLS_QUERY   = "Tq:"
	REQ_TOOL_QUERY    = "TQ:"
	REQ_DO_FMT        = "DF:"
	REQ_DO_RENAME     = "DR:"
	REQ_FILES_OPENED  = "FO:"
	REQ_FILES_CLOSED  = "FC:"
	REQ_FILES_WRITTEN = "FW:"
)

// globals set from main-app on init. 'bad style', but ok for this personal pet project
var (
	Out    *json.Encoder
	RawOut *bufio.Writer
)

func out(v interface{}) error {
	return Out.Encode(v)
}

func HandleRequest(queryln string) (e error) {
	var inlst []string
	var inmap map[string]interface{}
	var inint ReqIntel
	var zid string
	msgid, msgrest := ustr.BreakAt(queryln, 3)
	msgzids, msgargs := ustr.BreakOn(msgrest, ":")
	zids := ustr.Split(msgzids, ",")
	if len(zids) > 0 {
		zid = zids[0]
	}

	if len(msgargs) > 1 {
		if msgargs[0] == '[' {
			json.Unmarshal([]byte(msgargs), &inlst)
		} else if msgargs[0] == '{' {
			if ustr.Pref(msgargs, "{\"Ffp\":\"") {
				json.Unmarshal([]byte(msgargs), &inint)
				inint.Ffp = normalizeFilePath(inint.Ffp)
			} else {
				json.Unmarshal([]byte(msgargs), &inmap)
			}
		}
	}
	switch msgid {
	//  each case is ideally just a single func-call out, rpc-like
	//  anything else in a case then is only to furnish proper func args from msg-argstr / json

	case REQ_INTEL_DEFLOC:
		e = out(Zengines[zid].IntelDefLoc(&inint, false))
	case REQ_INTEL_TDEFLOC:
		e = out(Zengines[zid].IntelDefLoc(&inint, true))
	case REQ_INTEL_IMPLS:
		e = out(Zengines[zid].IntelImpls(&inint))
	case REQ_INTEL_REFS:
		e = out(Zengines[zid].IntelRefs(&inint))
	case REQ_INTEL_HOVER:
		e = out(Zengines[zid].IntelHovs(&inint))
	case REQ_INTEL_CMPL:
		e = out(Zengines[zid].IntelCmpl(&inint))
	case REQ_INTEL_CMPLDOC:
		e = out(Zengines[zid].IntelCmplDoc(&inint))
	case REQ_INTEL_HILITES:
		e = out(Zengines[zid].IntelHiLites(&inint))
	case REQ_INTEL_SYM:
		e = out(Zengines[zid].IntelSymbols(&inint, false))
	case REQ_INTEL_WSYM:
		e = out(Zengines[zid].IntelSymbols(&inint, true))
	case REQ_TOOL_INTEL:
		µ := Zengines[zid]
		if srcrefs, ran := µ.B().intelTool(µ, &inint); ran {
			e = out(srcrefs)
		} else if srcrefs, err := µ.IntelTool(&inint); err != nil && len(srcrefs) == 0 {
			e = out(err.Error())
		} else {
			e = out(srcrefs)
		}
	case REQ_TOOLS_INTEL:
		µ := Zengines[msgargs]
		e = out(append(µ.B().intelTools(µ), µ.IntelTools()...))
	case REQ_TOOLS_QUERY:
		e = out(Zengines[msgargs].QueryTools())
	case REQ_TOOL_QUERY:
		e = out(Zengines[zid].QueryTool(&inint))

	case REQ_FILES_WRITTEN:
		onFilesWritten(Zengines[zid], inlst)
		e = out(jsonLiveDiags(zid, nil, nil))
	case REQ_FILES_OPENED:
		onFilesOpened(Zengines[zid], inlst)
		e = out(jsonLiveDiags(zid, nil, inlst))
	case REQ_FILES_CLOSED:
		onFilesClosed(Zengines[zid], inlst)
		e = out(jsonLiveDiags(zid, inlst, nil))

	case REQ_DO_FMT:
		if resp, err := doFmt(zid, umisc.S(inmap["s"]), umisc.S(inmap["c"]), uint8(umisc.F(inmap["t"]))); err != nil {
			e = out(err.Error())
		} else {
			e = out(resp)
		}
	case REQ_DO_RENAME:
		if resp, err := doRename(zid, umisc.S(inmap["c"]), umisc.S(inmap["rfp"]), uint64(ustr.ParseInt(umisc.S(inmap["o"]))), umisc.S(inmap["nn"]), umisc.S(inmap["e"]), umisc.S(inmap["no"]), uint64(ustr.ParseInt(umisc.S(inmap["o1"]))), uint64(ustr.ParseInt(umisc.S(inmap["o2"])))); err != nil {
			e = out(err.Error())
		} else {
			e = out(resp)
		}

	case REQ_QUERY_TOOL:
		cmdargs := ustr.Split(msgrest, " ")
		cmdstdout, cmdstderr, cmderr := urun.CmdExecStdin("", "", cmdargs[0], cmdargs[1:]...)
		cmdstdout = strings.TrimSpace(cmdstdout)
		errstr := ""
		if cmderr != nil {
			errstr = cmderr.Error()
		}
		if len(errstr) == 0 && len(cmdstderr) == 0 && ((ustr.Pref(cmdstdout, "{") && ustr.Suff(cmdstdout, "}")) || (ustr.Pref(cmdstdout, "[") && ustr.Suff(cmdstdout, "]"))) {
			RawOut.Write([]byte(strings.Replace(cmdstdout, "\n", " ", -1) + "\n"))
		} else {
			e = out(map[string]string{"_stdout": cmdstdout, "_stderr": cmdstderr, "_err": errstr})
		}
	case REQ_QUERY_DIAGS:
		e = out(jsonLiveDiags("", nil, nil))
	case REQ_QUERY_CAPS:
		resp := map[string][]*RespCmd{}
		for _, zid := range zids {
			if µ := Zengines[zid]; µ != nil {
				resp[zid] = µ.Caps(msgargs)
			}
		}
		e = out(resp)

	case REQ_ZEN_LANGS:
		e = out(jsonZengines())
	case REQ_ZEN_STATUS:
		e = out(jsonStatus())
	case REQ_ZEN_CONFIG:
		Zengines[zid].OnCfg(inmap)
		e = out(nil)

	//  nothing matched? a bug in client, throw at client
	default:
		e = out(jsonErrMsg("Unknown MSG-ID in: `" + queryln + "`"))
	}
	return
}
