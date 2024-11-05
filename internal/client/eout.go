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
	Auth      authOut                `json:"Auth,omitempty"`
}

type authOut struct {
	Type    string                     `json:"Type,omitempty"`
	Details []pcollection.KeyValueType `json:"Details,omitempty"`
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

func buildVarStr(col pcollection.Collection, rp pcollection.Req) string {

	var vo varOut

	vo.Variables = col.Variables
	vo.Env = col.Env

	authType, det, err := rp.Auth.Get()
	if err != nil {
		return err.Error()
	}
	vo.Auth = authOut{Type: authType, Details: det}

	marshaled, err := json.MarshalIndent(vo, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(marshaled)

}
