package pcollection

func (i *Item) i2ReqLt() ([]Req, error) {
	var lr []Req

	if i.Request != nil {
		lr = append(lr, buildReq(i.Request))
	}

	if i.Item == nil {
		return lr, nil
	}

	for _, v := range i.Item {
		tr, err := v.i2ReqLt()
		if err != nil {
			return nil, err
		}

		lr = append(lr, tr...)
	}

	return lr, nil
}
