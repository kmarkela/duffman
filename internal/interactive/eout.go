package interactive

import (
	"encoding/json"

	"github.com/kmarkela/duffman/internal/pcollection"
)

type varOut struct {
	Variables []pcollection.KeyValue `json:"Variables,omitempty"`
	Env       []pcollection.KeyValue `json:"Enviroment,omitempty"`
}

func buildReqStr(r pcollection.Req) string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(marshaled)

}

func buildVarStr(col pcollection.Collection) string {

	var vo varOut

	vo.Variables = col.Variables
	vo.Env = col.Env

	marshaled, err := json.MarshalIndent(vo, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(marshaled)

}
