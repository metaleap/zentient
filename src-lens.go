package z

type SrcLens struct {
	FilePath   string  `json:"fp,omitempty"`
	SrcFull    string  `json:"sf,omitempty"`
	SrcSel     string  `json:"ss,omitempty"`
	Pos        *SrcPos `json:"p,omitempty"`
	RangeStart *SrcPos `json:"r0,omitempty"`
	RangeEnd   *SrcPos `json:"r1,omitempty"`
}

// All fields are 1-based, so 0 means 'missing'
type SrcPos struct {
	Off int `json:"o,omitempty"`
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
}
