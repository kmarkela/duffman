package pcollection

import (
	"github.com/kmarkela/duffman/internal/auth"
	"github.com/kmarkela/duffman/internal/internalTypes"
)

// filtered out version
type Collection struct {
	Variables []internalTypes.KeyValue
	Requests  []Req
	Env       []internalTypes.KeyValue
	Schema    Schema
}

type Req struct {
	Method      string            `json:"Method,omitempty"`
	URL         string            `json:"URL,omitempty"`
	Headers     map[string]string `json:"Headers,omitempty"`
	Body        string            `json:"Body,omitempty"`
	ContentType string            `json:"ContentType,omitempty"`
	Parameters  Parameters        `json:"-"`
	Auth        *auth.Auth        `json:"Auth"`
}

type Parameters struct {
	Get  map[string]string
	Post map[string]string
	Path map[string]string
}
