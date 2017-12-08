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
	Icon   string
	Title  string `json:",omitempty"`
	Status struct {
		Flag CaddyStatus
		Desc string `json:",omitempty"`
	}
	OnReady         func() `json:"-"`
	OnStatusChanged func() `json:"-"`
}

func (me *Caddy) onInit() {
	me.Status.Flag, me.Status.Desc = CADDY_PENDING, "pending"
	me.OnStatusChanged = me.onStatusChanged
}

func (me *Caddy) onStatusChanged() {
	send(&msgResp{CaddyUpdate: me})
}