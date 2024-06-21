package fuzz

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kmarkela/duffman/internal/pcollection"
)

type Fuzzer struct {
	workers, maxReq           int
	excludeParams, parameters []string
	verbose                   bool
	headers                   map[string]string
	tr                        *http.Transport
}

func New(workers, maxReq int, headers, excludeParams, parameters []string, proxy string, verbose bool) (Fuzzer, error) {

	var fuzzer = Fuzzer{}

	// parse headers
	h, err := pheaders(headers)
	if err != nil {
		return fuzzer, err
	}

	// parse proxy
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return fuzzer, err
	}
	return Fuzzer{
		headers:       h,
		workers:       workers,
		maxReq:        maxReq,
		excludeParams: excludeParams,
		parameters:    parameters,
		verbose:       v,
		tr:            &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}, nil
}

func pheaders(headers []string) (map[string]string, error) {
	rh := make(map[string]string)

	for _, h := range headers {
		p := strings.Split(h, ":")
		if len(p) < 2 {
			return nil, fmt.Errorf("%s is wrong header format", h)
		}
		rh[strings.TrimSpace(p[0])] = p[1]
	}

	return rh, nil
}

func (f *Fuzzer) Run(col *pcollection.Collection, fname string) {

	// read wordlist
	wordlist, err := pwlist(filename)
	if err != nil {
		return nil, err
	}
}
