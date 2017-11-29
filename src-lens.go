package z

type SrcLens struct {
	FilePath string  `json:"fp,omitempty"`
	SrcFull  string  `json:"sf,omitempty"`
	SrcSel   string  `json:"ss,omitempty"`
	Pos0     *SrcPos `json:"p0,omitempty"`
	Pos1     *SrcPos `json:"p1,omitempty"`
}

type SrcPos struct {
	Off int `json:"o,omitempty"`
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
}
