package z

import (
	"fmt"
	"sort"
	"text/scanner"
	"unicode/utf8"

	"github.com/go-leap/str"
)

type ISrcIntel interface {
	iDispatcher

	CanIntel(*SrcIntelLex) bool
	ComplDetails(*SrcLens, string) *SrcIntelCompl
	ComplItems(*SrcLens) SrcIntelCompls
	ComplItemsShouldSort(*SrcLens) bool
	DefSym(*SrcLens) SrcLocs
	DefType(*SrcLens) SrcLocs
	DefImpl(*SrcLens) SrcLocs
	Highlights(*SrcLens, string) SrcLocs
	Hovers(*SrcLens) []InfoTip
	InfoBits(*SrcLens) []*SrcIntelAnnotaction
	References(*SrcLens, bool) SrcLocs
	Signature(*SrcLens) *SrcIntelSigHelp
	Symbols(*SrcLens, string, bool) SrcLenses
}

type SrcIntelLex struct {
	Ident   string
	Int     string
	Float   string
	Char    string
	String  string
	Comment string
	Other   string
}

func (me SrcIntelCompls) Len() int               { return len(me) }
func (me SrcIntelCompls) Swap(i int, j int)      { me[i], me[j] = me[j], me[i] }
func (me SrcIntelCompls) Less(i int, j int) bool { return me[i].SortText < me[j].SortText }

type SrcIntelBase struct {
	Impl ISrcIntel
}

func (*SrcIntelBase) Init() {
}

func (me *SrcIntelBase) dispatch(req *ipcReq, resp *ipcResp) bool {
	switch req.IpcID {
	case IPCID_SRCINTEL_HOVER:
		me.onHover(req, resp.withSrcIntel())
	case IPCID_SRCINTEL_INFOBITS:
		me.onInfoBits(req, resp.withSrcIntel())
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
	if lex := me.posLex(req.SrcLens); me.Impl.CanIntel(lex) {
		resp.SrcIntel.Cmpl = me.Impl.ComplItems(req.SrcLens)
		if me.Impl.ComplItemsShouldSort(req.SrcLens) {
			for _, c := range resp.SrcIntel.Cmpl {
				if c.SortText == "" {
					c.SortText = ustr.Lo(c.Label)
				}
				c.SortText = Strf("%02d", c.SortPrio) + c.SortText
			}
			sort.Sort(resp.SrcIntel.Cmpl)
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

func (me *SrcIntelBase) onInfoBits(req *ipcReq, resp *ipcResp) {
	resp.SrcIntel.InfoBits = me.Impl.InfoBits(req.SrcLens)
}

func (me *SrcIntelBase) onHover(req *ipcReq, resp *ipcResp) {
	lex := me.posLex(req.SrcLens)
	if me.Impl.CanIntel(lex) {
		resp.SrcIntel.InfoTips = me.Impl.Hovers(req.SrcLens)
	}

	var hov InfoTip
	if lex.Char != "" {
		hov.Value = Strf("`%s` — byte length %d", lex.Char, len(lex.Char[:len(lex.Char)-1][1:]))
	} else if lex.Int != "" || lex.Float != "" {
		if i, ui, f := ustr.ToI64(lex.Int, 0, 0), ustr.ToUi64(lex.Int, 0, 0), ustr.ToF64(lex.Float, 0); ui != 0 || i != 0 || f != 0 {
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
		if ustr.Pref(lex.Comment, "//") {
			hov.Value = ustr.Trim(ustr.TrimPref(lex.Comment, "//"))
		} else {
			hov.Value = ustr.TrimSuff(ustr.TrimPref(lex.Comment, "/*"), "*/")
		}
	}
	if hov.Value != "" {
		resp.SrcIntel.InfoTips = append(resp.SrcIntel.InfoTips, hov)
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

func (me *SrcIntelBase) posLex(srcLens *SrcLens) (posLex *SrcIntelLex) {
	if srcLens.EnsureSrcFull(); srcLens.Txt != "" {
		var scan scanner.Scanner
		scan.Init(ustr.Reader(srcLens.Txt)).Filename = srcLens.FilePath
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
			posLex = &SrcIntelLex{}
			switch lr {
			case scanner.Ident:
				posLex.Ident = ls
			case scanner.Int:
				posLex.Int = ls
			case scanner.Float:
				posLex.Float = ls
			case scanner.Char:
				posLex.Char = ls
			case scanner.String:
				posLex.String = ls
			case scanner.RawString:
				posLex.String = ls
			case scanner.Comment:
				posLex.Comment = ls
			default:
				posLex.Other = ls
			}
		}
	}
	return
}

func (*SrcIntelBase) posLexErrNoOp(*scanner.Scanner, string)       {}
func (*SrcIntelBase) CanIntel(*SrcIntelLex) bool                   { return true }
func (*SrcIntelBase) ComplItems(*SrcLens) SrcIntelCompls           { return nil }
func (*SrcIntelBase) ComplDetails(*SrcLens, string) *SrcIntelCompl { return nil }
func (*SrcIntelBase) ComplItemsShouldSort(*SrcLens) bool           { return false }
func (*SrcIntelBase) DefImpl(*SrcLens) SrcLocs                     { return nil }
func (*SrcIntelBase) DefSym(*SrcLens) SrcLocs                      { return nil }
func (*SrcIntelBase) DefType(*SrcLens) SrcLocs                     { return nil }
func (*SrcIntelBase) Highlights(*SrcLens, string) SrcLocs          { return nil }
func (*SrcIntelBase) Hovers(*SrcLens) []InfoTip                    { return nil }
func (*SrcIntelBase) InfoBits(*SrcLens) []*SrcIntelAnnotaction     { return nil }
func (*SrcIntelBase) References(*SrcLens, bool) SrcLocs            { return nil }
func (*SrcIntelBase) Signature(*SrcLens) *SrcIntelSigHelp          { return nil }
func (*SrcIntelBase) Symbols(*SrcLens, string, bool) SrcLenses     { return nil }
