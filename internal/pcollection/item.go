package pcollection

import (
	"github.com/kmarkela/duffman/internal/auth"
)

func (i *Item) i2ReqLt(n *Node, auth *auth.Auth) ([]Req, error) {
	var lr []Req

	if i.Auth != nil {
		auth = i.Auth
	} else {
		i.Auth = auth
	}

	if i.Request != nil {
		r := buildReq(i.Request)
		n.Req = &r

		if i.Request.Auth != nil {
			r.Auth = i.Request.Auth
		} else {
			r.Auth = auth
		}

		lr = append(lr, r)
	}

	if i.Item == nil {
		return lr, nil
	}
	for _, v := range i.Item {

		tn := Node{}
		tn.Name = v.Name
		tr, err := v.i2ReqLt(&tn, i.Auth)
		if err != nil {
			return nil, err
		}

		lr = append(lr, tr...)
		n.Node = append(n.Node, tn)

	}

	return lr, nil
}
