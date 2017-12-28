package zgo

import (
	"path/filepath"
	"sort"
	"strings"
	"time"

	gurujson "golang.org/x/tools/cmd/guru/serial"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var (
	symsPatterns = map[string]z.Symbol{
		"*":          z.SYM_NULL,
		"map[":       z.SYM_OBJECT,
		"[]":         z.SYM_ARRAY,
		"func(":      z.SYM_EVENT,
		"interface{": z.SYM_INTERFACE,
		"struct{":    z.SYM_CLASS,
	}
)

func (*goSrcIntel) References(srcLens *z.SrcLens, includeDeclaration bool) (refs z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	bytepos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	if gr := udevgo.QueryRefs_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(bytepos)); len(gr) > 0 {
		refs = make(z.SrcLenses, 0, len(gr))
		for _, gref := range gr {
			if srcref := udev.SrcMsgFromLn(gref.Pos); srcref != nil {
				refs.AddFrom(srcref, nil)
			}
		}
	}
	return
}

func (*goSrcIntel) DefSym(srcLens *z.SrcLens) (defs z.SrcLenses) {
	var refloc *udev.SrcMsg
	bytepos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	spos := ustr.FromInt(bytepos)

	if refloc == nil && tools.godef.Installed {
		refloc = udevgo.QueryDefLoc_Godef(srcLens.FilePath, srcLens.Txt, spos)
	}
	if refloc == nil && tools.gogetdoc.Installed {
		refloc = udevgo.QueryDefLoc_Gogetdoc(srcLens.FilePath, srcLens.Txt, spos)
	}
	if refloc == nil && tools.guru.Installed {
		if gd, _ := udevgo.QueryDesc_Guru(srcLens.FilePath, srcLens.Txt, spos); gd != nil {
			if gd.Type != nil && len(gd.Type.NamePos) > 0 {
				if rl := udev.SrcMsgFromLn(gd.Type.NamePos); rl != nil {
					refloc = rl
				}
			}
			if gd.Value != nil && len(gd.Value.ObjPos) > 0 {
				if rl := udev.SrcMsgFromLn(gd.Value.ObjPos); rl != nil {
					refloc = rl
				}
			}
		}
	}
	if refloc != nil {
		defs.AddFrom(refloc, nil)
	}
	return
}

func (me *goSrcIntel) DefType(srcLens *z.SrcLens) (defs z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	var refloc *udev.SrcMsg
	bytepos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	spos := ustr.FromInt(bytepos)

	if gd, _ := udevgo.QueryDesc_Guru(srcLens.FilePath, srcLens.Txt, spos); gd != nil {
		if gd.Type != nil && len(gd.Type.NamePos) > 0 {
			if rl := udev.SrcMsgFromLn(gd.Type.NamePos); rl != nil {
				refloc = rl
			}
		}
		if refloc == nil && gd.Value != nil && len(gd.Value.Type) > 0 {
			// some hackery to adequately service 98+% of cases --- otherwise, no-go-to-type-def-for-you, not the end of the world
			for {
				if strings.HasPrefix(gd.Value.Type, "[]") {
					gd.Value.Type = gd.Value.Type[2:]
				} else if strings.HasPrefix(gd.Value.Type, "*") {
					gd.Value.Type = gd.Value.Type[1:]
				} else if strings.HasPrefix(gd.Value.Type, "map[") {
					gd.Value.Type = gd.Value.Type[strings.IndexRune(gd.Value.Type, ']')+1:]
				} else {
					break
				}
			}
			pkgimppath, typename := ustr.BreakOnLast(gd.Value.Type, ".")
			pkgname := ustr.AfterLast(pkgimppath, "/", false)
			if udevgo.PkgsByImP != nil {
				if pkg := udevgo.PkgsByImP[pkgimppath]; pkg != nil && len(pkg.Name) > 0 {
					pkgname = pkg.Name
				}
			}
			hacky1 := z.Strf("\n\nfunc Zentient%d () *", time.Now().UnixNano())
			hacky2 := " { return nil }\n"
			if len(pkgname) > 0 {
				hacky1 = hacky1 + pkgname + "."
			}
			srcLens.Pos = &z.SrcPos{Off: len(srcLens.Txt) + len(hacky1) + 1}
			srcLens.Txt = srcLens.Txt + hacky1 + typename + hacky2
			return me.DefSym(srcLens)
		}
	}
	if refloc != nil {
		defs.AddFrom(refloc, nil)
	}
	return
}

