package client

import (
	"bytes"
	"encoding/json"

	"github.com/kmarkela/duffman/internal/internalTypes"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/internal/req"
)

type varOut struct {
	Variables []internalTypes.KeyValue `json:"Variables,omitempty"`
	Env       []internalTypes.KeyValue `json:"Environment,omitempty"`
	// Auth      *auth.Auth               `json:"Auth,omitempty"`
}

func buildReqStr(rp pcollection.Req, env, vars []internalTypes.KeyValue) string {

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

	// if rp.Auth != nil {
	// 	auth := auth.ResolveVars(col.Env, col.Variables, rp.Auth)
	// 	vo.Auth = &auth

	// }

	marshaled, err := json.MarshalIndent(vo, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(marshaled)

}
