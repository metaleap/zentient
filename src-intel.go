package z

import (
	"fmt"
	"sort"
	"strings"
	"text/scanner"
	"unicode/utf8"

	"github.com/metaleap/go-util/str"
)

type ISrcIntel interface {
	iDispatcher

	ComplDetails(*SrcLens, string) *SrcIntelCompl
	ComplItems(*SrcLens) SrcIntelCompls
	DefSym(*SrcLens) SrcLocs
	DefType(*SrcLens) SrcLocs
	DefImpl(*SrcLens) SrcLocs
	Highlights(*SrcLens, string) SrcLocs
	Hovers(*SrcLens) []InfoTip
	References(*SrcLens, bool) SrcLocs
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) SrcLenses
}

type SrcIntels struct {
	Info []InfoTip `json:",omitempty"`
	Refs SrcLocs   `json:",omitempty"`
}

type srcIntelResp struct {
	SrcIntels
	Sig  *SrcIntelSigHelp `json:",omitempty"`
	Cmpl SrcIntelCompls   `json:",omitempty"`
	Syms SrcLenses        `json:",omitempty"`
}

type SrcIntelCompl struct {
	Kind          Completion   `json:"kind,omitempty"`
	Label         string       `json:"label"`
	Documentation *SrcIntelDoc `json:"documentation,omitempty"`
	Detail        string       `json:"detail,omitempty"`
	SortText      string       `json:"sortText,omitempty"`
	// FilterText    string       `json:"filterText,omitempty"`
	// InsertText    string       `json:"insertText,omitempty"`
	// CommitChars   []string     `json:"commitCharacters,omitempty"` // basically in all languages always operator/separator/punctuation (that is, "non-identifier") chars --- no need to send them for each item, for each language --- the client-side will do it
	SortPrio int `json:"-"`
}

type srcIntelLex struct {
	Ident   string
	Int     string
	Float   string
	Char    string
	String  string
	Comment string
	Other   string
}

func (me *srcIntelLex) canIntel() bool { return me == nil || me.Ident != "" || me.Other != "" }

type SrcIntelCompls []*SrcIntelCompl

func (me SrcIntelCompls) Len() int               { return len(me) }
func (me SrcIntelCompls) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me SrcIntelCompls) Less(i int, j int) bool { return me[i].SortPrio < me[j].SortPrio }

type SrcIntelDoc struct {
	Value     string `json:"value,omitempty"`
	IsTrusted bool   `json:"isTrusted,omitempty"`
}

type SrcIntelSigHelp struct {
	ActiveSignature int               `json:"activeSignature"`
	ActiveParameter int               `json:"activeParameter,omitempty"`
	Signatures      []SrcIntelSigInfo `json:"signatures,omitempty"`
}

type SrcIntelSigInfo struct {
	Label         string             `json:"label"`
	Documentation SrcIntelDoc        `json:"documentation,omitempty"`
	Parameters    []SrcIntelSigParam `json:"parameters"`
}

type SrcIntelSigParam struct {
	Label         string      `json:"label"`
	Documentation SrcIntelDoc `json:"documentation,omitempty"`
}

type SrcIntelBase struct {
	Impl ISrcIntel
}

func (*SrcIntelBase) Init() {
}

