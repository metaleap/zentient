package zgo
import (
	"github.com/metaleap/go-devgo"
	"github.com/metaleap/go-util-dev"
	"github.com/metaleap/go-util-misc"
	"github.com/metaleap/zentient/z"
)

func (self *zgo) IntelDefLoc (ffp string, srcin string, offset string) (*udev.SrcMsg, error) {
	if !devgo.Has_godef { return nil , ugo.E("No applicable Code Intel tools available.") }
	return devgo.QueryDefLoc_Godef(ffp, srcin, offset)
}


func (self *zgo) IntelHovs (ffp string, srcin string, offset string) (hovs []*z.RespHov) {
	if devgo.Has_godef { if defdecl := devgo.QueryDefDecl_GoDef(ffp, srcin, offset)  ;  len(defdecl)>0 {
		hovs = append(hovs, &z.RespHov { Lang: "go", Txt: defdecl })
	} }
	if len(hovs)==0 { hovs = append(hovs, &z.RespHov { Txt: "No applicable Code Intel tools available." }) }
	return
}
