package pcollection

import (
	"fmt"
	"strings"
)

func buildReq(r *Request) Req {
	req := Req{}
	req.Headers = make(map[string]string)
	// parse headers
	for _, v := range r.Header {
		req.Headers[strings.ToLower(v.Key)] = v.Value
	}

	// url
	req.URL = resolveURLVar(r.URL)

	//remove get parameters
	req.URL = strings.Split(req.URL, "?")[0]

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

	req.ContentType = ct
	req.Body = body
	req.Method = r.Method

	return req

}

func resolveURLVar(url URL) string {
	if url.Variables == nil {
		return url.Raw
	}

	u := url.Raw
	for _, v := range url.Variables {
		key := fmt.Sprintf(":%s", v.Key)
		u = strings.ReplaceAll(u, key, v.Value)
	}

	return u
}

func parseBody(b Body) (map[string]string, string, string) {
	params := make(map[string]string)

	if b.FormData != nil {
		for _, v := range b.FormData {
			params[v.Key] = v.Value
		}
		return params, "multipart/form-data; boundary=------border", ""
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
		params, _ := UnmarshalJSONBody(b.Raw)
		return params, "application/json", ""
	// TODO: add xml parser
	// case "xml":
	default:
		return params, b.Options.Raw.Lang, b.Raw
	}

}