func (me *SrcIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCINTEL_HOVER:
		me.onHover(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_SYMS_FILE, IPCID_SRCINTEL_SYMS_PROJ:
		me.onSyms(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_CMPL_ITEMS:
		me.onCmplItems(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_CMPL_DETAILS:
		me.onCmplDetails(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_HIGHLIGHTS:
		me.onHighlights(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_SIGNATURE:
		me.onSignature(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_REFERENCES:
		me.onReferences(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_DEFIMPL:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefImpl)
	case IPCID_SRCINTEL_DEFSYM:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefSym)
	case IPCID_SRCINTEL_DEFTYPE:
		me.onDefinition(req, resp.withSrcIntel(), me.Impl.DefType)
	default:
		return false
	}
	return true
}

func (me *SrcIntelBase) onCmplItems(req *ipcReq, resp *ipcResp) {
	if lex := me.posLex(req.SrcLens); lex.canIntel() {
		resp.SrcIntel.Cmpl = me.Impl.ComplItems(req.SrcLens)
		var shouldsort bool
		for _, c := range resp.SrcIntel.Cmpl {
			if shouldsort = c.SortPrio != 0; shouldsort {
				break
			}
		}
		if shouldsort {
			sort.Sort(resp.SrcIntel.Cmpl)
			for i, c := range resp.SrcIntel.Cmpl {
				c.SortText = Strf("%03d", i)
			}
		}
	}
}

func (me *SrcIntelBase) onCmplDetails(req *ipcReq, resp *ipcResp) {
	itemtext, _ := req.IpcArgs.(string)
	if cmpl := me.Impl.ComplDetails(req.SrcLens, itemtext); cmpl != nil {
		resp.SrcIntel.Cmpl = SrcIntelCompls{cmpl}
	}
}

func (*SrcIntelBase) onDefinition(req *ipcReq, resp *ipcResp, def func(*SrcLens) SrcLocs) {
	resp.SrcIntel.Refs = def(req.SrcLens)
}

func (me *SrcIntelBase) onHighlights(req *ipcReq, resp *ipcResp) {
	curword, _ := req.IpcArgs.(string)
	resp.SrcIntel.Refs = me.Impl.Highlights(req.SrcLens, curword)
}

func (me *SrcIntelBase) onHover(req *ipcReq, resp *ipcResp) {
	if lex := me.posLex(req.SrcLens); lex.canIntel() {
		resp.SrcIntel.Info = me.Impl.Hovers(req.SrcLens)
	} else {
		var hov InfoTip
		if lex.Char != "" {
			hov.Value = Strf("`%s` — byte length %d", lex.Char, len(lex.Char[:len(lex.Char)-1][1:]))
		} else if lex.Int != "" || lex.Float != "" {
			if i, ui, f := ustr.ParseInt(lex.Int), ustr.ParseUint(lex.Int), ustr.ParseFloat(lex.Float); ui != 0 || i != 0 || f != 0 {
				const strf = "`%s` — `%s`\n\n"
				formats := []string{"%v", "%d", "%x", "%X", "%o", "%b", "%c", "%U", "%q"}
				if i == 0 && ui == 0 {
					formats = []string{"%v", "%g", "%G", "%f", "%0f", "%.f", "%9.6f", "%b", "%e", "%E"}
				}
				for _, format := range formats {
					if ui > 0 {
						hov.Value += Strf(strf, format, Strf(format, ui))
					} else if i != 0 {
						hov.Value += Strf(strf, format, Strf(format, i))
					} else {
						hov.Value += Strf(strf, format, Strf(format, f))
					}
				}
			}
		} else if lex.String != "" {
			var str string
			if n, e := fmt.Sscanf(lex.String, "%q", &str); e != nil {
				hov.Value = e.Error()
			} else if n > 0 && str != "" {
				hov.Value = Strf("Byte-length: %d — rune count: %d\n\n---------------------------------------\n\n%s", len(str), utf8.RuneCountInString(str), str)
			}
		} else if lex.Comment != "" {
			if strings.HasPrefix(lex.Comment, "//") {
				hov.Value = strings.TrimLeft(strings.TrimPrefix(lex.Comment, "//"), " \t")
			} else {
				hov.Value = strings.TrimSuffix(strings.TrimPrefix(lex.Comment, "/*"), "*/")
			}
		}
		if hov.Value != "" {
			resp.SrcIntel.Info = []InfoTip{hov}
		}
	}
}

func (me *SrcIntelBase) onReferences(req *ipcReq, resp *ipcResp) {
	includeDeclaration := false
	if ctx, _ := req.IpcArgs.(map[string]interface{}); ctx != nil {
		if incldecl, ok := ctx["includeDeclaration"]; ok {
			includeDeclaration, _ = incldecl.(bool)
		}
	}
	resp.SrcIntel.Refs = me.Impl.References(req.SrcLens, includeDeclaration)
}

func (me *SrcIntelBase) onSignature(req *ipcReq, resp *ipcResp) {
	if resp.SrcIntel.Sig = me.Impl.Signature(req.SrcLens); resp.SrcIntel.Sig != nil {
		for i := range resp.SrcIntel.Sig.Signatures { // vsc can't handle `null` for `parameters` but can handle `[]`
			if resp.SrcIntel.Sig.Signatures[i].Documentation.IsTrusted = true; resp.SrcIntel.Sig.Signatures[i].Parameters == nil {
				resp.SrcIntel.Sig.Signatures[i].Parameters = []SrcIntelSigParam{}
			}
		}
	}
}

func (me *SrcIntelBase) onSyms(req *ipcReq, resp *ipcResp) {
	var query string
	if req.IpcID == IPCID_SRCINTEL_SYMS_PROJ {
		query, _ = req.IpcArgs.(string)
	}
	resp.SrcIntel.Syms = me.Impl.Symbols(req.SrcLens, query, req.IpcID == IPCID_SRCINTEL_SYMS_FILE)
}

func (*SrcIntelBase) ComplItems(srcLens *SrcLens) SrcIntelCompls                    { return nil }
func (*SrcIntelBase) ComplDetails(srcLens *SrcLens, itemText string) *SrcIntelCompl { return nil }
func (*SrcIntelBase) DefImpl(srcLens *SrcLens) SrcLenses                            { return nil }
func (*SrcIntelBase) DefSym(srcLens *SrcLens) SrcLenses                             { return nil }
func (*SrcIntelBase) DefType(srcLens *SrcLens) SrcLenses                            { return nil }
func (*SrcIntelBase) Highlights(srcLens *SrcLens, curWord string) SrcLenses         { return nil }
func (*SrcIntelBase) Hovers(srcLens *SrcLens) []InfoTip                             { return nil }
func (*SrcIntelBase) posLexErrNoOp(*scanner.Scanner, string)                        {}

func (me *SrcIntelBase) posLex(srcLens *SrcLens) (poslex *srcIntelLex) {
	if srcLens.EnsureSrcFull(); srcLens.Txt != "" {
		var scan scanner.Scanner
		scan.Init(strings.NewReader(srcLens.Txt)).Filename = srcLens.FilePath
		scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments
		scan.Error = me.posLexErrNoOp
		var lr rune
		var ls string
		for r := scan.Scan(); scan.ErrorCount == 0 && r != scanner.EOF; r = scan.Scan() {
			if scan.Line > srcLens.Pos.Ln || (scan.Line == srcLens.Pos.Ln && scan.Column > srcLens.Pos.Col) {
				break
			} else {
				lr, ls = r, scan.TokenText()
			}
		}
		if ls != "" {
			poslex = &srcIntelLex{}
			switch lr {
			case scanner.Ident:
				poslex.Ident = ls
			case scanner.Int:
				poslex.Int = ls
			case scanner.Float:
				poslex.Float = ls
			case scanner.Char:
				poslex.Char = ls
			case scanner.String:
				poslex.String = ls
			case scanner.RawString:
				poslex.String = ls
			case scanner.Comment:
				poslex.Comment = ls
			default:
				poslex.Other = ls
			}
		}
	}
	return
}

func (*SrcIntelBase) References(srcLens *SrcLens, includeDeclaration bool) SrcLenses {
	return nil
}

func (*SrcIntelBase) Signature(srcLens *SrcLens) *SrcIntelSigHelp {
	return nil
}

func (*SrcIntelBase) Symbols(srcLens *SrcLens, query string, curFileOnly bool) SrcLenses {
	return nil
}
