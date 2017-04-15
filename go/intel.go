package zgo
import (
	"strings"

	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/go-util-str"
	"github.com/metaleap/zentient/z"
)

func (me *zgo) may (cmdname string) bool {
	return me.Base.CfgIntelToolEnabled(cmdname)
}


func (me *zgo) IntelDefLoc (req *z.ReqIntel, typedef bool) (refloc *udev.SrcMsg) {
	req.RunePosToBytePos()
	//	go to definition
	if (!typedef) {
		if refloc==nil && devgo.Has_guru && me.may("guru") { if gd := devgo.QueryDescribe_Guru(req.Ffp, req.Src, req.Pos)  ;  gd!=nil {
			if gd.Type!=nil && len(gd.Type.NamePos)>0 { if rl,ok := udev.SrcMsgFromLn(gd.Type.NamePos)  ;  ok { refloc = &rl } }
			if gd.Value!=nil && len(gd.Value.ObjPos)>0 { if rl,ok := udev.SrcMsgFromLn(gd.Value.ObjPos)  ;  ok { refloc = &rl } }
		} }
		if refloc==nil && devgo.Has_gogetdoc && me.may("gogetdoc") { refloc = devgo.QueryDefLoc_Gogetdoc(req.Ffp, req.Src, req.Pos) }
		if refloc==nil && devgo.Has_godef && me.may("godef") { refloc = devgo.QueryDefLoc_Godef(req.Ffp, req.Src, req.Pos) }
		return
	}
	//	go to type definition
	if devgo.Has_guru && me.may("guru") {
		if gd := devgo.QueryDescribe_Guru(req.Ffp, req.Src, req.Pos)  ;  gd!=nil {
			if gd.Type!=nil && len(gd.Type.NamePos)>0 {
				if rl,ok := udev.SrcMsgFromLn(gd.Type.NamePos)  ;  ok { refloc = &rl }
			} else if gd.Value!=nil && len(gd.Value.Type)>0 {
				possiblyfullyqualified := strings.TrimLeft(strings.TrimPrefix(strings.TrimLeft(gd.Value.Type, "*"), "[]"), "*")
				pkgimppath,typename := ustr.BreakOnLast(possiblyfullyqualified, ".")  ;  pkgname := ustr.AfterLast(pkgimppath, "/", false)
				if devgo.PkgsByImP!=nil { if pkg := devgo.PkgsByImP[pkgimppath]  ;  pkg!=nil && len(pkg.Name)>0 {  pkgname = pkg.Name  } }
				hacky1 := "\n\nfunc Zen" + req.Id + " () *"  ;  hacky2 := " { return nil }\n"  ;  if len(pkgname)>0 {  hacky1 = hacky1 + pkgname + "."  }
				req.Pos = ugo.SPr(len( (req.Src)) + len(hacky1))  ;  req.Src = req.Src + hacky1 + typename + hacky2
				refloc = me.IntelDefLoc(req, false)
			}
		}
	}
	return
}


func (me *zgo) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	req.RunePosToBytePos()
	var ggd *devgo.Gogetdoc
	var decl string
	if devgo.Has_gogetdoc && me.may("gogetdoc") { if ggd = devgo.Query_Gogetdoc(req.Ffp, req.Src, req.Pos)  ;  ggd!=nil && len(ggd.Doc)>0 {
		d := ggd.ImpN  ;  if len(d)>0  {  d = "**" + d + "**\n\n"  }
		d = d + ggd.Doc
		hovs = append(hovs, &z.RespHov { Txt: d })
	} }
	if ggd!=nil && len(ggd.Decl)>0 { decl = ggd.Decl }
	if len(decl)==0 && devgo.Has_godef && me.may("godef") { decl = devgo.QueryDefDecl_GoDef(req.Ffp, req.Src, req.Pos) }
	if decl = ustr.Trim(decl)  ;  len(decl)>0 {  declhov := &z.RespHov { Lang: "go", Txt: decl }
		if ustr.Has(decl, "\n") { hovs = append(hovs, declhov) } else {
			hovs = append([]*z.RespHov{ declhov }, hovs...) } }
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

func (me *zgo) IntelCmplDoc(req *z.ReqIntel) *z.RespTxt {
	req.RunePosToBytePos()
	curword := req.Sym1	 ;  replword := req.Sym2  ;  wordpos := int(ustr.ParseInt(req.Pos))
	if curword!=replword { if wp := wordPos(req.Src, curword, wordpos)  ;  wp>=0 {
		wordpos = wp  ;  req.Pos = ugo.SPr(wp)
		req.Src = req.Src[:wordpos] + replword + req.Src[wordpos+len(curword):]
	} }
	if devgo.Has_gogetdoc && me.may("gogetdoc") { if ggd := devgo.Query_Gogetdoc(req.Ffp, req.Src, req.Pos)  ;  ggd!=nil {
		if d := ustr.Trim(ggd.Doc)  ;  len(d)>0 { return &z.RespTxt { Id: req.Id, Result: d } } } }
	return nil
}


func wordPos (src string, word string, wordpos int) (wp int) {
	for i,l := wordpos , len(word)+1  ;  i>=0 && i>wordpos-l  ;  i-- {
		if idx := ustr.Idx(src[i:], word)  ;  idx==0 { wp = i  ;  break } }
	return
}
