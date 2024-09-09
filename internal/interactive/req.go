package interactive

import (
	"fmt"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func buildString(r pcollection.Req) string {
	out := fmt.Sprintf("Method:: %s\n", r.Method)
	out += fmt.Sprintf("URL:: %s\n", r.URL)
	if len(r.Headers) > 0 {
		out += "HEADERS:: "

		for k, v := range r.Headers {
			out += fmt.Sprintf("%s: %s\n", k, v)
		}
	}

	return out

}
