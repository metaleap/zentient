package z

type CaddyStatus uint8

const (
	CADDY_PENDING CaddyStatus = iota
	CADDY_ERROR
	CADDY_BUSY
	CADDY_GOOD
)

type Caddy struct {
	ID     string `json:",omitempty"`
	LangID string `json:",omitempty"`
	Icon   string
	Title  string `json:",omitempty"`
	Status struct {
		Flag CaddyStatus
		Desc string `json:",omitempty"`
	}
	Details                 string `json:",omitempty"`
	UxActionID              string `json:",omitempty"`
	ShowTitle               bool   `json:",omitempty"`
	ShouldReRunWhenNextDone bool   `json:"-"`

	ready bool

	OnReady func() `json:"-"`
}

func (me *Caddy) onInit() {
	me.Status.Flag, me.Status.Desc = CADDY_PENDING, "pending"
}

func (me *Caddy) OnStatusChanged() {
	go send(&ipcResp{CaddyUpdate: me})
}

func (me *Caddy) IsPendingOrBusy() bool {
	return me.Status.Flag == CADDY_BUSY || me.Status.Flag == CADDY_PENDING
}
func (me *Caddy) IsReady() bool {
	return me.ready
}
