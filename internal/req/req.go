package req

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/kmarkela/duffman/internal/internalTypes"
	"github.com/kmarkela/duffman/internal/pcollection"
)

func CreateEndpoint(url string, getParam map[string]string, pathParam map[string]string) string {

	for k, v := range pathParam {
		key := fmt.Sprintf(":%s", k)
		url = strings.ReplaceAll(url, key, v)
	}

	endpoint := fmt.Sprintf("%s?", url)
	for gk, gv := range getParam {
		endpoint = fmt.Sprintf("%s%s=%s&", endpoint, gk, gv)
	}

	return endpoint
}

// TODO: Return instead of modify
func ResolveVars(env, vars []internalTypes.KeyValue, req *pcollection.Req) {

	allVars := append(vars, env...)

	for _, v := range allVars {

		vk := fmt.Sprintf("{{%s}}", v.Key)
		req.URL = strings.ReplaceAll(req.URL, vk, v.Value)
		req.Body = strings.ReplaceAll(req.Body, vk, v.Value)

		for hk, hv := range req.Headers {
			req.Headers[strings.ReplaceAll(hk, vk, v.Value)] = strings.ReplaceAll(hv, vk, v.Value)
		}

		for pk, pv := range req.Parameters.Get {
			req.Parameters.Get[strings.ReplaceAll(pk, vk, v.Value)] = strings.ReplaceAll(pv, vk, v.Value)
		}

		for pk, pv := range req.Parameters.Post {
			req.Parameters.Post[strings.ReplaceAll(pk, vk, v.Value)] = strings.ReplaceAll(pv, vk, v.Value)
		}

		for pk, pv := range req.Parameters.Path {
			req.Parameters.Path[strings.ReplaceAll(pk, vk, v.Value)] = strings.ReplaceAll(pv, vk, v.Value)
		}

	}
}

// Deep copy function for Req struct
func DeepCopyReq(original *pcollection.Req) *pcollection.Req {
	if original == nil {
		return nil
	}

	// Create a new Req object
	copy := &pcollection.Req{
		Method:      original.Method,
		URL:         original.URL,
		Headers:     make(map[string]string),
		Body:        original.Body,
		ContentType: original.ContentType,
		Parameters: pcollection.Parameters{
			Get:  make(map[string]string),
			Post: make(map[string]string),
			Path: make(map[string]string),
		},
	}

	// Copy map values for Headers
	for k, v := range original.Headers {
		copy.Headers[k] = v
	}

	// Copy map values for Parameters
	for k, v := range original.Parameters.Get {
		copy.Parameters.Get[k] = v
	}
	for k, v := range original.Parameters.Post {
		copy.Parameters.Post[k] = v
	}
	for k, v := range original.Parameters.Path {
		copy.Parameters.Path[k] = v
	}

	return copy
}

func DoRequest(endpoint string, body io.Reader, r pcollection.Req, tr *http.Transport) (int, int64, time.Duration, error) {

	start := time.Now()

	res, err := DoRequestFull(endpoint, body, r, tr)
	if err != nil {
		return 0, 0, 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, res.ContentLength, time.Duration(time.Since(start).Truncate(time.Millisecond)), nil

}

func DoRequestFull(endpoint string, body io.Reader, r pcollection.Req, tr *http.Transport) (*http.Response, error) {

	req, err := http.NewRequest(r.Method, endpoint, body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	if r.ContentType != "" {
		req.Header.Set("Content-Type", r.ContentType)
	}

	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
