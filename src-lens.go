package z

import (
	"path/filepath"
	"unicode/utf8"

	"github.com/go-leap/dev"
	"github.com/go-leap/dev/lex"
	"github.com/go-leap/fs"
	"github.com/go-leap/str"
)

func (me *SrcPos) SetRune1OffFromByte0Off(byte0Off int, src []byte) {
	if byte0Off > len(src) {
		byte0Off = len(src)
	} else {
		me.byteoff = true
	}
	me.byteOff = byte0Off
	me.Off = 1 + utf8.RuneCount(src[:byte0Off])
}

func (me *SrcPos) isBetween(sr *SrcRange) bool {
	return (sr.isEmpty() && me.isEquivTo(&sr.Start)) ||
		me.isSameOrGreaterThan(&sr.Start) && sr.End.isSameOrGreaterThan(me)
}

func (me *SrcPos) isEquivTo(pos *SrcPos) bool {
	if me.Off > 0 && pos.Off > 0 {
		return me.Off == pos.Off
	}
	return me.Ln == pos.Ln && me.Col == pos.Col
}

func (me *SrcPos) isSameOrGreaterThan(pos *SrcPos) bool {
	if me.Off > 0 && pos.Off > 0 {
		return me.Off >= pos.Off
	}
	return me.Ln > pos.Ln ||
		(me.Ln == pos.Ln && me.Col >= pos.Col)
}

func (me *SrcPos) String() string {
	if me.Ln > 0 && me.Col > 0 {
		return Strf("%d,%d", me.Ln, me.Col)
	}
	return Strf("#%d", me.Off-1)
}

func (me *SrcRange) isEmpty() bool {
	return me.Start.isEquivTo(&me.End) || (me.End.Col == 0 && me.End.Ln == 0 && me.End.Off == 0)
}

func (me *SrcRange) overlapsWith(sr *SrcRange) bool {
	if is0me, is0sr := me.isEmpty(), sr.isEmpty(); is0me && is0sr {
		return me.Start.isEquivTo(&sr.Start)
	} else if is0me {
		return me.Start.isBetween(sr)
	} else if is0sr {
		return sr.Start.isBetween(me)
	}
	return (!(me.Start.isEquivTo(&sr.End) || me.End.isEquivTo(&sr.Start))) &&
		(me.Start.isBetween(sr) || me.End.isBetween(sr) || sr.Start.isBetween(me) || sr.End.isBetween(me))
}

func (me *SrcLocs) Add(srcFilePath string, pos *udevlex.Pos) (loc *SrcLoc) {
	loc = &SrcLoc{FilePath: srcFilePath, Pos: &SrcPos{Ln: pos.Ln1, Col: pos.Col1}}
	*me = append(*me, loc)
	return
}

func (me *SrcLocs) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (loc *SrcLoc) {
	if srcRefLoc != nil {
		loc = &SrcLoc{}
		loc.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		*me = append(*me, loc)
	}
	return
}

func (me SrcLenses) Len() int          { return len(me) }
func (me SrcLenses) Swap(i int, j int) { me[i], me[j] = me[j], me[i] }
func (me SrcLenses) Less(i int, j int) bool {
	return !me[i].Pos.isSameOrGreaterThan(me[j].Pos)
}

func (me *SrcLenses) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (lens *SrcLens) {
	if srcRefLoc != nil {
		lens = &SrcLens{}
		lens.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		*me = append(*me, lens)
	}
	return
}

func (me *SrcLens) EnsureSrcFull() {
	if me.Txt == "" {
		me.Txt = ufs.ReadTextFileOr(me.FilePath, "")
	}
}

func (me *SrcLens) ByteOffsetForPos(pos *SrcPos) int {
	if (!pos.byteoff) && (pos.Off > 0 || (pos.Ln > 0 && pos.Col > 0)) {
		pos.byteoff = true
		if pos.Off > 1 || pos.Col > 1 || pos.Ln > 1 {
			me.EnsureSrcFull()
		}
		if pos.Off == 0 {
			if pos.Col == 1 && pos.Ln == 1 {
				pos.Off = 1
			} else {
				ln := 1
				for _, r := range me.Txt {
					if ln == pos.Ln {
						break
					} else if r == '\n' {
						ln++
					}
					pos.Off++
				}
				pos.Off += pos.Col
			}
		}
		if pos.Off > 1 {
			r := 1
			for i := range me.Txt {
				if r == pos.Off {
					pos.byteOff = i
					return pos.byteOff
				}
				r++
			}
			pos.byteOff = len([]byte(me.Txt))
		}
	}
	return pos.byteOff
}

func (me *SrcLens) ByteOffsetForFirstLineBeginningWith(prefix string) int {
	if ustr.Pref(me.Txt, prefix) {
		return 0
	} else if idx := ustr.Pos(me.Txt, "\n"+prefix); idx >= 0 {
		return len([]byte(me.Txt[:idx+1])) // want byte-pos not rune-pos
	}
	return -1
}

func (me *SrcLens) Rune1OffsetForByte0Offset(byte0off int) (rune1off int) {
	return 1 + utf8.RuneCountInString(me.Txt[:byte0off])
}

func (me *SrcLoc) SetFilePathAndPosOrRangeFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	me.setFilePathFrom(srcRef, fallbackFilePath)
	me.setPosOrRangeFrom(srcRef, true)
}

func (me *SrcLoc) setFilePathFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	if me.FilePath = srcRef.Ref; me.FilePath != "" && !filepath.IsAbs(me.FilePath) {
		if absfilepath, err := filepath.Abs(me.FilePath); err == nil {
			me.FilePath = absfilepath
		} else if fallbackFilePath != nil {
			me.FilePath = fallbackFilePath()
		}
	}
	if (fallbackFilePath != nil) && (me.FilePath == "" || !ufs.IsFile(me.FilePath)) {
		me.FilePath = fallbackFilePath()
	}
}

func (me *SrcLoc) setPosOrRangeFrom(srcRef *udev.SrcMsg, preferRange bool) {
	me.Pos, me.Range = nil, nil
	if preferRange && srcRef.Pos2Ch > 0 && srcRef.Pos2Ln > 0 && srcRef.Pos1Ch > 0 && srcRef.Pos1Ln > 0 {
		me.Range = &SrcRange{Start: SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch},
			End: SrcPos{Ln: srcRef.Pos2Ln, Col: srcRef.Pos2Ch}}
	} else {
		me.Pos = &SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch}
	}
}
