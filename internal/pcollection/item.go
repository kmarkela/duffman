package pcollection

func (i *Item) filter() ([]*Request, error) {
	var lr []*Request

	if i.Request != nil {
		lr = append(lr, i.Request)
	}

	if i.Item == nil {
		return lr, nil
	}

	for _, v := range i.Item {
		tr, err := v.filter()
		if err != nil {
			return nil, err
		}

		lr = append(lr, tr...)
	}

	return lr, nil
}
