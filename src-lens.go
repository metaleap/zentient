package z

import (
	"path/filepath"

	"github.com/go-leap/dev"
	"github.com/go-leap/fs"
	"github.com/go-leap/str"
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

func (this *SrcPos) isBetween(sr *SrcRange) bool {
	return (sr.isEmpty() && this.isEquivTo(&sr.Start)) ||
		this.isSameOrGreaterThan(&sr.Start) && sr.End.isSameOrGreaterThan(this)
}

func (this *SrcPos) isEquivTo(pos *SrcPos) bool {
	if this.Off > 0 && pos.Off > 0 {
		return this.Off == pos.Off
	}
	return this.Ln == pos.Ln && this.Col == pos.Col
}

func (this *SrcPos) isSameOrGreaterThan(pos *SrcPos) bool {
	if this.Off > 0 && pos.Off > 0 {
		return this.Off >= pos.Off
	}
	return this.Ln > pos.Ln ||
		(this.Ln == pos.Ln && this.Col >= pos.Col)
}

func (this *SrcPos) String() string {
	if this.Ln > 0 && this.Col > 0 {
		return Strf("%d,%d", this.Ln, this.Col)
	}
	return Strf("#%d", this.Off-1)
}

type SrcRange struct {
	Start SrcPos `json:"s"`
	End   SrcPos `json:"e,omitempty"`
}

func (this *SrcRange) isEmpty() bool {
	return this.Start.isEquivTo(&this.End) || (this.End.Col == 0 && this.End.Ln == 0 && this.End.Off == 0)
}

func (this *SrcRange) overlapsWith(sr *SrcRange) bool {
	if is0me, is0sr := this.isEmpty(), sr.isEmpty(); is0me && is0sr {
		return this.Start.isEquivTo(&sr.Start)
	} else if is0me {
		return this.Start.isBetween(sr)
	} else if is0sr {
		return sr.Start.isBetween(this)
	}
	return (!(this.Start.isEquivTo(&sr.End) || this.End.isEquivTo(&sr.Start))) &&
		(this.Start.isBetween(sr) || this.End.isBetween(sr) || sr.Start.isBetween(this) || sr.End.isBetween(this))
}

type SrcLocs []*SrcLoc

func (this *SrcLocs) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (loc *SrcLoc) {
	if srcRefLoc != nil {
		loc = &SrcLoc{}
		loc.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		*this = append(*this, loc)
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

func (this *SrcLenses) AddFrom(srcRefLoc *udev.SrcMsg, fallbackFilePath func() string) (lens *SrcLens) {
	if srcRefLoc != nil {
		lens = &SrcLens{}
		lens.SetFilePathAndPosOrRangeFrom(srcRefLoc, fallbackFilePath)
		*this = append(*this, lens)
	}
	return
}

type SrcLens struct {
	SrcLoc
	Txt  string `json:"t,omitempty"`
	Str  string `json:"s,omitempty"`
	CrLf bool   `json:"l,omitempty"`
}

func (this *SrcLens) EnsureSrcFull() {
	if this.Txt == "" {
		this.Txt = ufs.ReadTextFileOr(this.FilePath, "")
	}
}

func (this *SrcLens) ByteOffsetForPos(pos *SrcPos) int {
	if (!pos.byteoff) && (pos.Off > 0 || (pos.Ln > 0 && pos.Col > 0)) {
		pos.byteoff = true
		if pos.Off > 1 || pos.Col > 1 || pos.Ln > 1 {
			this.EnsureSrcFull()
		}
		if pos.Off == 0 {
			if pos.Col == 1 && pos.Ln == 1 {
				pos.Off = 1
			} else {
				ln := 1
				for _, r := range this.Txt {
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
			for i := range this.Txt {
				if r == pos.Off {
					pos.byteOff = i
					return pos.byteOff
				}
				r++
			}
			pos.byteOff = len([]byte(this.Txt))
		}
	}
	return pos.byteOff
}

func (this *SrcLens) ByteOffsetForFirstLineBeginningWith(prefix string) int {
	if ustr.Pref(this.Txt, prefix) {
		return 0
	} else if idx := ustr.Pos(this.Txt, "\n"+prefix); idx >= 0 {
		return len([]byte(this.Txt[:idx+1])) // want byte-pos not rune-pos
	}
	return -1
}

func (this *SrcLens) Rune1OffsetForByte0Offset(byte0off int) (rune1off int) {
	return 1 + ustr.NumRunes(this.Txt[:byte0off])
	// for byteoff := range this.Txt {
	// 	rune1off++
	// 	if byteoff >= byte0off {
	// 		return
	// 	}
	// }
	// return
}

func (this *SrcLoc) SetFilePathAndPosOrRangeFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	this.setFilePathFrom(srcRef, fallbackFilePath)
	this.setPosOrRangeFrom(srcRef, true)
}

func (this *SrcLoc) setFilePathFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	if this.FilePath = srcRef.Ref; this.FilePath != "" && !filepath.IsAbs(this.FilePath) {
		if absfilepath, err := filepath.Abs(this.FilePath); err == nil {
			this.FilePath = absfilepath
		} else if fallbackFilePath != nil {
			this.FilePath = fallbackFilePath()
		}
	}
	if (fallbackFilePath != nil) && (this.FilePath == "" || !ufs.IsFile(this.FilePath)) {
		this.FilePath = fallbackFilePath()
	}
}

func (this *SrcLoc) setPosOrRangeFrom(srcRef *udev.SrcMsg, preferRange bool) {
	this.Pos, this.Range = nil, nil
	if preferRange && srcRef.Pos2Ch > 0 && srcRef.Pos2Ln > 0 && srcRef.Pos1Ch > 0 && srcRef.Pos1Ln > 0 {
		this.Range = &SrcRange{Start: SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch},
			End: SrcPos{Ln: srcRef.Pos2Ln, Col: srcRef.Pos2Ch}}
	} else {
		this.Pos = &SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch}
	}
}
