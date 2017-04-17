package zgo
import (
	"bytes"
	"sort"
	"strings"

	gurujson "golang.org/x/tools/cmd/guru/serial"

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
		if refloc==nil && devgo.Has_guru && me.may("guru") { if gd := devgo.QueryDesc_Guru(req.Ffp, req.Src, req.Pos)  ;  gd!=nil {
			if gd.Type!=nil && len(gd.Type.NamePos)>0 { if rl := udev.SrcMsgFromLn(gd.Type.NamePos)  ;  rl!=nil { refloc = rl } }
			if gd.Value!=nil && len(gd.Value.ObjPos)>0 { if rl := udev.SrcMsgFromLn(gd.Value.ObjPos)  ;  rl!=nil { refloc = rl } }
		} }
		if refloc==nil && devgo.Has_gogetdoc && me.may("gogetdoc") { refloc = devgo.QueryDefLoc_Gogetdoc(req.Ffp, req.Src, req.Pos) }
		if refloc==nil && devgo.Has_godef && me.may("godef") { refloc = devgo.QueryDefLoc_Godef(req.Ffp, req.Src, req.Pos) }
		return
	}
	//	go to type definition
	if devgo.Has_guru && me.may("guru") {
		if gd := devgo.QueryDesc_Guru(req.Ffp, req.Src, req.Pos)  ;  gd!=nil {
			if gd.Type!=nil && len(gd.Type.NamePos)>0 {
				if rl := udev.SrcMsgFromLn(gd.Type.NamePos)  ;  rl!=nil { refloc = rl }
			} else if gd.Value!=nil && len(gd.Value.Type)>0 {
				for ustr.Pref(gd.Value.Type, "map[") {  gd.Value.Type = gd.Value.Type[ustr.Idx(gd.Value.Type, "]")+1:]  }
				possiblyfullyqualified := strings.TrimLeft(strings.TrimPrefix(strings.TrimLeft(gd.Value.Type, "*"), "[]"), "*")
				pkgimppath,typename := ustr.BreakOnLast(possiblyfullyqualified, ".")  ;  pkgname := ustr.AfterLast(pkgimppath, "/", false)
				if devgo.PkgsByImP!=nil { if pkg := devgo.PkgsByImP[pkgimppath]  ;  pkg!=nil && len(pkg.Name)>0 {  pkgname = pkg.Name  } }
				hacky1 := "\n\nfunc Zen" + req.Id + " () *"  ;  hacky2 := " { return nil }\n"  ;  if len(pkgname)>0 {  hacky1 = hacky1 + pkgname + "."  }
				req.Pos = ugo.SPr(len(req.Src) + len(hacky1))  ;  req.Src = req.Src + hacky1 + typename + hacky2
				refloc = me.IntelDefLoc(req, false)
			}
		}
	}
	return
}


func (me *zgo) IntelImpls (req *z.ReqIntel) (srcrefs udev.SrcMsgs) {
	req.RunePosToBytePos()
	if devgo.Has_guru && me.may("guru") { if gi := devgo.QueryImpl_Guru(req.Ffp, req.Src, req.Pos)  ;  gi!=nil && (len(gi.AssignableTo)>0 || len(gi.AssignableFrom)>0 || len(gi.AssignableFromPtr)>0) {
		add := func (desc string, its ...gurujson.ImplementsType) {
			for _,it := range its { if srcref := udev.SrcMsgFromLn(it.Pos)  ;  srcref!=nil {
				srcref.Msg = devgo.ShortenImPs(it.Name)  ;  srcref.Misc = it.Kind + " " + desc
				srcrefs = append(srcrefs, srcref)
			} }
		}
		add("type of current selection", gi.T)  ;  tname := devgo.ShortenImPs(gi.T.Name)
		add("type implementing `" + tname + "`", gi.AssignableTo...)
		add("type implemented by `" + tname + "`", gi.AssignableFrom...)
		add("type implemented by `*" + tname + "`", gi.AssignableFromPtr...)
	} }
	return
}


