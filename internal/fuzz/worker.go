package fuzz

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

func startWorker(wg *sync.WaitGroup, wq <-chan workUnit, tr *http.Transport) {

	for wu := range wq {

		if !wu.parBody {

			wu.r.Parameters.Get[wu.param] = wu.word

			endpoint := fmt.Sprintf("%s?", wu.r.URL)
			for gk, gv := range wu.r.Parameters.Get {
				endpoint = fmt.Sprintf("%s%s=%s&", endpoint, gk, gv)
			}

			var r io.Reader = strings.NewReader(wu.r.Body)
			doRequest(endpoint, r, tr)

			continue
		}

	}

	wg.Done()

}

func doRequest(endpoint string, body io.Reader, tr *http.Transport) {}
