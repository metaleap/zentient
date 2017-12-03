package z

import (
	"github.com/metaleap/go-util/fs"
)

type SrcLens struct {
	FilePath   string  `json:"fp,omitempty"`
	SrcFull    string  `json:"sf,omitempty"`
	SrcSel     string  `json:"ss,omitempty"`
	Pos        *SrcPos `json:"p,omitempty"`
	RangeStart *SrcPos `json:"r0,omitempty"`
	RangeEnd   *SrcPos `json:"r1,omitempty"`
}

func (me *SrcLens) ensureSrcFull() {
	if me.SrcFull == "" {
		me.SrcFull = ufs.ReadTextFile(me.FilePath, true, "")
	}
}

func (me *SrcLens) ByteOffsetForPosWithRuneOffset(pos *SrcPos) int {
	if !pos.byteoff {
		pos.byteoff = true
		if pos.Off > 1 {
			me.ensureSrcFull()
			r := 1
			for i, _ := range me.SrcFull {
				if r == pos.Off {
					pos.byteOff = i
					return pos.byteOff
				}
				r++
			}
			pos.byteOff = len([]byte(me.SrcFull))
		}
	}
	return pos.byteOff
}

// All public fields are 1-based (so 0 means 'missing') and rune-not-byte-based
type SrcPos struct {
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
	Off int `json:"o,omitempty"`

	// if & when this is computed, it'll be 0-based
	byteOff int
	byteoff bool
}