package z

type SrcLoc struct {
	FilePath string `json:"fp,omitempty"`
	SrcFull  string `json:"sf,omitempty"`
	SrcSel   string `json:"ss,omitempty"`
	PosOff   int    `json:"po,omitempty"`
	PosLn    int    `json:"pl,omitempty"`
	PosCol   int    `json:"pc,omitempty"`
}
