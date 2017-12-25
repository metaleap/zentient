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

type SrcRange struct {
	Start SrcPos `json:"s"`
	End   SrcPos `json:"e,omitempty"`
}

type SrcLenses []*SrcLens

type SrcLens struct {
	FilePath string    `json:"f,omitempty"`
	Txt      string    `json:"t,omitempty"`
	Str      string    `json:"s,omitempty"`
	Pos      *SrcPos   `json:"p,omitempty"`
	Range    *SrcRange `json:"r,omitempty"`
	CrLf     bool      `json:"l,omitempty"`
	Flag     int       `json:"e"` // don't omitempty
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

func (me *SrcLens) SetFilePathAndPosOrRangeFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
	me.SetFilePathFrom(srcRef, fallbackFilePath)
	me.SetPosOrRangeFrom(srcRef, true)
}

func (me *SrcLens) SetFilePathFrom(srcRef *udev.SrcMsg, fallbackFilePath func() string) {
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

func (me *SrcLens) SetPosOrRangeFrom(srcRef *udev.SrcMsg, preferRange bool) {
	me.Pos, me.Range = nil, nil
	if preferRange && srcRef.Pos2Ch > 0 && srcRef.Pos2Ln > 0 {
		me.Range = &SrcRange{Start: SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch},
			End: SrcPos{Ln: srcRef.Pos2Ln, Col: srcRef.Pos2Ch}}
	} else {
		me.Pos = &SrcPos{Ln: srcRef.Pos1Ln, Col: srcRef.Pos1Ch}
	}
}
