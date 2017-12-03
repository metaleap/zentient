package z

import (
	"github.com/metaleap/go-util/dev"
)

type iSrcIntel interface {
	iHandler

	Hovers(*SrcLens) []SrcIntelHover
	Symbols(*SrcLens, string, bool) udev.SrcMsgs
}

type srcIntelResp struct {
	Hovers  []SrcIntelHover `json:"hovs,omitempty"`
	Symbols udev.SrcMsgs    `json:"syms,omitempty"`
}

type SrcIntelHover struct {
	Value string `json:"value"`

	// If empty, clients default to 'markdown'
	Language string `json:"language,omitempty"`
}

type SrcIntelBase struct {
	Self iSrcIntel
}

func (me *SrcIntelBase) Init() {
	handlers = append(handlers, me.Self)
}

func (_ *SrcIntelBase) Hovers(srcLens *SrcLens) (hovs []SrcIntelHover) {
	hovs = append(hovs,
		SrcIntelHover{Value: Strf("Hovers not yet implemented for **%s** by `%s`", Lang.Title, Prog.name)},
		SrcIntelHover{Value: Strf("File: %s", srcLens.FilePath), Language: "plaintext"},
		SrcIntelHover{Value: Strf("Line/Char/Offset: %v", *srcLens.Pos)},
	)
	return
}

func (*SrcIntelBase) Symbols(srcLens *SrcLens, query string, curFileOnly bool) (srcRefs udev.SrcMsgs) {
	if curFileOnly {
		for icon := 0; icon <= 25; icon++ {
			srcRefs = append(srcRefs,
				&udev.SrcMsg{Flag: icon, Msg: Strf("%s", Symbol(icon)), Ref: srcLens.FilePath,
					Misc:   Strf("flag: %d", icon),
					Pos1Ch: 1, Pos1Ln: icon + 1, Pos2Ch: 1, Pos2Ln: icon + 1,
				},
			)
		}
	}
	return
}

func (me *SrcIntelBase) handle(req *msgReq, resp *msgResp) bool {
	switch req.MsgID {
	case MSGID_SRCINTEL_HOVER:
		me.handle_Hover(req, resp)
	case MSGID_SRCINTEL_SYMS_FILE, MSGID_SRCINTEL_SYMS_PROJ:
		me.handle_Syms(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) handle_Syms(req *msgReq, resp *msgResp) {
	var query string
	if req.MsgID == MSGID_SRCINTEL_SYMS_PROJ {
		query, _ = req.MsgArgs.(string)
	}
	resp.SrcIntel = &srcIntelResp{Symbols: me.Self.Symbols(req.SrcLens, query, req.MsgID == MSGID_SRCINTEL_SYMS_FILE)}
}

func (me *SrcIntelBase) handle_Hover(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Hovers: me.Self.Hovers(req.SrcLens)}
}