func (me *zgo) IntelHovs (req *z.ReqIntel) (hovs []*z.RespHov) {
	req.RunePosToBytePos()
	var ggd *devgo.Gogetdoc
	var decl string
	if devgo.Has_gogetdoc && me.may("gogetdoc") { if ggd = devgo.Query_Gogetdoc(req.Ffp, req.Src, req.Pos)  ;  ggd!=nil && len(ggd.Doc)>0 {
		d := ggd.ImpN  ;  if len(d)>0  {  d = "### " + d + " [🕮](http://godoc.org/" + ggd.DocUrl + ")\n\n"  }
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
				cmpl := &z.RespCmpl{ RespIntel: z.RespIntel { Label: n, Doc: c }, Detail: t }
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


func (me *zgo) IntelHiLites(req *z.ReqIntel) (srcrefs udev.SrcMsgs) {
	req.RunePosToBytePos()
	if devgo.Has_guru && me.may("guru") { if gw := devgo.QueryWhat_Guru(req.Ffp, req.Src, req.Pos1)  ;  gw!=nil {
		for _,sameid := range gw.SameIDs { if srcref := udev.SrcMsgFromLn(sameid)  ;  srcref!=nil {
			srcrefs = append(srcrefs, srcref)
		} }
		if len(srcrefs)==0 && len(gw.Enclosing)>0 {
			bpos2rpos := func(bytepos int) int {
				return bytepos // *technically* the below is "correcter" but *practically* we get the same off-by-n quirks with utf8 chars with the below as with the left.. thanks, guru, for friggin "byte offsets" everywhere
				var strbuf bytes.Buffer  ;  for byteidx,char := range req.Src {
					if byteidx>=bytepos { return strbuf.Len() }  ;  strbuf.WriteRune(char) }
				return bytepos
			}
			check := func (num int, checks ...string) bool {
				if ustr.AnyOf(gw.Enclosing[0].Description, checks[:num]...) {
					for _,syntaxnode := range gw.Enclosing { if ustr.AnyOf(syntaxnode.Description, checks[num:]...) {
						srcrefs = append(srcrefs, &udev.SrcMsg { Pos2Ln: bpos2rpos(syntaxnode.Start), Pos2Ch: bpos2rpos(syntaxnode.End) })
						return true
					} }
				}
				return false
			}
			if check(2, "defer statement", "return statement", "function literal", "function declaration") { return }
			if check(1, "break statement", "range loop", "for loop", "select statement", "switch statement") { return }
			if check(1, "continue statement", "range loop", "for loop") { return }
		}
	} }
	return
}


func (me *zgo) IntelSymbols(req *z.ReqIntel, allfiles bool) (srcrefs udev.SrcMsgs) {
	req.EnsureSrc()
	if ustr.Pref(req.Src, "package ") { req.Pos = "8" } else {
		j := 0  ;  lns := ustr.Split(req.Src, "\n")
		for i,ln := range lns { if ustr.Pref(ustr.Trim(ln), "package ") {
			for bytepos,char := range req.Src { if char=='\n' { if j++  ;  j==i {
				req.Pos = ugo.SPr(bytepos + 9)
				break
			} } }
			break
		} }
	}
	if devgo.Has_guru && me.may("guru") { if gd := devgo.QueryDesc_Guru(req.Ffp, req.Src, req.Pos)  ;  gd!=nil && gd.Package!=nil {
		fbreak := func (fdecl string) (fargs string , fret string) {
			fdecl = devgo.ShortenImPs(fdecl)
			if p3 := ustr.Idx(fdecl, ") ")  ;  p3<=0 {  fargs,fret = fdecl,"void"  } else {
				fret = fdecl[p3+2:]  ;  fargs = fdecl[:p3+1]
			}
			return
		}
		fpathok := func (fpath string) bool {  return fpath==req.Ffp || (allfiles && ustr.Pref(fpath, srcDir))  }
		for _,mem := range gd.Package.Members {
			if srcref := udev.SrcMsgFromLn(mem.Pos)  ;  srcref!=nil {
				if fpathok(srcref.Ref) {
					mem.Type = devgo.ShortenImPs(mem.Type)  ;  srcref.Msg = mem.Kind + " " + mem.Name  ;  srcref.Flag = z.SYM_PACKAGE
					if mem.Kind=="const" {  srcref.Flag = z.SYM_CONSTANT  ;  srcref.Misc = "= " + mem.Value }
					if mem.Kind=="var" {  srcref.Flag = z.SYM_VARIABLE  ;  srcref.Misc = mem.Type }
					if mem.Kind=="func" {
						srcref.Flag = z.SYM_FUNCTION
						fargs,fret := fbreak(mem.Type)
						srcref.Misc = fret
						srcref.Msg = srcref.Msg + " " + strings.TrimPrefix(fargs, "func")
					}
					if mem.Kind=="type" {
						srcref.Misc = mem.Type  ;  srcref.Flag = z.SYM_CLASS
						if ustr.Pref(mem.Type, "struct{") { srcref.Flag = z.SYM_STRUCT }
						if ustr.Pref(mem.Type, "interface{") { srcref.Flag = z.SYM_INTERFACE }
						if ustr.Pref(mem.Type, "func(") { srcref.Flag = z.SYM_CONSTRUCTOR }
						if ustr.Pref(mem.Type, "[]") { srcref.Flag = z.SYM_ARRAY }
						if ustr.Pref(mem.Type, "map[") { srcref.Flag = z.SYM_NAMESPACE }
						if ustr.Pref(mem.Type, "*") { srcref.Flag = z.SYM_NULL }
						if ustr.AnyOf(mem.Type, "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "float", "complex") { srcref.Flag = z.SYM_NUMBER }
						if ustr.AnyOf(mem.Type, "string", "rune") { srcref.Flag = z.SYM_STRING }
						switch mem.Type {
						case "bool": srcref.Flag = z.SYM_BOOLEAN
						}
					}
					srcrefs = append(srcrefs, srcref)
				}
				for _,method := range mem.Methods { if mref := udev.SrcMsgFromLn(method.Pos)  ;  mref!=nil && fpathok(mref.Ref) {
					p1 , p2 := ustr.Idx(method.Name, " (") , ustr.Idx(method.Name, ") ")
					mref.Msg = method.Name[:p2][p1+2:] + "·" + method.Name[p2+2:] ;  mref.Flag = z.SYM_METHOD
					mref.Msg,mref.Misc = fbreak(mref.Msg)  ;  if i := ustr.Idx(mref.Msg, "(")  ;  i>0 {  mref.Msg = mref.Msg[:i] + " " + mref.Msg[i:]  }
					srcrefs = append(srcrefs, mref)
				} }
			}
		}
		if allfiles { sort.Sort(srcrefs)  ;  for _,srcref := range srcrefs {
			srcref.Msg = "[ " + strings.TrimLeft(srcref.Ref[len(srcDir):], "/\\") + " ]\t\t" + srcref.Msg
		} }
	} }
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
