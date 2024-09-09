package pcollection

// filtered out version
type Collection struct {
	Variables []KeyValue
	Requests  []Req
	Env       []KeyValue
	Schema    Schema
}

type Req struct {
	Method      string            `json:"Method,omitempty"`
	URL         string            `json:"URL,omitempty"`
	Headers     map[string]string `json:"Headers,omitempty"`
	Body        string            `json:"Body,omitempty"`
	ContentType string            `json:"ContentType,omitempty"`
	Parameters  Parameters        `json:"-"`
}

type Parameters struct {
	Get  map[string]string
	Post map[string]string
	Path map[string]string
}
