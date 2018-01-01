package z

import (
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
)

// All public fields are 1-based (so 0 means 'missing') and rune-not-byte-based
type SrcPos struct {
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
	Off int `json:"o,omitempty"`

	// if & when this is computed, it'll be 0-based
	byteOff int
	byteoff bool
}

func (oneRangeStart *SrcPos) ComesBehind(anotherRangeEnd *SrcPos) bool {
	return oneRangeStart.Off >= anotherRangeEnd.Off ||
		oneRangeStart.byteOff >= anotherRangeEnd.byteOff ||
		oneRangeStart.Ln >= anotherRangeEnd.Ln ||
		(oneRangeStart.Ln == anotherRangeEnd.Ln && oneRangeStart.Col >= anotherRangeEnd.Col)
}

func (me *SrcPos) EquivTo(pos *SrcPos) bool {
	if me.byteoff && pos.byteoff {
		return me.byteOff == pos.byteOff
	} else if me.Off > 0 && pos.Off > 0 {
		return me.Off == pos.Off
	}
	return me.Ln == pos.Ln && me.Col == pos.Col
}

func (me *SrcPos) IsBetween(sr *SrcRange) bool {
	if sr.IsEmpty() {
		return me.EquivTo(&sr.Start)
	}
	return me.ComesBehind(&sr.Start) && sr.End.ComesBehind(me)
}

func (me *SrcPos) String() string {
	if me.Ln > 0 && me.Col > 0 {
		return Strf("%d,%d", me.Ln, me.Col)
	}
	return Strf("#%d", me.Off-1)
}

type SrcRange struct {
	Start SrcPos `json:"s"`
	End   SrcPos `json:"e,omitempty"`
}

func (me *SrcRange) IsEmpty() bool {
	return me.Start.EquivTo(&me.End) || (me.End.Col == 0 && me.End.Ln == 0 && me.End.Off == 0)
}

func (me *SrcRange) OverlapsWith(sr *SrcRange) bool {
	return (!me.Start.EquivTo(&sr.End) || me.End.EquivTo(&sr.Start)) &&
		(me.Start.IsBetween(sr) || me.End.IsBetween(sr) || sr.Start.IsBetween(me) || sr.End.IsBetween(me))
}

type SrcLocs []*SrcLoc

func (me *SrcLocs) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (loc *SrcLoc) {
	if srcRefLoc != nil {
		loc = &SrcLoc{}
		loc.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		(*me) = append(*me, loc)
	}
	return
}

type SrcLoc struct {
	Flag     int       `json:"e"` // don't omitempty
	FilePath string    `json:"f,omitempty"`
	Pos      *SrcPos   `json:"p,omitempty"`
	Range    *SrcRange `json:"r,omitempty"`
}

type SrcLenses []*SrcLens

func (me *SrcLenses) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (lens *SrcLens) {
	if srcRefLoc != nil {
		lens = &SrcLens{}
		lens.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		(*me) = append(*me, lens)
	}
	return
}

type SrcLens struct {
	SrcLoc
	Txt  string `json:"t,omitempty"`
	Str  string `json:"s,omitempty"`
	CrLf bool   `json:"l,omitempty"`
}

func (me *SrcLens) EnsureSrcFull() {
	if me.Txt == "" {
		me.Txt = ufs.ReadTextFile(me.FilePath, true, "")
	}
}

func (me *SrcLens) ByteOffsetForPosWithRuneOffset(pos *SrcPos) int {
	if !pos.byteoff {
		pos.byteoff = true
		if pos.Off > 1 {
			me.EnsureSrcFull()
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
	if strings.HasPrefix(me.Txt, prefix) {
		return 0
	} else if idx := strings.Index(me.Txt, "\n"+prefix); idx >= 0 {
		return len([]byte(me.Txt[:idx+1])) // want byte-pos not rune-pos
	}
	return -1
}

func (me *SrcLoc) SetFilePathAndPosOrRangeFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	me.SetFilePathFrom(srcRef, fallbackFilePath)
	me.SetPosOrRangeFrom(srcRef, true)
}

func (me *SrcLoc) SetFilePathFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	if me.FilePath = srcRef.Ref; me.FilePath != "" && !filepath.IsAbs(me.FilePath) {
		if absfilepath, err := filepath.Abs(me.FilePath); err == nil {
			me.FilePath = absfilepath
		} else if fallbackFilePath != nil {
			me.FilePath = fallbackFilePath()
		}
	}
	if (fallbackFilePath != nil) && (me.FilePath == "" || !ufs.FileExists(me.FilePath)) {
		me.FilePath = fallbackFilePath()
	}
}

func (me *SrcLoc) SetPosOrRangeFrom(srcRef *udev.SrcMsg, preferRange bool) {
	me.Pos, me.Range = nil, nil
	if preferRange && srcRef.Pos2Ch > 0 && srcRef.Pos2Ln > 0 && srcRef.Pos1Ch > 0 && srcRef.Pos1Ln > 0 {
		me.Range = &SrcRange{Start: SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch},
			End: SrcPos{Ln: srcRef.Pos2Ln, Col: srcRef.Pos2Ch}}
	} else {
		me.Pos = &SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch}
	}
}
