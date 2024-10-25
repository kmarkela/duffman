package pcollection

import "fmt"

// type StringOrBool struct {
// 	Value interface{}
// }

type KeyValueType struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type AuthType int

const (
	None AuthType = iota
	Oauth2
	Bearer
)

type Auth struct {
	Type   string         `json:"type"`
	Oauth2 []KeyValueType `json:"oauth2"`
	Bearer []KeyValueType `json:"bearer"`
}

func (a Auth) Get() (AuthType, []KeyValueType, error) {

	switch a.Type {
	case "bearer":
		return Oauth2, a.Bearer, nil
	case "oauth2":
		return Bearer, a.Oauth2, nil
	}

	return None, []KeyValueType{}, fmt.Errorf("Auth method (%s) is not implemented", a.Type)
}
