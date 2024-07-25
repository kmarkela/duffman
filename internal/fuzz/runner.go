package fuzz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/kmarkela/duffman/internal/output"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/pkg/jsonparser"
)

type fuzzType int

const (
	POST fuzzType = iota + 1
	GET
	PATH
)

func (f *Fuzzer) Run(col *pcollection.Collection, fname string) {

	// read wordlist
	wordlist, err := pwlist(fname)
	if err != nil {
		log.Fatalf("cannot open wordlist: %s", err)
	}

	// Create rate limiter if maxReq > 0
	var rateLimiter <-chan time.Time
	if f.maxReq > 0 {
		// TODO:  rateLimiter: time.NewTicker(time.Second / time.Duration(maxReq)),
		rateLimiter = time.Tick(time.Second / time.Duration(f.maxReq))
	}

	var wg sync.WaitGroup
	var rg sync.WaitGroup
	var wq = make(chan workUnit)
	var wr = make(chan output.Results)

	output.Header(col, len(wordlist), f.responseT, f.blacklist)

	// consume the results
	go func() {
		var lt []output.Results
		var le []output.Results
		for r := range wr {
			if r.Err != nil {
				le = append(le, r)
				continue
			}
			if slices.Contains(f.blacklist, r.Code) {
				continue
			}

			if f.responseT > 0 && time.Duration(f.responseT)*time.Millisecond > r.Time {
				continue
			}

			lt = append(lt, r)
			output.RenderTable(lt)
		}

		if len(le) > 0 {
			output.RenderErrors(le)
		}
		rg.Done()
	}()
	rg.Add(1)

	// start workers
	for i := 0; i < f.workers; i++ {
		wg.Add(1)
		go startWorker(&wg, wq, wr, f.tr)
	}

	for _, v := range col.Requests {
		resolveVars(col.Env, col.Variables, &v)

		for hk, hv := range f.headers {
			v.Headers[hk] = hv
		}

		for key := range v.Parameters.Path {
			fmt.Println(v.Parameters.Path)
			fmt.Println(v.URL)
			distrWU(key, wordlist, v, rateLimiter, wq, PATH)

		}

		for key := range v.Parameters.Get {
			distrWU(key, wordlist, v, rateLimiter, wq, GET)

		}

		for key := range v.Parameters.Post {
			distrWU(key, wordlist, v, rateLimiter, wq, POST)
		}
	}

	close(wq)
	wg.Wait()

	close(wr)
	rg.Wait()

}

func distrWU(key string, wordlist []string, r pcollection.Req, rl <-chan time.Time, wq chan workUnit, ft fuzzType) {

	for _, word := range wordlist {

		if key == jsonparser.Slice {
			continue
		}

		if rl != nil {
			<-rl // Wait for rate limit if provided
		}

		wq <- workUnit{
			r:     r,
			word:  word,
			param: key,
			fuzzT: ft,
		}
	}
}

func pwlist(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func resolveVars(env, vars []pcollection.KeyValue, req *pcollection.Req) {

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
