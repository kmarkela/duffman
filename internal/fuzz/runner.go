package fuzz

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kmarkela/duffman/internal/output"
	"github.com/kmarkela/duffman/internal/pcollection"
)

type workUnit struct {
	r           pcollection.Req
	word, param string
	parBody     bool
}

// type WorkResults struct {
// 	endpoint, param, word, method string
// 	time                          time.Duration
// 	code                          int
// 	length                        int64
// 	err                           error
// }

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
	var wq = make(chan workUnit)
	var wr = make(chan output.Results)

	output.Header(col, len(wordlist))

	// consume the results
	go func() {
		var lt []output.Results
		for r := range wr {
			lt = append(lt, r)
			output.RenderTable(lt)
			// fmt.Printf("%s\t%s\t%s\t%s\t%d\t%s\t%d\n", r.method, r.endpoint, r.param, r.word, r.code, r.time, r.length)
		}
	}()

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

		for key := range v.Parameters.Get {
			distrWU(key, wordlist, v, rateLimiter, wq, false)

		}

		for key := range v.Parameters.Post {
			distrWU(key, wordlist, v, rateLimiter, wq, true)
		}
	}

	close(wq)
	wg.Wait()
}

func distrWU(key string, wordlist []string, r pcollection.Req, rl <-chan time.Time, wq chan workUnit, parBody bool) {

	for _, word := range wordlist {

		if rl != nil {
			<-rl // Wait for rate limit if provided
		}

		wq <- workUnit{
			r:       r,
			word:    word,
			param:   key,
			parBody: parBody,
		}
	}
}

func pwlist(filename string) ([]string, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// TODO: remove EOF
	lines := strings.Split(string(content), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("%s is empty", filename)
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

	}
}
