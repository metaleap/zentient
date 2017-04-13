package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/zentient/z"
)




func (self *zgo) IntelDefLoc (ffp string, srcin string, offset string) (*udev.SrcMsg, error) {
	if !devgo.Has_godef { return nil , ugo.E("No applicable Code Intel tools available.") }
	return devgo.QueryDefLoc_Godef(ffp, srcin, offset)
}


func (self *zgo) IntelHovs (ffp string, srcin string, offset string) (hovs []*z.RespHov) {
	if devgo.Has_godef { if defdecl := devgo.QueryDefDecl_GoDef(ffp, srcin, offset)  ;  len(defdecl)>0 {
		hovs = append(hovs, &z.RespHov { Lang: "go", Txt: defdecl })
	} }
	if len(hovs)==0 { hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools available." }) }
	return
}

func (self *zgo) IntelCmpl (ffp string, srcin string, offset string) (cmpls []*z.RespCmpl) {
	if devgo.Has_gocode {
		if rawresp := devgo.QueryCmplSugg_Gocode(ffp, srcin, offset)  ;  len(rawresp)>0 {
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
