package auth

import (
	"encoding/json"
	"fmt"
)

type KVT struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type Auth struct {
	Type    string            `json:"Type"`
	Details map[string]string `json:"Details"`
}

type AuthType int

const (
	None AuthType = iota
	Oauth2
	Bearer
)

func (at AuthType) String() string {

	switch at {
	case None:
		return "None"
	case Oauth2:
		return "Oauth2"
	case Bearer:
		return "Bearer"
	}

	return "NotSupported"

}

func (a *Auth) UnmarshalJSON(data []byte) error {

	// check if alredy in custom format
	var tempPrFormat struct {
		Type    string            `json:"Type"`
		Details map[string]string `json:"Details"`
	}

	json.Unmarshal(data, &tempPrFormat)
	if tempPrFormat.Details != nil {
		*a = Auth{Type: tempPrFormat.Type, Details: tempPrFormat.Details}
		return nil
	}

	var temp struct {
		Type   string `json:"type"`
		Oauth2 []KVT  `json:"oauth2"`
		Bearer []KVT  `json:"bearer"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var details map[string]string

	switch temp.Type {
	case "oauth2":
		details = getDet(temp.Oauth2)
	case "bearer":
		details = getDet(temp.Bearer)

		// default:
	}

	*a = Auth{Type: temp.Type, Details: details}

	return nil

}

// func (a Auth) MarshalJSON() ([]byte, error) {
// 	type Alias Auth
// 	return json.Marshal(Alias(a))
// }

func getDet(k []KVT) map[string]string {

	result := map[string]string{}

	for _, el := range k {
		// convert value to string
		result[el.Key] = fmt.Sprintf("%v", el.Value)
	}

	return result
}
