package pcollection

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/pkg/jsonparser"
)

func buildReq(r *Request) Req {
	req := Req{}
	req.Headers = make(map[string]string)
	// parse headers
	for _, v := range r.Header {
		req.Headers[strings.ToLower(v.Key)] = v.Value
	}

	//remove get parameters
	req.URL = strings.Split(r.URL.Raw, "?")[0]

	// check schema
	if !strings.HasPrefix(req.URL, "http") {
		req.URL = fmt.Sprintf("http://%s", req.URL)
	}

	// parse params
	get := make(map[string]string)
	for _, v := range r.URL.Query {
		if v.Key == "" {
			continue
		}
		get[strings.ToLower(v.Key)] = v.Value
	}

	post, ct, body := parseBody(r.Body)

	req.Parameters.Get = get
	req.Parameters.Post = post
	req.Parameters.Path = parseVariables(r.URL.Variables)

	req.ContentType = ct
	req.Body = body
	req.Method = r.Method

	return req

}

func parseVariables(vars []KeyValue) map[string]string {

	rv := make(map[string]string)
	for _, v := range vars {
		rv[v.Key] = v.Value
	}

	return rv
}

func parseBody(b Body) (map[string]string, string, string) {
	params := make(map[string]string)

	if b.FormData != nil {
		for _, v := range b.FormData {
			params[v.Key] = v.Value
		}
		return params, "multipart/form-data", ""
	}

	if b.URLEncoded != nil {
		for _, v := range b.URLEncoded {
			params[v.Key] = v.Value
		}
		return params, "application/x-www-form-urlencoded", ""
	}

	switch b.Options.Raw.Lang {
	case "json":
		// TODO: log error
		params, _ := jsonparser.Unmarshal(b.Raw)
		return params, "application/json", ""
	// TODO: add xml parser
	// case "xml":
	default:
		return params, b.Options.Raw.Lang, b.Raw
	}

}
