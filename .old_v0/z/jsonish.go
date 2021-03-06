package z

import (
	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/str"
)

const (
	DIAG_SEV_ERR  = 0
	DIAG_SEV_WARN = 1
	DIAG_SEV_INFO = 2
	DIAG_SEV_HINT = 3
)

const (
	SYM_FILE          = 0
	SYM_MODULE        = 1
	SYM_NAMESPACE     = 2
	SYM_PACKAGE       = 3
	SYM_CLASS         = 4
	SYM_METHOD        = 5
	SYM_PROPERTY      = 6
	SYM_FIELD         = 7
	SYM_CONSTRUCTOR   = 8
	SYM_ENUM          = 9
	SYM_INTERFACE     = 10
	SYM_FUNCTION      = 11
	SYM_VARIABLE      = 12
	SYM_CONSTANT      = 13
	SYM_STRING        = 14
	SYM_NUMBER        = 15
	SYM_BOOLEAN       = 16
	SYM_ARRAY         = 17
	SYM_OBJECT        = 18
	SYM_KEY           = 19
	SYM_NULL          = 20
	SYM_ENUMMEMBER    = 21
	SYM_STRUCT        = 22
	SYM_EVENT         = 23
	SYM_OPERATOR      = 24
	SYM_TYPEPARAMETER = 25
)

const (
	CMPL_TEXT          = 0
	CMPL_METHOD        = 1
	CMPL_FUNCTION      = 2
	CMPL_CONSTRUCTOR   = 3
	CMPL_FIELD         = 4
	CMPL_VARIABLE      = 5
	CMPL_CLASS         = 6
	CMPL_INTERFACE     = 7
	CMPL_MODULE        = 8
	CMPL_PROPERTY      = 9
	CMPL_UNIT          = 10
	CMPL_VALUE         = 11
	CMPL_ENUM          = 12
	CMPL_KEYWORD       = 13
	CMPL_SNIPPET       = 14
	CMPL_COLOR         = 15
	CMPL_FILE          = 16
	CMPL_REFERENCE     = 17
	CMPL_FOLDER        = 18
	CMPL_ENUMMEMBER    = 19
	CMPL_CONSTANT      = 20
	CMPL_STRUCT        = 21
	CMPL_EVENT         = 22
	CMPL_OPERATOR      = 23
	CMPL_TYPEPARAMETER = 24
)

type RespIntel struct {
	Label string `json:"label,omitempty"`
	Doc   string `json:"documentation,omitempty"`
}

type RespCmpl struct {
	RespIntel
	Kind        int      `json:"kind,"` // CMPL_FOO
	Detail      string   `json:"detail,omitempty"`
	SortTxt     string   `json:"sortText,omitempty"`
	FilterTxt   string   `json:"filterText,omitempty"`
	InsertTxt   string   `json:"insertText,omitempty"`
	CommitChars []string `json:"commitCharacters,omitempty"`
}

type RespCmd struct {
	Name string   `json:",omitempty"` //	actual cmd name
	Args []string `json:",omitempty"` //	args

	Title  string `json:",omitempty"` //	display name, eg: N = "go vet" when C = "go" with A = ["vet"]  ;  if empty fall back to C
	Exists bool   `json:",omitempty"` //	installed?
	Hint   string `json:",omitempty"` //	install hint
	More   string `json:",omitempty"`

	f func() //	tmp field used in Base.DoFmt()
}

type RespPick struct {
	Label  string `json:"label,omitempty"`
	Desc   string `json:"description,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type RespTxt struct {
	Result   string   `json:",omitempty"`
	Id       string   `json:",omitempty"`
	Warnings []string `json:",omitempty"`
}

type RespHov struct {
	Txt string `json:"value,omitempty"`
}

type ReqIntel struct {
	Ffp  string `json:",omitempty"`
	Pos  string `json:",omitempty"`
	Src  string `json:",omitempty"`
	Sym1 string `json:",omitempty"`
	Sym2 string `json:",omitempty"`
	Pos1 string `json:",omitempty"`
	Pos2 string `json:",omitempty"`
	Id   string `json:",omitempty"`

	r2b_ bool
}

func (me *ReqIntel) EnsureSrc() {
	if len(me.Src) == 0 {
		me.Src = ufs.ReadTextFile(me.Ffp, false, "")
	}
}

func (me *ReqIntel) RunePosToBytePos() {
	if !me.r2b_ {
		me.EnsureSrc()
		srcraw := []byte(me.Src)
		reoff := func(off int) int {
			r := 0
			for i, _ := range me.Src {
				if r == off {
					return i
				}
				{
					r++
				}
			}
			return len(srcraw)
		}
		rpos2bpos := func(off int, posfield *string) {
			if off > 0 && len(me.Src) > 0 {
				*posfield = umisc.Str(reoff(off))
			}
		}
		rpos2bpos(ustr.ToInt(me.Pos), &me.Pos)
		rpos2bpos(ustr.ToInt(me.Pos1), &me.Pos1)
		rpos2bpos(ustr.ToInt(me.Pos2), &me.Pos2)
		me.r2b_ = true
	}
}

var newlivediags = true

func jsonLiveDiags(frpszid string, closedfrps []string, openedfrps []string) (jld map[string]map[string]udev.SrcMsgs) {
	if len(closedfrps) > 0 || len(openedfrps) > 0 {
		newlivediags = true
	}
	if newlivediags {
		diagsready := true
		jld = map[string]map[string]udev.SrcMsgs{}
		var fc, fo []string
		for zid, µ := range Zengines {
			if !µ.ReadyToBuildAndLint() {
				diagsready = false
			}
			if zid == frpszid {
				fc, fo = closedfrps, openedfrps
			} else {
				fc, fo = nil, nil
			}
			jld[zid] = µ.B().liveDiags(µ, fc, fo)
		}
		if diagsready {
			newlivediags = false
		}
	}
	return // if diags haven't changed since last req, send nil
}

// the ONLY jsonish func to return a string-encoded-as-JSON-value
// thereby establishing convention/protocol for clients:
// if the response is such, it's to be interpreted as a reportable error
func jsonErrMsg(msg string) interface{} {
	return msg
}

func jsonStatus() interface{} {
	resp := map[string]interface{}{}
	// resp["livediags"] = Zengines["go"].B().livediags
	// resp["lintdiags"] = Zengines["go"].B().lintdiags
	// resp["builddiags"] = Zengines["go"].B().builddiags
	resp["Ctx"] = Ctx
	resp["OpenFiles"] = OpenFiles
	resp["AllFiles"] = AllFiles
	// resp["Zengines"] = jsonZengines()
	// for zid, zengine := range Zengines {
	// 	resp["Zengines["+zid+"]"] = zengine
	// }
	return resp
}

func jsonZengines() interface{} {
	list := map[string][]string{} // ouch =)
	for zid, zengine := range Zengines {
		list[zid] = zengine.EdLangIDs()
	}
	return list
}
