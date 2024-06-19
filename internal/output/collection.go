package output

import (
	"fmt"

	"github.com/kmarkela/duffman/internal/pcollection"
)

// Output types

type OutputType int

const (
	Brief OutputType = iota
	Req
	Full
)

func PrintCol(t OutputType, c *pcollection.Collection) {

	if c.Env != nil {
		fmt.Printf("[*] Envoriment:\n")
		for _, v := range c.Env {
			fmt.Printf("  - %s: %s\n", v.Key, v.Value)
		}
	}

	if c.Variables != nil {
		fmt.Printf("[*] Variables:\n")
		for _, v := range c.Variables {
			fmt.Printf("  - %s: %s\n", v.Key, v.Value)
		}
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
			for key, value := range v.Headers {
				fmt.Printf("      > %s: %s\n", key, value)
			}
		}

		if len(v.Parameters.Get) > 0 {
			fmt.Printf("    * Get Params:\n")
			for key, value := range v.Parameters.Get {
				fmt.Printf("      > %s: %s\n", key, value)
			}
		}

		if len(v.Parameters.Post) > 0 {
			fmt.Printf("    * Post Params:\n")
			for key, value := range v.Parameters.Post {
				fmt.Printf("      > %s: %s\n", key, value)
			}
		}
	}

}
