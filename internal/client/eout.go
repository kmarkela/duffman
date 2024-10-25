package client

import (
	"bytes"
	"encoding/json"

	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/internal/req"
)

type varOut struct {
	Variables []pcollection.KeyValue `json:"Variables,omitempty"`
	Env       []pcollection.KeyValue `json:"Environment,omitempty"`
}

func buildReqStr(rp pcollection.Req, env, vars []pcollection.KeyValue) string {

	r := req.DeepCopyReq(&rp)

	req.ResolveVars(env, vars, r)
	r.URL = req.CreateEndpoint(r.URL, r.Parameters.Get, r.Parameters.Path)
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "   ")
	if err := encoder.Encode(r); err != nil {
		return err.Error()
	}
	return buf.String()

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
