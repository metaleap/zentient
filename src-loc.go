package z

type SrcLoc struct {
	FilePath string     `json:"fp,omitempty"`
	SrcFull  string     `json:"sf,omitempty"`
	SrcSel   string     `json:"ss,omitempty"`
	Pos0     *SrcLocPos `json:"p0,omitempty"`
	Pos1     *SrcLocPos `json:"p1,omitempty"`
}

type SrcLocPos struct {
	Off int `json:"o,omitempty"`
	Ln  int `json:"l,omitempty"`
	Col int `json:"c,omitempty"`
}
