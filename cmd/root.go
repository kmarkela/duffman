package cmd

import "fmt"

const version = "v0.1.0"

func Execute() {
	// Get Cli Parameters
	var flags = &CliFlags{}
	flags.Parse()

	// print version and exit
	if flags.Version {
		fmt.Printf("DuffMan - version %s", version)
		return
	}

}
