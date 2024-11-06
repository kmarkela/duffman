package pcollection

import "github.com/kmarkela/duffman/internal/auth"

// The file contains structs to unmarshal Collection and Environment Json files.

type RawCollection struct {
	Info     Info       `json:"info"`
	Items    []Item     `json:"item"`
	Variable []KeyValue `json:"variable,omitempty"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Item struct {
	Name    string     `json:"name"`
	Item    []Item     `json:"item,omitempty"`
	Request *Request   `json:"request,omitempty"`
	Auth    *auth.Auth `json:"auth,omitempty"`
}

type Request struct {
	Method string     `json:"method"`
	Header []KeyValue `json:"header,omitempty"`
	Body   Body       `json:"body,omitempty"`
	URL    URL        `json:"url"`
}

type URL struct {
	Raw       string     `json:"raw"`
	Protocol  string     `json:"protocol"`
	Host      []string   `json:"host"`
	Path      []string   `json:"path"`
	Query     []KeyValue `json:"query,omitempty"`
	Variables []KeyValue `json:"variable"`
}

type Body struct {
	Mode       string     `json:"mode"`
	Raw        string     `json:"raw,omitempty"`
	FormData   []KeyValue `json:"formdata,omitempty"`
	URLEncoded []KeyValue `json:"urlencoded,omitempty"`
	Options    Options    `json:"options"`
}

type Options struct {
	Raw Raw `json:"raw"`
}

type Raw struct {
	Lang string `json:"language"`
}

type Environment struct {
	Values []KeyValue `json:"values"`
}

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
}
