package z

func handleMetaCmdsListAll(req *MsgReq, resp *MsgResp) {
	resp.Menu = &MsgRespMenu{Desc: "Choose wisely, mister:"}
	menu := resp.Menu.Choices
	menu = append(menu, &MsgRespPick{ID: 1, Title: "Choice One", Desc: "First choice.. First choice.. First choice.. First choice.. First choice.. First choice.. First choice.. ", Detail: "1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, 1 detail, "})
	menu = append(menu, &MsgRespPick{ID: 2, Title: "Choice Two", Desc: "Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. Second choice.. ", Detail: "2 details, 2 details, 2 details, 2 details, 2 details, 2 details, 2 details, 2 details, "})
	menu = append(menu, &MsgRespPick{ID: 3, Title: "Choice Tri", Desc: "Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. Third choice.. ", Detail: "3 details, 3 details, 3 details, 3 details, 3 details, 3 details, 3 details, 3 details, "})
	resp.Menu.Choices = menu
}
