package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/zentient/z"
)




func (self *zgo) IntelDefLoc (req *z.ReqIntel) *udev.SrcMsg {
	if !devgo.Has_godef { return nil }
	return devgo.QueryDefLoc_Godef(req.Ffp, req.Src, req.Pos, req.EoL)
}


func (self *zgo) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	if devgo.Has_godef { if defdecl := devgo.QueryDefDecl_GoDef(req.Ffp, req.Src, req.Pos, req.EoL)  ;  len(defdecl)>0 {
		hovs = append(hovs, &z.RespHov { Lang: "go", Txt: defdecl })
	} }
	if len(hovs)==0 { hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools available." }) }
	return
}

func (self *zgo) IntelCmpl (req *z.ReqIntel) (cmpls []*z.RespCmpl) {
	if devgo.Has_gocode {
		if rawresp := devgo.QueryCmplSugg_Gocode(req.Ffp, req.Src, req.Pos)  ;  len(rawresp)>0 {
			for _,raw := range rawresp { if c,n,t := raw["class"] , raw["name"] , raw["type"] ; len(n)>0 {
				cmpl := &z.RespCmpl{ Label: n, Detail: c, Doc: t }
				switch c {
				case "func": cmpl.Kind = z.CMPL_FUNCTION
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
