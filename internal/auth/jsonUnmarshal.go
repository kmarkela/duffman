package auth

import "encoding/json"

type Auth struct {
	Type    string            `json:"Type"`
	Details map[string]string `json:"Details"`
}

type KVT struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// type AuthType int

// func (at AuthType) String() string {

// }

// const (
// 	None AuthType = iota
// 	Oauth2
// 	Bearer
// )

func (a *Auth) UnmarshalJSON(data []byte) error {

	var temp struct {
		Type   string `json:"type"`
		Oauth2 []KVT  `json:"oauth2"`
		Bearer []KVT  `json:"bearer"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return nil
	}

	switch temp.Type {
	case "oauth2":

	default:
	}

	*a = Auth{Type: temp.Type}

	return nil

}

// func (a Auth) Get() (string, []KeyValueType, error) {

// 	switch a.Type {
// 	case "bearer":
// 		return "Bearer", a.Bearer, nil
// 	case "oauth2":
// 		return "Oauth2", a.Oauth2, nil
// 	}

// 	return "None", []KeyValueType{}, fmt.Errorf("Auth method (%s) is not implemented", a.Type)
// }