func (*goSrcIntel) DefImpl(srcLens *z.SrcLens) (defs z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	bytepos := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	if gi := udevgo.QueryImpl_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(bytepos)); gi != nil {
		if defs = make(z.SrcLenses, 0, len(gi.AssignableFrom)+len(gi.AssignableTo)+len(gi.AssignableFromPtr)+len(gi.AssignableFromMethod)+len(gi.AssignableFromPtrMethod)+len(gi.AssignableToMethod)); cap(defs) > 0 {
			addtypes := func(impltypes []gurujson.ImplementsType) {
				for _, it := range impltypes {
					if srcref := udev.SrcMsgFromLn(it.Pos); srcref != nil {
						defs.AddFrom(srcref, nil)
					}
				}
			}
			addmethods := func(methods []gurujson.DescribeMethod) {
				for _, m := range methods {
					if srcref := udev.SrcMsgFromLn(m.Pos); srcref != nil {
						defs.AddFrom(srcref, nil)
					}
				}
			}
			if gi.Method != nil {
				addmethods(gi.AssignableToMethod)      // "implements `"+gi.Method.Name+"`",
				addmethods(gi.AssignableFromMethod)    // "implemented by `"+gi.Method.Name+"`",
				addmethods(gi.AssignableFromPtrMethod) // "implemented by `"+gi.Method.Name+"`",
			} else {
				addtypes(gi.AssignableTo)      // "type implementing `"+gi.T.Name+"`",
				addtypes(gi.AssignableFrom)    // "type implemented by `"+gi.T.Name+"`",
				addtypes(gi.AssignableFromPtr) // "type implemented by `*"+gi.T.Name+"`",
			}
		}
	}
	return
}

func (*goSrcIntel) Highlights(srcLens *z.SrcLens, curWord string) (all z.SrcLenses) {
	if !tools.guru.Installed {
		return
	}
	byteoff := srcLens.ByteOffsetForPosWithRuneOffset(srcLens.Pos)
	gw, err := udevgo.QueryWhat_Guru(srcLens.FilePath, srcLens.Txt, ustr.FromInt(byteoff))
	if err != nil {
		panic(err)
	}
	all = make(z.SrcLenses, 0, len(gw.SameIDs))
	for _, sameid := range gw.SameIDs {
		if srcref := udev.SrcMsgFromLn(sameid); srcref != nil {
			if sl := all.AddFrom(srcref, nil); sl.FilePath == srcLens.FilePath && (sl.Range != nil || sl.Pos != nil) {
				sl.FilePath = ""
			}
		}
	}
	if len(all) == 0 && len(gw.Enclosing) > 0 {
		srcraw := []byte(srcLens.Txt)
		bpos2rpos := func(bytepos int) int {
			length := 0
			for _ = range string(srcraw[:bytepos]) {
				length++
			}
			return length // here's what *won't* work for multi-byte/unicode/etc: just bytepos, or even len(string(srcraw[:bytepos]))
		}
		var check func(num int, checks ...string) bool
		check = func(num int, checks ...string) bool {
			if num <= 0 {
				for _, chk := range checks {
					if check(1, chk, chk) {
						return true
					}
				}
				return false
			}
			if ustr.AnyOf(gw.Enclosing[0].Description, checks[:num]...) {
				for _, syntaxnode := range gw.Enclosing {
					if ustr.AnyOf(syntaxnode.Description, checks[num:]...) {
						all = append(all, &z.SrcLens{Range: &z.SrcRange{
							Start: z.SrcPos{Off: 1 + bpos2rpos(syntaxnode.Start)}, End: z.SrcPos{Off: 1 + bpos2rpos(syntaxnode.End)}}})
						return true
					}
				}
			}
			return false
		}
		if check(1, "break statement",
			"range loop", "for loop", "select statement", "switch statement") {
			return
		}
		if check(1, "case clause",
			"select statement", "switch statement") {
			return
		}
		if check(1, "continue statement",
			"range loop", "for loop") {
			return
		}
		if check(2, "defer statement", "return statement",
			"function literal", "function declaration") {
			return
		}
		check(-1, "if statement", "select statement", "switch statement", "go statement", "range loop", "for loop", "struct type", "interface type", "map type", "function type", "slice type", "type specification", "type declaration", "function declaration", "field/method/parameter", "field/method/parameter list", "function call", "function call (or conversion)", "basic literal", "composite literal", "variable declaration", "constant declaration")
	}
	return
}

