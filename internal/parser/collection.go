package parser

type Collection struct {
	Items    []Item     `json:"item"`
	Variable []KeyValue `json:"variable,omitempty"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Item struct {
	Name    string   `json:"name"`
	Item    []Item   `json:"item,omitempty"`
	Request *Request `json:"request,omitempty"`
}

type Request struct {
	Method string     `json:"method"`
	Header []KeyValue `json:"header,omitempty"`
	Body   Body       `json:"body,omitempty"`
	URL    URL        `json:"url"`
}

type URL struct {
	Raw      string     `json:"raw"`
	Protocol string     `json:"protocol"`
	Host     []string   `json:"host"`
	Path     []string   `json:"path"`
	Query    []KeyValue `json:"query,omitempty"`
}

type Body struct {
	Mode       string     `json:"mode"`
	Raw        string     `json:"raw,omitempty"`
	FormData   []KeyValue `json:"formdata,omitempty"`
	URLEncoded []KeyValue `json:"urlencoded,omitempty"`
}
