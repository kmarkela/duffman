package output

import (
	"fmt"

	"github.com/kmarkela/duffman/internal/pcollection"
)

// Output types

type OutputType int

const (
	Brief OutputType = iota
	Vars
	Req
	Full
)

func PrintCol(t OutputType, c *pcollection.Collection) {
	printBr(c)

}

func printBr(c *pcollection.Collection) {
	if c.Env != nil {
		fmt.Printf("[*] Envoriment:\n")
		for _, v := range c.Env {
			fmt.Printf(" - %s: %s\n", v.Key, v.Value)
		}
	}

	if c.Variables != nil {
		fmt.Printf("[*] Variables:\n")
		for _, v := range c.Variables {
			fmt.Printf(" - %s: %s\n", v.Key, v.Value)
		}
	}

	fmt.Printf("[*] Req amount: %d\n", len(c.Requests))
}
