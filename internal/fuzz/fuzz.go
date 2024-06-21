package fuzz

import "github.com/kmarkela/duffman/internal/pcollection"

type Fuzzer struct{}

func New(workers, maxReq int,
	headers, excludeParam, parameters []string,
	fname, proxy string,
	verbose bool) (Fuzzer, error) {

	return Fuzzer{}, nil
}

func (f *Fuzzer) Run(col *pcollection.Collection) {}
