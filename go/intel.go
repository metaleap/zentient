package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)

func (me *zgo) may (cmdname string) bool {
	return me.Base.CfgIntelToolEnabled(cmdname)
}


func (me *zgo) IntelDefLoc (req *z.ReqIntel) (refloc *udev.SrcMsg) {
	req.RunePosToBytePos()
	if refloc==nil && devgo.Has_godef && me.may("godef") { refloc = devgo.QueryDefLoc_Godef(req.Ffp, req.Src, req.Pos) }
	if refloc==nil && devgo.Has_gogetdoc && me.may("gogetdoc") { refloc = devgo.QueryDefLoc_Gogetdoc(req.Ffp, req.Pos) }
	return
}


func (me *zgo) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	req.RunePosToBytePos()
	var ggd *devgo.Gogetdoc
	var decl string
	if devgo.Has_gogetdoc && me.may("gogetdoc") { if ggd = devgo.Query_Gogetdoc(req.Ffp, req.Pos)  ;  ggd!=nil {
		d := ggd.ImpN  ;  if len(d)>0  {  d = "**" + d + "**\n\n"  }
		d = d + ggd.Doc
		hovs = append(hovs, &z.RespHov { Txt: d })
	} }
	if ggd!=nil && len(ggd.Decl)>0 { decl = ggd.Decl }
	if len(decl)==0 && devgo.Has_godef && me.may("godef") { decl = devgo.QueryDefDecl_GoDef(req.Ffp, req.Src, req.Pos) }
	if len(decl)>0 { hovs = append(hovs, &z.RespHov { Lang: "go", Txt: decl }) }
	return
}

func (me *zgo) IntelCmpl (req *z.ReqIntel) (cmpls []*z.RespCmpl) {
	if devgo.Has_gocode && me.may("gocode") {
		if rawresp := devgo.QueryCmplSugg_Gocode(req.Ffp, req.Src, "c" + req.Pos)  ;  len(rawresp)>0 {
			for _,raw := range rawresp { if c,n,t := raw["class"] , raw["name"] , raw["type"] ; len(n)>0 {
				cmpl := &z.RespCmpl{ Label: n, Detail: t, Doc: c }
				switch c {
				case "func": cmpl.Kind = z.CMPL_FUNCTION   ;  cmpl.SortTxt = "9" + cmpl.Label  ;  cmpl.CommitChars = []string { "(" }
				case "package": cmpl.Kind = z.CMPL_FOLDER  ;  cmpl.SortTxt = "1" + cmpl.Label
				case "var": cmpl.Kind = z.CMPL_VARIABLE  ;  cmpl.SortTxt = "4" + cmpl.Label
				case "const": cmpl.Kind = z.CMPL_CONSTANT  ;  cmpl.SortTxt = "3" + cmpl.Label
				case "type": cmpl.SortTxt = "2" + cmpl.Label  ;  switch t {
					case "struct": cmpl.Kind = z.CMPL_STRUCT
					case "interface": cmpl.Kind = z.CMPL_INTERFACE
					default: if ustr.Pref(t, "func(") {
						cmpl.Kind = z.CMPL_METHOD } else { cmpl.Kind = z.CMPL_CLASS }
				}
				default: cmpl.Kind = z.CMPL_COLOR  ;  cmpl.SortTxt = "0" + cmpl.Label
				}
				if (len(raw) > 3) { for k,v := range raw { if k!="class" && k!="name" && k!="type" {
					cmpl.Doc = "❬" + k + "=" + v + "❭ " + cmpl.Doc
				} } }
				cmpls = append(cmpls, cmpl)
			} }
		}

	}
	return
}

func (_ *zgo) IntelCmplDoc(req *z.ReqIntel) *z.RespTxt {
	return &z.RespTxt { Result: "foo `" + req.Sym1 + "` > `" + req.Sym2 + "` dis" }
}
