package output

import "time"

type Results struct {
	Endpoint, Param, Word, Method string
	Time                          time.Duration
	Code                          int
	Length                        int64
	Err                           error
}
