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
	if srcLens == nil {
		srcRefs = append(srcRefs,
			&udev.SrcMsg{Flag: SYM_FILE, Msg: "The Proj Symbol", Ref: "/home/__/c/go/src/github.com/metaleap/zentient/z.go",
				Misc:   "query: " + query,
				Pos1Ch: 2, Pos1Ln: 3, Pos2Ch: 5, Pos2Ln: 3,
			},
		)
	} else {
		srcRefs = append(srcRefs,
			&udev.SrcMsg{Flag: SYM_FILE, Msg: "The File Symbol", Ref: srcLens.FilePath,
				Misc:   "query: " + query,
				Pos1Ch: 2, Pos1Ln: 3, Pos2Ch: 5, Pos2Ln: 3,
			},
		)
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
