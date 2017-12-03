package z

import (
	"github.com/metaleap/go-util/dev"
)

type iSrcIntel interface {
	iHandler

	ComplItems(*SrcLens) []SrcIntelCompl
	Hovers(*SrcLens) []SrcIntelHover
	Symbols(*SrcLens, string, bool) udev.SrcMsgs
}

type srcIntelResp struct {
	Cmpl    []SrcIntelCompl `json:"cmpl,omitempty"`
	Hovers  []SrcIntelHover `json:"hovs,omitempty"`
	Symbols udev.SrcMsgs    `json:"syms,omitempty"`
}

type SrcIntelCompl struct {
	Label         string     `json:"label"`
	Kind          Completion `json:"kind,omitempty"`
	Detail        string     `json:"detail,omitempty"`
	Documentation string     `json:"documentation,omitempty"`
	SortText      string     `json:"sortText,omitempty"`
	FilterText    string     `json:"filterText,omitempty"`
	InsertText    string     `json:"insertText,omitempty"`
	CommitChars   []string   `json:"commitCharacters,omitempty"`
	// Range               Range      `json:"Range,omitempty"`
	// AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
	// Command Command `json:"command,omitempty"`
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

func (_ *SrcIntelBase) ComplItems(srcLens *SrcLens) (cmpls []SrcIntelCompl) {
	cmpls = make([]SrcIntelCompl, 25)
	for i := 0; i < len(cmpls); i++ {
		cmplkind := Completion(i)
		cmpls[i].Label = cmplkind.String()
		cmpls[i].Kind = cmplkind
		cmpls[i].Detail = Strf("Detail for %s", cmplkind)
		cmpls[i].Documentation = Strf("Doc for %s", cmplkind)
		// srcRefs = append(srcRefs,
		// 	&udev.SrcMsg{Flag: icon, Msg: Strf("%s", Symbol(icon)), Ref: srcLens.FilePath,
		// 		Misc:   Strf("flag: %d", icon),
		// 		Pos1Ch: 1, Pos1Ln: icon + 1, Pos2Ch: 1, Pos2Ln: icon + 1,
		// 	},
		// )
	}
	return
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
		for i := 0; i <= 25; i++ {
			srcRefs = append(srcRefs,
				&udev.SrcMsg{Flag: i, Msg: Strf("%s", Symbol(i)), Ref: srcLens.FilePath,
					Misc:   Strf("flag: %d", i),
					Pos1Ch: 1, Pos1Ln: i + 1, Pos2Ch: 1, Pos2Ln: i + 1,
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
	case MSGID_SRCINTEL_CMPL_ITEMS:
		me.handle_CmplItems(req, resp)
	case MSGID_SRCINTEL_CMPL_DETAILS:
		me.handle_CmplDetails(req, resp)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) handle_CmplItems(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Cmpl: me.Self.ComplItems(req.SrcLens)}
}

func (me *SrcIntelBase) handle_CmplDetails(req *msgReq, resp *msgResp) {
}

func (me *SrcIntelBase) handle_Hover(req *msgReq, resp *msgResp) {
	resp.SrcIntel = &srcIntelResp{Hovers: me.Self.Hovers(req.SrcLens)}
}

func (me *SrcIntelBase) handle_Syms(req *msgReq, resp *msgResp) {
	var query string
	if req.MsgID == MSGID_SRCINTEL_SYMS_PROJ {
		query, _ = req.MsgArgs.(string)
	}
	resp.SrcIntel = &srcIntelResp{Symbols: me.Self.Symbols(req.SrcLens, query, req.MsgID == MSGID_SRCINTEL_SYMS_FILE)}
}
