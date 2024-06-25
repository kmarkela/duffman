package fuzz

import (
	"net/http"
	"sync"
)

func startWorker(wg *sync.WaitGroup, wq <-chan workUnit, tr *http.Transport) {

}
