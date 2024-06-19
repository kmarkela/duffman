package pcollection

import "strings"

func buildReq(r Request) (Req, error) {
	req := Req{}

	// parse headers
	for _, v := range r.Header {
		req.Headers[strings.ToLower(v.Key)] = v.Value
	}

	// parse params
	get := make(map[string]string)
	// post := make(map[string]string)

	for _, v := range r.URL.Query {
		get[strings.ToLower(v.Key)] = v.Value
	}

	//

}

func resolveLocalVar(v KeyValue) string {

}
