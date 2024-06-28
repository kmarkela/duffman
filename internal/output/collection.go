package output

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/internal/pcollection"
)

// Output types

type OutputType int

const (
	Brief OutputType = iota
	Req
	Full
)

func printKV(tab int, s string, lv []pcollection.KeyValue) {
	for _, v := range lv {
		fmt.Printf("%s%s %s: %s\n", strings.Repeat(" ", tab), s, v.Key, v.Value)
	}
}

func printMap(tab int, s string, lv map[string]string) {
	for k, v := range lv {
		fmt.Printf("%s%s %s: %s\n", strings.Repeat(" ", tab), s, k, v)
	}
}

func PrintCol(t OutputType, c *pcollection.Collection) {

	if c.Env != nil {
		fmt.Printf("[*] Envoriment:\n")
		printKV(2, "-", c.Env)

	}

	if c.Variables != nil {
		fmt.Printf("[*] Variables:\n")
		printKV(2, "-", c.Variables)

	}

	fmt.Printf("[*] Req amount: %d\n", len(c.Requests))

	if t == Brief {
		return
	}

	fmt.Printf("[*] Requests: \n")

	if t == Req {

		for _, v := range c.Requests {

			fmt.Printf("  - URL: %s\n", v.URL)
		}

		return

	}

	for _, v := range c.Requests {

		fmt.Printf("  - URL %s: \n", v.URL)
		fmt.Printf("    * Method: %s\n", v.Method)

		if v.ContentType != "" {
			fmt.Printf("    * Content-Type: %s\n", v.ContentType)
		}

		if v.Body != "" {
			fmt.Printf("    * Body: %s\n", v.Body)
		}

		if len(v.Headers) > 0 {
			fmt.Printf("    * Headers:\n")
			printMap(4, ">", v.Headers)

		}

		if len(v.Parameters.Get) > 0 {
			fmt.Printf("    * Get Params:\n")
			printMap(6, ">", v.Parameters.Get)

		}

		if len(v.Parameters.Post) > 0 {
			fmt.Printf("    * Post Params:\n")
			printMap(6, ">", v.Parameters.Post)
		}
	}

}