func (me *goSrcIntel) Symbols(sL *z.SrcLens, query string, curFileOnly bool) (allsyms z.SrcLenses) {
	onerr := func(label string, detail string) z.SrcLenses {
		return z.SrcLenses{&z.SrcLens{Flag: int(z.SYM_EVENT), Str: label, Txt: detail, FilePath: sL.FilePath, Pos: sL.Pos, Range: sL.Range}}
	}
	if !tools.guru.Installed {
		return onerr(z.ToolsMsgGone("guru"), z.ToolsMsgMore("guru"))
	}
	sL.EnsureSrcFull()
	srclns := strings.Split(sL.Txt, "\n")
	bytepos := 8 + sL.ByteOffsetForFirstLineBeginningWith("package ")
	gd, err := udevgo.QueryDesc_Guru(sL.FilePath, sL.Txt, ustr.FromInt(bytepos))
	if err != nil {
		return onerr("Error running guru:", err.Error())
	} else if gd.Package == nil {
		return onerr("Error running guru:", "not in a Go package")
	}

	// no more early-returns, now get busy
	if !curFileOnly {
		sort.Sort(gd) // sort doesn't seem to help improve vsc's ctrl+t ux for now, but maybe in some future vsc release..
	}
	anyfilegoes, curpkgdir, numpms := !curFileOnly, filepath.Dir(sL.FilePath), len(gd.Package.Members)
	query, allsyms = strings.ToLower(query), make(z.SrcLenses, 0, numpms) // numpms will never be a 'good' cap in any of these, but hey any number beats the default cap of 0..
	for _, pm := range gd.Package.Members {
		ispmlisted := false
		if srcref := udev.SrcMsgFromLn(pm.Pos); srcref != nil && (anyfilegoes || srcref.Ref == sL.FilePath) && (query == "" || gd.Matches(pm, query)) {
			ispmlisted = true
			pmtype, sym := pm.Type, &z.SrcLens{Str: pm.Name}
			pmuntyped := strings.HasPrefix(pmtype, "untyped ")
			{
				if pmuntyped {
					pmtype = pmtype[8:]
				} else if strings.HasPrefix(pmtype, "struct{") {
					next := func() int { return strings.Index(pmtype, ` "json:\"`) }
					for ij1 := next(); ij1 > 0; ij1 = next() {
						if ij2 := strings.Index(pmtype[ij1+9:], `\""`); ij2 >= 0 {
							pref, suff := pmtype[:ij1], pmtype[ij1+9+3+ij2:]
							pmtype = pref + suff
						} else {
							break
						}
					}
				}
				pmtype = udevgo.PkgImpPathsToNamesInLn(pmtype, curpkgdir)
			}
			sym.SetFilePathAndPosOrRangeFrom(srcref, nil)
			switch pm.Kind {
			case "const":
				if sym.Flag, sym.Txt = int(z.SYM_CONSTANT), pmtype+" = "+pm.Value; !pmuntyped {
					sym.Str, sym.Flag = "▶   "+pm.Name, int(z.SYM_NUMBER)
				}
			case "var":
				sym.Flag, sym.Txt = int(z.SYM_VARIABLE), pmtype
			case "func":
				fnargs, fnret := me.symFuncSigBreak(pmtype)
				sym.Flag, sym.Txt = int(z.SYM_FUNCTION), udevgo.PkgImpPathsToNamesInLn(fnret, curpkgdir)
				sym.Str += "  " + udevgo.PkgImpPathsToNamesInLn(strings.TrimPrefix(fnargs, "func"), curpkgdir)
			case "type":
				sym.Txt, sym.Flag = pmtype, int(z.SYM_TYPEPARAMETER)
				switch pmtype {
				case "float32", "float64", "float", "complex64", "complex128", "int64", "uint64":
					sym.Flag = int(z.SYM_NUMBER)
				case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
					sym.Flag = int(z.SYM_ENUMMEMBER)
				case "rune", "string", "[]byte":
					sym.Flag = int(z.SYM_STRING)
				case "bool":
					sym.Flag = int(z.SYM_BOOLEAN)
				case "uintptr":
					sym.Flag = int(z.SYM_NULL)
				default:
					for pref, enum := range symsPatterns {
						if strings.HasPrefix(pmtype, pref) {
							sym.Flag = int(enum)
							break
						}
					}
				}
			default:
				z.BadPanic("guru.DescribeMember.Kind", pm.Kind)
			}
			allsyms = append(allsyms, sym)
		}
		for _, method := range pm.Methods {
			if isok := ispmlisted || !strings.HasPrefix(pm.Type, "interface{"); isok && (query == "" || strings.Contains(strings.ToLower(method.Name), query)) {
				if srcref := udev.SrcMsgFromLn(method.Pos); srcref != nil && (anyfilegoes || srcref.Ref == sL.FilePath) {
					p1, p2 := strings.Index(method.Name, " ("), strings.Index(method.Name, ") ")
					methodtype, methodtitle := method.Name[:p2][p1+2:], method.Name[p2+2:]
					methodname := strings.TrimSpace(methodtitle[:strings.Index(methodtitle, "(")])
					if strings.HasPrefix(pm.Type, "interface{") && !(strings.Contains(pm.Type, "{"+methodname+"(") || strings.Contains(pm.Type, "; "+methodname+"(")) {
						// guru reports an embedded interface's methods redundantly for each embedder, all pointing to the embeddee's original loc. we skip these
						continue
					}
					if curFileOnly && srcref.Ref == sL.FilePath && strings.HasPrefix(pm.Type, "struct{") && srcref.Pos1Ln > 0 && srcref.Pos1Ln <= len(srclns) {
						if srcln := srclns[srcref.Pos1Ln-1]; !strings.Contains(srcln, methodtype+") "+methodname+"(") {
							// guru reports an embedded struct's methods redundantly for each embedder, all pointing to the embeddee's original loc. we skip these --- at least in "file symbols" mode
							continue
						}
					}
					lens := allsyms.AddFrom(srcref, nil)
					lens.Flag, lens.Str = int(z.SYM_METHOD), "▶  "+methodtitle
					// if !ispmlisted { // if method's receiver type not in the symbols listing, prepend it's name to the pretend-indentation
					lens.Str = methodtype + "  " + lens.Str
					// }
					lens.Str, lens.Txt = me.symFuncSigBreak(lens.Str)
					lens.Str, lens.Txt = udevgo.PkgImpPathsToNamesInLn(lens.Str, curpkgdir), udevgo.PkgImpPathsToNamesInLn(lens.Txt, curpkgdir)
					if i := strings.Index(lens.Str, "("); i > 0 { // insert some spacing between name and args
						lens.Str = lens.Str[:i] + "  " + lens.Str[i:]
					}
				}
			}
		}
	}
	return
}

func (*goSrcIntel) symFuncSigBreak(fnsig string) (fnargs string, fnret string) {
	fnargs, fnret = fnsig, " "
	co, cc, pos := 0, 0, 0
	for i, r := range fnsig {
		if r == '(' {
			co++
		} else if r == ')' {
			cc++
		}
		if cc > 0 && co > 0 && cc == co {
			pos = i
			break
		}
	}
	if pos > 0 && pos < len(fnsig)-1 {
		fnargs, fnret = fnsig[:pos+1], fnsig[pos+2:]
	}
	return
}