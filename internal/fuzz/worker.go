package fuzz

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/kmarkela/duffman/internal/output"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/pkg/jsonparser"
)

type workUnit struct {
	r           pcollection.Req
	word, param string
	fuzzT       fuzzType
}

func startWorker(wg *sync.WaitGroup, wq <-chan workUnit, wr chan<- output.Results, tr *http.Transport) {

	for wu := range wq {

		result := output.Results{
			Param:  wu.param,
			Word:   wu.word,
			Method: wu.r.Method,
		}

		switch wu.fuzzT {
		case PATH:

			pathParam := make(map[string]string)
			for k, v := range wu.r.Parameters.Path {
				pathParam[k] = v
			}
			pathParam[wu.param] = wu.word

			endpoint := createEndpoint(wu.r.URL, wu.r.Parameters.Get, pathParam)
			result.Code, result.Length, result.Time, result.Err = doRequest(endpoint, strings.NewReader(wu.r.Body), wu, tr)
			result.Endpoint = strings.Split(endpoint, "?")[0]

			wr <- result

		case GET:

			getParam := make(map[string]string)
			for k, v := range wu.r.Parameters.Get {
				getParam[k] = v
			}
			getParam[wu.param] = wu.word

			endpoint := createEndpoint(wu.r.URL, getParam, wu.r.Parameters.Path)
			result.Code, result.Length, result.Time, result.Err = doRequest(endpoint, strings.NewReader(wu.r.Body), wu, tr)
			result.Endpoint = strings.Split(endpoint, "?")[0]

			wr <- result

		case POST:

			endpoint := createEndpoint(wu.r.URL, wu.r.Parameters.Get, wu.r.Parameters.Path)

			body, err := encodeBody(&wu)
			result.Code, result.Length, result.Time, result.Err = doRequest(endpoint, body, wu, tr)
			if err != nil {
				result.Err = err
			}

			result.Endpoint = strings.Split(endpoint, "?")[0]

			wr <- result
		}

	}

	wg.Done()
}

func createEndpoint(url string, getParam map[string]string, pathParam map[string]string) string {

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

func doRequest(endpoint string, body io.Reader, wu workUnit, tr *http.Transport) (int, int64, time.Duration, error) {

	req, err := http.NewRequest(wu.r.Method, endpoint, body)
	if err != nil {
		return 0, 0, 0, err
	}

	for k, v := range wu.r.Headers {
		req.Header.Set(k, v)
	}

	if wu.r.ContentType != "" {
		req.Header.Set("Content-Type", wu.r.ContentType)
	}

	client := &http.Client{Transport: tr}

	start := time.Now()

	res, err := client.Do(req)
	if err != nil {
		return 0, 0, 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, res.ContentLength, time.Duration(time.Since(start).Truncate(time.Millisecond)), nil

}

func encodeBody(wu *workUnit) (io.Reader, error) {

	// to restore param back
	postParam := make(map[string]string)
	for k, v := range wu.r.Parameters.Post {
		postParam[k] = v
	}

	postParam[wu.param] = wu.word

	// encode Form
	if strings.HasPrefix(wu.r.ContentType, "application/x-www-form-urlencoded") {
		form := url.Values{}
		for k, v := range postParam {
			form.Add(k, v)
		}
		return strings.NewReader(form.Encode()), nil
	}

	if strings.HasPrefix(wu.r.ContentType, "multipart/form-data") {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)
		for k, v := range postParam {
			writer.WriteField(k, v)
			// TODO: error handler
			// err := writer.WriteField(k, v)

		}
		wu.r.ContentType = writer.FormDataContentType()
		return &buf, nil
	}

	// encode json
	if wu.r.ContentType == "application/json" {
		b, _ := jsonparser.Marshal(postParam)
		return bytes.NewBuffer(b), nil
	}

	// unknown content type
	return strings.NewReader(wu.r.Body), fmt.Errorf("no encoder for: %s", wu.r.ContentType)
}
