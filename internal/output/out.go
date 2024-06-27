package output

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func Header(col *pcollection.Collection, wl int) {

	var length int = 55

	fmt.Printf("%s#\n", strings.Repeat("#", length))

	f := (length - 7) / 2
	line := fmt.Sprintf("%sDuffMan%s", strings.Repeat(" ", f), strings.Repeat(" ", length-f-8))
	fmt.Printf("#%s#\n", line)

	line = fmt.Sprintf("# [*] Wordlist count: %d", wl)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	line = fmt.Sprintf("# [*] Amount of request: %d", len(col.Requests))
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	var r int
	for _, v := range col.Requests {
		r += len(v.Parameters.Get) + len(v.Parameters.Post)
	}

	line = fmt.Sprintf("# [*] Amount of parameters: %d", r)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	line = fmt.Sprintf("# [*] Total to fuzz: %d", r*wl)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	fmt.Printf("%s#\n", strings.Repeat("#", length))

}
