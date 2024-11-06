package auth

import (
	"encoding/json"
	"fmt"

	"github.com/kmarkela/duffman/internal/internalTypes"
)

type Auth struct {
	Type    string            `json:"Type"`
	Details map[string]string `json:"Details"`
}

type AuthType int

const (
	None AuthType = iota
	Oauth2
	// Bearer
)

func (at AuthType) String() string {

	switch at {
	case None:
		return "None"
	case Oauth2:
		return "Oauth2"
		// case Bearer:
		// 	return "Bearer"
	}

	return "NotSupported"

}

func (a *Auth) UnmarshalJSON(data []byte) error {

	var temp struct {
		Type   string                   `json:"type"`
		Oauth2 []internalTypes.KeyValue `json:"oauth2"`
		// Bearer []KVT  `json:"bearer"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return nil
	}

	var details map[string]string

	switch temp.Type {
	case "oauth2":
		details = getDet(temp.Oauth2)
		// default:
	}

	*a = Auth{Type: temp.Type, Details: details}

	return nil

}

func getDet(k []internalTypes.KeyValue) map[string]string {

	result := map[string]string{}

	for _, el := range k {
		// convert value to string
		result[el.Key] = fmt.Sprintf("%v", el.Value)
	}

	return result
}
