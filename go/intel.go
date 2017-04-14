package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/zentient/z"
)

func (me *zgo) may (cmdname string) bool {
	return me.Base.CfgIntelToolEnabled(cmdname)
}


func (me *zgo) IntelDefLoc (req *z.ReqIntel) *udev.SrcMsg {
	if devgo.Has_godef && me.may("godef") { return devgo.QueryDefLoc_Godef(req.Ffp, req.Src, req.Pos) }
	return nil
}


func (me *zgo) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
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
		if rawresp := devgo.QueryCmplSugg_Gocode(req.Ffp, req.Src, req.Pos)  ;  len(rawresp)>0 {
			for _,raw := range rawresp { if c,n,t := raw["class"] , raw["name"] , raw["type"] ; len(n)>0 {
				cmpl := &z.RespCmpl{ Label: n, Detail: c, Doc: t }
				switch c {
				case "func": cmpl.Kind = z.CMPL_FUNCTION  ;  cmpl.CommitChars = []string { "(" }
				case "package": cmpl.Kind = z.CMPL_FOLDER
				case "var": cmpl.Kind = z.CMPL_VARIABLE
				case "const": cmpl.Kind = z.CMPL_CONSTANT
				case "type": switch t {
					case "struct": cmpl.Kind = z.CMPL_STRUCT
					case "interface": cmpl.Kind = z.CMPL_INTERFACE
					default: cmpl.Kind = z.CMPL_CLASS
				}
				default: cmpl.Kind = z.CMPL_COLOR
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
