package pcollection

func (i *Item) i2ReqLt(n *Node) ([]Req, error) {
	var lr []Req

	if i.Request != nil {
		r := buildReq(i.Request)
		n.Req = &r
		lr = append(lr, r)
	}

	if i.Item == nil {
		return lr, nil
	}

	for _, v := range i.Item {
		tn := Node{}
		// fmt.Println(v.Name)
		tn.Name = v.Name
		tr, err := v.i2ReqLt(&tn)
		if err != nil {
			return nil, err
		}

		lr = append(lr, tr...)
		n.Node = append(n.Node, tn)

	}

	return lr, nil
}
